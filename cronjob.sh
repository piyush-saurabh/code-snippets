#!/bin/sh
# This script will check if any new process is spawned because of cron job

# Loop by line
IFS=$'\n'
old_process=$(ps -eo command)

while true; do
    new_process=$(ps -eo command)
    diff <(echo "$old_process") <(echo "$new_process")
    sleep 1
    old_process=$new_process
done
