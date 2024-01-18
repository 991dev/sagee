package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
)

const baseURL = "http://localhost:1234" // Замените на ваш URL

var (
	greenBold = color.New(color.FgGreen, color.Bold).SprintFunc()
	redBold   = color.New(color.FgRed, color.Bold).SprintFunc()
)

func main() {
	for {
		printMenu()

		fmt.Print("Выберите действие: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			getUsers()
		case "2":
			createUser()
		case "3":
			getChats()
		case "4":
			createChat()
		case "5":
			getMessages()
		case "6":
			sendMessage()
		case "0":
			fmt.Println("Программа завершена.")
			return
		default:
			fmt.Println(redBold("Неверный ввод. Попробуйте еще раз."))
		}
	}
}

func printMenu() {
	fmt.Println(greenBold("1. Получить пользователей"))
	fmt.Println(greenBold("2. Создать пользователя"))
	fmt.Println(greenBold("3. Получить чаты"))
	fmt.Println(greenBold("4. Создать чат"))
	fmt.Println(greenBold("5. Получить сообщения"))
	fmt.Println(greenBold("6. Отправить сообщение"))
	fmt.Println(greenBold("0. Выйти"))
}

func getUsers() {
	resp, err := resty.New().R().Get(baseURL + "/users")
	handleResponse(resp, err)
}

func createUser() {
	fmt.Print("Введите имя пользователя: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()

	resp, err := resty.New().R().
		SetBody(map[string]interface{}{"username": username}).
		Post(baseURL + "/users")
	handleResponse(resp, err)
}

func getChats() {
	resp, err := resty.New().R().Get(baseURL + "/chats")
	handleResponse(resp, err)
}

func createChat() {
	fmt.Print("Введите название чата: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	chatName := scanner.Text()

	resp, err := resty.New().R().
		SetBody(map[string]interface{}{"chatName": chatName}).
		Post(baseURL + "/chats")
	handleResponse(resp, err)
}

func getMessages() {
	resp, err := resty.New().R().Get(baseURL + "/messages")
	handleResponse(resp, err)
}

func sendMessage() {
	fmt.Print("Введите ID пользователя: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	senderID := scanner.Text()

	fmt.Print("Введите текст сообщения: ")
	scanner.Scan()
	messageText := scanner.Text()

	fmt.Print("Введите ID чата: ")
	scanner.Scan()
	chatID := scanner.Text()

	resp, err := resty.New().R().
		SetBody(map[string]interface{}{
			"senderID":    senderID,
			"messageText": messageText,
			"chatID":      chatID,
		}).
		Post(baseURL + "/messages")
	handleResponse(resp, err)
}

func handleResponse(resp *resty.Response, err error) {
	if err != nil {
		fmt.Println(redBold("Ошибка при выполнении запроса:"), err)
		return
	}

	if resp.StatusCode() == 200 {
		fmt.Println(greenBold("Успешно выполнено!"))
	} else {
		fmt.Println(redBold("Ошибка:"), resp.Status())
		fmt.Println("Тело ответа:", resp.String())
	}
}
