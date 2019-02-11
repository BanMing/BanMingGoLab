http包包含：请求行line，头文件Header，主体Body

get与post区别：
1.get把需要传输的数据直接放在url中，直接加上？参数，&分隔参数，
post中的数据是放在http包中的body中。

2.get传输的数据是有限的，URL长度有限。post传输的数据是无限。

3.post比get更加安全一些。get参数是在URL中，可以让别人看见。
