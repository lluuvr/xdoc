package node

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Daemon start node daemon
func Daemon() bool {
	log.Debug("start node daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	if err := rpc.GetValidators(); err != nil {
		log.Fatal(err)
	}

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
