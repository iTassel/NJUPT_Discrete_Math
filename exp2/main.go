package main

import (
	"fmt"
	"os"
	"errors"
	"strings"
)

func main() {
	var choice uint64
New:num,table,err := get_input()
	if err != nil {
		fmt.Printf("变量代号越界,请重新输入\n")
		goto New
	}
	for true {
		choice = menu()
		switch choice {
			case 1:
				print_Ans(table,num)
			case 2:
				goto New
			case 3:
				fmt.Printf("Bye~Bye~\n")
				os.Exit(0)
			default:
				fmt.Printf("错误输入,请重新选择 :(\n")
		}
	}
}
func print_Ans(table [][] bool, num uint64) {
	if(checkReflexivity(num,table)) {
		fmt.Printf("满足自反性\t:)\n")
	}
	if(checkAntiReflexivity(num,table)) {
		fmt.Printf("满足反自反性\t:)\n")
	}
	if(checkSymmetry(num,table)) {
		fmt.Printf("满足对称性\t:)\n")
	}
	if(checkAntiSymmetry(num,table)) {
		fmt.Printf("满足反对称性\t:)\n")
	}
	if(checkTransitivity(num,table)) {
		fmt.Printf("满足传递性\t:)\n")
	}
}
func get_input() (uint64, [][]bool,error){
	var i,x,y,num,relation_num uint64
	fmt.Printf("输入元素数量:  ")
	fmt.Scanf("%d",&num)
	fmt.Printf("输入关系数量:  ")
	fmt.Scanf("%d",&relation_num)
	
	table := make([][]bool,num + 1)
	for i = 0 ;i < num; i++ {
		table[i] = make([]bool,num + 1)
	}
	//var Map map[string] uint64
	//Map = make(map[string] uint64)
	
	for i = 0; i < relation_num ; i++ {
		fmt.Println("输入第 ",i + 1," 组关系 (x,y)")
		fmt.Scanf("%d %d",&x,&y)
		if x < num && y < num {
			table[x][y] = true;
			table[y][x] = true;
		} else {
			return num,table,errors.New("Error Input :(")
		}
	}
	return num,table,nil
}
func parseInput(input string) (string,uint64) {
	tmp := input
	var count uint64
	var varName string = ""
	var table string =  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for _,v := range tmp {
		if strings.Index(table,string(v)) != -1 {
			table = strings.Replace(table, string(v), "", -1 )
			varName += string(v)
			count = count + 1
		}
	}
	return varName,count
}
func menu() uint64 {
	var choice uint64
	fmt.Printf("1. 输出判断结果\n")
	fmt.Printf("2. 重新输入公式\n")
	fmt.Printf("3. 退出\n")
	fmt.Printf("Command: ")
	fmt.Scanf("%d",&choice)
	return choice
}

func checkReflexivity(num uint64, table [][] bool) bool{
	var i uint64
    for i = 0;i < num; i++ {
        if !table[i][i] {
			return false
        }
    }
    return true
}
func checkAntiReflexivity(num uint64, table [][] bool) bool{
	var i uint64
    for i = 0;i < num; i++ {
        if table[i][i] {
			return false
        }
    }
    return true
}
func checkSymmetry(num uint64, table [][] bool) bool{
	var i,j uint64
    for i = 0;i < num; i++ {
        for j = 0;j < num; j++ {
            if table[i][j] != table[j][i] {
                return false
            }
        }
    }
    return true
}
func checkAntiSymmetry(num uint64, table [][] bool) bool{
	var i,j uint64
    for i = 0;i < num; i++ {
        for j = 0;j < num; j++ {
            if table[i][j] == table[j][i] {
                return false
            }
        }
    }
    return true
}
func checkTransitivity(num uint64, table [][] bool) bool{
	var i,j,k uint64
    for i = 0;i < num; i++ {
        for j = 0;j < num; j++ {
            for k = 0;k < num; k++ {
                if(!table[i][j] && (table[i][k] && table[k][j])){
                	return false
                }
            }
        }
    }
    return true
}


