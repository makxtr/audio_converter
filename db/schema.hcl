schema "project_db" {
  table "users" {
    column "id" {
      type = int
      primary_key = true
      auto_increment = true
    }
    column "name" {
      type = string
      null = false
      length = 100
    }
    column "email" {
      type = string
      unique = true
      null = false
      length = 100
    }
    column "password" {
      type = string
      null = false
      length = 255
    }
  }
}