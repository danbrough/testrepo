package main

import (
  "github.com/danbrough/testrepo/ipfs"
  "github.com/danbrough/testrepo/testing"
)

func main() {
  ipfs.StartIPFS(testing.TestLog)
}
