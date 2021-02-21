package amqp

import (
	"testing"
	"time"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestAMQPConfig(t *testing.T) {
	cfg := amqp.Config{Locale: "123"}
	c := &Component{}
	assert.NoError(t, AMQPConfig(cfg)(c))
	assert.Equal(t, cfg, c.cfg)
}

func TestBatching(t *testing.T) {
	type args struct {
		count   uint
		timeout time.Duration
	}
	tests := map[string]struct {
		args        args
		expectedErr string
	}{
		"success":         {args: args{count: 2, timeout: 2 * time.Millisecond}},
		"invalid count":   {args: args{count: 1, timeout: 2 * time.Millisecond}, expectedErr: "count should be larger than 1 message"},
		"invalid timeout": {args: args{count: 2, timeout: -3}, expectedErr: "timeout should be a positive number"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := &Component{}
			err := Batching(tt.args.count, tt.args.timeout)(c)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.batchCfg.count, tt.args.count)
				assert.Equal(t, c.batchCfg.timeout, tt.args.timeout)
			}
		})
	}
}