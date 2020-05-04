package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	e "github.com/8qfx1ai5/viewcrypt/internal/types/export"
	"github.com/8qfx1ai5/viewcrypt/test/data/keyregexdata"
)

var exportPath = filepath.Join("web", "static", "export")

var testsuites []e.Export = []e.Export{
	keyregexdata.TestCases,
}

func main() {
	exportTestCasesToGit()
}

// function writes files
func exportTestCasesToGit() {
	for _, t := range testsuites {
		f, err := os.Create(filepath.Join(exportPath, t.FileName()))
		if err != nil {
			log.Print(fmt.Errorf("file creation failed: %v", err))
			continue
		}
		defer f.Close()

		// write content
		content, err := t.Export()
		if err != nil {
			log.Print(fmt.Errorf("test case export failed: %v", err))
			continue
		}
		_, err = f.WriteString(content)
		if err != nil {
			log.Print(fmt.Errorf("file write failed: %v", err))
			continue
		}

		err = f.Sync()
		if err != nil {
			log.Print(fmt.Errorf("sync file to disk failed: %v", err))
			continue
		}
	}
}
