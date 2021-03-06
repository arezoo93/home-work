package input

import "fmt"

func GetUrl() (string, error) {
	var url string
	fmt.Println("Enter A URL: ")

	_, err := fmt.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}
