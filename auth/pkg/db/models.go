package db

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateBaseUserRequest struct {
	Email    string
	Password string
	Name     string
	Role     string
}

type CreatePublisherRequest struct {
	CreateBaseUserRequest
	PlanId int
}

type GenerateTokenRequest struct {
	Id   int
	Role string
}

type JWT_claim struct {
	Id       int
	Password string
	Role     string
}