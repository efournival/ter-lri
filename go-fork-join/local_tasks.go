package fj

import "sync"

type (
	localTask struct {
		worker taskWorker
		result WorkType
	}

	LocalTasks struct {
		sync.WaitGroup
		tasks []*localTask
	}
)

func NewLocalTasks() *LocalTasks {
	return &LocalTasks{}
}

func (t *LocalTasks) Fork(f taskWorker, w WorkType) {
	tsk := &localTask{f, nil}
	t.tasks = append(t.tasks, tsk)
	t.Add(1)

	go func() {
		tsk.result = f(w)
		t.Done()
	}()
}

func (t *LocalTasks) Join() {
	t.Wait()
}

func (t *LocalTasks) Result(i int) WorkType {
	return t.tasks[i].result
}

func (t *LocalTasks) Results() (res []WorkType) {
	for _, tsk := range t.tasks {
		res = append(res, tsk.result)
	}
	return
}
