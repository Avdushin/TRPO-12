package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	dataFolder = "data"
	ciperFile  = "data/шифр.txt"
	decFile    = "data/расшифровка.txt"
)

func main() {
	// Шифровальная таблица Полибия
	poliby := [][]string{
		{"A", "B", "C", "D", "E"},
		{"F", "G", "H", "I", "J", "K"},
		{"L", "M", "N", "O", "P"},
		{"Q", "R", "S", "T", "U"},
		{"V", "W", "X", "Y", "Z"},
	}

	// Создадим папку для файликов
	if _, err := os.Stat(dataFolder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dataFolder, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	// a) Зашифровываем текст в файл
	fmt.Print("Введите текст для шифрования: ")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	cipherText := ""

	for _, char := range strings.ToUpper(message) {
		if char >= 'A' && char <= 'Z' {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if poliby[i][j] == string(char) {
						cipherText += fmt.Sprintf("%v%v", i+1, j+1)
					}
				}
			}
		} else if char == ' ' {
			cipherText += " "
		}
	}

	file, err := os.Create(ciperFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Шифр: %v\n", cipherText)

	_, err = file.WriteString(cipherText)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fmt.Println("Текст успешно зашифрован и сохранен в файле", ciperFile)

	// б) Считываем зашифрованный текст из файла и расшифровываем его

	cipherFile, err := os.Open(ciperFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cipherFile.Close()

	cipherReader := bufio.NewReader(cipherFile)
	cipherText, _ = cipherReader.ReadString('\n')

	// Расшифровываем текст
	var plainText string
	for i := 0; i < len(cipherText); i += 2 {
		if i+1 < len(cipherText) {
			row, col := cipherText[i:i+1], cipherText[i+1:i+2]
			rowNum, colNum := int(row[0]-'0')-1, int(col[0]-'0')-1

			// Жалкая попытка обработать все прочие символы, которые не входят в диапазон алфавита
			if rowNum >= 0 && rowNum < 5 && colNum >= 0 && colNum < 5 {
				plainText += poliby[rowNum][colNum]
			} else {
				plainText += " "
			}

		}
	}

	fmt.Println("Расшифрованный текст: ", plainText)

	decrptFile, err := os.Create(decFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = decrptFile.WriteString(plainText)
	if err != nil {
		fmt.Println(err)
		return
	}
}
