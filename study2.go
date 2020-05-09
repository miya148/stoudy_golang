package main
import "fmt"
import "strings"

func main(){
    //タイピングの秒計算
    //条件：{"a","s","d","f","g","h","j","k","l"}の入力のみ
    //1文字以上100文字以下
    
    required_values := []string{"a","s","d","f","g","h","j","k","l"}
    var flag int  =  1
    var input string
    
    for flag == 1{
        fmt.Scan(&input)
        
        if len(input) >0 && len(input) <= 100{
            flag = verification(input, required_values)
            if flag == 0{
                break
            }
        } 
        fmt.Println("入力をやり直してください")
	}
	fmt.Println("入力チェック完了")
	result := calculation(input, required_values)
	fmt.Printf("妖精がボタンを押した時間と、移動した時間の合計は%d秒\n",result)
}

func verification(input string, required_values []string)(int){
	
	slice := strings.Split(input, "")
	count := 0
	for _,val := range slice{
		for _,required := range required_values{
			if val == required{
				count ++
				fmt.Printf("カウントは%d\n",count)
			}
		}
	}
	fmt.Printf("文字数は%d\n",len(slice))
	if count == len(slice){
		return 0
	}else {
		return 1
	}
}

func calculation(input string, required_values []string) (int){
	slice := strings.Split(input, "")
	count := 0
	var before int = 0
	for _,val := range slice{
		for i,required_value := range required_values{
			if val == required_value{
				move := i - before
				if move < 0{
					move = move / -1
				}
				count += 1 + move
				before = i
			}
		}
	}
	return count
}

