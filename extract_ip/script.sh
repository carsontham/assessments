#!/bin/bash

echo "extracting files.."

mkdir temp && tar -xf archive_log.tar.gz -C temp

access_logs=$(find . -name 'access.[0-9][0-9].log')

for log_file in $access_logs; do
    # Extract records with non-empty IP address and append to access.log
    grep -E '[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+' "$log_file" >> access.log
done