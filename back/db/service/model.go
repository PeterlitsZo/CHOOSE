package service

// `form:"password" binding:"required"`
type Answer struct {
	Id         int64  `gorm:"id;primary_key;AUTO_INCREMENT" json:"id"`
	UserName   string `gorm:"user_name" json:"user_name"`
	QuestionId string `gorm:"question_id" json:"question_id"`
	Answer     string `gorm:"answer" json:"answer"`
}

type Question struct {
	Id       int64  `gorm:"id" json:"id"`
	Context  string `gorm:"context" json:"context"`
	Question string `gorm:"question" json:"question"`
}
