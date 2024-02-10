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

func (h *Handler) UserProfile(c *gin.Context) {
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}

	s, err := h.SessionManager.GetSession(c, cookie)
	if err != nil {
		handleError(c, "failed to get session", http.StatusBadRequest)
		return
	}
	todos, err := h.Service.GetTodos(c, s.ID)
	if err != nil {
		handleError(c, "failed to get todos", http.StatusBadRequest)
		return
	}

	var todoCompleted []model.TodoInfo
	for _, t := range todos {
		TodoInfo := model.TodoInfo{
			ID:        t.ID,
			Completed: t.Completed,
		}
		todoCompleted = append(todoCompleted, TodoInfo)
	}

	// Render the profile template
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"ID":    s.ID,
		"Name":  s.Name,
		"Email": s.Email,
		"Total": sumCompleted(todoCompleted),
	})
}

func sumCompleted(todos []model.TodoInfo) int {
	totalCompleted := 0
	for _, todo := range todos {
		if todo.Completed {
			totalCompleted++
		}
	}
	return totalCompleted
}

// Get the information from ninja API
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
	year, _ := strconv.Atoi(b[:4])
	age := time.Now().Year() - year
	if time.Now().YearDay() < time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC).YearDay() {
		age--
	}
	return age
}
