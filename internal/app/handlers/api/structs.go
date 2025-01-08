package api_handler

import "likemind/internal/domain"

type SignInRequst struct {
	Email    string `in:"form=email;nonzero"`
	Password string `in:"form=password;nonzero"`
}

type SignUpRequest struct {
	Name     string `in:"form=name;nonzero"`
	Surname  string `in:"form=surname;nonzero"`
	Nickname string `in:"form=nickname;nonzero"`
	Email    string `in:"form=email;nonzero"`
	Password string `in:"form=password;nonzero"`
}

func (r *SignUpRequest) getCreateUserRequest() domain.User {
	return domain.User{
		Name:     r.Name,
		Surname:  r.Surname,
		Nickname: r.Nickname,
	}
}
