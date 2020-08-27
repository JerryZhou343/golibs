package envoy_tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnvoyTracer struct{}
var _ opentracing.Tracer= &EnvoyTracer{}

func (e EnvoyTracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	return nil
}

func (e EnvoyTracer) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	return nil
}

func (e EnvoyTracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	payload,ok := carrier.(map[string]string)
	if ok{
		for k,v := range payload{
			logrus.Info("key:%v value:%v",k,v)
		}
	}

	return nil,nil
}

