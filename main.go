package main

import (
	"fmt"
	"os"

	"github.com/RoaringBitmap/roaring/roaring64"
)

func loadBitmap(filename string) *roaring64.Bitmap {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	bm := roaring64.New()
	if err := bm.UnmarshalBinary(data); err != nil {
		panic(err)
	}
	return bm
}

func main() {
	// 加载分类数据
	sports := loadBitmap("sports.roar")
	music := loadBitmap("music.roar")
	movie := loadBitmap("movie.roar")

	fmt.Println("sports size:", sports.GetCardinality())
	fmt.Println("music size:", music.GetCardinality())
	fmt.Println("movie size:", movie.GetCardinality())

	// --- 正确的交集写法（就地 And）---
	// 方式1：先克隆，再就地 And（推荐，避免改动原集合）
	inter := sports.Clone()
	inter.And(music)

	fmt.Println("sports ∩ music cardinality:", inter.GetCardinality())

	it := inter.Iterator()
	for it.HasNext() {
		fmt.Println("common UID:", it.Next())
	}

	// 方式2：如果只想要交集数量而不需要结果集合
	// card := sports.Clone() // 避免修改 sports
	// card.And(music)
	// fmt.Println("cardinality:", card.GetCardinality())

	// 方式3：用 Or 初始化再 And（不想 Clone 时）
	// inter2 := roaring64.New()
	// inter2.Or(sports) // inter2 = sports
	// inter2.And(music) // inter2 = sports ∩ music
}
