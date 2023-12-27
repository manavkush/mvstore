package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path"
)

type PathTransformFunc func(string) string

func DefaultPathTransformFunc(key string) string {
	return key
}

func CASPathTransformFunc(key string) string {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	slicelength := len(hashStr) / blocksize
	paths := make([]string, slicelength)

	for i := 0; i < slicelength; i++ {
		from, to := i*blocksize, (i+1)*blocksize
		paths[i] = hashStr[from:to]
	}

	return path.Join(paths...)
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathName := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return nil
	}

	buff := new(bytes.Buffer)
	io.Copy(buff, r)

	fileNameBytes := md5.Sum(buff.Bytes())
	fileName := hex.EncodeToString(fileNameBytes[:])
	filePath := path.Join(pathName, fileName)

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, buff)
	if err != nil {
		return err
	}

	log.Printf("Written (%d) bytes to disk: %s", n, filePath)

	return nil
}
