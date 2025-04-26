package main
import "fmt"

func main(){
	var b int
	b = 5
	def(b, 10)
}
func def(x, n int){
	if x > 0{
		fmt.Print(n-x, " DE ")
		def(x/2, n)
		fmt.Print(n+x, " EF ")
	}
}