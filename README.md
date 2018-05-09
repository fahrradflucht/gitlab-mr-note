# gitlab-mr-note

> CLI utility to post a merge request note from Gitlab CI

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

[Create a Gitlab access
token](https://docs.gitlab.com/ce/user/profile/personal_access_tokens.html)
for the user who should be used for making the notes. Then [configure it as a
CI environment
variable](https://docs.gitlab.com/ce/ci/variables/#secret-variables) named
`GITLAB_ACCESS_TOKEN`.

In CI:

```
$ wget https://github.com/fahrradflucht/gitlab-mr-note/releases/download/v1.0.0/gitlab-mr-note_linux_amd64 -O gitlab-mr-note
$ chmod +x gitlab-mr-note
```

Most likely `linux_amd64` will be the correct architecture but if your runner
has a different architecture get the correct link from the [release
page](https://github.com/fahrradflucht/gitlab-mr-note/releases).

## Usage
Just pipe your desired note text into it:

```
$ echo "Hallo Gitlab" | gitlab-mr-note
```

If the current CI jobs `HEAD` is the same as the `HEAD` of an open MR this
will post a note to the MR containing the piped in text wrapped in a code
block. If no MR which fulfils this condition is found the program exits with
a non-zero exit code.

This is as close as it gets for matching MR's and CI jobs until
[gitlab-org/gitlab-ce#23902](https://gitlab.com/gitlab-org/gitlab-ce/issues/23902)
and
[gitlab-org/gitlab-ce#15280](https://gitlab.com/gitlab-org/gitlab-ce/issues/15280)
are implemented.

## Maintainers

[@fahrradflucht](https://github.com/fahrradflucht)

## Contribute

PRs accepted.

## License

MIT © 2018 Mathis Wiehl
