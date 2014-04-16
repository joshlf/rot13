// Copyright 2014 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"os"

	"github.com/joshlf13/rot13"
)

func main() {
	io.Copy(os.Stdout, rot13.NewReader(os.Stdin))
}
