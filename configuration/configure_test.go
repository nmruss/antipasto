package configuration

import (
	"testing"
)

// Tests if ParseConfigurationFile
// Returns a correct APConfig Object
func TestParseConfigurationFile(t *testing.T) {
	configurationObj := ParseConfigurationFile("../testdata/test.apconfig")
	expectedConfigurationObject := APConfig{}
	expectedConfigurationObject.DefaultHTML = `<!DOCTYPE html> <html lang="en"><html>`
	expectedConfigurationObject.DefaultJS = `function main(){ console.log('hello') }`
	expectedConfigurationObject.DefaultStyling = `.div { position: absolute; }`

	if configurationObj.DefaultHTML != expectedConfigurationObject.DefaultHTML {
		t.Fatalf(`ParseConfigurationFile filled DefaultHTML property incorrectly. Expected: %s; Got: %s`, expectedConfigurationObject.DefaultHTML, configurationObj.DefaultHTML)
	}

	if configurationObj.DefaultJS != expectedConfigurationObject.DefaultJS {
		t.Fatalf(`ParseConfigurationFile filled DefaultJS property incorrectly. Expected: %s; Got: %s`, expectedConfigurationObject.DefaultJS, configurationObj.DefaultJS)
	}

	if configurationObj.DefaultStyling != expectedConfigurationObject.DefaultStyling {
		t.Fatalf(`ParseConfigurationFile filled DefaultStyling property incorrectly. Expected: %s; Got: %s`, expectedConfigurationObject.DefaultStyling, configurationObj.DefaultStyling)
	}
}
