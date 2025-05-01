package main

import (
	"errors"
	"fmt"
)


func main() {
	fmt.Println(club(17))
}
func club(age uint)(string, error){
	if(age<18){
		return "Not all0wed \n",errors.New("too young")
		}
		return "C0me in \n",nil
		}