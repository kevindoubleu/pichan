package habits

type Scorecard struct {
	Id          int
	Name        string
	Connotation string
	Time        string
	Order       int
}

const (
	SCORECARD_SCHEMA = `
	id SERIAL,
	name VARCHAR(50) NOT NULL,
	connotation HABIT_CONNOTATION NOT NULL DEFAULT 'Neutral',
	time TIME,
	sortOrder INT NOT NULL DEFAULT 1000
	`
)
