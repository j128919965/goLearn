package server

import "github.com/google/wire"

var PSet = wire.NewSet(NewHttpSerever)