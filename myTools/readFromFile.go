package myTools

import (
	"io"
	"os"
)

func MustReadContentFile(filepath string) string {
	content, err := ReadContentFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func ReadContentFile(filepath string) ([]byte, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	payload, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
