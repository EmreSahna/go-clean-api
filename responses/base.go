package responses

type Response struct {
	Status int
	Body   interface{}
}

func (r *Response) SetStatusCode(statusCode int) *Response {
	r.Status = statusCode
	return r
}

func (r *Response) SetBody(body interface{}) *Response {
	r.Body = body
	return r
}
