package main

import (
	"flag"
	"fmt"
	"log"
)

func seedDog(store Storage, name string, path string) *Dog {
	dog := NewDog(name, path)
	if err := store.AddDog(dog); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new dog => ", dog.Name)

	return dog
}

func seedDogs(s Storage) {
	seedDog(s, "affenpinscher", "./ASSETS/affenpinscher.jpg")
	seedDog(s, "dachshund", "./ASSETS/dachshund.jpg")
	seedDog(s, "pitbull", "./ASSETS/pitbull.jpg")
	seedDog(s, "labrador", "./ASSETS/labrador.jpg")
	seedDog(s, "african", "./ASSETS/african.jpg")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	if *seed {
		fmt.Println("seeding the database")
		seedDogs(store)
	}

	server := NEWAPIServer("8080", store)
	server.Run()
}
