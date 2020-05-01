package pork

import "testing"

func TestCloneRepository(t *testing.T) {
  if err := CloneRepository("davidlares/testrepo", "master", false); err != nil {
    t.Fail()
  }
}
