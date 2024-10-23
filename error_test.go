package lago

import (
	"encoding/json"
	"testing"
)

func TestErrorErr(t *testing.T) {
	var hasErr error = HTTPError{
		HTTPStatusCode: 422,
		Message:        "Type assertion failed",
	}
	t.Logf("%s", hasErr.Error())
}

func TestErrorNoErr(t *testing.T) {
	var noErr error = HTTPError{
		HTTPStatusCode: 500,
		Message:        "500",
	}
	t.Logf("%s", noErr.Error())
}

func TestErrorDetails(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{
			name: "Single detail",
			input: `{
  "status": 422,
  "error": "Unprocessable Entity",
  "code": "validation_errors",
  "error_details": {
    "transaction_id": [
      "value_already_exist"
    ]
  }
}`,
			want: &HTTPError{
				HTTPStatusCode: 422,
				Message:        "Unprocessable Entity",
				ErrorCode:      "validation_errors",
				ErrorDetail: map[int]map[string][]string{
					0: {
						"transaction_id": {
							"value_already_exist",
						},
					},
				},
			},
		},
		{
			name: "Multiple details",
			input: `{
  "status": 422,
  "error": "Unprocessable Entity",
  "code": "validation_errors",
  "error_details": {
    "0": {
      "transaction_id": [
        "value_already_exist"
      ]
    },
    "1": {
      "transaction_id": [
        "value_already_exist"
      ]
    },
    "2": {
      "transaction_id": [
        "value_already_exist"
      ]
    }
  }
}`,
			want: &HTTPError{
				HTTPStatusCode: 422,
				Message:        "Unprocessable Entity",
				ErrorCode:      "validation_errors",
				ErrorDetail: map[int]map[string][]string{
					0: {
						"transaction_id": {
							"value_already_exist",
						},
					},
					1: {
						"transaction_id": {
							"value_already_exist",
						},
					},
					2: {
						"transaction_id": {
							"value_already_exist",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errObj := &HTTPError{}
			err := json.Unmarshal([]byte(tt.input), errObj)
			if err != nil {
				t.Errorf("got error %s", err.Error())
				return
			}

			expectErr, err := json.Marshal(tt.want)
			if err != nil {
				t.Errorf("got error %s", err.Error())
				return
			}

			gotErr, err := json.Marshal(errObj)
			if err != nil {
				t.Errorf("got error %s", err.Error())
				return
			}

			if string(expectErr) != string(gotErr) {
				t.Errorf("got error %s, but expected error %s", string(gotErr), string(expectErr))
				return
			}
		})
	}
}
