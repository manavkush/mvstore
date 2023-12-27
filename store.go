package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

type PathTransformFunc func(string) PathKey

type PathKey struct {
	PathName string
	FileName string
}

func (p PathKey) FullPath() string {
	return path.Join(p.PathName, p.FileName)
}

func DefaultPathTransformFunc(key string) string {
	return key
}

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	slicelength := len(hashStr) / blocksize
	paths := make([]string, slicelength)

	for i := 0; i < slicelength; i++ {
		from, to := i*blocksize, (i+1)*blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		PathName: path.Join(paths...),
		FileName: hashStr,
	}
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

func (s *Store) getFirstPathDir(key string) string {
	pathKey := s.PathTransformFunc(key)

	firstPathDir := strings.Split(pathKey.PathName, "/")[0]
	return firstPathDir
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)
	return os.Open(pathKey.FullPath())
}

func (s *Store) Has(key string) bool {
	pathKey := s.PathTransformFunc(key)

	_, err := os.Stat(pathKey.FullPath())
	if err == fs.ErrNotExist {
		return false
	}
	return true
}

func (s *Store) Delete(key string) error {
	defer func() {
		log.Printf("Deleted key: %s", key)
	}()
	firstPathDir := s.getFirstPathDir(key)
	return os.RemoveAll(firstPathDir)
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)

	return buf, err
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	pathName := pathKey.PathName

	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}

	fullPath := pathKey.FullPath()
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("Written (%d) bytes to disk: %s", n, fullPath)

	return nil
}
