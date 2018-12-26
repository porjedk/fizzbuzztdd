package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	fb "github.com/porjedk/fizzbuzztdd/fizzbuzz"
)

func main() {
	// Echo instance
	e := echo.New()

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//getFizzBuzz API
	e.GET("/getFizzBuzz/:number", func(c echo.Context) error {
		//get number parameter from Context
		numberParam := c.Param("number")
		//Convert number parameter to integer and 
		//assign return value to number parameter and error parameter
		number, err := strconv.Atoi(numberParam)
		//If can not convert number parameter then 
		//return bad request status with error message
		if err != nil {
			return c.String(http.StatusBadRequest, "Please input only number!!")
		}

		//Calling FizzBuzz function from fizzbuzz package 
		//and return the result as String
		return c.String(http.StatusOK, fb.FizzBuzz(number))
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}