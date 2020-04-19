package encode

import (
	"fmt"
	"os"
	"os/exec"
)

type filePath string

const (
	inFile     filePath = "input.tmp.xml"
	scriptFile filePath = "script.tmp.xsl"
)

// HTML function encodes content HTML text based on the specified key and add css classes
func HTML(in string, keyFrom string, keyTo string, cssClass string) (out string, err error) {

	// TODO: use golang xslt packages (maybe with in memory files)

	// write input file
	f, err := os.Create(string(inFile))
	if err != nil {
		return "", fmt.Errorf("file creation failed: %v", err)
	}
	defer f.Close()

	// write content
	_, err = f.WriteString(fmt.Sprintf("<foo>%s</foo>", in))
	if err != nil {
		//fmt.Errorf("decompress %v: %v", name, err)
		//errors.New("can't work with 42")
		return "", fmt.Errorf("file write failed: %v", err)
	}

	// write script file
	f, err = os.Create(string(scriptFile))
	if err != nil {
		return "", fmt.Errorf("file creation failed: %v", err)
	}
	defer f.Close()

	// write content
	_, err = f.WriteString(script)
	if err != nil {
		//fmt.Errorf("decompress %v: %v", name, err)
		//errors.New("can't work with 42")
		return "", fmt.Errorf("file write failed: %v", err)
	}

	err = f.Sync()
	if err != nil {
		return "", fmt.Errorf("sync file to disk failed: %v", err)
	}

	//xsltproc src/encode.xsl var/input.xml > var/output.xml
	cmd := exec.Command("xsltproc", "--stringparam", "translateFrom", keyFrom, "--stringparam", "translateTo", keyTo, "--stringparam", "cssClass", cssClass, string(scriptFile), string(inFile))
	outByte, err := cmd.CombinedOutput()
	out = string(outByte)
	if err != nil {
		return "", fmt.Errorf("command execution failed: %v", err)
	}

	return out, err

}
