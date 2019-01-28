package main
import("fmt")

func main() {
	baseinfo := make(map[string]map[string]string)

	baseinfo["1"] = make(map[string]string)
	baseinfo["1"]["name"] = "luyuan"
	baseinfo["1"]["age"] = "18"
	baseinfo["1"]["sex"] = "boy"

	baseinfo["2"] = make(map[string]string)
	baseinfo["2"]["name"] = "xiaoming"
	baseinfo["2"]["age"] = "16"
	baseinfo["2"]["sex"] = "boy"

	for key, value := range baseinfo {
		for k2, v2 := range value {
			fmt.Printf("key=%v, k2=%v, v2=%v\n", key, k2, v2)
		}
	}

	fmt.Println("长度", len(baseinfo))

	fmt.Println(baseinfo)
}
