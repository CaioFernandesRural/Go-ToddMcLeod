package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "batido, não mexido"}

	fmt.Println(jb)

	mp := []string{"miss", "moneypenny", "coe, menor"}

	fmt.Println(mp)

	xp := [][]string{jb, mp}

	fmt.Println(xp)

	for i, x := range xp {
		fmt.Println(i)
		for j, v := range x {
			fmt.Println(j, v)
		}
	}
}
