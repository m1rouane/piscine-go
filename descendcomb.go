package main

import (
	"fmt"
)
func DescendComb() {
	for nbr1 := '0';nbr1 <= '9';nbr1++ {
		for nbr2 := '0';nbr2 <= '9';nbr2++ {
			for nbr3 := '0';nbr3 <= '9';nbr3++ {
				for nbr4 := '0';nbr4 <= '9';nbr4++ {
					if nbr1 <= nbr2 && nbr3 <= nbr4 {
						continue
					}
						fmt.Println(nbr1)
						fmt.Println(nbr2)
						fmt.Println(' ')
						fmt.Println(nbr3)
						fmt.Println(nbr4)
						if nbr1 != '8'|| nbr2 != '9'|| nbr3 != '9'|| nbr4 != '9' {
							fmt.Println(',')
							fmt.Println(' ')
					}
				}
			}
		}
	}
	fmt.Println('\n')
}

func main() {
	DescendComb()
}

