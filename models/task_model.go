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
	_, err := Db.Exec("", name, description, startDate, endDate, assignedTo, project)
	return err == nil
}

func GetTasks() ([]Task, error) {
	var task []Task
	query, err := Db.Query("")
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
	err := Db.QueryRow("", id).Scan(&task)
	if err != nil {
		return task, err
	}

	return task, nil
}

func UpdateStatus(id int64, status string) bool {
	_, err := Db.Exec("", status, id)
	return err == nil
}
