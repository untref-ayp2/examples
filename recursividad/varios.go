package recursividad

import "fmt"

func imprimirArray(a []int) {
	for _, valor := range a {
		fmt.Println(valor)
	}
}

func ImprimirRec(a []int) { // 1,2,3
	if len(a) == 0 {
		return
	}
	//1
	if len(a) == 1 {
		fmt.Print(a[0])
		return
	}
	// array tiene 2 o mas
	fmt.Println(a[0])
	ImprimirRec(a[1:]) //3

}
