# main app phonies
.PHONY: deploy build run prep-do clear clean
# -------------------------------------------------------------------------

# deploy from local with production config to remote
# just select ip address from digital ocean droplet
# call like: make deploy dir=$(pwd) ip=64.225.104.7
deploy: prep-do
	ssh root@$(ip) "cd /webgrinch; make build run"


# build container image from container manifest file
build:
	docker build -t webgrinch-alpha -f ./build/container-image/Dockerfile .


# run container based on an image
run:
	docker run --restart=always -d \
		-v $(shell readlink -f /etc/letsencrypt/live/webgrinch.8qfx1ai5.de/cert.pem):/certs/cert.pem \
		-v $(shell readlink -f /etc/letsencrypt/live/webgrinch.8qfx1ai5.de/privkey.pem):/certs/privkey.pem \
		-p "80:80" \
		-p "443:443" \
		webgrinch-alpha


# prepare droplet on digital ocean from remote
# call like: make prep-do dir=$(pwd) ip=64.225.104.7
prep-do: clean
	ssh root@$(ip) "apt-get update; apt-get -y upgrade; apt-get -y install docker.io; systemctl start docker; systemctl enable docker"
	#ssh root@$(ip) "mkdir webgrinch"
	rsync -v --archive --delete --exclude=.git* --compress $(dir) root@$(ip):/
	#scp -r $(dir) root@$(ip):/
	ssh root@$(ip) "apt-get -y install make; cd /webgrinch; make clear"


# stop and remove all docker container
# call like: make deploy dir=$(pwd) ip=64.225.104.7
clear:
	-docker stop `docker ps -a -q` 2>/dev/null
	-docker rm `docker ps -a -q` 2>/dev/null
	#docker system prune -f


# remove temp files
clean:
	find . -type f -name '*.tmp.*' -print0 | xargs -0 rm
	find . -type f -name 'webgrinch*' -print0 | xargs -0 rm
	find . -type f -name '*.test' -print0 | xargs -0 rm
	rm -rf tmp


# main app development phonies
.PHONY: rundev serve login loginforce access ps runtapid
# -------------------------------------------------------------------------


# run container for dev local based on an image
rundev:
	docker run --rm -d --name api \
		-v $(shell pwd)/tmp/certs:/certs \
		-p "80:80" \
		-p "443:443" \
		webgrinch-alpha


# run local api tests
# call like: make runt url=http://api/api
# docker run --rm --link api -t postman/newman run
# docker run --rm --link api newman run /Contract_Tests.postman_collection.json --env-var baseUrl=http://api/api
runtapid:
	docker build -t newman -f ./build/container-image-newman/Dockerfile .
	docker run --rm --link api newman run /Contract_Tests.postman_collection.json --env-var baseUrl=http://api/api


# run service on local docker env for development
serve: clean clear tls build rundev itestd


# ssh to digital ocean
# call like: make access ip=64.225.104.7
access:
	ssh -t root@$(ip) "cd /webgrinch ; bash"


# get shell inside of the first running docker container
# call like: make login
login:
	docker exec -it `docker ps -a -q | head -n 1` /bin/sh


# get shell inside of a stoped docker container
# call like: make loginforce
loginforce:
	docker run -it --env TLSCERT --env TLSCERTKEY --entrypoint /bin/sh webgrinch-alpha -s


# list all docker containers
ps:
	docker ps -a


# test phonies
.PHONY: utest utestd itestd itest btest test
# -------------------------------------------------------------------------


# run go unit tests
utest:
	docker build -t utest -f build/container-image-utest/Dockerfile .
	docker run utest


# run go unit tests during deploy
utestd:
	@echo "RUN go unit tests..."
	@go test -count=1 ./internal/... | egrep "^(FAIL.|ok)" | sed ''/ok/s//`printf "\033[32mok\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/'' | sed ''/?/s//`printf "\033[33m?\033[0m"`/''


# run go integration tests during deploy
itestd:
	@echo "RUN go integration tests..."
	@sleep 2 # starting the server is to slow
	@go test -count=1 ./test/... | egrep "^(FAIL.|ok)" | sed ''/ok/s//`printf "\033[32mok\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/'' | sed ''/?/s//`printf "\033[33m?\033[0m"`/''


# run go integration tests during deploy
itest:
	@echo "RUN go integration tests..."
	@go test -v -count=1 ./test/... | sed ''/PASS/s//`printf "\033[32mPASS\033[0m"`/'' | sed ''/ok/s//`printf "\033[32mok\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/'' | sed ''/?/s//`printf "\033[33m?\033[0m"`/''


# run go benchmark tests
btest:
	cd internal/encode; go test -v -count=1 -bench=Encoding -cpuprofile=cpu.tmp.out
	cd internal/encode; go tool pprof cpu.tmp.out
	#web > ../../test/results/web_result.tmp.svg
	#top50 > ../../test/results/top50_result.tmp.txt


# run all tests
test:
	go test -count=1 ./... | egrep "^(FAIL.|ok|?)" | sed ''/ok/s//`printf "\033[32mok\033[0m"`/'' | sed ''/FAIL/s//`printf "\033[31mFAIL\033[0m"`/'' | sed ''/?/s//`printf "\033[33m?\033[0m"`/''


# service phonies
.PHONY: swagger sass tls cirunner le-list le-create le-renew
# -------------------------------------------------------------------------


# recreate swagger static content from the files
swagger:
	go get github.com/rakyll/statik
	cd third_party/swagger-ui; statik -src=`pwd`/dist


# run sass conversion
sass:
	cd build/container-image-sass; docker build -t sass .
	cat web/static/example/scss/main.scss | docker run -i sass > web/static/example/css/main.css


# create self signed certificate for local development
tls:
	cd build/container-image-tls; docker build -t tls .
	mkdir -p tmp/certs
	docker run -i tls cert.pem > tmp/certs/cert.pem
	docker run -i tls privkey.pem > tmp/certs/privkey.pem


# list all the registered certificates on the system from letsencrypt
# run on remote host
le-list:
	sudo docker run -it --rm --name certbot \
            -v "/etc/letsencrypt:/etc/letsencrypt" \
            -v "/var/lib/letsencrypt:/var/lib/letsencrypt" \
            certbot/certbot certificates


# create a new letsencrypt certificate (only if you do not have an old one)
# run on remote host
le-create:
	sudo docker run -it --rm --name certbot \
            -p 80:80 \
            -p 443:443 \
            -v "/etc/letsencrypt:/etc/letsencrypt" \
            -v "/var/lib/letsencrypt:/var/lib/letsencrypt" \
            certbot/certbot certonly --standalone


# renew a letsencrypt certificate specified by name (also see list certificates)
# run on remote host
le-renew:
	# stop the current running server (downtime)
	docker stop `docker ps -a -q`
	# renew the cert
	sudo docker run -it --rm --name certbot \
            -p "80:80" \
            -p "443:443" \
            -v "/etc/letsencrypt:/etc/letsencrypt" \
            -v "/var/lib/letsencrypt:/var/lib/letsencrypt" \
            certbot/certbot certonly --standalone --force-renew --cert-name webgrinch.8qfx1ai5.de
	# restart the server
	make run
	## if you want to deploy to the server and for ssh access, you maybe need to update your .ssh/known_hosts file

