package Config

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Priority: low-normal-critical - overdue.
// status: in progress - completed - overdue.

func DB_CreateTable() {
	statement, _ := DB.Prepare("CREATE TABLE IF NOT EXISTS ToDotable (ID INTEGER PRIMARY KEY, Title TEXT UNIQUE, Description TEXT, Category TEXT, Progress TEXT, Deadline datetime, Priority TEXT, CreatedTime datetime default current_timestamp, UpdatedTime datetime default current_timestamp)")
	statement.Exec()

	createdTime := time.Now()
	statement, _ = DB.Prepare("INSERT INTO ToDotable (Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime) VALUES (?, ?, ? ,?, ?, ?, ?, ?)")
	deadline := time.Date(2021, 10, 15, 12, 0, 0, 0, time.Local)
	statement.Exec("Marketten 1 kg. elma alınacak", "Bir sonraki alışverişte", "Alışveriş", "In progress", deadline, "Low.", createdTime, createdTime)

	statement, _ = DB.Prepare("INSERT INTO ToDotable (Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime) VALUES (?, ?, ? ,?, ?, ?, ?, ?)")
	deadline = time.Date(2021, 10, 17, 12, 0, 0, 0, time.Local)
	statement.Exec("Quin AI Case: ToDo App yazılacak", "Deadline içinde tamamlanmalı.", "Job interviews", "Completed.", deadline, "Completed.", createdTime, createdTime)

	statement, _ = DB.Prepare("INSERT INTO ToDotable (Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime) VALUES (?, ?, ? ,?, ?, ?, ?, ?)")
	deadline = time.Date(2021, 12, 3, 12, 0, 0, 0, time.Local)
	statement.Exec("ALES başvurusu", "Başvuru tarihi geçmeden önce, acil!.", "Yüksek Lisans", "Overdue.", deadline, "Overdue.", createdTime, createdTime)

	statement, _ = DB.Prepare("INSERT INTO ToDotable (Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime) VALUES (?, ?, ? ,?, ?, ?, ?, ?)")
	deadline = time.Date(2021, 10, 19, 12, 0, 0, 0, time.Local)
	statement.Exec("Aybüke'ye doğum günü hediyesi alınacak", "Acilll", "Doğum günleri", "In progress", deadline, "Critical.", createdTime, createdTime)

	rows, _ := DB.Query("SELECT id, Title, Description, Category, Progress, Deadline, Priority, CreatedTime, UpdatedTime FROM ToDotable")
	var id int
	var Title string
	var Description string
	var Category string
	var Progress string
	var Deadline *time.Time
	var Priority string
	var CreatedTime *time.Time
	var UpdatedTime *time.Time

	for rows.Next() {
		rows.Scan(&id, &Title, &Description, &Category, &Progress, &Deadline, &Priority, &CreatedTime, &UpdatedTime)
	}
}
