package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseTimestamp_ValidValue(t *testing.T) {
	timestamp, err := ParseTimestamp("2020-11-23T15:44:55Z")

	assert.Nil(t, err)
	assert.NotNil(t, timestamp)
	assert.Equal(t, 2020, timestamp.Year())
	assert.Equal(t, time.Month(11), timestamp.Month())
	assert.Equal(t, 23, timestamp.Day())
	assert.Equal(t, 15, timestamp.Hour())
	assert.Equal(t, 44, timestamp.Minute())
	assert.Equal(t, 55, timestamp.Second())
}

func TestParseTimestamp_ValidValueWithZone(t *testing.T) {
	timestamp, err := ParseTimestamp("2020-11-23T15:44:55+03:30")

	assert.Nil(t, err)
	assert.NotNil(t, timestamp)
	assert.Equal(t, 2020, timestamp.Year())
	assert.Equal(t, time.Month(11), timestamp.Month())
	assert.Equal(t, 23, timestamp.Day())
	assert.Equal(t, 15, timestamp.Hour())
	assert.Equal(t, 44, timestamp.Minute())
	assert.Equal(t, 55, timestamp.Second())
	zone, offset := timestamp.Zone()
	assert.Equal(t, "", zone)
	assert.Equal(t, int(3.5*60*60), offset)
}

func TestParseTimestamp_InvalidValue(t *testing.T) {
	testData := []string{
		"202-11-23T15:44:55Z",
		"2020-14-23T15:44:55Z",
		"2020-11-37T15:44:55Z",
		"2020-11-23T35:44:55Z",
		"2020-11-23T105:44:55Z",
		"2020-11-23T15:74:55Z",
		"2020-11-23T15",
		"2020-11-23T15:74:55Z",
		"2020-11-23T15:44:95Z",
		"2020-11-23",
	}
	for _, value := range testData {
		timestamp, err := ParseTimestamp(value)
		t.Run(value, func(t *testing.T) {
			assert.NotNil(t, err)
			assert.NotNil(t, timestamp)
			assert.Equal(t, 1, timestamp.Year())
		})
	}
}

func TestParseTimestampBy_ValidShortValue(t *testing.T) {
	timestamp, err := ParseTimestampBy("2006-01-02", "2020-11-23")
	assert.Nil(t, err)
	assert.NotNil(t, timestamp)
	assert.Equal(t, 2020, timestamp.Year())
	assert.Equal(t, time.Month(11), timestamp.Month())
	assert.Equal(t, 23, timestamp.Day())
	assert.Equal(t, 0, timestamp.Hour())
	assert.Equal(t, 0, timestamp.Minute())
	assert.Equal(t, 0, timestamp.Second())
}

func TestParseTimestampBy_ValidLongValue(t *testing.T) {
	timestamp, err := ParseTimestampBy("2006-01-02", "2020-11-23T15:44:55Z")
	assert.Nil(t, err)
	assert.NotNil(t, timestamp)
	assert.Equal(t, 2020, timestamp.Year())
	assert.Equal(t, time.Month(11), timestamp.Month())
	assert.Equal(t, 23, timestamp.Day())
	assert.Equal(t, 15, timestamp.Hour())
	assert.Equal(t, 44, timestamp.Minute())
	assert.Equal(t, 55, timestamp.Second())
}

func TestParseTimestampBy_InvalidValue(t *testing.T) {
	testData := []string{
		"202-11-23T15:44:55Z",
		"2020-14-23T15:44:55Z",
		"2020-11-37T15:44:55Z",
		"2020-11-23T35:44:55Z",
		"2020-11-23T105:44:55Z",
		"2020-11-23T15:74:55Z",
		"2020-11-23T15",
		"2020-11-23T15:74:55Z",
		"2020-11-23T15:44:95Z",
	}
	for _, value := range testData {
		t.Run(value, func(t *testing.T) {
			timestamp, err := ParseTimestampBy("2006-01-02", value)
			assert.NotNil(t, err)
			assert.NotNil(t, timestamp)
			assert.Equal(t, 1, timestamp.Year())
		})
	}
}
