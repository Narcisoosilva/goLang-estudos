package slices
import "fmt"

func main()  {
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)

	slice := []int{1,2,3,4,5}
	fmt.Println(slice)

	/*
	 a diferença de array  para slice  é que no slice vc pode aumentar seu tamanho,
	 mais nesta operação voce pederá o slice anterior e criará um novo.
	 */
	slice2 :=append(slice, 6)
	fmt.Println(slice2)

}
