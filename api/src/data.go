package src

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

var booksDB = []Book{
	{Id: 123, Title: "The Hobbit", Author: "J. R. R. Tolkien", Genre: "Fantasy"},
	{Id: 456, Title: "Harry Potter and the Philosopher's Stone", Author: "J. K. Rowling", Genre: "Fantasy"},
	{Id: 789, Title: "The Little Prince", Author: "Antoine de Saint-Exupéry", Genre: "Novella"},
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
