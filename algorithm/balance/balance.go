package balance

import "errors"

var (
	NoFoundErr = errors.New("no found instance")
)

type Balance interface {
	// 负载均衡算法
	DoBalance([]*Instance, ...string) (*Instance, error)
}