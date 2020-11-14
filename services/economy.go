package services

import "github.com/MruV-RP/mruv-api-placeholder/generator"

type EconomyServer struct {
	gen generator.IGenerator
}

func NewEconomyServer(gen generator.IGenerator) *EconomyServer {
	return &EconomyServer{gen: gen}
}
