package cn

type Chinese struct {
}

func (e *Chinese) Hello() string {
	return "你好"
}

var avoidOpt interface{}

func init() {
	avoidOpt = &Chinese{}
}
