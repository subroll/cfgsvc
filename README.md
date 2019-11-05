# cfgsvc
This is a hypothetical centralized configuration service.

## Requirement
- Golang: 1.13
- Database: mysql
- DB Migration: [pressly/goose](https://github.com/pressly/goose)

## How to run
### Clone this repository
```
git clone git@github.com:subroll/cfgsvc.git
```
### Install all required dependencies
```
go mod tidy && go mod vendor
```
### Comply config.yaml with the appropriate values
> you can copy from config.yaml.dist and fill in the right values

### Migrate data structure
- Install [pressly/goose](https://github.com/pressly/goose)
- Go to `migrations` directory
- Use migration command as instructed

### Build the main package
- Go to root directory of this repository
- Execute `go build -o <output name>` command 

### Usage
```
$ ./<output name>
```
The above command will give you any available command and flags in this service

```
$ ./<output name> serve
```
The `serve` command will run this service by searching `config.yaml` file in the 
same directory and start the http server on port **8080**

If you want to specify the port you want to use or 
if you want to specify the config file location, 
please execute the following command to see how to do it.
```
$ ./<output name> serve -h
```
