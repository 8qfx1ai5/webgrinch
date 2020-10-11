

# image to create a self signed cert for development
FROM ubuntu:20.04 AS tls

# install openssl
RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get -y install openssl

# required params
ENV domain=webgrinch.8qfx1ai5.de
ENV commonname=webgrinch.8qfx1ai5.de
ENV password=abcdefghijklmn

# optional params
ENV country=DE
ENV state=Berlin
ENV locality=Berlin
ENV organization=webgrinch.8qfx1ai5.de
ENV organizationalunit=IT
ENV email=webgrinch@8qfx1ai5.de

# The below command will create a file named 'server.pass.key' and place it in the same folder where the command is executed. 
# RUN openssl genrsa -des3 -passout pass:${password} -out server.pass.key -aes256 4096

# The below command will use the 'server.pass.key' file that just generated and create 'server.key'.
# RUN openssl rsa -passin pass:${password} -in server.pass.key -out server.key
 
# We no longer need the 'server.pass.key'
# RUN rm server.pass.key

# Create a cert request. The below command will ask you for information that would be included in the certificate. Since this is a self-signed certificate, there is no need to provide the 'challenge password' (to leave it blank, press enter).
#RUN openssl req -new -key server.key -out server.csr -passin pass:${password} -subj "/C=${country}/ST=${state}/L=${locality}/O=${organization}/OU=${organizationalunit}/CN=${commonname}/emailAddress=${email}"

# Create a cert and private key pair for a local development webserver
RUN openssl req -new -passout pass:${password} -x509 -newkey rsa:4096 -days 365 -keyout privkey.pem -out cert.pem -subj "/C=${country}/ST=${state}/L=${locality}/O=${organization}/OU=${organizationalunit}/CN=${commonname}/emailAddress=${email}"

# Remove the password from the private key
RUN openssl rsa -in privkey.pem -out privkey.pem -passin pass:${password}

ENTRYPOINT ["/usr/bin/cat"]
