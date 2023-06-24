package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"workspace/types"
)

func TestFetchData(t *testing.T) {
	// Create a temporary output file for testing
	tempFile, err := ioutil.TempFile("", "test-output-*.json")
	if err != nil {
		t.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(tempFile.Name())

	// Fetch data
	FetchData(tempFile.Name())

	// Read the saved data from the output file
	savedData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal("Error reading saved data:", err)
	}

	// Unmarshal the saved data into a Data struct
	var data types.Data
	err = json.Unmarshal(savedData, &data)
	if err != nil {
		t.Fatal("Error unmarshaling saved data:", err)
	}

	// Verify the fetched data is not empty, above marshall was success shows data receive in correct format
	if data == (types.Data{}) {
		t.Errorf("Fetched data does not match expected data. Got %+v, want %+v", types.Data{}, data)
	}
}
