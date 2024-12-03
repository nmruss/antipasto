package validation

import (
	"testing"
)

func TestFolderValid(t *testing.T) {
	//Test FolderValid with a valid input
	//FolderValid tests a banner output folder for validity == true
	validFolderPath := "../testdata/valid_banner_folder"
	status := FolderValid(validFolderPath)
	if status != true {
		t.Fatalf(`FolderValid('../testData/valid_banner_folder') failed`)
	}
}

func TestFolderInvalid(t *testing.T) {
	//Test FolderValid with an invalid input
	//FolderValid tests a banner output folder for validity == false
	invalidFolderPath := "../testdata/invalid_banner_folder"
	status := FolderValid(invalidFolderPath)
	if status != false {
		t.Fatalf(`FolderValid('../testData/invalid_banner_folder') failed`)
	}
}

func TestProjectInvalid(t *testing.T) {
	//Test ProjectValid with an invalid input
	//ProjectValid tests a project folder structure for validity == false
	invalidProjectPath := "../testdata/testInvalidProjectFolder"
	status := ProjectValid(invalidProjectPath)
	if status != false {
		t.Fatalf(`FolderValid('../testData/testInvalidProjectFolder') failed`)
	}
}

func TestProjectValid(t *testing.T) {
	//Test ProjectValid with an invalid input
	//ProjectValid tests a project folder structure for validity == false
	validProjectPath := "../testdata/testValidProjectFolder"
	status := ProjectValid(validProjectPath)
	if status != true {
		t.Fatalf(`FolderValid('../testData/testInvalidProjectFolder') failed`)
	}
}
