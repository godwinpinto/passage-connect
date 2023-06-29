package main

import (
	"fmt"
	"os/user"
	"path/filepath"
)

const configFile = ".connect"

func homeDir(username string) (string, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return "", err
	}

	return u.HomeDir, nil
}

func ReadUserConfig(username string) (string, error) {
	hd, err := homeDir(username)
	if err != nil {
		return "", err
	}
	fmt.Println("Home directory: ", hd)
	fmt.Println(filepath.Join(hd, configFile))
	fmt.Println("-------------")
	return filepath.Join(hd, configFile), nil
}
