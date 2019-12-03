// model.go

package main

import (
	"errors"
)

// Skill - struct modelling a skill
type Skill struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// Skills - global var containing all skills
var Skills = []Skill{
	Skill{ID: 1, Name: "Natural", Level: 10},
	Skill{ID: 2, Name: "NaturalONE", Level: 10},
	Skill{ID: 3, Name: "JavaScript", Level: 8},
	Skill{ID: 4, Name: "Java", Level: 7},
	Skill{ID: 5, Name: "C#", Level: 6},
	Skill{ID: 6, Name: "Go", Level: 3},
	Skill{ID: 7, Name: "HTML", Level: 8},
	Skill{ID: 8, Name: "CSS", Level: 8},
	Skill{ID: 9, Name: "React", Level: 5},
	Skill{ID: 10, Name: "Docker", Level: 7},
	Skill{ID: 11, Name: "Kubernetes", Level: 6},
	Skill{ID: 12, Name: "AWS", Level: 5},
}

func getSkills() ([]Skill, error) {
	return Skills, nil
}

func getSkill(id int) (Skill, error) {
	for _, skill := range Skills {
		if skill.ID == id {
			return skill, nil
		}
	}

	return Skill{}, errors.New("Skill not found")
}
