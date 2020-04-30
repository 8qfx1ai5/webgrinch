package encodehtml

const (
	script string = `<?xml version="1.0" encoding="iso-8859-1" ?>
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
)
