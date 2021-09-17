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
	{ID: 3, Name: "JavaScript", Level: 9},
	{ID: 9, Name: "React", Level: 8},
	{ID: 4, Name: "Java", Level: 7},
	{ID: 5, Name: "C#", Level: 6},
	{ID: 6, Name: "Go", Level: 3},
	{ID: 7, Name: "HTML", Level: 9},
	{ID: 8, Name: "CSS", Level: 8},
	{ID: 12, Name: "AWS", Level: 7},
	{ID: 11, Name: "Kubernetes", Level: 8},
	{ID: 10, Name: "Docker", Level: 8},
	{ID: 1, Name: "Natural", Level: 10},
	{ID: 2, Name: "NaturalONE", Level: 10},
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
