package team

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"
)

type Service interface {
	CreateTeam(c *gin.Context, ids []int, questID int) (int, error)
	GetTeam(c *gin.Context, teamId int) ([]domain.UserDTO, error)
	AddCompletion(c *gin.Context, id int, questId int, startYear int, startMonth time.Month, startDay int, startHour int, startMinutes int, startSeconds int) error
	GetRanking(c *gin.Context, questId int) ([]domain.QuestTeamCompletionDTO, error)
	GetWaitRoomAccepted(c *gin.Context, teamId int, questId int) ([]domain.UserDTO, error)
	DeleteTeam(c *gin.Context, teamId int) error
	AcceptQuestTeam(c *gin.Context, teamId int, userId int) error
}

type service struct {
	repo      Repository
	userRepo  user.Repository
	questRepo quest.Repository
}

func NewService(rep Repository, userRepo user.Repository, questRepo quest.Repository) Service {
	return &service{
		repo:      rep,
		userRepo:  userRepo,
		questRepo: questRepo,
	}
}

func (s *service) CreateTeam(c *gin.Context, ids []int, questID int) (int, error) {

	teamID, err := s.repo.CreateTeam(c)

	if err != nil {
		return 0, err
	}

	for i := range ids {
		err = s.repo.AddPlayer(c, teamID, ids[i], questID)
		if err != nil {
			return 0, err
		}
	}

	return teamID, err
}

func (s *service) AddCompletion(c *gin.Context, id int, questId int, startYear int, startMonth time.Month, startDay int, startHour int, startMinutes int, startSeconds int) error {

	startTime := time.Date(startYear, startMonth, startDay, startHour, startMinutes, startSeconds, 00, time.UTC).Add(time.Hour * 3)

	actualTime := time.Now()

	if err := s.repo.AddCompletion(c, id, questId, startTime, actualTime); err != nil {
		return err
	}

	return nil
}

func (s *service) GetRanking(c *gin.Context, questId int) ([]domain.QuestTeamCompletionDTO, error) {

	completions, err := s.repo.GetRanking(c, questId)
	if err != nil {
		return nil, err
	}

	sort.Sort(QuestTeamCompletions(completions))

	questTeamCompletionsDTO := make([]domain.QuestTeamCompletionDTO, len(completions))

	for i := range completions {
		team, err := s.repo.GetTeam(c, completions[i].TeamID)
		if err != nil {
			return nil, err
		}

		for j := range team {
			userDTO, _, err := s.userRepo.GetUser(c, team[j].UserID)
			if err != nil {
				return nil, err
			}
			questTeamCompletionsDTO[i].Users = append(questTeamCompletionsDTO[i].Users, userDTO.Username)
		}

		questTeamCompletionsDTO[i].StartTime = completions[i].StartTime
		questTeamCompletionsDTO[i].EndTime = completions[i].EndTime

		diff := questTeamCompletionsDTO[i].EndTime.Sub(questTeamCompletionsDTO[i].StartTime)
		stringDiff := diff.String()

		splitHours := strings.Split(stringDiff, "h")

		var minsSeconds string

		if diff.Hours() < (time.Hour.Hours() * 1) {
			questTeamCompletionsDTO[i].Hours = 0
			minsSeconds = splitHours[0]
		} else {
			hoursFloat, _ := strconv.ParseFloat(splitHours[0], 64)
			questTeamCompletionsDTO[i].Hours = hoursFloat
			minsSeconds = splitHours[1]
		}

		splitMinsSeconds := strings.Split(minsSeconds, "m")

		minutesFloat, _ := strconv.ParseFloat(splitMinsSeconds[0], 64)

		questTeamCompletionsDTO[i].Minutes = minutesFloat

		secondsString := strings.Replace(splitMinsSeconds[1], "s", "", -1)

		secondsFloat, _ := strconv.ParseFloat(secondsString, 64)

		questTeamCompletionsDTO[i].Seconds = secondsFloat

	}

	return questTeamCompletionsDTO, nil
}

func (s *service) DeleteTeam(c *gin.Context, teamId int) error {

	err := s.repo.DeleteTeam(c, teamId)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) AcceptQuestTeam(c *gin.Context, teamId int, userId int) error {

	err := s.repo.AcceptQuestTeam(c, teamId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetWaitRoomAccepted(c *gin.Context, teamId int, questId int) ([]domain.UserDTO, error) {

	waitRoom, err := s.repo.GetWaitRoomAccepted(c, teamId, questId)
	if err != nil {
		return nil, err
	}

	waitRoomAccepted := make([]domain.UserDTO, 0)

	for i := range waitRoom {
		userDTO, _, err := s.userRepo.GetUser(c, waitRoom[i].UserID)
		if err != nil {
			return nil, err
		}
		waitRoomAccepted = append(waitRoomAccepted, userDTO)
		fmt.Println(waitRoomAccepted)
	}

	fmt.Println(waitRoomAccepted)

	return waitRoomAccepted, nil
}

func (s *service) GetTeam(c *gin.Context, teamId int) ([]domain.UserDTO, error) {

	team, err := s.repo.GetTeam(c, teamId)
	if err != nil {
		return nil, err
	}

	teamPlayers := make([]domain.UserDTO, 0)

	for i := range team {
		userDTO, _, err := s.userRepo.GetUser(c, team[i].UserID)
		if err != nil {
			return nil, err
		}
		teamPlayers = append(teamPlayers, userDTO)
		fmt.Println(teamPlayers)
	}

	return teamPlayers, nil
}
