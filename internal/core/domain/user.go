package domain

type User struct {
	Id       *uint64 `json:"id,omitempty"`
	Uuid     *string `json:"uuid,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

func NewUser() *User {
	return &User{}
}

func NewUserFill(id *uint64, uuid *string, userName *string) *User {
	return &User{Id: id, Uuid: uuid, UserName: userName}
}

func (model *User) Valid() bool {
	if model == nil {
		return false
	}

	if model.UserName != nil && model.Uuid != nil {
		return len(*model.UserName) > 0 && len(*model.Uuid) > 0
	} else {
		return true
	}
}
