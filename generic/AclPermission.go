package generic

type AclPermission string

const (
	AclPermissionRead   = "READ"
	AclPermissionWrite  = "WRITE"
	AclPermissionManage = "MANAGE"
)
