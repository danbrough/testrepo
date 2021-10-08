package main

import (
  "github.com/danbrough/repotest/ipfs"
  "github.com/danbrough/repotest/testing"
)

func main() {
  ipfs.StartIPFS(testing.TestLog)
}
