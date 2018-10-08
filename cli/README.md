# go实现selpg

### go基础语法学习

* 变量声明
* 循环条件语句
* 文件读写
* 多线程

### flag包学习

[项目概览](https://pmlpml.github.io/ServiceComputingOnCloud/ex-cli-basic)

[flag入门](https://segmentfault.com/a/1190000014935402)

[flag和pflag](https://o-my-chenjian.com/2017/09/20/Using-Flag-And-Pflag-With-Golang/)

flag包其实就是一个把命令行输入参数读取解析的中间件

### selpg项目分析

[内有源码，非常详细](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)

#### 项目结构

其实使用flagp，把selpg.c翻译成go版本就完成了，项目结构如下：

##### flag初始化参数

##### 读入参数

##### 合法性检查

没解决第一个参数必须是-s，似乎flag没提供访问第一个参数的方法

##### 读入文件/命令行数据

错误处理未完成

##### 输出/写出

| < > 等涉及子进程的指令都是linux自带的

-d要求使用管道输出到打印机，需要新建管道给读取的代码使用，我直接赋值了别人的代码

### 测试

使用[这里](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)的指南进行测试

除了-d没有打印机测试，其余都完成

