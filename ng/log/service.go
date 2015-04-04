package log

import "github.com/gopherjs/gopherjs/js"

type Service struct {
	*js.Object `ajs-service:"$log"`
}

func (p *Service) Service(format string, params ...interface{}) {
	p.Call("log", fmt.Sprintf(format, params...))
}

func (p *Service) Warn(format string, params ...interface{}) {
	p.Call("warn", fmt.Sprintf(format, params...))
}

func (p *Service) Info(format string, params ...interface{}) {
	p.Call("info", fmt.Sprintf(format, params...))
}

func (p *Service) Error(format string, params ...interface{}) {
	p.Call("error", fmt.Sprintf(format, params...))
}

func (p *Service) Debug(format string, params ...interface{}) {
	p.Call("debug", fmt.Sprintf(format, params...))
}
