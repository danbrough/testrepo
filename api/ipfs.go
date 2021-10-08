package api

import (
  "github.com/danbrough/repotest/ipfs"
  "github.com/danbrough/repotest/misc"
)

func StartIPFS(sink misc.LogSink) {
  ipfs.StartIPFS(createLog(sink))
}
