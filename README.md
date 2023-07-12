# fmp4
Go language implementation of simple byte-range manifest generator for fragmented mp4 files.  
Disclaimer: This project is mainly for learning purpose, as I'm new to Go language and I wanted to learn more about ISO base media file format like MP4. So you might find some bugs or unused code, etc... Any advice or suggestion is welcome.

## Usage
To generate simple example manifests for fragmented mp4 files, run the shell script (Mac only). It will download test video, parse Atoms (aka Boxes), and generate byte-range manifests for video track. For convenience, it will also start a simple http server using [Gin](https://github.com/gin-gonic/gin) to serve the manifests and video files with byte-rage retrieve possibility, and open it in VLC player.
```shell
sh run.sh
```