package tester

import (
	"context"
	"io"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

const toCatchValue = "37355060-8ef2-47c3-833d-c78c3663cd66"
const toCatchValueRes = "ssh: handshake failed: " + toCatchValue

type toCatch struct{}
type NotCatched struct{}

func (e toCatch) Error() string    { return toCatchValue }
func (e NotCatched) Error() string { return "HostKeyCallback not called" }

func Test(ctx context.Context, addr string) (handshake time.Duration, readkey time.Duration, err error) {
	xz := &xzTester{}
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(xz),
		},
		HostKeyCallback: xz.HostKeyCallback,
	}

	start := time.Now()
	conn, err := net.DialTimeout("tcp", addr, config.Timeout)
	handshake = time.Since(start)

	if err != nil {
		return handshake, readkey, err
	}

	start = time.Now()
	_, _, _, err = ssh.NewClientConn(conn, addr, config)
	readkey = time.Since(start)

	if err == nil {
		return handshake, readkey, NotCatched{}
	}
	if err.Error() != toCatchValueRes {
		return handshake, readkey, err
	}
	return handshake, readkey, nil
}

type xzTester struct{}

func (s *xzTester) PublicKey() ssh.PublicKey {
	return nil // will never be called
}

func (s *xzTester) Sign(rand io.Reader, data []byte) (*ssh.Signature, error) {
	return nil, nil // will never be called
}

func (s *xzTester) HostKeyCallback(_ string, _ net.Addr, key ssh.PublicKey) error {
	return toCatch{}
}
