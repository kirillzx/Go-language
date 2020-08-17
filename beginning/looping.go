package main

import (
	"fmt"
)

func main(){
	// for i:=0; i<10;i=i+2{
	// 	fmt.Printf("%v\n",i)
	// }
	//another way
	// j := 0
	// for ;j<5;j++{
	// 	fmt.Printf("%v\n",j)
	// }
Loop:
	for j:=1;j<10;j++{
		for k:=1;k<10;k=k+2{
			fmt.Printf("%v\n",k*j)
			if k*j > 4{
				break Loop
			}

		}
	}

	list := []int{1,14,27}
	for k, v := range list{
		fmt.Println(k, v)
	}
}
