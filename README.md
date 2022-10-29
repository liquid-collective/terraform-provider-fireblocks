# Terraform Provider OpenZeppelin Defender

[![GoDoc](https://pkg.go.dev/badge/github.com/liquid-collective/terraform-provider-fireblocks.svg)](https://pkg.go.dev/github.com/liquid-collective/terraform-provider-fireblocks)
[![Go Report Card](https://goreportcard.com/badge/github.com/liquid-collective/terraform-provider-fireblocks)](https://goreportcard.com/report/github.com/liquid-collective/terraform-provider-fireblocks)
[![Release](https://img.shields.io/github/v/release/liquid-collective/terraform-provider-fireblocks?logo=terraform&include_prereleases&style=flat-square)](https://github.com/liquid-collective/terraform-provider-fireblocks/releases)
[![Codecov](https://img.shields.io/codecov/c/github/liquid-collective/terraform-provider-fireblocks?logo=codecov&style=flat-square)](https://codecov.io/gh/liquid-collective/terraform-provider-fireblocks)
[![License](https://img.shields.io/github/license/liquid-collective/terraform-provider-fireblocks.svg?logo=fossa&style=flat-square)](https://github.com/liquid-collective/terraform-provider-fireblocks/blob/master/LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/liquid-collective/terraform-provider-fireblocks/Main/master?logo=github&style=flat-square)](https://github.com/liquid-collective/terraform-provider-fireblocks/actions?query=branch%3Amaster)

Fireblocks Terraform Provider is a plugin for managing resources on Fireblocks using
[Terraform](https://www.terraform.io/).

---

## Documentation

- [Official Docs](https://registry.terraform.io/providers/liquid-collective/fireblocks/latest/docs)

## Getting Started

### Requirements

- [Terraform](https://www.terraform.io/downloads)
- A [Fireblocks](https://fireblocks.com/) account

### Installation

This provider is available on [Terraform Registry](https://registry.terraform.io/). 

To use it this provider, copy and paste the following code into your Terraform configuration.

```terraform
terraform {
  required_providers {
    fireblocks = {
      source = "liquid-collective/fireblocks"
    }
  }
}

provider "fireblocks" {
  api_key = "<api-key>"
  rsa_private_key = "<api-key>"
}
```

Then, run 

```sh
$ terraform init
```
