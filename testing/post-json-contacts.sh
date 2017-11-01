#!/bin/bash
curl -v -H "Accept: application/json" -H "Content-type: application/json" -F "data=@contacts.json" \
-X POST http://localhost:8000/json_contacts
