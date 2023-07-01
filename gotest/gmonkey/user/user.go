package user


type User struct  {
	Name    string
	Age     int
}

type UserSer struct {

}

func (s *UserSer) GetUserINfo(name string) (*User, error) {
	// do something

	user := &User{
		Name: "test",
		Age:  0,
	}

	return user, nil
}