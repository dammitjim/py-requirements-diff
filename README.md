# Requirements diff
Quick and dirty script to compare a git diff of a requirements.txt file for including in changelogs.

Could this have been done using fancy shell commands? Yeah probably.

# Install

* go install github.com/dammitjim/py-requirements-diff

# Usage
```bash
git diff {version} requirements/common.txt | py-requirements-diff -prefix="* "

// i want it sorted 
git diff {version} requirements/common.txt | py-requirements-diff -prefix="* " | sort

// i have a file instead 
py-requirements-diff -prefix="* " diff.txt | sort
```

## Example output

```bash
⟩ git diff 2.41.0 requirements/common.txt | py-requirements-diff | sort
Changed: amqp == 5.1.1  -> 5.2.0
Changed: billiard == 4.1.0  -> 4.2.0
Changed: cattrs == 23.1.2  -> 23.2.3
Changed: celery == 5.3.4  -> 5.3.6
Changed: certifi == 2023.7.22  -> 2023.11.17
Changed: charset-normalizer == 3.3.0  -> 3.3.2
Changed: django == 3.2.22  -> 3.2.23
Changed: django-cache-url == 3.4.4  -> 3.4.5
Changed: envier == 0.4.0  -> 0.5.0
Changed: freezegun == 1.2.2  -> 1.3.1
Changed: idna == 3.4  -> 3.6
Changed: importlib-metadata == 6.8.0  -> 6.11.0
Changed: kombu == 5.3.2  -> 5.3.4
Changed: opentelemetry-api == 1.20.0  -> 1.21.0
Changed: packaging == 21.3  -> 23.2
Changed: platformdirs == 3.11.0  -> 4.1.0
Changed: prompt-toolkit == 3.0.39  -> 3.0.43
Changed: protobuf == 4.24.4  -> 4.25.1
Changed: pypdf == 3.16.4  -> 3.17.2
Changed: setuptools == 68.2.2  -> 69.0.2
Changed: typing-extensions == 4.8.0  -> 4.9.0
Changed: urllib3 == 1.26.17  -> 1.26.18
Changed: vine == 5.0.0  -> 5.1.0
Changed: wcwidth == 0.2.8  -> 0.2.12
Changed: wrapt == 1.15.0  -> 1.16.0
Removed: pyparsing == 3.1.1
```
