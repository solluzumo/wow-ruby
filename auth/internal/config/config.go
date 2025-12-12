package config

import (
	"os"
	"strconv"

	"github.com/solluzumo/wow-ruby/pkg"
)

type Config struct {
	ArgonParams      *pkg.Argon2Params
	TaskChanCap      int
	HashWorkersCount int
}

var DefaultParams = &pkg.Argon2Params{
	Memory:      16 * 1024,
	Iterations:  1,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   16,
}

func NewConfig() (*Config, error) {
	taskChanCap, err := strconv.Atoi(os.Getenv("TASK_CHANNEL_CAP"))
	if err != nil {
		return nil, err
	}
	hashWorkersCount, err := strconv.Atoi(os.Getenv("HASH_WORKERS_COUNT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		ArgonParams:      DefaultParams,
		TaskChanCap:      taskChanCap,
		HashWorkersCount: hashWorkersCount,
	}, nil
}
