package options

import "os"

type Options struct {
	InstanceDetailEndpoint string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) ApplyAndValidate() error {
	o.InstanceDetailEndpoint = os.Getenv("INSTANCE_DETAIL_ENDPOINT")
	if o.InstanceDetailEndpoint == "" {
		o.InstanceDetailEndpoint = "https://price.cloudpilot.ai"
	}
	return nil
}
