package domain

type Repository interface {
	Create(*Lesson) (Lesson, error)
	Find(string) (Lesson, error)
	Update(*Lesson) (Lesson, error)
	Delete(string) error
	List(map[string]interface{}) ([]Lesson, error)
}
