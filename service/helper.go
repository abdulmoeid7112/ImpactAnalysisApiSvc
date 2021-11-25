package service

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/abdulmoeid7112/impact-analysis-api-svc/config"
	"github.com/abdulmoeid7112/impact-analysis-api-svc/models"
)

type countryCases struct {
	Name       string
	CasesCount int
}

// generateJWT for token generation.
func generateJWT(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["issued_at"] = user.TokenIssuesAt
	claims["expired_at"] = calculateExpireTime(user.TokenIssuesAt)

	tokenString, err := token.SignedString( []byte(viper.GetString(config.JWTSecretKey)))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ScheduleNextBackup schedule next backup time
func calculateExpireTime(issueTime time.Time) time.Time {

	backupTime := strings.Split(fmt.Sprintf("%s", issueTime), ":")
	hours, _ := strconv.Atoi(backupTime[0])
	minutes, _ := strconv.Atoi(backupTime[1])
	seconds, _ := strconv.Atoi(backupTime[2])

	return timeToHourOfDay(issueTime.AddDate(0, 0, 0), hours, minutes, seconds)

}

func timeToHourOfDay(nextTime time.Time, hour, minute, seconds int) time.Time {
	data := nextTime.Add((time.Duration(nextTime.Hour()*-1) * time.Hour) + (time.Duration(nextTime.Minute()*-1) * time.Minute))
	data = data.Round(time.Hour)
	data = data.Add(time.Duration(hour) * time.Hour).Add(time.Duration(minute+30) * time.Minute).Add(time.Duration(seconds) * time.Second)

	return data
}
