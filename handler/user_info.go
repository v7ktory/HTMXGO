package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/v7ktory/htmx+go/model"
)

const (
	apiURL = "https://api.api-ninjas.com/v1/randomuser"
)

func (h *Handler) FetchUser(c *gin.Context) {

	// Create an HTTP client
	client := &http.Client{}

	// Prepare a GET request to the API endpoint
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		handleError(c, "failed to create request", http.StatusInternalServerError)
		return
	}
	// Add the API key to the request headers
	req.Header.Add("X-Api-Key", os.Getenv("API"))

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
	var user model.User
	if err := json.Unmarshal(body, &user); err != nil {
		handleError(c, "failed to parse response body", http.StatusInternalServerError)
		return
	}

	// Render the user data using a template
	c.HTML(http.StatusOK, "user_info", gin.H{
		"Sex":      CheckSex(user.Sex),
		"Name":     user.Name,
		"Email":    user.Email,
		"Birthday": CalculateAge(user.Birthday),
	})
}

func CheckSex(sex string) string {
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

func CalculateAge(b string) int {
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
