package main

import (
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	EmailAddress string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("[Email] Sending to %s: %s\n", e.EmailAddress, message)
	return nil
}

type SlackNotifier struct {
	ChannelID string
}

func (s SlackNotifier) Send(message string) error {
	fmt.Printf("[Slack] Pinging channel %s: %s\n", s.ChannelID, message)
	return nil
}

func AlertSystem(n Notifier, msg string) {
	fmt.Println("System is preparing to send alert...")

	err := n.Send(msg)
	if err != nil {
		fmt.Println("Failed to send alert:", err)
		return
	}
	fmt.Println("Alert sent successfully!\n---")
}

func main() {
	myEmail := EmailNotifier{EmailAddress: "hodtduy.work@gmail.com"}
	mySlack := SlackNotifier{ChannelID: "#alert-prod"}

	AlertSystem(myEmail, "CPU Usage > 90%")
	AlertSystem(mySlack, "Database connection lost!")
	
	fmt.Println(">>> Broadcasting to all channels:")

	allNotifiers := []Notifier{myEmail, mySlack}
	for _, notifier := range allNotifiers {
		notifier.Send("System Rebooting in 5 min. Save your work!")
	}
}