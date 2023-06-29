package main

// #cgo LDFLAGS: -lpam
// #include <security/pam_appl.h>
// #include <security/pam_modules.h>
import "C"

import (
	"strings"
	"time"
	"unsafe"

	"github.com/godwinpinto/gatepass/client/util"
	"github.com/muesli/go-pam"
	log "github.com/sirupsen/logrus"
)

var (
	logLevel = log.InfoLevel
)

func logError(args ...interface{}) {
	log.Error(args...)
	time.Sleep(time.Second)
}

func logErrorf(s string, args ...interface{}) {
	log.Errorf(s, args...)
	time.Sleep(time.Second)
}

//export goAuthenticate
func goAuthenticate(handle *C.pam_handle_t, flags C.int, argv []string) C.int {
	for _, arg := range argv {
		if strings.ToLower(arg) == "debug" {
			logLevel = log.DebugLevel
		}
	}
	log.SetLevel(logLevel)
	log.Debugf("argv: %+v", argv)

	hdl := pam.Handle{Ptr: unsafe.Pointer(handle)}
	username, err := hdl.GetUser()
	if err != nil {
		return C.PAM_AUTH_ERR
	}
	configFileLocation, err := ReadUserConfig(username)
	if err != nil {
		logError(err)
		return C.PAM_SUCCESS
	}
	/*
		commented for now since you dont want to lock out other users
		if err != nil {
			logError(err)
			switch err.(type) {
			case user.UnknownUserError:
				return C.PAM_USER_UNKNOWN
			default:
				return C.PAM_AUTHINFO_UNAVAIL
			}
		} */

	if Controller(configFileLocation) == util.AUTH_SUCCESS {
		return C.PAM_SUCCESS
	}
	return C.PAM_AUTH_ERR
}

//export setCred
func setCred(handle *C.pam_handle_t, flags C.int, argv []string) C.int {
	return C.PAM_SUCCESS
}

// main is for testing purposes only, the PAM module has to be built with:
// go build -buildmode=c-shared
func main() {
}
