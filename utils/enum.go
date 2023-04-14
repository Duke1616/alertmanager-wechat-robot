package utils

func ParamType(t string) map[string]string {
	return map[string]string{"type": t}
}

type EnumDescribe struct {
	Name     string            `json:"name"`
	Value    string            `json:"value"`
	Describe string            `json:"describe"`
	Meta     map[string]string `json:"meta"`
	SubItems []*EnumDescribe   `json:"sub_items"`
}
