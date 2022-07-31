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

const monitoringTimes = 3
const delay = 5

func main() {

	showStartMessage()

	for {
		showMenu()

		choice := registerUserChoice()

		fmt.Println("You chose option ", choice)

		switch choice {

		case 1:
			fmt.Println("Monitoring")
			startMonitoring()

		case 2:
			fmt.Println("Logs")
			showLogs()

		case 0:
			fmt.Println("Exiting the program")
			os.Exit(0)

		default:
			os.Exit(-1)

		}
	}

}

func showStartMessage() {
	name := "Douglas"
	version := 1.2
	fmt.Println("Hi mr.", name)
	fmt.Println("This program is in version: ", version)

}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func registerUserChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

func startMonitoring() {

	sites := readSitesFromFile()

	for i := 0; i < monitoringTimes; i++ {
		for _, site := range sites {
			fmt.Println("Monitoring site", site)
			monitorSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

}

func monitorSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ops! Something is wrong: ", err)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("The website is working correctly")
			registerLog(site, true)
		} else {
			fmt.Println("The website id down")
			registerLog(site, false)
		}
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ops! Something is wrong: ", err)
	} else {

		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			sites = append(sites, line)
			fmt.Println(line)

			if err == io.EOF {
				break
			}

			// if err == io.EOF {
			// 	break
			// } else {
			// 	line = strings.TrimSpace(line)
			// 	sites = append(sites, line)
			// 	fmt.Println(line)
			// }

		}

		fmt.Println(sites)

	}

	file.Close()

	return sites
}

func registerLog(site string, online bool) {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		println("Ops! Something is wrong: ", err)
	} else {
		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + "Site: " + site + " - " + "Online: " + strconv.FormatBool(online) + "\n")
	}

	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Ops! SOmething is wrong: ", err)
	} else {
		fmt.Println(string(file))
	}

}
