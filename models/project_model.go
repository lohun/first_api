package models

type Project struct {
	Id                 int64  `json:"id"`
	ProjectName        string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
	Status             string `json:"status"`
}

func AddProject(name string, description string, startDate string, endDate string) bool {
	_, err := Db.Exec("INSERT INTO projects (project_name, project_description, start_date, end_date, status) VALUES (?,?,?,?, 'pending')", name, description, startDate, endDate)
	return err == nil
}

func GetProjects() ([]Project, error) {
	var project []Project
	query, err := Db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}

	err = query.Scan(&project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func GetProject(id int64) (Project, error) {
	var project Project
	err := Db.QueryRow("SELECT * FROM projects WHERE project_id = ?", id).Scan(&project)
	if err != nil {
		return project, err
	}

	return project, nil
}

func UpdateProjectStatus(id int64, status string) bool {
	_, err := Db.Exec("UPDATE projects SET status = ? WHERE project_id = ?", status, id)
	return err == nil
}
