package poll

import (
	"errors"
	"time"

	"appengine/datastore"
)

type Poll struct {
	UID      *datastore.Key `json:"uid" datastore:"-"`
	Name     string         `json:"name"`
	Question string         `json:"question"`
	Answers  []Answer       `json:"answers,omitempty"`
	Created  time.Time      `json:"created"`
	Updated  time.Time      `json:"updated"`
	Status   int            `json:"status"`
}

type Answer struct {
	Answer string `json:"answer"`
	Votes  int    `json:"votes"`
}

func (p *Poll) IsValid() error {
	if p.Name == "" {
		p.Name = "[" + time.Now().Format(time.RFC822) + "] " + p.Question
	}
	if p.Question == "" {
		return errors.New("Question can't be empty")
	}
	if p.Answers == nil {
		return errors.New("Answers can't be empty")
	}
	if len(p.Answers) < 2 {
		return errors.New("A poll must have two answers at least.")
	}

	for _, element := range p.Answers {
		if element.Answer == "" {
			return errors.New("A poll option can't be empty.")
		}
	}

	return nil
}
