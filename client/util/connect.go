package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/passageidentity/passage-go"
)

type AuthStatus int

const (
	AUTH_SUCCESS AuthStatus = iota
	AUTH_FAILURE
	AUTH_TIMEOUT
)

type AuthBean struct {
	UserID  string
	MacCode string
	IPAddr  string
}

func Authenticate(authBean AuthBean) (status AuthStatus) {
	attempts := 2
	var response string
	var err error
	for i := 0; i < attempts; i++ {
		response, err = getData()
		if err == nil {
			break
		}
	}
	if response == "" {
		return AUTH_TIMEOUT
	}
	fmt.Println(response)
	return passageAuthentication(response)
}

func getData() (response string, err error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("HTTP Status is incorrect")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	return string(body), nil
}

func passageAuthentication(authToken string) (authStatus AuthStatus) {
	psg, _ := passage.New("<PASSAGE_APP_ID>", nil)
	_, err := psg.ValidateAuthToken(authToken)
	if err {
		return AUTH_FAILURE
	}
	return AUTH_SUCCESS

}
