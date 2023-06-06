package main
import "fmt"

type User struct{
	Name string
	Email string
	active bool
}

type Student struct{
	user User		//Colocamos una variables para acceder a sus datos
	code string
}

//Relacion uno a muchos
type Curso struct{
	name string
	videos[]Video
}

type Video struct{
	title string
	curso Curso
}



func main(){
	//Relacion uno a uno
	/*
	alex := User{
		Name : "Alex",
		Email: "alex@gmail.com",
		active: true,
}
	klen := User{
		Name : "Klen",
		Email: "Klen@gmail.com",
		active: true,
	}

	alexStudent := Student{
		user: alex,
		code: "123",
	}

	fmt.Println(klen)
	fmt.Println(alexStudent.user.Name)
	*/

	//Relacion uno a muchos
	video1 := Video{
		title: "Video1",
	}
	video2 := Video{
		title: "Video2",
	}

	curso:= Curso{
		name: "Curso de Go",
		videos: []Video{video1, video2},
	}
	video1.curso = curso	//Asignamos el curso al video1
	video2.curso = curso	//Asignamos el curso al video2

	fmt.Println(video1.curso.name)

	for _,video := range curso.videos{
		fmt.Println(video.title)
	}

}