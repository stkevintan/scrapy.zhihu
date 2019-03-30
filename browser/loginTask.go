package browser

import (
	"context"
	"log"
	"net/http"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

const switchJS = `
		window.onload = function() {
			var parent = document.querySelector('.Login-content');
			if(parent != null) return;
			var elem = document.querySelector('.SignContainer-switch span');
			if(elem != null) elem.click();
		}
`

// LoginTask load the zhihu login page, and grap the cookies
func LoginTask(account Account, cookies *[]*http.Cookie) chromedp.Tasks {
	var ret *runtime.RemoteObject
	usernameSel := `//input[@name="username"]`
	passwordSel := `//input[@name="password"]`
	buttonSel := `//button[@type="submit"]`
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.zhihu.com/signup`),
		chromedp.Evaluate(switchJS, &ret),
		chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
			log.Printf("ret: %v\n", ret)
			return nil
		}),
		chromedp.WaitVisible(usernameSel),
		chromedp.SendKeys(usernameSel, account.Username),
		chromedp.SendKeys(passwordSel, account.Password),
		chromedp.Click(buttonSel, chromedp.NodeVisible),
		chromedp.WaitReady(`img.Avatar`),
		chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
			raw, err := network.GetAllCookies().Do(ctxt, h)
			if err != nil {
				return err
			}

			for _, item := range raw {
				cookie := &http.Cookie{}
				cookie.Name = item.Name
				cookie.Value = item.Value
				*cookies = append(*cookies, cookie)
			}

			return nil
		}),
	}
}
