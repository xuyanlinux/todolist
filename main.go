package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	id = "id"
	name = "name"
	startTime = "startTime"
	endTime = "endTime"
	status = "status"
	user = "user"
)

var menu []string = []string{
	"0: 打印任务清单",
	"1: 查询任务",
	"2: 添加任务",
	"3: 删除任务",
	"4: 修改任务",
}

const (
	statusNew = "未执行"
	_statusComplete = "已执行"
)

var todos []map[string]string = []map[string]string{
{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "", "status": statusNew, "user": "kk"},
{"id": "2", "name": "备课", "startTime": "21:00", "endTime": "", "status": statusNew, "user": "kk"},
{"id": "4", "name": "复习", "startTime": "09:00", "endTime": "", "status": statusNew, "user": "kk"},
}

func printTask(task map[string]string){
	fmt.Println(strings.Repeat("-",30))
	fmt.Println("ID: ",task[id])
	fmt.Println("NAME: ",task[name])
	fmt.Println("startTime: ",task[startTime])
	fmt.Println("endTime: ",task[endTime])
	fmt.Println("status: ",task[status])
	fmt.Println("user: ",task[user])

}

func getId() string{
	var rt int = 0
	for _, task := range todos{
		val, _ := strconv.Atoi(task[id])
		if rt < val{
			rt = val
		}
	}
	return strconv.Itoa(rt + 1)
}

func input(str string)string  {
	var rt string = ""
	fmt.Printf("请输入%s:", str)
	fmt.Scan(&rt)
	return rt
}

func findTaskFromId(targetId string)int{
	for index, task := range  todos{
		if targetId == task[id]{
			fmt.Printf("找到的task id:%s\n 在todos中的索引是：%d\n", targetId, index)
			return index
		}
	}
	return -1
}
func queryTask()  {
	for _,task := range todos{
		printTask(task)
		//fmt.Printf("%v\n",task)
	}
}

func printTasks() {
	for _, task := range todos {
		fmt.Printf("%v\n", task)
	}
}

func creatTask() map[string]string {
	newTask := make(map[string]string)
	newTask[id] = getId()
	newTask[name] = ""
	newTask[startTime] = ""
	newTask[endTime] = ""
	newTask[status] = statusNew
	newTask[user] = ""
	return newTask
}

func delTask()  {
	inputID := input("将删除任务的ID")
	delIndex := findTaskFromId(inputID)
	if delIndex == -1{
		fmt.Println("输入的任务ID有错")
		return
	}
	fmt.Println(delIndex)
	todos = append(todos[:delIndex], todos[delIndex+1:]...)
}
func addTask()  {
	task := creatTask()
	fmt.Println("请输入任务信息")
	task[name] = input("任务名称")
	task[startTime] = input("开始时间")
	task[user] = input("负责人")
	todos = append(todos, task)
}

func selectTask() {
	targetID := input("输入要查找的任务ID")
	targetIndex := findTaskFromId(targetID)
	if targetIndex == -1{
		fmt.Println("输入的任务ID有错")
		return
	}
	printTask(todos[targetIndex])
}
func editTask() {
	targetID := input("输入要查找的任务ID")
	targetIndex := findTaskFromId(targetID)
	if targetIndex == -1{
		fmt.Println("输入的任务ID有错")
		return
	}
	task := todos[targetIndex]
	input1 := input("任务状态,0 表示未完成，1表示完成")
	if input1 == "0"{
		task[status] = statusNew
	} else if input1 == "1"{
		task[status] = _statusComplete
		task[endTime] = time.Now().Format("03:04")
	} else {
		fmt.Println("输入的状态有误")
	}

}

func main() {
	for {
		fmt.Println("输入你需要进行操作的编号(输入q退出程序)：")
		for _, value := range menu{
			fmt.Printf("%v\n",value)
			//fmt.Println(index,":",value[index])
		}

		optNum := input("操作代号")
		switch optNum {
		case "0":
			printTasks()
		case "1":
			selectTask()
		case "2":
			addTask()
		case "3":
			delTask()
		case "4":
			editTask()
		case "q":
			return
		default:
			continue
		}
	con := input("退出(n|N)?继续(其他任意键)?")
	if strings.ToLower(con) != "n"{
		return
	}
	}


}