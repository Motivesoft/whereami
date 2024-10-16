package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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

	// Add headers
	// TODO READ THESE FROM A DOTFILE CALLED .env
	req.Header.Add("User-ID", "")
	req.Header.Add("API-Key", "")

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
