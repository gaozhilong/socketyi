package socketyi

import (
	
)

type Config struct {
	MaxConnection int
	QueueLength int
	ReadBufferSize int
	Heartbeat int64
	Timeout int64
}

var Default = Config{
	MaxConnection: 0,
	QueueLength: 9,
	ReadBufferSize: 1024,
	Heartbeat: 10e9,
	Timeout: 10e9,
}