package api

type EchoReq struct {
}

type EchoResp struct {
	Message string `json:"message"`
}

type SaveReq struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
}

type SaveResp struct {
}
