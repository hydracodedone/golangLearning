package service

import (
	"context"
)

type StudentInfoQueryService struct {
}

func (studentInfoQueryService *StudentInfoQueryService) QueryStudentInfo(context.Context, *StudentQueryId) (*Student, error) {
	return &Student{
		Id:   1,
		Name: "Hydra",
		Sex:  Sex_FEMALE,
		StudentGrade: &Grade{
			Chinese: 50,
			Math:    100,
			English: 100,
		},
		StudentPersonalInfo: &Student_StudentPersonalInfo{
			Address: "CN_CD",
			Rich:    true,
		},
		StudentTeacher: []*Teacher{
			{
				Id: 1,
			},
			{
				Id: 2,
			},
		},
	}, nil
}
