package main

import (
  "github.com/danbrough/testrepo/misc"
  "github.com/danbrough/testrepo/repo"
  "github.com/danbrough/testrepo/testing"
  "os"
  "path/filepath"
)

func repoCheck(log misc.Log) {
  log.Info("HOME: %s", os.Getenv("HOME"))
  var path = os.Getenv("HOME")
  var initialized = repo.RepoIsInitialized(path)
  log.Trace("%s is an initialized repo: %t", path, initialized)
  path = filepath.Join(path, ".ipfs")
  initialized = repo.RepoIsInitialized(path)
  log.Trace("%s is an initialized repo: %t", path, initialized)
  if !initialized {
    return
  }
  log.Warn("loading config ...")

}

func main() {
  repoCheck(testing.TestLog)
}
