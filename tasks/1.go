package tasks

import (
	"fmt"
	"math"
)

/*Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
 */

// структура человек с полями имя, возраст, рост, масса тела
type Human struct {
	name   string
	age    int
	height float64
	weight float64
}

// функция, которая вычисляет индекс массы тела человека
func (man *Human) Body_Mass_Index() float64 {
	result := man.weight / (math.Exp(2 * math.Log(man.height)))
	fmt.Println("Body Mass Index", result, "kg/m^2")
	return result
}

// структура, которая "Наследуется" от Human
type Action struct {
	basis Human
}

// функция анализа результата на основе индекса массы тела
func (act *Action) Output_result() string {
	var res = ""
	IMT := act.basis.Body_Mass_Index()
	if IMT < 0 {
		res = "Error"
	} else {
		if IMT < 16 {
			res = "Severe underweight"
		} else if IMT < 18.5 {
			res = "Insufficient (deficit) body weight"
		} else if IMT < 25 {
			res = "Normal weight"
		} else if IMT < 30 {
			res = "Excess body weight (pre-obesity)"
		} else if IMT < 35 {
			res = "Obesity of the first degree"
		} else if IMT < 40 {
			res = "Obesity of the second degree"
		} else {
			res = "Obesity of the third degree"
		}
	}
	return res
}

func L1_1() {
	var masha = Human{name: "Masha", age: 22, height: 1.60, weight: 50}
	fmt.Println(masha)
	//masha.Body_Mass_Index()
	var action = Action{masha}
	fmt.Println("Result", action.Output_result(), "\n\n")

	//var kate = Human{name: "Kate", age: 24, height: 1.70, weight: 150} // Obesity of the third degree
	//var kate = Human{name: "Kate", age: 24, height: 1.70, weight: 100} // Obesity of the first degree
	//var kate = Human{name: "Kate", age: 24, height: 1.70, weight: 110} //Obesity of the second degree
	var kate = Human{name: "Kate", age: 24, height: 1.70, weight: 45} //Severe underweight
	fmt.Println(kate)
	var action_kate = Action{kate}
	fmt.Println("Result: ", action_kate.Output_result())

}
