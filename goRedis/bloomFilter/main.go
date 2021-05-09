package main

import (
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"

	"github.com/demdxx/gocast"
)

type Encryptor struct {
}

func NewEncryptor() *Encryptor {
	return &Encryptor{}
}

func (e *Encryptor) Encrypt(origin string) int32 {
	hasher := murmur3.New32()
	_, _ = hasher.Write([]byte(origin))
	return int32(hasher.Sum32() % math.MaxInt32)
}

type LocalBloomService struct {
	//m 布隆过滤器的bitmap长度
	//k 映射函数的个数,也表示一个输入需要经过k次映射,占用k个bit位
	//n 表示bitmap已经容纳了n个输入
	m, k, n   int32
	Bitmap    []int
	encryptor *Encryptor
}

func NewLocalBloomService(m, k int32, encryptor *Encryptor) *LocalBloomService {
	return &LocalBloomService{
		m:         m,
		k:         k,
		Bitmap:    make([]int, m/32+1),
		encryptor: encryptor,
	}
}
func (l *LocalBloomService) Exist(val string) bool {
	for _, offset := range l.getKEncrypted(val) {
		index := offset / 32     // 等价于 / 32
		bitOffset := offset % 32 // 等价于 % 32

		if l.Bitmap[index]&(1<<bitOffset) == 0 {
			return false
		}
	}

	return true
}

func (l *LocalBloomService) Set(val string) {
	l.n++
	res := l.getKEncrypted(val)
	for _, offset := range res {
		index := offset / 32     // 等价于 / 32
		bitOffset := offset % 32 // 等价于 % 32

		l.Bitmap[index] |= (1 << bitOffset)
	}
}
func (l *LocalBloomService) getKEncrypted(val string) []int32 {
	encrypteds := make([]int32, 0, l.k)
	origin := val
	for i := 0; int32(i) < l.k; i++ {
		encrypted := l.encryptor.Encrypt(origin)
		encrypteds = append(encrypteds, encrypted%l.m)
		if int32(i) == l.k-1 {
			break
		}
		origin = gocast.ToString(encrypted)
	}
	return encrypteds
}

func main() {
	e := NewEncryptor()
	bf := NewLocalBloomService(129, 5, e)
	bf.Set("hello,world")
	fmt.Println(bf.Bitmap)

	bf.Set("hello,worlds")
	fmt.Println(bf.Bitmap)

	fmt.Println(bf.Exist("hello,wrolds"))
	fmt.Println(bf.Exist("hello,worlds"))

}
