package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.Use(CORSMiddleware())

	r.GET("/files/*file_path", func(c *gin.Context) {
		filePath := "static" + c.Param("file_path")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"detail": "File not found"})
			return
		}

		rangeHeader := c.GetHeader("Range")
		rangeParts := strings.Split(strings.Split(rangeHeader, "=")[1], "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])
		fmt.Println("Client requests byte range:", start, end)

		file, _ := os.Open(filePath)
		defer file.Close()

		_, _ = file.Seek(int64(start), 0)
		buf := make([]byte, (end+1)-start)

		if _, err := file.Read(buf); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fileInfo, _ := file.Stat()
		reponseHeader := map[string]string{
			"Content-Range": fmt.Sprintf("bytes %d-%d/%d", start, end, fileInfo.Size()),
			"Accept-Ranges": "bytes",
		}
		c.DataFromReader(
			http.StatusPartialContent,
			int64(len(buf)),
			"video/mp4",
			bytes.NewReader(buf), reponseHeader,
		)
	})

	_ = r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Next()
	}
}
