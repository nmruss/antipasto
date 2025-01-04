package configuration

import (
	"testing"
)

// Tests if ParseConfigurationFile
// Returns a correct APConfig Object
func TestParseConfigurationFile(t *testing.T) {
	configurationObj := ParseConfigurationFile("../testdata/test.apconfig")
	expectedConfigurationObject := APConfig{}
	expectedConfigurationObject.DefaultHTML = []string{`<!DOCTYPE html> <html lang="en"><html>` + "\n"}
	expectedConfigurationObject.DefaultJS = []string{`function main(){ console.log('hello') }` + "\n"}
	expectedConfigurationObject.DefaultStyling = []string{`.div { position: absolute; }` + "\n"}

	for i := range configurationObj.DefaultHTML {
		if configurationObj.DefaultHTML[i] != expectedConfigurationObject.DefaultHTML[i] {
			t.Fatalf(`ParseConfigurationFile filled DefaultHTML property incorrectly. Expected:%s Got:%s`, expectedConfigurationObject.DefaultHTML[i], configurationObj.DefaultHTML[i])
		}
	}

	for i := range configurationObj.DefaultJS {
		if configurationObj.DefaultJS[i] != expectedConfigurationObject.DefaultJS[i] {
			t.Fatalf(`ParseConfigurationFile filled DefaultHTML property incorrectly. Expected:%s Got:%s`, expectedConfigurationObject.DefaultJS[i], configurationObj.DefaultJS[i])
		}
	}

	for i := range configurationObj.DefaultStyling {
		if configurationObj.DefaultStyling[i] != expectedConfigurationObject.DefaultStyling[i] {
			t.Fatalf(`ParseConfigurationFile filled DefaultHTML property incorrectly. Expected:%s Got:%s`, expectedConfigurationObject.DefaultStyling[i], configurationObj.DefaultStyling[i])
		}
	}
}
