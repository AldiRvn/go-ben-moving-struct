package gobenmovingstruct

import "time"

type DeviceSession struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	DeviceID     string    `json:"device_id"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	DisplayName  string    `json:"display_name"`
	AvatarURL    string    `json:"avatar_url"`
	SessionToken string    `json:"session_token"`
	RefreshToken string    `json:"refresh_token"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	DeviceModel  string    `json:"device_model"`
	OSVersion    string    `json:"os_version"`
	AppVersion   string    `json:"app_version"`
	Region       string    `json:"region"`
	Timezone     string    `json:"timezone"`
	LoginMethod  string    `json:"login_method"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}

type DeviceSessionNew struct {
	New          bool      `json:"new"`
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	DeviceID     string    `json:"device_id"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	DisplayName  string    `json:"display_name"`
	AvatarURL    string    `json:"avatar_url"`
	SessionToken string    `json:"session_token"`
	RefreshToken string    `json:"refresh_token"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	DeviceModel  string    `json:"device_model"`
	OSVersion    string    `json:"os_version"`
	AppVersion   string    `json:"app_version"`
	Region       string    `json:"region"`
	Timezone     string    `json:"timezone"`
	LoginMethod  string    `json:"login_method"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}
