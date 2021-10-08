package api

import (
  "github.com/danbrough/testrepo/ipfs"
  "github.com/danbrough/testrepo/misc"
)

func StartIPFS(sink misc.LogSink) {
  ipfs.StartIPFS(createLog(sink))
}
