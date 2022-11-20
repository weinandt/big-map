package bigmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
