package service

import (
	"github.com/labstack/echo/v4"
	"log"
	"mentoring-golang-rest-api/entity"
	"mentoring-golang-rest-api/entity/dto"
	"mentoring-golang-rest-api/repository"
	"net/http"
)

func CreateTodo(ctx echo.Context) error {
	var todoRequest dto.TodosRequest
	var todoResponse dto.TodosResponse
	var todo entity.Todo

	err := ctx.Bind(&todoRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "bad request")
	}

	todo = entity.Todo{
		Title:       todoRequest.Title,
		Description: todoRequest.Description,
	}

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	sqlQuery := "INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING id, check_flag"

	var id, checkFlag int
	err = db.QueryRow(sqlQuery, todo.Title, todo.Description).Scan(&id, &checkFlag)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "internal server error")
	}

	var checklFlagBool bool
	checkFlagBool := checklFlagBool

	todoResponse = dto.TodosResponse{
		ID:          id,
		Title:       todo.Title,
		Description: todo.Description,
		CheckFlag:   checkFlagBool,
	}

	return ctx.JSON(http.StatusOK, todoResponse)
}

func GetTodos(ctx echo.Context) error {
	var todoListResponse []dto.TodosResponse

	query := "SELECT id, title, description, check_flag FROM todos"

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	rows, err := db.Query(query)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "internal server error")
	}

	for rows.Next() {
		var todoResponse dto.TodosResponse
		err := rows.Scan(&todoResponse.ID, &todoResponse.Title, &todoResponse.Description, &todoResponse.CheckFlag)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, "internal server error")
		}
		todoListResponse = append(todoListResponse, todoResponse)
	}

	err = rows.Close()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "internal server error")
	}

	return ctx.JSON(http.StatusOK, todoListResponse)
}

func UpdateTodo(ctx echo.Context) error {
	var todoRequest dto.TodosRequest
	var todoResponse dto.TodosResponse

	id := ctx.Param("id")

	err := ctx.Bind(&todoRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "bad request")
	}

	query := "UPDATE todos SET title = $1, description = $2 WHERE id = $3 RETURNING id, title, description, check_flag"

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	err = db.QueryRow(query, todoRequest.Title, todoRequest.Description, id).Scan(&todoResponse.ID, &todoResponse.Title, &todoResponse.Description, &todoResponse.CheckFlag)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "internal server error")
	}

	return ctx.JSON(http.StatusOK, todoResponse)
}

func UpdateCheckFlag(ctx echo.Context) error {
	var todoRequest dto.TodoCheckFlag
	var todoResponse dto.TodosResponse

	id := ctx.Param("id")

	err := ctx.Bind(&todoRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "bad request")
	}

	var checkFlag int
	if todoRequest.CheckFlag == true {
		checkFlag = 1
	} else {
		checkFlag = 0
	}

	query := "UPDATE todos SET check_flag = $1 WHERE id =  $2 RETURNING id, title, description, check_flag"

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	err = db.QueryRow(query, checkFlag, id).Scan(&todoResponse.ID, &todoResponse.Title, &todoResponse.Description, &checkFlag)

	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, "internal server error")
	}
	if checkFlag == 1 {
		todoResponse.CheckFlag = true
	} else {
		todoResponse.CheckFlag = false
	}
	return ctx.JSON(http.StatusOK, todoResponse)
}

func DeleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	query := "DELETE FROM todos WHERE id = $1"

	db := repository.ConnectDB()
	defer repository.CloseDB(db)

	result, err := db.Exec(query, id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "internal server error")
	}

	res, err := result.RowsAffected()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "internal server error")
	}

	if res == 0 {
		return ctx.JSON(http.StatusNotFound, "not found")
	}

	return ctx.JSON(http.StatusOK, "success")
}
