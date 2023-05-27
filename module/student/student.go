package module

import (
	"github.com/jinzhu/copier"
	"goweb-quick-start/api/model"
	dao2 "goweb-quick-start/dao"
)

func SaveStudent(student model.Student) error {
	studentDB := dao2.NewTableStudent(dao2.GetMysqlDB())
	var stuEntity dao2.Student
	copier.Copy(&stuEntity, student)
	err := studentDB.Insert(stuEntity)
	return err
}

func UpdateStudent(student model.Student) error {
	studentDB := dao2.NewTableStudent(dao2.GetMysqlDB())
	var stuEntity dao2.Student
	copier.Copy(&stuEntity, student)
	stuEntity.ID = student.Id
	err := studentDB.Update(stuEntity)
	return err
}

func DeleteStudent(id int64) error {
	studentDB := dao2.NewTableStudent(dao2.GetMysqlDB())
	err := studentDB.Delete(id)
	return err
}

func FindStudent(stu model.StudentListReq) ([]dao2.Student, error) {
	studentDB := dao2.NewTableStudent(dao2.GetMysqlDB())
	var stuEntity dao2.Student
	copier.Copy(&stuEntity, stu)
	return studentDB.Find(stuEntity)
}
