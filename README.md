# Sample Function: Go "Random"

## Introduction

This repository contains a sample "Random" function written in Go. You can deploy it on DigitalOcean's App Platform as a Serverless Function component or as a standalone Function. Documentation is available at https://docs.digitalocean.com/products/functions.

**Note: Following these steps may result in charges for the use of DigitalOcean services.**

### Requirements

* You need a DigitalOcean account. If you don't already have one, you can sign up at [https://cloud.digitalocean.com/registrations/new](https://cloud.digitalocean.com/registrations/new).
* To deploy from the command line, you will need the [DigitalOcean `doctl` CLI](https://github.com/digitalocean/doctl/releases).

## Deploying the Function

Clone this repo:

```shell
git clone git@github.com:digitalocean/sample-functions-golang-random.git
```

Deploy the project, using a remote build so that compiled executable matches the runtime environment:

```shell
doctl serverless deploy sample-functions-golang-random --remote-build
```

The output from the deploy command will resemble the following:

```
Deploying 'sample-functions-golang-random'
  to namespace 'fn-...'
  on host '...'
Submitted function 'sample/random' for remote building and deployment in runtime go:default (id: ...)
Deployment status recorded in 'sample-functions-golang-random/.deployed'

Deployed functions ('doctl sls fn get <funcName> --url' for URL):
  - sample/random
```

## Using the Function

Invoking the function will return a random number from 0 to 100.

```
doctl serverless functions invoke sample/random
```

Use this command to retrieve the URL for your function and use it as an API.
```
doctl sls fn get sample/random --url
```

You can use that API directly in your browser, with `curl` or with an API platform such as Postman.
Parameters may be passed as query parameters, or as JSON body. Here is an example using `curl`.

```
curl `doctl sls fn get sample/random --url`
```

### Learn More

You can learn more about Functions by reading the [Functions Documentation](https://docs.digitalocean.com/products/functions). 
