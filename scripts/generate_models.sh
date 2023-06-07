#!/bin/bash

# Array of names
names=("ask" "comment" "job" "poll" "pollopt" "story" "item" "user")

# Executable path
executable="gojsonstruct"

# Iterate over the array
for name in "${names[@]}"
do
    # Capitalize first letter
    # https://stackoverflow.com/a/12487465
    struct_name="$(tr '[:lower:]' '[:upper:]' <<< ${name:0:1})${name:1}"
    # Call the executable with parameters
    $executable --packagename "models" -omitempty always -typename "$struct_name" -o "./api/models/$name.go" "./api/json/$name.json"
done