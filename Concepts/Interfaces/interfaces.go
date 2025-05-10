package main

func main() {

}

type UsersMethods interface {
}

type Values struct {
	um UsersMethods
}

func application(un UsersMethods) *Values {
	return &Values{}
}
