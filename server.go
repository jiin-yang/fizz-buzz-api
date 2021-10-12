package main

import (
	"fizz-buzz-api/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tylerb/graceful"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	e *echo.Echo
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	e := echo.New()
	e.HideBanner = true
	server := &Server{e: e, config: config}

	return server
}

func (s *Server)Start() error {
	s.e.Server.Addr = fmt.Sprintf(":%d", s.config.Server.Port)
	s.e.GET("/:count", GetResult)
	s.e.GET("/health", s.healthCheck)

	return graceful.ListenAndServe(s.e.Server, time.Second * 10)
}

func (s *Server)healthCheck(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func GetResult(ctx echo.Context) error {
	number := ctx.Param("count")
	count, _ := strconv.Atoi(number)

	var fizzBuzzArray []string

	//result := map[string]interface{}{}
	result := map[string][]string{}

	for i := 1; i <= count; i++ {
		if i % 15 == 0{
			fizzBuzzArray = append(fizzBuzzArray, "FizzBuzz")
		} else if i % 3 == 0{
			fizzBuzzArray = append(fizzBuzzArray, "Fizz")
		} else if i % 5 == 0{
			fizzBuzzArray = append(fizzBuzzArray, "Buzz")
		} else{
			fizzBuzzArray = append(fizzBuzzArray, strconv.Itoa(i))
		}
	}

	result["fizzbuzz"] = fizzBuzzArray

	status := http.StatusOK
	return ctx.JSON(status, result)
}