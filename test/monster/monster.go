package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 给 Monster 绑定方法 Store, 可以将一个 Monster 变量(对象),序列化后保存到文件中
func (_this *Monster) Store() bool {
	//先序列化
	data, err := json.Marshal(_this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}

	//保存到文件
	filePath := "../monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true
}

func (_this *Monster) ReStore() bool {
	//1. 先从文件中，读取序列化的字符串
	filePath := "../monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}

	//2.使用读取到 data []byte ,对反序列化
	err = json.Unmarshal(data, _this)
	if err != nil {
		fmt.Println("UnMarshal err =", err)
		return false
	}
	return true
}
