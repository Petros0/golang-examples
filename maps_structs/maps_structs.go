package main

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Hey maps & structs!")

	
	//with make
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 3943101,
		"Texas": 43242,
	}
	//add 
	statePopulations["Georgia"] = 10000000
	//delete
	delete(statePopulations,"Georgia")
	fmt.Println(statePopulations)
	//get
	fmt.Println(statePopulations["Texas"])

	//the first variable is the value of the key
	//and the second is a boolean that says if the key exists or not
	value, exists := statePopulations["Ohio"]
	fmt.Println(value, exists)
	fmt.Println(len(statePopulations))

	//maps are also call by reference.


	//STRUCTS

	aDoctor := Doctor{
		number :3, 
		actorName: "Petros Stergioulas", 
		companions: []string {
			"Alkis",
			"Christos",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	//anonymosu structs
	//they are not with pass by reference
	//except: we use '&' e.g. anotherDoctor := &bDoctor
	bDoctor := struct{name string}{name: "Petros!"}
	anotherDoctor := &bDoctor
	fmt.Println(bDoctor)
	anotherDoctor.name = "New Petros"
	fmt.Println(bDoctor)


	//composition
	bird := Bird{}
	bird.name = "Petros"
	bird.origin = "Australia"
	bird.speedKPH = 100
	bird.canFly = false
	fmt.Println(bird,"\n")



	bird = Bird{
		Animal: Animal{name: "Petros", origin: "Greece"},
		speedKPH: 34,
		canFly: true,
	}

	fmt.Println(bird)

	//tag
	t:= reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}

type Animal struct {
	name string `required max:"100"`
	origin string
}

type Bird struct {
	Animal //we could specify a name, but we dont have to.
	speedKPH float32
	canFly bool
}


type Doctor struct {
	number int
	actorName string
	companions []string
}