package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var db *gorm.DB

func initDB() {
    dsn := "root:@tcp(localhost:3306)/tugas-mysql?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("gagal terhubung database:", err)
    }

    // Migrasi model ke database
    db.AutoMigrate(&Users{}, &Roles{}, &ACS{}, &Services{})
}
