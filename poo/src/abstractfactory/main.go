package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

////////////////////////////////////////////////////////////
type SMS struct {
}
type Email struct {
}

func (SMS) SendNotification() {
	fmt.Println("Send via SMS")
}
func (Email) SendNotification() {
	fmt.Println("Send via Email")
}

/////////////////////////////////////////////////////////////////
func (SMS) GetSender() ISender {
	return SMSnotificationSender{}
}

func (Email) GetSender() ISender {
	return EmailnotificationSender{}
}

//////////////////////////////////////////////////////////////////
type SMSnotificationSender struct {
}

func (SMSnotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSnotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailnotificationSender struct {
}

func (EmailnotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailnotificationSender) GetSenderChannel() string {
	return "Gmail"
}

//_______________________________________________________
//
func getNotificationFactory(notificationtype string) (INotificationFactory, error) {
	if notificationtype == "SMS" {
		return &SMS{}, nil
	}
	if notificationtype == "Email" {
		return &Email{}, nil
	}
	return nil, fmt.Errorf("No notification factory")
}

//

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}
func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

//

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")
	SendNotification(smsFactory)
	SendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
