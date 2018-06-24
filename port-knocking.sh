#!/bin/bash
# This is the script for port knocking
# After sending the ping request to 'n' ports in a particular sequence, it will open up a particular port. e.g ssh
# Usage : ./port-knocking.sh <ip> <port1> <port2> <port3> && ssh username@ip

HOST=$1
shift
for ARG in "$@"
do
    nmap -Pn --max-retries 0 -p $ARG $HOST
done