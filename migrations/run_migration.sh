#!/usr/bin/bash
curl -LO https://edu.postgrespro.com/demo-small-en.zip && unzip demo-small-en.zip && rm demo-small-en.zip
# replace with your own credentials here
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=test -e POSTGRES_DB=demo -e POSTGRES_PASSWORD=secret -d postgres:13
# apply test data to newly created DB
psql demo -h localhost -U test -f demo-small-en-20170815.sql 