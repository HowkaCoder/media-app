package otp

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/go-redis/redis/v8"
)

const maxAttempts = 5

var ErrTooManyRequests = errors.New("code was sent recently, try again later")

type Store struct {
	rdb      *redis.Client
	ttl      time.Duration
	cooldown time.Duration
}

func NewStore(rdb *redis.Client) *Store {
	return &Store{rdb: rdb, ttl: 5 * time.Minute, cooldown: time.Minute}
}

func GenerateCode() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}

func (s *Store) Issue(ctx context.Context, phone string) (string, error) {
	ok, err := s.rdb.SetNX(ctx, "otp:cooldown:"+phone, 1, s.cooldown).Result()
	if err != nil {
		return "", err
	}
	if !ok {
		return "", ErrTooManyRequests
	}

	code, err := GenerateCode()
	if err != nil {
		return "", err
	}

	if err := s.rdb.Set(ctx, "otp:code:"+phone, code, s.ttl).Err(); err != nil {
		return "", err
	}
	if err := s.rdb.Del(ctx, "otp:attempts:"+phone).Err(); err != nil {
		return "", err
	}
	return code, nil
}

func (s *Store) Verify(ctx context.Context, phone, code string) (bool, error) {
	key := "otp:code:" + phone
	attemptsKey := "otp:attempts:" + phone

	attempts, err := s.rdb.Incr(ctx, attemptsKey).Result()
	if err != nil {
		return false, err
	}
	if attempts == 1 {
		s.rdb.Expire(ctx, attemptsKey, s.ttl)
	}
	if attempts > maxAttempts {
		s.rdb.Del(ctx, key)
		return false, nil
	}

	stored, err := s.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if stored != code {
		return false, nil
	}
	return true, s.rdb.Del(ctx, key, attemptsKey).Err()
}
