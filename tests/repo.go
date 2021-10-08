package main

import (
  "github.com/danbrough/repotest/misc"
  "github.com/danbrough/repotest/repo"
  "github.com/danbrough/repotest/testing"
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
