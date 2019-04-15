package singleton

import "fmt"

type singletObjectNonSafety struct {
}

func (this *singletObjectNonSafety) Output() {
	fmt.Println("我是SingletonObjectNonSafety,效率高，非线程安全")
}

var instanceNonSafety *singletObjectNonSafety

func GetInstanceNonSafety() *singletObjectNonSafety {
	if instanceNonSafety == nil {
		instanceNonSafety = &singletObjectNonSafety{}
	}
	return instanceNonSafety
}
