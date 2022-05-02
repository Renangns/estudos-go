package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5

func main() {
	for {
		showMenu()
		comando := userInput()
		executeCommand(comando)
	}
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func userInput() (comando int) {
	fmt.Scan(&comando)
	return
}

func executeCommand(comando int) {
	switch comando {
	case 1:
		initMonitoring()
	case 2:
		readLogs()
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Comando não encontrado")
		os.Exit(-1)
	}
}

func initMonitoring() {
	fmt.Println("Monitorando...")

	urls := readUrls()

	for {
		for _, url := range urls {
			httpCall(url)
		}
		fmt.Println("")
		time.Sleep(delay * time.Minute)
	}

}

func httpCall(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode != 200 {
		fmt.Println("Site", url, "com problema status code", response.StatusCode)
		registerLog(url, false)
		return
	}

	fmt.Println("Site", url, "carregado com sucesso")
	registerLog(url, true)
}

func readUrls() []string {
	var urls []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return urls
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func readLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}
