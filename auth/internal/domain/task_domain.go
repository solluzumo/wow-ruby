package domain

type HashTaskDomain struct {
	Password string
	Result   chan string
}
