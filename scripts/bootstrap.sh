#!/bin/bash

go mod tidy

cargo build

pip install -r ml/requirements.txt