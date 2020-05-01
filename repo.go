package pork

import (
  git "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/config"
  "gopkg.in/src-d/go-git.v4/plumbing"
  "strings"
  "path/filepath"
  "fmt"
)

type GHRepo struct {
  RepoDir string
  owner string
  project string
  repo *git.Repository
}

func NewGHRepo(repository string) (*GHRepo, error) {
  values := strings.Split(repository, "/")
  if len(values) != 2 {
    return nil, fmt.Errorf("repository must have owner/project format")
  }
  return &GHRepo {
    owner: values[0],
    project: values[1],
  }, nil
}

// this function will return the GH repository endpoint
func (g *GHRepo) RepositoryURL() string {
    return fmt.Sprintf("https://github.com/%s/%s.git", g.owner, g.project)
}

// expects a local disk destination
func (g *GHRepo) Clone(dest string) error {
  fullPath := filepath.Join(dest, fmt.Sprintf("%s-%s", g.owner, g.project))
  // provide the path - passing clone options
  repo, err := git.PlainClone(fullPath, false, &git.CloneOptions{
      URL: g.RepositoryURL(),
  })
  if err != nil {
    return err
  }
  g.repo = repo
  g.RepoDir = fullPath
  return nil
}

// adding checkout
func (g *GHRepo) Checkout(ref string, create bool) error {
  // setting arg options
  opts := &git.CheckoutOptions{
      Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", ref)),
      Create: create,
  }

  if create {
    head, err := g.repo.Head()
    if err != nil {
      return err
    }
    opts.Hash = head.Hash()
  }

  tree, err := g.repo.Worktree()
  if err != nil {
    return err
  }
  return tree.Checkout(opts)
}

// remotes
func(g *GHRepo) AddUpStream(repository *GHRepo) error {
  // upstream and local
   _, err := g.repo.CreateRemote(&config.RemoteConfig{
    Name: "upstream",
    URLs: []string {repository.RepositoryURL()},
  })
  return err
}
