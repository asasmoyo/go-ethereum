package me

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
)

var BlocksFilePath string
var blockFile *os.File

func prepareBlockWrite() error {
	if BlocksFilePath == "" {
		Log("BlocksFilePath is empty")
		return errors.New("failed preparing block write")
	}

	if _, err := os.Stat(BlocksFilePath); os.IsNotExist(err) {
		os.Create(BlocksFilePath)
	}

	var err error
	blockFile, err = os.OpenFile(BlocksFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		Log("Failed opening file for block write", "err", err)
		return err
	}

	return nil
}

func WriteNewBlock(block *types.Block) {
	Log("Writing block", "number", block.Number(), "hash", block.Hash())
	if err := prepareBlockWrite(); err != nil {
		return
	}

	j, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"transactions": block.Body().Transactions,
			"uncles":       block.Body().Uncles,
		},
		"header": block.Header(),
	})
	if err != nil {
		Log("Failed converting block to json", "err", err)
		return
	}

	_, err = blockFile.WriteString(string(j) + "\n")
	if err != nil {
		Log("Failed writing block", "err", err)
		return
	}
	Log("Write OK!")
}
