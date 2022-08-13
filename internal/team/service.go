package team

import "github.com/gin-gonic/gin"

type Service interface {
	CreateTeam(c *gin.Context, ids []int) (int, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
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
