package filemgmt

import (
	"testing"
)

func TestUpdateCSSTokenList(t *testing.T) {
	//Tests if the WriteCSS function is writing valid CSS to the test files
	//under filemgmt
	cssPath := "../testdata/filemgmt/test.css"
	properties := []CSSToken{
		{Type: "HASH", Value: "#copy1"},
		{Type: "CHAR", Value: "{"},
		{Type: "S", Value: "      /n"},
		{Type: "HASH", Value: "      /n"},
		{Type: "IDENT", Value: "top"},
		{Type: "CHAR", Value: ":"},
		{Type: "DIMENSION", Value: "10px"},
		{Type: "CHAR", Value: ";"},
	}

	inserts := []CSSProperyInsert{
		{ParentName: "#copy1", PropertyName: "top", Value: "44px"},
	}

	UpdateCSSTokenList(&cssPath, &properties, &inserts)

	for _, property := range properties {
		if property.Type == "DIMENSION" && property.Value != "44px" {
			t.Fatalf("UpdateCSSTokenList() failed to update DIMENSION token on #copy1")
		}
	}
}

// Tests if CSS parsing functionality is valid
func TestTokenizeCSSFromFile(t *testing.T) {
	cssPath := "../testdata/filemgmt/test.css"
	TokenizeCSSFromFile(&cssPath)
}
