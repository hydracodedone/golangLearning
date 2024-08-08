package common

import (
	"strings"

	"google.golang.org/grpc/metadata"
)

// 实现TextMapWriter接口
// type TextMapWriter interface {
// 	Set(key, val string)
// 	ForeachKey(handler func(key, val string) error) error
// }

// metadata 读写
type MDReaderWriter struct {
	metadata.MD
}

// 为了 opentracing.TextMapReader ，参考 opentracing 代码
func (c MDReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range c.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// 为了 opentracing.TextMapWriter，参考 opentracing 代码
func (c MDReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	c.MD[key] = append(c.MD[key], val)
}
