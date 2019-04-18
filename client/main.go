package main

import (
	"log"
	"net/rpc"

	"go.RPC_S/tasks"
)

func taskSample() {
	var err error

	var reply tasks.ToDo
	var slice []tasks.ToDo

	client, err := rpc.DialHTTP("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	finishApp := tasks.ToDo{"Finish App", "Started"}
	makeDinner := tasks.ToDo{"Make Dinner", "Not Started"}
	walkDog := tasks.ToDo{"Walk the dog", "Not Started"}

	memUsage := tasks.ToDo{"Mem usage", ""}
	client.Call("Task.MakeToDo", memUsage, &reply)
	log.Println(reply)
	client.Call("Task.MakeToDo", finishApp, &reply)
	log.Println(reply)
	client.Call("Task.MakeToDo", makeDinner, &reply)
	log.Println(reply)
	client.Call("Task.MakeToDo", walkDog, &reply)
	log.Println(reply)

	client.Call("Task.DeleteToDo", makeDinner, &reply)
	log.Println(reply)

	client.Call("Task.MakeToDo", makeDinner, &reply)
	log.Println(reply)

	client.Call("Task.GetToDo", "Finish App", &reply)
	log.Println("Finish App: ", reply)

	client.Call("Task.GetResource", "Resources", &reply)
	log.Println("Resources: ", reply)

	err = client.Call("Task.EditToDo", tasks.EditTodo{"Finish App", "Finish App", "Completed"}, &reply)
	if err != nil {
		log.Fatal("Problem editing ToDo: ", err)
	}

	client.Call("Task.GetSlice", "", &slice)
	log.Println("Slice: ", slice)
}

func rsrcSample() {
	var err error

	client, err := rpc.DialHTTP("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	var rsrcReply tasks.Rsrc
	var rsrcSlice []tasks.Rsrc

	memUsage := tasks.RsrcStruct{"Mem", "Get"}
	CPUUsage := tasks.RsrcStruct{"CPU", "Get"}
	IOUsage := tasks.RsrcStruct{"IO", "Get"}

	client.Call("Rsrc.MakeRsrcUsage", memUsage, &rsrcReply)
	log.Println(rsrcReply)
	client.Call("Rsrc.MakeRsrcUsage", CPUUsage, &rsrcReply)
	client.Call("Rsrc.MakeRsrcUsage", IOUsage, &rsrcReply)

	client.Call("Rsrc.GetSlice", "", &rsrcSlice)
	log.Println("rsrcSlice: ", rsrcSlice)
}

func main() {
	taskSample()
	// rsrcSample()
}
