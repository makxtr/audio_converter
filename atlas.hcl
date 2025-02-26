env "local" {
  src = "file:///app/db/schema.hcl"
  dev = "mysql://root:rootpass@tcp(db_mysql_user:3306)/project_db?parseTime=true"
}