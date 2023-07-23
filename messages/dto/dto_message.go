package dto

type MessageDto struct {
	Id        int    `json:"id"`
	Userid    int    `json:"userid"`
	Itemid    string `json:"itemid"`
	Content   string `json:"content"`
	Createdat string `json:"createdat"`
}

type MessagesDto []MessageDto
