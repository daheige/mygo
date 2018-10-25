# RabbitMQ安装步骤
    虚拟机：VMware workstation 12.0
    Linux系统：CentOS 7.0
    安装Erlang环境
    由于RabbitMQ是采用Erlang编写的，首先需要安装该语言库，以便运行代理服务器，可以参考Erlang官方文档。

# erlang-solution配置信息安装
    wget http://packages.erlang-solutions.com/erlang-solutions-1.0-1.noarch.rpm
    sudo rpm -Uvh erlang-solutions-1.0-1.noarch.rpm
    rpm --import http://packages.erlang-solutions.com/rpm/erlang_solutions.asc
# 第三方yum源依赖
    wget http://packages.sw.be/rpmforge-release/rpmforge-release-0.5.2-2.el6.rf.x86_64.rpm
    rpm –import http://apt.sw.be/RPM-GPG-KEY.dag.txt
    sudo rpm -i rpmforge-release-0.5.2-2.el6.rf.*.rpm
# 安装erlang
    sudo yum install erlang
# 运行erl命令进行测试
    安装RabbitMQ
    首先下载最新版的RabbitMQ
    wget http://www.rabbitmq.com/releases/rabbitmq-server/v3.6.1/rabbitmq-server-3.6.1-1.noarch.rpm
    使用rpm和yum进行安装
    rpm --import http://www.rabbitmq.com/rabbitmq-signing-key-public.asc
    yum install rabbitmq-server-3.6.1-1.noarch.rpm
# 启动RabbitMQ管理插件，用于web界面管理
    rabbitmq-plugins enable rabbitmq_management
    service rabbitmq-server restart
# 测试安装完成的RabbitMQ
    rabbitmqctl status
    具体内容可以参考RabbitMQ安装官方文档。

    注意：在Ubuntu下安装RabbitMQ非常简单，系统已经默认安装Erlang环境，使用apt-get install rabbitmq-server即可安装。

# RabbitMQ使用和管理
    后台操作命令管理RabbitMQ
    rabbitmqctl是RabbitMQ中间件的一个命令行管理工具，原理就是通过连接一个中间件的节点执行所有的动作，本地节点默认为“rabbit”，rabbitmqctl来指定RabbitMQ中间件在本地节点rabbit@localhost进行管理操作。

    注意：在使用rabbitmqctl命令时，可以用-n标志来明确指定的节点，比如rabbitmqctl -n rabbit@localhost …，在使用默认节点的情况下，这个可以省略。

# 应用管理

    rabbitmqctl status //显示RabbitMQ中间件的所有信息
    rabbitmqctl stop //停止RabbitMQ应用，关闭节点
    rabbitmqctl stop_app //停止RabbitMQ应用
    rabbitmqctl start_app //启动RabbitMQ应用
    rabbitmqctl restart //重置RabbitMQ节点
    rabbitmqctl force_restart //强制重置RabbitMQ节点
# 用户管理
    rabbitmqctl add_user username password //添加用户
    rabbitmqctl delete_user username //删除用户
    rabbitmqctl change_password username newpassword //修改密码
    rabbitmqctl list_users //列出所有用户
# 权限控制管理
     rabbitmqctl add_vhost vhostpath //创建虚拟主机
     rabbitmqctl delete_vhost vhostpath //删除虚拟主机
     rabbitmqctl list_vhosts //列出所有虚拟主机
     rabbitmqctl set_permissions [-p vhostpath] username <conf> <write> <read> //设置用户权限
     rabbitmqctl clear_permissions [-p vhostpath] username //删除用户权限
     rabbitmqctl list_permissions [-p vhostpath] //列出虚拟机上的所有权限
     rabbitmqctl list_user_permissions username //列出用户权限
# 集群管理
    rabbitmqctl cluster_status //获得集群配置信息
    rabbitmqctl join_cluster rabbit@localhost --ram | --disc //加入到rabbit节点中，使用内存模式或者磁盘模式
    rabbitmqctl change_cluster_node_type disc | ram //修改存储模式
    rabbitmqctl set_cluster_name newname //修改名字
# 查看管理
    rabbitmqctl list_queues [-p <vhostpath>]  //查看所有队列
    rabbitmqctl list_exchanges [-p <vhostpath>] //查看所有交换机
    rabbitmqctl list_bindings [-p <vhostpath>] //查看所有绑定
    rabbitmqctl list_connections //查看所有连接
    rabbitmqctl list_channels //查看所有信道
    rabbitmqctl list_consumers //查看所有消费者信息
# Web界面管理RabbitMQ
    RabbitMQ通过使用RabbitMQ Management 插件的Web界面来管理用户、队列和交换器。

# Web界面包含的内容

    服务器数据统计概览
    导入/导出服务器配置
    监控服务器连接
    信道列表
    交换器列表、添加交换器
    队列列表、添加队列
    修改队列绑定
    用户列表、添加用户
    查看vhost、添加vhost
    注意：使用rabbitmq-plugins enable rabbitmq_management来启动Management插件。 默认是可以本地登录localhost:15672，用户名：guest；密码：guest；端口默认15672。

# CLI管理
    在web界面上还有两个选项，HTTP API和CLT。

    HTTP API：
        提供了一个关于REST接口的文档界面，Web界面可以完成的功能，都可以通过使 
        用curl并调用API命令来完成。比如需要列出服务器上的vhost的话，在终端执行下列代码即可：
        curl -i -u guest:guest http://localhost:15672/api/vhosts  
    CLI：
        主要是Python脚本，相比于REST的API好处是，不需要手工编写请求，rabbitmqadmin会包装REST API，使用干净的接口与其交互，举例来说：
        curl -i -u guest:guest http://localhost:15672/api/queues //使用REST API
        ./rabbitmqadmin list queues   //使用CLI
    rabbitmqadmin脚本安装
        wget http://localhost:15672/cli/rabbitmqadmin
        chmod +x rabbitmqadmin
        $ chmod +x rabbitmqadmin
        $ sudo mv rabbitmqadmin /usr/bin/

# 总结
    三种管理方式各有特点：

    1. Web UI对于日常的开发更加简单，可以通过视图查看服务器的状态，方便观察。
    2. REST API可以自动化这些任务，并通过curl来调用，得到JSON对象后就可以集成到当前的工具和语言中。
    3. rabbitmqadmin脚本不需要手工构造通过curl发送的HTTP请求，获得更加简介易懂的格式化输出，帮助管理和监控RabbitMQ。
