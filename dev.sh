DIR=$(PWD)
snippetgo -pkg=docsrc > ./docsrc/examples-generated.go

cd $(PWD)/docsrc
go run ./build/main.go

function docsRestart() {
  echo "=================>"
  killall docgodocs
#  export DEV_CORE_JS=1
#  export DEV_VUETIFY_JS=1
#  export DEV_PRESETS=1
  go build -o /tmp/docgodocs ./docsmain/main.go && /tmp/docgodocs
}

export -f docsRestart
find . -name "*.go" | entr -r bash -c "docsRestart"

