package utility

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
)

func JsonFormat(str string) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(str), "", "    ")

	if err != nil {
		panic(err)
	}

	return buf.String()
}

func LoadEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Printf("環境変数を読み込めませんでした: %v", err)
	}
}
