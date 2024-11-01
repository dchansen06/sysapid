package main

import (
	"os"	// For hostname in getData and collectData
	"os/exec"	// For execCommand

	"github.com/gin-gonic/gin"	// As REST-API server framework
)

/* JSON format to send data in */
type data struct {
	ID	string	`json:"id"`	// Hostname
	CPU	string	`json:"cpu"`	// Percent
	TEMP	string	`json:"temp"`	// Celcius
	CONN	string	`json:"conn"`	// Number active
	RAM	string	`json:"ram"`	// Percent
	SWAP	string	`json:"swap"`	// Percent
	DISK	string	`json:"disk"`	// Percent
}

/* Set up server and get it running */
func main() {
	gin.SetMode(gin.ReleaseMode)	// Disables all error messages

	router := gin.Default()
	router.SetTrustedProxies(nil)	// Block proxy requests

	router.GET("/", getData)
	router.GET("/index.html", getIndex)	// Joke
	router.GET("/favicon.ico", getFavicon)	// Joke

	hostname, _ := os.Hostname()
	switch hostname {
		case "raspberrypi":
			router.Run("127.0.0.0:80")	// Change to fit your needs ("hides" from other devices on network)
			break
		default:
			router.Run("localhost:8080")	// Change to fit your needs (does not hide from other devices on network)
	}
}

/* Joke function */
func getIndex(context *gin.Context) {
	context.String(418, "Error 418: I'm A Teapot\n\nThe server refuses the attempt to brew coffee with a teapot.")
}

/* Joke function */
func getFavicon(context *gin.Context) {
	context.String(415, "Error 415: Unsupported Media Type\n\nThe media format of the requested data is not supported by the server, so the server is rejecting the request.")
}

/* Request handler for root directory */
func getData(context *gin.Context) {
	dataSelf := data{}

	switch context.Query("fast") {
		case "cpu":
			dataSelf.CPU = collectData("cpu")
			break
		case "temp":
			dataSelf.TEMP = collectData("temp")
			break
		case "conn":
			dataSelf.CONN = collectData("conn")
			break
		case "ram":
			dataSelf.RAM = collectData("ram")
			break
		case "swap":
			dataSelf.SWAP = collectData("swap")
			break
		case "disk":
			dataSelf.DISK = collectData("disk")
			break
		default:
			dataSelf.CPU = collectData("cpu")
			dataSelf.TEMP = collectData("temp")
			dataSelf.CONN = collectData("conn")
			dataSelf.RAM = collectData("ram")
			dataSelf.SWAP = collectData("swap")
			dataSelf.DISK = collectData("disk")
			break
	}

	hostname, _ := os.Hostname()
	dataSelf.ID = hostname
	context.SecureJSON(200, dataSelf)	// Supposed to be safer, same output as normal JSON though
}

/* Command executor, returns STDOUT as string */
func execCommand(command string) string {
	dat, err := exec.Command("/usr/bin/bash", "-c", command).Output()	// Default to run as bash

	if err != nil {
		return "Error, execCommand failed somewhere"	// Command failed to run
	}

	return string(dat)
}

/* Determine which variation(s) if any the host supports, then run commands to collect data */
func collectData(dataType string) string {
	switch dataType {
		case "cpu":
			return string(execCommand("{ /usr/bin/cat /proc/stat; /usr/bin/sleep 0.06; /usr/bin/cat /proc/stat; } | /usr/bin/awk '/^cpu / {usr=$2-usr; sys=$4-sys; idle=$5-idle; iow=$6-iow} END {total=usr+sys+idle+iow; printf \"%.2f%%\", (total-idle)*100/total}'"))
		case "conn":
			return string(execCommand("/usr/bin/netstat -an | /usr/bin/grep ESTABLISHED | /usr/bin/wc -l | /usr/bin/sed 's/$/ active/' | /usr/bin/tr -d '\\n'"))
		case "ram":
			return string(execCommand("/usr/bin/free -m | /usr/bin/awk 'FNR == 2 {printf(\"%0.2f%%\", $3/$2*100)}'"))
		case "swap":
			return string(execCommand("/usr/bin/free -m | /usr/bin/awk 'FNR == 3 {printf(\"%0.2f%%\", $3/$2*100)}'"))
		case "disk":
			return string(execCommand("/usr/bin/df --type=ext4 | /usr/bin/awk 'FNR == 2 {printf $5\"%\"}'"))
		case "temp":	// Not everything is the same for all computers
			hostname, _ := os.Hostname()
			switch hostname {
				case "ubuntu":	// Change to fit your needs
					return string(execCommand("/usr/bin/sensors | /usr/bin/grep 'Tctl:' | /usr/bin/awk '{printf $2}' | /usr/bin/tr -d +"))
				case "raspberrypi":	// Change to fit your needs
					return string(execCommand("/usr/bin/cat /sys/class/thermal/thermal_zone0/temp | /usr/bin/awk '{printf(\"%0.2f°C\", $1/1000)}'"))
				default:
					return "Error, unknown hostname in collectData"	// Unknown
			}
		default:
			return "Error, unknown dataType in collectData"	// Request for data collection was malformed and incorrect
	}
}
