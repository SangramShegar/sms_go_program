
package main 
 import "fmt"

 type author struct { 
 name string 
 branch string 
 particles int 
} 
func (a *author) show(abranch string) { 
 (*a).branch = abranch 
} 
func main(){
 res := author{ 
 name: "ABC", 
 branch: "CSE", 
 } 
 fmt.Println("Author's name: ", res.name) 
 fmt.Println("Branch Name(Before): ", res.branch) 
 p := &res  
 p.show("ECE") 
 fmt.Println("Author's name: ", res.name) 
 fmt.Println("Branch Name(After): ", res.branch) 
}