package database

import "github.com/jmoiron/sqlx"

type TaskStore struct {
	db *sqlx.DB
}

func NewTaskStore(db *sqlx.DB) *TaskStore
 return &TaskStore{db: db}

func (s *TaskStore) GetAll() ([]models.Task, error){
	tasks:=[]models.Task

     query :=
`SELECT id, title, description, completed, created_at, update_at FROM tasks order by created_at desc;`

   err := s.db.Select(&tasks, query)
   if err != nil {
	return nil, err
   }

   return tasks, nil
}

func (s *TaskStore) GetByID(id int) (models.Task, error) {
	var task models.Task

	query :=
`SELECT id, title, description, completed, created_at, update_at FROM tasks where id = $1;`

err := s.db.Get(&task, query, id)

if err == sql.ErrNoRows {
	return nil, fmt.Errorf("task with id %d not found', id")
}

if err != nil{
	return nil, err
}

return &task, nil

}

func (s *TaskStore) Create(input models.CreatedtaskInput)(*models.Task,error){
	query := 
`INSERT INTO tasks (title, description, completed, created_at, update_at)
VALUES ($1, $2, $3, $4, $5)
returning id, title, description, completed, created_at, update_at;`

now := time.Now()
var task models.Task()

err := s.db.QueryRowx(query, input.Title, input.Description, input.Completed, now, now,).StructScan(&task)

if err != nil{
	return nil, err
}
return &task, nil
}

func (s *TaskStore) Update(id int, input models.UpdateTaskInput)(*models.Task, error){
	task, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if input.Tittle != nil{
		task.Tittle = *input.Tittle
	}
	if input.Description != nil{
		task.Description = *input.Description
	}
	if input.Completed != nil{
		task.Completed = *input.Completed
	}

	task.Update = time.Now()

	query :=
	`UPDATE tasks
	SET title = $1, description = $2, completed = $3, update = $4
	where id = $5
	returning id, title, description, completed, created_at, update_at;`

	var updateTask models.Task
	
	err = s.db.QueryRowx(query, task.Title, task.Description, task.Completed, task.UpdateAt, id).StructScan(&updateTask)
	if err != nil{
		return nil, err
	}
	return &updateTask, nil
}

func (s *TaskStore) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = $1;`

	result, err := s.db.Exec(query, id)

	if err != nil{
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil{
		return err
	}
	if rows == 0{
		return fmt.Errorf("task with id %d not found", id)
	}
	return nil
}