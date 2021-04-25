package producer

type Producer interface {
	Produce() error
}

type defaultProducer struct {
}

func NewProducer() Producer {
	return &defaultProducer{}
}

func (p *defaultProducer) Produce() error {
	return nil
}
