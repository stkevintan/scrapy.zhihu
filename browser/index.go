package browser

import (
	"context"
	"log"
	"net/http"

	"github.com/chromedp/chromedp"
)

//Account is a func
type Account struct {
	Username, Password string
}

//Launch is a func
func Launch(ctx context.Context, account Account) []*http.Cookie {
	if account.Username == "" || account.Password == "" {
		return []*http.Cookie{}
	}

	ctxt, cancel := context.WithCancel(ctx)
	defer cancel()
	c, err := chromedp.New(ctxt)

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
