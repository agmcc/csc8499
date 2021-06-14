#!/bin/bash
source ./venv/bin/activate
locust --headless \
--host http://localhost:8080 \
--users 1000 \
--spawn-rate 10 \
--run-time 300s \
--csv test \
--csv-full-history \
--html test.html
