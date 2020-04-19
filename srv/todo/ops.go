package todo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Load -
func Load() (*TaskList, error) {
	//var taskList TaskList
	taskList := NewTaskList()

	taskFile, err := os.Open(getStorageDir() + "/task_list.json")
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
