package aws

type PrincipalName struct {
	UserName  string `json:"user_name,omitempty"`
	GroupName string `json:"group_name,omitempty"`
}