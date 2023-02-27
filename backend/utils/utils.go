package utils

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func GetEnviromentalVariable(variableName string) string {
	envMap, err := godotenv.Read(".env")

	HandleError(err)
	return envMap[variableName]

}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func HandleErrorTest(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
