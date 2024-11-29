package handlers

import (
	"fmt"
	"io"
	"net/http"
	contexts "wrench/app/contexts"
)

type HttpRequestClientHandler struct {
	next Handler
}

func HttpRequestClient(wrenchContext contexts.WrenchContext, bodyContext contexts.BodyContext) {
	resp, err := http.Get("http://example.com/")

	if err != nil {
		fmt.Printf("error")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		str := string(body[:])
		fmt.Println(str)
	}
}
