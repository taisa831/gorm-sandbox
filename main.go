package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
  //gorm.Model
  ID uint `gorm:"primary_key"`
  Code  string
  Price uint
}

func main() {
  // db, err := gorm.Open("sqlite3", "test.db")
  db, err := gorm.Open("mysql", "gorm:gorm@/sandbox?charset=utf8mb4&parseTime=True&loc=Local")
  if err != nil {
    panic("データベースへの接続に失敗しました")
  }
  defer db.Close()

  db.LogMode(true)

  // db.SingularTable(true)

  // スキーマのマイグレーション
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "L1212", Price: 1000})

  // Read
  var product Product
  db.First(&product, 1)                   // idが1の製品を探します
  db.First(&product, "code = ?", "L1212") // codeがL1212の製品を探します

  // Update - 製品価格を2,000に更新します
  db.Model(&product).Update("Price", 2000)

  // Delete - 製品を削除します
  product.ID = 2
  db.Delete(&product)
}
