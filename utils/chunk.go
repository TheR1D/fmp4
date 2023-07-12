package utils

import (
	"os"
)

func ReadChunk(file *os.File, byteRange ByteRange) ([]byte, error) {
	buf := make([]byte, byteRange.Length)
	_, err := file.ReadAt(buf, int64(byteRange.Start))
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func WriteChunk(fileName string, buf []byte) error {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func MustReadChunk(file *os.File, byteRange ByteRange) []byte {
	buf, err := ReadChunk(file, byteRange)
	if err != nil {
		panic(err)
	}
	return buf
}

func MustWriteChunk(fileName string, buf []byte) {
	err := WriteChunk(fileName, buf)
	if err != nil {
		panic(err)
	}
}
