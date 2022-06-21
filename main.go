package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

type Users struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Arguments map[string]string

func parseArgs() Arguments {
	var id = flag.String("id", "", "id flag")
	var operation = flag.String("operation", "", "operation flag")
	var item = flag.String("item", "", "item flag")
	var fileName = flag.String("filename", "", "filename flag")
	flag.Parse()

	return Arguments{
		"id":        *id,
		"operation": *operation,
		"item":      *item,
		"fileName":  *fileName,
	}
}

func Perform(args Arguments, writer io.Writer) error {
	var operation func(Arguments) (string, error)
	switch args["operation"] {
	case "add":
		operation = add
	case "list":
		operation = list
	case "remove":
		operation = remove
	case "findById":
		operation = findById
	case "":
		return fmt.Errorf("-operation flag has to be specified")
	default:
		return fmt.Errorf("Expect error to be '%s', but got '%s'", args["operation"])
	}
	str, err := operation(args)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}

func add(args Arguments) (string, error) {

}
func list(args Arguments) (string, error) {

}
func remove(args Arguments) (string, error) {

}
func findById(args Arguments) (string, error) {
	jsonUsers := `[{"id": "1", "email": "test@test.com", "age": 31},
{"id": "2", "email": "test2@test.com", "age": 41}]`

	user := []Users{}

	err := json.Unmarshal([]byte(jsonUsers), &user)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, t := range user {
		fmt.Println(t.Id)
	}
	return "", err
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
