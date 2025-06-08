NAME=$1
LOWER=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')

cat <<EOF >validation/${LOWER}_validation.go
package validation

import (
  "context"
  validation "github.com/go-ozzo/ozzo-validation/v4"
  "github.com/keigo-saito0602/go_learning_for_poke_api/domain/model"
)

type ${NAME}Validator struct{}

func New${NAME}Validator() *${NAME}Validator {
  return &${NAME}Validator{}
}

func (v *${NAME}Validator) ValidateCreate(ctx context.Context, m *model.${NAME}) error {
  return validation.ValidateStruct(m,
    validation.Field(&m.UserID, validation.Required),
    validation.Field(&m.Content, validation.Required),
  )
}
EOF

echo "✅ validation/${LOWER}_validation.go created"
