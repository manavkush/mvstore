package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "my random key for testing"
	pathName := CASPathTransformFunc(key)
	fmt.Println(pathName)

	assert.Equal(t, pathName, "d7c6c/38aaa/c9532/b2501/ee7c2/3967a/d017c/ea879")
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(opts)

	data := bytes.NewReader([]byte("some file data. Might be some image.jpg or text file"))

	if err := store.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
