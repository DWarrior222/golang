package main
import(
	"fmt"
	"net"
	"../common/message"
	"encoding/binary"
	"encoding/json"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")

	// conn.Read 在conn没有关闭的时候才会阻塞
	// 如果客户端关闭了conn，则不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}

	// 根据buf[:4]转为一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	// 根据pkgLen读取内容
	n, err := conn.Read(buf[:pkgLen])

	if n != int(pkgLen) || err != nil {
		return
	}

	// 把pkgLen反序列化成 -> message.Message
	err = json.Unmarshal(buf[:pkgLen], err)
	if err != nil {
		fmt.Println("json.Unmarshal err is", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	// 发送data本身
	n, err = conn.Write(data)

	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Wirte fail", err)
		return
	}
	return
}
