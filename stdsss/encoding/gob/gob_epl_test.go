package gob_test

import (
	"bytes"
	"encoding/gob"
	"testing"
)

// Transmit
// Export Fields Or UnExport Fields
type P struct {
	X    int
	Y    int
	Z    int
	Name string
}

// Reviver
// U Can Pointer And Omit Field
type Q struct {
	X    *int32
	Y    *int32
	Name string
}

// This example shows the basic usage of the package:
// 1. Create an encoder,
// 2. transmit some values
// 3. receive them with a decoder.
func TestNormalGob(t *testing.T) {
	// Network
	var network bytes.Buffer

	// Initialize the encoder and decoder.
	// Write to Network
	encoder := gob.NewEncoder(&network)

	// encode
	encoder.Encode(P{X: 1, Y: 2, Z: 3, Name: "GodYao"})

	encoder.Encode(P{X: 4, Y: 5, Z: 6, Name: "Mike"})

	// ----------------------------------------------------------------

	// Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.

	// ----------------------------------------------------------------

	// Initialize the decoder
	// Read From Network
	var q Q
	decoder := gob.NewDecoder(&network)

	// decode
	if err := decoder.Decode(&q); err != nil {
		t.Fatal(err)
	}
	t.Logf("%d, %d, %s", *q.X, *q.Y, q.Name)
	if err := decoder.Decode(&q); err != nil {
		t.Fatal(err)
	}
	t.Logf("%d, %d, %s", *q.X, *q.Y, q.Name)
}
