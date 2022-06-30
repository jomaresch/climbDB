gen-orm:
	cd ./configs && sqlboiler sqlite3 -o ./../internal/database/models
