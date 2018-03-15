package main

import "fmt"

//Function that converts infixed regular expressions to
//post fix regular expressions
func intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	pofix := []rune{}
	s := []rune{}

	//For loop to loop over input
	for _, r := range infix{
		
	}

	}

	return string(pofix)
}

func main() {
	//Answer: ab.c*.
	fmt.Println("Infix:	  ", "a.b.c*")
	fmt.Println("Postfix:	", intopost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix:	  ", "a.(b|d))*")
	fmt.Println("Postfix:	", intopost("a.(b|d))*"))

	//Answer: abd|.c*
	fmt.Println("Infix:	  ", "a.(b|d).c*")
	fmt.Println("Postfix:	", intopost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix:	  ", "a.(b.b)+.c")
	fmt.Println("Postfix:	", intopost("a.(b.b)+.c"))

}
