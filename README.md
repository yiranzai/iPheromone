# golang-project-template

[![codecov](https://codecov.io/gh/yiranzai/golang-project-template/branch/main/graph/badge.svg)](https://codecov.io/gh/yiranzai/golang-project-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/yiranzai/golang-project-template)](https://goreportcard.com/report/github.com/yiranzai/golang-project-template)
[![Sourcegraph](https://sourcegraph.com/github.com/yiranzai/golang-project-template/-/badge.svg)](https://sourcegraph.com/github.com/yiranzai/golang-project-template?badge)
[![Open Source Helpers](https://www.codetriage.com/yiranzai/golang-project-template/badges/users.svg)](https://www.codetriage.com/yiranzai/golang-project-template)
[![Release](https://img.shields.io/github/release/yiranzai/golang-project-template.svg?style=flat-square)](https://github.com/yiranzai/golang-project-template/releases)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2Fgolang-project-template.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2Fgolang-project-template?ref=badge_shield)

Golang project general template for me.

## Table of contents

______________________________________________________________________

<!--ts-->
* [golang-project-template](#golang-project-template)
  * [目录](#目录)
  * [Setup](#setup)
    * [Setup by Command](#setup-by-command)
    * [Setup on GitHub](#setup-on-github)
  * [Usage](#usage)
    * [Install](#install)
    * [Test](#test)
    * [<a href="https://pre-commit.com/" rel="nofollow">Pre-commit</a>](#pre-commit)
  * [Github Workflows](#github-workflows)
    * [Golang Test And Coverage](#golang-test-and-coverage)
    * [<a href="https://github.com/pantheon-systems/autotag">Autotag</a>](#autotag)
    * [<a href="https://github.com/goreleaser/goreleaser-action">Goreleaser</a>](#goreleaser)
    * [<a href="https://github.com/yiranzai/github-markdown-toc">Github Markdown TOC</a>](#github-markdown-toc)
  * [License](#license)

<!-- Added by: runner, at: Sat Apr 10 09:08:32 UTC 2021 -->

<!--te-->

______________________________________________________________________

## Setup

### Setup by Command

1. `git clone https://github.com/yiranzai/golang-project-template your_awesome_project`
1. Replace all strings `yiranzai/golang-project-template` in this repository to `your_github_username/your_awesome_project`
1. Replace all strings `golang-project-template` in this repository to `your_awesome_project`

### Setup on GitHub

Click "Use this template" button on GitHub project page.

## Usage

Something.

### Install

```sh
go get github.com/yiranzai/golang-project-template
```

### Test

```sh
go get gotest.tools/v3
```

### [Pre-commit](https://pre-commit.com/)

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

## Github Workflows

This repo used some workflows

### Golang Test And Coverage

Golang Test And Coverage upload to [Codecov](https://codecov.io)

### [Autotag](https://github.com/pantheon-systems/autotag)

Automatically increment version tags to a git repo based on commit messages.

### [Goreleaser](https://github.com/goreleaser/goreleaser-action)

GitHub Action for GoReleaser

### [ghaction-import-gpg](https://github.com/crazy-max/ghaction-import-gpg)

GitHub Action to easily import a GPG key.

[New Repository secret](https://github.com/yiranzai/golang-project-template/settings/secrets/actions/new)

add `YOUR_PRIVATE_KEY` and `PASSPHRASE` secrets.

### [Github Markdown TOC](https://github.com/yiranzai/github-markdown-toc)

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

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2Fgolang-project-template.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2Fgolang-project-template?ref=badge_large)


---

## base on [lupguo/ddd-layout](https://github.com/lupguo/ddd-layout)

### DDD布局

#### 背景

最近一年一直在尝试DDD落地，也在重新思考DDD的的布局如何更合理，结合最近项目落地后，重新构思了一版DDD的目录结构，后续会结合项目以及DDD理解，还会持续不断演进

相关文章参考: https://tkstorm.com/posts-list/software-engineering/cloud-native/ddd-think/

#### 目录说明

```
.
├── Makefile
├── README.md
├── app
│    ├── application    # 应用层，内部文件以`_app.go`结尾
│    │    ├── image_search_app.go
│    │    └── image_upload_app.go
│    ├── domain         # 领域层，依赖仓储接口
│    │    ├── entity        # 领域实体，以`_entity.go`结尾
│    │    ├── repository    # 仓储接口，以`_repos.go`结尾
│    │    ├── service       # 领域服务，以`_service.go`结尾
│    │    └── valobj        # 值对象，以`_valobj.go`结尾
│    ├── infrastructure # 基础上设施层，实现仓储接口
│    │    ├── dbs
│    │    ├── esearch
│    │    ├── mqs
│    │    └── rds
│    └── interfaces     # 接口层，依赖application层
│        ├── search_intf.go
│        ├── tag_intf.go
│        └── upload_intf.go
├── build               # 多操作系统，编译文件生成
├── cmd                 
│    └── myapp          # 目录文件夹为应用名称，期间仅包含一个main.go文件
├── configs             # 服务配置文件，包括配置中心、错误码等信息
│    ├── apollo
│    ├── confd
│    └── errcode
├── deployments         # CI/CD持续部署相关的一些脚本
├── docs                # 项目相关文档
├── go.mod
├── internel            # 从项目内抽离的包，可能被复用到其他项目中去的包
│    └── reusable
└── test
    └── testdata
```

#### 分层说明

##### 领域层（业务核心）：

- 领域层是业务核心，按先分析，然后设计，最后才是开发，这样才能分析核心业务
- DDD 反模式，领域划分不清楚，导致领域内有多个上下文相关信息，代码必然会混乱
- 仓储接口（依赖接口编程，而非实现编程）：
    - 业务领域代码会有存储需求，通过依赖注入方式，解耦对基础设施层的依赖
- 实体、值对象、聚合、聚合根：
    - 实体可以简单理解为领域的核心对象，有属性、操作方法（可以理解为 OOP 中的实体），比如登录领域（实体账号，有属性还有登录方法）、商品领域（商品实体）、订单领域（用户、购物车、商品、订单）
    - 实体采用充血模型实现所有与之相关的业务，若单一实体或值对象无法实现领域中的功能，可以借助**领域服务**组合多个实体（或值对象）实现复杂的业务逻辑（比如下单服务，涉及用户、商品、订单、购物车实体交互）
- 值对象：
    - 领域中的特殊对象，把一组相关属性组合在一起的对象，可以简单理解为一组属性的集合，比如用户收货地址，可以打包成一个值对象；
    - 值对象可以嵌入到实体属性中，比如收货地址值对象嵌入到用户收货属性中
- 聚合（领域聚合服务）：
    - 在一个限定上下文中（在某个领域内），紧密相关的实体、值对象进行组合，就构成了聚合，有更强的表现
    - 注意：如果是两个完全不同领域的信息聚合，应该通过应用服务来组合
- 聚合根（根实体）：
    - 聚合根也是称根实体，是聚合中的主负责人，作为聚合管理者，对外统一（比如商品详情页的商品聚合根）
    - 聚合根说白了就是在一组聚合中选出一个对外代表领域的根实体

##### 应用层（流程编排）：

- 应用层做哪些内容：
    - 做多个领域服务的聚合（课详内容：包含课程资料、课程类型、课程价格、课程销量，一定会存在业务聚合的情况，这块应用层做这个事情的）
    - 做业务流程的编排处理，说明清楚业务 1234，但实际不应该有逻辑，逻辑在领域内（比如发布课程：课程资产入库、图片上传、老师人员、客服人员安排等）
    - 事件通知、发布订阅、权限校验、安全认证等（比如上课提醒消息）
- 反模式：在应用层写业务逻辑代码，这样 Domain 层和应用层就职责不清，代码难维护

##### 接口层：

- 做参数的基本处理，比如入参校验，回参 DTO 转换（拆包、组包）

##### 基础层：

- 基础层是基于依赖倒置设计的，封装基础资源服务，通过依赖注入方式解耦
- 基础组件都放这提供对其他层的支撑（包括第三方工具、MQ、文件、缓存、数据库、搜索等）
- 实现仓储接口 DB，通常真正读写 DB，实现仓储接口的语句都是写在这块
