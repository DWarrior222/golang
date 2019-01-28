// 练习
package main
import(
	"fmt"
)

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name
	}
}

func main () {
	var user map[string]map[string]string

	user = make(map[string]map[string]string, 10)

	// user := make(map[string]map[string]string, 10)

	user["luyuan"] = make(map[string]string, 2)
	user["luyuan"]["nickname"] = "luyuan的昵称"
	user["luyuan"]["pwd"] = "2345"

	modifyUser(user, "luyuan")
	modifyUser(user, "xiaoming")

	fmt.Println(user)
}
