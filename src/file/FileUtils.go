package file

import (
	"bufio"
	"giogii/src/entity"
	"io"
	"log"
	"os"
	"strings"
)

func ReadConfig(fileName string, flashbackInfo entity.FlashbackInfo) (entity.FlashbackInfo, error) {
	path := "./" + fileName
	f, err := os.Open(path)
	if err != nil {
		log.Println("Failed to open config file")
		log.Fatal(err)
		os.Exit(-1)
	}
	reader := bufio.NewReader(f)
	for {
		readLine, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		line := string(readLine)
		if line == "" {
			break
		}
		key := strings.Split(line, "=")[0]
		value := strings.Split(line, "=")[1]

		switch key {
		case "sourceUserInfo":
			flashbackInfo.SetSourceUserInfo(value)
		case "targetUserInfo":
			flashbackInfo.SetTargetUserInfo(value)
		case "sshInfo":
			flashbackInfo.SetSshInfo(value)
			flashbackInfo.SetSshUser(strings.Split(value, ":")[0])
			flashbackInfo.SetSshPass(strings.Split(value, ":")[1])
		case "sshPort":
			if strings.Trim(value, " ") == "" {
				value = "22"
			}
			flashbackInfo.SetSshPort(value)
		default:
			break
		}
	}

	return flashbackInfo, nil
}
