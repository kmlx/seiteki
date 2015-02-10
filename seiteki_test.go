/*
  Copyright 2015 Adrian Stanescu. All rights reserved.
  Use of this source code is governed by the MIT License (MIT) that can be found in the LICENSE file.
 */

package main

import (
  "net/http"
  "os/exec"
  "testing"
)

func TestStatic(t *testing.T) {
  exec.Command("/usr/local/go/bin/go", "run", "./static.go", "9999", ".")
  _, err := http.Get("http://127.0.0.1:9999/static.go")
  if err != nil {
    t.Errorf("0: Retrieval test failed")
  }
}
