package model

type Statement interface {
	New(a []string) Statement
}

type Stmt1 struct {
	Date      string
	Statement string
	Detail    string
	NumCheck  string
	Amount    string
	Tax       string
	Balance   string
	Branch    string
}

func (_ Stmt1) New(a []string) Statement {
	return &Stmt1{a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]}
}

type Stmt2 struct {
	Date      string
	Amount    string
	Statement string
	Detail    string
}

func (_ Stmt2) New(a []string) Statement {
	return &Stmt2{a[0], a[1], a[2], a[3]}
}

type Stmt3 struct {
	date        string
	price       string
	description string
}

func (_ Stmt3) New(a []string) Statement {
	return &Stmt3{a[0], a[1], a[2]}
}
