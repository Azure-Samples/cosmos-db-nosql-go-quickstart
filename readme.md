<!--
---
page_type: sample
name: "Quickstart: Azure Cosmos DB for NoSQL and Azure SDK for Go"
description: This is a simple web application to illustrate common basic usage of Azure Cosmos DB for NoSQL and the Azure SDK for Go.
urlFragment: template
languages:
- go
- azdeveloper
products:
- azure-cosmos-db
---
-->

# Quickstart: Azure Cosmos DB for NoSQL client library for Go

This is a simple web application to illustrate common basic usage of Azure Cosmos DB for NoSQL's client library for Go. This sample application accesses an existing account, database, and container using the [`azcosmos`](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos) and [`azidentity`](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity) packages.

## Prerequisites

- [Docker](https://www.docker.com/)
- [Azure Developer CLI](https://aka.ms/azd-install)
- [Go 1.21 or newer](https://go.dev/dl/)

## Quickstart

1. Log in to Azure Developer CLI.

    ```bash
    azd auth login
    ```

    > [!TIP]
    > This is only required once per-install.

1. Initialize this template (`cosmos-db-nosql-go-quickstart`) using `azd init`

    ```bash
    azd init --template cosmos-db-nosql-go-quickstart
    ```

1. Ensure that **Docker** is running in your environment.

1. Use `azd up` to provision your Azure infrastructure and deploy the web application to Azure.

    ```bash
    azd up
    ```

1. Observed the deployed web application

    ![Screenshot of the deployed web application.](assets/web.png)
