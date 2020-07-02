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
