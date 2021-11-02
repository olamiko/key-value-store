package server

import (
	"github.com/olamiko/key-value-store/utils"
)

type Listener int

type Server struct {
	log     string
	store   utils.Store
	address string
}
