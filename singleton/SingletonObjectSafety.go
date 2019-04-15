package singleton

import "fmt"

type singletonObjectsafety struct {
}

func (this *singletonObjectsafety) Output() {
	fmt.Println("我是SingletonObjectSafety,线程安全，初始化的时候需要内存（浪费内存")
}

var instanceSafety *singletonObjectsafety = &singletonObjectsafety{}

func GetInstanceSafety() *singletonObjectsafety {
	return instanceSafety
}
