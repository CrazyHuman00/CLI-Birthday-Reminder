#!/bin/bash

# Check if command argument is provided
MAIN=$HOME/bin-birthday/main.go

case "$1" in
    ("list")
        go run $MAIN list
        ;;
    ("today")
        go run $MAIN today
        ;;
    ("add")
        if [ -z "$2" ] || [ -z "$3" ]; then
            echo "Usage: birthday.sh add [name] [date]"
            exit 1
        fi
        go run $MAIN add "$2" "$3"
        ;;
    ("remove")
        if [ -z "$2" ]; then
            echo "Usage: birthday.sh remove [name]"
            exit 1
        fi
        go run $MAIN remove "$2"
        ;;
    ("update")
        if [ -z "$2" ] || [ -z "$3" ]; then
            echo "Usage: birthday.sh update [name] [date]"
            exit 1
        fi
        go run $MAIN update "$2" "$3"
        ;;
    (*)
        echo "Invalid command. Use 'list' or 'today'"
        exit 1
        ;;
esac
