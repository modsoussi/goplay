// Playing around with mandrill go package.
// Author: modsoussi
//
// (C) 2017. modsoussi.

package main

import mail "github.com/keighl/mandrill"
import "fmt"

func main() {
	client := mail.ClientWithKey("VmvF8-Qd5AiZEVl16PHIoQ")
	message := &mail.Message{}
	message.AddRecipient("mo@daycationapp.com", "Mo Soussi", "to")
	message.FromEmail = "no-reply@daycationapp.com"
	message.FromName = "Daycation"
	message.Subject = "Welcome to Daycation"
	message.Text = "You won!"

	r, err := client.MessagesSend(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r)
}
