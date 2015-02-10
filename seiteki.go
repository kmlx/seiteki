/*
  Copyright 2015 Adrian Stanescu. All rights reserved.
  Use of this source code is governed by the MIT License (MIT) that can be found in the LICENSE file.
 */

/*
  seiteki
  Go program that serves files statically over http

  Usage:
  go run ./seiteki.go [listen port / 8080] [location to serve / .]
 */

package main

import (
  "log"
  "net/http"
  "os"
)

func main() {
  port          := "8080"
  configpath    := "."

  if len(os.Args) > 2 {
    port = os.Args[1]
    configpath = os.Args[2]
  } else {
    log.Printf("Usage: go run ./seiteki.go [listen port] [location to serve]")
  }

  // Simple static webserver
  log.Printf("Seiteki static http server is listening on port " + port + " and serving " + configpath)
  log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir(configpath))))
}
