package service

type Answer struct {
	Id             int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserName       string `gorm:"column:user_name" json:"user_name"`
	QuestionId     string `gorm:"column:question_id" json:"question_id"`
	Answer         string `gorm:"column:answer" json:"answer"`
	IsHardToAnswer bool   `gorm:"column:is_hard_to_answer" json:"is_hard_to_answer"`
}

type Question struct {
	Id       string   `gorm:"column:id" json:"id"`
	Context  string   `gorm:"column:context" json:"context"`
	Question string   `gorm:"column:question" json:"question"`
	Choices  []Choice `gorm:"column:choices;foreignKey:QuestionId;references:Id" json:"choices"`
}

type Choice struct {
	Id         int64  `gorm:"column:id" json:"id"`
	QuestionId string `gorm:"column:question_id" json:"question_id"`
	Content    string `gorm:"column:content" json:"content"`
}
