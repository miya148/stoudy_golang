// あなたはストップウォッチでx時間y分z秒を計りました。この時間が何秒か計算せよ
package main
import "fmt"
func main(){
    var x int
    var y int
    var z int
    
    fmt.Scan(&x, &y, &z)
    fmt.Printf("%d時間%d分%d秒\n",x,y,z)
    
    minutes := x * 60 + y
    L := minutes * 60 + z
    
    fmt.Printf("%d秒",L)
    
}
