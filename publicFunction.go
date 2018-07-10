package golibs

import (
	"net"
	"os"
)

//检查文件是否存在
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//获得本机一张网卡的地址
func GetHostIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok &&
			!ipnet.IP.IsLoopback() &&
			!ipnet.IP.IsMulticast() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}


//如果目录不存在，则创建一个目录
func MkDir(path string) error {
	_, err := os.Open(path)
	if err != nil {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
		}
	}
	return err
}
