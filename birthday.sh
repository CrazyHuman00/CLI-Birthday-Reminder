#!/bin/bash

# This script prints a birthday message.
cd $HOME/bin-birthday

# build birthday command.
go build -o birthday

# run birthday today command.
./birthday today