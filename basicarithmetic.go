package main
import "fmt"

func arithmetic(a int, b int) (int, int, int){
	sum := a + b
	diff := a - b
	product := a * b
	 return sum, diff , product
	}
func main(){
	var num1,num2 int
	fmt.Println("enter first number:  ")
	fmt.Scanf("%d", &num1)
	fmt.Println("enter second  number:  ")
	fmt.Scanf("%d",&num2)
	sum,diff,product := arithmetic(num1,num2)
	fmt.Println("the sum is = ", sum)
	fmt.Println("the diffrence  is = ", diff)
	fmt.Println("the product  is = ", product)
	}
	
	
	
	
