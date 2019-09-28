package saver

import (
	"goLearning/learn/helloworld/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	user := model.Profile{Name: "心灵", ID: "1124690023"}

	save(user)
}
