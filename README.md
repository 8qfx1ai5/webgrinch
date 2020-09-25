# Project "webgrinch" for html/content encoding

Trigger: the idea was publish content readable for all people, but which is not processable by automated systems and crallers.

This repository contains all required code to run an api server for the encoding part and some other handy scripts to use the functionality.

## How to use

<a href="https://app.getpostman.com/run-collection/0c3bbddf36204db54b25#?env%5Blocal%5D=W3sia2V5IjoiYmFzZVVybCIsInZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdC9hcGkiLCJlbmFibGVkIjp0cnVlfV0=" target="_blank"><img src="https://run.pstmn.io/button.svg" height="30px" alt="Run in Postman" /></a>

<a href="http://localhost" target="_blank"><img src="https://static1.smartbear.co/swagger/media/assets/images/swagger_logo.svg" height="30px" alt="Run in Swagger-UI" /></a>

### Deploy on PRODUCTION (docker required)

**ATTENTION**: this command will **stop and remove** all your docker containers on the remote machine and start the server on port 80. Unit tests are executed automatically during the deploy.

To ship the code to an remote docker droplet with known ip and ssh access.

    make deploy dir=$(pwd) ip=<address>

### Deploy on LOCAL for dev (docker required)

**ATTENTION**: this command will **stop and remove** all your local docker containers and start the server on port 80. Unit tests and integration tests are executed automatically during the deploy.

You can access over <http://localhost/>

    make serve

### How to include the text into a web page

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

    cd cmd/webgrinchserver; go build

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

## Special thanks

My special thanks goes to my friend and devops kubernetes specialist [Tino](https://github.com/pandorasNox). Without his knowledge and go workshop lessons the project would not look like the same.
