package memstream

import (
	"bytes"
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	toWrite := []byte("abcdefg")

	buff := make([]byte, 100)
	s := New()

	// Empty read check
	n, err := s.Read(buff)
	if err != io.EOF {
		t.Errorf("Empty read failed, should have reported EOF. Error returned was %v", err)
	}
	if n != 0 {
		t.Errorf("Should have read 0 bytes, got %v", n)
	}

	// Write
	n, err = s.Write(toWrite)
	if err != nil {
		t.Errorf("Error writing: %v", err)
	}
	if n != len(toWrite) {
		t.Errorf("Did not write enough bytes (wrote %v, should have written %v)", n, len(toWrite))
	}

	// Rewind
	pos, err := s.Seek(0, 0)
	if err != nil {
		t.Errorf("Error rewinding: %v", err)
	}
	if pos != 0 {
		t.Errorf("Error rewinding, not at beginning (at %v)", pos)
	}

	// Read back data
	n, err = s.Read(buff)
	if err != nil && err != io.EOF {
		t.Errorf("Empty read failed: %v", err)
	}
	if n != len(toWrite) {
		t.Errorf("Should have read %v bytes, got %v", len(toWrite), n)
	}

	if !bytes.Equal(buff[:n], toWrite) {
		t.Errorf("Bytes do not match")
	}
}

func TestWrite(t *testing.T) {
	s := New()

	// Write until it expands
	singlebyte := make([]byte, 1)
	totalNeeded := DefaultCapacity + 10
	for totalWritten := 0; totalWritten < totalNeeded; {
		n, err := s.Write(singlebyte)
		totalWritten += n

		if err != nil {
			t.Errorf("Error writing: %v", err)
		}
		if n != 1 {
			t.Errorf("Did not write enough bytes (wrote %v, should have written %v)", n, 1)
			return
		}
	}

	// Rewind
	pos, err := s.Seek(0, 0)
	if err != nil {
		t.Errorf("Error rewinding: %v", err)
	}
	if pos != 0 {
		t.Errorf("Error rewinding, not at beginning (at %v)", pos)
	}

	// Read back data
	buff := make([]byte, DefaultCapacity*2)
	_ = "breakpoint"
	n, err := s.Read(buff)
	if err != nil && err != io.EOF {
		t.Errorf("Empty read failed: %v", err)
	}
	if n != totalNeeded {
		t.Errorf("Should have read %v bytes, got %v", totalNeeded, n)
	}
}

func TestSeek(t *testing.T) {
	toWrite := []byte("abcdefg")
	s := New()

	// Get position via seek
	pos, err := s.Seek(0, 1)
	if err != nil {
		t.Errorf("Error getting current position: %v", err)
	}
	if pos != 0 {
		t.Errorf("Position read via seek is not %v, got %v", len(toWrite), pos)
	}

	// Write
	n, err := s.Write(toWrite)
	if err != nil {
		t.Errorf("Error writing: %v", err)
	}
	if n != len(toWrite) {
		t.Errorf("Did not write enough bytes (wrote %v, should have written %v)", n, len(toWrite))
	}

	// Get position via seek
	pos, err = s.Seek(0, 1)
	if err != nil {
		t.Errorf("Error getting current position: %v", err)
	}
	if pos != int64(len(toWrite)) {
		t.Errorf("Position read via seek is not %v, got %v", len(toWrite), pos)
	}
}

func TestBytes(t *testing.T) {
	toWrite := []byte("abcdefg")
	s := New()

	// Write
	n, err := s.Write(toWrite)
	if err != nil {
		t.Errorf("Error writing: %v", err)
	}
	if n != len(toWrite) {
		t.Errorf("Did not write enough bytes (wrote %v, should have written %v)", n, len(toWrite))
	}

	written := s.Bytes()
	if !bytes.Equal(toWrite, written) {
		t.Error("Bytes() output did not match written bytes")
	}
}
