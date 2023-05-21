package models

type StakeHolder struct {
	ID          string
	Name        string
	Role        string
	ProjectName string
	ProjectId   int64
}

func GetStakeholders() ([]StakeHolder, error) {
	var stakeholders []StakeHolder
	query, err := Db.Query("")
	if err != nil {
		return nil, err
	}

	err = query.Scan(&stakeholders)
	if err != nil {
		return nil, err
	}

	return stakeholders, nil
}

func GetStakeholder(id int64) (StakeHolder, error) {
	var currentStakeHolder StakeHolder

	err := Db.QueryRow("").Scan(&currentStakeHolder)
	if err != nil {
		return currentStakeHolder, err
	}

	return currentStakeHolder, nil

}

func AddStakeholder(name string, email string, phone int64, password string, role string, project int64) bool {
	_, err := Db.Exec("", name, role, project)
	return err == nil
}

func AssignStakeHolder(project int64, stakeholder int64, role int64) bool {
	_, err := Db.Exec("", stakeholder, role, project)
	return err == nil
}
