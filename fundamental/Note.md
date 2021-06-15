###make和new的区别
- make和new都是用来申请内存的
- new很少用，一般用来给基本数据类型申请内存，`string`, `int`等，一般返回的是对应类型的指针(*string, *int)
- make是用来给`slice`、`map`、`chan`申请内存的，make函数返回的是对应的这三个类型本身