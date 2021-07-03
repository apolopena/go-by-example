#!/bin/bash
# b.sh: Build a go binary from src/ into bin/, execute that binary if a -e flag is passed

check() {
  [[ ! -f $f ]] && echo "Error: $f does not exist" && exit 1
  return 0
}

build () {
  check "$f" && go build "$f" && mv "$(basename "$b1")" "$b2"
}

main() {
  if [[ $1 == '-e' ]]; then
    f="src/$2.go" && b1="${f%.*}" && b2="bin/$(basename "$b1")"
    build && ./"$b2"
  else
    f="src/$1.go" && b1="${f%.*}" && b2="bin/$(basename "$b1")"
    build
  fi
}

main "$@"