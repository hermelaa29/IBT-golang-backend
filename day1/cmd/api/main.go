package main

import ("fmt"; "internal/delivery"; "internal/order")

type Dish struct {
	Name string
	price int
}
type Order struct {
	Dishes []Dish
	isPaid bool
}
Shiro := Dish{
	Name: "shiro",
	price: 100,
}
order1 := Order{
     Dishes: []Dish{shiro},
	 isPaid: false,	
}
func (o order) AddDish(d Dish) order{
	o.Dishes = append(o.Dishes, d)
	return o
}

func main(){
	var name1 string = "hermi"
	 name2  := "hermi"
	var age int = 22
	 age2  := 23

	fmt.Print("hello " , name1 , "ur age is ", age , "right?")
	fmt.Println("hello " , name2 , "ur age is ", age2 , "right?")
}

