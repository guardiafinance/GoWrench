package handlers

import (
	"fmt"
	"io"
	"net/http"
	contexts "wrench/app/contexts"
)

type HttpRequestClientHandler struct {
	Next Handler
}

func (httpRequestClientHandler *HttpRequestClientHandler) Do(wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {
	resp, err := http.Get("http://localhost:8085")

	if err != nil {
		fmt.Printf("error")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body[:])

	if err != nil {
		fmt.Println(bodyString)
	}

	bodyContext.Body = bodyString
	bodyContext.HttpStatusCode = resp.StatusCode

	if httpRequestClientHandler.Next != nil {
		httpRequestClientHandler.Next.Do(wrenchContext, bodyContext)
	}
}

func (httpRequestClientHandler *HttpRequestClientHandler) SetNext(handler Handler) {
	httpRequestClientHandler.Next = handler
}
