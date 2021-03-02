# kamgo Microservice

## Description

kamgo is a microservice that provides API for managing kamailio database.

## Requirements

- Go 1.12 or higher

- Do NOT disable [Go modules](https://github.com/golang/go/wiki/Modules) (`export GO111MODULE=on`)

## Configuration (config/config.toml)

- release_mode - this will enable and disable the debug mode. default is false.
- log_level - logging level, it can be DEBUG,INFO,WARN,ERROR,OFF
- [app] name - name of the application
- [app] version - version of the application
- [server] graceful - graceful shutdown
- [server] addr - socket address
- [kam_database] name - kamailio database name
- [kam_database] user_name - database username
- [kam_database] host - database host
- [kam_database] port - database port

### sample configuration

Sample configuration is config/config.toml. You can modify the database credentials as your enviornment.

## Installation

There are 2 methods explained in this document to install kamgo:

1. Manual Installation Steps
2. Docker Based Installation

### Method 1: Manual Installation Steps

1. Clone the kamgo into your go path and change directory to cloned repository.

   ```bash
   
   git clone https://@github.com/siprtc-io/kamgo.git
   
   cd kamgo
   ```

2. Build the kamgo microservice

   ```bash
   
   go build -o kamgo main.go
   
   ```

3. Set the configuration as per your requirement

    ```bash
    
    vim config/config.toml
    
    ```

4. Now you are ready the run the kamgo microservice

   ```bash
   ./kamgo
   ```

### Method 2: Docker Based Installation

1. Clone the kamgo into your go path and change directory to cloned repository.

   ```bash
   
   git clone https://@github.com/siprtc-io/kamgo.git
   
   cd kamgo
   ```

2. Build the kamgo microservice

   ```bash
   
   docker build -t kamgo .
   
   ```

3. Now you are ready the run the kamgo microservice

    ```bash
    
    docker run -p 9093:9093 kamgo
    
    ```

## API's

API Documentation link is coming soon.