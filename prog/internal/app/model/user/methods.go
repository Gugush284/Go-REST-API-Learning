package model_user

func (s *User) Add_User(id int, login string, password string) {
	s.login = login
	s.password = password
	s.id = id
}
