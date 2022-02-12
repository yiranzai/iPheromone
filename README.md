# iPheromone 🏷

[![codecov](https://codecov.io/gh/yiranzai/iPheromone/branch/main/graph/badge.svg)](https://codecov.io/gh/yiranzai/iPheromone)
[![Go Report Card](https://goreportcard.com/badge/github.com/yiranzai/iPheromone)](https://goreportcard.com/report/github.com/yiranzai/iPheromone)
[![Sourcegraph](https://sourcegraph.com/github.com/yiranzai/iPheromone/-/badge.svg)](https://sourcegraph.com/github.com/yiranzai/iPheromone?badge)
[![Open Source Helpers](https://www.codetriage.com/yiranzai/iPheromone/badges/users.svg)](https://www.codetriage.com/yiranzai/iPheromone)
[![Release](https://img.shields.io/github/release/yiranzai/iPheromone.svg?style=flat-square)](https://github.com/yiranzai/iPheromone/releases)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2FiPheromone.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2FiPheromone?ref=badge_shield)

🐜 Wait for me to write.

## Table of contents 💿

______________________________________________________________________

<!--ts-->
   * [iPheromone <g-emoji class="g-emoji" alias="label" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f3f7.png">🏷</g-emoji>](#iPheromone-)
      * [Table of contents <g-emoji class="g-emoji" alias="cd" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f4bf.png">💿</g-emoji>](#table-of-contents-)
      * [Setup <g-emoji class="g-emoji" alias="electric_plug" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f50c.png">🔌</g-emoji>](#setup-)
         * [Setup by Command](#setup-by-command)
         * [Setup on GitHub](#setup-on-github)
      * [Usage <g-emoji class="g-emoji" alias="airplane" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/2708.png">✈️</g-emoji>](#usage-️)
         * [Install <g-emoji class="g-emoji" alias="anchor" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/2693.png">⚓️</g-emoji>](#install-️)
         * [Test <g-emoji class="g-emoji" alias="vertical_traffic_light" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f6a6.png">🚦</g-emoji>](#test-)
         * [<a href="https://pre-commit.com/" rel="nofollow">Pre-commit</a> <g-emoji class="g-emoji" alias="swimmer" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f3ca.png">🏊</g-emoji>](#pre-commit-)
      * [Recommend <g-emoji class="g-emoji" alias="star" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/2b50.png">⭐</g-emoji><g-emoji class="g-emoji" alias="star" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/2b50.png">⭐</g-emoji><g-emoji class="g-emoji" alias="star" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/2b50.png">⭐</g-emoji><g-emoji class="g-emoji" alias="star" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/2b50.png">⭐</g-emoji><g-emoji class="g-emoji" alias="star" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/2b50.png">⭐️</g-emoji>](#recommend-️)
         * [Packages <g-emoji class="g-emoji" alias="package" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f4e6.png">📦</g-emoji>](#packages-)
         * [Github Workflows <g-emoji class="g-emoji" alias="chains" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/26d3.png">⛓</g-emoji>](#github-workflows-)
            * [Golang Test And Coverage](#golang-test-and-coverage)
            * [<a href="https://github.com/pantheon-systems/autotag">Autotag</a>](#autotag)
            * [<a href="https://github.com/goreleaser/goreleaser-action">Goreleaser</a>](#goreleaser)
            * [<a href="https://github.com/crazy-max/ghaction-import-gpg">ghaction-import-gpg</a>](#ghaction-import-gpg)
            * [<a href="https://github.com/yiranzai/github-markdown-toc">Github Markdown TOC</a>](#github-markdown-toc)
      * [License](#license)
      * [base on <a href="https://github.com/lupguo/ddd-layout">lupguo/ddd-layout</a> <g-emoji class="g-emoji" alias="pray" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f64f.png">🙏</g-emoji><g-emoji class="g-emoji" alias="pray" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f64f.png">🙏</g-emoji><g-emoji class="g-emoji" alias="pray" fallback-src="https://github.githubassets.com/images/icons/emoji/unicode/1f64f.png">🙏</g-emoji>](#base-on-lupguoddd-layout-)
         * [DDD 布局](#ddd-布局)
            * [背景](#背景)
            * [目录说明](#目录说明)
            * [分层说明](#分层说明)
               * [领域层（业务核心）：](#领域层业务核心)
               * [应用层（流程编排）：](#应用层流程编排)
               * [接口层：](#接口层)
               * [基础层：](#基础层)

<!-- Added by: runner, at: Sat Feb 12 09:14:10 UTC 2022 -->

<!--te-->

______________________________________________________________________

## Setup 🔌

Something.

## Usage ✈️

Something.

### Install ⚓️

```sh
go get github.com/yiranzai/iPheromone
```

### Test 🚦

```sh
make test
```

### [Pre-commit](https://pre-commit.com/) 🏊

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

## Relations ⭐

### Packages 📦

- [github.com/spf13/cobra](https://github.com/spf13/cobra) powerful modern CLI
- [github.com/spf13/viper](https://github.com/spf13/viper) Go configuration with fangs
- [github.com/stretchr/testify](https://github.com/stretchr/testify) Go code (golang) set of packages that provide many tools for testifying that your code will behave as you intend.

### Github Workflows ⛓

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
