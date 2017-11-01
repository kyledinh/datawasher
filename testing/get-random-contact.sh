#!/bin/bash
curl -H "Accept: application/json" -H "Content-type: application/json" -X GET \
http://localhost:8000/random_contact
