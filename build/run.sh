#!/bin/sh

# start all required services
echo "start server" >> /server.log
/webgrinchserver -p 80 >> /server.log 2>&1 &
sleep 2
