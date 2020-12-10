package models

import (
	"fmt"
	"strconv"
)

func TeacherScore(project string, goal, completed, nextMonthGoal string) (rate, score float64) {
	goal1, _ := strconv.ParseFloat(goal, 64)
	completed1, _ := strconv.ParseFloat(completed, 64)

	switch project {
	case "1":
		rate = completed1 / goal1
		score = 30 * rate
		break
	case "2":
		rate = completed1 / goal1
		score = 30 * rate
		break
	case "3":
		rate = completed1 / goal1
		score = 20 * rate
		break
	case "4":
		rate = completed1 / goal1
		score = 20 * rate
		break
	default:
		rate = 0
		score = 0
		break
	}
	rate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", rate), 64)
	score, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", score), 64)
	return
}

func SalesExecutive(project string, goal, completed string) (rate, score float64) {
	goal1, _ := strconv.ParseFloat(goal, 64)
	completed1, _ := strconv.ParseFloat(completed, 64)

	switch project {
	case "1":
		rate = completed1 / goal1
		score = 40 * rate
		break
	case "2":
		rate = completed1 / goal1
		score = 30 * rate
		break
	case "3":
		rate = completed1 / goal1
		score = 30 * rate
		break
	default:
		rate = 0
		score = 0
		break
	}
	return
}

func CourseConsultant(project string, goal, completed string) (rate, score float64) {
	goal1, _ := strconv.ParseFloat(goal, 64)
	completed1, _ := strconv.ParseFloat(completed, 64)

	switch project {
	case "1":
		rate = completed1 / goal1
		score = 30 * rate
		break
	case "2":
		rate = completed1 / goal1
		score = 25 * rate
		break
	case "3":
		rate = completed1 / goal1
		score = 25 * rate
		break
	case "4":
		rate = completed1 / goal1
		score = 20 * rate
		break
	default:
		rate = 0
		score = 0
		break
	}
	return
}
func Marketing(project string, goal, completed string) (rate, score float64) {
	goal1, _ := strconv.ParseFloat(goal, 64)
	completed1, _ := strconv.ParseFloat(completed, 64)

	switch project {
	case "1":
		rate = completed1 / goal1
		score = 40 * rate
		break
	case "2":
		rate = completed1 / goal1
		score = 60 * rate
		break
	default:
		rate = 0
		score = 0
		break
	}
	return
}
func ProductManager(project string, goal, completed string) (rate, score float64) {
	goal1, _ := strconv.ParseFloat(goal, 64)
	completed1, _ := strconv.ParseFloat(completed, 64)

	switch project {
	case "1":
		rate = completed1 / goal1
		score = 30 * rate
		break
	case "2":
		rate = completed1 / goal1
		score = 30 * rate
		break
	case "3":
		rate = completed1 / goal1
		score = 20 * rate
		break
	case "4":
		rate = completed1 / goal1
		score = 20 * rate
		break
	default:
		rate = 0
		score = 0
		break
	}
	return
}

