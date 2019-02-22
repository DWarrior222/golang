package people

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type People struct {
	Name string
	Age  int
}

func (_this *People) Serialize() bool {
	// 序列化
	res, err := json.Marshal(_this)
	if err != nil {
		fmt.Println("marshal err is ", err)
		return false
	}
	// 存储到文件
	serPath := "../people.ser"
	err = ioutil.WriteFile(serPath, res, 0666)
	if err != nil {
		fmt.Println("write file err is ", err)
		return false
	}
	return true
}

func (_this *People) Unserialize() bool {
	// 获取内容
	serPath := "../people.ser"
	data, err := ioutil.ReadFile(serPath)
	if err != nil {
		fmt.Println("read file err is ", err)
		return false
	}
	// 反序列化
	err = json.Unmarshal(data, _this)
	if err != nil {
		fmt.Println("unmarshal err is ", err)
		return false
	}
	return true
}
