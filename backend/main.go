package main

import (
	"fmt"
	"github.com/huchengbei/for-my-girl/backend/pkg/setting"
	"github.com/huchengbei/for-my-girl/backend/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()


	s := &http.Server{
		Addr: 			fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: 		router,
		ReadTimeout:	setting.ReadTimeout,
		WriteTimeout: 	setting.WriteTimeout,
		MaxHeaderBytes:	1 << 20,
	}

	s.ListenAndServe()
}
