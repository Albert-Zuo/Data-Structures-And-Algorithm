package sort

import (
	"errors"
	"fmt"
	"sync"
)

type SorterMgr struct {
	m map[string]Sorter

	mu sync.Mutex
	f  func(currType string, sortedSlice []int)
}

var (
	mgr = &SorterMgr{
		m: make(map[string]Sorter),
		f: func(currType string, sortedSlice []int) {
			fmt.Println(fmt.Sprintf("%s: %v", currType, sortedSlice))
		},
	}
	KeySet       = make(map[string]struct{})
	UndefinedErr = errors.New("this Sorter Undefined")
)

func (s *SorterMgr) register(typeName string, sort Sorter) {
	s.m[typeName] = sort
}

func RegisterSorter(typeName string, sort Sorter) {
	mgr.register(typeName, sort)
	KeySet[typeName] = struct{}{}
}

func MgrF(f func(currType string, sortedSlice []int)) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.f = f
}

func DoSort(typeName string, slice []int) error {
	return mgr.do(typeName, slice)
}

func (s *SorterMgr) do(typeName string, slice []int) (err error) {

	sorter, ok := s.m[typeName]
	if !ok {
		err = UndefinedErr
		return
	}
	sorter.Sort(slice)
	mgr.mu.Lock()

	defer func() {
		mgr.mu.Unlock()
		if err := recover(); err != nil {
			err = errors.New(fmt.Sprintf("panic err: %s", err))
		}
	}()
	s.f(typeName, slice)
	return
}


