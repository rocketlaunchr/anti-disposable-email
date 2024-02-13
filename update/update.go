// Copyright 2020-24 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package update

import (
	"bufio"
	"context"
	"sync"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Update can be used to update the list of disposable email domains.
// It uses the regularly updated list found here: https://github.com/martenson/disposable-email-domains.
func Update(ctx context.Context, list *map[string]struct{}, lock ...sync.Locker) error {

	fs := memfs.New()

	opts := &git.CloneOptions{
		URL:   "https://github.com/disposable-email-domains/disposable-email-domains",
		Depth: 0,
	}

	_, err := git.CloneContext(ctx, memory.NewStorage(), fs, opts)
	if err != nil {
		return err
	}

	file, err := fs.Open("disposable_email_blocklist.conf")
	if err != nil {
		return err
	}

	newList := make(map[string]struct{}, 3500)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newList[scanner.Text()] = struct{}{}
	}

	err = file.Close()
	if err != nil {
		return err
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	if len(lock) > 0 && lock[0] != nil {
		lock[0].Lock()
		defer lock[0].Unlock()
	}

	*list = newList

	return nil
}
