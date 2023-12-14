# Requirements diff
Quick and dirty script to compare a git diff of a requirements.txt file for including in changelogs.

Could this have been done using fancy shell commands? Yeah probably.

# Install

* go install github.com/dammitjim/py-requirements-diff

# Usage
```bash
git diff {version} requirements/common.txt | py-requirements-diff -prefix="* "

// but i want it sorted 
git diff {version} requirements/common.txt | py-requirements-diff -prefix="* " | sort
```
