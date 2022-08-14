package team

import (
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"
)

type Service interface {
	CreateTeam(c *gin.Context, ids []int) (int, error)
	AddCompletion(c *gin.Context, id int, questId int, startYear int, startMonth time.Month, startDay int, startHour int, startMinutes int, startSeconds int) error
	GetRanking(c *gin.Context, questId int) ([]domain.QuestTeamCompletionDTO, error)
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

func (s *service) CreateTeam(c *gin.Context, ids []int) (int, error) {

	teamID, err := s.repo.CreateTeam(c)

	if err != nil {
		return 0, err
	}

	for i := range ids {
		err = s.repo.AddPlayer(c, teamID, ids[i])
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
			questTeamCompletionsDTO[i].Users = append(questTeamCompletionsDTO[i].Users, userDTO)
		}
		questTeamCompletionsDTO[i].StartTime = completions[i].StartTime
		questTeamCompletionsDTO[i].EndTime = completions[i].EndTime
	}

	return questTeamCompletionsDTO, nil
}
