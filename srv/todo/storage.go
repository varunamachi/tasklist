package todo

// TaskList -
type TaskList struct {
	Tasks  []*TaskItem `json:"tasks"`
	NextID int         `json:"nextID"`
}

func (tl *TaskList) nextID() int {
	id := tl.NextID
	tl.NextID++
	return id
}

// Add -
func (tl *TaskList) Add(items ...*TaskItem) {
	for _, task := range items {
		task.ID = tl.nextID()
		tl.Tasks = append(tl.Tasks, task)
	}
}

// Remove -
func (tl *TaskList) Remove(id int) {
	index := -1
	for i, task := range tl.Tasks {
		if task.ID == id {
			index = i
			break
		}
	}
	tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
}

// Iterate -
func (tl *TaskList) Iterate(operation func(*TaskItem)) {
	for _, task := range tl.Tasks {
		operation(task)
	}
}

// NewTaskList -
func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: make([]*TaskItem, 0, 100),
	}
}

// JSONStorage -
type JSONStorage struct {
}

// Init -
func (pg *JSONStorage) Init() error {
	return nil
}

// Add -
func (pg *JSONStorage) Add(ti *TaskItem) error {
	return nil
}

// Remove -
func (pg *JSONStorage) Remove(id int) error {
	return nil
}

// Name -
func (pg *JSONStorage) Name() string {
	return "pg"
}

// Update -
func (pg *JSONStorage) Update(item *TaskItem) error {
	return nil
}

// Retrieve -
func (pg *JSONStorage) Retrieve(id int) error {
	return nil
}

// Bulk -
func (pg *JSONStorage) Bulk(op BulkOp) error {
	return nil
}

// RetrieveAll -
func (pg *JSONStorage) RetrieveAll() []*TaskItem {
	return nil
}
