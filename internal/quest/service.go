package quest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"
	"gorm.io/datatypes"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetQuests(c *gin.Context) ([]*domain.QuestDTO, error)
	GetQuest(c *gin.Context, id string) (domain.QuestDTO, error)
	CreateQuest(c *gin.Context, id string, scene int, inventory []string, logs []string, points float64) error
	CreateQuestProgression(c *gin.Context, id int, teamId int) error
	GetQuestProgression(c *gin.Context, id int, teamId int) (datatypes.JSON, error)
	UpdateQuestProgression(c *gin.Context, id int, teamId int, scene int, inventory []string, logs []string, objects map[string]int, points float32, finished bool) error
	GetTimeDifference(c *gin.Context, id int, teamId int, compareTime int64) (int64, error)
	SendUpdate(c *gin.Context, teamID int, userID int, itemName string) error
	UpdateQuest(c *gin.Context, quest domain.QuestDTO, paramId string) error
	DeleteQuest(c *gin.Context, id string) error
	CreateCompletion(c *gin.Context, questID int, userID int, startYear int, startMonth time.Month,
		startDay int, startHour int, startMinutes int, startSeconds int) error
	GetRating(c *gin.Context, questID int, userID int) (domain.Rating, error)
	CreateRating(c *gin.Context, questID int, userID int, rating int) error
	GetRanking(c *gin.Context, id int) ([]domain.QuestCompletionDTO, error)
	GetQuestRanking(c *gin.Context, id int) ([]domain.QuestProgressDTO, error)
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

func (s *service) CreateQuest(c *gin.Context, id string, scene int, inventory []string, logs []string, points float64) error {
	err := s.repo.CreateQuest(c, id, scene, inventory, logs, points)

	return err
}

func (s *service) CreateQuestProgression(c *gin.Context, id int, teamId int) error {

	_, err := s.repo.GetQuestProgression(c, id, teamId)

	now := time.Now()
	sec := now.Unix()

	//if quest already exists, start new quest
	if err == nil {
		err := s.repo.UpdateQuestProgression(c, id, teamId, 0, []string{}, []string{}, map[string]int{}, 0, false, sec)
		if err != nil {
			return err
		}
		return nil
	}

	err = s.repo.CreateQuestProgression(c, id, teamId, 0, []string{}, []string{}, map[string]int{}, 0, false, sec)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetQuestProgression(c *gin.Context, id int, teamId int) (datatypes.JSON, error) {
	questProgress, err := s.repo.GetQuestProgression(c, id, teamId)

	if err != nil {
		return nil, err
	}
	return questProgress, nil
}

func (s *service) GetTimeDifference(c *gin.Context, id int, teamId int, compareTime int64) (int64, error) {
	questProgress, err := s.repo.GetQuestProgressionInfo(c, id, teamId)

	if err != nil {
		return 0, err
	}

	difference := compareTime - questProgress.StartTime

	return difference, nil
}

func (s *service) UpdateQuestProgression(c *gin.Context, id int, teamId int, scene int, inventory []string, logs []string, objects map[string]int, points float32, finished bool) error {
	questProgress, err := s.repo.GetQuestProgressionInfo(c, id, teamId)

	if err != nil {
		return err
	}

	err = s.repo.UpdateQuestProgression(c, id, teamId, scene, inventory, logs, objects, points, finished, questProgress.StartTime)

	return err
}

func (s *service) SendUpdate(c *gin.Context, teamID int, userID int, itemName string) error {
	// fmt.Println("Team id: ", teamID)
	// fmt.Println("User id: ", userID)
	// fmt.Println("Item name: ", itemName)

	team, err := s.repo.GetTeam(c, teamID)
	if err != nil {
		return err
	}

	senderDTO, _, err := s.userRepo.GetUser(c, userID)

	for i := range team {
		userDTO, _, err := s.userRepo.GetUser(c, team[i].UserID)
		if err != nil {
			return err
		}

		//Encode the data
		postBody, _ := json.Marshal(map[string]string{
			"team_id":   strconv.Itoa(teamID),
			"sender":    senderDTO.Name,
			"token":     userDTO.FirebaseToken,
			"item_name": itemName,
		})
		responseBody := bytes.NewBuffer(postBody)
		//Leverage Go's HTTP Post function to make request
		resp, err := http.Post(config.GetConfig("dev").APP_NOTIFICATIONS_URL+"notifications/quest_update", "application/json", responseBody)
		//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
	}

	return nil
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

	err = s.repo.UpdateQuestInfo(c, quest)
	if err != nil {
		return err
	}

	err = s.userRepo.UnlockAchivement(c, userID, "RatedQuest_ac")
	if err != nil {
		return err
	}

	return nil
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

func (s *service) GetQuestRanking(c *gin.Context, id int) ([]domain.QuestProgressDTO, error) {

	quests, err := s.repo.GetQuestProgressions(c, id)

	questProgresses := make([]domain.QuestProgressDTO, len(quests))

	for i := range quests {
		if err != nil {
			return nil, err
		}
		questProgresses[i].Info = quests[i].Info
		questProgresses[i].TeamID = quests[i].TeamID
		questProgresses[i].Points = quests[i].Points

		team, err := s.repo.GetTeam(c, quests[i].TeamID)
		if err != nil {
			return nil, err
		}

		for j := range team {
			userDTO, _, err := s.userRepo.GetUser(c, team[j].UserID)
			if err != nil {
				return nil, err
			}
			questProgresses[i].Users = append(questProgresses[i].Users, domain.UserDTO{Username: userDTO.Username, Image: userDTO.Image})
		}
	}

	sort.Sort(QuestsProgresses(questProgresses))

	return questProgresses, err
}

func isBestTime(startTime1 time.Time, endTime1 time.Time, startTime2 time.Time, endTime2 time.Time) bool {

	diff1 := endTime1.Sub(startTime1)
	diff2 := endTime2.Sub(startTime2)

	if diff1 < diff2 {
		return true
	}
	return false
}
