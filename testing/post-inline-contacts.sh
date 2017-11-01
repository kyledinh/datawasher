#!/bin/bash
curl -v -H "Accept: application/json" -H "Content-type: application/json" -d '[{"first_name":"Jerry","last_name":"Gergich","email":"jerry.gergich@boeing.com","phone_number":"555-792-3506","street_address":"9 Street Rd","city":"Pawnee","state":"IN"},
{"first_name":"Leslie","last_name":"Knope","email":"leslie.knope@cocaola.com","phone_number":"555-843-8116","street_address":"68 Street Rd","city":"Pawnee","state":"IN"},
{"first_name":"Andy","last_name":"Dwyer","email":"andy.dwyer@netflix.com","phone_number":"555-715-8261","street_address":"73 Street Rd","city":"Pawnee","state":"IN"},
{"first_name":"Betty","last_name":"Baker","email":"bbaker@gmail.com","phone_number":"555-395-7681","street_address":"56 Silver Street","city":"Pawnee","state":"IN"},
{"first_name":"Ann","last_name":"Perkins","email":"ann.perkins@primeconsulting.com","phone_number":"555-207-1944","street_address":"72 Street Rd","city":"Pawnee","state":"IN"},
{"first_name":"Donna","last_name":"Meagle","email":"donna.meagle@boeing.com","phone_number":"555-595-7884","street_address":"21 Street Rd","city":"Pawnee","state":"IN"}
]' \
-X POST http://localhost:8000/json_contacts
