package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyHandler struct{}

const (
	UNIT_G = int64(1073741824)
	UNIT_M = int64(1048576)
	UNIT_K = int64(1024)
)

func (p *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	list, err := ioutil.ReadDir(*dirname)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
		return
	}

	for _, info := range list { //遍历目录下的内容，获取文件详情，同os.Stat(filename)获取的信息
		unit := ""
		size := info.Size()
		if size >= UNIT_G { //1024*1024*1024
			unit = "G"
			size = size / UNIT_G
		} else if size >= UNIT_M {
			unit = "M"
			size = size / UNIT_M
		} else if size >= UNIT_K {
			unit = "K"
			size = size / UNIT_K
		} else {
			unit = "B"
		}

		fmt.Fprintf(w, "%s | %d%s\n",
			info.Name(), size, unit)
		/*
			if info.IsDir() == true {
				fmt.Println("是目录")
			}
		*/
	}
}

var dirname = flag.String("d", ".", "file server dir")
var hostport = flag.String("-s", ":5321", "file server dir")

func main() {
	flag.Parse()

	http.Handle("/files/", &MyHandler{})
	fmt.Println(http.ListenAndServe(*hostport, nil))
}
