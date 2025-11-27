package entity

type Permission struct {
	ID    uint
	Title string
}

const (
	UserListPermission   = "user-list"
	UserDeletePermission = "user-delete"
)
