package configuration

import (
	"testing"
)

// Tests if ParseConfigurationFile returns a valid APConfig Object
func TestParseConfigurationFile(t *testing.T) {
	configurationObj := ParseConfigurationFile("../testdata/test.apconfig")
	expectedConfigurationObject := APConfig{}
	expectedConfigurationObject.DefaultHTML = []string{`<!DOCTYPE html> <html lang="en"><html>` + "\n"}
	expectedConfigurationObject.DefaultJS = []string{`function main(){ console.log('hello') }` + "\n"}
	expectedConfigurationObject.DefaultStyling = []string{`.div { position: absolute; }` + "\n"}
	expectedConfigurationObject.Size = []int{300, 250}

	for i := range configurationObj.DefaultHTML {
		if configurationObj.DefaultHTML[i] != expectedConfigurationObject.DefaultHTML[i] {
			t.Fatalf(`ParseConfigurationFile filled DefaultHTML property incorrectly. Expected:%s Got:%s`, expectedConfigurationObject.DefaultHTML[i], configurationObj.DefaultHTML[i])
		}
	}

	for i := range configurationObj.DefaultJS {
		if configurationObj.DefaultJS[i] != expectedConfigurationObject.DefaultJS[i] {
			t.Fatalf(`ParseConfigurationFile filled DefaultJS property incorrectly. Expected:%s Got:%s`, expectedConfigurationObject.DefaultJS[i], configurationObj.DefaultJS[i])
		}
	}

	for i := range configurationObj.DefaultStyling {
		if configurationObj.DefaultStyling[i] != expectedConfigurationObject.DefaultStyling[i] {
			t.Fatalf(`ParseConfigurationFile filled DefaultStyling property incorrectly. Expected:%s Got:%s`, expectedConfigurationObject.DefaultStyling[i], configurationObj.DefaultStyling[i])
		}
	}

	for i := range configurationObj.Size {
		if configurationObj.Size[i] != expectedConfigurationObject.Size[i] {
			t.Fatalf(`ParseConfigurationFile filled Size property incorrectly. Expected:%d Got:%d`, expectedConfigurationObject.Size[i], configurationObj.Size[i])
		}
	}
}
