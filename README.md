# Desc

这个库提供了对json类型（或者说类似的）数据结构的一个封装，封装的作用就是你不用定义里边的每一个数据的类型，就是提供了像
js那种弱类型的处理方式。核心的点就是这个库自己实现了Map对象，来model json类似的数据结构，

    json -to-> Map -to->json # 提供了这种转换

核心的API，...

    * FromJSON， FromURLQuery
    * JSON, URLQuery 




# objx

  * Jump into the [API Documentation](http://godoc.org/github.com/stretchr/objx)


## files

    ├── accessors.go                     # implementation of retrieve value by key.
    ├── codegen
    │   ├── array-access.txt
    │   ├── index.html
    │   ├── template.txt
    │   └── types_list.txt
    ├── constants.go                      # define some constants
    ├── conversions.go                    # encoding Map object to various target
    ├── doc.go                            # provide overall package documentation
    ├── map.go                            # a custom map implemenation. type Map = map[string] interface{}, all constructors specified in this file
    ├── mutations.go                      # provide some method to create a new Map out of exsiting one
    ├── security.go                       # provide a simple implementation for hashing key & data
    ├── tests.go                          # test key in Map or not ts_test.go
    ├── type_specific_codegen.go          # 这个文件对go 语言中的每种类型都做了一次封装, 看看2k的代码行数，
    │                                     # 就知道go语言中缺失泛型，导致的问题是多么严重了，代码的重复率太高了
    │                                     # 这个文件看懂一个type的实现就可以了，其他都是类似的
    ├── value.go                          # 封装了 Value 类型的一些方法