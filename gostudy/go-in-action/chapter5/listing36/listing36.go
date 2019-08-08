package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	var u notifier
	// u = &user{"Bill", "bill@email.com"}
	u = &user{"Bill", "bill@email.com"}

	sendNotification(u)
}

func sendNotification(n notifier) {
	n.notify()
}
