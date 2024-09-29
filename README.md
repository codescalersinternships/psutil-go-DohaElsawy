
# psutil-go Package 
psutil-go is a lightweight, library in Go that provides essential system information.

## Installation 
- 1. Download project
```golang
   git get https://github.com/codescalersinternships/psutil-go-DohaElsawy.git
```
- 2. import package :
```golang
   import pokeClient "github.com/codescalersinternships/psutil-go-DohaElsawy.git"
```
### Functions
- 1 -  Get CPU informations:
```golang

  c, err := GetCpuInfo()
```

- 2 - Get Memory informations:
 ```golang
  m, err := GetMemInfo()
```

- 3 - Get all running process list:
 ```golang
 listproc, err := ListProc()
```

- 2 - Get process details by its id:
 ```golang
  procDetails, err := GetProcDetails(PID)
```

### Test
- to run all tests
```golang
  make test
```
### Format
- format all files inside project
```golang
  make format
```
