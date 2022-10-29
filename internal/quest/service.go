package quest

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetQuests(c *gin.Context) ([]*domain.QuestDTO, error)
	GetQuest(c *gin.Context, id string) (domain.QuestDTO, error)
	CreateQuest(c *gin.Context, id string, scene int, inventory []string) error
	UpdateQuest(c *gin.Context, quest domain.QuestDTO, paramId string) error
	DeleteQuest(c *gin.Context, id string) error
	CreateCompletion(c *gin.Context, questID int, userID int, startYear int, startMonth time.Month,
		startDay int, startHour int, startMinutes int, startSeconds int) error
	GetRating(c *gin.Context, questID int, userID int) (domain.Rating, error)
	CreateRating(c *gin.Context, questID int, userID int, rating int) error
	GetRanking(c *gin.Context, id int) ([]domain.QuestCompletionDTO, error)
}

type service struct {
	repo     Repository
	userRepo user.Repository
}

func NewService(rep Repository, userRepo user.Repository) Service {
	return &service{
		repo:     rep,
		userRepo: userRepo,
	}
}

func (s *service) GetQuests(c *gin.Context) ([]*domain.QuestDTO, error) {
	quests, err := s.repo.GetQuests(c)

	return quests, err
}

func (s *service) GetQuest(c *gin.Context, id string) (domain.QuestDTO, error) {
	quests, err := s.repo.GetQuest(c, id)

	return quests, err
}

func (s *service) CreateQuest(c *gin.Context, id string, scene int, inventory []string) error {
	fmt.Println("id: ", id)
	err := s.repo.CreateQuest(c, id, scene, inventory)

	return err
}

func (s *service) UpdateQuest(c *gin.Context, quest domain.QuestDTO, paramId string) error {

	err := s.repo.UpdateQuest(c, quest, paramId)

	return err
}

func (s *service) DeleteQuest(c *gin.Context, id string) error {
	err := s.repo.DeleteQuest(c, id)

	return err
}

func (s *service) CreateCompletion(c *gin.Context, questID int, userID int, startYear int, startMonth time.Month,
	startDay int, startHour int, startMinutes int, startSeconds int) error {

	startTime := time.Date(startYear, startMonth, startDay, startHour, startMinutes, startSeconds, 00, time.UTC).Add(time.Hour * 3)

	actualTime := time.Now()

	completion, err := s.repo.GetCompletion(c, questID, userID)

	if err != nil {
		err = s.repo.AddCompletion(c, questID, userID, startTime, actualTime)
		if err != nil {
			return err
		}
	}
	//Update quest completion quantity
	quest, err := s.repo.GetQuestInfo(c, questID)

	if err != nil {
		return err
	}

	quest.Completions++

	if err := s.repo.UpdateQuestInfo(c, quest); err != nil {
		return err
	}

	if !isBestTime(startTime, actualTime, completion.StartTime, completion.EndTime) {
		return nil
	}

	completion.StartTime = startTime
	completion.EndTime = actualTime

	err = s.repo.SaveCompletion(c, completion)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetRating(c *gin.Context, questID int, userID int) (domain.Rating, error) {
	rating, err := s.repo.GetRating(c, questID, userID)

	return rating, err
}

func (s *service) CreateRating(c *gin.Context, questID int, userID int, rating int) error {

	ratingRecord, err := s.repo.GetRating(c, questID, userID)

	if err != nil {
		ratingRecord.QuestID = questID
		ratingRecord.UserID = userID
		ratingRecord.Rate = rating
		err = s.repo.AddRating(c, ratingRecord)
	} else {
		ratingRecord.Rate = rating
		err = s.repo.AddRating(c, ratingRecord)
	}

	if err != nil {
		return err
	}

	quest, err := s.repo.GetQuestInfo(c, questID)

	if err != nil {
		return err
	}

	ratings, err := s.repo.GetQuestRatings(c, questID)

	if err != nil {
		return err
	}

	var ratingsSum int = 0
	for _, rating := range ratings {
		ratingsSum += rating.Rate
	}

	quest.Qualification = float32(ratingsSum) / float32(len(ratings))

	return s.repo.UpdateQuestInfo(c, quest)
}

func (s *service) GetRanking(c *gin.Context, id int) ([]domain.QuestCompletionDTO, error) {

	quests, err := s.repo.GetQuestsCompletions(c, id)

	sort.Sort(QuestsCompletions(quests))

	questCompletionsDTO := make([]domain.QuestCompletionDTO, len(quests))

	for i := range quests {

		userDTO, _, err := s.userRepo.GetUser(c, quests[i].UserID)

		if err != nil {
			return nil, err
		}

		questCompletionsDTO[i].Username = userDTO.Username
		questCompletionsDTO[i].UserImage = userDTO.Image
		questCompletionsDTO[i].StartTime = quests[i].StartTime
		questCompletionsDTO[i].EndTime = quests[i].EndTime

		questCompletionsDTO[i].StartTime = quests[i].StartTime
		questCompletionsDTO[i].EndTime = quests[i].EndTime

		diff := questCompletionsDTO[i].EndTime.Sub(questCompletionsDTO[i].StartTime)
		stringDiff := diff.String()

		splitHours := strings.Split(stringDiff, "h")

		var minsSeconds string

		if diff.Hours() < (time.Hour.Hours() * 1) {
			questCompletionsDTO[i].Hours = 0
			minsSeconds = splitHours[0]
		} else {
			hoursFloat, _ := strconv.ParseFloat(splitHours[0], 64)
			questCompletionsDTO[i].Hours = hoursFloat
			minsSeconds = splitHours[1]
		}

		splitMinsSeconds := strings.Split(minsSeconds, "m")

		minutesFloat, _ := strconv.ParseFloat(splitMinsSeconds[0], 64)

		questCompletionsDTO[i].Minutes = minutesFloat

		secondsString := strings.Replace(splitMinsSeconds[1], "s", "", -1)

		secondsFloat, _ := strconv.ParseFloat(secondsString, 64)

		questCompletionsDTO[i].Seconds = secondsFloat
	}

	return questCompletionsDTO, err
}

func isBestTime(startTime1 time.Time, endTime1 time.Time, startTime2 time.Time, endTime2 time.Time) bool {

	diff1 := endTime1.Sub(startTime1)
	diff2 := endTime2.Sub(startTime2)

	if diff1 < diff2 {
		return true
	}
	return false
}
