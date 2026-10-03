package tasks

import "time"

type Task struct {
	Id          int
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

const (
	UNDONE = iota
	DONE
	INPROGRESS
)

func CreateTask(description string, lastId int, taskList []Task) []Task {
	task := Task{Id: lastId,
		Description: description,
		Status:      UNDONE,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil}
	taskList = append(taskList, task)
	return taskList
}

func ChangeTaskDescription(description string, id int, taskList []Task) ([]Task, error) {
	for count := range taskList {
		if taskList[count].Id == id {
			taskList[count].Description = description

			timeNow := time.Now()
			taskList[count].UpdatedAt = &timeNow
			return taskList, nil
		}
	}
	return taskList, ErrTaskNotFound
}

func DeleteTask(id int, taskList []Task) ([]Task, error) {
	for count := range taskList {
		if taskList[count].Id == id {
			taskList = append(taskList[:count], taskList[count+1:]...)
			return taskList, nil
		}
	}
	return nil, ErrTaskNotFound
}

func ChangeStatusTask(id int, taskList []Task, status int) ([]Task, error) {
	for count := range taskList {
		if taskList[count].Id == id {
			taskList[count].Status = status
			return taskList, nil
		}
	}
	return taskList, ErrTaskNotFound
}

func GetFilteredTasks(taskList []Task, status int) []Task {
	var filteredTasks []Task
	for count := range taskList {
		if taskList[count].Status == status {
			filteredTasks = append(filteredTasks, taskList[count])
		}
	}
	return filteredTasks
}
