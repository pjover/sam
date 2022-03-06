package model

import "github.com/pjover/sam/internal/domain/model/sequence_type"

type Sequence struct {
	Id      sequence_type.SequenceType
	Counter int
}
