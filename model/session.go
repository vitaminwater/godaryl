package model

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vitaminwater/daryl/protodef"
)

type SessionSlice struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Habit Habit     `json:"habit"`
}

func (s SessionSlice) ToProtodef() (*protodef.SessionSlice, error) {
	start, err := ptypes.TimestampProto(s.Start)
	if err != nil {
		return nil, err
	}
	end, err := ptypes.TimestampProto(s.End)
	if err != nil {
		return nil, err
	}
	h, err := s.Habit.ToProtodef()
	if err != nil {
		return nil, err
	}

	return &protodef.SessionSlice{
		Start: start,
		End:   end,
		Habit: h,
	}, nil
}

type Session struct {
	Start  time.Time      `json:"start"`
	End    time.Time      `json:"end"`
	Slices []SessionSlice `json:"slices"`
}

func (s Session) ToProtodef() (*protodef.Session, error) {
	start, err := ptypes.TimestampProto(s.Start)
	if err != nil {
		return nil, err
	}
	end, err := ptypes.TimestampProto(s.End)
	if err != nil {
		return nil, err
	}
	slices := []*protodef.SessionSlice{}
	for _, slice := range s.Slices {
		sl, err := slice.ToProtodef()
		if err != nil {
			return nil, err
		}
		slices = append(slices, sl)
	}
	return &protodef.Session{
		Start:  start,
		End:    end,
		Slices: slices,
	}, nil
}

func NewSessionFromProtodef(d Daryl, ps *protodef.Session) (Session, error) {
	start, err := ptypes.Timestamp(ps.Start)
	if err != nil {
		return Session{}, err
	}
	end, err := ptypes.Timestamp(ps.End)
	if err != nil {
		return Session{}, err
	}
	s := Session{
		Start:  start,
		End:    end,
		Slices: make([]SessionSlice, 0),
	}
	for _, ss := range ps.Slices {
		ss, err := NewSessionSliceFromProtodef(d, ss)
		if err != nil {
			return s, err
		}
		s.Slices = append(s.Slices, ss)
	}
	return s, nil
}

func NewSessionSliceFromProtodef(d Daryl, ss *protodef.SessionSlice) (SessionSlice, error) {
	start, err := ptypes.Timestamp(ss.Start)
	if err != nil {
		return SessionSlice{}, err
	}
	end, err := ptypes.Timestamp(ss.End)
	if err != nil {
		return SessionSlice{}, err
	}
	h, err := NewHabitFromProtodef(d.Id, ss.Habit)
	if err != nil {
		return SessionSlice{}, err
	}
	return SessionSlice{
		Start: start,
		End:   end,
		Habit: h,
	}, nil
}
