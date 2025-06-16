#!/bin/bash

# FUNCTIONS
# Cleanup function
function cleanup {
	true
}

function error {
	echo "$1"
	exit 1
}

# Trap exit.
trap cleanup EXIT

# Move folder
SOURCE="${BASH_SOURCE[0]}"
while [ -h "${SOURCE}" ]; do # resolve $SOURCE until the file is no longer a symlink
  BIN_DIR="$(cd -P "$(dirname "$SOURCE}")" && pwd)"
  readonly BIN_DIR
  SOURCE="$(readlink "${SOURCE}")"
  [[ ${SOURCE} != /* ]] && SOURCE="${BIN_DIR}/${SOURCE}" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$(cd -P "$(dirname "${SOURCE}")" && pwd)"
readonly DIR
# Move to this directory
cd "${DIR}" || exit

cd .. || exit

echo "Going to delete from $(pwd)"
echo 'SURE? (y)'
read -r INPUT
if [[ $INPUT != 'y' ]]
then
  echo quitting
  exit
fi
rm -rf jira_old
rm -rf checklists
rm -rf people/*.md
rm -rf people/*.last_updated
rm -rf people/*.pdf
rm -rf people/*last_contacted
rm -rf people/*birthday
rm -rf priority/*/[0-9]*
rm -rf status/*/[0-9]*
rm -rf tasks/[0-9]*
rm -rf webserver
rm -rf meetings/*
echo 0 > .metadata/lasttask
rm -rf .git
git init
