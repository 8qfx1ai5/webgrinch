# Project "webgrinch" for html/content encoding

Trigger: the idea was publish content readable for all people, but not processable by automated systems and crallers.

To be more specific: the text is only understandable with visual analysis like ML for image-text recognition or specialized text analysis software based on letter frequencies or so.

This repository contains all required code to run an api server for the encoding part and some other handy scripts to use the functionality.

## How to use

1. **Here is an example html page to show how it works:**
<br/>http://webgrinch.8qfx1ai5.de

2. **Here is a working instance (+ Swagger) of the API to convert text:**
<br/>http://webgrinch.8qfx1ai5.de/api

3. **Postman tests and documentation for the API:**
<br/><a href="https://app.getpostman.com/run-collection/0c3bbddf36204db54b25#?env%5Blocal%5D=W3sia2V5IjoiYmFzZVVybCIsInZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdC9hcGkiLCJlbmFibGVkIjp0cnVlfV0=" target="_blank"><img src="https://run.pstmn.io/button.svg" height="30px" alt="Run in Postman" /></a>

## Deployment of the API

You can setup your own instance of the API on every machine with a docker installation.

### Deploy on LOCAL for development (docker required)

**ATTENTION**: this command will **stop and remove** all your local docker containers and start the server on port 80. Unit tests and integration tests are executed automatically during the deploy. Use the following command in the project root directory:

    make serve

You can access the service over "http://localhost:80"

### Deploy on REMOTE (docker required)

**ATTENTION**: this command will **stop and remove** all your docker containers on the remote machine and start the server on port 80. Unit tests are executed automatically during the deploy.

Ship the code to an remote docker droplet with known IP and ssh access. The local code from the selected local **directory (dir)** will be copied with rsync to the **IP**. Then a build process is started on the remote machine. Use the following command in the project root directory:

    make deploy dir=$(pwd) ip=<address>

You can access the service over the "http://IP-address"

## How to include the text into a web page

To make the text readable you need to load the corresponding web font to the text. For instance:

    <style>
        @font-face {
            font-family: <your-selected-font-family-name>;
            src: url(<select-a-WebGrinch-font.ttf>) format('truetype');
        }

        .decode {
            font-family: <your-selected-font-family-name>;
        }
    </style>

### Run builds

    cd cmd/webgrinchserver; go build

### Run the tests

    make utest # for debugging
    make utestd # for deploy

#### All tests

    go test -v -count=1 ./...

#### Unit tests

    go test -v -count=1 ./internal/...

#### Integration tests

    go test -v -count=1 ./test/...

#### Benchmark tests

    $ cd internal/encode; go test -v -count=1 -bench=Encoding -cpuprofile=cpu.tmp.out
    $ go tool pprof cpu.tmp.out
    (pprof)$ web > ../../test/results/web_result.tmp.svg
    (pprof)$ top50 > ../../test/results/top50_result.tmp.txt

## Project evolution

- v0.4 was build on plain xslt running with php
- v0.5 was switched to golang to enable webservice functionality, but running xsltproc over cli exec
- v0.6 add simple api webserver
- v0.7 use docker deployment
- v0.8 enable concurrent xslt processing
- v0.9 use logging
- v0.10 use contract testing for the api + swagger
- v0.11 moved product organisation to GitHub projects
- v1.0 decision to make the project open source under the name "webgrinch"

## Features

- plain text conversion
- html fragments conversion (dirty)
- xhtml conversion
- (xml conversion)

## Missing Features

- create a font with a special key
- advanced regex support for the key generation
- advanced html understanding of the API

## Special thanks

My special thanks goes to my friend and devops kubernetes specialist [Tino](https://github.com/pandorasNox). Without his knowledge and go workshop lessons the project would not look like the same.
