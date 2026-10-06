package booking

type handler struct {
	service BookingService
}

func NewHandler(s BookingService) *handler {
	return &handler{service: s}
}
