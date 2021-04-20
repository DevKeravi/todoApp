package model

import (
	"sort"
	"time"
)

type TodoData struct {
	Id        int       `json:"id"`
	Msg       string    `json:"message"`
	IsDone    bool      `json:"isdone"`
	CreatedAt time.Time `json:"createdat"`
}

type Success struct {
	Success bool `json:"success"`
}

var TodoList map[int]*TodoData
var size int

func listsort(list map[int]*TodoData) (index []int) {
	for k, _ := range list {
		index = append(index, k)
	}
	sort.Ints(index)
	return index
}

func NewModel() {
	TodoList = make(map[int]*TodoData)
	NewDB()
	size = 0
}

func Create(msg string) *TodoData {
	data := &TodoData{
		Id:        size,
		Msg:       msg,
		IsDone:    false,
		CreatedAt: time.Now(),
	}

	success := addDB(data)
	if success {
		size = getSize()
		return data
	}
	return nil

}

func Read() []*TodoData {
	/*
		list := []TodoData{}
		index := listsort(TodoList)

		for i := 0; i < len(TodoList); i++ {
			num := index[i]
			list = append(list, *(TodoList[num]))
		}
	*/
	list := getDB()

	return list

}

func Delete(id int) Success {
	/*
		if _, ok := TodoList[id]; ok {
			delete(TodoList, id)
			size--
			return Success{Success: true}
		}
		return Success{Success: false}
	*/
	if success := deleteDB(id); success {
		size = getSize()
		return Success{Success: true}
	}
	return Success{Success: false}

}
