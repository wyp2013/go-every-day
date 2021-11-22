package workerpool

type Job interface {}
type ProcessJobFunc func(Job) error