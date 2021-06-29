package main

//User Is Main Table
type User struct {
	UId  int `gorm:"primary_key,auto_increment"`
	Name string
}

//CreditCard Is Dependent Table
// CreditCard Belongs To User
type CreditCard struct {
	CId      int `gorm:"primary_key,auto_increment"`
	CardName string
	UserUId  int
	User     User
}

func main() {

}
