package quest

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetQuests(c *gin.Context) ([]*domain.QuestDTO, error)
	GetQuest(c *gin.Context, id string) (domain.QuestDTO, error)
	CreateQuest(c *gin.Context, name string) error
	UpdateQuest(c *gin.Context, id string, quest domain.QuestDTO) error
	DeleteQuest(c *gin.Context, id string) error
	CreateCompletion(c *gin.Context, questID int, userID int, startYear int, startMonth time.Month,
		startDay int, startHour int, startMinutes int, startSeconds int) error
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) GetQuests(c *gin.Context) ([]*domain.QuestDTO, error) {
	quests, err := s.repo.GetQuests(c)

	return quests, err
}

func (s *service) GetQuest(c *gin.Context, id string) (domain.QuestDTO, error) {
	quests, err := s.repo.GetQuest(c, id)

	return quests, err
}

func (s *service) CreateQuest(c *gin.Context, name string) error {
	err := s.repo.CreateQuest(c, name)

	return err
}

func (s *service) UpdateQuest(c *gin.Context, id string, quest domain.QuestDTO) error {

	err := s.repo.UpdateQuest(c, id, quest)

	return err
}

func (s *service) DeleteQuest(c *gin.Context, id string) error {
	err := s.repo.DeleteQuest(c, id)

	return err
}

func (s *service) CreateCompletion(c *gin.Context, questID int, userID int, startYear int, startMonth time.Month,
	startDay int, startHour int, startMinutes int, startSeconds int) error {

	fmt.Println(startYear)
	startTime := time.Date(startYear, startMonth, startDay, startHour, startMinutes, startSeconds, 00, time.UTC)

	fmt.Println(startTime)
	actualTime := time.Now()
	fmt.Println(actualTime)
	diff := actualTime.Sub(startTime)

	res1 := strings.Split(diff.String(), "h")

	hours, _ := strconv.ParseFloat(res1[0], 32)

	res2 := strings.Split(res1[1], "m")

	mins, _ := strconv.ParseFloat(res2[0], 32)

	res3 := strings.Split(res2[1], "m")

	segsString := strings.Replace(res3[0], "s", "", -1)

	segs, _ := strconv.ParseFloat(segsString, 32)

	err := s.repo.CreateCompletion(c, questID, userID, actualTime, hours, mins, segs)

	return err
}
