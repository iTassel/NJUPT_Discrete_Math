package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"stack"
	"errors"
)

func main() {
	var choice,count,table_Num,i uint64
	fmt.Printf("Author:\t 风沐云烟\n")
	var input,var_Table string
New:var_Table,input,count,table_Num = get_input()
	fmt.Println("Input: ",input,"\t","Var_Table: ",var_Table,"Var_Num: ",count,"\t","Table_Num: ",table_Num)
	table := make([][]uint64,table_Num)
	for i = 0 ;i < table_Num; i++ {
		table[i] = make([]uint64,count + 1)
	}
	err := print_Truth_Table(input,var_Table,table,count,table_Num)
	if err != nil {
		fmt.Printf("公式错误,请重新输入\n")
		goto New
	}
	for true {
		choice = menu()
		switch choice {
			case 1:
				Disjunctive(input,var_Table,table,count,table_Num)
			case 2:
				Conjunction(input,var_Table,table,count,table_Num)
			case 3:
				goto New
			case 4:
				fmt.Printf("Bye~Bye~\n")
				os.Exit(0)
			default:
				fmt.Printf("错误输入,请重新选择\n")
		}
	}
}

func Disjunctive(input string,var_Table string,table [][]uint64,count uint64,table_Num uint64) {
	var i,j uint64
	var t bool = true
	for i = 0; i < table_Num ;i++ {
		if(table[i][count] != 0) {
			continue
		}
		if(t) {
			t = false
		} else {
			fmt.Printf("∨")
		}
		fmt.Printf("(")
		for j = 0; j < count ; j++ {
			if (table[i][j] != 0) {
				fmt.Printf("¬")
			}
			fmt.Printf(string(var_Table[j]))
			if j != count-1 {
				fmt.Printf("∧")
			}
		}
		fmt.Printf(")")
	}
	fmt.Printf("\n")
}

func Conjunction(input string,var_Table string,table [][]uint64,count uint64,table_Num uint64) {
	var i,j uint64
	var t bool = true
	for i = 0; i < table_Num ;i++ {
		if(table[i][count] == 0) {
			continue
		}
		if(t) {
			t = false
		} else {
			fmt.Printf("∧")
		}
		fmt.Printf("(")
		for j = 0; j < count ; j++ {
			if (table[i][j] == 0) {
				fmt.Printf("¬")
			}
			fmt.Printf(string(var_Table[j]))
			if j != count-1 {
				fmt.Printf("∨")
			}
		}
		fmt.Printf(")")
	}
	fmt.Printf("\n")
}

func print_Truth_Table (input string,var_Table string,table [][]uint64,count uint64,table_Num uint64) error {
	var i,j,num uint64
    var T bool = true
    var F bool = false
	var Map map[string] int
	Map = make(map[string] int)
	for i = 0; i< table_Num;i++ {
		num = i
		for j = 0;j < count ;j++ {
			table[i][j] = num%2
			num /= 2
		}
	}
	fmt.Printf("Truth Table\n")
	for i = 0; i < count; i++ {
		fmt.Printf(string(var_Table[i]))
		fmt.Printf("\t")
	}
	fmt.Println(input)
	input = Change_Op(input)
	for i = 0; i< count + 1;i++ {
		fmt.Printf("-\t")
	}
	fmt.Printf("\n")
	for i = 0; i< table_Num;i++ {
		num = i
		for j = 0;j < count;j++ {
			if table[i][j] == 0 {
				fmt.Printf("T")
				Map[string(var_Table[j])] = 0
			} else {
				fmt.Printf("F")
				Map[string(var_Table[j])] = 1
			}
			fmt.Printf("\t")
		}
		ret,err := calc(input,Map)
		if err == nil {
			if(ret == T) {
				table[i][count] = 0
				fmt.Printf("T")
			} else if(ret == F) {
				fmt.Printf("F")
				table[i][count] = 1
			}
		} else {
			return err
		}
		fmt.Printf("\n")
	}
	return nil
}
func calc(input string,Map map[string] int) (bool,error) {
	inffix_Table := input
	for i,v := range Map {
		inffix_Table = strings.Replace(inffix_Table, i, strconv.Itoa(v), -1 )
	}
	suffix_Table := inffixToSuffix(inffix_Table)
	ret,err := calcSuffixExpr(suffix_Table)
	if err == nil {
		return ret,nil
	} else {
		return false,err
	}
}
func get_input()(string,string,uint64,uint64) {
	var input string
	var count uint64
	var table_Num uint64
	var varName string

	fmt.Printf("Please Use \t¬ ∧ ∨ → ↔\n")
	fmt.Printf("For Example:  (P∧Q)∨R\n")
	fmt.Printf("输入公式\n")
	fmt.Scanf("%s",&input)
	varName,count = parseInput(input)
	table_Num = uint64(pow(2,count))
	return varName,input,count,table_Num
}
func parseInput(input string) (string,uint64) {
	tmp := input
	var count uint64
	var varName string = ""
	var table string =  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for _,v := range tmp {
		if strings.Index(table,string(v)) != -1 {
			table = strings.Replace(table, string(v), "", -1 )
			varName += string(v)
			count = count + 1
		}
	}
	return varName,count
}
func Change_Op (input string) string {
	var symbol map[string] string
	symbol = make(map[string] string)
	symbol["¬"] = "!"
	symbol["∧"] = "&"
	symbol["∨"] = "|"
	symbol["→"] = ">"
	symbol["↔"] = "="
	tmp := input
	for i,v := range symbol {
		tmp = strings.Replace(tmp, i, v, -1 )
	}
	return tmp
}
func menu() uint64 {
	var choice uint64
	fmt.Printf("1. 主析取范式\n")
	fmt.Printf("2. 主合取范式\n")
	fmt.Printf("3. 重新输入公式\n")
	fmt.Printf("4. 退出\n")
	fmt.Printf("Command: ")
	fmt.Scanf("%d",&choice)
	return choice
}
func pow(x float64, n uint64) float64 {
	if x == 0 {
		return 0
	}
	result := calcPow(x, n)
	if n < 0 {
		result = 1 / result
	}
	return result
}
func calcPow(x float64, n uint64) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	result := calcPow(x, n>>1)
	result *= result
	
	if n&1 == 1 {
		result *= x
	}
 
	return result
}


func inffixToSuffix(input string) (string) {
	S := stack.NewEmptyStack()
	var op string = ""
	for _,v := range input {
		switch(v) {
			case '1':
				op += string(v)
			case '0':
				op += string(v)
			case '(':
				S.Push(string(v))
			case ')':
				for !S.IsEmpty() {
					if (S.Top() == "(") {
						S.Pop()
						break
					}
					op += S.Top().(string)
					S.Pop()
				}
			case '!':
				S.Push("!")
			case '|':
				fallthrough
		    case '&':
				for !S.IsEmpty() && S.Top() != "(" && S.Top() != ">" && S.Top() != "=" {
		            op += S.Top().(string)
		            S.Pop()
		        }
		        S.Push(string(v))
		    case '>':
		    	fallthrough
		    case '=':
		    	for !S.IsEmpty() && S.Top() != "(" {
		            op += S.Top().(string)
		            S.Pop()
		        }
		        S.Push(string(v))
		}
	}
	for !S.IsEmpty() {
        op += S.Top().(string)
        S.Pop()
    }
    return op
}

func calcSuffixExpr(input string) (bool,error) {
    S := stack.NewEmptyStack()
    var T bool = true
    var F bool = false
    var x,y,res bool
    for _,v := range input {
        if (v == '0') {
            S.Push(T)
        } else if (v == '1') {
        	S.Push(F)
        } else if (v == '!' && S.Top() != nil) {
            x = S.Top().(bool)
            S.Pop()
            S.Push(!x)
        } else {
        	if S.Top() != nil {
	            x = S.Top().(bool)
		        S.Pop();
		        if S.Top() != nil {
				    y = S.Top().(bool)
				    S.Pop();
				} else {
					return false,errors.New("No item in stack's top is bool :(")
				}
            } else {
            	return false,errors.New("No item in stack's top is bool :(")
            }
            switch (v) {
            case '|':
                res = x || y
            case '&':
                res = x && y
            case '>':
                res = !x && y
            case '=':
                res = (x == y)
            }
            S.Push(res)
        }
    }
    if S.Top() != nil {
    	return S.Top().(bool),nil
    } else {
    	return false,errors.New("Error Input :(")
    }
}

