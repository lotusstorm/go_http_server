package src

// type Book struct {
// 	Id     int    `json:"id" gorm:"primaryKey"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// 	Genre  string `json:"genre"`
// }

type Quiz struct {
	ID       int      `json:"id" gorm:"primaryKey"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers" gorm:"foreignKey:QuizID"`
	// SelectedAnswer []SelectedAnswer `json:"answer" gorm:"foreignKey:AnswerID"`
	MultySelect bool `json:"multySelect"`
}

type Answer struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	QuizID int
}

type SelectedAnswer struct {
	ID       int
	AnswerID int
}

type CustomResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}

var responseCodes = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	409: "Conflict",
	422: "Validation Error",
	429: "Too Many Requests",
	500: "Internal Server Error",
}
