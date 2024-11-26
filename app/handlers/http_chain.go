package handlers

type Handler interface {
	setNext(Handler)
}
