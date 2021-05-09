package Binding

type JsonBindingStruct struct {
	Name string `json:"user_name"`
	Age  int    `json:"user_age"`
}
type FormBindingStruct struct {
	Name string `form:"user_name"`
	Age  int    `form:"user_age"`
}
type JsonValidateStruct struct {
	Name string `json:"user_name"`
	Age  int    `json:"user_age"`
}
