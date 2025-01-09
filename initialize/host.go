package initialize

import (
	http2 "github.com/advanced-go/agency/http"
	"github.com/advanced-go/agency/module"
	"github.com/advanced-go/common/access"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/host"
	"net/http"
	"time"
)

func Host(cmdLine []string) error {
	// Initialize host proxy for all HTTP handlers,and add intermediaries
	host.SetHostTimeout(time.Second * 3)
	host.SetAuthExchange(AuthHandler, nil)
	err := host.RegisterExchange(module.Authority, host.NewAccessLogIntermediary(access.InternalTraffic, http2.Exchange))
	return err
}

func AuthHandler(r *http.Request) (*http.Response, *core.Status) {
	/*
		if r != nil {
			tokenString := r.Header.Get(host.Authorization)
			if tokenString == "" {
				status := core.NewStatus(http.StatusUnauthorized)
				return &http.Response{StatusCode: status.HttpCode()}, status
				//w.WriteHeader(http.StatusUnauthorized)
				//fmt.Fprint(w, "Missing authorization header")
			}
		}


	*/
	return &http.Response{StatusCode: http.StatusOK}, core.StatusOK()
}
