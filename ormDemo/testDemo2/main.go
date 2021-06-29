package main

//CreditCard Is Dependent Table
type CreditCard struct {
	CId      int `gorm:"primary_key,auto_increment"`
	CardName string
	UserUId  int
}

//User Is Main Table
//User Has A CreditCard
type User struct {
	UId        int `gorm:"primary_key,auto_increment"`
	Name       string
	CreditCard CreditCard
}

func main() {

}
