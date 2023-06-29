#!/bin/bash

content="This is the line to append"
file="somefile"

echo "" >> "$file"
echo "$content" >> "$file"