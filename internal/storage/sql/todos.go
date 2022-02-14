package storage

import "fmt"

// Todo represents the todo entity stored in database.
type Todo struct {
	Common
	Description string `gorm:"column:description"`
	Completed   bool   `gorm:"column:completed"`
}

// TableName returns the table name with the schema.
func (Todo) TableName() string {
	return fmt.Sprintf("%s.%s", schema, "Todos")
}

// GetTodo returns a todo with the provided id.
func (s *SQLStorage) GetTodo(id string) (*Todo, error) {
	var todo Todo
	if err := first(s.db.Where("Id = ?", NewUUID(id)), &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

// ListTodos returns a todos list.
func (s *SQLStorage) ListTodos(ids []string) ([]*Todo, error) {
	var todos []*Todo
	if err := s.db.
		Order("CreatedAt desc").
		Find(&todos, ids).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
