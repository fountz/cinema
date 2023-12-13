package main

import (
	"fmt"

	"github.com/fountz/cinema/movie"
	"github.com/fountz/cinema/ticket"
)

func init() {
	fmt.Println("movieName")
}

func main() {
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println(movieName)
	movie.ReviewMovie(movieName, 9.4)
	ticket.BuyTicket(movieName)
}
