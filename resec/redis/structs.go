package redis

import (
	"context"
	"github.com/jpillora/backoff"
	"github.com/redis/go-redis/v9"
	"github.com/seatgeek/resec/resec/state"
	log "github.com/sirupsen/logrus"
)

type Manager struct {
	backoff        *backoff.Backoff // exponential backoff helper
	logger         *log.Entry       // logging specificall for Redis
	client         *redis.Client    // redis client
	config         *Config          // redis config
	state          *state.Redis     // redis state
	ctx            context.Context  // context
	stateCh        chan state.Redis // redis state channel to publish updates to the reconciler
	commandCh      chan Command
	stopCh         chan interface{}
	watcherRunning bool
}

type Config struct {
	Address string // address (IP+Port) used to talk to Redis
}
