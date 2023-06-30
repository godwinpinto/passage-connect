package util

//currently this file is not used and is developed for future development
import (
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
)

func GetInstanceID() (string, error) {
	data, err := os.ReadFile("/var/lib/cloud/data/instance-id")
	if err != nil {
		return "", err
	}

	instanceID := strings.TrimSpace(string(data))
	return instanceID, nil
}

func GetUUID() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuidObj.String(), nil
}

func GetIPAddresses() ([]string, error) {
	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	// Retrieve IP addresses associated with the hostname
	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return nil, err
	}

	// Filter and collect IPv4 addresses
	var ipv4Addresses []string
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Addresses = append(ipv4Addresses, ipv4.String())
		}
	}

	return ipv4Addresses, nil
}
