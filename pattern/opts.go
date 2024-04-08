package pattern

func defaultOpts() Opts {
	return Opts{
		Attribute1: 10,
		Attribute2: "default",
		Attribute3: false,
	}
}

func WithAttribute1(n int) OptsFunc {
	return func(opts *Opts) {
		opts.Attribute1 = n
	}
}

func WithAttribute2(id string) OptsFunc {
	return func(opts *Opts) {
		opts.Attribute2 = id
	}
}

func WithAttribute3() OptsFunc {
	return func(o *Opts) {
		o.Attribute3 = true
	}
}

func NewCustomStruct(opts ...OptsFunc) *CustomStruct {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}
	return &CustomStruct{
		Opts: o,
	}
}
