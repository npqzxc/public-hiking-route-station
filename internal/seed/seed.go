package seed

import "public-hiking-route-station/internal/models"

func DefaultRecords() []models.Record {
	return []models.Record{
		{ID: "r1", Title: "Hiking Route Station 首轮安排", Owner: "kai", Status: "进行中", Priority: "高", Note: "先处理高优先级事项。"},
		{ID: "r2", Title: "Hiking Route Station 协同会", Owner: "mila", Status: "待处理", Priority: "中", Note: "同步当前问题和下一步动作。"},
	}
}
