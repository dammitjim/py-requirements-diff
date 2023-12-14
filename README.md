# Requirements diff
Quick and dirty script to compare a git diff of a requirements.txt file for including in changelogs.

Could this have been done using fancy shell commands? Yeah probably.

# Usage
```
```bash
git diff {version} requirements/common.txt | go run main.go -prefix="* "

// but i want it sorted 
git diff {version} requirements/common.txt | go run main.go -prefix="* " | sort
```
