package Service

import (
	"ToDoProject/Config"
	"ToDoProject/Models"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var id int
var Title string
var Description string
var Category string
var Progress string
var Deadline *time.Time
var Priority string
var CreatedTime *time.Time
var UpdatedTime *time.Time

// Insert a new task to database.
func Insert(input Models.InputModel) (error, []Models.TaskModel) {
	createdTime := time.Now()
	updatedTime := createdTime

	task := DecidePriorty(input)

	statement, _ := Config.DB.Prepare("INSERT INTO ToDotable (Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime) VALUES (?, ?, ? ,?, ?, ?, ?, ?)")
	statement.Exec(task.Title, task.Description, task.Category, task.Progress, task.Deadline, task.Priority, createdTime, updatedTime)

	rows, _ := Config.DB.Query("SELECT id, Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime FROM ToDotable")
	for rows.Next() {
		rows.Scan(&id, &Title, &Description, &Category, &Progress, &Deadline, &Priority, &CreatedTime, &UpdatedTime)
	}
	_, list := GetList()
	return nil, list
}

// Update an existing task in database.
func Update(input Models.InputModel) (error, []Models.TaskModel) {

	var updatedTime time.Time
	updatedTime = time.Now()

	task := DecidePriorty(input)

	statement, _ := Config.DB.Prepare("UPDATE ToDotable SET Description = ?, Category = ?, Progress = ?, Deadline = ?, Priority = ?, UpdatedTime = ? WHERE Title = ? RETURNING id, Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime")
	statement.Exec(task.Description, task.Category, task.Progress, task.Deadline, task.Priority, updatedTime, task.Title)

	rows, _ := Config.DB.Query("SELECT id, Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime FROM ToDotable")
	for rows.Next() {
		rows.Scan(&id, &Title, &Description, &Category, &Progress, &Deadline, &Priority, &CreatedTime, &UpdatedTime)
	}
	_, list := GetList()
	return nil, list
}

// Decide Priorty level of the task.
func DecidePriorty(input Models.InputModel) (task Models.TaskModel) {
	task.Title = input.Title
	task.Description = input.Description
	task.Category = input.Category
	task.Progress = input.Progress
	task.Deadline = input.Deadline
	HoursLeft := input.Deadline.Sub(time.Now()).Hours()

	if input.Progress == "Completed" {
		task.Priority = "Completed"
	} else if input.Progress == "Overdue" {
		task.Priority = "Overdue"
	} else if HoursLeft < 0 {
		task.Priority = "Overdue"
		task.Progress = task.Priority
	} else if input.Progress == "In progress" {
		if strings.Contains(input.Description, "acil") || strings.Contains(input.Description, "Acil") {
			task.Priority = "Critical"
			return task
		}

		if input.Priority == "" {
			if HoursLeft >= 0 && HoursLeft <= 4 {
				task.Priority = "Critical"
			} else if HoursLeft > 4 && HoursLeft <= 24 {
				task.Priority = "Important"
			} else if HoursLeft > 24 && HoursLeft <= 24*3 {
				task.Priority = "Normal"
			} else {
				task.Priority = "Low"
			}
		} else {
			task.Priority = input.Priority
		}
	}
	return task
}

// Get all tasks from the database.
func GetList() (err error, list []Models.TaskModel) {

	rows, _ := Config.DB.Query("SELECT id, Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime FROM ToDotable")

	for rows.Next() {
		rows.Scan(&id, &Title, &Description, &Category, &Progress, &Deadline, &Priority, &CreatedTime, &UpdatedTime)
		list = append(list, Models.TaskModel{Id: id, Title: Title, Description: Description, Category: Category, Progress: Progress, Deadline: *Deadline, Priority: Priority, CreatedTime: *CreatedTime, UpdatedTime: *UpdatedTime})
	}
	return nil, list
}

// Get the task with given id from database.
func GetTaskByID(input Models.IdModel) (err error, Result Models.TaskModel) {
	rows, _ := Config.DB.Query("SELECT * FROM ToDotable WHERE id =  ?", input.Id)
	for rows.Next() {
		rows.Scan(&id, &Title, &Description, &Category, &Progress, &Deadline, &Priority, &CreatedTime, &UpdatedTime)
		Result = Models.TaskModel{Id: id, Title: Title, Description: Description, Category: Category, Progress: Progress, Deadline: *Deadline, Priority: Priority, CreatedTime: *CreatedTime, UpdatedTime: *UpdatedTime}
	}
	return nil, Result
}

// For control purposes.
func Get() (error, string) {
	msg := "Hello world"
	return nil, msg
}
