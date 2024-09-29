package application

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

// Интерфейс для репозитория задач
type TaskRepository interface {
	Create(task domain.Task) (domain.Task, error)
}

// TaskUsecase - структура для работы с задачами
type TaskUsecase struct {
	repo TaskRepository
}

// NewTaskUsecase создает новый экземпляр юзкейса
func NewTaskUsecase(repo TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

// CreateTask создает новую задачу
func (u *TaskUsecase) CreateTask(title string) (domain.Task, error) {
	task := domain.Task{Title: title}
	return u.repo.Create(task)
}
