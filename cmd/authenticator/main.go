package main

import (
	"authenticator/internal/totp"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Secret key arg missing!")
		return
	}
	totp.Run(os.Args[1], time.Now().Unix())
}
