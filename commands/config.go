package commands

type config struct {
	next     string
	previous string
}

var globalConfig config = config{}
