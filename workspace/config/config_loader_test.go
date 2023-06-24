package config

import (
	"io/ioutil"
	"os"
	"testing"
	"workspace/types"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file for testing
	tempFile, err := ioutil.TempFile("", "test-config.json")
	if err != nil {
		t.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test configuration to the config file
	configData := []byte(`{
"frequency": 1000000000,
"output_file": "test-output.txt"
}`)
	err = ioutil.WriteFile(tempFile.Name(), configData, 0644)
	if err != nil {
		t.Fatal("Error writing test configuration:", err)
	}

	// Load the test configuration
	config, err := LoadConfig(tempFile.Name())
	if err != nil {
		t.Fatal("Error loading test configuration:", err)
	}

	// Verify the loaded configuration matches the expected configuration
	expectedConfig := &types.Config{
		Frequency:  1000000000,
		OutputFile: "test-output.txt",
	}
	if config.Frequency != expectedConfig.Frequency || config.OutputFile != expectedConfig.OutputFile {
		t.Errorf("Loaded configuration does not match expected configuration. Got %+v, want %+v", config, expectedConfig)
	}
}
