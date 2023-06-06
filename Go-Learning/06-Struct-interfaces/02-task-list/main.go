package main

import "fmt"


type TaskList struct{
	tasks []*Task		//Aca se van a guardar objetos de la estructura Task
}

func (tl *TaskList) appendTask(t *Task){
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask(index int){
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)		//esta linea es para eliminar un elemento de un slice
}

//Tasks
type Task struct{
	name string
	description string
	done bool
}

func (t *Task) toPrint(){
	fmt.Printf("Nombre ; %s,\n Descripcion: %s, Estado: %t\n", t.name, t.description, t.done)
}

func (t *Task)markDone(){
	t.done = true
}

func main(){
	t1 := Task{
		name: "Tarea1",
		description: "Descripcion de la tarea1",
		done: false,
	}
	t2 := Task{
		name: "Tarea2",
		description: "Descripcion de la tarea2",
		done: true,
	}
	

	list:= TaskList{}
	list.appendTask(&t1)
	list.appendTask(&t2)
	

	t3 := Task{
		name: "Tarea3",
		description: "Descripcion de la tarea3",
		done: false,
	}

	list.appendTask(&t3)
	fmt.Println(list)
	list.deleteTask(1)	//se elimina la tarea 2

	for i, task:= range list.tasks{
		fmt.Println(i,task.name)
		
	}

	
}