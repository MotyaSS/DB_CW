package main

import (
  "log"

  "github.com/gin-gonic/gin"
)

func main() {
  a := gin.Default()
  log.Fatal(a.Run(":8080"))
}
