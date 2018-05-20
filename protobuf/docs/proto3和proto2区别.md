# proto3和proto2区别
    总的来说，proto3 比 proto2 支持更多语言但 更简洁。去掉了一些复杂的语法和特性，更强调约定而弱化语法。如果是首次使用 Protobuf ，建议使用 proto3 。

    1 在第一行非空白非注释行，必须写：

        syntax = "proto3";

    2 字段规则移除了 “required”，并把 “optional” 改名为 “singular”；

    3 在 proto2 中 required 也是不推荐使用的。proto3 直接从语法层面上移除了 required 规则。其实可以做的更彻底，把所有字段规则描述都撤销，原来的 repeated 改为在类型或字段名后加一对中括号。这样是不是更简洁？语言增加 Go、Ruby、JavaNano 支持；

    4 移除了 default 选项；

        在 proto2 中，可以使用 default 选项为某一字段指定默认值。在 proto3 中，字段的默认值只能根据字段类型由系统决定。也就是说，默认值全部是约定好的，而不再提供指定默认值的语法。

        在字段被设置为默认值的时候，该字段不会被序列化。这样可以节省空间，提高效率。

        但这样就无法区分某字段是根本没赋值，还是赋值了默认值。这在 proto3 中问题不大，但在 proto2 中会有问题。

        比如，在更新协议的时候使用 default 选项为某个字段指定了一个与原来不同的默认值，旧代码获取到的该字段的值会与新代码不一样。

        另一个重约定而弱语法的例子是 Go 语言里的公共/私有对象。Go 语言约定，首字母大写的为公共对象，否则为私有对象。所以在 Go 语言中是没有 public、private 这样的语法的。

    5 枚举类型的第一个字段必须为 0 ；这也是一个约定。

    6 移除了对分组的支持；

        分组的功能完全可以用消息嵌套的方式来实现，并且更清晰。在 proto2 中已经把分组语法标注为『过期』了。这次也算清理垃圾了。

    7 旧代码在解析新增字段时，会把不认识的字段丢弃，再序列化后新增的字段就没了；

        在 proto2 中，旧代码虽然会忽视不认识的新增字段，但并不会将其丢弃，再序列化的时候那些字段会被原样保留。

        我觉得还是 proto2 的处理方式更好一些。能尽量保持兼容性和扩展能力，或许实现起来也更简单。proto3 现在的处理方式，没有带来明显的好处，但丢掉了部分兼容性和灵活性。

    8 移除了对扩展的支持，新增了 Any 类型；

        Any 类型是用来替代 proto2 中的扩展的。目前还在开发中。

        proto2 中的扩展特性很像 Swift 语言中的扩展。理解起来有点困难，使用起来更是会带来不少混乱。

        相比之下，proto3 中新增的 Any 类型有点想 C/C++ 中的 void* ，好理解，使用起来逻辑也更清晰。

    9 增加了 JSON 映射特性；

        语言的活力来自于与时俱进。当前，JSON 的流行有其充分的理由。很多『现代化』的语言都内置了对 JSON 的支持，比如 Go、PHP 等。而 C++ 这种看似保罗万象的学院派语言，因循守旧、故步自封，以致于现出了式微的苗条。

# proto2字段修饰符
    required：字段必填。
    optional: 字段选填，不填就会使用默认值，默认数值类型的默认值为0，string类型为空字符串，枚举类型为第一个枚举值。
    repeated: 数组类型，可以放入多个类型实例,会转换为slice
    enum: 枚举类型,会转换为常量在常量前面加上enum的名称
    message: 会转换为struct结构体,在一个proto文件中可以存放多个message
             message内部也可以定义message，外部如需调用需要指明对应的层级关系
# import关键字
    //引入外部proto文件
    import "other.proto";
    //引入外部proto文件，并让引入了该文件的proto文件也能访问被引入类型。
    import public "other.proto";

# enum枚举类型
    protobuf可以定义枚举类型：
    enum EnumType {
        TYPEA = 0;
        TYPEB = 1;
        TYPEC = 2;
    }
    enum的每行字段都是一个枚举值，等号后面跟的是实际值，默认实际值是不能一样的，但只需要增加一个option配置就可以设置一样的值：

    enum EnumType {
    option allow_alias = true;
    TYPEA = 0;
    TYPEB = 0;
    TYPEC = 2;
    }

    proto3 enum可以有alias。enum的值不能为负数。

    enum EnumAllowingAlias {
        option allow_alias = true;

        UNKNOWN = 0;

        STARTED = 1;

        RUNNING = 1;

    }

# proto3 枚举默认值一定是0

# message结构体(自定义类型)
    自定义的message类型：

    message MessageType {
        repeated string str = 1;
    }

    //proto2语法 可选
    message CompositeType {
        optional MessageType message = 1;
    }
# double stream
    http://lib.csdn.net/article/go/68547

# grpc action
    https://smallnest.gitbooks.io/go-rpc-programming-guide/content/part1/grpc.html

# 参考
    https://blog.csdn.net/menggucaoyuan/article/details/43602915
