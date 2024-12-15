package idgen

import (
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/jonboulle/clockwork"
	"strings"
)

type UseCase int

const (
	User UseCase = 0
)

var prefixes = map[UseCase]string{
	User: "USER",
}

type IDGenerator interface {
	Get(u UseCase) (string, error)
}

type UseCaseIDGenerator struct {
	clk clockwork.Clock
}

func (u *UseCaseIDGenerator) Get(useCase UseCase) (string, error) {
	id, err := uuid.New().MarshalBinary()
	if err != nil {
		return "", err
	}
	return prefixes[useCase] + u.encodeToStringWithDate(id), nil
}

func (u *UseCaseIDGenerator) encodeToStringWithDate(b []byte) string {
	shortTime := u.clk.Now().Format("060102")

	base64Opt := base64.StdEncoding.EncodeToString(b)
	paddingSize := strings.Count(base64Opt, string(base64.StdPadding))

	if paddingSize == 0 {
		return base64Opt + shortTime
	}

	trimmedOpt := strings.TrimRight(base64Opt, string(base64.StdPadding))

	return trimmedOpt + shortTime + strings.Repeat(string(base64.StdPadding), paddingSize)
}
