package quest

import "github.com/proyecto-final-2022/geoquest-backend/internal/domain"

type QuestsCompletions []domain.QuestCompletion

func (q QuestsCompletions) Len() int { return len(q) }
func (q QuestsCompletions) Less(i, j int) bool {
	return q[i].EndTime.Sub(q[i].StartTime) < q[j].EndTime.Sub(q[j].StartTime)
}
func (q QuestsCompletions) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

type QuestsProgresses []domain.QuestProgressDTO

func (q QuestsProgresses) Len() int { return len(q) }
func (q QuestsProgresses) Less(i, j int) bool {
	return q[i].Points > q[j].Points
}
func (q QuestsProgresses) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
