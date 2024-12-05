package contexts

type BodyContext struct {
	Body           string
	HttpStatusCode int
	ContentType    string
	Headers        map[string]string
}
