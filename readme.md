## Contact me
if you have any question, please contact me.
* Email: [mfkhao2009@outlook.com](mfkhao2009@outlook.com)
  
## Config
Default config file is `./config/application.yaml`
we can change it by `--config /path/to/config/file`

## Command Tool
### 1. Install this cmd tool `seedgo`
```bash
 go install github.com/seedgo/seedgo-cli/cmd/seedgo@v0.0.6 // install specific version
 // or
 go install github.com/seedgo/seedgo-cli/cmd/seedgo@latest // install latest version
```
this tool will be installed in $GOPATH/bin directory, you need adding this folder to your PATH env variable.

check the GOPATH
```bash
go env | grep GOPATH
```

add $GOPATH/bin to your PATH env variable.
```bash
# ususally $GOPATH is `~/go`
export PATH=~/go/bin:$PATH
```

### 2. Using cmd tool to create one project
```bash
seedgo h
seedgo create project name {projectname}
```
replace this {projectname} to your actual projectname.

### 3. Run your project and check
```bash
curl http://localhost:10016/health
```
