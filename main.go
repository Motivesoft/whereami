package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// curl https://neutrinoapi.net/ip-info \
// --header "User-ID: <your-user-id>" \
// --header "API-Key: <your-api-key>" \
// --data-urlencode "ip=1.1.1.1" \
// --data-urlencode "reverse-lookup=false"
func main() {
	apiURL := "https://neutrinoapi.net/ip-info"

	if len(os.Args) < 2 {
		fmt.Println("Pass IP address as command line parameter")
		return
	}

	// Create URL with query parameters
	baseURL, err := url.Parse(apiURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	params := url.Values{}
	params.Add("ip", os.Args[1])
	params.Add("reverse-lookup", "false")
	baseURL.RawQuery = params.Encode()

	// Create a new request
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Read user-specific header values from a dotfile
	headers, err := readHeadersFromDotfile(".env")
	if err != nil {
		fmt.Println("Failed to read headers from .env:", err)
		return
	}

	// Put all values read from the dotfile as header entries
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	Print(string(body))
}

func Print(input string) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(input), "", "    "); err != nil {
		fmt.Println("Error prettifying JSON:", err)
		return
	}
	fmt.Println(prettyJSON.String())
}

func readHeadersFromDotfile(filename string) (map[string]string, error) {
	headers := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return headers, nil
}
