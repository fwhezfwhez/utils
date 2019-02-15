# utils
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/fwhezfwhez/utils)
[![Build Status]( https://www.travis-ci.org/fwhezfwhez/utils.svg?branch=master)]( https://www.travis-ci.org/fwhezfwhez/utils)

common functional utils for public

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [utils](#utils)
  - [1. start](#1-start)
  - [2. plugins](#2-plugins)
      - [2.1 redis](#21-redis)
      - [2.2 jwt](#22-jwt)
      - [2.3 jsoncrack](#23-jsoncrack)
      - [2.4 superchecker](#24-superchecker)
  - [3. Notes:](#3-notes)
      - [3.1 All plugins should be inited well before the app starts,so we don't add locks to make it concurrently safe.Neither can we change global tools's properties while the app is running.](#31-all-plugins-should-be-inited-well-before-the-app-startsso-we-dont-add-locks-to-make-it-concurrently-safeneither-can-we-change-global-toolss-properties-while-the-app-is-running)
      - [3.2 Add your plugins as you wish, but make sure a package with an `init.go` and at least a `xxx_test.go` file, when pull request is open, make sure travis-ci pass.](#32-add-your-plugins-as-you-wish-but-make-sure-a-package-with-an-initgo-and-at-least-a-xxx_testgo-file-when-pull-request-is-open-make-sure-travis-ci-pass)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 1. start
`go get github.com/fwhezfwhez/utils`

## 2. plugins
Most tools as global vals can be put into utils.Welcome to pull your request to fullfill the utils.

#### 2.1 redis
<a href="https://github.com/fwhezfwhez/utils/tree/master/util_redis">util_redis</a>

#### 2.2 jwt
<a href="https://github.com/fwhezfwhez/utils/tree/master/util_jwt">util_jwt</a>

#### 2.3 jsoncrack
<a href="https://github.com/fwhezfwhez/utils/tree/master/util_jsoncrack">util_jsoncrack</a>

#### 2.4 superchecker
<a href="https://github.com/fwhezfwhez/utils/tree/master/util_superchecker">util_superchecker</a>

## 3. Notes:
#### 3.1 All plugins should be inited well before the app starts,so we don't add locks to make it concurrently safe.Neither can we change global tools's properties while the app is running.
#### 3.2 Add your plugins as you wish, but make sure a package with an `init.go` and at least a `xxx_test.go` file, when pull request is open, make sure travis-ci pass.


