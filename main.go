package main

import (
	"fmt"
	"go-ADT/model"
)

// NOTE: https://eli.thegreenplace.net/2018/go-and-algebraic-data-types/
func main() {
	data := mapping()
	checkingADT(data)
}

func mapping() []model.Statement {
	// mock1
	// stmt1 := [][]string{
	// 	{"date", "statement", "detail", "numCheck", "amount", "tax", "balance", "branch"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "ATM annual Fee", "", "-0.99", "0.000000", "0.00", "0"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "ATM annual Fee", "", "-0.99", "0.000000", "0.00", "0"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "ATM annual Fee", "", "-0.99", "0.000000", "0.00", "0"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "ATM annual Fee", "", "-0.99", "0.000000", "0.00", "0"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "ATM annual Fee", "", "-0.99", "0.000000", "0.00", "0"},
	// }

	// mock2
	stmt2 := [][]string{
		{"date", "statement", "amount", "detail"},
		{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99", "ATM annual Fee"},
		{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99", "ATM annual Fee"},
		{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99", "ATM annual Fee"},
		{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99", "ATM annual Fee"},
		{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99", "ATM annual Fee"},
	}

	// mock3
	// stmt3 := [][]string{
	// 	{"date", "description", "price"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99"},
	// 	{"21-09-2019 07:22:07", "ค่าธรรมเนียมบัตร", "-0.99"},
	// }

	current := stmt2
	switch {
	case len(current[0]) == 8:
		return mapper(current, model.Stmt1{})
	case len(current[0]) == 4:
		return mapper(current, model.Stmt2{})
	default:
		return mapper(current, model.Stmt3{})
	}

}

func mapper(data [][]string, instance model.Statement) []model.Statement {
	result := make([]model.Statement, len(data)-1)
	for index, lane := range data {
		if index > 0 {
			result = append(result, instance.New(lane))
		}
	}
	return result
}

func checkingADT(data []model.Statement) {
	for _, v := range data {
		switch vt := v.(type) {
		case *model.Stmt1:
			fmt.Printf("%v / %T\n", vt, vt)
		case *model.Stmt2:
			fmt.Printf("%v / %T\n", vt, vt)
		case *model.Stmt3:
			fmt.Printf("%v / %T\n", vt, vt)
		}
	}
}
