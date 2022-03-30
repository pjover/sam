package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/sequence_type"
)

func ConvertSequencesToDbo(sequences []model.Sequence) []interface{} {
	var out []interface{}
	for _, sequence := range sequences {
		out = append(out, ConvertSequenceToDbo(sequence))
	}
	return out
}

func ConvertSequenceToDbo(sequence model.Sequence) Sequence {
	return Sequence{
		Id:      sequence.Id.String(),
		Counter: sequence.Counter,
	}
}

func ConvertSequencesToModel(sequences []Sequence) []model.Sequence {
	var out []model.Sequence
	for _, sequence := range sequences {
		out = append(out, ConvertSequenceToModel(sequence))
	}
	return out
}

func ConvertSequenceToModel(sequence Sequence) model.Sequence {
	return model.Sequence{
		Id:      sequence_type.NewSequenceType(sequence.Id),
		Counter: sequence.Counter,
	}
}
