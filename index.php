<?php

//$content = $_GET['content'];

//$cmd = "xsltproc src/encode.xsl var/input.xml > var/output.xml";
//$result = shell_exec ( $cmd );

$xml = new DOMDocument;
//$xml->load('var/input.xml');

// TODO:

// FIXME:

$xml->loadXML("<wrapper/>");
$f = $xml->createDocumentFragment();

$content = file_get_contents('var/input.html');
// $content = $_POST['content'];
$f->appendXML($content);
$xml->documentElement->appendChild($f);

$xsl = new DOMDocument;
$xsl->load('src/encode.xsl');

// Prozessor instanziieren und konfigurieren
$proc = new XSLTProcessor;
$proc->importStyleSheet($xsl); // XSL Document importieren

echo $proc->transformToXML($xml);

