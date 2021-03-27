package main

import (
	"fmt"
	"github.com/Albert-Zuo/Data-Structures-And-Algorithm/algorithm/balance"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	balanceName = loadConfig()
)

func loadConfig() string {
	return "hash"
}

func main() {

	var is []*balance.Instance
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port, _ := strconv.Atoi(fmt.Sprintf("880%d", i))
		one := balance.NewInstance(host, port)
		is = append(is, one)
	}
	// 参数获取
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(balanceName, is)
		if err != nil {
			fmt.Println("do balance err")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}


