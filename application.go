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
const monitoring = 5

func main() {
	greetings()
	for  {
		showMenu()
		command := readCommand()
		executeCommand(command)
	}
}

func executeCommand(command int) {
	switch command {
	case 1:
		initMonitoring()
	case 2:
		fmt.Println("Show logs")
		printLogs()
	case 0:
		fmt.Println("Quit application")
		os.Exit(0)
	default:
		fmt.Println("Command not found")
		os.Exit(-1)
	}
}

func showMenu() {
	fmt.Println("1 - Initialize Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Quit Application")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("Command selected:", command)
	return command
}

func greetings() {
	fmt.Println("Ola Mr. ? What is your name ?")
	name, _ := getNameAndAge()
	version := 1.1
	fmt.Println("Hello Mr.", name)
	fmt.Println("Program version ", version)
}

func getNameAndAge() (string, int) {
	name := "Victor Alcantara"
	age := 31
	return name, age
}

func initMonitoring() {
	fmt.Println("Monitoring applications...")

	sites := readSitesFromFile()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Monitoring site ", i, ":", site)
			checkSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error has occurred during the file reading")
	}

	fileReader := bufio.NewReader(file)
	for  {
		line, err := fileReader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func checkSite(site string) {
	fmt.Println("checking website:", site)
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("An error has occurred during the get request to web site:", site)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "is alive. Status code:", resp.StatusCode)
		createLogs(site, true)
	} else {
		fmt.Println("Site", site, "cannot be accessed.", "Status code:", resp.StatusCode)
		createLogs(site, false)
	}
}

func createLogs(site string, status bool)  {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs()  {
	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}