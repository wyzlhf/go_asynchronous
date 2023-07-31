package main

import (
	"fmt"
	"time"
)

func putInTea() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "茶叶已经放好了..."
	}()
	return vegetables
}
func boilingWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(10 * time.Second)
		water <- "水已经烧好了..."
	}()
	return water
}
func main() {
	teaCh := putInTea()
	waterCh := boilingWater()
	fmt.Println("已经安排好放茶叶和水，休息两秒吧。……")
	time.Sleep(3 * time.Second)

	tea := <-teaCh
	water := <-waterCh
	fmt.Println("准备好了，可以沏茶了：", tea, water)
}
