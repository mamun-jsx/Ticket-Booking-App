package dto

import "time"

type ResponseEvent struct {
	ID               uint64    `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Location         string    `json:"location"`
	StartsAt         time.Time `json:"start_at"`
	TotalTickets     int       `json:"total_tickets"`
	AvailableTickets int       `json:"available_tickets"`
	Price            int       `json:"price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}


// type PaginationResponse struct {
// 	Total      int             `json:"total"`
// 	Current    int             `json:"current_page"`
// 	Limit      int             `json:"per_page"`
// 	TotalPages int             `json:"total_pages"`
// 	Events     []ResponseEvent `json:"events"`
// }