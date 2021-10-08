package api

import "github.com/danbrough/repotest/misc"

func createLog(sink misc.LogSink) misc.Log {
  return misc.Logger{Name: "KIPFS_GO", Skip: 2, LogSink: sink}
}
