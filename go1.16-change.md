# go1.16变化
	
	主要体现在go fs,io,embed包等以及标准包的改变

# 支持arm64
	
	m1芯片可谓是最近的焦点，golang自然也不会落下。
	在golang1.16中官方已经支持darwin/arm64平台，cgo和编译成c语言可调用的动态/静态链接库的功能也已支持。
	同样受益的还有bsd家族的arm64版本,现在可以在新版mac上尝试golang了。
	不过plugin模式的支持仍在进行中，想要完整支持arm64还需要一段时间。

# GO111MODULE现在默认为on
	
	1.16开始默认启用modules，这在1.15的时候已经预告过了。现在GO111MODULE的默认值为on。
	不过golang还是提供了一个版本的适应期，如果你还不习惯modules，可以把GO111MODULE设置回auto。
	在1.17中这个环境变量将会被删除。

# go build不再更改mod相关文件

	以前的教程里我提到过go build会自动下载依赖，这会更新mod文件。

	现在这一行为被禁止了。想要安装、更新依赖只能使用go get命令，go build和go test将不会再做这类工作。

# go install的变化
	
	go install在1.16中也有了不小的变化。
	首先是通过go install my.module/tool@1.0.0 这样在module末尾加上版本号，可以在不影响当前mod的依赖的情况下安装golang程序。

	go install是未来唯一可以安装golang程序的命令，go get的编译安装功能现在可以靠-d选项关闭，而未来编译安装功能会从go get移除。

	也就是说go的命令各司其职，不再长臂管辖了。

# 新的GOVCS环境变量

	新的GOVCS环境变量指定了golang用什么版本控制工具下载源代码。

	其格式为：GOVCS=<module prefix>:<tool name>,[<module prefix>:<tool name>, ...]

	其中module prefix为github.com等，而tool name就是版本控制工具的名字，比如git，svn。

	一个更具体的例子是：GOVCS=github.com:git,evil.com:off,*:git|hg

	module prefix也可以用*通配任何模块的前缀。

	tool name还可以设置为all和off，all代表允许使用任何可用的工具，而off则表示不允许使用任何版本控制工具。

	不过现在设置为off的模块的代码仍然可能会被下载。

# 相对路径导入不在被允许
	
	golang1.16开始禁止import导入的模块以.开头，模块路径中也不允许出现任何非ASCII字符，所以下面的代码不再合法：
	```go
	import (
	    "./tools/factory"
	    "../models/user"
	    "some.pkg.com/杀马特/音乐工厂"
	)
	```

	对非ASCII字符一如既往的不友好，不过也只能按规矩办事了。

# 标准库的变化
	
	golang1.16除了对标准库进行通常的功能更新和修复，还引入了一些重大变化。

# testing
	
	testing包主要的变化是在测试用例里调用os.Exit(0)会从程序终止变成测试失败。

	比如这个：
	```go
	package main

	import (
	    "os"
	    "testing"
	)

	func TestXXX(t *testing.T) {
	    t.Log("exit")
	    os.Exit(0)
	}

	```
# ioutils包已经废弃
	
	1.16已经标记io/ioutil为废弃，函数被转移到了os和io这两个包里，具体见下表：
	| ioutil旧函数 | 	新函数 |
	| :---- 	  |  :----    |
	| Discard 	| io.Discard |
	| NopCloser | io.NopCloser| 
	| ReadAll 	| io.ReadAll |
	| ReadDir 	| os.ReadDir |
	| ReadFile 	| os.ReadFile |
	| WriteFile | os.WriteFile |
	| TempDir 	| os.MkdirTemp |
	| TempFile 	| os.CreateTemp |

	现在开始可以做移植了。

# tcp半连接队列扩容
	
	在Linux kernel 4.1以前，golang设置tcp的listen队列的长度是从/proc/sys/net/core/somaxconn获取的，通常为4096。

	而在4.1以后golang会直接设置半连接队列的长度为2^32 - 1也就是4294967295。

	更大的半连接队列意味着可以同时处理更多的新加入请求，而且不用再读取配置文件性能也会略微提升。

# 重大更新io/fs
	
	1.16除了支持嵌入静态资源外，最大的变化就是引入了io/fs包。

	golang认为文件的io操作是依赖于文件系统（filesystem，fs）的，所以决定模仿Linux的vfs做一套基于fs的io接口。

	这样做的目的有三个：
    1) os包应该专注于和系统交互而不是包含一部分io接口
    2) io包和os包分别包含了io接口的一部分，导致互相依赖职责不清晰
    3) 可以把有关联的一部分文件或者数据组成虚拟文件系统，供通用接口处理提升程序的可扩展性，比如zip打包的文件

    所以io/fs诞生了。

	fs包中主要包含了下面几种数据类型（都是接口类型）：
	| 名称 			| 作用   	|
	| :----  		| :----  	|
	| FS 			| 文件系统的抽象，有一个Open方法用来从FS打开获取文件数据 |
	| DirEntry 		| 描述目录项目（包含目录自身）的数据结构 |
	| File 			| 描述文件数据的结构，包含Stat，Read，Close方法 |
	| ReadDirFile 	| 在File的基础上支持ReadDir，可以代表目录自身 |
	| FileMode 		| 描述文件类型，比如是通常文件还是套接字或者是管道 |
	| FileInfo 		| 文件的元数据，例如创建时间等 |

	其中有一些接口和os包中的同名，实际上是os包引入fs包后起的别名。

	对于FS，还有以下的扩展，以便增量描述文件系统允许的操作：
	| 名称 			| 作用   	|
	| :----  		| :----  	|
	| GlobFS 			| 增加Glob方法，可以用通配符查找文件 |
	| ReadDirFS 		| 增加ReadDir方法，可以遍历目录 |
	| ReadFileFS 			| 增加ReadFile方法，可以用文件名读取文件所有内容 |
	| StatFS 	| 增加Stat方法，可以获得文件/目录的元信息 |
	| SubFS 		| 增加Sub方法，Sub方法接受一个文件/目录的名字，从这个名字作为根目录返回一个新的文件系统对象 |

	fs包还提供了诸如Glob，WalkDir等传统的文件操作接口。
	fs的主要威力在于处理zip、tar文件，以及http的文件接口时可以大幅简化代码。而且新的embed静态资源嵌入也是依赖fs实现的

# 可以嵌入静态资源（文件/文件目录）
	
	利用embed这个新的标准库。在声明静态资源的文件里我们需要引入这个库
	对于我们想要嵌入进程序的资源，需要使用//go:embed指令进行声明，注意//之后不能有空格。具体格式如下：

	//go:embed pattern
	// pattern是path.Match所支持的路径通配符

	| 通配符 			| 作用   	|
	| :----  		| :----  	|
	| ? 			| 代表任意一个字符（不包括半角中括号） |
	| * 		| 代表0至多个任意字符组成的字符串（不包括半角中括号） |
	| [...]和[!...] 			| 代表任意一个匹配方括号里字符的字符，!表示任意不匹配方括号中字符的字符 |
	| [a-z]、[0-9] 	| 代表匹配a-z任意一个字符的字符或是0-9中的任意一个数字 |
	| ** 		| 部分系统支持，*不能跨目录匹配，**可以，不过目前个golang中和*是同义词 |

	golang的embed默认的根目录从module的目录开始，路径开头不可以带/，不管windows还是其他系统路径分割副一律使用/。
	如果匹配到的是目录，那么目录下的所有文件都会被嵌入（有部分文件夹和文件会被排除，后面详细介绍），如果其中包含有子目录，
	则对子目录进行递归嵌入。

	对于一个完整的嵌入资源，代码中的声明是这样的：
	```go
	//go:embed images
	var imgs embed.FS

	//go:embed a.txt
	var txt []byte

	//go:embed b.txt
	var txt2 string

	```
	一共有三种数据格式可选：
	| 数据类型 			| 说明   	|
	| :----  		| :----  	|
	| []byte 			| 代表示数据存储为二进制格式，如果只使用[]byte和string需要以import (_ "embed")的形式引入embed标准库 |
	| string		| 表示数据被编码成utf8编码的字符串，因此不要用这个格式嵌入二进制文件比如图片，引入embed的规则同[]byte |
	| embed.FS 			| 表示存储多个文件和目录的结构，[]byte和string只能存储单个文件 |

	实际上接受嵌入文件数据的变量也可以是string和[]byte的类型别名或基于他们定义的新类型，例如下面的代码那样：
	```go
	type StringAlias = string

	//go:embed a.txt
	var text1 StringAlias

	type NewBytes []byte

	//go:embed b.txt
	var text2 NewBytes
	```
	处理单个文件
	```go
	package main

	import (
	    "fmt"
	    _ "embed"
	)

	//go:embed macbeth.txt
	var macbeth string

	//go:embed texts/en.txt
	var hello string

	func main() {
	    fmt.Println(len(macbeth)) // 麦克白的总字符数
	    fmt.Println(hello) // Output: Hello, world
	}
	```

	二进制文件的例子，embed_img.go如下所示：
	```go
	package main

	import (
	    "fmt"
	    _ "embed"
	)

	//go:embed imgs/screenrecord.gif
	var gif []byte

	//go:embed imgs/png/a.png
	var png []byte

	func main() {
	    fmt.Println("gif size:", len(gif)) // gif size: 81100466
	    fmt.Println("png size:", len(png)) // png size: 4958264
	}
	```

	处理多个文件和目录
	```go
	package main

	import (
	    "fmt"
	    "embed"
	)

	//go:embed texts
	var dir embed.FS

	// 两者没什么区别
	//go:embed texts/*
	var files embed.FS

	func main(){
	    zh, err := files.ReadFile("texts/zh.txt")
	    if err != nil {
	        fmt.Println("read zh.txt error:", err)
	    } else {
	        fmt.Println("zh.txt:", string(zh))
	    }

	    jp, err := dir.ReadFile("jp.txt")
	    if err != nil {
	        fmt.Println("read  jp.txt error:", err)
	    } else {
	        fmt.Println("jp.txt:", string(jp))
	    }
	    
	    jp, err = dir.ReadFile("texts/jp.txt")
	    if err != nil {
	        fmt.Println("read  jp.txt error:", err)
	    } else {
	        fmt.Println("jp.txt:", string(jp))
	    }
	}
	```

	我们想读取单个文件需要用ReadFile方法，它接受一个path字符串做参数，从中查找对应的文件然后返回([]byte, error)

	要注意的是文件路径必须要明确写出自己的父级目录，否则会报错，因为嵌入资源是按它存储路径相同的结构存储的，和通配符怎么指定无关。
	Open是和ReadFile类似的方法，只不过返回了一个fs.File类型的io.Reader，因此这里就不再赘述，
	需要使用Open还是ReadFile可以由开发者根据自身需求决定。

	embed.FS自身是只读的，所以我们不能在运行时添加或删除嵌入的文件，fs.File也是只读的，所以我们不能修改嵌入资源的内容。

	如果只是提供了一个查找读取资源的能力，那未免小看了embed。在golang1.16里任意实现了io/fs.FS接口的
	类型都可以表现的像是真实存在于文件系统中的目录一样，哪怕它其实是在内存里的类map数据结构。
	因此我们也可以像遍历目录一样去处理embed.FS

	```go
	package main

	import (
		"embed"
		"fmt"
	)

	// 更推荐直接用imgs去匹配
	//go:embed imgs/**
	var dir embed.FS

	// 遍历当前目录，有兴趣你可以改成递归版本的
	func printDir(name string) {
		// 返回[]fs.DirEntry
		entries, err := dir.ReadDir(name)
		if err != nil {
			panic(err)
		}

		fmt.Println("dir:", name)
		for _, entry := range entries {
			// fs.DirEntry的Info接口会返回fs.FileInfo，这东西被从os移动到了io/fs，接口本身没有变化
			info, _ := entry.Info()
			fmt.Println("file name:", entry.Name(), "\tisDir:", entry.IsDir(), "\tsize:", info.Size())
		}
		fmt.Println()
	}

	func main() {
		printDir("imgs")
		printDir("imgs/jpg")
		printDir("imgs/png")
	}
	```
	运行结果：
	```
	dir: imgs
	file name: jpg  isDir: true     size: 0
	file name: png  isDir: true     size: 0
	file name: screenrecord.gif     isDir: false    size: 81100466

	dir: imgs/jpg
	file name: a.jpg        isDir: false    size: 620419
	file name: b.jpg        isDir: false    size: 999162
	file name: c.jpg        isDir: false    size: 349725

	dir: imgs/png
	file name: a.png        isDir: false    size: 4958264
	file name: b.png        isDir: false    size: 1498303
	file name: c.png        isDir: false    size: 1751934
	```
	唯一和真实的目录不一样的地方是目录文件的大小，在ext4等文件系统上目录会存储子项目的元信息，所以大小通常不为0。
	如果想要内嵌整个module，则在引用的时候需要使用"."这个名字，但除了单独使用之外路径里不可以包含..或者.，换而言之，
	embed.FS不支持相对路径，把上面的代码稍加修改：

	```go
	package main

	import (
	    "fmt"
	    "embed"
	)

	//go:embed *
	var dir embed.FS

	func main() {
	    printDir(".")
	    //printDir("./texts/../imgs") panic: open ./texts/../imgs: file does not exist
	}
	```
	程序输出：
	```
	dir: .
	file name: embed_fs.go  isDir: false    size: 484
	file name: embed_img.go         isDir: false    size: 235
	file name: embed_img2.go        isDir: false    size: 187
	file name: embed_img_fs.go      isDir: false    size: 692
	file name: embed_text.go        isDir: false    size: 211
	file name: embed_text_fs.go     isDir: false    size: 603
	file name: go.mod       isDir: false    size: 30
	file name: imgs         isDir: true     size: 0
	file name: macbeth.txt  isDir: false    size: 100095
	file name: texts        isDir: true     size: 0
	```
	因为使用了错误的文件名或路径会在运行时panic，所以要格外小心。
    （当然//go:embed是在编译时检查的，而且同样不支持相对路径，同时也不支持超出了module目录的任何路径，
	比如go module在/tmp/proj，我们指定了/tmp/proj2）

	你也可以用embed.FS处理单个文件，但我个人认为单个文件就没必要再多包装一层了。

	由于是golang内建的支持，所以上述的代码无需调用任何第三方工具，也没有烦人的生成代码，
	不得不说golang对工程控制的把握上还是相当可靠的。

# 其他改进
	
	其他的改进包括Unicode更新到了13.0、新增加了runtime/metrics包已提供更好更规范的运行时信息等。

	同时1.16优化了链接器，现在它在linux/amd64上比1.15快了20-25%，内存占用减少了5-15%。

	在Windows上已经全面支持了地址空间布局随机化（ASLR），此前不支持将golang编译为dll时启用ASLR。

# 参考
	
	https://golang.google.cn/doc/devel/release.html#go1.16
	https://www.cnblogs.com/apocelipes/p/13907858.html
	https://www.cnblogs.com/apocelipes/p/14409273.html#go-install%E7%9A%84%E5%8F%98%E5%8C%96
