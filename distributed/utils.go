package distributed

import (
	"errors"
	"fmt"
)

func FindDarylServer(identifier string) (string, error) {
	url, err := GetKey(fmt.Sprintf("daryl_%s", identifier))
	if err != nil {
		return "", err
	}
	return url, nil
}

func ListDarylServers() ([]string, error) {
	urls, err := ListPrefix("private_")
	if err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, errors.New("No daryl servers")
	}
	return urls, nil
}
