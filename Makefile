

# build container image from container manifest file
.PHONY: build
build:
	docker build -t viewcrypt-alpha -f ./build/container-image/Dockerfile .


# run container based on an image
.PHONY: run
run:
	docker run --restart=always -d -p "80:80" viewcrypt-alpha


# run local api tests
.PHONY: runtapid
# call like: make runt url=http://api/api
# docker run --rm --link api -t postman/newman run
# docker run --rm --link api newman run /Contract_Tests.postman_collection.json --env-var baseUrl=http://api/api
runtapid:
	docker build -t newman -f ./build/container-image-newman/Dockerfile .
	docker run --rm --link api newman run /Contract_Tests.postman_collection.json --env-var baseUrl=http://api/api


# run container for dev local based on an image
.PHONY: rundev
rundev:
	docker run --rm  -d -p "80:80" --name api viewcrypt-alpha


# run service on local docker env for development
.PHONY: serve
serve: clear build rundev itestd


# run go unit tests
.PHONY: utest 
utest:
	go test -v -count=1 ./internal/... | sed ''/PASS/s//`printf "\033[32mPASS\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/''


# run go unit tests during deploy
.PHONY: utestd 
utestd:
	@echo "RUN go unit tests..."
	@go test -count=1 ./internal/... | egrep "^(FAIL.|ok)" | sed ''/ok/s//`printf "\033[32mok\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/'' | sed ''/?/s//`printf "\033[33m?\033[0m"`/''


# run go integration tests during deploy
.PHONY: itestd
itestd:
	@echo "RUN go integration tests..."
	@go test -count=1 ./test/... | egrep "^(FAIL.|ok)" | sed ''/ok/s//`printf "\033[32mok\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/'' | sed ''/?/s//`printf "\033[33m?\033[0m"`/''


# run go integration tests during deploy
.PHONY: itest
itest:
	@echo "RUN go integration tests..."
	@go test -v -count=1 ./test/... | sed ''/PASS/s//`printf "\033[32mPASS\033[0m"`/'' | sed ''/ok/s//`printf "\033[32mok\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/'' | sed ''/?/s//`printf "\033[33m?\033[0m"`/''


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
	go test -count=1 ./... | egrep "^(FAIL.|ok|?)" | sed ''/ok/s//`printf "\033[32mok\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/'' | sed ''/?/s//`printf "\033[33m?\033[0m"`/''


.PHONY: ps
ps:
	docker ps -a


# ssh to digital ocean
.PHONY: access
# call like: make access ip=64.225.104.7
access:
	ssh -t root@$(ip) "cd /viewcrypt ; bash"


# deploy from local with production config to remote
# just select ip address from digital ocean droplet
.PHONY: deploy
 # call like: make deploy dir=$(pwd) ip=64.225.104.7
 deploy: prep-do 
	ssh root@$(ip) "cd /viewcrypt; make build run"


# prepare droplet on digital ocean from remote
 .PHONY: prep-do
 # call like: make prep-do dir=$(pwd) ip=64.225.104.7
 prep-do: clean
	#ssh root@$(ip) "mkdir viewcrypt"
	rsync -v --archive --delete --exclude=.git* --compress $(dir) root@$(ip):/
	#scp -r $(dir) root@$(ip):/
	ssh root@$(ip) "apt install make; cd /viewcrypt; make clear"


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


# swagger
.PHONY: swagger
swagger:
	go get github.com/rakyll/statik
	cd third_party/swagger-ui; statik -src=`pwd`/dist

