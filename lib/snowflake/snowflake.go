package snowflake

import (
	"kevinku/go-forum/config"
	"kevinku/go-forum/lib/logger"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

var node *snowflake.Node

func Init() {
	var err error
	if node, err = snowflake.NewNode(config.Cfg.Snowflake.MachineID); err != nil {
		logger.Logger.Panic("snowflake create new node failed", zap.Any("error", err))
	}
}

func NewID() (id int64) {
	return node.Generate().Int64()
}
