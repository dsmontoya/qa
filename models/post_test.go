package models

import (
	"testing"

	"gorm.io/gorm"
)

func TestPost_Answer(t *testing.T) {
	type fields struct {
		Model      gorm.Model
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
				gorm.Model{},
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				Model:      tt.fields.Model,
				Content:    tt.fields.Content,
				QuestionID: tt.fields.QuestionID,
				Answers:    tt.fields.Answers,
			}
			if err := p.Answer(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Post.Answer() error = %v, wantErr %v", err, tt.wantErr)
			}
			db.Model(p).Association("Answers").Clear()
			db.Delete(p)
		})
	}
	if err := Close(); err != nil {
		t.Errorf("Close() error = %v", err)
	}
}
