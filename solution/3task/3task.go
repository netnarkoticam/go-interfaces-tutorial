package main

import (
	"fmt"
	"os"
)

type Writer interface {
	Write([]byte) (int, error)
}

type File struct {
	Filename string 
}

type Console struct{}

func (f File) Write(data []byte) (int, error) {
	file, err := os.OpenFile(f.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer func() {
		if file != nil {
			err := file.Close()
			if err != nil {
				fmt.Println("Ошибка закрытия файла:", err) 
			}
		}
	}() 

	return file.Write(data)
}

func (c Console) Write(data []byte) (int, error) {
	return fmt.Print(string(data))
}

func main() {
	fileWriter := File{Filename: "output.txt"}
	consoleWriter := Console{}

	text := "ЭЩКЕРЕЕЕЕЕЕЕЕЕЕЕЕЕЕ \n"

	n, err := fileWriter.Write([]byte(text))
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
		return
	}
	fmt.Printf("Записано %d байт в файл.\n", n)

	n, err = consoleWriter.Write([]byte(text))
	if err != nil {
		fmt.Printf("Ошибка записи в консоль: %v\n", err)
		return
	}
	fmt.Printf("Записано %d байт в консоль.\n", n)
}