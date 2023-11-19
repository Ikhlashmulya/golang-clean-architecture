package test

import (
	"testing"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	v := config.NewConfig()
	assert.Equal(t, "Golang RESTful API", v.GetString("app.name"))
}
