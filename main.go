package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func main() {
	prefixPtr := flag.String("prefix", "", "Prefix to add to each line")
	flag.Parse()

	stream, err := getStream()
	if err != nil {
		panic(err)
	}

	added := make(map[string]Requirement)
	removed := make(map[string]Requirement)

	scanner := bufio.NewScanner(bufio.NewReader(stream))
	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == '+' || line[1] == '-' {
			continue
		}
		// if first character is + or - then print it
		if line[0] == '+' {
			req := getRequirement(line)
			added[req.Name] = req
		} else if line[0] == '-' {
			req := getRequirement(line)
			removed[req.Name] = req
		}
	}

	changed := make(map[string][]Requirement)
	for _, req := range added {
		if _, ok := removed[req.Name]; ok {
			changed[req.Name] = []Requirement{req, removed[req.Name]}
			delete(added, req.Name)
			delete(removed, req.Name)
		}
	}

	var output []string
	for _, req := range added {
		output = append(output, fmt.Sprintf("%sAdded: %s == %s", *prefixPtr, req.Name, req.Version))
	}
	for _, req := range removed {
		output = append(output, fmt.Sprintf("%sRemoved: %s == %s", *prefixPtr, req.Name, req.Version))
	}
	for _, reqs := range changed {
		output = append(output, fmt.Sprintf("%sChanged: %s == %s -> %s", *prefixPtr, reqs[0].Name, reqs[1].Version, reqs[0].Version))
	}

	fmt.Println(strings.Join(output, "\n"))
}

type Requirement struct {
	Name    string
	Version string
}

func getRequirement(line string) Requirement {
	lineParts := strings.Split(line, ";")
	bitWeCareAbout := lineParts[0]
	requirementParts := strings.Split(bitWeCareAbout, "==")
	return Requirement{
		Name:    requirementParts[0][1:],
		Version: requirementParts[1],
	}
}

func getStream() (io.Reader, error) {
	if isInputFromPipe() {
		return os.Stdin, nil
	} else {
		panic("Unsupported")
	}
}
