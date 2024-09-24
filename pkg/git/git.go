package git

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"os/exec"
	"sync"

	"go.uber.org/zap"
)

type GitClient struct {
	logger      *zap.Logger
	username    string
	email       string
	token       string
	destination string
	mutex       *sync.Mutex
}

type Git interface {
	Status() error
	Add(path string) error
	Commit(message string) error
	Push(target string) error
	CheckOut(target string) error
	Clone(url string) error
}

func NewGitClient(logger *zap.Logger, username string, email string, destination string, token string) Git {
	if username != "" && email != "" {
		if err := setUsername(logger, username); err != nil {
			logger.Warn("fail to set git username", zap.Error(err))
			return nil
		}
		if err := setEmail(logger, email); err != nil {
			logger.Warn("fail to set git set email", zap.Error(err))
			return nil
		}
	}
	return &GitClient{
		username:    username,
		email:       email,
		destination: destination,
		token:       token,
		mutex:       &sync.Mutex{},
	}
}

func (g GitClient) Status() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command("git", "status")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = g.destination
	err := cmd.Run()
	if err != nil {
		g.logger.Error("fail to run git status", zap.Error(err))
		return errors.New("fail to run git status")
	}
	return nil
}

func (g GitClient) Add(path string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command("git", "add", path)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = g.destination
	err := cmd.Run()
	if err != nil {
		g.logger.Error("fail to run git add", zap.Error(err))
		return errors.New("fail to run git add")
	}
	return nil
}

func (g GitClient) Commit(message string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command("git", "commit", "-m", message)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = g.destination
	err := cmd.Run()
	if err != nil {
		g.logger.Error("fail to run git commit", zap.Error(err))
		return errors.New("fail to run git commit")
	}
	return nil
}

func (g GitClient) Push(target string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command(
		"git",
		"-c",
		g.createAuthHeader(),
		"push",
		"origin",
		target,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = g.destination
	err := cmd.Run()
	if err != nil {
		g.logger.Error("fail to git push", zap.Error(err))
		return errors.New("fail to run git push")
	}
	return nil
}

func (g GitClient) CheckOut(target string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command("git", "checkout", "-b", target)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = g.destination
	err := cmd.Run()
	if err != nil {
		g.logger.Error("fail to run git checkout", zap.Error(err))
		return errors.New("fail to run git checkout")
	}
	return nil
}

func (g GitClient) Clone(url string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command(
		"git",
		"-c",
		g.createAuthHeader(),
		"clone",
		url,
		g.destination,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		g.logger.Error("fail to run git clone", zap.Error(err))
		return errors.New("fail to run git clone")
	}
	return nil
}

func setUsername(logger *zap.Logger, username string) error {
	cmd := exec.Command("git", "config", "--global", "user.name", username)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		logger.Error("fail to set git username", zap.Error(err))
		return errors.New("fail to set git username")
	}
	return nil
}

func setEmail(logger *zap.Logger, email string) error {
	cmd := exec.Command("git", "config", "--global", "user.email", email)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		logger.Error("fail to set git email", zap.Error(err))
		return errors.New("fail to set git email")
	}
	return nil
}

func (g GitClient) createAuthHeader() string {
	token := fmt.Sprintf("%s:%s", g.username, g.token)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))
	header := fmt.Sprintf("Authorization: Basic %s", encodedToken)
	return fmt.Sprintf("http.extraHeader=%s", header)
}
