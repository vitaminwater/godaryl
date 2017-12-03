package model

import "errors"

func ToProtodef(m interface{}) (interface{}, error) {
	switch msg := m.(type) {
	case Habit:
		return msg.ToProtodef()
	case Message:
		return msg.ToProtodef()
	case Session:
		return msg.ToProtodef()
	}
	return nil, errors.New("Not a model")
}
