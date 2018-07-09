# Hello World Serverless Golang

## Introduction

Hello World to Serverless.com + Golang on AWS

## Contents

- [Everything](#everything)

## Everything

(Usually I don't have an "Everything" section.. but the steps are just too easy..)

```bash
# Prepare AWS
aws s3 mb s3://org-name-us-west-2-serverless-deploys

# Vendor, test
go get -u github.com/kardianos/govendor
govendor sync
go test ./... -cover

# Install, login, build, deploy
npm install serverless -g
sls login
make
export AWS_PROFILE=personal
sls deploy

# Invoke lambda
sls invoke -f world

# cURL API
curl -X POST -d "hello world" https://REDACTED.execute-api.us-west-2.amazonaws.com/dev/hello
```
