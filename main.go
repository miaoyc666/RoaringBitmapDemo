package main

import (
	"fmt"
	"os"

	"github.com/RoaringBitmap/roaring/roaring64"
)

// XBitmap 封装了 roaring64.Bitmap 并添加自定义方法
type XBitmap struct {
	*roaring64.Bitmap
}

// NewXBitmap 创建新的XBitmap
func NewXBitmap() *XBitmap {
	return &XBitmap{
		Bitmap: roaring64.New(),
	}
}

// NewXBitmapFromFile 从文件加载创建XBitmap
func NewXBitmapFromFile(filename string) *XBitmap {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	bm := roaring64.New()
	if err := bm.UnmarshalBinary(data); err != nil {
		panic(err)
	}
	return &XBitmap{
		Bitmap: bm,
	}
}

// List 遍历并打印所有数据（自定义方法）
func (xb *XBitmap) List(maxDisplay int) {
	fmt.Printf("\n=== 遍历数据 ===\n")

	iterator := xb.Iterator()
	count := 0
	for iterator.HasNext() {
		uid := iterator.Next()
		fmt.Printf("UID[%d]: %d\n", count, uid)
		count++

		// 如果数据太多，可以限制输出数量
		if maxDisplay > 0 && count >= maxDisplay {
			remaining := xb.GetCardinality() - uint64(count)
			if remaining > 0 {
				fmt.Printf("... (还有 %d 个数据，省略显示)\n", remaining)
			}
			break
		}
	}
	fmt.Printf("总共遍历了 %d 个数据\n", count)
}


func main() {
	// 使用XBitmap加载分类数据
	sports := NewXBitmapFromFile("sports.roar")
	music := NewXBitmapFromFile("music.roar")
	movie := NewXBitmapFromFile("movie.roar")

	// 打印基本信息
	fmt.Printf("sports size: %d\n", sports.GetCardinality())
	fmt.Printf("music size: %d\n", music.GetCardinality())
	fmt.Printf("movie size: %d\n", movie.GetCardinality())

	// 使用自定义的List方法遍历sports数据
	sports.List(10) // 显示前10个数据

	// --- 演示原生roaring64.Bitmap方法的复用 ---
	fmt.Println("\n=== 演示原生方法复用 ===")
	fmt.Printf("sports包含UID 1000: %v\n", sports.Contains(1000))
	fmt.Printf("music包含UID 2000: %v\n", music.Contains(2000))

	// 交集操作（使用原生方法）
	intersection := &XBitmap{Bitmap: sports.Clone()}
	intersection.And(music.Bitmap)
	fmt.Printf("sports ∩ music size: %d\n", intersection.GetCardinality())
	intersection.List(5)

	// 添加新数据
	testBitmap := NewXBitmap()
	testBitmap.Add(100)
	testBitmap.Add(200)
	testBitmap.Add(300)
	fmt.Printf("testBitmap size: %d\n", testBitmap.GetCardinality())
	testBitmap.List(0) // 显示所有数据

	// 演示更多原生方法
	fmt.Printf("testBitmap是否为空: %v\n", testBitmap.IsEmpty())
	fmt.Printf("testBitmap最小值: %d\n", testBitmap.Minimum())
	fmt.Printf("testBitmap最大值: %d\n", testBitmap.Maximum())
}
