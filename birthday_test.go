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
		should.NoError((err))
		should.Equal(test.age, birthday.Age)
		should.Equal(test.constellation, birthday.Constellation)
	}
}
