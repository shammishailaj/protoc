//+build !gogit

package main

import (
	"fmt"
	"log"
)

type gitRepo struct {
	url string
	dir string
}

func gitCmd(args ...string) error {
	code := execute("git", args...)
	if code != 0 {
		return fmt.Errorf("git failed: exit code %d", code)
	}
	return nil
}

func gitOpenDir(url, dir string) (repo, error) {
	err := gitCmd("-C", dir, "rev-parse")
	return &gitRepo{url: url, dir: dir}, err
}

func gitCloneDir(url, dir string) (repo, error) {
	err := gitCmd("clone", "https://"+url, dir)
	return &gitRepo{url: url, dir: dir}, err
}

func (r *gitRepo) Checkout(rev string) error {
	if err := gitCmd("-C", r.dir, "checkout", "master"); err != nil {
		return err
	}
	if rev == "" || rev == latestRev {
		if rev == "" {
			if err := gitCmd("-C", r.dir, "pull"); err != nil {
				// Ignore if pull fails, try our best to work offline
				log.Println("Git pull failed:", err)
			}
		}
		rev = "HEAD"
	}
	return gitCmd("-C", r.dir, "checkout", "-q", rev)
}