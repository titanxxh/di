package en

type english struct {
}

func (e english) Hello() string {
	return "hello"
}

var avoidOpt interface{}

func init() {
	// use this symbol to avoid compiler optimization
	avoidOpt = english{}
}
