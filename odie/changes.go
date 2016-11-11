u := test_repo.UserModel{
	Id:   1,
	Name: "David Jones",
}

u.Commit()

u.Name = "David Bowie"

fmt.Printf("%+v\n", u.Changes(true))
=> map[name:[David Jones David Bowie]]