package user

import (
	"context"

	"github.com/mozart-rue/gobid/internal/validator"
)

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req SignInUserRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator
	eval.CheckField(validator.NotBlank(req.Email), "email", "must be provided")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRx), "email", "must be a valid email address")
	eval.CheckField(validator.NotBlank(req.Password), "password", "must be provided")
	return eval
}
