package singleton

import (
	"sync"
	"fmt"
)

type singletonObjectSelectSynSafety struct {

}

func (this *singletonObjectSelectSynSafety) Output() {
	fmt.Println("我是singletonObjectSelectSynSafety,线程安全，通过缩小加锁的范围提高效率")
}

var instanceSelectSynSafety *singletonObjectSelectSynSafety

var syn sync.Mutex

func GetInstanceSelectSynSafety() *singletonObjectSelectSynSafety {
	if instanceSelectSynSafety == nil {
		syn.Lock()
		defer syn.Unlock()
		if instanceSelectSynSafety == nil {
			instanceSelectSynSafety = &singletonObjectSelectSynSafety{}
		}
	}
	return instanceSelectSynSafety
}
