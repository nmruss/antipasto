package navigation

import (
	"testing"
)

func TestFolderValid(t *testing.T) {
	//Test FolderValid with a valid input
	//FolderValid tests a banner output folder for validity
	validFolderPath := "../testData/valid_banner_folder"
	status := FolderValid(validFolderPath)
	if status != true {
		t.Fatalf(`FolderValid('../testData/valid_banner_folder') failed`)
	}
}

func TestFolderInvalid(t *testing.T) {
	//Test FolderValid with an invalid input
	//FolderValid tests a banner output folder for validity
	invalidFolderPath := "../testData/invalid_banner_folder"
	status := FolderValid(invalidFolderPath)
	if status != false {
		t.Fatalf(`FolderValid('../testData/invalid_banner_folder') failed`)
	}
}
