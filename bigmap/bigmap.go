package bigmap

import (
	"errors"
	"syscall"
)

func mmap(sizeInBytes int) ([]byte, error) {
	return syscall.Mmap(-1, 0, sizeInBytes, syscall.PROT_WRITE|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_SHARED)
}

type locator struct {
	index  int // Starting position for the string the the mmap'd array
	length int // Lenght of the string in the mmap'd array
}

type bigMap struct {
	lookup          map[string]locator
	storage         []byte
	currentPosition int // Only support adding strings to the dictionary for now.
}

func NewBigMap(backingSizeInBytes int) (*bigMap, error) {
	bytes, err := mmap(backingSizeInBytes)
	if err != nil {
		return nil, err
	}

	lookup := make(map[string]locator)
	newBigMap := &bigMap{
		lookup:          lookup,
		storage:         bytes,
		currentPosition: 0,
	}

	return newBigMap, nil
}

func (bm *bigMap) Add(key string, value string) error {
	// TODO: add checks to make sure string will fit.
	valueLen := len(value)
	location := locator{
		index:  bm.currentPosition,
		length: valueLen,
	}

	bm.lookup[key] = location

	// TODO: this can be done way more efficiently do avoid the double copy.
	newBytes := []byte(value)
	for i, byte := range newBytes {
		bm.storage[bm.currentPosition+i] = byte
	}

	// Adjust pointer so next insert is good to go.
	bm.currentPosition += valueLen

	return nil
}

func (bm *bigMap) Get(key string) (string, error) {
	location, ok := bm.lookup[key]
	if !ok {
		// TODO: should this return error or empty string, or comma okay idiom?
		return "", errors.New("Could not find key")
	}

	return string(bm.storage[location.index:location.length]), nil
}
