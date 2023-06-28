package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/magiconair/properties"
	passage "github.com/passageidentity/passage-go"
)

type UserData struct {
	data      string
	ready     bool
	mutex     sync.Mutex
	condition *sync.Cond
}

type AuthRequest struct {
	Token string `json:"token"`
}

type AuthResponse struct {
	Message string `json:"message"`
}

type ConnectRequest struct {
	UserID string `json:"user_id"`
}

type ConnectResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type ValidateRequest struct {
	Token string `json:"token"`
}

type SettingsRequest struct {
	Setting1 string `json:"setting1"`
	Setting2 string `json:"setting2"`
}

var usersData sync.Map // Map to store user-specific data

var passageClient *passage.App

func main() {
	var passageAppId string
	var passageApiKey string

	filePath := "config.properties"
	data, err := readPropertiesFile(filePath)
	if err != nil {
		fmt.Println("Error reading config.properties file:", err)
		fmt.Println("Fall back to environment variables", err)
		passageAppId = os.Getenv("PASSAGE_APP_ID")
		passageApiKey = os.Getenv("PASSAGE_API_KEY")
	} else {
		passageAppId = data["PASSAGE_APP_ID"]
		passageApiKey = data["PASSAGE_API_KEY"]
	}

	if passageAppId == "" {
		log.Fatal("PASSAGE_APP_ID environment variable not set")
		return
	}
	if passageApiKey == "" {
		log.Fatal("PASSAGE_API_KEY environment variable not set")
		return
	}

	passageClient, _ = passage.New(passageAppId, &passage.Config{
		APIKey:     passageApiKey,
		HeaderAuth: true,
	})

	//fs := http.FileServer(http.Dir("../web/build"))
	fs := http.FileServer(http.Dir("web"))

	//	http.Handle("/", fs)
	//	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/dashboard", http.StripPrefix("/dashboard", fs))
	http.Handle("/app/", http.StripPrefix("/app/", fs))
	http.Handle("/", http.StripPrefix("/", fs))

	//http.Handle("/dashboard", fs)

	http.HandleFunc("/connect", handleConnect)
	http.HandleFunc("/login", handleLogin)

	log.Println("Server listening on port 443...")
	//log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServeTLS(":443", "../cert/cert.pem", "../cert/key.pem", nil))
}

func handleConnect(w http.ResponseWriter, r *http.Request) {
	var connectReq ConnectRequest
	err := json.NewDecoder(r.Body).Decode(&connectReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	userID := connectReq.UserID
	if strings.TrimSpace(userID) == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	fmt.Println(userID)
	_, err = passageClient.GetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	userData, found := getUserData(userID)
	if !found {
		userData = createUserData(userID)
	}

	done := make(chan struct{}) // Channel to signal completion

	go func() {
		userData.mutex.Lock()
		defer userData.mutex.Unlock()

		if !userData.ready {
			fmt.Println("Waiting for data to be set")
			userData.condition.Wait()
			fmt.Println("Data set complete")
		}
		fmt.Println("closing func")

		done <- struct{}{} // Signal completion
	}()

	select {
	case <-done:
		if userData.data == "" {
			fmt.Println("Data is set")
			connectRes := ConnectResponse{
				Token: userData.data,
			}
			jsonData, err := json.Marshal(connectRes)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error: %v", err)

			} else {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonData)
			}
		}
	case <-time.After(5 * time.Second):
		userData.mutex.Lock()
		userData.data = ""
		userData.ready = true
		userData.condition.Broadcast()
		userData.mutex.Unlock()

		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Request has timed out")
	}

	clearUserData(userID)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var authReq AuthRequest
	err := json.NewDecoder(r.Body).Decode(&authReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	token := authReq.Token
	userID, err := passageClient.AuthenticateRequest(r)
	if err != nil {
		// Authentication failed!
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userData, found := getUserData(userID)
	if !found {
		//userData = createUserData(userID)
	} else {
		userData.mutex.Lock()
		userData.data = token
		userData.ready = true
		userData.mutex.Unlock()
		userData.condition.Broadcast() // Notify that data is set
	}

	//	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//	w.Write(jsonData)
	fmt.Fprintf(w, "No active shell sessions found.")
}

func getUserData(userID string) (*UserData, bool) {
	data, found := usersData.Load(userID)
	if found {
		return data.(*UserData), true
	}
	return nil, false
}

func createUserData(userID string) *UserData {
	userData := &UserData{}
	userData.condition = sync.NewCond(&userData.mutex)
	usersData.Store(userID, userData)
	return userData
}

func clearUserData(userID string) {
	usersData.Delete(userID)
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
