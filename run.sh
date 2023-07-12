#!/bin/bash

# Mac only.

FILE=static/main.mp4
URL=https://devstreaming-cdn.apple.com/videos/streaming/examples/img_bipbop_adv_example_fmp4/v9/main.mp4

if [ -f "$FILE" ]; then
    echo "$FILE exists."
else
    echo "$FILE does not exist, starting download."
    curl -o $FILE $URL
    if [ $? -eq 0 ]; then
        echo "Download successful."
    else
        echo "Download failed."
    fi
fi

echo "Installing server dependencies"
go get github.com/gin-gonic/gin

# Run the Go file
echo "Running main.go"
go run main.go

echo "Starting HTTP server and VLC player"
open -a VLC http://localhost:8080/static/manifest_main.m3u8

# Start HTTP Server
go run server.go
