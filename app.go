package main

import (
	"strings"
	"taskmanager/storage"
	"taskmanager/tasks"
	"taskmanager/user"
)

type App struct {
	Tasks       *TaskContainer
	CurrentUser string
}

type TaskContainer struct {
	LastId    int          `json:"last_id"`
	TasksList []tasks.Task `json:"tasks"`
}

func NewApp() *App {
	return &App{Tasks: &TaskContainer{}}
}

func (app *App) InitApp(username string) error {
	err := user.InitUserDir()
	if err != nil {
		return err
	}

	app.CurrentUser = username
	if !user.UserExist(username) {
		err = user.CreateUser(username)
		if err != nil {
			return err
		}
	} else {
		err = storage.ReadFromFile(user.GetUserTasks(username), app.Tasks)
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *App) CreateTask(taskDescription string) error {
	app.Tasks.LastId += 1
	app.Tasks.TasksList = tasks.CreateTask(taskDescription, app.Tasks.LastId, app.Tasks.TasksList)
	err := storage.SaveToFile(user.GetUserTasks(app.CurrentUser), app.Tasks)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) StatusToCode(status string) (int, error) {
	switch strings.ToLower(status) {
	case "done":
		return tasks.DONE, nil
	case "undone":
		return tasks.UNDONE, nil
	case "in-progress":
		return tasks.INPROGRESS, nil
	}
	return 0, tasks.ErrUnknownStatusCode
}

func (app *App) GetAllTasks() []tasks.Task {
	return app.Tasks.TasksList
}

func (app *App) GetTaskByStatus(status int) []tasks.Task {
	return tasks.GetFilteredTasks(app.Tasks.TasksList, status)
}

func (app *App) UpdateDescriptionTask(taskDescription string, id int) error {
	TasksList, err := tasks.ChangeTaskDescription(taskDescription, id, app.Tasks.TasksList)
	if err != nil {
		return err
	}
	app.Tasks.TasksList = TasksList
	err = storage.SaveToFile(user.GetUserTasks(app.CurrentUser), app.Tasks)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) UpdateTaskStatus(id int, status int) error {
	TasksList, err := tasks.ChangeStatusTask(id, app.Tasks.TasksList, status)
	if err != nil {
		return err
	}
	app.Tasks.TasksList = TasksList
	err = storage.SaveToFile(user.GetUserTasks(app.CurrentUser), app.Tasks)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) DeleteTask(id int) error {
	TasksList, err := tasks.DeleteTask(id, app.Tasks.TasksList)
	if err != nil {
		return err
	}
	app.Tasks.TasksList = TasksList
	err = storage.SaveToFile(user.GetUserTasks(app.CurrentUser), app.Tasks)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) ExitApp() error {
	if app.CurrentUser != "" {
		err := storage.SaveToFile(user.GetUserTasks(app.CurrentUser), app.Tasks)
		if err != nil {
			return err
		}
	}
	return nil
}
