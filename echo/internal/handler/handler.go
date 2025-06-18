package handler

import (
	"cmp"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"web_app/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type responseOk struct {
	Message string `json:"message"`
}

type responseError struct {
	Error string `json:"error"`
}

type TodoHandler struct {
	todoService *service.TodoService
	session     db.Session
}

func NewTodoHandler() *TodoHandler {
	databaseUrl := os.Getenv("GOOSE_DBSTRING")
	settings, err1 := postgresql.ParseURL(databaseUrl)
	session, err2 := postgresql.Open(settings)
	if err := cmp.Or(err1, err2); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	todoService := service.NewTodoService(session)
	return &TodoHandler{
		todoService: todoService,
		session:     session,
	}
}

func (h *TodoHandler) CloseDB() {
	h.session.Close()
}

func (h *TodoHandler) GetTodosHandler(c echo.Context) error {
	strPage := c.QueryParam("page")
	if len(strPage) == 0 {
		strPage = "1"
	}
	page, err := strconv.Atoi(strPage)
	if err != nil {
		log.Printf("TodoHandler: GetTodosHandler: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, responseError{Error: "Invalid page."})
	}

	todos, err := h.todoService.GetTodos(uint(page))
	if err != nil {
		return c.JSON(http.StatusNotFound, responseError{Error: "Todos not found."})
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("TodoHandler: GetTodoHandler: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, responseError{Error: "Invalid id."})
	}

	todo, err := h.todoService.GetTodoById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, responseError{Error: "Todo not found."})
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodoHandler(c echo.Context) error {
	title := c.FormValue("title")
	err := h.todoService.CreateTodo(title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseError{Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, responseOk{Message: "Todo successfully created."})
}

func (h *TodoHandler) UpdateTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("TodoHandler: UpdateTodoHandler: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, responseError{Error: "Invalid id."})
	}

	title := c.FormValue("title")
	err = h.todoService.UpdateTodo(uint(id), title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseError{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, responseOk{Message: "Todo successfully updated."})
}

func (h *TodoHandler) DeleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("TodoHandler: DeleteTodoHandler: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, responseError{Error: "Invalid id."})
	}

	err = h.todoService.DeleteTodo(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, responseError{Error: "Todo not found."})
	}

	return c.JSON(http.StatusOK, responseOk{Message: "Todo successfully deleted."})
}
