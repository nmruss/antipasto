package filemgmt

import (
	"testing"
)

// Tests if the UpdateCSSTokenList function is
// updating properties in the correct place
func TestUpdateCSSTokenList_PropertyUpdate(t *testing.T) {
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
func TestUpdateCSSTokenList_PropertyAdd(t *testing.T) {
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
		{Type: "HASH", Value: "#copy3"},
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
		Value:        "10px",
		Type:         "DIMENSION",
	})

	UpdateCSSTokenList(&cssPath, &properties, &inserts)

	if len(properties) < 13 {
		t.Fatalf("UpdateCSSTokenList() did not insert a new property correctly")
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

// Tests if a CSSPropertyInsert that does not have a parent
// in the current list of selectors is being added properly
func TestUpdateCSSTokenList_SelectorAdd(t *testing.T) {
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
		{Type: "HASH", Value: "#copy2"},
		{Type: "CHAR", Value: "{"},
		{Type: "S", Value: "\n"},
		{Type: "S", Value: "      "},
		{Type: "IDENT", Value: "left"},
		{Type: "CHAR", Value: ":"},
		{Type: "DIMENSION", Value: "30px"},
		{Type: "CHAR", Value: ";"},
		{Type: "S", Value: "\n"},
		{Type: "CHAR", Value: "}"},
	}

	var inserts []CSSPropertyInsert

	inserts = append(inserts, CSSPropertyInsert{
		ParentName:   "#newCopyID",
		PropertyName: "top",
		Value:        "20px",
		Type:         "DIMENSION",
	})

	UpdateCSSTokenList(&cssPath, &properties, &inserts)

	if len(properties) < 13 {
		t.Fatalf("UpdateCSSTokenList() did not insert a new selector properly")
	}

	if properties[10].Type != "HASH" || properties[8].Value != inserts[0].PropertyName {
		t.Fatalf("UpdateCSSTokenList() did not insert a new HASH '#newCopyID' properly")
	}

	if properties[11].Type != "CHAR" || properties[9].Value != "{" {
		t.Fatalf("UpdateCSSTokenList() did not insert a new CHAR '{' properly")
	}

	if properties[12].Type != "S" || properties[10].Value != "\n    " {
		t.Fatalf("UpdateCSSTokenList() did not insert a new S token properly")
	}

	if properties[13].Type != "IDENT" || properties[11].Value != "top" {
		t.Fatalf("UpdateCSSTokenList() did not insert a new IDENT 'top' properly")
	}

	if properties[12].Type != "CHAR" || properties[12].Value != ":" {
		t.Fatalf("UpdateCSSTokenList() did not insert a new CHAR ':' properly")
	}

	if properties[13].Type != "CHAR" || properties[13].Value != ":" {
		t.Fatalf("UpdateCSSTokenList() did not insert a new DIMENSION '20px' properly")
	}

	if properties[14].Type != "S" || properties[14].Value != "\n    " {
		t.Fatalf("UpdateCSSTokenList() did not insert a new S token properly")
	}

	if properties[15].Type != "CHAR" || properties[15].Value != "}" {
		t.Fatalf("UpdateCSSTokenList() did not insert a new CHAR '}' token properly")
	}
}

// Tests if CSS parsing functionality is valid
func TestTokenizeCSSFromFile(t *testing.T) {
	cssPath := "../testdata/filemgmt/test.css"
	TokenizeCSSFromFile(&cssPath)
}
