#!/bin/sh

# start all required services
echo "start server" >> /server.log
/webgrinchserver -p 8080 >> /server.log 2>&1 &
sleep 2

echo "start proxy" >> /server.log
/webgrinchreverseproxy >> /server.log 2>&1
