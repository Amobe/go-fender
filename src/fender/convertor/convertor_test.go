package convertor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineEndConvertor_dos2unix(t *testing.T) {
	originLine := "thi\r\ns is \r\n\na line\r\n"
	expectLine := "thi\ns is \n\na line\n"

	actualLine := lineEndConvertor(originLine, directionDos2Unix)

	assert.Equal(t, expectLine, actualLine)
}

func TestLineEndConvertor_unix2dos(t *testing.T) {
	originLine := "thi\ns is \r\na line\n"
	expectLine := "thi\r\ns is \r\na line\r\n"

	actualLine := lineEndConvertor(originLine, directionUnix2Dos)

	assert.Equal(t, expectLine, actualLine)
}

func TestLineEndConvertor_2way(t *testing.T) {
	originLine := "thi\ns is \r\na line\n"
	expectLine := "thi\ns is \na line\n"

	var actualLine string
	actualLine = lineEndConvertor(originLine, directionUnix2Dos)
	actualLine = lineEndConvertor(originLine, directionDos2Unix)

	assert.Equal(t, expectLine, actualLine)
}
