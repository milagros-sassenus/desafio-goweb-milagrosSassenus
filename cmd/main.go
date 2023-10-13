package main

import (
	"desafio-goweb-milagrosSassenus/cmd/internal/domain"
	"desafio-goweb-milagrosSassenus/cmd/internal/tickets"
	"desafio-goweb-milagrosSassenus/cmd/server/handler"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	list, err := LoadTicketsFromFile("tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	repo := tickets.NewRepository(list)
	s := tickets.NewService(repo)
	h := handler.NewTicket(s)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	// Rutas a desarollar:

	// GET - “/ticket/getByCountry/:dest”
	r.GET("/ticket/getByCountry/:dest", h.GetTicketsByCountry())
	// GET - “/ticket/getAverage/:dest”
	r.GET("/ticket/getAverage/:dest", h.AverageDestination())

	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
