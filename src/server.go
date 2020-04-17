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
	in := `
		<!-- wp:paragraph -->
		<p>Schon mehrfach wurde ich gefragt ob ich auch bereit wäre selbstständig (also als Freelancer) für Firmen zu arbeiten. Bisher hat mir das immer Unbehagen bereitet. Ich kenne einfach keinen Selbstständigen, der nicht sehr viel Leid und Elend über seine Familie gebracht hätte mit dieser Art von Tätigkeit.</p>
		<!-- /wp:paragraph -->
		
		<!-- wp:paragraph -->
		<p>Doch jetzt ist alles anders. Seit dem 11.02.2020 bekomme ich kein ALG1 mehr und auf ALG2 hab ich keinen Anspruch. Ich plage mich also mit den gleichen Umständen, wie sie auch ein Selbstständiger hat. Da ich seit meiner Kindheit in Kontakt mit dem Sozialsystem bin, hab ich mit Arbeitslosigkeit nie Probleme gehabt. Der berühmte <strong>soziale Abstieg in die Armut</strong> hat bei mir niemals irgendwelche Ängste ausgelöst, da ich es sehr gut kannte, wusste was es für Möglichkeiten gibt und was zu tun ist. Doch ich hatte bis jetzt immer Angst vor dem <strong>sozialen Abstieg in die Selbstständigkeit</strong>.</p>
		<!-- /wp:paragraph -->
		
		<!-- wp:quote {"className":"is-style-large"} -->
		<blockquote class="wp-block-quote is-style-large"><p>"Wenn wir unseren tiefsten Punkt erreichen, sind wir bereit für die größten Veränderungen."</p><cite>unbekannt</cite></blockquote>
		<!-- /wp:quote -->
		
		<!-- wp:paragraph -->
		<p>Wissen bestimmt unsere Entscheidungen und ich habe das Gefühl, dass das Wissen um die Selbstständigkeit meinen zukünftigen Weg erheblich beeinflussen wird. Ich wurde gezwungen dies zu lernen, doch freue mich auf das was kommt.</p>
		<!-- /wp:paragraph -->
		
		<!-- wp:paragraph -->
		<p></p>
		<!-- /wp:paragraph -->
	`
	xCfg := XSLTConfig{}
	out, err := EncodeHtml(in, xCfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}

func testController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bar")
}

type XSLTConfig struct {
	// exceptionsCfg ...
	// keyCfg
}

type path string

const (
	IN_FILE path = "var/input.xml"
)

//EncodeHtml does stuff...
func EncodeHtml(in string, xCfg XSLTConfig) (out string, err error) {
	_ = xCfg

	f, err := os.Create(string(IN_FILE))
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
