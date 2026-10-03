package dto

type CreateRequestEvent struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Location     string `json:"location" binding:"required"`
	StartsAt     string `json:"start_at" binding:"required"`
	TotalTickets int    `json:"total_tickets" binding:"required"`
	Price        int    `json:"price" binding:"required"`
}


type UpdateRequestEvent struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Location     string `json:"location"`
	StartsAt     string `json:"start_at"`
	TotalTickets int    `json:"total_tickets"`
	Price        int    `json:"price"`
}