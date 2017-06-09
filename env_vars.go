package cf_connectors

import "encoding/json"

type VcapServices map[string][]*json.RawMessage

type MinimalVcapService struct {
	Tags []string `json:"tags"`
}
