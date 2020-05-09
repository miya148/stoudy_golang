package main

import "fmt"
import "strings"
/*
* 文字列Sと、その文字列から一文字だけ取り除いて作られた文字列Tが与えられる。
* Sから取り除かれた文字は何番目の文字でしょうか？可能性がいくるあるか求めよ
* 条件：1 <= T <= 1000
* T + 1 = S
* T,Sは小文字の文字列
* 可能性は必ず1以上
*/
func main(){
	var S string
	var T string

	fmt.Scan(&S, &T)

	result := comparison(S, T)

	fmt.Println(result)

}

func comparison(S string, T string) (int){
	var count int
	var a string

	s_slice := strings.Split(S,"")

	
	for j := 0; j <len(S); j ++ {
		for i,slice := range s_slice{
			if i != j{
				a += slice
			}
		}
		if a == T{
			count ++
		}
		a = ""
	}

	return count
}