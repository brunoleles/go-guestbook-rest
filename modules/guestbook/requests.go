package guestbook

type guestbookPostRequest struct {
	//models.GuestbookModel
	Name    string `json:"name" form:"name"`
	Message string `json:"message" form:"message"`
}

type guestbookDeleteRequest struct {
	ID string `json:"id" form:"id"`
}
