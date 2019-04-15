package main

import (
	"designPatter/prototype"
	"fmt"
)

func main() {
	personalInfoOne := prototype.PersonalInfo{"小明", "男", 20}
	workExperienceOne :=prototype.WorkExperience{"2010-2012", "xxx公司"}
	resumeShallow := &prototype.Resume{"游泳",personalInfoOne,workExperienceOne}

	resumeShallowCopy :=resumeShallow.CloneByShallow()
	resumeShallowCopy.SetSkill("跳高")
	resumeShallowCopy.SetPersonalInfo("小红","女",18)
	resumeShallowCopy.SetWorkExperience("2011-2032","YYY公司")

	fmt.Println("浅拷贝输出对比：")
	resumeShallow.Display();
	fmt.Println("拷贝后：")
	resumeShallowCopy.Display();
	fmt.Println("")

    personalInfoTwo := prototype.PersonalInfo{"小黑","男",30}
	workExperienceTwo :=prototype.WorkExperience{"2015-2018","AAA公司"}
	resumeDeep := &prototype.Resume{"羽毛球",personalInfoTwo,workExperienceTwo}

	resumeDeepCopy :=resumeDeep.CloneByDeep()
	resumeDeepCopy.SetSkill("乒乓球")
	resumeDeepCopy.SetPersonalInfo("小美","女",27,)
	resumeDeepCopy.SetWorkExperience("2016-2019","BBB公司")

	fmt.Println("深拷贝拷贝输出对比：")
	resumeDeep.Display()
	fmt.Println("拷贝后：")
	resumeDeepCopy.Display()
}