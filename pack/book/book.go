package book

// หากขึ้นต้นด้วยตัวใหญ่จะเป็นการ public และตัวเล็กเป็น private
// มีผลถึงชื่อทั้งหมด เช่น func var struct field interface
type Book struct{}

func New() Book {
	return Book{}
}
