#!/bin/sh

if [ "$APP_SERVICE" == "customer" ]
then
    exec "/app/customer"
elif [ "$APP_SERVICE" == "flight" ]
then
    exec "/app/flight"
else
    exec "/app/booking"
fi