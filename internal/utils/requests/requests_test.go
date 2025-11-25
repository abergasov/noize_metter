package requests_test

import (
	"github.com/google/uuid"
)

type TestResponse struct {
	ID1 uuid.UUID `json:"id"`
	ID2 uuid.UUID `json:"id_2"`
	ID3 uuid.UUID `json:"id_3"`
}

type TestRequest struct {
	ID1 uuid.UUID `json:"id"`
	ID2 uuid.UUID `json:"id_2"`
	ID3 uuid.UUID `json:"id_3"`
}
