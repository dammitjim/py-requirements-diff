package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	prefixPtr := flag.String("prefix", "", "Prefix to add to each line")
	flag.Parse()

	file, err := getFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	added := make(map[string]Requirement)
	removed := make(map[string]Requirement)

	scanner := bufio.NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue // skip empty lines
		}

		// skip noise lines
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

	// changelog of requirements, before and after
	changed := make(map[string][]Requirement)
	for _, req := range added {
		// if the requirement is in both added and removed then it's changed
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

// getRequirement parses a line to return a requirement for output
func getRequirement(line string) Requirement {
	lineParts := strings.Split(line, ";")
	bitWeCareAbout := lineParts[0]
	requirementParts := strings.Split(bitWeCareAbout, "==")
	return Requirement{
		Name:    requirementParts[0][1:],
		Version: requirementParts[1],
	}
}

// getFile returns a file from either stdin or the first cli argument
func getFile() (*os.File, error) {
	isPipe, err := isInputFromPipe()
	if err != nil {
		return nil, err
	}

	if isPipe {
		return os.Stdin, nil
	}

	if len(os.Args) < 2 {
		return nil, fmt.Errorf("No file provided")
	}

	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		return nil, err
	}
	return file, nil
}

// isInputFromPipe inspects the stdin file info to determine if the input is piped from a previous command
func isInputFromPipe() (bool, error) {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return fileInfo.Mode()&os.ModeCharDevice == 0, nil
}
