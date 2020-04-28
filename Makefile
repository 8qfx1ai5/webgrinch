

# build container image from container manifest file
.PHONY: build
build:
	docker build -t viewcrypt-alpha -f ./build/container-image/Dockerfile .


# run container based on an image
.PHONY: run
run:
	docker run --restart=always -d -p "80:8080" viewcrypt-alpha


# run container for dev local based on an image
.PHONY: rundev
rundev:
	docker run --rm  -d -p "80:8080" viewcrypt-alpha


# run service on local docker env for development
.PHONY: serve
serve: clear build rundev test


# run go unit tests
.PHONY: utest
utest:
	go test -v -count=1 ./internal/...


# run go integration tests
.PHONY: itest
itest:
	go test -v -count=1 ./test/...


# run go benchmark tests
.PHONY: btest
btest:
	cd internal/encode; go test -v -count=1 -bench=Encoding -cpuprofile=cpu.tmp.out
	cd internal/encode; go tool pprof cpu.tmp.out
	#web > ../../test/results/web_result.tmp.svg
	#top50 > ../../test/results/top50_result.tmp.txt


# run all tests
.PHONY: test
test:
	go test -count=1 ./...


.PHONY: ps
ps:
	docker ps -a


# ssh to digital ocean
.PHONY: access
# call like: make access ip=64.225.104.7
access:
	ssh -t root@$(ip) "cd /viewcrypt ; bash"


# deploy on digital ocean
 .PHONY: deploy
 # call like: make deploy dir=$(pwd) ip=64.225.104.7
 deploy:
	scp -r $(dir) root@$(ip):/
	ssh root@$(ip) "apt install make; cd /viewcrypt; docker system prune -f; make build; make run"


# get shell inside of the first running docker container
 .PHONY: login
 # call like: make login
 login:
	docker exec -it `docker ps -a -q | head -n 1` /bin/sh


# stop and remove all docker container
.PHONY: clear
# call like: make deploy dir=$(pwd) ip=64.225.104.7
clear:
	-docker stop `docker ps -a -q` 2>/dev/null
	-docker rm `docker ps -a -q` 2>/dev/null
	#docker system prune -f


# remove temp files
.PHONY: clean
clean:
	find . -type f -name '*.tmp.*' -print0 | xargs -0 rm
	find . -type f -name 'vcrypt*' -print0 | xargs -0 rm
	find . -type f -name '*.test' -print0 | xargs -0 rm

