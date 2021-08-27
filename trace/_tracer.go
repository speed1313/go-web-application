package trace

import (
	"fmt"
	"io"
)
/*
New()...create tracer instance
Off()...set niltracer (default)

*/

//Tracer is interface which record the event in code
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}


func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}){}
//OffはTraceメソッドの呼び出しを無視するTracerを返す.
func Off() Tracer{
	return &nilTracer{}
}