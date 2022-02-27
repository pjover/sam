package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/sequence_type"
	"strings"
)

func ConvertSequencesToDbo(sequences []model.Sequence) []interface{} {
	var out []interface{}
	for _, sequence := range sequences {
		out = append(out, convertSequenceToDbo(sequence))
	}
	return out
}

func convertSequenceToDbo(sequence model.Sequence) Sequence {
	return Sequence{
		Id:      sequenceTypeValues[sequence.Id],
		Counter: sequence.Counter,
	}
}

func ConvertSequencesToModel(sequences []Sequence) []model.Sequence {
	var out []model.Sequence
	for _, sequence := range sequences {
		out = append(out, convertSequenceToModel(sequence))
	}
	return out
}

func convertSequenceToModel(sequence Sequence) model.Sequence {
	return model.Sequence{
		Id:      newSequenceType(sequence.Id),
		Counter: sequence.Counter,
	}
}

var sequenceTypeValues = []string{
	"",
	"STANDARD_INVOICE",
	"SPECIAL_INVOICE",
	"RECTIFICATION_INVOICE",
	"CUSTOMER",
}

func newSequenceType(value string) sequence_type.SequenceType {
	value = strings.ToLower(value)
	for i, val := range sequenceTypeValues {
		if strings.ToLower(val) == value {
			return sequence_type.SequenceType(i)
		}
	}
	return sequence_type.Invalid
}
