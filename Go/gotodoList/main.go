package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var comandlist = []string{"help/查看帮助","edit/编辑模式","delete/删除模式" ,"exit/退出"}
var todoList = []string{}

func addToDo(reader bufio.Reader) {
	for {
		fmt.Print("请输入内容:")
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("输入内容有误，请重新输入:")
			continue
		}
		todoList = append(todoList, content)
		fmt.Println("是否继续编辑y/n")
		check,_ := reader.ReadString('\n')
		if check == "n\n" {
			break
		}
		continue
	}
}

func header() {
	fmt.Println("------------------------")
	if len(todoList) != 0 {
		fmt.Printf("%-10s %s\n","序号","待办事项")
		for k,v:=range todoList{
			fmt.Printf("%-15d %s",k+1,v)
		}
	}else{
		fmt.Println("还没有待办事项噢")
	}
	fmt.Println("请输入命令:")
}

func deleteItem(reader bufio.Reader){
	for{
		fmt.Println("请选择要删除的项目:")
		num,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("输入内容有误")
			continue
		}
		// fmt.Println(num[0:len(num)-1])
		// tmp := num[0:len(num)]
		currentNum,_:= strconv.Atoi(num[0:len(num)-1])
		if currentNum <= 0 || len(todoList) < int(currentNum) {
			fmt.Println("序号有误！")
			continue
		}else{
			if currentNum == 1 {
				todoList = todoList[1:]
			}else if currentNum == len(todoList){
				todoList = todoList[0:len(todoList)-1]
			}else {
				todoList = append(todoList[:currentNum - 1],todoList[currentNum:]...)
			}
			header()
			fmt.Println("是否继续删除?y/n")
			check,_ := reader.ReadString('\n')
			if check == "n\n"{
				break
			}else{
				continue
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		header()
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			panic("出现异常")
		}
		if command == "exit\n" {
			break
		} else if command == "delete\n"{
			deleteItem(*reader)
		} else if command == "edit\n" {
			addToDo(*reader)
			continue
		} else if command == "help\n" {
			for k, v := range comandlist {
				fmt.Println(k+1, v)
			}
			continue
		} else {
			fmt.Println("未知命令，请输入help查看命令列表")
			continue
		}
	}
	fmt.Println("退出程序")
}
