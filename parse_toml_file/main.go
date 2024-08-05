package main

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"k8s.io/apimachinery/pkg/api/resource"
)

type Config struct {
	Timeout   time.Duration
	MemoryStr string    `toml:"memory_str"`
	MemoryInt int64     `toml:"memory_int"`
	StartTime time.Time `toml:"start_time"`
}

func main() {
	bs, err := os.ReadFile("demo.toml")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

	var c Config
	err = toml.Unmarshal(bs, &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)

	ms, err := resource.ParseQuantity(c.MemoryStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(ms.String())
	mi := resource.NewQuantity(c.MemoryInt, resource.BinarySI)
	fmt.Println(mi.String())
}
