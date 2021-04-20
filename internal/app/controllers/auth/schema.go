package auth


type Params struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshParams struct {
	Token string `json:"token" binding:"required"`
}