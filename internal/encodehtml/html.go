package encodehtml

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type filePath string

const (
	scriptFile filePath = "script.tmp.xsl"
)

var scriptFileCreated bool = false

// Run function encodes content HTML text based on the specified key and add css classes
func Run(in string, keyFrom string, keyTo string, cssClass string) (out string, err error) {

	// TODO: use golang xslt packages (maybe with in memory files)

	// write script file the first time
	if !scriptFileCreated {
		_, err := WriteXSLFile()
		if err != nil {
			return "", err
		}
		scriptFileCreated = true
	}

	out, err = runCliCommand(in, keyFrom, keyTo, cssClass)

	if err != nil {
		return "", fmt.Errorf("command execution failed: %v", err)
	}

	return out, err
}

func runCliCommand(in string, keyFrom string, keyTo string, cssClass string) (string, error) {
	cmd := exec.Command("xsltproc", "--stringparam", "translateFrom", keyFrom, "--stringparam", "translateTo", keyTo, "--stringparam", "cssClass", cssClass, string(scriptFile), "-")

	cmd.Stdin = strings.NewReader(fmt.Sprintf("<foo>%s</foo>", in))
	outByte, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command execution failed: %v", err)
	}
	out := string(outByte)
	return out, nil
}

// WriteXSLFile creates a file with the xsl transformation templates
func WriteXSLFile() (success bool, err error) {
	f, err := os.Create(string(scriptFile))
	if err != nil {
		return false, fmt.Errorf("file creation failed: %v", err)
	}
	defer f.Close()

	// write content
	_, err = f.WriteString(script)
	if err != nil {
		return false, fmt.Errorf("file write failed: %v", err)
	}

	err = f.Sync()
	if err != nil {
		return false, fmt.Errorf("sync file to disk failed: %v", err)
	}

	return true, nil
}
