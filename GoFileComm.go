package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

const ( //цветовые коды
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func main() {
	for {
		var an1 string
		fmt.Println("удалить, прочитать, создать файл, создать папку, просмотреть деректорию, выполнить файл или показать дополнительный текст? (1/2/3/4/5/6/help) напишите exit для выхода") //меню
		fmt.Scan(&an1)
		if an1 == "help" || an1 == "h" {
			fmt.Println("GoFileComm, это файловый менеджер созданный новичком в Go, так что не судите строго\n Функционал:\n 1. удаление файлов\n 2. прочтение файлов (не знаю зачем, но пусть будет)\n 3. создание файлов\n 4. создание папки\n 5. просмотр деректории (папки)\n 6. выполнение файла\n \n разделение создания файла и деректории обусловленно удобством, чтобы не дописывать всегда .dir в конце названия.")
			continue
		}
		if an1 == "exit" || an1 == "e" {
			return
		}
		an, err := strconv.Atoi(an1)
		if err != nil {
			fmt.Println(Red+"errot", err)
		}

		if an != 1 && an != 2 && an != 3 && an != 4 && an != 5 && an != 6 {
			fmt.Println(Green + "разрешены только варианты в меню!" + Reset)
			continue
		}
		if an == 1 { //удаление файлов
			var name string
			fmt.Println("введите название файла")
			fmt.Scan(&name)
			finalpath := filepath.Join(name)
			err := os.Remove(finalpath)
			if err != nil {
				fmt.Println(Red+"error"+Reset, err)

			} else {
				fmt.Println(Green + "удаление завершено" + Reset)
			}
		}
		if an == 2 { //прочетние файлов (не знаю зачем, но пусть будет)
			var name string
			fmt.Println("введите путь к файлу")
			fmt.Scan(&name)
			finalpath := filepath.Join(name)
			text, err := ioutil.ReadFile(finalpath)
			if err != nil {
				fmt.Println(Red+"error"+Reset, err)
			}
			fmt.Println("текст файла:\n", Yellow+string(text)+Reset)
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
				fmt.Println(Red+"ошибка создания папки"+Reset, err)
			}
		}
		if an == 5 { //просмотр папки (dir в консоли)
			var path string
			fmt.Println(Cyan + "введите путь для просмотра (или . для тукущей папки)" + Reset)
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
				fmt.Println(Cyan+"Содержимое директории:", finalpath)
				fmt.Println("-----------------------------------")
				for _, file := range files {
					if file.IsDir() {
						fmt.Println("\033[34m[DIR]\033[0m ", file.Name()+Reset)
					} else {
						fmt.Println("[FILE] ", file.Name())
					}
				}
				fmt.Println("-----------------------------------")
			}
		}
		if an == 6 {
			var finalpath string
			var filename string
			fmt.Println("укажите путь для запуска или название, если он находится в текущей папке")
			fmt.Scan(&filename)
			if !filepath.IsAbs(filename) {
				currentdir, _ := os.Getwd()
				finalpath = filepath.Join(currentdir, filename)
			} else {
				finalpath = filename
			}
			cmd := exec.Command("cmd", "/C", "start", "", finalpath)
			err := cmd.Run()
			if err != nil {
				fmt.Println("\033[31m[ERROR]\033[0m ошибка запуска:", err)
			} else {
				fmt.Println(Blue + "запуск произведён успешно")
			}
		}
		if an == 3 { //создание файлов (да, вот и 3)
			scanner := bufio.NewScanner(os.Stdin)
			var text string
			var name string
			fmt.Println("введите название файла или путь")
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
				fmt.Println(Red+"error\n", err)
			}
			fmt.Println(Yellow+"путь к файлу:"+Reset, finalpath)
		}
	}
}
