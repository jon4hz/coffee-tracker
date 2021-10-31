package database

import (
	"fmt"

	"gorm.io/gorm"
)

// Project represents a project
type Project struct {
	gorm.Model
	Name    string
	Owner   int64
	Coffees int64
}

// NewProject creates a new project
func NewProject(name string, owner int64) (*Project, error) {
	p := Project{
		Name:    name,
		Owner:   owner,
		Coffees: 0,
	}
	var projects []Project
	r := db.Where("name = ? AND owner = ?", name, owner).Find(&projects)
	if r.Error != nil {
		return nil, fmt.Errorf("error while searching for project: %s", r.Error)
	}
	if r.RowsAffected == 0 {
		r = db.Save(&p)
		if r.Error != nil {
			return nil, fmt.Errorf("error while creating project: %s", r.Error)
		}
		return &p, nil
	}
	return &projects[0], nil
}

// GetCoffees returns the number of coffees of the project, this method doesn't search in the database, only inmemory
func (p *Project) GetCoffees() int64 {
	return p.Coffees
}

// AddCoffee adds a coffee to the project
func (p *Project) AddCoffee() error {
	p.Coffees++
	r := db.Save(p)
	if r.Error != nil {
		return fmt.Errorf("error while saving project: %s", r.Error)
	}
	return nil
}
