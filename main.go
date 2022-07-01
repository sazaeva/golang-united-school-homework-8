package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
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
	var filename = flag.String("filename", "", "filename flag")
	flag.Parse()

	return Arguments{
		"id":        *id,
		"operation": *operation,
		"item":      *item,
		"filename":  *filename,
	}
}

func Perform(args Arguments, writer io.Writer) error {
	var operation func(Arguments) (string, error)

	if args["operation"] == "" {
		return fmt.Errorf("-operation flag has to be specified")
	}

	if args["filename"] == "" {
		return fmt.Errorf("-filename flag has to be specified")
	}

	switch args["operation"] {
	case "add":
		if args["item"] == "" {
			return fmt.Errorf("-item flag has to be specified")
		}
		operation = add
	case "list":
		operation = list
	case "remove":
		if args["id"] == "" {
			return fmt.Errorf("-id flag has to be specified")
		}
		operation = remove
	case "findById":
		if args["id"] == "" {
			return fmt.Errorf("-id flag has to be specified")
		}
		operation = findById
	default:
		return fmt.Errorf("Operation %s not allowed!", args["operation"])
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
	item := args["item"]

	file, err := os.OpenFile(args["filename"], os.O_RDWR|os.O_CREATE, os.ModeAppend)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(args["filename"])
	if err != nil {
		return "", err
	}

	_, err = file.WriteAt([]byte(item), int64(len(content)))
	if err != nil {
		return "", err
	}

	return "", nil
}
func list(args Arguments) (string, error) {
	filename := args["filename"]

	if !checkFileExists(filename) {
		fmt.Println("File does not exists")
		return "", nil
	}

	content, err := readFile(filename)
	if err != nil {
		return "", err
	}

	for _, val := range content {
		fmt.Println(val)
	}

	return "", nil
}
func remove(args Arguments) (string, error) {
	//userId := args["id"]

	return "", nil
}

func findById(args Arguments) (string, error) {
	userId := args["id"]
	filename := args["filename"]

	users := make([]Users, 0)

	content, err := readFile(filename)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(content, &users)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, t := range users {
		if t.Id == userId {
			fmt.Printf("User found: %s", t.Email)
			return t.Email, nil
		}
	}

	return "", err
}

func readFile(filename string) ([]byte, error) {
	if !checkFileExists(filename) {
		fmt.Println("File does not exists")
		return nil, errors.New("File does not exists")
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return content, err
	}

	return content, nil
}

func checkFileExists(filename string) bool {
	_, err := os.OpenFile(filename, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return false
	}

	return true
}

func writeFile(filename string, fileMode fs.FileMode, message []byte) bool {
	err := os.WriteFile(filename, message, fileMode)
	if err != nil {
		return false
	}

	return true
}

func main() {
	//existingItems := "[{\"id\":\"1\",\"email\":\"test@test.com\",\"age\":34},{\"id\":\"2\",\"email\":\"test2@test.com\",\"age\":32}]"
	//if !checkFileExists(args["filename"]) {
	//  createFile(args["filename"], filePermission, []byte(existingItems))
	//} else {
	//  writeFile(args["filename"], filePermission, []byte(existingItems))
	//}

	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
