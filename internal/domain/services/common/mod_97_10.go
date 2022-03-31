package common

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

// Mod9710 calculates the check digits according to ISO-7604
type Mod9710 struct {
	codes []string
}

func NewMod9710(codes ...string) Mod9710 {
	return Mod9710{
		codes: codes,
	}
}

func (i Mod9710) CheckDigits() string {
	preparedParams := i.prepareCodes()
	assignedWeightsToLetters := i.assignWeightsToLetters(preparedParams)
	return i.apply9710Model(assignedWeightsToLetters)
}

func (i Mod9710) prepareCodes() string {
	rawCode := strings.Join(i.codes, "")
	return i.prepareCode(rawCode)
}

func (i Mod9710) prepareCode(code string) string {
	var preparedCode string
	if code != "" {
		preparedCode = strings.ReplaceAll(code, " ", "")
		preparedCode = strings.ReplaceAll(preparedCode, "-", "")
	}
	return fmt.Sprintf("%s00", preparedCode)
}

func (i Mod9710) assignWeightsToLetters(code string) string {
	var buffer bytes.Buffer
	for _, letter := range []rune(code) {
		weight := i.assignWeightToLetter(letter)
		buffer.WriteString(strconv.Itoa(weight))
	}
	return buffer.String()
}

func (i Mod9710) assignWeightToLetter(letter rune) int {
	intValue := int(letter)
	if letter >= 'A' {
		return intValue - 'A' + 10
	} else {
		return intValue - '0'
	}
}

func (i Mod9710) apply9710Model(input string) string {
	in, ok := new(big.Int).SetString(input, 10)
	if !ok {
		log.Fatalf("cannot convert %s to big integer", input)
	}
	mod97 := new(big.Int).Mod(in, big.NewInt(97)).Int64()
	return fmt.Sprintf("%02d", 98-mod97)
}
