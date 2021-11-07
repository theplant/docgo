DIR=$(PWD)
snippetgo -pkg=docs > ./docs/examples-generated.go

function docsRestart() {
  echo "=================>"
  killall docgodocs
#  export DEV_CORE_JS=1
#  export DEV_VUETIFY_JS=1
#  export DEV_PRESETS=1
  go build -o /tmp/docgodocs docs/docsmain/main.go && /tmp/docgodocs
}

export -f docsRestart
find . -name "*.go" | entr -r bash -c "docsRestart"
