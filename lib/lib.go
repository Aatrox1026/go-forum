package lib

import (
	"kevinku/go-forum/lib/logger"
	"kevinku/go-forum/lib/snowflake"
)

func Init() {
	logger.Init()
	snowflake.Init()
}
