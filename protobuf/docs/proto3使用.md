# proto
    protobuf 是google开源的一个序列化框架，类似xml，json，最大的特点是基于二进制，比传统的XML表示同样一段内容要短小得多。还可以定义一些可选字段，用于服务端与客户端通信。
    
# install proto3
    安装方式1: source install
        下载网站: https://github.com/google/protobuf/releases 
        cd /usr/local/src
        sudo wget https://github.com/google/protobuf/releases/download/v3.5.1/protobuf-all-3.5.1.tar.gz
        sudo tar zxvf protobuf-all-3.5.1.tar.gz
        需要编译, 在新版的 PB 源码中，是不包含 .configure 文件的，需要生成，此时先执行 ./autogen.sh 
        脚本说明如下：
        # Run this script to generate the configure script and other files that will
        # be included in the distribution. These files are not checked in because they
        # are automatically generated.

        此时生成了 .configure 文件，可以开始编译了
        sudo ./configure --prefix=/usr/local/protobuf
        sudo make && make install

        安装完成后,查看版本:
        cd /usr/local/protobuf/bin
        ./protoc --version
            libprotoc 3.5.1
        
        建立软链接
            sudo ln -s /usr/local/protobuf/bin/protoc /usr/bin/protoc

    安装方式2: centos yum 安装
              ubuntu用apt安装(可能是proto2,升级到proto3,建议用源码方式安装)
                sudo apt install protobuf-c-compiler  protobuf-compiler

# 定义消息类型
    首先让我们来看一个非常简单的例子。假定我们有这样的需求，我们要定义一个搜索请求消息，每个消息都包含一个查询字符串，和你感兴趣的特定页面编号，以及每个页面的命中个数。
    syntax = "proto3";
    message SearchRequest {
        string query = 1;
        int32 page_number = 2;
        int32 result_per_page = 3;
    }
    文件的第一行指定了你使用的是proto3的语法：如果你不指定，protocol buffer 编译器就会认为你使用的是proto2的语法。这个语句必须出现在.proto文件的非空非注释的第一行。
    我们看到，搜索请求消息结构中定义指定了三个字段（name/value pairs）。每个字段都有一个名称和类型。

# 指定字段类型
    在上述例子中，所有的字段都是值类型：两个整形(page_number and result_per_page) 和一个字符串 (query)。然而，你也可以指定你的字段的组合类型，包括枚举和其他消息类型。

# 分配标识——tag
    如你所见，消息中的每个字段都有一个唯一的数字标识。这些标识用来在消息的二进制格式中识别你的字段，并且，一旦你的消息投入使用，这些标识就不应该再被修改。

    注意，标识是由1到15使用一个字节来编码，包括标识数字和字段类型（你可以在Protocol Buffer 编码中查看更多详细）。
    标识16到2047占用两个字节。所以你应该保留1到15，用作出现最频繁的消息类型的标识。记得为将来会继续增加并可能频繁出现的元素留一点儿标识区间，也就是说，不要一下子把1—15全部用完，为将来留一点儿哦。

    你可以指定的最小的标识数字是1，最大是228，或者536,870,911。你也不能使用19000 到 19999之间的数字(FieldDescriptor::kFirstReservedNumber through FieldDescriptor::kLastReservedNumber)，因为它们被Protocol Buffers保留使用——如果你在自己的.proto文件中使用了一个保留数字，protocol buffer 编译器将会提示。同样的，你不能使用任何之前保留的标识。

    1到15使用一个字节来编码，包括标识数字和字段类型（你可以在Protocol Buffer 编码中查看更多详细）；16到2047占用两个字节。因此定义proto文件时应该保留1到15，用作出现最频繁的消息类型的标识。记得为将来会继续增加并可能频繁出现的元素留一点儿标识区间，也就是说，不要一下子把1—15全部用完，为将来留一点儿。
    标识数字的合法范围：最小是1，最大是 229 - 1，或者536,870,911。
    另外，不能使用19000 到 19999之间的数字(FieldDescriptor::kFirstReservedNumber through FieldDescriptor::kLastReservedNumber)，因为它们被Protocol Buffers保留使用

# 指定字段规则
    消息字段可以是下边中的一种：
        1. singular (单个): 符合语法规则的消息包含零个或者一个这样的字段（最多一个）。
        2. repeated (重复)： 一个字段在合法的消息中可以重复出现一定次数（包括零次）。重复出现的值的次序将被保留。在proto3中，重复出现的值类型字段默认采用压缩编码。你可以在这里找到更多关于压缩编码的东西： Protocol Buffer Encoding。
        3. proto3不支持proto2中的required和optional
# proto3默认值
    字符串类型默认为空字符串

    字节类型默认为空字节

    布尔类型默认false

    数值类型默认为0值

    enums类型默认为第一个定义的枚举值，必须是0

# 添加更多消息类型
    多个消息类型可以定义在一个.proto文件中。这个在你定义多个关联的消息的时候非常有用，——这样，举个例子吧，如果你想定义你的搜索消息类型的响应消息格式，你可以在同一个.proto文件中添加如下的内容：
    message SearchRequest {
        string query = 1;
        int32 page_number = 2;
        int32 result_per_page = 3;
    }

    message SearchResponse {
    ...
    }

# 添加注释
    在你的 .proto 文件中添加注释, 使用C/C++-风格的 // 语法，像下边这样：
    message SearchRequest {
        string query = 1;
        int32 page_number = 2;  // Which page number do we want?
        int32 result_per_page = 3;  // Number of results to return per page.
    }

# 保留字段
    如果你通过删除整个字段更新了消息类型，或者将整个字段其注释掉，未来用户在编写新的类型的时候能够复用这些注释掉的标识数字。然而，这会引起一些严重的问题，如果他们后来加载了同一个.proto的旧版，包括数据损坏，安全隐私bug等等。一个确保这种问题不会发生的办法是，保留你要删除的字段的标识。Protocol buffer 编译器将会提示以后用户使用这些保留的字段标识。
    message Foo {
    reserved 2, 15, 9 to 11;
    reserved "foo", "bar";
    }
    注意: 你不要混淆同一个保留语句中的字段名称和标识。

# 你的.proto编译之后会生成什么？
    当你运行对一个.proto文件的编译之后，编译器会为你选择的语言生成代码。你在文件中描述的消息类型，包括获取和设置字段的值，序列化你的消息到一个输出流，以及从一个输入流中转换出你的消息。
    1. 对于C++,编译器会为每个.proto文件生成一个.h 和一个.cc的文件,为每一个给出的消息类型生成一个类。
    2. 对于Java，编译器会生成一个java文件，其中为每一个消息类型生成一个类，还有特殊的用来创建这些消息类实例的Builder类，
    3. Python有一点小不同——Python编译器生成一个模块儿，其中为每一个消息类型生成一个静态的描述器，在运行时，和一个metaclass一起使用来创建必要的Python数据访问类。
    4. 对于Go，编译器为每个消息类型生成一个.pb.go文件。

    代理类生成:
        C++, 每一个.proto 文件可以生成一个 .h 文件和一个 .cc 文件
        Java, 每一个.proto文件可以生成一个 .java 文件
        Python, 每一个.proto文件生成一个模块，其中为每一个消息类型生成一个静态的描述器，在运行时，和一个metaclass一起使用来创建必要的Python数据访问类
        Go, 每一个.proto生成一个 .pb.go 文件
        Ruby, 每一个.proto生成一个 .rb 文件
        Objective-C, 每一个.proto 文件可以生成一个 pbobjc.h 和一个pbobjc.m 文件
        C#, 每一个.proto文件可以生成一个.cs文件.

# 值类型
    值类型的消息字段可以是一下类型中的一种——这个表格展示了可以在.proto文件中使用的类型，以及自动生成的相应语言的类型：
        .proto Type	            说明	                        C++Type    Java Type   Python Type[2]  Go Type
        double		                                            double	    double	    float	        float64
        float		                                            float	    float	    float	        float32
        int32	    使用可变长度编码。对负数进行编码时比较低效
                    如果你的字段要使用负数值，请使用sint32来代替。	   int32	    int	        int	            int

        int64	    使用可变长度编码。对负数进行编码时比较低效
                    如果你的字段要使用负数值，请使用sint64来代替。	   int64	    long	    int/long[3]	    int64

        uint32	    使用可变长度编码	                           uint32	   int[1]	    int/long[3]	    uint32
        uint64	    使用可变长度编码	                           uint64	   long[1]	    int/long[3]	    uint64
# 官方字段类型对应
    .proto	    C++	        Java	    Python	        Go	    
    double	    double	    double	    float	        float64	
    float	    float	    float	    float	        float32
    int32	    int32	    int	        int	            int32
    int64	    int64	    long	    ing/long[3]	    int64
    uint32	    uint32	    int[1]	    int/long[3]	    uint32
    uint64	    uint64	    long[1]	    int/long[3]	    uint64
    sint32	    int32	    int	        int	            int32	
    sint64	    int64	    long	    int/long[3]	    int64
    fixed32	    uint32	    int[1]	    int	            uint32	
    fixed64	    uint64	    long[1]	    int/long[3]	    uint64  
    sfixed32	int32	    int	        int	            int32	    
    sfixed64	int64	    long	    int/long[3]	    int64   
    bool	    bool	    boolean	    boolean	        bool	    
    string	    string	    String	    str/unicode[4]	string
    bytes	    string	    ByteString	str	            []byte

# 字段修饰符
    字段修饰符：
    required: 值不可为空
    optional: 可选字段 proto3字段规则移除了 “required”，并把 optional 改名为 singular
    singular: 符合语法规则的消息包含零个或者一个这样的字段（最多一个）
    repeated: 一个字段在合法的消息中可以重复出现一定次数（包括零次）。
             重复出现的值的次序将被保留。在proto3中，重复出现的值类型字段默认采用压缩编码。
             你可以在这里找到更多关于压缩编码的东西： Protocol Buffer Encoding。
    默认值：  optional PhoneType type = 2 [default = HOME];
# 字段默认值
    1. strings, 默认值是空字符串（empty string）
    2. bytes, 默认值是空bytes（empty bytes）
    3. bools, 默认值是false
    4. numeric, 默认值是0
    5. enums, 枚举默认值是第一个枚举值（value必须为0）
    6. message fields, the field is not set. Its exact value is langauge-dependent. See the generated code guide for details.
    7. repeated fields，默认值为empty，通常是一个空list

# install go proto
    go get -u github.com/golang/protobuf/proto
    go get -u github.com/golang/protobuf/protoc-gen-go

# go 文件生成
    下面这几种方式生成都可以：

    $ protoc --go_out=./go/ ./proto/helloworld.proto

    $ protoc --go_out=./go/ -I proto ./proto/helloworld.proto

        这次多了一个参数 -I ， -I=IMPORT_PATH  can be used as a short form of --proto_path.

        -IPATH, --proto_path=PATH   Specify the directory in which to search for imports.  May be specified multiple times; directories will be searched in order.  If not given, the current working directory is used.

        IMPORT_PATH specifies a directory in which to look for .proto files when resolving import directives. If omitted, the current directory is used. Multiple import directories can be specified by passing the --proto_path option multiple times; they will be searched in order.

    简单来说，就是如果多个proto文件之间有互相依赖，生成某个proto文件时，需要import其他几个proto文件，这时候就要用-I来指定搜索目录。

    如果没有指定 –I 参数，则在当前目录进行搜索。

# go grpc demo
    https://studygolang.com/articles/4370

