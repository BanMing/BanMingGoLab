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

Interface:
与c#不同的是，如果一个类里实现了一个接口里的所有方法，那么就说明这个类是继承于这个类。

所有类型都实现了一个空类型的接口，所以可以把这个空的接口想做c#中的object或者是泛型

Comma-ok断言
value,ok:=element.(T)
这个有点像tryget一个参数的类型,value是element的值，T是需要判断的类型，ok是返回是否是指定类型T