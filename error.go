package lago

import (
	"encoding/json"
	"io"
)

type ErrorDetail map[int]map[string][]string

func (ed *ErrorDetail) UnmarshalJSON(data []byte) error {
	// First attempt to unmarshal singular.
	var singularErr map[string][]string
	err := json.Unmarshal(data, &singularErr)
	if err == nil {
		*ed = ErrorDetail{0: singularErr}
		return nil
	}

	// Then attempt to unmarshal multiple.
	var multipleErr map[int]map[string][]string
	err = json.Unmarshal(data, &multipleErr)
	if err == nil {
		*ed = ErrorDetail(multipleErr)
		return nil
	}

	return err
}

func (ed ErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[int]map[string][]string(ed))
}

type Error struct {
	Err error `json:"-"`

	HTTPStatusCode int    `json:"status"`
	Message        string `json:"error"`
	ErrorCode      string `json:"code"`

	ErrorDetail ErrorDetail `json:"error_details,omitempty"`
}

func (e Error) Error() string {
	type alias struct {
		Error
		Err string `json:"err,omitempty"`
	}
	err := alias{Error: e}
	if e.Err != nil {
		err.Err = e.Err.Error()
	}
	msg, _ := json.Marshal(&err)
	return string(msg)
}

func readError(r io.Reader) *Error {
	var e Error
	if err := json.NewDecoder(r).Decode(&e); err != nil {
		return &Error{Err: err}
	}
	return &e
}
