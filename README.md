# Project "viewcrypt" for html/content encoding

Trigger: the idea was publish content readable for all people, but which is not processable by automated systems and crallers.

This repository contains all required code to run an api server for the encoding part and some other handy scripts to use the functionality.

## How to use

### Deploy in production

Currently there is no automated pibeline. But it is comming soon.

### Deploy local for dev (docker required)

**ATTENTION**: this command will **stop and remove** all your local docker containers and start the server on port 80. Unit tests and integration tests are executed automatically during the deploy.

You can access over <http://localhost/>

    make serve

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

### Run builds

    cd cmd/vcryptcli; go build
    cd cmd/vcryptserver; go build

### Run the tests

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

## Products

- plain text conversion
- html fragments conversion (dirty)
- xhtml conversion
- (xml conversion)
- epub

## Special thanks

My special thanks goes to my friend and devops kubernetes specialist [Tino](https://github.com/pandorasNox). Without his knowledge and go workshop lessons the project would not look like the same.
