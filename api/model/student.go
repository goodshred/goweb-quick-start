package model

type Student struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	Sex  string `json:"sex"`
}

type StudentListReq struct {
	Sex  string `json:"sex" form:"sex" validate:"required"`
	Name string `json:"name" form:"name" validate:"required"`
	Age  int    `json:"age" form:"age"`
}
