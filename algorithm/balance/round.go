package balance

func init() {
	RegisterBalance("round", &RoundRobinBalance{})
}
// 轮询
type RoundRobinBalance struct {
	currIndex int
}

func (r *RoundRobinBalance) DoBalance(is [] *Instance, key ...string) (inst *Instance, err error) {
	if len(is) == 0 {
		err = NoFoundErr
		return
	}

	lens := len(is)
	if r.currIndex >= lens {
		r.currIndex = 0
	}
	inst = is[r.currIndex]
	r.currIndex++
	return
}
