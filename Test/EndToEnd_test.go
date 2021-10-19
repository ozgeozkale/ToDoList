package test

import (
	"fmt"
	"quinAI/Config"
	"quinAI/Models"
	"quinAI/Service"
	"testing"
)

// func TestUpsert(t *testing.T) {
// 	Config.InitDB()
// 	Config.DB_CreateTable()

// 	//	low priorty
// 	input1 := Models.InputModel{Title: "diş doktoru randevusu", Description: "dolgu yaptırılacak", Category: "Sağlık", Progress: "In progress", Deadline: time.Date(2021, 10, 30, 0, 0, 0, 0, time.Local)}
// 	// overdue
// 	input2 := Models.InputModel{Title: "seramik kursu malzemeler alınacak", Category: "Eğlence", Progress: "In progress", Deadline: time.Date(2021, 10, 14, 0, 0, 0, 0, time.Local)}
// 	//acil
// 	input3 := Models.InputModel{Title: "proje tamamlanacak", Description: "hemen acil", Category: "İş", Progress: "In progress", Deadline: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local)}
// 	// update
// 	input4 := Models.InputModel{Title: "diş doktoru randevusu", Description: "genel kontrol yaptırılacak", Category: "Sağlık", Progress: "In progress", Deadline: time.Date(2021, 10, 16, 23, 30, 0, 0, time.Local)}

// 	Service.Upsert(input1)
// 	Service.Upsert(input2)
// 	Service.Upsert(input3)
// 	Service.Upsert(input4)
// 	Service.GetList()
// }

func TestGetTaskByID(t *testing.T) {
	Config.InitDB()
	Config.DB_CreateTable()
	var id = Models.IdModel{Id: 5}
	_, result := Service.GetTaskByID(id)
	fmt.Println(result)
}

// func TestGetList(t *testing.T) {
// 	Config.InitDB()
// 	Config.DB_CreateTable()
// 	_, result := Service.GetList()
// 	fmt.Println(result)
// }
