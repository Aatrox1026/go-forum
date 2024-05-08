package snowflake

import (
	"kevinku/go-forum/config"
	l "kevinku/go-forum/lib/logger"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

var cfg = config.Cfg
var logger = l.Logger

var node *snowflake.Node

func Init() {
	var err error
	if node, err = snowflake.NewNode(cfg.Snoeflake.MachineID); err != nil {
		logger.Panic("snowflake create new node failed", zap.Any("error", err))
	}
}

func NewID() (id int64) {
	return node.Generate().Int64()
}
