// Copyright 2014 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rot13 provides functionality for rot13-ing alphabetic
// ASCII characters in ASCII and UTF-8 text.
package rot13

import "io"

type byteReader struct {
	r io.Reader
}

type byteWriter struct {
	w io.Writer
}

type runeReader struct {
	r io.RuneReader
}

func (r byteReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)

	for i := 0; i < n; i++ {
		p[i] = Rot13(p[i])
	}

	return n, err
}

func (w byteWriter) Write(p []byte) (int, error) {
	q := make([]byte, len(p))

	for i, b := range p {
		q[i] = Rot13(b)
	}

	n, err := w.w.Write(q)

	return n, err
}

func (r runeReader) ReadRune() (rune, int, error) {
	rn, n, err := r.r.ReadRune()
	return Rot13Rune(rn), n, err
}

// If b is an alphabetic character,
// Rot13 rotates it 13 alphabetic places forward
// (wrapping around if necessary).
func Rot13(b byte) byte {
	switch {
	case b >= 'A' && b <= 'Z':
		return 'A' + (((b - 'A') + 13) % 26)
	case b >= 'a' && b <= 'z':
		return 'a' + (((b - 'a') + 13) % 26)
	default:
		return b
	}
}

// If r is an alphabetic ASCII character,
// Rot13Rune rotates it 13 alphabetic places
// forward (wrapping around if necessary).
func Rot13Rune(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (((r - 'A') + 13) % 26)
	case r >= 'a' && r <= 'z':
		return 'a' + (((r - 'a') + 13) % 26)
	default:
		return r
	}
}

// Rot13Bytes applies Rot13 to each byte in b.
func Rot13Bytes(b []byte) {
	for i, v := range b {
		b[i] = Rot13(v)
	}
}

// Rot13Runes applies Rot13Rune to each rune in r.
func Rot13Runes(r []rune) {
	for i, v := range r {
		r[i] = Rot13Rune(v)
	}
}

// NewReader returns an io.Reader whose Read method
// calls r's Read method and rot13's the returned
// bytes.
func NewReader(r io.Reader) io.Reader {
	return byteReader{r}
}

// NewWriter returns an io.Writer whose Write method
// rot13's the given bytes before passing them to
// w's Write method.
func NewWriter(w io.Writer) io.Writer {
	return byteWriter{w}
}

// NewReader returns an io.Reader whose ReadRune method
// calls r's ReadRune method and rot13's the returned rune.
func NewRuneReader(r io.RuneReader) io.RuneReader {
	return runeReader{r}
}
