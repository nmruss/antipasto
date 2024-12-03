package validation

import (
	"testing"
)

func TestIsFolderValid(t *testing.T) {
	//Test FolderValid with a valid input
	//FolderValid tests a banner output folder for validity == true
	validFolderPath := "../testdata/valid_banner_folder"
	invalidFolderPath := "../testdata/invalid_banner_folder"
	validFolderStatus := IsFolderValid(validFolderPath)
	invalidFolderStatus := IsFolderValid(invalidFolderPath)

	if validFolderStatus == false {
		t.Fatalf(`FolderValid('../testData/valid_banner_folder') failed`)
	}

	if invalidFolderStatus == true {
		t.Fatalf(`FolderValid('../testData/invalid_banner_folder') failed`)
	}
}

func TestIsProjectValid(t *testing.T) {
	//Test ProjectValid with both valid, and invalid input
	//ProjectValid tests a project folder structure for validity == false
	invalidProjectPath := "../testdata/testInvalidProjectFolder"
	validProjectPath := "../testdata/testValidProjectFolder"

	invalidStatus := IsProjectValid(invalidProjectPath)
	validStatus := IsProjectValid(validProjectPath)

	if invalidStatus == true {
		t.Fatalf(`FolderValid('../testData/testInvalidProjectFolder') failed`)
	}

	if validStatus == false {
		t.Fatalf(`FolderValid('../testData/testValidProjectFolder') failed`)
	}
}
