安装：
https://www.cnblogs.com/ficow/p/6785905.html
权限：
https://blog.csdn.net/yemao_guyue/article/details/80575532

调试：
brew install go-delve/delve/delve
这个很详细：
https://www.jianshu.com/p/137854be2458


类java的package机制

go语言数据底层的存储：
https://research.swtch.com/godata

大写就是公有
小写就是私有

数组是为值类型
slice map 为引用类型
在函数传递中不需要取地址后传递指针，但是需要改变slice的长度是就需要取地址指针

new：返回指针
make:返回初始化后的（非零）值，内部构建，make只能创建slice、map和channel，
因为这个三个他们内部是由一个数据指针、长度和容量所组成，这几个在被初始化之前是nil。

int转字符串
strconv.Itoa()

switch中自带break