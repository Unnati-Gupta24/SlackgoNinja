package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main(){
	os.Setenv("SLACK_BOT_TOKEN","xapp-1-A07MECABRB7-77339-5b4e56548a79307da0e70189541847719b1ca29da7")
	os.Setenv("CHANNEL_ID","D07MKPECL")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"linkedin.txt"}

	for i := 0; i<len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil{
			fmt.Printf("%s\n",err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n",file.Name,file.URL)
	}

}
