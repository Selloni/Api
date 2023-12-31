package user

// bson - внутренее моонгоДБ, _id -   системное поле сомо генерирует уникальное поле, omid - может быть пустым
type User struct {
	Id           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
	Email        string `json:"email"bson:"email"`
}

type CreateUserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
