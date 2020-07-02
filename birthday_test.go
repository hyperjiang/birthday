package birthday

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBirthday(t *testing.T) {
	should := require.New(t)

	year := time.Now().Format("2006")

	tests := []struct {
		date          string
		age           int
		constellation string
	}{
		{
			year + "-01-01",
			0,
			"Capricorn",
		},
		{
			year + "-3-01",
			0,
			"Pisces",
		},
	}

	for _, test := range tests {
		birthday, err := ParseBirthday(test.date)
		should.NoError(err)
		should.Equal(test.age, birthday.Age)
		should.Equal(test.constellation, birthday.Constellation)
	}
}

func TestFormat(t *testing.T) {
	should := require.New(t)

	birthday, err := ParseBirthday("2020-07-02")
	should.NoError(err)
	should.Equal("2/7/2020", birthday.Format("2/1/2006"))
	should.Equal("Jul 2nd, 2020", birthday.Format("Jan 2nd, 2006"))
}

func TestParseConstellation(t *testing.T) {
	should := require.New(t)

	tests := []struct {
		month         time.Month
		day           int
		constellation string
	}{
		{
			time.January,
			1,
			"Capricorn",
		},
		{
			time.January,
			28,
			"Aquarius",
		},
		{
			time.February,
			1,
			"Aquarius",
		},
		{
			time.February,
			28,
			"Pisces",
		},
		{
			time.March,
			1,
			"Pisces",
		},
		{
			time.March,
			28,
			"Aries",
		},
		{
			time.April,
			1,
			"Aries",
		},
		{
			time.April,
			28,
			"Taurus",
		},
		{
			time.May,
			1,
			"Taurus",
		},
		{
			time.May,
			28,
			"Gemini",
		},
		{
			time.June,
			1,
			"Gemini",
		},
		{
			time.June,
			28,
			"Cancer",
		},
		{
			time.July,
			1,
			"Cancer",
		},
		{
			time.July,
			28,
			"Leo",
		},
		{
			time.August,
			1,
			"Leo",
		},
		{
			time.August,
			28,
			"Virgo",
		},
		{
			time.September,
			1,
			"Virgo",
		},
		{
			time.September,
			28,
			"Libra",
		},
		{
			time.October,
			1,
			"Libra",
		},
		{
			time.October,
			28,
			"Scorpio",
		},
		{
			time.November,
			1,
			"Scorpio",
		},
		{
			time.November,
			28,
			"Sagittarius",
		},
		{
			time.December,
			1,
			"Sagittarius",
		},
		{
			time.December,
			28,
			"Capricorn",
		},
		{
			13,
			1,
			"Unknown",
		},
	}

	for _, test := range tests {
		should.Equal(test.constellation, parseConstellation(test.month, test.day))
	}
}
