dump:
	sqlite3 db.sqlite .dump > dump.sql

restore:
	rm -f old.db.sqlite

	$(if $(wildcard db.sqlite), mv db.sqlite old.db.sqlite)

	sqlite3 db.sqlite < dump.sql