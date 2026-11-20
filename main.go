package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	/** Flags **/
	AdminUser := flag.String("admin-user", "ubuntu", "Specify the user's name for the admin account")
	
	flag.Parse()

	state, err := GetUserState(*AdminUser)

	if err != nil {
		fmt.Printf("User `%s` could not be checked. \n", *AdminUser)
		fmt.Printf("An error has ocurred: %v \n", err)
		os.Exit(1)
	}

	if !state.UserExists {
		fmt.Printf("User `%s` does not exist. (This is good, we can create it) \n", *AdminUser)
		os.Exit(1)
	}

	if !state.IsUserSudo {
		fmt.Printf("User `%s` is not sudo. (This is good, we can assign it) \n", *AdminUser)
		os.Exit(1)
	}

	fmt.Printf("User `%s` exists and is sudo (We do not change anything) \n", *AdminUser)

}