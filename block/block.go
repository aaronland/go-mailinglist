package block

const RULE_TYPE_ADDRESS int = 1
const RULE_TYPE_HOST int = 2
const RULE_TYPE_IP int = 3
const RULE_TYPE_CIDR int = 4

type Rule struct {
	Type    int    `json:"type"`
	Value   string `json:"value"`
	Created int64  `json:"created"`
}
