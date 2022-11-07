# Terraform Provider Fireblocks

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
  abi_path = "<path-to-abi-json>"
}
```

Then, run 

```sh
$ terraform init
```

### Features

#### Supported Fireblocks API objects

| **API object**        | **Terraform Method** | **Supported**      | **Comment**                                                                                                   |
|-----------------------|----------------------|--------------------|---------------------------------------------------------------------------------------------------------------|
| Vault Account         | Create               | :green_circle:     | Some Fireblocks API parameters are not supported yet                                                          |
| Vault Account         | Update               | :green_circle:     | Vault account can be renamed and hidden from console                                                          |
| Vault Account         | Delete               | :orange_circle:    | Fireblocks API does not allow to archive vault account (while it is possible on the Fireblocks console)       |
| Vault Account         | Import               | :green_circle:     |                                                                                                               |
| Vault Account Asset   | Create               | :green_circle:     | Only for listed assets. Fireblocks API does not allow to create an asset with custom address                  |
| Vault Account Asset   | Update               | :heavy_minus_sign: | Fireblocks API does not allow to update Vault Account Wallet                                                  |
| Vault Account Asset   | Delete               | :orange_circle:    | Fireblocks API does not allow to archive vault account asset (while it is possible on the Fireblocks console) |
| Vault Account Asset   | Import               | :green_circle:     |                                                                                                               |
| External Wallet       | Create               | :green_circle:     |                                                                                                               |
| External Wallet       | Update               | :heavy_minus_sign: | Fireblocks API does not allow to update External Wallet                                                       |
| External Wallet       | Delete               | :green_circle:     |                                                                                                               |
| External Wallet       | Import               | :green_circle:     |                                                                                                               |
| External Wallet Asset | Create               | :green_circle:     |                                                                                                               |
| External Wallet Asset | Update               | :heavy_minus_sign: | Fireblocks API does not allow to update External Wallet Asset                                                 |
| External Wallet Asset | Delete               | :green_circle:     |                                                                                                               |
| External Wallet Asset | Import               | :green_circle:     |                                                                                                               |
| Transaction           | Create               | :green_circle:     |                                                                                                               |
| Transaction           | Udpate               | :heavy_minus_sign: | Fireblocks API does not allow to update Transaction                                                           |
| Transaction           | Delete               | :green_circle:     |                                                                                                               |
| Transaction           | Import               | :green_circle:     |                                                                                                               |

#### Ethereum support

This provider allows to load an ABI and craft any Ethereum transaction using terraform configuration to parametrize method and argument.

