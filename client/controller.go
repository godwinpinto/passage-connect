package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/common-nighthawk/go-figure"
	util "github.com/godwinpinto/gatepass/client/util"
	"github.com/magiconair/properties"
)

func Controller(propertiesFilePath string) util.AuthStatus {
	var passageAppId string
	var passageUserId string

	data, err := readPropertiesFile(propertiesFilePath)
	if err != nil {
		fmt.Println("Error reading config.properties file:", err)
		fmt.Println("Fall back to environment variables", err)
		passageAppId = os.Getenv("PASSAGE_APP_ID")
		passageUserId = os.Getenv("PASSAGE_USER_ID")
	} else {
		passageAppId = data["PASSAGE_APP_ID"]
		passageUserId = data["PASSAGE_USER_ID"]
	}

	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get current user: %v\n", err)
	}
	fmt.Printf("Welcome %v", currentUser.Name)
	fmt.Println("")
	fmt.Println("")
	myFigure := figure.NewFigure("Passage Connect Login", "", true)
	fmt.Println("")
	myFigure.Print()
	fmt.Println("")
	fmt.Println("Visit https://connect.coauth.dev to authenticate this login.")
	fmt.Println("")
	connectBean := util.ConnectRequest{
		UserID: passageUserId,
	}
	cancelCounter := util.NewCancelCounter(30)
	go util.CountdownProgressBar(cancelCounter)
	authStatus := util.Authenticate(connectBean, passageAppId)
	cancelCounter.Cancel()
	fmt.Println("")
	fmt.Println("")

	if authStatus == util.AUTH_SUCCESS {
		myFigure := figure.NewFigure("WELCOME", "", true)
		fmt.Println("")
		myFigure.Print()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("You are now logged in!!!")
	} else if authStatus == util.AUTH_TIMEOUT {
		fmt.Println("Authentication timeout!!! Please login again")
	} else {
		fmt.Println("Authentication denied")
	}
	return authStatus
}

func readPropertiesFile(filePath string) (map[string]string, error) {
	p, err := properties.LoadFile(filePath, properties.UTF8)
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)

	// Iterate over the properties and store them in the map
	for _, key := range p.Keys() {
		value := p.GetString(key, "")
		data[key] = value
	}

	return data, nil
}
