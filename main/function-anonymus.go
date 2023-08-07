package main

import "fmt"

type Member func(string) bool

func registerMember(name string, member Member) {
	if member(name) {
		fmt.Println("Hallo",name)
	}else {
		fmt.Println("Register Member")
	}
}

// func member(name string) bool {
// 	return name == "member"
// }

// func nonMember(name string) bool {
// 	return name == "non member"
// }


func main () {
	memberList := func(name string) bool {
		return name == "member"
	}

	registerMember("member", memberList)
	registerMember("doni", memberList)
}