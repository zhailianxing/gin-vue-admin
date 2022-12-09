package myPkg

import "fmt"

type MyApiService struct {
}

func (m *MyApiService) CreateApiService() {
	fmt.Println("this is service")
}
