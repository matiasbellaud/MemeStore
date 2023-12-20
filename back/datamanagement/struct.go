package datamanagement

type MemeFromDB struct {
	IdMeme      int
	IdUser      int
	MemeName    string
	UrlMeme     string
	Description string
	Topic       string
	Favorite    int
}

type UserFromDB struct {
	IdUser           int
	Description      string
	Username         string
	Password         string
	VerificationWord string
}
