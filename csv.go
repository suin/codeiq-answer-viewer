package main

import (
	"encoding/csv"
	"github.com/axgle/mahonia"
	"github.com/suin/ioreplacer"
	"os"
)

type Reader struct {
	file *os.File
	csv  *csv.Reader
}

func NewReader(filename string) (this *Reader, err error) {
	this = new(Reader)

	this.file, err = os.Open(filename)

	if err != nil {
		return
	}

	converter := mahonia.NewDecoder("shiftjis").NewReader(this.file)

	replacer := ioreplacer.NewReader(converter, map[string]string{"\r": "\n", "\r\n": "\n"})

	this.csv = csv.NewReader(replacer)
	this.csv.FieldsPerRecord = -1

	return
}

func (this *Reader) Read() (record []string, err error) {
	return this.csv.Read()
}

func (this *Reader) Close() {
	defer this.file.Close()
}
