package decorator

import "fmt"

// User represents a user entity.
type User struct {
	ID   string
	Name string
}

// UserService defines the business contract.
type UserService interface {
	GetUser(id string) (*User, error)
}

// Logger abstracts logging dependency.
type Logger interface {
	Info(message string)
}

// Concrete implementation.
type userService struct {
	users map[string]*User
}

// NewUserService creates base service.
func NewUserService() UserService {

	return &userService{
		users: map[string]*User{
			"1": {
				ID:   "1",
				Name: "Imran",
			},
		},
	}
}

func (u *userService) GetUser(
	id string,
) (*User, error) {

	user, ok := u.users[id]

	if !ok {
		return nil, fmt.Errorf(
			"user not found",
		)
	}

	return user, nil
}

// Logging decorator.

type LoggingDecorator struct {
	service UserService
	logger  Logger
}

func NewLoggingDecorator(
	service UserService,
	logger Logger,
) UserService {

	return &LoggingDecorator{
		service: service,
		logger:  logger,
	}
}

func (l *LoggingDecorator) GetUser(
	id string,
) (*User, error) {

	l.logger.Info(
		fmt.Sprintf("fetching user: %s", id),
	)

	return l.service.GetUser(id)
}

// Metrics decorator.
type Metrics interface {
	Increment(metric string)
}

type MetricsDecorator struct {
	service UserService
	metrics Metrics
}

func NewMetricsDecorator(
	service UserService,
	metrics Metrics,
) UserService {

	return &MetricsDecorator{
		service: service,
		metrics: metrics,
	}
}

func (m *MetricsDecorator) GetUser(
	id string,
) (*User, error) {

	m.metrics.Increment("user_fetches")

	return m.service.GetUser(id)
}
