module local_record_check

go 1.19

require record_mysql_barcheck v0.0.1

require query_bar_update v0.0.1

require global v0.0.1

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

replace record_mysql_barcheck => ../../record/record_mysql_barcheck

replace global => ../../global

replace query_bar_update => ../../query/query_bar_update
