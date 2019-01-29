#!/bin/bash

curl -s -d '{"args": "help"}' http://localhost:6666/cmd

echo
curl -s -d '{"args": "echo this is a test | md5 | base64"}' http://localhost:6666/cmd
