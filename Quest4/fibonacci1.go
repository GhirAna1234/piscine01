package main 
import("fmt")
func Fibonacci(a int) int {
if a<0{
	return-1
}
if a==0{
	return 0
}
if a==1{
	return 1
}
return Fibonacci(a-1)+Fibonacci(a-2)
}
func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}