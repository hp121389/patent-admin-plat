package dto

import (
	"fmt"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/my_config"
	"net"
	"os"
	"testing"
)

func TestGenAndDelFiles(t *testing.T) {
	my_config.LoadPatentConfig()

	ureq := PackageUpdateReq{
		Files: []string{"/a/b/c/123-123.1.txt", "/a/b/c/123-123.2.txt", "/a/b/c/123-123.3.txt"},
	}

	pm := models.Package{
		Files: "",
	}

	ureq.GenerateAndAddFiles(&pm)
	fmt.Println(pm.Files) //添加成功

	ureq.Files = []string{"/a/b/c/123-123.1.txt", "/a/b/c/123-123.3.txt"}
	ureq.GenerateAndDeleteFiles(&pm)
	fmt.Println(pm.Files)
}

func TestIPV4(t *testing.T) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
