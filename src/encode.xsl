<?xml version="1.0" encoding="iso-8859-1"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:param name="chars-config-file" select="'../config/key.xml'"/>
<xsl:param name="exceptions-config-file" select="'../config/exceptions.xml'"/>

<xsl:variable name="chars" select="document($chars-config-file)/key/dictionary/*"/>
<xsl:variable name="cssclass" select="document($exceptions-config-file)/config/css/class/text()"/>

<xsl:variable name="translateFrom">
	<xsl:for-each select="$chars">
		<xsl:value-of select="original/text()"/>
	</xsl:for-each>
</xsl:variable>

<xsl:variable name="translateTo">
	<xsl:for-each select="$chars">
		<xsl:value-of select="linked/text()"/>
	</xsl:for-each>
</xsl:variable>

<xsl:template match="node() | @* | comment()">
	<xsl:copy>
		<xsl:apply-templates select="node() | @* | comment()"/>
	</xsl:copy>
</xsl:template>

<xsl:template match="/*/node()">
	<xsl:copy>
		<xsl:choose>
			<xsl:when test="$cssclass  != ''">
				<xsl:attribute name="class"><xsl:value-of select="normalize-space(concat($cssclass,' ',@class))" />
				</xsl:attribute> 
				<xsl:apply-templates select="node() | @*[name()!='class'] | comment()"/>
			</xsl:when>
			<xsl:otherwise>
				<xsl:apply-templates select="node() | @* | comment()"/>
			</xsl:otherwise>
		</xsl:choose>
	</xsl:copy>
</xsl:template>

<xsl:template match="text()[not(ancestor::*[name()='h2'])]">
	<xsl:value-of select="translate(.,$translateFrom,$translateTo)"/>
</xsl:template>

</xsl:stylesheet>
