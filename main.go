package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func getStatement(accountNumber, token string) error {
	// Вычисляем вчерашнюю дату в UTC
	yesterday := time.Now().UTC().AddDate(0, 0, -1)
	fromDate := yesterday.Format("2006-01-02T15:04:05Z")

	url := fmt.Sprintf("https://business.tbank.ru/openapi/api/v1/statement?accountNumber=%s&withBalances=true&from=%s", accountNumber, fromDate)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}
	fmt.Println(string(body))
	return nil
}

func main() {
	token := os.Getenv("TB_TOKEN")
	if token == "" {
		fmt.Println("Error: TB_TOKEN environment variable is not set")
		return
	}

	accountNumber := "40817810500000586627"
	if err := getStatement(accountNumber, token); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
