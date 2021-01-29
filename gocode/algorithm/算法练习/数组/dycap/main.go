package main

import (
	"bytes"
	"fmt"
)

/*
Dynamic capacity
实现一个支持动态扩容的数组
数组设计之初是在形式上依赖内存分配而成的，所以必须在使用前预先请求空间。这使得数组有以下特性：

1.请求空间以后大小固定，不能再改变（数据溢出问题）；
2.在内存中有空间连续性的表现，中间不会存在其他程序需要调用的数据，为此数组的专用内存空间；
3.在旧式编程语言中（如有中阶语言之称的C），程序不会对数组的操作做下界判断，也就有潜在的越界操作的风险（比如会把数据写在运行中程序需要调用的核心部分的内存上）。

传统数组的局限性导致了动态数组的诞生。然而动态数组也不是使用动态的内存，依旧是一块连续的内存。那它是如何实现数组大小不固定的呢？原因是当超过数组容量时，程序将自动执行扩容操作：

1.重新开辟一块大小为当前数组容量两倍的内存
2.把原数组的数据拷贝到此内存空间
3.释放原数组的内存

扩/缩容规则
数组 resize 指当数组元素超过数组容量，或者元素小于数组容量时，需要完成的扩容和缩容规则：

1.超过数组容量，按照当前容量的 2 倍扩容。
2.数组元素个数为当前容量 1/4 时，缩容为当前容量的一半。

为什么缩容不是 1/2？

如果在 1/2 时缩容，会导致在扩容的临界点添加、删除一个元素都是 O(n) 复杂度的情况（临界点添加一个元素，导致扩容为 2 倍，此时删除刚添加的元素，又会缩容为 1/2）。
*/

func main() {
	var arr ArrayInterface = GetArray(10)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)
	arr.Add(1, 100)
	fmt.Println(arr)
	arr.Remove(1)
	fmt.Println(arr)
}

type Array struct {
	Data []interface{}
	Size int
}

type ArrayInterface interface {
	// 添加
	Add(int, interface{}) // 插入元素
	AddFirst(interface{})
	AddLast(interface{})
	// 删除
	Remove(int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	// 查找
	Find(interface{}) int      // 查找元素返回第一个索引
	FindAll(interface{}) []int // 查找元素返回所有索引
	Contains(interface{}) bool // 查找是否有元素
	Get(int) interface{}
	// 修改
	Set(int, interface{})
	// 基本方法
	GetCapacity() int // 获取数组容量
	GetSize() int     // 获取元素个数
	IsEmpty() bool    // 查看数组是否为空
}

// GetArray 获取自定义数组,参数为数组初始容量
func GetArray(capacity int) *Array {
	arr := &Array{}
	arr.Data = make([]interface{}, capacity)
	arr.Size = 0
	return arr
}

// GetCapacity 获取数组容量
func (a *Array) GetCapacity() int {
	return len(a.Data)
}

// GetSize 获取元素个数
func (a *Array) GetSize() int {
	return a.Size
}

// IsEmpty 查看数组是否为空
func (a *Array) IsEmpty() bool {
	return a.Size == 0
}

// newCapacity 新的数组容量
// 逻辑：声明一个新的数组,然后将原数组的值copy到新数组
func (a *Array) resize(newCapacity int) {
	newArr := make([]interface{}, newCapacity)
	for i := 0; i < a.Size; i++ {
		newArr[i] = a.Data[i]
	}
	a.Data = newArr
}

// Find 获得元素的首个索引,不存在则返回-1
func (a *Array) Find(element interface{}) int {
	for i := 0; i < a.Size; i++ {
		if a.Data[i] == element {
			return i
		}
	}
	return -1
}

// FindAll 获得元素的所有索引，返回索引组成的切片
func (a *Array) FindAll(element interface{}) (indexes []int) {
	for i := 0; i < a.Size; i++ {
		if a.Data[i] == element {
			indexes = append(indexes, i)
		}
	}
	return
}

// 查看数组是否包含元素,返回bool
func (a *Array) Contains(element interface{}) bool {
	return a.Find(element) != -1
}

// Get 获取索引对应元素,需判断索引是否合法
func (a *Array) Get(index int) interface{} {
	if index < 0 || index > a.Size-1 {
		panic("Get failed, index is illegal")
	}
	return a.Data[index]
}

// Set 修改索引对应元素值
func (a *Array) Set(index int, element interface{}) {
	if index < 0 || index > a.Size-1 {
		panic("set failed, index is illegal")
	}
	a.Data[index] = element
}

// Add 指定索引添加元素,需要考虑扩容问题
func (a *Array) Add(index int, element interface{}) {
	if index < 0 || index > a.GetCapacity() {
		panic("add failed, index is illegal")
	}
	// 数组已满则扩容
	if a.Size == a.GetCapacity() {
		a.resize(2 * a.Size)
	}
	// 将插入索引后的元素后移,腾出插入位置
	for i := a.Size - 1; i >= index; i-- {
		a.Data[i+1] = a.Data[i]
	}
	a.Data[index] = element
	a.Size++
}

// AddFirst 数组头部插入元素
func (a *Array) AddFirst(element interface{}) {
	a.Add(0, element)
}

// AddLast 数组尾部插入元素
func (a *Array) AddLast(element interface{}) {
	a.Add(a.Size, element)
}

// Remove 删除指定索引元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.Size {
		panic("remove failed, index is illegal")
	}
	removeEle := a.Data[index]
	// 将删除索引后的元素前移,覆盖删除位置
	for i := index; i < a.Size-1; i++ {
		a.Data[i] = a.Data[i+1]
	}
	a.Size--
	a.Data[a.Size] = nil
	// 考虑边界情况,不能resize为0
	if a.Size == len(a.Data)/4 && len(a.Data)/2 != 0 {
		a.resize(len(a.Data) / 2)
	}
	return removeEle
}

// RemoveFirst 删除头部元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// RemoveLast 删除尾部元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.Size - 1)
}

// String 重写数组打印时的展示形式
func (a *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capaticy= %d\n", a.Size, a.GetCapacity()))
	buffer.WriteString("[")
	for i := 0; i < a.Size; i++ {
		buffer.WriteString(fmt.Sprintf("%v", a.Data[i]))
		if i != a.Size-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
