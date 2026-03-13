package service

func (s *Service) Login(username, password string) (string, error) {

	user, err := s.storage.GetByUsername(username)
	if err != nil {
		return "", err
	}

	jwt, err := s.authentication.ValidateUser(user, password)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
