dump:
	sqlite3 db.sqlite .dump > dump.sql

restore:
	rm -f old.db.sqlite

	if [ -f db.sqlite ]; then
		mv sb.sqlite old.db.sqlite
	fi

	sqlite3 db.sqlite < dump.sql