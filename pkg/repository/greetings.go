package repository

type GreetingsRepo struct {
	// db *sql.DB
}

func NewGreetingsRepository() *GreetingsRepo {
	return &GreetingsRepo{}
}

func (r *GreetingsRepo) GetGreeting() string {
	return "Hello,"
}
