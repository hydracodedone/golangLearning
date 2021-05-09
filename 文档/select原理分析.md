# select原理分析

## case的数据结构
```
type scase struct {
	c           *hchan          //chan
	elem        unsafe.Pointer  //data element
	kind        uint16
	pc          uintptr  //race pc (for race detector / msan)
	releasetime int64
}
```
scase.c为当前case语句所操作的channel指针，这也说明了一个case语句只能操作一个channel。
scase.kind表示该case的类型，分为读channel、写channel和default，三种类型分别由常量定义：
```
const (
	caseNil = iota
	caseRecv
	caseSend
	caseDefault
)

```
caseRecv：case语句中尝试读取scase.c中的数据；
caseSend：case语句中尝试向scase.c中写入数据；
caseDefault： default语句
scase.elem表示缓冲区地址，跟据scase.kind不同，有不同的用途：

scase.kind == caseRecv ： scase.elem表示读出channel的数据存放地址；
scase.kind == caseSend ： scase.elem表示将要写入channel的数据存放地址；

## select选择case的函数
```
func selectgo(cas0 *scase, order0 *uint16, ncases int) (int, bool)
```
cas0为scase数组的首地址，selectgo()就是从这些scase中找出一个返回。
order0为一个两倍cas0数组长度的buffer，保存scase随机序列pollorder和scase中channel地址序列lockorder
pollorder：每次selectgo执行都会把scase序列打乱，以达到随机检测case的目的。
lockorder：所有case语句中channel序列，以达到去重防止对channel加锁时重复加锁的目的。
ncases表示scase数组的长度
函数返回值：

int： 选中case的编号，这个case编号跟代码一致
bool: 是否成功从channle中读取了数据，如果选中的case是从channel中读数据，则该返回值表示是否读取成功。
### 选择逻辑
    1. 锁定scase语句中所有的channel
    2. 按照随机顺序检测scase中的channel是否ready
       2.1 如果case可读，则读取channel中数据，解锁所有的channel，然后返回(case index, true)
       2.2 如果case可写，则将数据写入channel，解锁所有的channel，然后返回(case index, false)
       2.3 所有case都未ready，则解锁所有的channel，然后返回（default index, false）
    3. 所有case都未ready，且没有default语句
       3.1 将当前协程加入到所有channel的等待队列
       3.2 当将协程转入阻塞，等待被唤醒
    4. 唤醒后返回channel对应的case index
       4.1 如果是读操作，解锁所有的channel，然后返回(case index, true)
       4.2 如果是写操作，解锁所有的channel，然后返回(case index, false)