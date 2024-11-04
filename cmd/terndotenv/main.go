package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command("tern", "migrate", "--migrations", "./internal/store/pgstore/migrations", "--config", "./internal/store/pgstore/migrations/tern.conf")

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Command execution failed with error: ", err)
		fmt.Println("Output: ", string(output))
		panic(err)
	}

	fmt.Println("Command executed successfully: ", string(output))
}
