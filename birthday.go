package birthday

import (
	"time"

	"github.com/hyperjiang/php"
)

// Birthday is the birthday detail
type Birthday struct {
	Date          time.Time
	Year          int
	Month         time.Month
	Day           int
	Age           int
	Constellation string
}

// ParseBirthdayFromTime parse birthday from given time
func ParseBirthdayFromTime(t time.Time) *Birthday {
	birthday := &Birthday{
		Date:  t,
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}

	now := time.Now()

	birthday.Age = now.Year() - birthday.Year
	if now.YearDay() < t.YearDay() && birthday.Age > 0 {
		birthday.Age--
	}

	birthday.Constellation = parseConstellation(birthday.Month, birthday.Day)

	return birthday
}

// ParseBirthday parse birthday from given string date
func ParseBirthday(date string) (*Birthday, error) {
	t, err := php.DateCreate(date)
	if err != nil {
		return nil, err
	}

	return ParseBirthdayFromTime(t), nil
}

func parseConstellation(month time.Month, day int) string {
	switch month {
	case time.January:
		if day >= 20 {
			return "Aquarius"
		}
		return "Capricorn"
	case time.February:
		if day >= 19 {
			return "Pisces"
		}
		return "Aquarius"
	case time.March:
		if day >= 21 {
			return "Aries"
		}
		return "Pisces"
	case time.April:
		if day >= 20 {
			return "Taurus"
		}
		return "Aries"
	case time.May:
		if day >= 21 {
			return "Gemini"
		}
		return "Taurus"
	case time.June:
		if day >= 22 {
			return "Cancer"
		}
		return "Gemini"
	case time.July:
		if day >= 23 {
			return "Leo"
		}
		return "Cancer"
	case time.August:
		if day >= 23 {
			return "Virgo"
		}
		return "Leo"
	case time.September:
		if day >= 23 {
			return "Libra"
		}
		return "Virgo"
	case time.October:
		if day >= 24 {
			return "Scorpio"
		}
		return "Libra"
	case time.November:
		if day >= 23 {
			return "Sagittarius"
		}
		return "Scorpio"
	case time.December:
		if day >= 22 {
			return "Capricorn"
		}
		return "Sagittarius"
	}

	return "Unknown"
}
