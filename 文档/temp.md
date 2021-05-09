# GO语言容器

## nil

nil 标识符是不能比较的:nil == nil报错
nil 并不是Go语言的关键字或者保留字，也就是说我们可以定义一个名称为 nil 的变量
nil 没有默认类型
不同类型 nil 的指针是一样的,地址都是0x0
不同类型的 nil 是不能比较的
相同类型的 nil 值也可能无法比较
     map、slice 和 function 类型的 nil 值不能比较
     也就是说,chan,interface的nil是可以比较的,且结果是true

    package main
    import "fmt"
    type some interface{}
    func main() {
        var temp1 chan int
        var temp2 chan int
        fmt.Println(temp1 == temp2)
        var temp3 some
        var temp4 some
        fmt.Println(temp3 == temp4)
    }
不同类型的 nil 值占用的内存大小可能是不一样的

## MAKE AND NEW

make 用于创建切片、哈希表和管道等内置数据结构，new 用于分配并创建一个指向对应类型的指针。
make 关键字的主要作用是创建Slice、Hash和 Channel 等内置的数据结构，而 new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针。
在编译期间的类型检查阶段，Go语言其实就将代表 make 关键字的 OMAKE 节点根据参数类型的不同转换成了 OMAKESLICE、OMAKEMAP 和 OMAKECHAN 三种不同类型的节点，这些节点最终也会调用不同的运行时函数来初始化数据结构。

## container.List

双链表

## sync.Map

Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。
无须初始化，直接声明即可。
sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false
sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。

## 接口

函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体，当类型方法被调用时，还需要调用函数本体
// 函数定义为类型
type FuncCaller func(interface{})
// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
    // 调用f()函数本体
    f(p)
}

## 反射

Go语言程序中，使用 reflect.TypeOf() 函数可以获得任意值的反射类型对象（reflect.Type）

## Name Kind

    类型名称对应的反射获取方法是 reflect.Type 中的 Name() 方法，返回表示类型名称的字符串；类型归属的种类（Kind）使用的是 reflect.Type 中的 Kind() 方法，返回 reflect.Kind 类型的常量。

### Kind

    Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。
    Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr
    type Kind uint
    const (
        Invalid Kind = iota  // 非法类型
        Bool                    //    布尔型
        Int                       //   有符号整型
        Int8                    // 有符号8位整型
        Int16                 // 有符号16位整型
        Int32                 // 有符号32位整型
        Int64                // 有符号64位整型
        Uint                 // 无符号整型
        Uint8                // 无符号8位整型
        Uint16               // 无符号16位整型
        Uint32               // 无符号32位整型
        Uint64               // 无符号64位整型
        Uintptr              // 指针
        Float32              // 单精度浮点数
        Float64              // 双精度浮点数
        Complex64            // 64位复数类型
        Complex128           // 128位复数类型
        Array                // 数组
        Chan                 // 通道
        Func                 // 函数
        Interface            // 接口
        Map                  // 映射
        Ptr                  // 指针
        Slice                // 切片
        String               // 字符串
        Struct               // 结构体
        UnsafePointer        // 底层指针
    )

### Name

    type Enum int
    type Student struct{
        Name string
        Age int
    }
    如果 对上述两种变量使用reflect.TypeOf().Kind(),reflect.TypeOf().Name(),前者获得的是 int和struct,后者获得的是Enum和Student

### Elem

    reflect.Elem() 方法获取这个指针指向的元素类型

### Struct

    如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息

    StructField 的结构如下

    type StructField struct {
        Name string          // 字段名
        PkgPath string       // 字段路径
        Type      Type       // 字段反射类型对象
        Tag       StructTag  // 字段的结构体标签
        Offset    uintptr    // 字段在结构体中的相对偏移
        Index     []int      // Type.FieldByIndex中的返回的索引值
        Anonymous bool       // 是否为匿名字段
    }

    Field(i int) StructField
        根据索引返回索引对应的结构体字段的信息，当值不是结构体或索引超界时发生宕机
    NumField() int
        返回结构体成员字段数量，当类型不是结构体或索引超界时发生宕机
    FieldByName(name string) (StructField, bool)
        根据给定字符串返回字符串对应的结构体字段的信息，没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机
    FieldByIndex(index []int) StructField
        多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生宕机
    FieldByNameFunc(match func(string) bool) (StructField,bool)
        根据匹配函数匹配需要的字段，当值不是结构体或索引超界时发生宕机
