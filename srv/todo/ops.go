package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Store -
func Store(taskList *TaskList) error {
	raw, err := json.MarshalIndent(taskList, "", "    ")
	if err != nil {
		return err
	}

	taskFile, err := os.Create("/home/varun/task_list.json")
	if err != nil {
		return err
	}
	defer func() {
		taskFile.Close()
	}()

	fmt.Fprint(taskFile, string(raw))
	return nil
}

// Load -
func Load() (*TaskList, error) {
	//var taskList TaskList
	taskList := NewTaskList()

	taskFile, err := os.Open("/home/varun/task_list.json")
	if err != nil {
		return taskList, err
	}
	defer taskFile.Close()

	data, err := ioutil.ReadAll(taskFile)
	if err != nil {
		return taskList, err
	}

	err = json.Unmarshal(data, taskList)
	return taskList, err
}
