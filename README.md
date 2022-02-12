# 🐜 iPheromone

[![codecov](https://codecov.io/gh/yiranzai/iPheromone/branch/main/graph/badge.svg)](https://codecov.io/gh/yiranzai/iPheromone)
[![Go Report Card](https://goreportcard.com/badge/github.com/yiranzai/iPheromone)](https://goreportcard.com/report/github.com/yiranzai/iPheromone)
[![Sourcegraph](https://sourcegraph.com/github.com/yiranzai/iPheromone/-/badge.svg)](https://sourcegraph.com/github.com/yiranzai/iPheromone?badge)
[![Open Source Helpers](https://www.codetriage.com/yiranzai/ipheromone/badges/users.svg)](https://www.codetriage.com/yiranzai/ipheromone)
[![Release](https://img.shields.io/github/release/yiranzai/iPheromone.svg?style=flat-square)](https://github.com/yiranzai/iPheromone/releases)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2FiPheromone.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2FiPheromone?ref=badge_shield)

🐜 Wait for me to write.

## 💿 Table of contents

______________________________________________________________________

<!--ts-->
   * [iPheromone](#-ipheromone)
      * [Table of contents](#-table-of-contents)
      * [Setup](#-setup)
      * [Usage](#️-usage)
         * [Install](#-install)
         * [Test](#-test)
         * [<a href="https://pre-commit.com/" rel="nofollow">Pre-commit</a>](#-pre-commit)
      * [Relations](#-relations)
         * [Packages](#-packages)
         * [Github Workflows](#-github-workflows)
            * [Golang Test And Coverage](#golang-test-and-coverage)
            * [<a href="https://github.com/pantheon-systems/autotag">Autotag</a>](#autotag)
            * [<a href="https://github.com/goreleaser/goreleaser-action">Goreleaser</a>](#goreleaser)
            * [<a href="https://github.com/crazy-max/ghaction-import-gpg">ghaction-import-gpg</a>](#ghaction-import-gpg)
            * [<a href="https://github.com/yiranzai/github-markdown-toc">Github Markdown TOC</a>](#github-markdown-toc)
      * [License](#license)

<!-- Added by: runner, at: Sat Feb 12 13:15:53 UTC 2022 -->

<!--te-->

______________________________________________________________________

## 🔌 Setup

Something.

## ✈️ Usage

Something.

### ⚓ Install

```sh
go get github.com/yiranzai/iPheromone
```

### 🚦 Test

```sh
make test
```

### 🏊 [Pre-commit](https://pre-commit.com/)

check or fix code style.

see [.pre-commit-config.yaml](.pre-commit-config.yaml)

e.g:

- go fmt
- golines
- go mod tiny
- go vet

Install the [pre-commit](https://pre-commit.com/)

```sh
pip install pre-commit
pre-commit install
vim .pre-commit-config.yaml
```

## ⭐ Relations

### 📦 Packages

- [github.com/spf13/cobra](https://github.com/spf13/cobra) powerful modern CLI
- [github.com/spf13/viper](https://github.com/spf13/viper) Go configuration with fangs
- [github.com/stretchr/testify](https://github.com/stretchr/testify) Go code (golang) set of packages that provide many tools for testifying that your code will behave as you intend.

### ⛓ Github Workflows

This repo used some workflows

#### Golang Test And Coverage

Golang Test And Coverage upload to [Codecov](https://codecov.io)

#### [Autotag](https://github.com/pantheon-systems/autotag)

Automatically increment version tags to a git repo based on commit messages.

#### [Goreleaser](https://github.com/goreleaser/goreleaser-action)

GitHub Action for GoReleaser

#### [ghaction-import-gpg](https://github.com/crazy-max/ghaction-import-gpg)

GitHub Action to easily import a GPG key.

[New Repository secret](https://github.com/yiranzai/iPheromone/settings/secrets/actions/new)

add `YOUR_PRIVATE_KEY` and `PASSPHRASE` secrets.

#### [Github Markdown TOC](https://github.com/yiranzai/github-markdown-toc)

This [Github Markdown TOC](https://github.com/yiranzai/github-markdown-toc) fork for [@ekalinin](https://github.com/ekalinin)'s [Github Markdown TOC](https://github.com/ekalinin/github-markdown-toc).

I Added flags to support for more features.

See [ekalinin/github-markdown-toc#110](https://github.com/ekalinin/github-markdown-toc/issues/110) and [ekalinin/github-markdown-toc#115](https://github.com/ekalinin/github-markdown-toc/pull/115)

```ini
--all Find all Markdown files for non-hidden folders
--auto Ignore ts/te tags, Automatically at the end/head of the file
--head The TOC is generated in the header of the file, requires --auto
```

## License

This project is licensed under the MIT License.
See the [LICENSE](./LICENSE) file
for the full license text.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2FiPheromone.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2FiPheromone?ref=badge_large)
