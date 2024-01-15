# Scanner Server
Scan blockchain data to DB and provide api to search

# Folder Structure
---
    |-- Project
        |-- cmd
            |-- {command name}          ---如果有多個service要啟動的話
                |-- main.go
        |-- conf.d 
            |-- app.yaml                ---config file
        |-- internal
            |-- constant
            |-- model                   ---define struct
            |-- db                      ---db連線
            |-- utils                   ---專案內部的通用工具
        |-- service
            |-- {service name}          ---如果有多個service要開發的話
                |-- config              ---service config define
                |-- app                 ---啟動點＋註冊
                |-- controller          ---handler function。request 和 response 的物件宣告在handler function裡面

---

# API Document
```
1. open docs/index.html
2. select local json file to open docs/swagger.json
```

# Start Service
```
go run ./cmd/api-server/main.go
go run ./cmd/scanner-server/main.go
```

or
```
docker-compose up -d
```

# Sequence Diagram
```mermaid
sequenceDiagram
title API-Service

actor client
participant api_service
participant app
participant rest
participant restctl
participant core
participant repository
participant DB

opt Server Start
    api_service->>app: new application
    app->>rest: new restful api object
    rest->>DB: new db connection
    DB-->>rest: return
    rest->>repository: init repository object
    repository-->>rest: dependency injection
    rest->>core: init core object
    core-->>rest: dependency injection
    rest->>restctl: init restctl object
    restctl-->>rest: dependency injection
    rest-->app: return 

    app->>rest: run restful api server
end

client->>rest: call api
rest->>restctl: call api handle function implement
restctl->>core: call core implement
core->>repository: call repo implement
repository->>DB: get data

DB-->>client: return result

```

```mermaid
sequenceDiagram

title Scanner-Service

participant scanner_service
participant app
participant core
participant repository
participant DB
participant eth

opt Server Start
    scanner_service->>app: new application
    app->>DB: new db connection
    DB-->>app: return
    app->>repository: init repository object
    repository-->>app: dependency injection
    app->>core: init core object
    core-->>app: dependency injection

    app->>core: run scanner-server
    core->>eth: start reat-time pipeline to get latest blocks
    core->>eth: start history pipeline to get history blocks

    eth-->>core: return result

    core->>repository: send data
    repository->>DB: storage data
    DB-->>repository: return result
end

```

## TODO List

- [x] update docker-compose.yml
- [ ] unit test
- [ ] optimizer scanner-server eth request
