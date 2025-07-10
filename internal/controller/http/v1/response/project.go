package response

import "task-trail/internal/usecase/dto"

type ProjectRes struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TaskCount   int    `json:"tasksCount"`
}

type projectCreateRes struct {
	ID int `json:"id"`
}

func NewProjectResFromDTO(data *dto.ProjectRes) *ProjectRes {
	return &ProjectRes{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		TaskCount:   data.TaskCount,
	}
}

func NewProjectResFromDTOBatch(data []*dto.ProjectRes) []*ProjectRes {
	if len(data) == 0 {
		return []*ProjectRes{}
	}
	var retVal []*ProjectRes
	for _, v := range data {
		retVal = append(retVal, NewProjectResFromDTO(v))
	}
	return retVal
}

func NewProjectCreateResFromDTO(projectID int) *projectCreateRes {
	return &projectCreateRes{ID: projectID}
}
