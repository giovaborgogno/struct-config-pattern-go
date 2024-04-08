package pattern

type Opts struct {
	Attribute1 int
	Attribute2 string
	Attribute3 bool
}

type OptsFunc func(*Opts)

type CustomStruct struct {
	Opts
}
