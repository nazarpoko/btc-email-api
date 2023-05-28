//the servise consists of 3 functions
//rate
//subscribe
//sendEmails
//emails are stored in emails.txt

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const BtcAddress = "https://btc-trade.com.ua/api/ticker/btc_uah"

func main() {
	router := gin.Default()
	router.GET(rate)
	router.POST(subscribe)
	router.POST(sendEmails)

	router.Run("localhost:8080")
}

func rate(c *gin.Context) {
	resp, err := http.Get(BtcAddress)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	//Convert the body to type string
	sb := string(body)
	c.IndentedJSON(http.StatusOK, sb)
}

func subscribe(c *gin.Context) {

	// aimed to read existing emails form txt and check for duplicates

}

func sendEmails(c *gin.Context) {
	// Set your SendGrid API key
	apiKey := "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"
	body := rate()
	// Create a SendGrid client
	client := sendgrid.NewSendClient(apiKey)

	// Define the email content
	from := mail.NewEmail("Sender Name", "jackfortiethmile@gmail.com")
	subject := "Test Email"

	// Iterate over the recipients and send the email individually
	filePath := "emails.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	str_content := string(content)
	maillist := strings.Split(str_content, "\n")
	for _, recipient := range maillist {
		to := mail.NewEmail("", recipient)
		message := mail.NewSingleEmail(from, subject, to, body, body)
		response, err := client.Send(message)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", recipient, err)
			continue
		}
		if response.StatusCode >= 200 && response.StatusCode < 300 {
			log.Printf("Email sent successfully to %s", recipient)
		} else {
			log.Printf("Failed to send email to %s. Status: %d", recipient, response.StatusCode)
		}
	}
}
