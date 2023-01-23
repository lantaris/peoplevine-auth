package cmdargs

var (
	Args struct {
		Loggig string `arg:"-d,--debug" help:"Debug level"`
		Listen string `arg:"-l,--listen" help:"REST listen"`
	}
)
