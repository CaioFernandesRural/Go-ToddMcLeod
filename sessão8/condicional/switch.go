package main

import "fmt"

func main() {

	/*switch {
	case false:
		fmt.Println("não printa")
	case 2 == 4:
		fmt.Println("não printa")
	case 2 == 2:
		fmt.Println("printa")
	case 3 == 3:
		fmt.Println("printa")
	case 4 == 4:
		fmt.Println("printa")
	}*/
	switch {
	case false:
		fmt.Println("não printa")
	case 2 == 4:
		fmt.Println("não printa")
	case 2 == 2:
		fmt.Println("printa")
		fallthrough
	case 3 == 3:
		fmt.Println("printa")
		fallthrough
	case 4 == 4:
		fmt.Println("printa")
	default:
		fmt.Println("a")
	}
}
