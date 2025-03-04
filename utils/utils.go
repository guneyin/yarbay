package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ilyakaznacheev/cleanenv"
	"go.opentelemetry.io/otel/trace"
)

func LoadConfig(cfg any) error {
	errFile := cleanenv.ReadConfig(".env", cfg)
	if errFile != nil {
		errEnv := cleanenv.ReadEnv(cfg)
		if errEnv != nil {
			return errEnv
		}
	}

	return nil
}

func Convert[T any](from any, to T) (T, error) {
	res, err := deepCopy(from, to)
	if err != nil {
		return to, err
	}

	if rt, ok := res.(T); ok {
		return rt, nil
	}

	return to, fmt.Errorf("cannot convert from %T to %T", from, to)
}

func deepCopy(src, dest any) (any, error) {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(src)
	if err != nil {
		return nil, err
	}
	err = gob.NewDecoder(&buf).Decode(dest)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func SpanFromContext(c *fiber.Ctx) trace.Span {
	return trace.SpanFromContext(c.Context())
}
