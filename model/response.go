package model

import "encoding/json"

type ErrorResponse struct {
	Info *ErrorInfo `json:"error,omitempty"`
}

type ErrorInfo struct {
	RootCause []*ErrorInfo
	Type      string
	Reason    string
	Phase     string
}

type IndexResponse struct {
	Index   string `json:"_index"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Result  string
}

type SearchResponse struct {
	Took int64
	Hits struct {
		Total struct {
			Value int64
		}
		Hits []*SearchHit
	}
}

type SearchHit struct {
	Score   float64 `json:"_score"`
	Index   string  `json:"_index"`
	Type    string  `json:"_type"`
	Version int64   `json:"_version,omitempty"`

	Source User `json:"_source"`
}


type User struct {
	UserID    int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	IsActive  bool   `json:"isActive"`
	Balance   int    `json:"balance"`
	Phone     string `json:"phone"`
}

func (user User) ToString() string {
	res, err := json.MarshalIndent(user, "", "")
	if err != nil {
		return "Data Is empty"
	}
	return string(res)
}