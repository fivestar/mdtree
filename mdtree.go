package main

import (
    "fmt"
    "log"
    "math"
    "os"
    "os/exec"
    "regexp"
    "strings"
)

func main() {
    args := os.Args[1:]

    hasColorArg := false
    for _, arg := range args {
        if arg == "-C" || arg == "-n" {
            hasColorArg = true
            break
        }
    }
    if !hasColorArg {
        args = append(args, "-C")
    }

    out, err := exec.Command("tree", args...).Output()
    if err != nil {
        log.Fatal(err)
    }

    var needles []string
    var nameRegex *regexp.Regexp
    if os.Getenv("LANG") == "C" {
        needles = []string{"|", "|-", "`-"}
        nameRegex = regexp.MustCompile("-- (.+)$")
    } else {
        needles = []string{"│", "├", "└"}
        nameRegex = regexp.MustCompile("── (.+)$")
    }

    currentPos := -1
    indentLv := 0

    lines := strings.Split(string(out[:]), "\n")
    for _, line := range lines {
        if line == "" {
            fmt.Println()
            continue
        }

        pos := -1

        for _, needle := range needles {
            p := strings.LastIndex(line, needle)
            if p > -1 {
                pos = int(math.Max(float64(p), float64(pos)))
            }
        }

        switch {
        case pos == -1:
            indentLv = 0
            break
        case pos == 0:
            indentLv = 1
            break
        case currentPos < pos:
            indentLv++
            break
        case currentPos > pos:
            indentLv--
            break
        }

        currentPos = pos

        name := ""
        matches := nameRegex.FindStringSubmatch(line)
        if len(matches) > 0 {
            name = matches[1]
        } else {
            name = line
        }

        fmt.Println(strings.Repeat("    ", indentLv) + "* " + name)
    }
}
