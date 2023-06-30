#!/bin/bash

/home/passage-connect-client

if [ $? -eq 0 ]; then
  # Success: Do nothing, let the session continue
  exit 0
else
  # Failure: Terminate the session
  exit 1
fi
