package pckStruct

type Student struct {
	Name string
}

type stu struct {
	Name string
}

func NewStu(name string) *stu {
	return &stu{Name: name}
}
