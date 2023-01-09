package model

import (
	"time"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/ryo-funaba/example_echo/internal/utils/timeutil"
)

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func EmployeeFromDomainModel(m *model.Employee) *Employee {
	return &Employee{
		ID:        m.EmpNo,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Age:       calcAge(m.BirthDate),
	}
}

// calcAge calculates age from birthday.
func calcAge(birthday time.Time) int {
	now := timeutil.GetTruncatedNowJSTTime()

	thisYear, thisMonth, thisDay := now.Date()
	age := thisYear - birthday.Year()

	// when birthday is yet to come
	if thisMonth < birthday.Month() || thisMonth == birthday.Month() && thisDay < birthday.Day() {
		age--
	}

	return age
}
