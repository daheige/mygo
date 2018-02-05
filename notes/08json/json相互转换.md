```
Go语言的JSON 库

　　Go语言自带的JSON转换库为 encoding/json
　　1.1）其中把对象转换为JSON的方法（函数）为 json.Marshal()，其函数原型如下

　　　　func Marshal(v  interface{}) ([]byte, error)

　　　　也就是说，这个函数接收任意类型的数据 v，并转换为字节数组类型，返回值就是我们想要的JSON数据和一个错误代码。当转换成功的时候，这个错误代码为nil

　　　　在进行对象转换为 JSON 的过程中，会遵循如下几条规则：

　　　　布尔型转换为 JSON 后仍是布尔型　， 如true -> true

　　　　浮点型和整数型转换后为JSON里面的常规数字，如 1.23 -> 1.23

　　　　字符串将以UTF-8编码转化输出为Unicode字符集的字符串，特殊字符比如<将会被转义为\u003c

　　　　数组和切片被转换为JSON 里面的数组，[]byte类会被转换为base64编码后的字符串，slice的零值被转换为null

　　　　结构体会转化为JSON对象，并且只有结构体里边以大写字母开头的可被导出的字段才会被转化输出，而这些可导出的字段会作为JSON对象的字符串索引

　　　　转化一个map 类型的数据结构时，该数据的类型必须是 map[string]T（T 可以是encoding/json 包支持的任意数据类型）



　　1.2）把 JSON 转换回对象的方法（函数）为 json.Unmarshal()，其函数原型如下

　　　　func Unmarshal(data [] byte, v interface{}) error

　　　　这个函数会把传入的 data 作为一个JSON来进行解析，解析后的数据存储在参数 v 中。这个参数 v 也是任意类型的参数（但一定是一个类型的指针），原因是我们在是以此函数进行JSON 解析的时候，这个函数不知道这个传入参数的具体类型，所以它需要接收所有的类型。

　　　　那么，在进行解析的时候，如果JSON 和 对象的结构不对口会发生什么呢，这就需要解析函数json.Unmarshal()遵循以下规则

　　　　json.Unmarshal() 函数会根据一个约定的顺序查找目标结构中的字段，如果找到一个即发生匹配。那什么是找到了呢？关于“找到了”又有如下的规则：假设一个JSON对象有个名为"Foo"的索引，要将"Foo"所对应的值填充到目标结构体的目标字段上，json.Unmarshal() 将会遵循如下顺序进行查找匹配

　　　　　　§ 一个包含Foo 标签的字段

　　　　　　§  一个名为Foo 的字段

　　　　　　§ 一个名为Foo 或者Foo 或者除了首字母其他字母不区分大小写的名为Foo 的字段。 这些字段在类型声明中必须都是以大写字母开头、可被导出的字段。

　　　　注意：如果JSON中的字段在Go目标类型中不存在，json.Unmarshal() 函数在解码过程中会丢弃该字段。

　　　　当JSON 的结构是未知的时候，会遵循如下规则：

　　　　　　§ JSON中的布尔值将会转换为Go中的bool类型

　　　　　　§ 数值会被转换为Go中的float64类型

　　　　　　§ 字符串转换后还是string类型

　　　　　　§ JSON数组会转换为[]interface{} 类型

　　　　　　§ JSON对象会转换为map[string]interface{}类型

　　　　　　§ null值会转换为nil

　　　　注意：在Go的标准库encoding/json包中，允许使用map[string]interface{}和[]interface{} 类型的值来分别存放未知结构的JSON对象或数组
```
