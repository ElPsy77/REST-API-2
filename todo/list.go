package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

// когда в конструкторе лучше передовать значение по указателю а когда копию
func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

// что если будеть одинаковое название задач
func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExist
	}

	l.tasks[task.Title] = task

	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return task, nil
}

// чем черевато передовать мапу по указател/
func (l *List) ListTask() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Task, len(l.tasks))

	for k, v := range l.tasks {
		tmp[k] = v
	}

	return tmp
}

func (l *List) ListUncompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	uncompletedTasks := make(map[string]Task)

	for k, v := range l.tasks {
		if !v.Completed {
			uncompletedTasks[k] = v
		}
	}

	return uncompletedTasks
}

func (l *List) CompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Complete()

	l.tasks[title] = task

	return l.tasks[title], nil
}

func (l *List) UncompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Uncomplete()

	l.tasks[title] = task

	return l.tasks[title], nil
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[title]; !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
