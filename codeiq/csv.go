package codeiq

import (
	"encoding/csv"
	"github.com/qiniu/iconv"
	"github.com/suin/ioreplacer"
	"os"
)

type Reader struct {
	file           *os.File
	csv            *csv.Reader
	iconvConverter iconv.Iconv
}

func NewReader(filename string) (this *Reader, err error) {
	this = new(Reader)

	this.file, err = os.Open(filename)

	if err != nil {
		return
	}

	this.iconvConverter, err = iconv.Open("utf-8", "sjis")

	if err != nil {
		return
	}

	iconvReader := iconv.NewReader(this.iconvConverter, this.file, 0)
	replacer := ioreplacer.NewReader(iconvReader, map[string]string{"\r": "\n", "\r\n": "\n"})

	this.csv = csv.NewReader(replacer)
	this.csv.FieldsPerRecord = -1

	return
}

func (this *Reader) Read() (record []string, err error) {
	return this.csv.Read()
}

func (this *Reader) Close() {
	defer this.file.Close()
	defer this.iconvConverter.Close()
}
