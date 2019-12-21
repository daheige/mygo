package main

import "log"

/**
1、结构体和匿名字段结构体组合（或指针类型）嵌入
2、结构体方法调用和方法重写
*/

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	log.Println("hello,", p.Name, " age: ", p.Age)
}

type Student struct {
	Person // 匿名字段
	Sex    int
}

// 重写Person method
func (s *Student) SayHello() {
	log.Println("rewrite SayHello method")
	log.Println("student name: ", s.Name, "your age: ", s.Age)
}

type Student2 struct {
	*Person // 匿名字段，以指针类型嵌入
	Sex     int
}

// 重写SayHello方法
func (s *Student2) SayHello() {
	s.Person.SayHello() // 调用Person上面的SayHello方法
	log.Println("s3 haha")
}

func main() {
	// s := Student{}
	var s Student
	s.Person.Name = "122"
	// s.Name = "daheige"
	s.Age = 23
	s.Sex = 1
	log.Println("s: ", s)
	log.Println("s name: ", s.Name)
	s.SayHello()

	//如果Person被Student以指针类型嵌入
	// 使用方式1
	// Person必须初始化
	var s2 Student2
	s2.Person = new(Person)
	s2.Person.Name = "fefe"
	log.Println("s2: ", s2.Name)
	s2.SayHello() //实际上是(&s2).SayHello() golang会自动解引用，就可以调用到SayHello方法

	//使用方式2
	s3 := Student2{
		Person: &Person{
			Name: "fefe",
			Age:  18,
		},
	}

	//这里修改下age
	s3.Age = 12

	log.Println("s3 name: ", s3.Name)
	log.Println("s3 age: ", s3.Age) //这里可以省略Person
	log.Println("s3 person age: ", s3.Person.Age)
	s3.SayHello()

	s4 := &Student2{}
	/**
		不初始化Person就会出现panic
		panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x497660]
	*/

	// s4.Person = new(Person)
	// s4.Name = "fefefe" //省略了Person
	// s4.Age = 28

	// 也可以先把Person赋值好，然后赋值给s4
	p := Person{
		Name: "fefefe",
		Age:  28,
	}

	s4.Person = &p

	s4.Sex = 2
	log.Println("s4 student: ", s4.Name, s4.Age, s4.Sex) //2019/12/21 21:54:14 s4 student:  fefefe 28 2

}

/**
$ go run struct_demo.go
2019/12/21 22:03:12 s:  {{122 23} 1}
2019/12/21 22:03:12 s name:  122
2019/12/21 22:03:12 rewrite SayHello method
2019/12/21 22:03:12 student name:  122 your age:  23
2019/12/21 22:03:12 s2:  fefe
2019/12/21 22:03:12 hello, fefe  age:  0
2019/12/21 22:03:12 s3 haha
2019/12/21 22:03:12 s3 name:  fefe
2019/12/21 22:03:12 s3 age:  12
2019/12/21 22:03:12 s3 person age:  12
2019/12/21 22:03:12 hello, fefe  age:  12
2019/12/21 22:03:12 s3 haha
2019/12/21 22:03:12 s4 student:  fefefe 28 2
*/
