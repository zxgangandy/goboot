package model

type GetUserReq struct {
	UserID int64 `json:"userId"`
}

type GetStudentReq struct {
	StudentID int64 `json:"studentId"`
}
