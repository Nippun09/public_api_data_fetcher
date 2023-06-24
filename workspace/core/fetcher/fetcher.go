package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"workspace/types"
)

func FetchData(outputfile string) {
	// Fetch data from API
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Parse response body into Data struct
	var data types.Data
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		return
	}

	// Write data to output file
	file, err := os.OpenFile(outputfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding data to JSON:", err)
		return
	}

	fmt.Println("Data saved to", outputfile)
}
