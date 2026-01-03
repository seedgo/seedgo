# Usage
## 1. Install this cmd tool `seedgo`
```bash
 go install github.com/seedgo/seedgo@v0.0.5 // install specific version
 // or
 go install github.com/seedgo/seedgo@latest // install latest version
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

## 2. Using cmd tool to create one project
```bash
seedgo h
seedgo create project name {projectname}
```
replace this {projectname} to your actual projectname.
