package main
import "fmt"

func main() {
  var prev,current,sum,n int
  fmt.Println("Enter the number of terms:")
  fmt.Scan(&n)
  prev=0
  current=1
  sum=0
  
  for prev+current<n{
      next:=prev+current
      if next%2==0{
          sum+=next
          
      }
      prev=current
      current=next
  }
  fmt.Println("The sum of even fibonnocci is:",sum)
 
}