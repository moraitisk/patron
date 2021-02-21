package amqp

import (
	"errors"
	"time"

	"github.com/streadway/amqp"
)

// OptionFunc definition for configuring the component in a functional way.
type OptionFunc func(*Component) error

// Batching option for setting up batching.
// Allowed values for count is > 1 and timeout > 0.
func Batching(count uint, timeout time.Duration) OptionFunc {
	return func(c *Component) error {
		if count == 0 || count == 1 {
			return errors.New("count should be larger than 1 message")
		}
		if timeout <= 0 {
			return errors.New("timeout should be a positive number")
		}

		c.batchCfg.count = count
		c.batchCfg.timeout = timeout
		return nil
	}
}

// AMQPConfig option for setting AMQP configuration.
func AMQPConfig(cfg amqp.Config) OptionFunc {
	return func(c *Component) error {
		c.cfg = cfg
		return nil
	}
}
