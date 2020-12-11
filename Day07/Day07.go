package Day07

type bagContains struct{
	color string
	number int
}

type bag struct{
	color string
	contains []bagContains
	containedBy []string
	canHoldGold bool
}


