package pattern2

func NewDefaultConfig() Config {
	return Config{
		Attribute1: 10,
		Attribute2: "default",
		Attribute3: false,
	}
}

func (c Config) WithAttribute1(n int) Config {
	c.Attribute1 = n
	return c
}

func (c Config) WithAttribute2(id string) Config {
	c.Attribute2 = id
	return c
}

func (c Config) WithAttribute3True() Config {
	c.Attribute3 = true
	return c
}

func NewCustomStruct() (*CustomStruct, error) {
	return &CustomStruct{
		config: NewDefaultConfig(),
	}, nil
}

func NewCustomStructWithConfig(c Config) (*CustomStruct, error) {
	return &CustomStruct{
		config: c,
	}, nil
}
