package main

import(
	"testing"
)

func TestParseTree2Markdown(t *testing.T) {
	treeString := `.
├── main.go
└── README.md
`

	expected := `* .
    * main.go
    * README.md
`

	actual := ParseTree2Markdown(treeString)
	if expected != actual {
		t.Errorf("Not equals to:\n----\n%s", actual)
	}
}
