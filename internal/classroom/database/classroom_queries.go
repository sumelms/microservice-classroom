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
		createClassroom: `INSERT INTO classrooms (code, name, description, format, can_subscribe, 
					starts_at, ends_at, subject_id, course_id) 
			VALUES (:code, :name, :description, :format, :can_subscribe, 
			        :starts_at, :ends_at, :subject_id, :course_id) 
			RETURNING *`,
		deleteClassroom: "UPDATE classrooms SET deleted_at = NOW() WHERE uuid = :uuid",
		getClassroom:    "SELECT * FROM classrooms WHERE uuid = :uuid",
		listClassroom:   "SELECT * FROM classrooms",
		updateClassroom: `UPDATE classrooms 
			SET code = :code, name = :name, description = :description, format = :format, can_subscribe = :can_subscribe,
			    starts_at = :starts_at, ends_at = :ends_at, subject_id = :subject_id, course_id = :course_id
			WHERE uuid = :uuid 
			RETURNING *`,
		addLesson:    "INSERT INTO classroom_lessons (classroom_id, lesson_id) VALUES(:classroom_id, :lesson_id)",
		removeLesson: "UPDATE classroom_lessons SET deleted_at = NOW() WHERE uuid = :uuid",
	}
}
