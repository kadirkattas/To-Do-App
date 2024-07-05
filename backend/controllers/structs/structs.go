package structs

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type ToDo struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userid"`
	UserName    string `json:"username"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PostDate    string `json:"postdate"`
}
