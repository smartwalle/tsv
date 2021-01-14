package tsv

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
)

type parser struct {
	mu   sync.RWMutex
	sep  string
	rows []*Row
}

func newParser(sep string) *parser {
	var p = &parser{}
	p.sep = sep
	return p
}

func (this *parser) loadFile(file string) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	var f, err = os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	err = this.load(f)
	f.Close()
	if err != nil {
		return err
	}

	return nil
}

func (this *parser) load(r io.Reader) error {
	var reader = bufio.NewReader(r)
	var line []byte
	var err error

	var index = 0
	var rows []*Row

	for {
		if line, _, err = reader.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if index == 0 {
			line = bytes.TrimPrefix(line, []byte("\xef\xbb\xbf"))
		}
		index++

		var sLine = strings.TrimSpace(string(line))

		// 如果是空行,则忽略
		if sLine == "" {
			continue
		}
		// 如果是注释,则忽略
		if strings.HasPrefix(sLine, "#") || strings.HasPrefix(sLine, ";") {
			continue
		}

		var items = strings.Split(sLine, this.sep)
		if len(items) == 0 {
			continue
		}
		var row = &Row{}
		row.values = items
		rows = append(rows, row)
	}
	this.rows = rows
	return nil
}

func (this *parser) Len() int {
	return len(this.rows)
}

func (this *parser) RowAtIndex(index int) *Row {
	if index >= len(this.rows) {
		return nil
	}
	return this.rows[index]
}

func (this *parser) Rows() []*Row {
	return this.rows
}
