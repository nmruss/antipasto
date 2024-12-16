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

	expectedProperties := []CSSToken{
		{Type: "HASH", Value: "#copy1"},
		{Type: "CHAR", Value: "{"},
		{Type: "S", Value: "\n"},
		{Type: "S", Value: "      "},
		{Type: "IDENT", Value: "top"},
		{Type: "CHAR", Value: ":"},
		{Type: "DIMENSION", Value: "44px"},
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

	for i, v := range expectedProperties {
		if properties[i] != v {
			t.Fatalf("UpdateCSSTokenList() Failed to place expected properties at Token Number %d; Expected: %s, Got: %s", i, v, properties[i])
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

	expectedProperties := []CSSToken{
		{Type: "HASH", Value: "#copy1"},
		{Type: "CHAR", Value: "{"},
		{Type: "S", Value: "\n"},
		{Type: "S", Value: "      "},
		{Type: "IDENT", Value: "top"},
		{Type: "CHAR", Value: ":"},
		{Type: "DIMENSION", Value: "10px"},
		{Type: "CHAR", Value: ";"},
		{Type: "S", Value: "\n"},
		{Type: "S", Value: "      "},
		{Type: "IDENT", Value: "left"},
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

	for i, v := range expectedProperties {
		if properties[i] != v {
			t.Fatalf("UpdateCSSTokenList() Failed to place expected properties at Token Number %d; Expected: %s, Got: %s", i, v, properties[i])
		}
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

	expectedProperties := []CSSToken{
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
		{Type: "HASH", Value: "#newCopyID"},
		{Type: "CHAR", Value: "{"},
		{Type: "S", Value: "\n"},
		{Type: "S", Value: "      "},
		{Type: "IDENT", Value: "top"},
		{Type: "CHAR", Value: ":"},
		{Type: "DIMENSION", Value: "20px"},
		{Type: "CHAR", Value: ";"},
		{Type: "S", Value: "\n"},
		{Type: "CHAR", Value: "}"},
	}

	for i, v := range expectedProperties {
		if properties[i] != v {
			t.Fatalf("UpdateCSSTokenList() Failed to place expected properties at Token Number %d; Expected: %s, Got: %s", i, v, properties[i])
		}
	}
}

// Tests if CSS parsing functionality is valid
func TestTokenizeCSSFromFile(t *testing.T) {
	cssPath := "../testdata/filemgmt/test.css"
	TokenizeCSSFromFile(&cssPath)
}
