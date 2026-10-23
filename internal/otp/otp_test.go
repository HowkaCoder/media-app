package otp

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

func newTestStore(t *testing.T) *Store {
	t.Helper()
	mr := miniredis.RunT(t)
	return NewStore(redis.NewClient(&redis.Options{Addr: mr.Addr()}))
}

func TestGenerateCode(t *testing.T) {
	re := regexp.MustCompile(`^\d{6}$`)
	seen := map[string]bool{}
	for i := 0; i < 200; i++ {
		code, err := GenerateCode()
		if err != nil {
			t.Fatal(err)
		}
		if !re.MatchString(code) {
			t.Fatalf("bad code %q", code)
		}
		seen[code] = true
	}
	if len(seen) < 150 {
		t.Fatalf("codes are not random enough: %d unique of 200", len(seen))
	}
}

func TestIssueAndVerify(t *testing.T) {
	ctx := context.Background()
	s := newTestStore(t)

	code, err := s.Issue(ctx, "998901234567")
	if err != nil {
		t.Fatal(err)
	}

	ok, err := s.Verify(ctx, "998901234567", "000000x")
	if err != nil || ok {
		t.Fatalf("wrong code accepted: ok=%v err=%v", ok, err)
	}

	ok, err = s.Verify(ctx, "998901234567", code)
	if err != nil || !ok {
		t.Fatalf("valid code rejected: ok=%v err=%v", ok, err)
	}

	ok, _ = s.Verify(ctx, "998901234567", code)
	if ok {
		t.Fatal("code can be used twice")
	}
}

func TestIssueCooldown(t *testing.T) {
	ctx := context.Background()
	s := newTestStore(t)

	if _, err := s.Issue(ctx, "998901234567"); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Issue(ctx, "998901234567"); !errors.Is(err, ErrTooManyRequests) {
		t.Fatalf("expected ErrTooManyRequests, got %v", err)
	}
}

func TestVerifyAttemptsLimit(t *testing.T) {
	ctx := context.Background()
	s := newTestStore(t)

	code, err := s.Issue(ctx, "998901234567")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < maxAttempts; i++ {
		s.Verify(ctx, "998901234567", "wrong")
	}

	ok, _ := s.Verify(ctx, "998901234567", code)
	if ok {
		t.Fatal("code accepted after too many attempts")
	}
}
