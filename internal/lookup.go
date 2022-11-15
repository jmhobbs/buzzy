package internal

import (
	"bufio"
	"encoding/base32"
	"os"
	"path"
)

type LookupTable interface {
	Append(word string) error
	Lookup(pattern string) ([]string, error)
}

type DiskLookupTable struct {
	base string
}

func NewDiskLookupTable(path string) *DiskLookupTable {
	return &DiskLookupTable{base: path}
}

func fieldToPath(field [4]byte) (string, string) {
	fileName := base32.StdEncoding.EncodeToString(field[:])
	return path.Join(string(fileName[:2]), string(fileName[2:4])), fileName
}

func (d *DiskLookupTable) Append(word string) error {
	field := pack(word)

	dir, fileName := fieldToPath(field)

	fullDir := path.Join(d.base, dir)
	err := os.MkdirAll(dir, 0744)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path.Join(fullDir, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte(word)); err != nil {
		f.Close()
		return err
	}

	if _, err := f.Write([]byte{'\n'}); err != nil {
		f.Close()
		return err
	}

	return f.Close()
}

func (d *DiskLookupTable) Lookup(pattern string) ([]string, error) {
	field := pack(pattern)

	dir, fileName := fieldToPath(field)
	fullDir := path.Join(d.base, dir)

	f, err := os.Open(path.Join(fullDir, fileName))
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		panic(err)
	}

	defer f.Close()

	words := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, nil
	}

	return words, nil
}
