package common

import (
	"log"
	"net/rpc"
)

/*
rpc对象
1 可导出
2 包含的方法必须接收两个参数,两个参数是可导出类型或者内建类型
3 方法的第二个参数必须是指针
4 方法只有一个返回值,返回值是error接口类型
5 可以认为第一个入参是远程调用的传入参数,第二个参数是返回值参数
*/

type Rect struct {
}

type Parameter struct {
	Length int
	Width  int
}
type RpcService interface {
	Area(*Parameter, *int) error
}

func RpcServiceRegisterService(service RpcService) {
	err := rpc.Register(service)
	if err != nil {
		log.Fatalf("rpc register fail:<%s>", err.Error())
	}
}
func (r *Rect) Area(p *Parameter, result *int) error {
	*result = p.Length * p.Width
	return nil
}
func (r *Rect) Perimeter(p *Parameter, result *int) error {
	*result = 2 * (p.Length + p.Width)
	return nil
}
