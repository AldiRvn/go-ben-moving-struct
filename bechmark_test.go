package gobenmovingstruct

import (
	"encoding/json"
	"testing"
)

func Benchmark_normalPerField(b *testing.B) {
	for b.Loop() {
		sessionNew := DeviceSessionNew{
			New:          true,
			ID:           session.ID,
			UserID:       session.UserID,
			DeviceID:     session.DeviceID,
			PhoneNumber:  session.PhoneNumber,
			Email:        session.Email,
			Username:     session.Username,
			DisplayName:  session.DisplayName,
			AvatarURL:    session.AvatarURL,
			SessionToken: session.SessionToken,
			RefreshToken: session.RefreshToken,
			IPAddress:    session.IPAddress,
			UserAgent:    session.UserAgent,
			DeviceModel:  session.DeviceModel,
			OSVersion:    session.OSVersion,
			AppVersion:   session.AppVersion,
			Region:       session.Region,
			Timezone:     session.Timezone,
			LoginMethod:  session.LoginMethod,
			IsActive:     session.IsActive,
			CreatedAt:    session.CreatedAt,
		}
		_ = sessionNew
	}
}

func Benchmark_marshallUnmarshall(b *testing.B) {
	for b.Loop() {
		sessionByte, _ := json.Marshal(session)
		newSession := DeviceSessionNew{}
		_ = json.Unmarshal(sessionByte, &newSession)
	}
}
