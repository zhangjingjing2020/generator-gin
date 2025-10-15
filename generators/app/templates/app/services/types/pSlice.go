package types

import "<%= moduleName %>/app/models"

type PlanSlice []models.OrderAutoDispatchPlan

func (p PlanSlice) Len() int {
	return len(p)
}

func (p PlanSlice) Less(i, j int) bool {
	return p[i].StartTime.Before(p[j].StartTime)
}

func (p PlanSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
