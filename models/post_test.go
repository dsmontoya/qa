package models

import (
	"testing"
)

func TestPost_Answer(t *testing.T) {
	type fields struct {
		Content    string
		QuestionID *uint
		Answers    []*Post
	}
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"answer",
			fields{
				"Is this a question?",
				nil,
				nil,
			},
			args{
				"Yes!",
			},
			false,
		},
	}
	if err := Connect(); err != nil {
		t.Errorf("Connect() error = %v", err)
	}
	defer Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				Content:    tt.fields.Content,
				QuestionID: tt.fields.QuestionID,
				Answers:    tt.fields.Answers,
			}
			answer, err := p.Answer(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post.Answer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if answer.ID == 0 {
				t.Errorf("invalid Post.Answer() id")
			}
			db.Model(p).Association("Answers").Clear()
			db.Delete(p)
		})
	}
}

func TestGetQuestions(t *testing.T) {
	if err := Connect(); err != nil {
		t.Errorf("Connect() error = %v", err)
	}
	defer Close()

	questions := []*Post{
		{
			Answers: []*Post{
				{},
				{},
				{},
			},
		},
		{
			Answers: []*Post{
				{},
			},
		},
		{
			Answers: []*Post{},
		},
	}
	db.Create(questions)
	defer func() {
		for _, question := range questions {
			db.Delete(question.Answers)
		}
		db.Delete(questions)
	}()
	getQuestions, err := GetQuestions()
	if err != nil {
		t.Errorf("GetQuestions() error = %v", err)
	}
	if len(getQuestions) != 3 {
		t.Errorf("no question was found")
	}
}
