package handler

type LoginRequest struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
