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
	Skill{ID: 1, Name: "HTML", Level: 8},
	Skill{ID: 2, Name: "CSS", Level: 8},
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
