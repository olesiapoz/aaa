package aaa

import (
	"fmt"

	"os"
)

func init() {
	fmt.Println("aaaInit")
	os.Setenv("USER_DATABASE", "mongodb")
	os.Setenv("MONGO_HOST", "192.168.0.3:27017")

}
func Load() {
	return 
}
