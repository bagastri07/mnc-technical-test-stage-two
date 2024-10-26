package healthcheck

type Handler struct {
}

func New() *Handler {
	return new(Handler)
}
