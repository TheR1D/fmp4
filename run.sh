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

# Compile the Go file
echo "Compiling main.go"
go build main.go

if [ $? -eq 0 ]; then
    echo "Compilation successful"
else
    echo "Compilation failed"
    exit 1
fi

# Run the Go file
echo "Running main.go"
./main

if [ $? -eq 0 ]; then
    echo "Running of main.go was successful"
else
    echo "Running of main.go failed"
    exit 1
fi

echo "Installing Python dependencies"
pip install "fastapi[all]"

echo "Starting Python HTTP server and VLC player"
open -a VLC http://localhost:8000/static/manifest_main.m3u8

# Start Python HTTP Server
uvicorn server:app --reload
