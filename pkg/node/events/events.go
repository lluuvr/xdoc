package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

var ChSendNewBlock = make(chan *types.Block)

// Listen listen for events
func Listen() error {
	var (
		chReadNewBlock = make(chan *types.Block)
		brk            = env.GetBroker()
	)

	if err := brk.Subscribe(broker.SBlock, chReadNewBlock); err != nil {
		return err
	}

	if err := brk.Publish(broker.SBlock, ChSendNewBlock); err != nil {
		return err
	}

	for {
		select {
		case block, ok := <-chReadNewBlock:
			if !ok {
				return nil
			}

			log.Debug(block.Index)

			ChSendNewBlock <- block
		}
	}
}
