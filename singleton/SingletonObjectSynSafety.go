package singleton

import (
	"fmt"
)

type singletonObjectSynSafety struct {
}

func (this *singletonObjectSynSafety) Output() {
	fmt.Println("我是singletonObjectSynSafety,线程安全，效率低下(方法加锁)")
}

var instanceSysSafety *singletonObjectSynSafety

func GetInstanceSysSafety() *singletonObjectSynSafety {
	syn.Lock()
	defer syn.Unlock()
	if instanceSysSafety == nil {
		instanceSysSafety = &singletonObjectSynSafety{}
	}
	return instanceSysSafety
}
