package encode

import (
	"fmt"
	"os"
	"os/exec"
)

// XsltConfig comming soon
type XsltConfig struct {
	// exceptionsCfg ...
	// keyCfg
}

type filePath string

const (
	inFile     filePath = "input.tmp.xml"
	scriptFile filePath = "script.tmp.xsl"
)

//HTML does stuff...
func HTML(in string, keyFrom string, keyTo string, cssClass string) (out string, err error) {

	script := `<?xml version="1.0" encoding="iso-8859-1" ?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

	<!-- suppress xml declaration on top of the output file -->
	<xsl:output omit-xml-declaration="yes" indent="no" />

	<!-- name of the css class given to the first level nodes -->
	<xsl:param name="cssClass" />

	<!-- given current alphabet -->
	<xsl:param name="translateFrom" />

	<!-- given resulting alphabet -->
	<xsl:param name="translateTo" />

	<!-- copy all nodes + attributes + comments by default -->
	<xsl:template match="node() | @*">
		<xsl:copy>
			<xsl:apply-templates select="node() | @* | comment()" />
		</xsl:copy>
	</xsl:template>

	<!-- ignore root node -->
	<xsl:template match="/node()">
		<xsl:apply-templates select="node() | @* | comment()" />
	</xsl:template>

	<!-- set css class to second level nodes -->
	<xsl:template match="/*/*">
		<xsl:copy>
			<xsl:choose>
				<xsl:when test="$cssClass  != ''">
					<xsl:attribute name="class">
						<xsl:value-of select="normalize-space(concat($cssClass,' ',@class))" />
					</xsl:attribute>
					<xsl:apply-templates select="node() | @*[name()!='class'] | comment()" />
				</xsl:when>
				<xsl:otherwise>
					<xsl:apply-templates select="node() | @* | comment()" />
				</xsl:otherwise>
			</xsl:choose>
		</xsl:copy>
	</xsl:template>

	<!-- do the hard encode work to the text -->
	<xsl:template match="text()">
		<xsl:value-of select="translate(.,$translateFrom,$translateTo)" />
	</xsl:template>

</xsl:stylesheet>`

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
