package input

import "fmt"

func GetUrl() (string, error) {
	fmt.Println("Enter A URL: ")

	var url string

	_, err := fmt.Scan(&url)

	if err != nil {
		return "", err
	}

	return url, nil
}
