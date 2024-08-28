package main


type Users struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Gender   string `json:"gender"`
    Photo    string `json:"photo"`
    Address  string `json:"address"`
    RoleID   uint   `json:"role"`
}

type Roles struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    RoleName string `json:"role_name"`
}

type ACS struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Brand string `json:"brand"`
    PK    string `json:"pk"`
    Price float64 `json:"price"`
}

type Services struct {
    ID           uint   `json:"id" gorm:"primaryKey"`
    TechnicianID uint   `json:"technician_id"`
    ClientID     uint   `json:"client_id"`
    ACID         uint   `json:"ac_id"`
    Date         string `json:"date"`
    Status       string `json:"status"`
}
