package fj

type (
	WorkType interface{}

	taskWorker func(WorkType) WorkType

	ForkJoinModeler interface {
		Fork(f taskWorker, w WorkType)
		Join()
		Result(i int) WorkType
		Results() (res []WorkType)
	}
)
