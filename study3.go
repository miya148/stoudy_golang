package main

import "fmt"
import "math"
/*
* N*G人で決議を行う
* 賛成する人がA人、反対する人がN*G-A人いるとわかっている
* N人ずつのGグループ作成し、賛成人数が多いグループを賛成グループとする
* グループ分けを自由にできる場合、賛成グループは最大いくつ、また最小でいくつできるのかを求める
* 
* 条件:N,Gは奇数
* 1 <= N*G <= 10^6
* 0 <= A <= N*G
* 入力は全て整数
*/

func main(){
	var N int
	var G int
	var A int
	var flag bool = false
	var max int 
	var min int

	for  flag == false{
		fmt.Scan(&N,&G,&A)
		if N * G > 0 && N * G <= 1000000 && A >= 0 && A <= N * G {
			flag = oddVerification(N, G)
		} 
		if !flag{
			fmt.Printf("入力をやり直してください")
		}
	}
	max, min = agreeTeam(N, G, A)
	fmt.Printf("%d, %d\n",max,min)
}

func oddVerification(N int, G int) (bool){
	if N % 2 == 0 || G % 2 == 0{
		return false
	}
	return true
}

func agreeTeam(N int, G int, A int) (int, int){
	var majority int 

	if N > 1{
		majority = int(math.Round(float64(N)/2.0))
	} else{
		majority = 1
	}
	max := A / majority
	min := (N * G - A) / majority

	return max, min
}