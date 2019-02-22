/**
Go2 Error Inspection前瞻
通过xerror.Errorf对已知error变量进行包裹后
返回的error变量所归属的类型实现了error、xerrors.Formatter和xerrors.Wrapper接口
携带了被包裹(wrap)变量的信息，而传统的通过fmt.Errorf生成的error变量仅仅实现了error接口
没有被包裹的error变量的任何信息。
这些携带的信息将在后续error变量值测试时(Is和As)以及error变量信息展示时被充分利用
%+v 每个错误变量生成时的位置信息，这是因为xerrors.Errorf生成的变量类型:wrapError中携带了“位置信息(frame Frame)"

Is方法是在error chain上做值测试。
有些时候我们更方便做类型测试，即某一个err是否是某error类型。xerror提供了As方法：
	func As(err error, target interface{}) bool

As会将err中的error chain上的每个error type与target的类型做匹配
如果相同，则返回true，并且将匹配的那个error var的地址赋值给target
相当于通过As的target将error chain中类型匹配的那个error变量析出

//小结：
以上就是xerrors提供的有关Go2 error inspection机制的主要功能。
注意：xerrors及其proposal仍然可能会变动(包括设计和具体的实现)，因此这里不能保证本文demo示例在后续的版本中依然可以编译运行。
本文中的示例代码可以在这里得到。目前go官方的golang.org/x/exp repo中有两个版本的实现：
golang.org/x/exp/errors和golang.org/x/exp/xerrors。差别在于前者没有提供errors.Errorf。
*/

package main

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/xerrors"
)

func function3() error {
	return xerrors.New("fn3 error")
}

func function2() error {
	err := function3()
	if err != nil {
		return xerrors.Errorf("wrap2: %w", err)
	}

	return nil
}

func function1() error {
	err := function2()
	if err != nil {
		return xerrors.Errorf("wrap1: %w", err)
	}
	return nil
}

//fn1 包裹fn2,fn2包裹fn3
func main() {
	err := errors.New("fefe")
	log.Println("fefe")
	log.Println(err)

	myerr2 := xerrors.Errorf("daheige")
	log.Println(myerr2.Error())

	e := function1()

	if e != nil {
		log.Printf("%v", e)    //2019/02/22 22:21:34 wrap1: wrap2: fn3 error
		log.Println(e.Error()) //wrap1: wrap2: fn3 error
		log.Printf("%+v", e)   //采用%+v，将打印具体的错误信息
	}

	err1 := xerrors.New("1")
	err2 := xerrors.Errorf("wrap2 err: %w", err1)
	err3 := xerrors.Errorf("wrap3 err: %w", err2)
	log.Println("err3 is err1: ", xerrors.Is(err3, err1))
	log.Println("err2 is err1: ", xerrors.Is(err2, err1))
	log.Println("err3 is err2: ", xerrors.Is(err3, err2))
	/*
		运行结果
		2019/02/22 22:37:42 err3 is err1:  true
		2019/02/22 22:37:42 err2 is err1:  true
		2019/02/22 22:37:42 err3 is err2:  true
	*/
	erra := xerrors.New("a")

	log.Println("erra is err1: ", xerrors.Is(erra, err1)) //false

	hgerr1 := MyError{}

	hgerr2 := xerrors.Errorf("hg wrap 2: %w", hgerr1) //包裹了hgerr3
	hgerr3 := xerrors.Errorf("hg wrap 3: %w", hgerr2)

	log.Println("hgerr2 is hgerr3: ", xerrors.Is(hgerr2, hgerr3)) //false
	log.Println("hgerr3 is hgerr2: ", xerrors.Is(hgerr3, hgerr2)) //true
	var hgerr MyError

	b := xerrors.As(hgerr3, &hgerr)

	fmt.Println("hgerr3 as MyError? -> ", b)
	fmt.Println("hgerr is hgerr1? -> ", xerrors.Is(hgerr, hgerr1)) //true

	err4 := xerrors.Opaque(hgerr3)
	b = xerrors.As(err4, &hgerr)
	fmt.Println("err4 as MyError? -> ", b)
	b = xerrors.Is(err4, hgerr3)
	fmt.Println("err4 is hgerr3? -> ", b)
	/*
		hgerr3 as MyError? ->  true
		hgerr is hgerr1? ->  true
		err4 as MyError? ->  false
		err4 is hgerr3? ->  false
	*/
	/*
		As方法从hgerr3的error chain中匹配到MyError类型的hgerr1，并将hgerr1赋值给err变量析出。
		在后续的Is测试也证实了这一点。代码中还调用了xerrors的Opaque方法
		该方法将传入的支持unwrap操作的error变量转换为一个不支持unwrap的类型的error变量。
		在最后的对err4（通过Opaque调用得到)的测试我们也可以看到：err4无法匹配MyError type，与hgerr3的等值测试也返回false。
	*/

}

type MyError struct{}

func (MyError) Error() string {
	return "MyError"
}
