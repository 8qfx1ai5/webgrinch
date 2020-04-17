package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// fmt.Println("foo")
	// http.HandleFunc("/test", testController)
	// http.Handle("/", http.FileServer(http.Dir("../static")))
	// http.ListenAndServe(":8888", nil)

	in := `<!-- this is a comment -->
	<p>Lorem ipsum dolor sit amet, consectetur <b>adipisicing</b> elit. Repellat, deleniti!</p>`
	xCfg := xsltConfig{}
	out, err := EncodeHTML(in, xCfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out)
}

func testController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bar")
}

type xsltConfig struct {
	// exceptionsCfg ...
	// keyCfg
}

type path string

const (
	inFile path = "var/input.xml"
)

//EncodeHTML does stuff...
func EncodeHTML(in string, xCfg xsltConfig) (out string, err error) {
	_ = xCfg

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

	err = f.Sync()
	if err != nil {
		return "", fmt.Errorf("sync file to disk failed: %v", err)
	}

	//xsltproc src/encode.xsl var/input.xml > var/output.xml
	outByte, err := exec.Command("xsltproc", "src/encode.xsl", "var/input.xml").Output()

	if err != nil {
		return "", fmt.Errorf("command execution failed: %v", err)
	}

	out = string(outByte)

	return out, err

}
