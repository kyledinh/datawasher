#!/bin/bash
curl -v -H "Accept: application/json" -H "Content-type: application/json" --data single-contact.json  \
-X POST http://localhost:8000/json_contacts
