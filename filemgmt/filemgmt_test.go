package filemgmt

import "testing"

func TestWriteCSS(t *testing.T) {
	//Tests if the WriteCSS function is writing valid CSS to the test files
	//under filemgmt
	cssPath := "../testdata/filemgmt/test.css"
	properties := []Property{{name: "top", value: "10"}}

	WriteCSS(&cssPath, "image1", properties)
}

func TestParseCSSFromFile(t *testing.T) {
	//Tests if CSS parsing functionality is valid

	cssPath := "../testdata/filemgmt/test.css"
	ParseCSSFromFile(&cssPath)
}
