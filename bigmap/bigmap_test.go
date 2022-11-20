package bigmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestInsertAndGet(t *testing.T) {
	// Arrange
	testMap, err := NewBigMap(10)
	if err != nil {
		t.Fatal("could not create map.")
	}

	key := "key"
	expectedValue := "myVal"
	testMap.Add(key, expectedValue)

	// Act
	returnedVal, err := testMap.Get("key")
	if err != nil {
		t.Fatal("Could not get value from map.")
	}

	// Assert
	assert.Equal(t, expectedValue, returnedVal, "Strings should have matched.")
}

func generate500MBRandomString() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 500*1024*1024)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestInsert16Gigs(t *testing.T) {
	bigString := generate500MBRandomString()
	testMap, err := NewBigMap(16 * 1024 * 1024 * 1024)
	if err != nil {
		t.Fatal("could not create map.", err)
	}

	for i := 0; i < 16*2; i++ {
		err := testMap.Add(fmt.Sprint(i), bigString)
		if err != nil {
			t.Fatal("Could not insert a big string into the map.", err)
		}
	}
}
