package consumer

type Consumer interface {
	Consume() error
}

type defaultConsumer struct {
}

func NewConsumer() Consumer {
	return &defaultConsumer{}
}

func (c *defaultConsumer) Consume() error {
	return nil
}
