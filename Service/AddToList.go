package Service

import (
	"fmt"
	"quinAI/Config"
	"quinAI/Models"
	"strconv"
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

func DecidePriorty(input Models.InputModel) (task Models.TaskModel) {
	task.Title = input.Title
	task.Description = input.Description
	task.Category = input.Category
	task.Progress = input.Progress
	task.Deadline = input.Deadline
	if input.Progress == "Completed" {
		task.Priority = "Completed"
	} else if input.Progress == "Overdue" {
		task.Priority = "Overdue"
	} else if strings.Contains(input.Description, "acil") || strings.Contains(input.Description, "Acil") {
		task.Priority = "Critical"
	} else {
		hoursLeft := input.Deadline.Sub(time.Now()).Hours()
		if hoursLeft < 0 {
			task.Priority = "Overdue"
			task.Progress = task.Priority
		} else if hoursLeft > 0 && hoursLeft <= 4 {
			task.Priority = "Critical"
		} else if hoursLeft > 4 && hoursLeft <= 24 {
			task.Priority = "Important"
		} else if hoursLeft > 24 && hoursLeft <= 24*3 {
			task.Priority = "Normal"
		} else {
			task.Priority = "Low"
		}
	}
	return task
}

func GetList() (err error, list []Models.TaskModel) {

	rows, _ := Config.DB.Query("SELECT id, Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime FROM ToDotable")

	for rows.Next() {
		rows.Scan(&id, &Title, &Description, &Category, &Progress, &Deadline, &Priority, &CreatedTime, &UpdatedTime)
		fmt.Println(strconv.Itoa(id) + ":" + " Title:" + Title + " Description:" + Description + " Category:" + Category + " Progress:" + Progress + " Deadline:" + Deadline.String() + " Priority:" + Priority + " createdTime:" + CreatedTime.String() + " updatedTime:" + UpdatedTime.String())
		list = append(list, Models.TaskModel{Id: id, Title: Title, Description: Description, Category: Category, Progress: Progress, Deadline: *Deadline, Priority: Priority, CreatedTime: *CreatedTime, UpdatedTime: *UpdatedTime})
	}
	fmt.Println(list)
	return nil, list
}

func Get() (error, string) {
	msg := "Hello world"
	return nil, msg
}

func GetTaskByID(input Models.IdModel) (err error, Result Models.TaskModel) {
	rows, _ := Config.DB.Query("SELECT * FROM ToDotable WHERE id =  ?", input.Id)
	for rows.Next() {
		rows.Scan(&id, &Title, &Description, &Category, &Progress, &Deadline, &Priority, &CreatedTime, &UpdatedTime)
		Result = Models.TaskModel{Id: id, Title: Title, Description: Description, Category: Category, Progress: Progress, Deadline: *Deadline, Priority: Priority, CreatedTime: *CreatedTime, UpdatedTime: *UpdatedTime}
	}
	return nil, Result
}
