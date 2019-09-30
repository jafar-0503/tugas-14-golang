package main

import "fmt"
import "os/exec"

func main(){
  var hasil,_ = exec.Command("IPCONFIG").Output()
  fmt.Println(string(hasil))

  var hasil1,_ = exec.Command("attrib").Output()
  fmt.Println(string(hasil1))

  var hasil2,_ = exec.Command("assoc").Output()
  fmt.Println(string(hasil2))
}
