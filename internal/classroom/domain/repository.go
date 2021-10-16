package domain

type Repository interface {
	Create(*Classroom) (Classroom, error)
	Find(string) (Classroom, error)
	Update(*Classroom) (Classroom, error)
	Delete(string) error
	List() ([]Classroom, error)
}
