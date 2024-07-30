package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

/*
*
字段个数、名称、类型完全相同可以强转
*/
type Person struct {
	Name   string
	Age    int
	School string
}

type Per Person

func main() {
	var t1 Teacher
	t1.Name = "李敬堂"
	t1.Age = 32
	t1.School = "安中医"
	fmt.Println(t1)

	var t2 Teacher = Teacher{"马老师", 45, "清华大学"}
	fmt.Println(t2)

	var t3 *Teacher = new(Teacher)
	(*t3).Name = "指针方式"
	t3.Age = 10
	fmt.Println(*t3)

	var t4 *Teacher = &Teacher{"嚣张", 1, "北大"}
	t4.Name = "寒气"
	fmt.Println(*t4)

	var p Person
	p = Person(t1)
	fmt.Println(p)

	var per Per = Per(p)
	fmt.Println(per)

}
