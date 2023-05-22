package models

type Task struct {
	Id              int64
	TaskName        string
	TaskDescription string
	StartDate       string
	EndDate         string
	Status          string
	AssignedTo      string
	AssignedToId    int64
	ProjectId       int64
	ProjectName     string
}

func AddTask(name string, description string, startDate string, endDate string, assignedTo int64, project int64) bool {
	_, err := Db.Exec("INSERT INTO tasks(task_name, task_description, start_date, end_date, status, assigned_to, project_id) VALUES(?,?,?,?,'pending',?,?)", name, description, startDate, endDate, assignedTo, project)
	return err == nil
}

func GetTasks(id int64) ([]Task, error) {
	var task []Task
	query, err := Db.Query("SELECT t.task_id, t.task_name, t.task_description, t.start_date, t.end_date, t.status, s.stakeholder_name, s.stakeholder_id, p.project_id, p.project_name FROM tasks t JOIN projects p ON t.project_id = p.project_id JOIN stakeholders s ON t.assigned_to = s.stakeholder_id WHERE p.project_id = ?", id)
	if err != nil {
		return nil, err
	}

	err = query.Scan(&task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func GetTask(id int64) (Task, error) {
	var task Task
	err := Db.QueryRow("SELECT t.task_id, t.task_name, t.task_description, t.start_date, t.end_date, t.status, s.stakeholder_name, s.stakeholder_id, p.project_id, p.project_name FROM tasks t JOIN projects p ON t.project_id = p.project_id JOIN stakeholders s ON t.assigned_to = s.stakeholder_id WHERE t.task_id = ?", id).Scan(&task)
	if err != nil {
		return task, err
	}

	return task, nil
}

func UpdateStatus(id int64, status string) bool {
	_, err := Db.Exec("UPDATE tasks SET status = ? WHERE task_id = ?", status, id)
	return err == nil
}
