package e2e

import (
	"testing"

	"github.com/DRK-Blutspende-BaWueHe/go-hl7"
	"github.com/DRK-Blutspende-BaWueHe/go-hl7/lib/hl7v23"
	"github.com/stretchr/testify/assert"
)

func TestMarshalMSH(t *testing.T) {
	// Arrange
	messageString := "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|0||ER|ER||8859/1|\u000d"

	var err error
	var message hl7v23.ORM_001
	err = hl7.Unmarshal(
		[]byte(messageString),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	// Act
	var marshalledMessageBytes [][]byte
	marshalledMessageBytes, err = hl7.Marshal(
		message,
		hl7.StandardFieldSeparator,
		hl7.EncodingASCII,
		hl7.TimezoneEuropeBerlin,
		hl7.StandardNotation)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, messageString, string(marshalledMessageBytes[0]))
}

func TestMarshalPID(t *testing.T) {
	// Arrange
	var mshData string = "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|0||ER|ER||8859/1|\u000d"
	var pidData string = "PID|1|a^b~^c|00100M56016||Smith^Harry||19500412|M\u000d"
	filedata := mshData + pidData

	var message hl7v23.ORM_001
	err := hl7.Unmarshal(
		[]byte(filedata),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	// Act
	marshalledMessageBytes, err := hl7.Marshal(
		message,
		hl7.StandardFieldSeparator,
		hl7.EncodingASCII,
		hl7.TimezoneEuropeBerlin,
		hl7.StandardNotation)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, mshData, string(marshalledMessageBytes[0]))
	assert.Equal(t, pidData, string(marshalledMessageBytes[1]))
}
