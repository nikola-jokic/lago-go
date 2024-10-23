package lago

import (
	"encoding/json"
	"fmt"
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

type HTTPError struct {
	HTTPStatusCode int    `json:"status"`
	Message        string `json:"error"`
	ErrorCode      string `json:"code"`

	ErrorDetail ErrorDetail `json:"error_details,omitempty"`
}

func (e HTTPError) Error() string {
	msg, _ := json.Marshal(&e)
	return string(msg)
}

func readError(r io.Reader) error {
	var e HTTPError
	if err := json.NewDecoder(r).Decode(&e); err != nil {
		return fmt.Errorf("failed to decode error response: %w", err)
	}
	return &e
}
