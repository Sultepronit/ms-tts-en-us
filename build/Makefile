dump:
	sqlite3 db.sqlite .dump > dump.sql

restore:
	rm -f db.old.sqlite
	$(if $(wildcard db.sqlite), mv db.sqlite db.old.sqlite)
	sqlite3 db.sqlite < dump.sql

edit:
	rm -f db.old.sqlite
	$(if $(wildcard db.sqlite), cp db.sqlite db.old.sqlite)
	sqlite3 db.sqlite < ../db/edit/actual.sql