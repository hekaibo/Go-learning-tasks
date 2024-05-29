package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// 定义一个任务
type Task struct {
	ID      int
	Message string
	Time    int
}

// 管理任务
type TaskManager struct {
	mu    sync.Mutex
	tasks map[int]*Task
}

// 创建一个新的TaskManager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[int]*Task),
	}
}

// 添加一个任务
func (tm *TaskManager) AddTask(task *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks[task.ID] = task
	go tm.runTask(task)
}

// 运行一个任务
func (tm *TaskManager) runTask(task *Task) {
	select {
	case <-time.After(time.Duration(task.Time) * time.Second):
		fmt.Printf("任务id:%d,延迟%ds完成,输出内容: %s\n", task.ID, task.Time, task.Message)

	}
}

// 处理添加任务的HTTP请求
func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	id, _ := strconv.Atoi(q["ID"][0])
	Msg := q["Message"][0]
	taskT, _ := strconv.Atoi(q["Time"][0])

	// 假设task已经被填充了ID、Message和Time
	task := &Task{
		ID:      id,
		Message: Msg,
		Time:    taskT,
	}

	taskManager.AddTask(task)

	fmt.Fprint(w, "任务添加成功")
}

var taskManager *TaskManager // 全局TaskManager实例（为了简单起见）

func main() {
	taskManager = NewTaskManager()

	http.HandleFunc("/add-task", addTaskHandler)
	fmt.Println("Server started on :8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
