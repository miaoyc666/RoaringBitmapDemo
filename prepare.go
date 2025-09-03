package main

import (
	"fmt"
	"os"

	"github.com/RoaringBitmap/roaring/roaring64"
)

func saveBitmap(filename string, uids []uint64) {
	bm := roaring64.New()
	for _, id := range uids {
		bm.Add(id)
	}

	data, err := bm.MarshalBinary()
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Saved %s with %d UIDs\n", filename, len(uids))
}

func main() {
	// 假设有几个分类，每个分类对应一些 UID（这里直接用数字）
	categories := map[string][]uint64{
		"sports": {1, 2, 3, 100, 200},
		"music":  {2, 3, 4, 300},
		"movie":  {5, 6, 100, 300, 400},
	}

	for name, uids := range categories {
		saveBitmap(name+".roar", uids)
	}
}

