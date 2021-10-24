package domain

type Repository interface {
	Create(*ClassroomLesson) (ClassroomLesson, error)
	Find(string) (ClassroomLesson, error)
	Update(*ClassroomLesson) (ClassroomLesson, error)
	Delete(string) error
	List(map[string]interface{}) ([]ClassroomLesson, error)
}
