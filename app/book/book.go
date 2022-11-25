package book

type Book struct {
	Id       string   `json:"id" firestore:"-"`
	Title    string   `json:"title" firestore:"title" binding:"required,gt=0"`
	Authors  []string `json:"authors" firestore:"authors" binding:"required,gt=0"`
	Year     int      `json:"year" firestore:"year" binding:"required,gt=0"`
	Editions int      `json:"editions" firestore:"editions" binding:"required,gt=0"`
	Category string   `json:"category" firestore:"category" binding:"required,gt=0"`
}

type Library struct {
	Books []Book `json:"books"`
}
