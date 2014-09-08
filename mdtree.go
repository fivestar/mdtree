package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"unicode/utf8"
)

type ListWords struct {
	Needles   []string
	NameRegex *regexp.Regexp
}

func hasColorArg(args []string) bool {
	for _, arg := range args {
		if arg == "-C" || arg == "-n" {
			return true
		}
	}

	return false
}

func resolveListWords() *ListWords {
	var needles []string
	var nameRegex *regexp.Regexp

	if os.Getenv("LANG") == "C" {
		needles = []string{"|", "|-", "`-"}
		nameRegex = regexp.MustCompile("-- (.+)$")
	} else {
		needles = []string{"│", "├", "└"}
		nameRegex = regexp.MustCompile("── (.+)$")
	}

	return &ListWords{Needles: needles, NameRegex: nameRegex}
}

func calculateIndentLv(line string, listWords *ListWords) int {
	indentLv := 0

	pos := -1
	for _, needle := range listWords.Needles {
		p := strings.LastIndex(line, needle)
		if p > -1 && pos < p {
			pos = p
		}
	}

	if pos > -1 {
		indentLv = int(utf8.RuneCountInString(line[:pos])/4) + 1
	}

	return indentLv
}

func parseName(line string, listWords *ListWords) string {
	matches := listWords.NameRegex.FindStringSubmatch(line)
	if len(matches) > 0 {
		return matches[1]
	}

	return line
}

func ParseTree2Markdown(treeString string) string {
	var mdStrings []string

	listWords := resolveListWords()

	lines := strings.Split(treeString, "\n")
	for _, line := range lines {
		if line == "" {
			mdStrings = append(mdStrings, "")
			continue
		}

		indentLv := calculateIndentLv(line, listWords)
		name := parseName(line, listWords)

		mdStrings = append(mdStrings, strings.Repeat("    ", indentLv)+"* "+name)
	}

	return strings.Join(mdStrings, "\n")
}

func main() {
	args := os.Args[1:]

	if !hasColorArg(args) {
		args = append(args, "-C")
	}

	out, err := exec.Command("tree", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(ParseTree2Markdown(string(out[:])))
}
