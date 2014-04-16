// Copyright 2014 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rot13

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

var strings = []pair{
	pair{"123", "123"},
	pair{"abc", "nop"},
	pair{"ABC", "NOP"},
	pair{"123abc", "123nop"},
	pair{"123ABC", "123NOP"},
	pair{"a123Bc", "n123Op"},
	pair{string([]byte{0x80, 'a'}), string([]byte{0x80, 'n'})},
}

type pair struct {
	before, after string
}

func (p pair) testByteRead(t *testing.T) {
	strBytes := []byte(p.before)
	buf := bytes.NewBuffer(strBytes)
	r := NewReader(buf)

	dst, err := ioutil.ReadAll(r)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(dst, []byte(p.after)) {
		t.Errorf("Expected \"%s\"; got \"%s\"", p.after, string(dst))
	}
}

func (p pair) testByteWrite(t *testing.T) {
	strBytes := []byte(p.before)
	var buf bytes.Buffer
	w := NewWriter(&buf)

	_, err := writeAll(strBytes, w)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	dst := buf.Bytes()

	if !reflect.DeepEqual(dst, []byte(p.after)) {
		t.Errorf("Expected \"%s\"; got \"%s\"", p.after, string(dst))
	}
}

func (p pair) testRuneRead(t *testing.T) {
	strBytes := []byte(p.before)
	buf := bytes.NewBuffer(strBytes)
	r := NewRuneReader(buf)

	dst, err := readAllRunes(r)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(dst, []rune(p.after)) {
		t.Errorf("Expected \"%s\"; got \"%s\"", p.after, string(dst))
	}
}

func (p pair) testRot13(t *testing.T) {
	strBytes := []byte(p.before)
	for i, v := range strBytes {
		strBytes[i] = Rot13(v)
	}

	if !reflect.DeepEqual(strBytes, []byte(p.after)) {
		t.Errorf("Expected \"%s\"; got \"%s\"", p.after, string(strBytes))
	}
}

func (p pair) testRot13Bytes(t *testing.T) {
	strBytes := []byte(p.before)
	Rot13Bytes(strBytes)

	if !reflect.DeepEqual(strBytes, []byte(p.after)) {
		t.Errorf("Expected \"%s\"; got \"%s\"", p.after, string(strBytes))
	}
}

func (p pair) testRot13Rune(t *testing.T) {
	strRunes := []rune(p.before)
	for i, v := range strRunes {
		strRunes[i] = Rot13Rune(v)
	}

	if !reflect.DeepEqual(strRunes, []rune(p.after)) {
		t.Errorf("Expected \"%s\"; got \"%s\"", p.after, string(strRunes))
	}
}

func (p pair) testRot13Runes(t *testing.T) {
	strRunes := []rune(p.before)
	Rot13Runes(strRunes)

	if !reflect.DeepEqual(strRunes, []rune(p.after)) {
		t.Errorf("Expected \"%s\"; got \"%s\"", p.after, string(strRunes))
	}
}

func TestByteReader(t *testing.T) {
	for _, p := range strings {
		p.testByteRead(t)
	}
}

func TestByteWriter(t *testing.T) {
	for _, p := range strings {
		p.testByteWrite(t)
	}
}

func TestRuneReader(t *testing.T) {
	for _, p := range strings {
		p.testRuneRead(t)
	}
}

func TestRot13(t *testing.T) {
	for _, p := range strings {
		p.testRot13(t)
	}
}

func TestRot13Bytes(t *testing.T) {
	for _, p := range strings {
		p.testRot13Bytes(t)
	}
}

func TestRot13Rune(t *testing.T) {
	for _, p := range strings {
		p.testRot13Rune(t)
	}
}

func TestRot13Runes(t *testing.T) {
	for _, p := range strings {
		p.testRot13Runes(t)
	}
}

func writeAll(p []byte, w io.Writer) (int, error) {
	n := 0
	for len(p) > 0 {
		nn, err := w.Write(p)
		n += nn
		if err != nil {
			return n, err
		}
		p = p[nn:]
	}

	return n, nil
}

func readAllRunes(r io.RuneReader) ([]rune, error) {
	p := make([]rune, 0)
	for {
		rn, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return p, nil
			}
			return p, err
		}
		p = append(p, rn)
	}

	return p, nil
}
