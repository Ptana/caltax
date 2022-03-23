package main

import "fmt"

func caltax(income float32, name string) {
	var t float32
	var rate = [][8]float32{
		{5000000, 2000000, 1000000, 750000, 500000, 300000, 150000, 0},
		{0.35, 0.30, 0.25, 0.20, 0.15, 0.10, 0.05, 0.00},
	}

	for i := 0; i < 8; i++ {
		if income > rate[0][i] {
			t += (income - rate[0][i]) * rate[1][i]
			// fmt.Printf("%5.0f %2.0f %5.0f %5.0f\n ",
			// 	rate[0][i],
			// 	rate[1][i]*100,
			// 	income-rate[0][i],
			// 	(income-rate[0][i])*rate[1][i])
			income = rate[0][i]
		}

	}

	fmt.Println(name, "ภาษีที่ต้องจ่ายคือ ,%.2f", t)

}

func main() {

	var income float32
	var insur float32 = 45000
	var ltf float32 = 97000
	var max_ltf float32 = 200000
	var max_insur float32 = 100000
	var name string

	fmt.Printf("enter your name and Income -> ")
	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &income)

	if ltf < max_ltf {
		income = income - ltf
	} else {
		income = income - max_ltf
	}
	if insur < max_insur {
		income = income - insur
	} else {
		income = income - max_insur
	}
	fmt.Println("net Income before Calculate Tax of ", name, "is ", income)
	caltax(income, name)
	// for i := 0; i < 8; i++ {
	// 	if income > rate[0][i] {
	// 		t += (income - rate[0][i]) * rate[1][i]
	// 		fmt.Printf("%5.0f %2.0f %5.0f %5.0f\n ",
	// 			rate[0][i],
	// 			rate[1][i]*100,
	// 			income-rate[0][i],
	// 			(income-rate[0][i])*rate[1][i])
	// 		income = rate[0][i]
	// 	}

	// }

	// fmt.Println("pay tax ", t)

}

// var rate [8][8]float64 {
// 		{5000000, 2000000, 1000000, 750000, 500000, 300000, 150000, 0} ,
// 		{0.35, 0.30, 0.25, 0.20, 0.15, 0.10, 0.05, 0.00}
// 	}

// 	fmt.Println("2d array:  ", rate)
// 	// for i := 0; i < 100; i++ {
// 	// 	if i == 5 {
// 	// 		continue

// 	// 	} else if i == 10 {
// 	// 		break
// 	// 	}
// 	// 	println(i)
// 	// }

// }
