package model

import "github.com/pjover/sam/internal/domain/model/sequence_type"

type Sequence struct {
	id      sequence_type.SequenceType
	counter int
}

func NewSequence(id sequence_type.SequenceType, counter int) Sequence {
	return Sequence{
		id:      id,
		counter: counter,
	}
}

func (s Sequence) Id() sequence_type.SequenceType {
	return s.id
}
func (s Sequence) Counter() int {
	return s.counter
}
