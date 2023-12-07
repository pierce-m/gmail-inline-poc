package main

import (
	"errors"
	"io"
	"log"
	"time"

	gosmtp "github.com/emersion/go-smtp"
	"github.com/jordan-wright/email"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// NewSession is called after client greeting (EHLO, HELO).
func (bkd *Backend) NewSession(c *gosmtp.Conn) (gosmtp.Session, error) {
	return &Session{}, nil
}

// A Session is returned after successful login.
type Session struct {
	Client *gosmtp.Client
}

// AuthPlain implements authentication using SASL PLAIN.
func (s *Session) AuthPlain(username, password string) error {
	if username != "username" || password != "password" {
		return errors.New("Invalid username or password")
	}
	return nil
}

func (s *Session) Mail(from string, opts *gosmtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string, opts *gosmtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	e, err := email.NewEmailFromReader(r)
	if err != nil {
		log.Fatal("error reading email", err)
		return err
	}

	e.Headers.Add("X-NIGHTFALL-SCANNED", "")
	err = e.Send("smtp-relay.gmail.com:587", nil)
	if err != nil {
		log.Fatal("error sending email", err)
		return err
	}

	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func main() {
	be := &Backend{}

	s := gosmtp.NewServer(be)

	s.Addr = "localhost:1025"
	s.Domain = "localhost"
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
