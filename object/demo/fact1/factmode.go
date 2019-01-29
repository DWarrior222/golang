package factmode

type student struct {
	Name  string
	score float64
}

func (s *student) GetScore() float64 {
	return s.score
}

func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}
