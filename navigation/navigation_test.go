package navigation

import(
	"testing"
)

func TestFolderValid(t *testing.T){
	//Test FolderValid with a valid input
	validFolderPath := "../testData/valid_banner_folder"
	status := FolderValid(validFolderPath)
	if(status != true){
		t.Fatalf(`FolderValid('../testData/valid_banner_folder') failed`)
	}
}

func TestFolderInvalid(t *testing.T){
	//Test FodlerValid with an invalid input
	invalidFolderPath := "../testData/invalid_banner_folder"
	status := FolderValid(invalidFolderPath)
	if(status != false){
		t.Fatalf(`FolderValid('../testData/invalid_banner_folder') failed`)
	}
}
