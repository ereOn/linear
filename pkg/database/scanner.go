package database

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Scanner scans for databases.
type Scanner struct {
	FileName  string
	Directory string
}

// Scan for databases.
func (s Scanner) Scan() (path string, database *Database, err error) {
	dir := s.Directory

	for {
		path = filepath.Join(dir, s.FileName)

		var f io.ReadCloser

		if f, err = os.Open(path); err == nil {
			defer f.Close()

			database, err = FromReader(f)

			return
		}

		previousDir := dir
		dir = filepath.Dir(dir)

		if previousDir == dir {
			err = fmt.Errorf("no `%s` database found in `%s`", s.FileName, s.Directory)
			return
		}
	}
}

// DefaultDatabaseFileName is the default database filename.
const DefaultDatabaseFileName = ".linear.yml"

// NewDefaultScanner create a new scanner that starts from the current
// directory, and uses the default database filename.
func NewDefaultScanner() (Scanner, error) {
	wd, err := os.Getwd()

	return Scanner{
		FileName:  DefaultDatabaseFileName,
		Directory: wd,
	}, err
}
