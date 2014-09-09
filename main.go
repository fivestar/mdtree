package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]

	if !HasColorArg(args) {
		args = append(args, "-C")
	}

	out, err := exec.Command("tree", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(ParseTree2Markdown(string(out[:])))
}
