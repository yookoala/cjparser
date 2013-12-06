#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f $0)")
SRC_DIR="$BASE_DIR/src"
DB_PATH="$BASE_DIR/data/cangjie.db"
DATA_DIR="$BASE_DIR/data"

# check dependencies and try to fix
echo -ne "Check dependencies ... "
go list github.com/mattn/go-sqlite3 2>&1 1>/dev/null
if [ $? -ne 0 ]; then
  echo -ne "install now ... "
  go get github.com/mattn/go-sqlite3
  if [ $? -ne 0 ]; then
    echo "Failed"
    echo "Cannot install package github.com/mattn/go-sqlite3"
    exit 1
  fi
fi
echo "Success"

# remove previously built database file
if [ -f "$DB_PATH" ]; then
  echo -ne "Remove old database ... "
  rm "$DB_PATH"
  echo "Success"
fi

# build the parser tool
echo -ne "building the parser ... "
cd "$SRC_DIR"
go fmt && go build -o ../parseCj
RET=$?
if [ $RET -eq 0 ]; then
  echo "Success"
else
  echo "Failed"
  exit $RET
fi

# build the database
echo -ne "build the database ... "
cd ..
./parseCj -f "$DATA_DIR" -d "$DB_PATH"
RET=$?
if [ $RET -eq 0 ]; then
  echo "Success"
else
  echo "Failed"
  exit $RET
fi

exit $RET

