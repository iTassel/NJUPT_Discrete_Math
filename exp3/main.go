package main

import (
		"fmt"
		"stack")

func gcd(a uint64,b uint64) uint64 {
	var tmp uint64
	if a < b {
		tmp = a
		a = b
		b = tmp
	}
	for b != 0 {   
		tmp = a % b
		a = b
		b = tmp
    }
    return a
}
func lcm(a uint64,b uint64) uint64 {
	return a*b/gcd(a,b)
}

func main() {
	var num uint64
	fmt.Println("Please Input The Num")
	fmt.Scanf("%d",&num)
	factors := stack.NewEmptyStack()
	for i := num; i > 0; i-- {
		if (num % i == 0) {
			factors.Push(i);
		}
	}
	fmt.Printf("因子:")
	for i := 0; i < factors.Size() ; i++  {
		fmt.Printf("%d ",factors.GetItem(i))
	}
	fmt.Println()
	
	matrix := make([][]bool,factors.Size())
	for i := 0 ;i < factors.Size(); i++ {
		matrix[i] = make([]bool,factors.Size())
	}
	for i := 0; i < factors.Size() ; i++ {
		for j := 0; j < factors.Size() ; j++ {
			matrix[i][j] = (factors.GetItem(i).(uint64) % factors.GetItem(j).(uint64) == 0)
		}
	}
	
	fmt.Println("盖住集：")
	
    for i:= 0 ; i < factors.Size() ; i++ {
        for j := 0;j < factors.Size() ; j++ {
            if(i == j || !matrix[i][j]) {
            	continue
            }
            var flag bool = true
            for k := 0 ; k < factors.Size() ; k++ {
                if(k==i || k == j) {
                	continue
                }
                if matrix[i][k] && matrix[k][j] {
                    flag = false
                }
            }
            if flag {
                fmt.Println("(",factors.GetItem(i),",",factors.GetItem(j),")  ")
            }
        }
    }
    
    
    var all_flag bool = true
    for i := 0 ; i < factors.Size(); i++ {
        var flag bool = false
        for j := 0 ; j < factors.Size() ; j++ {
            flag = flag || ( gcd(factors.GetItem(i).(uint64),factors.GetItem(j).(uint64)) == 1 && lcm(factors.GetItem(i).(uint64),factors.GetItem(j).(uint64)) == num )
        }
        all_flag = all_flag && flag
    }
    if all_flag {
        fmt.Println("该格是有补格.")
    } else {
        fmt.Println( "该格不是有补格.")
    }
}
