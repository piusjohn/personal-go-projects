package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, _ := os.ReadFile(filename)
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	splitted := strings.Split(content, "\n")
	asciimap := make(map[rune][]string)
	currentchar := ' '
	if len(splitted) == 0 {
		return nil, errors.New("empty file")
	}
	if len(splitted) < 855 {
		return nil, errors.New("corrupted bannerfiles")
	}
	for i := 1; i < len(splitted); i += 9 {
		if currentchar > '~' {
			break
		}
		asciimap[currentchar] = splitted[i : i+8]
		currentchar++
	}
	return asciimap, nil
}
