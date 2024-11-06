package user

import (
	"context"

	"github.com/mozart-rue/gobid/internal/validator"
)

type CreateUserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.UserName), "user_name", "must be provided")
	eval.CheckField(validator.NotBlank(req.Email), "email", "must be provided")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRx), "email", "must be a valid email address")
	eval.CheckField(validator.NotBlank(req.Bio), "bio", "must be provided")
	eval.CheckField(validator.MinChars(req.Bio, 10) && validator.MaxChars(req.Bio, 255), "bio", "This fiels must be between 10 and 255 characters")
	eval.CheckField(validator.MinChars(req.Password, 8), "password", "must have at least 8 characters")

	return eval
}
