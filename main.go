package main

import (
	"fmt"
	"github.com/therahulprasad/goAws/ses"
)

func main() {
	c, s, e1, e2 := ses.SendSingleMail("user@example.com", "Testing subject", "Testing body", "noreply@touchtalent.com", "Rahul", "Touchtalent", "", "")
	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(e1)
	fmt.Println(e2)
}
