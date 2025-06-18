package service

import (
	"cmp"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/upper/db/v4"
)

var validate *validator.Validate

type TodoService struct {
	session db.Session
}

type Todo struct {
	ID    uint   `db:"id,omitempty" json:"id"`
	Title string `db:"title"        json:"title" validate:"required,gte=5,lte=255"`
}

func (t *Todo) Store(sess db.Session) db.Store {
	return sess.Collection("todos_go")
}

func NewTodoService(session db.Session) *TodoService {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return &TodoService{session: session}
}

func (s *TodoService) GetTodos(page uint) ([]Todo, error) {
	res := s.session.Collection("todos_go").Find()
	var listTodos []Todo
	p := res.Paginate(10)
	err := p.Page(page).All(&listTodos)
	if err != nil {
		log.Printf("TodoService: GetTodos(%d): %s\n", page, err.Error())
		return nil, err
	}

	return listTodos, nil
}

func (s *TodoService) GetTodoById(id uint) (Todo, error) {
	var todo Todo
	err := s.session.Get(&todo, db.Cond{"id": id})
	if err != nil {
		log.Printf("TodoService: GetTodoById(%d): %s\n", id, err.Error())
		return Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) CreateTodo(title string) error {
	todo := Todo{Title: title}
	err := validate.Struct(todo)
	if err != nil {
		log.Printf("TodoService: CreateTodo(%s): %s\n", title, err.Error())
		return err
	}

	err = s.session.Save(&todo)
	if err != nil {
		log.Printf("TodoService: CreateTodo(%s): %s\n", title, err.Error())
		return err
	}

	return nil
}

func (s *TodoService) UpdateTodo(id uint, title string) error {
	todo, err1 := s.GetTodoById(id)
	todo.Title = title
	err2 := validate.Struct(todo)
	if err := cmp.Or(err1, err2); err != nil {
		log.Printf("TodoService: UpdateTodo(%d, %s): %s\n", id, title, err.Error())
		return err
	}

	err := s.session.Save(&todo)
	if err != nil {
		log.Printf("TodoService: UpdateTodo(%d, %s): %s\n", id, title, err.Error())
		return err
	}

	return nil
}

func (s *TodoService) DeleteTodo(id uint) error {
	todo, err := s.GetTodoById(id)
	if err != nil {
		return err
	}

	err = s.session.Delete(&todo)
	if err != nil {
		log.Printf("TodoService: DeleteTodo(%d): %s\n", id, err.Error())
		return err
	}

	return nil
}
