package git

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os/exec"
	"sync"
)

type GitClient struct {
	username        string
	email           string
	token           string
	tagDestinations map[string]string
	mutex           *sync.Mutex
}

type Git interface {
	Status(dir string) error
	Add(path string, dir string) error
	Commit(message string, dir string) error
	Push(target string, dir string) error
	CheckOut(target string, dir string) error
	Clone(url string, destination string) error
	// SetTagDestination(tag string, destination string)
}

func NewGitClient(username string, email string, token string) Git {
	if username != "" && email != "" {
		if err := setUsername(username); err != nil {
			return nil
		}
		if err := setEmail(email); err != nil {
			return nil
		}
	}
	return &GitClient{
		username: username,
		email:    email,
		token:    token,
		mutex:    &sync.Mutex{},
		// tagDestinations: make(map[string]string),
	}
}

// func (g GitClient) SetTagDestination(tag string, destination string) {
// 	g.tagDestinations[tag] = destination
// }

func (g GitClient) Status(dir string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command("git", "status")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fail to run status")
	}
	fmt.Println(out.String())
	return nil
}

func (g GitClient) Add(path string, dir string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command("git", "add", path)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fail to run add")
	}
	fmt.Println(out.String())
	return nil
}

func (g GitClient) Commit(message string, dir string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command("git", "commit", "-m", message)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fail to run commit")
	}
	fmt.Println(out.String())
	return nil
}

func (g GitClient) Push(target string, dir string) error {
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
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fail to run push")
	}
	fmt.Println(out.String())
	return nil
}

func (g GitClient) CheckOut(target string, dir string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command("git", "checkout", "-b", target)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fail to run checkout")
	}
	fmt.Println(out.String())
	return nil
}

func (g GitClient) Clone(url string, destination string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	cmd := exec.Command(
		"git",
		"-c",
		g.createAuthHeader(),
		"clone",
		url,
		destination,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		println(err.Error())
		return fmt.Errorf("fail to run clone")
	}
	fmt.Println(out.String())
	return nil
}

func setUsername(username string) error {
	cmd := exec.Command("git", "config", "--global", "user.name", username)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fail to set git username")
	}
	fmt.Println(out.String())
	return nil
}

func setEmail(email string) error {
	cmd := exec.Command("git", "config", "--global", "user.email", email)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fail to set git email")
	}
	fmt.Println(out.String())
	return nil
}

func (g GitClient) createAuthHeader() string {
	token := fmt.Sprintf("%s:%s", g.username, g.token)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))
	header := fmt.Sprintf("Authorization: Basic %s", encodedToken)
	return fmt.Sprintf("http.extraHeader=%s", header)
}
