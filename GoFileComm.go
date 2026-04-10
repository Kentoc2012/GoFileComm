package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	for {
		var an1 string
		fmt.Println("удалить, прочитать, создать файл, создать папку или просмотреть деректорию? (1/2/3/4/5)") //меню
		fmt.Scan(&an1)
		an, err := strconv.Atoi(an1)
		if err != nil {
			fmt.Println("errot", err)
		}
		if an != 1 && an != 2 && an != 3 && an != 4 && an != 5 {
			fmt.Println("разрешены только варианты в меню!")
			continue
		}
		if an == 1 { //удаление файлов
			var name string
			fmt.Println("введите название файла")
			fmt.Scan(&name)
			finalpath := filepath.Join(name)
			err := os.Remove(finalpath)
			if err != nil {
				fmt.Println("error", err)
			}
		}
		if an == 2 { //прочетние файлов (не знаю зачем, но пусть будет)
			var name string
			fmt.Println("введите путь к файлу")
			fmt.Scan(&name)
			finalpath := filepath.Join(name)
			text, err := ioutil.ReadFile(finalpath)
			if err != nil {
				fmt.Println("error", err)
			}
			fmt.Println("текст файла:\n", string(text))
		}
		if an == 4 { //создание папки (да 4, не 3)
			var dirname string
			fmt.Println("введите название папки или путь")
			fmt.Scan(&dirname)
			var finalpath string
			if !filepath.IsAbs(dirname) {
				currentdir, _ := os.Getwd()
				finalpath = filepath.Join(currentdir, dirname)
			} else {
				finalpath = dirname
			}
			err := os.Mkdir(finalpath, 0755)
			if err != nil {
				fmt.Println("ошибка создания файлов", err)
			}
		}
		if an == 5 { //просмотр папки (dir в консоли)
			var path string
			fmt.Println("введите путь для просмотра (или . для тукущей папки)")
			fmt.Scan(&path)
			var finalpath string
			if !filepath.IsAbs(path) {
				currentdir, _ := os.Getwd()
				finalpath = filepath.Join(currentdir, path)
			} else {
				finalpath = path
			}
			files, err := os.ReadDir(finalpath)
			if err != nil {
				fmt.Println("\033[31m[ERROR]\033[0m ошибка сканирования", err)
			} else {
				fmt.Println("Содержимое директории:", finalpath)
				fmt.Println("-----------------------------------")
				for _, file := range files {
					if file.IsDir() {
						fmt.Println("\033[34m[DIR]\033[0m ", file.Name())
					} else {
						fmt.Println("[FILE] ", file.Name())
					}
				}
				fmt.Println("-----------------------------------")
			}
		}
		if an == 3 { //создание файлов (да, вот и 3)
			scanner := bufio.NewScanner(os.Stdin)
			var text string
			var name string
			fmt.Println("введи название файла или путь")
			fmt.Scan(&name)
			currentdir, _ := os.Getwd()
			scanner.Scan() //продувка буфера от предыдущего Enter
			fmt.Println("введите текст файла")
			if scanner.Scan() {
				text = scanner.Text()
			}
			finalpath := filepath.Join(currentdir, name)
			text1 := []byte(text)
			ext := filepath.Ext(name) //авто замена пустоты на .txt
			if ext == "" {
				finalpath = name + ".txt"
			}
			err := ioutil.WriteFile(finalpath, text1, 0644) //создание файла
			if err != nil {
				fmt.Println("error\n", err)
			}
			fmt.Println("путь к файлу:", finalpath)
		}
	}
}
