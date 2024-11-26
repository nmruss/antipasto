package navigation

import(
	"testing"
)

func TestFolderValid(t *testing.T){
	//test the FolderValid function
	//this function should return true when the following is true:
	/*
		Any subfolder of the parent folder named in the format 'integer'x'integer' contains:
			-src/main.js
			-styles/main.css
			index.html

		If this is untrue return false

		Note: this function can ignore other folders intentionally, allowing for build flexibility
	*/
	validFolderPath := "../testData/test_banner_folder"
	status := FolderValid(validFolderPath)
	if(status != true){
		t.Fatalf(`FolderValid('../testData/test_banner_folder') returned false`)
	}
}
