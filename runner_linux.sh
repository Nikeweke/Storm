export GOPATH=$(pwd)
export PATH=$PATH:$GOPATH/bin

# go-assets-builder views -o app/helpers/views.go
# fresh

go run main.go
# go build main.go