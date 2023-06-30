package main

// #cgo LDFLAGS: -lpam
// #include <security/pam_appl.h>
// #include <security/pam_modules.h>
// void myCFunction(const char* str);
import "C"

import (
	"fmt"
	"os"
	"os/user"
	"unsafe"

	"github.com/godwinpinto/gatepass/client/util"
	"github.com/muesli/go-pam"
)

//export goAuthenticate
func goAuthenticate(handle *C.pam_handle_t, flags C.int, argv []string) C.int {
	//calling a C function from go
	//MyFunction("Hello from Go!")
	hdl := pam.Handle{Ptr: unsafe.Pointer(handle)}
	username, err := hdl.GetUser()
	if err != nil {
		return C.PAM_AUTH_ERR
	}
	configFileLocation, err := ReadUserConfig(username)
	if err != nil {
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

// main is for testing purposes only or deploying as a sshd ForceCommand option, the PAM module has to be built with:
// go build -buildmode=c-shared
func main() {

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Authentication allowed since failed to get username: ")
		//exit as succes
		return
	}

	configFileLocation, err := ReadUserConfig(currentUser.Username)
	if err != nil {
		fmt.Println("Authentication allowed No Passage configuration set for user: ", currentUser.Username)
		//exit as success
		return
	}
	authStatus := Controller(configFileLocation)

	if authStatus == util.AUTH_SUCCESS {
		//exit as failure and close session
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func MyFunction(message string) {
	cMessage := C.CString(message)
	C.myCFunction(cMessage)
}
