package samples

import "github.com/aacfactory/json"

type Settings struct {
	Key   string          `json:"key,omitempty"`
	Value json.RawMessage `json:"value,omitempty"`
}
