package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newadmin.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
