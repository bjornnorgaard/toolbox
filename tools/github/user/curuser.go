package user

import (
	"log"
	"os"
)

func Me() string {
	token := os.Getenv("CI")
	if len(token) == 0 {
		return "@me"
	}

	user := os.Getenv("GITHUB_REPOSITORY_OWNER")
	if len(user) == 0 {
		log.Fatalf("'GITHUB_REPOSITORY_OWNER' env var is not set")
	}

	return user
}
