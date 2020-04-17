<?xml version="1.0" encoding="iso-8859-1" ?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

	<!-- suppress xml declaration on top of the output file -->
	<xsl:output omit-xml-declaration="yes" indent="yes" />

	<!-- the locations of the config files (hard coded) -->
	<xsl:param name="chars-config-file" select="'../config/key.xml'" />
	<xsl:param name="exceptions-config-file" select="'../config/exceptions.xml'" />

	<!-- load configuration -->
	<xsl:variable name="chars" select="document($chars-config-file)/key/dictionary/*" />
	<xsl:variable name="cssclass" select="document($exceptions-config-file)/config/css/class/text()" />

	<!-- select current alphabet -->
	<xsl:variable name="translateFrom">
		<xsl:for-each select="$chars">
			<xsl:value-of select="original/text()" />
		</xsl:for-each>
	</xsl:variable>

	<!-- select resulting alphabet -->
	<xsl:variable name="translateTo">
		<xsl:for-each select="$chars">
			<xsl:value-of select="linked/text()" />
		</xsl:for-each>
	</xsl:variable>

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

	<!-- ignore blank lines -->
	<xsl:template match="*/text()[not(normalize-space())]" />

	<!-- set css class to second level nodes -->
	<xsl:template match="/*/*">
		<xsl:copy>
			<xsl:choose>
				<xsl:when test="$cssclass  != ''">
					<xsl:attribute name="class">
						<xsl:value-of select="normalize-space(concat($cssclass,' ',@class))" />
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
		<xsl:value-of select="translate(.,$translateFrom,$translateTo)"/>
	</xsl:template>

</xsl:stylesheet>
