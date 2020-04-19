# Project "viewcrypt" for html/content encoding

Trigger: the idea was publish content readable for all people, but which is not processable by automated systems and crallers.

This repository contains all required code to run an api server for the encoding part and some other handy scripts to use the functionality.

## How to use

### Run the cli script

    $ go run cmd/cli.go

### Run the server

    $ go run cmd/server.go

### Include the text into a web page

To make the text readable you need to load the corresponding web font to the text. For instance:

    <style>
        @font-face {
            font-family: OpenSans-Regular-vc;
            src: url(https://8qfx1ai5.de/vc/font/open_sans/OpenSans-Regular-vc.ttf) format('truetype');
        }

        .vc {
            font-family: OpenSans-Regular-vc, sans-serif;
        }
    </style>

### Run the tests

#### All tests

    $ go test -v -count=1 ./...

#### Benchmark tests

    $ cd internal/encode; go test -v -count=1 -bench=Encoding -cpuprofile=cpu.tmp.out
    $ go tool pprof perf.test cpu.tmp.out
    (pprof)$ web > ../../test/results/web_result.tmp.svg
    (pprof)$ top50 > ../../test/results/top50_result.tmp.svg

## Project evolution

- version 0.4 was build on plain xslt running with php
- version 0.5 was switched to golang to enable webservice functionality, but running xsltproc over cli exec

## Possible future next steps

- running api web server
- unsing go xslt packages instead of xsltproc
- enable decoding
- font generation for custom keys
- use contract testing for the api

## Special thanks

My special thanks goes to my friend and devops kubernetes specialist [Tino](https://github.com/pandorasNox). Without his knowledge and go workshop lessons the project would not look like the same.
