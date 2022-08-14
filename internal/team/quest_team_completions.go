package team

import "github.com/proyecto-final-2022/geoquest-backend/internal/domain"

type QuestTeamCompletions []domain.QuestTeamCompletion

func (q QuestTeamCompletions) Len() int { return len(q) }
func (q QuestTeamCompletions) Less(i, j int) bool {
	return q[i].EndTime.Sub(q[i].StartTime) < q[j].EndTime.Sub(q[j].StartTime)
}
func (q QuestTeamCompletions) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
