package model_user

func (s *User) Add_User(login string, password string) {
	s.login = login
	s.password = password
}

func (s *User) Add_id(id int) {
	s.id = id
}

func (s *User) GetLogin() string {
	return s.login
}

func (s *User) GetPassword() string {
	return s.password
}

func (s *User) Getid() int {
	return s.id
}
