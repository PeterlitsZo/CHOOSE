package service

// `form:"password" binding:"required"`
type Answer struct {
	Id         int64  `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	UserName   string `gorm:"user_name" json:"user_name" form:"user_name"`
	QuestionId string `gorm:"question_id" json:"question_id" form:"question_id"`
	Answer     string `gorm:"answer" json:"answer" form:"answer"`
}

type Question struct {
	Id       string   `gorm:"id" json:"id" form:"id"`
	Context  string   `gorm:"context" json:"context" form:"context"`
	Question string   `gorm:"question" json:"question" form:"question"`
	Choices  []string `gorm:"choices" json:"choices" form:"choices"`
}
