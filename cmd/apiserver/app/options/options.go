package options

type Options struct {
	APIQPS   float32
	APIBurst int
}

func NewOptions() *Options {
	return &Options{
		APIQPS:   50,
		APIBurst: 100,
	}
}

func (o *Options) ApplyAndValidate() error {
	return nil
}
