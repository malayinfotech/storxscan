#!/bin/sh

dbx schema -d pgx -d pgxcockroach storxscandb.dbx .
dbx golang -d pgx -d pgxcockroach -p dbx -t templates storxscandb.dbx .
( printf '%s\n' '//lint:file-ignore U1000,ST1012 generated file'; cat storxscandb.dbx.go ) > storxscandb.dbx.go.tmp && mv storxscandb.dbx.go.tmp storxscandb.dbx.go
gofmt -r "*sql.Tx -> tagsql.Tx" -w storxscandb.dbx.go
gofmt -r "*sql.Rows -> tagsql.Rows" -w storxscandb.dbx.go
perl -0777 -pi \
  -e 's,\t_ "github.com/jackc/pgx/v4/stdlib"\n\),\t_ "github.com/jackc/pgx/v4/stdlib"\n\n\t"private/tagsql"\n\),' \
  storxscandb.dbx.go
perl -0777 -pi \
  -e 's/type DB struct \{\n\t\*sql\.DB/type DB struct \{\n\ttagsql.DB/' \
  storxscandb.dbx.go
perl -0777 -pi \
  -e 's/\tdb = &DB\{\n\t\tDB: sql_db,/\tdb = &DB\{\n\t\tDB: tagsql.Wrap\(sql_db\),/' \
  storxscandb.dbx.go
