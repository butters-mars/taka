package idgen

import (
	"github.com/bwmarrin/snowflake"

	"github.com/butters-mars/tiki/logging"
)

// SnowFlakeIDGen id generator using snowflake
type SnowFlakeIDGen struct {
	nodeID int64
	node   *snowflake.Node
}

// NewSnowFlakeIDGen -
func NewSnowFlakeIDGen(nodeID int64) (Service, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		logging.WError("cannot init snowflake idgen", "err", err)
		return nil, err
	}

	return &SnowFlakeIDGen{
		nodeID: nodeID,
		node:   node,
	}, nil
}

// GenID generate a snowflake ID.
func (sf *SnowFlakeIDGen) GenID() (int64, error) {
	id := sf.node.Generate()
	return int64(id), nil
}
