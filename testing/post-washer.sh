#!/bin/bash
curl -v -H "Accept: application/json" -H "Content-type: application/json" --upload-file contacts.json  \
-X POST http://localhost:8000/washer
