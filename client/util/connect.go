package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/passageidentity/passage-go"
)

type AuthStatus int

type ConnectResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

const (
	AUTH_SUCCESS AuthStatus = iota
	AUTH_FAILURE
	AUTH_TIMEOUT
)

type ConnectRequest struct {
	UserID string `json:"user_id"`
}

func Authenticate(connectBean ConnectRequest, appID string) (status AuthStatus) {
	var response string
	var err error
	response, err = getData(connectBean)
	if response == "" || err != nil {
		return AUTH_TIMEOUT
	}
	var connectRes ConnectResponse

	err1 := json.Unmarshal([]byte(response), &connectRes)
	if err1 != nil {
		fmt.Println("Error:", err)
		return
	}

	return PassageAuthentication(connectRes.Token, appID, connectBean.UserID)
}

func getData(connectBean ConnectRequest) (response string, err error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	jsonPayload, err := json.Marshal(connectBean)
	if err != nil {
		fmt.Printf("Error encoding JSON payload: %v\n", err)
		return
	}

	resp, err := client.Post("https://connect.coauth.dev/connect", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("")
		log.Fatal(err)
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
	return string(body), nil
}

func PassageAuthentication(authToken string, appID string, userIDFromFile string) (authStatus AuthStatus) {
	psg, err2 := passage.New(appID, nil)
	if err2 != nil {
		fmt.Print(err2)
	}
	userID, success := psg.ValidateAuthToken(authToken)
	if userIDFromFile != userID {
		return AUTH_FAILURE
	}
	if success {
		return AUTH_SUCCESS
	} else {
		return AUTH_FAILURE
	}
}
