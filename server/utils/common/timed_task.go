package common

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
	"sophgo.com/mixmode/config"
)

type Timer interface {
	AddTaskByFunc(taskName string, spec string, task func(), option ...cron.Option) (cron.EntryID, error)
	AddTaskByJob(taskName string, spec string, job interface{ Run() }, option ...cron.Option) (cron.EntryID, error)
	FindCron(taskName string) (*cron.Cron, bool)
	StartTask(taskName string)
	StopTask(taskName string)
	Remove(taskName string, id int)
	Clear(taskName string)
	Close()
}

// TimerStruct 定时任务管理
type TimerStruct struct {
	TaskList map[string]*cron.Cron
	sync.Mutex
}

// AddTaskByFunc 通过函数的方法添加任务
func (t *TimerStruct) AddTaskByFunc(taskName string, spec string, task func(), option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.TaskList[taskName]; !ok {
		t.TaskList[taskName] = cron.New(option...)
	}
	id, err := t.TaskList[taskName].AddFunc(spec, task)
	t.TaskList[taskName].Start()
	return id, err
}

// AddTaskByFuncWithSeconds 通过函数的方法使用WithSeconds添加任务
func (t *TimerStruct) AddTaskByFuncWhithSecond(taskName string, spec string, task func(), option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.TaskList[taskName]; !ok {
		t.TaskList[taskName] = cron.New(option...)
	}
	id, err := t.TaskList[taskName].AddFunc(spec, task)
	t.TaskList[taskName].Start()
	return id, err
}

// AddTaskByJob 通过接口的方法添加任务
func (t *TimerStruct) AddTaskByJob(taskName string, spec string, job interface{ Run() }, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.TaskList[taskName]; !ok {
		t.TaskList[taskName] = cron.New(option...)
	}
	id, err := t.TaskList[taskName].AddJob(spec, job)
	t.TaskList[taskName].Start()
	return id, err
}

// AddTaskByJobWithSeconds 通过接口的方法添加任务
func (t *TimerStruct) AddTaskByJobWithSeconds(taskName string, spec string, job interface{ Run() }, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.TaskList[taskName]; !ok {
		t.TaskList[taskName] = cron.New(option...)
	}
	id, err := t.TaskList[taskName].AddJob(spec, job)
	t.TaskList[taskName].Start()
	return id, err
}

// FindCron 获取对应taskName的cron 可能会为空
func (t *TimerStruct) FindCron(taskName string) (*cron.Cron, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.TaskList[taskName]
	return v, ok
}

// StartTask 开始任务
func (t *TimerStruct) StartTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.TaskList[taskName]; ok {
		v.Start()
	}
}

// StopTask 停止任务
func (t *TimerStruct) StopTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.TaskList[taskName]; ok {
		v.Stop()
	}
}

// Remove 从taskName 删除指定任务
func (t *TimerStruct) Remove(taskName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.TaskList[taskName]; ok {
		v.Remove(cron.EntryID(id))
	}
}

// Clear 清除任务
func (t *TimerStruct) Clear(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.TaskList[taskName]; ok {
		v.Stop()
		delete(t.TaskList, taskName)
	}
}

// Close 释放资源
func (t *TimerStruct) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.TaskList {
		v.Stop()
	}
}

func NewTimerTask() Timer {
	return &TimerStruct{TaskList: make(map[string]*cron.Cron)}
}

func InitTimer(task func()) {
	period := config.Conf.Timer.Period
	if period <= 0 {
		period = 5
	}
	tm := NewTimerTask()
	tm.AddTaskByFunc("func", fmt.Sprintf("@every %ds", period), task)
}
