package prototype

import "fmt"

type Resume struct {
	Skill string
    PersonalInfo PersonalInfo
	WorkExperience WorkExperience
}

func (this *Resume) SetSkill(skill string)  {
	this.Skill=skill
}

func (this *Resume) SetPersonalInfo(name string,sex string,age int)  {
	this.PersonalInfo.Name=name
	this.PersonalInfo.Sex=sex
	this.PersonalInfo.Age=age
}

func (this *Resume) SetWorkExperience(timeArea string,company string)  {
	this.WorkExperience.TimeArea=timeArea
	this.WorkExperience.Company=company
}

func (this *Resume) Display() {
	fmt.Println("个人信息：",this.PersonalInfo.Name," ",this.PersonalInfo.Sex," ",this.PersonalInfo.Age)
	fmt.Println("工作经历:",this.WorkExperience.TimeArea," ",this.WorkExperience.Company )
	fmt.Println("技能：",this.Skill)
}

func (this *Resume) CloneByDeep() *Resume {
	resume :=*this
	return &resume
}

func (this *Resume) CloneByShallow() *Resume {
	resume :=this
	return resume
}