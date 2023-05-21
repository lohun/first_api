package models

type Project struct {
	Id                 int64
	ProjectName        string
	ProjectDescription string
	StartDate          string
	EndDate            string
	Status             string
}

func AddProject(name string, description string, startDate string, endDate string) bool {
	_, err := Db.Exec("", name, description, startDate, endDate)
	return err == nil
}

func GetProjects() ([]Project, error) {
	var project []Project
	query, err := Db.Query("")
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
	err := Db.QueryRow("", id).Scan(&project)
	if err != nil {
		return project, err
	}

	return project, nil
}

func UpdateProjectStatus(id int64, status string) bool {
	_, err := Db.Exec("", status, id)
	return err == nil
}
