package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/v7ktory/htmx+go/internal/model"
)

const (
	apiURL = "https://api.api-ninjas.com/v1/randomuser"
)

// Getting user info from ninja API
func (h *Handler) UserInfo(c *gin.Context) {

	// Create an HTTP client
	client := &http.Client{}

	// Prepare a GET request to the API endpoint
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		handleError(c, "failed to create request", http.StatusInternalServerError)
		return
	}
	// Add the API key to the request headers
	req.Header.Add("X-Api-Key", os.Getenv("API_KEY"))

	// Send the request and get the response
	response, err := client.Do(req)
	if err != nil {
		handleError(c, "failed to send request", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		handleError(c, "failed to read response body", http.StatusInternalServerError)
		return
	}

	// Unmarshal the response body into a User struct
	var user model.UserInfo

	if err := json.Unmarshal(body, &user); err != nil {
		handleError(c, "failed to parse response body", http.StatusInternalServerError)
		return
	}

	// Render the user data using a template
	c.HTML(http.StatusOK, "user_info", gin.H{
		"Sex":      checkSex(user.Sex),
		"Name":     user.Name,
		"Email":    user.Email,
		"Birthday": calculateAge(user.Birthday),
	})
}

func checkSex(sex string) string {
	switch sex {
	case "M":
		sex = "Гигачад"
	case "F":
		sex = "Вумен"
	default:
		sex = "Неизвестно"
	}
	return sex
}

func calculateAge(b string) int {
	yearStr := b[:4]

	birthYear, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0
	}

	now := time.Now()
	age := now.Year() - birthYear

	if now.YearDay() < time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC).YearDay() {
		age--
	}

	return age
}

// Getting user info from session
func (h *Handler) GetUserInfo(c *gin.Context) {

	// Get the session from the authorization header
	sessionHeader := c.GetHeader("Authorization")

	// Ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		handleError(c, "invalid session header", 400)
		return
	}

	// Get the session id
	sessionID := sessionHeader[7:]

	// Get the user data from the session
	user, err := h.SessionManager.GetSession(sessionID)
	if err != nil {
		handleError(c, "failed to get user session", 500)
		return
	}

	c.JSON(200, user)
}