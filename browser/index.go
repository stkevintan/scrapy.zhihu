package browser

import (
	"context"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
)

//Account is a func
type Account struct {
	Username string
	Password string
}

//Launch is a func
func Launch(account Account) []*http.Cookie {
	if account.Username == "" || account.Password == "" {
		return []*http.Cookie{}
	}

	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))

	if err != nil {
		log.Fatal(err)
	}
	var cookies = make([]*http.Cookie, 0)
	err = c.Run(ctxt, LoginTask(account, &cookies))

	if err != nil {
		log.Fatal(err)
	}

	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
	return cookies
}