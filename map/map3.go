package main
import("fmt")

func main() {
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "白骨精"
		monsters[1]["age"] = "1000"
	}

	newMonster := map[string]string{
		"name" : "火云邪神",
		"age" : "111",
	}

	monsters = append(monsters, newMonster)

	fmt.Println(monsters)

	var stored []map[string]map[string]string
	stored = make([]map[string]map[string]string, 2)

	stored[0] = make(map[string]map[string]string, 1)
	stored[0]["qq"] = make(map[string]string, 2)
	stored[0]["qq"]["word_id"] = "1111"
	stored[0]["qq"]["hints"] = "4123"
	stored[0]["qq"]["popular"] = "10"

	stored[1] = make(map[string]map[string]string, 1)
	stored[1]["douyin"] = make(map[string]string, 2)
	stored[1]["douyin"]["word_id"] = "2222"
	stored[1]["douyin"]["hints"] = "6123"
	stored[1]["douyin"]["popular"] = "110"
	fmt.Println(stored)
}
