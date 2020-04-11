package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

// Store -
func Store(taskList *TaskList) error {
	raw, err := json.MarshalIndent(taskList, "", "    ")
	if err != nil {
		return err
	}

	taskFile, err := os.Create(getStorageDir() + "/task_list.json")
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

func getStorageDir() string {
	user, err := user.Current()
	if err != nil {
		dir, _ := os.Getwd()
		return dir
	}
	return user.HomeDir
}
