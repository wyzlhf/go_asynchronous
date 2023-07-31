package main

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
