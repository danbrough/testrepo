package ipfs

import (
  "context"
  logging "github.com/ipfs/go-log"
  loggables "github.com/libp2p/go-libp2p-loggables"
  "github.com/danbrough/testrepo/misc"
  "math/rand"
  "time"
)

var log = logging.Logger("kipfs")

func StartIPFS(debugLog misc.Log) int {
  debugLog.Info("StartIPFS ...")
  rand.Seed(time.Now().UnixNano())
  ctx := logging.ContextWithLoggable(context.Background(), loggables.Uuid("session"))

  debugLog.Debug("Created ctx: %p", ctx)
  return 0
}
