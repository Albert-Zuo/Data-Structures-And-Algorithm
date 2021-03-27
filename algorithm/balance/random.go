package balance

import (
	"math/rand"
)

func init()  {
	RegisterBalance("random",&RandomBalance{})
}

// 随机
type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(is [] *Instance, key...string) (inst *Instance, err error) {
	if len(is) == 0 {
		err = NoFoundErr
		return
	}
	lens := len(is)

	index := rand.Intn(lens)
	inst = is[index]

	return
}
