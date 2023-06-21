package main
import "fmt"
import "math"
func largestPrimeFactor(n int) int{
    //Initialising largest prime number as 1
    largest:=1
    for n%2==0{
        largest=2
        n/=2
    }
    //check for divisibility by odd numbers starting from 3 up to the square root of n
    for i:=3;i<=int(math.Sqrt(float64(n)));i+=2{
        for n%i==0{
            largest=i
            n/=i
        }
    }
    //if n is prime number greater than 2,it is the largest prime factor
    if n>2{
        largest=n
    }
    return largest
}

func main() {
  var number int
  fmt.Println("Enter the number to find largest PF")
  fmt.Scanln(&number)
  largest:=largestPrimeFactor(number)
  fmt.Println("The largest prime factor:",largest)
 
}