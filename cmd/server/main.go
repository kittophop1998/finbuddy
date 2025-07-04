package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kittiphop/finbuddy/internal/delivery/http"
)

func main() {
	r := gin.Default()

	http.WalletRouter(r)

	r.Run(":3004")
}
