package pkg

type ErrorExisting struct {
	Text string
}

func (e *ErrorExisting) Error() string {
	return e.Text
}
