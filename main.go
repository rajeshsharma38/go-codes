package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p *person) printName() string {
	return p.firstName + " " + p.lastName
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

	m := make(map[string]string)
	m1 := map[string]string{}

	fmt.Println(m)
	fmt.Println(m1)

	alexPointer := &alex

	fmt.Println(alexPointer.printName())
	fmt.Println(alex.printName())

	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	// nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for i := range nums {
	// 	if nums[i]%2 == 0 {
	// 		fmt.Printf("%d is even\n", nums[i])
	// 	} else {
	// 		fmt.Printf("%d is odd\n0", nums[i])
	// 	}
	// }

}
