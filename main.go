package main

func main() {
	todos := Todos{}

	storage := NewStorage[Todos]("dk-todos.json")
	storage.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	// todos.print()

	storage.Save(todos)
}
