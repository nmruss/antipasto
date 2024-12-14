package filemgmt

import (
	"testing"
)

// Tests if the UpdateCSSTokenList function is
// updating properties in the correct place
func TestUpdateCSSTokenList(t *testing.T) {
	//under filemgmt
	cssPath := "../testdata/filemgmt/test.css"
	properties := []CSSToken{
		{Type: "HASH", Value: "#copy1"},
		{Type: "CHAR", Value: "{"},
		{Type: "S", Value: "\n"},
		{Type: "S", Value: "      "},
		{Type: "IDENT", Value: "top"},
		{Type: "CHAR", Value: ":"},
		{Type: "DIMENSION", Value: "10px"},
		{Type: "CHAR", Value: ";"},
		{Type: "S", Value: "\n"},
		{Type: "CHAR", Value: "}"},
	}

	var inserts []CSSPropertyInsert

	inserts = append(inserts, CSSPropertyInsert{
		ParentName:   "#copy1",
		PropertyName: "top",
		Value:        "44px",
		Type:         "DIMENSION",
	})

	UpdateCSSTokenList(&cssPath, &properties, &inserts)

	for _, property := range properties {
		if property.Type == "DIMENSION" && property.Value != "44px" {
			t.Fatalf("UpdateCSSTokenList() failed to update DIMENSION token on #copy1")
		}
	}
}

// Tests if the UpdateCSSTokenList() function
// is inserting new tokens properly
func TestUpdateCSSTokenListAddition(t *testing.T) {
	cssPath := "../testdata/filemgmt/test.css"
	properties := []CSSToken{
		{Type: "HASH", Value: "#copy1"},
		{Type: "CHAR", Value: "{"},
		{Type: "S", Value: "\n"},
		{Type: "S", Value: "      "},
		{Type: "IDENT", Value: "top"},
		{Type: "CHAR", Value: ":"},
		{Type: "DIMENSION", Value: "10px"},
		{Type: "CHAR", Value: ";"},
		{Type: "S", Value: "\n"},
		{Type: "CHAR", Value: "}"},
	}

	var inserts []CSSPropertyInsert

	inserts = append(inserts, CSSPropertyInsert{
		ParentName:   "#copy1",
		PropertyName: "left",
		Value:        "20px",
		Type:         "DIMENSION",
	})

	UpdateCSSTokenList(&cssPath, &properties, &inserts)

	if len(properties) < 13 {
		t.Fatalf("UpdateCSSTokenList() did not insert the right number of properties")
	}

	if properties[8].Type != "IDENT" || properties[8].Value != inserts[0].PropertyName {
		t.Fatalf("UpdateCSSTokenList() did not insert a new IDENT 'left' properly")
	}

	if properties[9].Type != "CHAR" || properties[9].Value != ":" {
		t.Fatalf("UpdateCSSTokenList() did not insert a new CHAR ':' properly")
	}

	if properties[10].Type != "DIMENSION" || properties[10].Value != inserts[0].Value {
		t.Fatalf("UpdateCSSTokenList() did not insert a new DIMENSION properly")
	}

	if properties[11].Type != "CHAR" || properties[11].Value != ";" {
		t.Fatalf("UpdateCSSTokenList() did not insert a new CHAR ';' properly")
	}
}

// Tests if CSS parsing functionality is valid
func TestTokenizeCSSFromFile(t *testing.T) {
	cssPath := "../testdata/filemgmt/test.css"
	TokenizeCSSFromFile(&cssPath)
}
