// map 排序

// golang 中 map 的排序，是先将 key 进行排序，然后根据 key 值遍历输出即可
package main
import(
	"fmt"
	"sort"
)

func main() {
	var map1 map[int]int
	// map 的容量达到后，再想 map 增加元素，会自动扩容，并不会发生 panic，也就是说 map 能动 态的增长 键值对(key-value)
	map1 = make(map[int]int)

	// map 的 value 也经常使用 struct 类型，更适合管理复杂的数据
	map1[10] = 100
	map1[3] = 333
	map1[2] = 122
	map1[5] = 111

	fmt.Println(map1)

	var keys []int

	for k, _ := range map1 {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map[%v]=%v \n", k, map1[k])
	}
}
