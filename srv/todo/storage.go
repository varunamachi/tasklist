package todo

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
