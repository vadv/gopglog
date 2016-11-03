package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Parser struct {
	fd          *os.File
	scanner     *bufio.Scanner
	line_prefix *regexp.Regexp
}

func Create(filename string) (*Parser, error) {

	result := &Parser{}
	if fd, err := os.Open(filename); err != nil {
		return nil, err
	} else {
		result.fd = fd
	}

	result.scanner = bufio.NewScanner(result.fd)

	return result, nil
}

func (p *Parser) DebugPrintAll() {
	for p.scanner.Scan() {
		source := p.scanner.Bytes()
		line, err := ParseLine(source)
		if err == nil {
			if line != nil {
				fmt.Printf("line:\t%s\n", line.ToString())
			}
		} else {
			fmt.Printf("error:\t%s for string: %s\n", err.Error(), source)
		}
	}
}
