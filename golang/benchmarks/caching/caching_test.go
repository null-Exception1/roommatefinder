package caching

import (
	initfuncs "golang/init"
	"testing"

	"github.com/joho/godotenv"
)

func TestCaching(t *testing.T) {
	godotenv.Load("../../.testenv")

	initfuncs.Database()

}
