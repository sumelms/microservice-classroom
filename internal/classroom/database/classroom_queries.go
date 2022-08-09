package database

const (
	createClassroom = "create classroom"
	deleteClassroom = "delete classroom by uuid"
	getClassroom    = "get classroom by uuid"
	listClassroom   = "list classrooms"
	updateClassroom = "update classroom by uuid"
	addLesson       = "adds lesson to classroom"
	removeLesson    = "remove lesson from classroom"
)

func queriesClassroom() map[string]string {
	return map[string]string{
		createClassroom: "INSERT INTO classrooms (name, description) VALUES ($1, $2) RETURNING *",
		deleteClassroom: "UPDATE classrooms SET deleted_at = NOW() WHERE uuid = $1",
		getClassroom:    "SELECT * FROM classrooms WHERE uuid = $1",
		listClassroom:   "SELECT * FROM classrooms",
		updateClassroom: "UPDATE classrooms SET name = $1, description = $2 WHERE uuid = $3 RETURNING *",
		addLesson:       "INSERT INTO classroom_lessons (classroom_id, lesson_id) VALUES($1, $2)",
		removeLesson:    "UPDATE classroom_lessons SET deleted_at = NOW() WHERE classroom_id = $1 AND lesson_id = $2",
	}
}
