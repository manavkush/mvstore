package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "my random key for testing"
	pathName := CASPathTransformFunc(key)
	fmt.Println(pathName)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(opts)

	data := []byte("some file data. Might be some image.jpg or text file")

	if err := store.writeStream("myspecialpicture", bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := store.Read("myspecialpicture")
	if err != nil {
		t.Error(err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("string b:", string(b))

	assert.Equal(t, data, b)
}

func TestStoreDelete(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(opts)

	data := []byte("some file data. Might be some image.jpg or text file")

	if err := store.writeStream("myspecialpicture", bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	err := store.Delete("myspecialpicture")
	if err != nil {
		t.Error(err)
	}
}
