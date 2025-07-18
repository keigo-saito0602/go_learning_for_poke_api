NAME=$1
LOWER=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')

cat <<EOF >interface/handler/${LOWER}_handler.go
package handler

import (
  "net/http"

  "github.com/keigo-saito0602/go_learning_for_poke_api/domain/model"
  "github.com/keigo-saito0602/go_learning_for_poke_api/usecase"
  "github.com/labstack/echo/v4"
)

type ${NAME}Handler struct {
  usecase usecase.${NAME}Usecase
}

func New${NAME}Handler(u usecase.${NAME}Usecase) *${NAME}Handler {
  return &${NAME}Handler{usecase: u}
}

// List${NAME}s, Get${NAME}, etc...
EOF

echo "✅ handler/${LOWER}_handler.go created"
