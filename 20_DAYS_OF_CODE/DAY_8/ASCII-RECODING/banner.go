package main

import (
	"errors"
	//"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, _ := os.ReadFile(filename)
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	splitted := strings.Split(content, "\n")
	hash := make(map[rune][]string)
	if len(splitted) == 0 {
		return nil, errors.New("empty files")
	}
	if len(splitted) < 855 {
		return nil, errors.New("corrupted files")
	}
	currentchar := ' '
	for i := 1; i < len(splitted); i += 9 {
		if currentchar > 126 {
			break
		}
		hash[currentchar] = splitted[i : i+8]
		currentchar++
	}
	return hash, nil
}
