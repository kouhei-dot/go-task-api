package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var db *sql.DB
var ConnectionString string

const selectSql = `select * from Tasks;`
const createSql = "insert into Tasks (uuid, label, content, created_at, updated_at) values ($1, $2, $3, $4, $5);"
const deleteSql = "delete from Tasks where id = $1"
const updateSql = "update Tasks set label = $1 where id = $2"
const findLastInsertIdSql = "select id from Tasks order by id desc limit 1;"

type Task struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Label     string    `json:"label"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Tasks []Task

type ReqTask struct {
	Label   string `json:"label"`
	Content string `json:"content"`
}

type ReqUpdateTask struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

func (reqTask *ReqTask) CreateTask() (Task, error) {
	log.Println("CreateTask start")
	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		log.Println("connection Error")
		return Task{}, err
	}
	defer db.Close()

	var task Task
	task.UUID = genUUID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	_, err = db.Exec(createSql, task.UUID, reqTask.Label, reqTask.Content, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		log.Println("sql exec Error")
		return Task{}, err
	}

	row := db.QueryRow(findLastInsertIdSql)
	var insertedId int
	row.Scan(&insertedId)

	task.ID = insertedId
	task.Label = reqTask.Label
	task.Content = reqTask.Content
	return task, nil
}

func (reqTask *ReqUpdateTask) UpdateTask() error {
	log.Println("UpdateTask start")
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(updateSql, reqTask.Label, reqTask.ID)
	if err != nil {
		log.Println("sql exec Error")
		return err
	}
	return nil
}

func GetTasks() (Tasks, error) {
	log.Println("GetTasks start")
	db, err := sql.Open("postgres", ConnectionString)
	defer db.Close()
	if err != nil {
		log.Println("connection Error")
		return nil, err
	}

	rows, err := db.Query(selectSql)
	if err != nil {
		log.Println("select Error")
		return nil, err
	}

	var tasks Tasks
	var task Task
	for rows.Next() {
		err := rows.Scan(
			&task.ID, &task.UUID, &task.Label, &task.Content, &task.CreatedAt, &task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func DeleteTask(id int) error {
	log.Println("DeleteTask start")
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(deleteSql, id)
	return nil
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		log.Println("connection Error")
		return nil, err
	}
	return db, nil
}

func genUUID() string {
	var uuid, err = uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return uuid.String()
}
