package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Data struct {
	Achievements []interface{} `json:"achievements"`
	CreationTime int64         `json:"creation_time"`
	HighestRank  int           `json:"highest_rank"`
	Pomodoros    []Pomodoro    `json:"pomodoros"`
	Projects     []Project     `json:"projects"`
	Rank         int           `json:"rank"`
	Version      int           `json:"version"`
	WorkUnits    []WorkUnit    `json:"workUnits"`
}

type Pomodoro struct {
	ID        int    `json:"_id"` // ID of the Pomodoro
	Duration  int    `json:"_d"`  // Duration of Pomodoro in seconds
	Timestamp int64  `json:"_t"`  // Timestamp
	Date      string `json:"_dt"`
}

type Project struct {
	ID        int64  `json:"_id"` // ID
	ParentID  int    `json:"_p"`  // Project parent id
	Name      string `json:"_n"`  // Name
	Timestamp int64  `json:"_t"`  // Timestamp
	Duration  int    `json:"_d"`  // ??
}

type WorkUnit struct {
	ID        int64  `json:"_id"` // ID
	ProjectID int64  `json:"_p"`
	Duration  int64  `json:"_d"`  // Duration of Pomodoro in seconds
	Date      string `json:"_dt"` // Date human readale
	Timestamp int64  `json:"_t"`  // Timestamp
}

type Pro struct {
	ID       int64
	Name     string
	Week     string
	Pomodoro int64
}

func main() {
	plan, _ := ioutil.ReadFile("backup.json")
	var data Data
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println(err)
	}

	projects := map[int64]Project{}
	for _, v := range data.Projects {
		projects[v.ID] = v
	}

	results := map[int64]map[string]*Pro{}
	for _, wu := range data.WorkUnits {
		po := projects[wu.ProjectID]
		tm := time.Unix(wu.Timestamp, 0)
		begin := tm

		// 1 = Monday
		weekDay := int(tm.Weekday())
		if weekDay > 1 {
			begin = tm.AddDate(0, 0, -(weekDay - 1))
		}

		if weekDay == 0 {
			begin = tm.AddDate(0, 0, weekDay+1)
		}

		bTime := begin.Format("2006-01-02")
		if _, ok := results[po.ID]; !ok {
			results[po.ID] = map[string]*Pro{}
		}

		if _, ok := results[po.ID][bTime]; !ok {
			results[po.ID][bTime] = &Pro{
				ID:       po.ID,
				Name:     po.Name,
				Pomodoro: 0,
			}
		}

		pro := results[po.ID][bTime]
		pro.Pomodoro += 1
	}

	for _, v := range results {
		for k, w := range v {
			fmt.Println(k, ":", w.Name, ":", w.Pomodoro)
		}
	}
}
