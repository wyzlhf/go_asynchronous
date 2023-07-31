package main

import "fmt"

//func Buy(n int) <-chan string{
//	out:=make(chan string)
//	go func() {
//		defer close(out)
//		for i:=1;i<=n;i++{
//			out<-fmt.Sprint("配件",i)
//		}
//	}()
//	return out
//}
//func Build(in <-chan string) <-chan string{
//	out:=make(chan string)
//	go func() {
//		defer close(out)
//		for c:=range in{
//			out<-"组装("+c+")"
//		}
//	}()
//	return out
//}
//func Pack(in <-chan string) <-chan string{
//	out:=make(chan string)
//	go func() {
//		defer close(out)
//		for c:=range in{
//			out<-"打包("+c+")"
//		}
//	}()
//	return out
//}
//func main() {
//	accessories:=Buy(6)
//	computers:=Build(accessories)
//	packs:=Pack(computers)
//	for p:=range packs{
//		fmt.Println(p)
//	}
//}

func Generator(max int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
func Square(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}
func Sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var Sum int
		for v := range in {
			Sum += v
		}
		out <- Sum
		close(out)
	}()
	return out
}
func main() {
	arr := Generator(555555555)
	squ := Square(arr)
	sum := <-Sum(squ)
	fmt.Println(sum)
}
