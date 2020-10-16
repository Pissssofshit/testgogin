package main

import (
	"fmt"
	"github.com/Pissssofshit/testgogin/routers"
	"github.com/Pissssofshit/testgogin/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
