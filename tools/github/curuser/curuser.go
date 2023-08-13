package curuser

import (
	"log"
	"os"
)

func Me() string {
	token := os.Getenv("CI")
	if len(token) == 0 {
		return "@me"
	}

	log.Printf("Detected CI env, continuing with repo owner as user")

	user := os.Getenv("GITHUB_REPOSITORY_OWNER")
	if len(user) == 0 {
		log.Fatalf("'GITHUB_REPOSITORY_OWNER' env var is not set")
	}

	return user
}
