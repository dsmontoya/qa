package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content    string
	QuestionID *uint
	Answers    []*Post `gorm:"foreignkey:QuestionID"`
}

func (p *Post) Answer(content string) error {
	return db.Model(p).Association("Answers").Append(&Post{Content: content})
}

func GetQuestions() ([]*Post, error) {
	questions := []*Post{}
	d := db.Where("question_id is NULL").Find(&questions)
	if err := d.Error; err != nil {
		return nil, err
	}
	for _, question := range questions {
		if err := db.Model(question).Association("Answers").Find(&question.Answers); err != nil {
			return nil, err
		}
	}
	return questions, nil
}
