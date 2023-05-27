package dao

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Age  int64  `gorm:"column:age" json:"age"`
	Sex  string `gorm:"column:sex" json:"sex"`
}

func (Student) TableName() string {
	return "student"
}

type TbaleStudent struct {
	db *gorm.DB
}

func NewTableStudent(db *gorm.DB) TbaleStudent {
	return TbaleStudent{
		db: db,
	}
}

func (t TbaleStudent) Find(stu Student) ([]Student, error) {
	var res []Student
	name := stu.Name
	stu.Name = ""
	// 全是and查询
	dbWhere := t.db.Where(&stu)
	if name != "" {
		dbWhere = dbWhere.Where("name like ?", name+"%")
	}
	err := dbWhere.Debug().Find(&res).Error
	return res, err
}

func (t TbaleStudent) Insert(student Student) error {
	return t.db.Create(&student).Error
}

func (t TbaleStudent) Delete(id int64) error {
	return t.db.Where("id=?", id).Delete(&Student{}).Error
}

func (t TbaleStudent) Update(student Student) error {
	// Update attributes with `struct`, will only update non-zero fields
	// 这里有个坑：假如status有0：运行，1：等待，那么要从1更新为0用下面方法status将不会被更新：所以对于数据库设计，不要使用0表示状态量
	// 假如真有这种情况，那就指定字段更新：return s.db.Model(&Student{}).Select("filed1", "filed2").Updates(student).Error
	return t.db.Model(&Student{}).Where("id=?", student.ID).Updates(&student).Error
}
