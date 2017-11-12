package me

import "github.com/ethereum/go-ethereum/log"

func Log(message string, context ...interface{}) {
	log.Info("[ME]: "+message, context...)
}
