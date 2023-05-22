package models

type StakeHolder struct {
	ID          string
	Name        string
	Role        string
	ProjectName string
	ProjectId   int64
}

func GetStakeholders(id int64) ([]StakeHolder, error) {
	var stakeholders []StakeHolder
	query, err := Db.Query("SELECT p.stakeholder_id, p.stakeholder_name, r.name, j.project_name, j.project_id  FROM projectstakeholders p JOIN user_roles u ON p.stakeholder_id = u.user JOIN roles r ON u.role = r.id JOIN projects j ON u.project_id = j.project_id WHERE u.project_id = ?", id)
	if err != nil {
		return nil, err
	}

	err = query.Scan(&stakeholders)
	if err != nil {
		return nil, err
	}

	return stakeholders, nil
}

func GetStakeholder(stakeHolderId int64, projectId int64) (StakeHolder, error) {
	var currentStakeHolder StakeHolder

	err := Db.QueryRow("SELECT p.stakeholder_id, p.stakeholder_name, r.name, j.project_name, j.project_id  FROM projectstakeholders p JOIN user_roles u ON p.stakeholder_id = u.user JOIN roles r ON u.role = r.id JOIN projects j ON u.project_id = j.project_id WHERE u.project_id = ? and p.stakeholder_id = ?", projectId, stakeHolderId).Scan(&currentStakeHolder)
	if err != nil {
		return currentStakeHolder, err
	}

	return currentStakeHolder, nil

}

func AddStakeholder(name string, email string, phone int64, password string, role int64, project int64) bool {
	query, err := Db.Exec("INSERT INTO projectstakeholders(stakeholder_name, email, phone_number, password, verified) VALUES(?,?,?,?,1) ON DUPLICATE KEY UPDATE stakeholder_name = ?, phone_number = ?", name, email, phone, password)
	if err == nil {
		stakeholder, _ := query.LastInsertId()
		_, err := Db.Exec("INSERT INTO user_roles(user, role, project) VALUES(?,?,?)", stakeholder, role, project)

		return err == nil
	}
	return false
}

func AssignStakeHolder(project int64, stakeholder int64, role int64) bool {
	_, err := Db.Exec("INSERT INTO user_roles(user, role, project) VALUES(?,?,?)", stakeholder, role, project)
	return err == nil
}
