package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type SalaryInfo struct {
	jobTitle        string
	experienceLevel string
	workModel       string
	workYear        int
	salaryUSD       int
}

type JobMedian struct {
	jobTitle string
	median   int
}

func main() {
	allJobs := readSalaries("salaries.csv")
	partOne(allJobs)
	partTwo(allJobs)
	partThree(allJobs)
	partFour(allJobs)
	partFive(allJobs)
}

func getFile(name string) []string {
	contents, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	bigString := string(contents)
	allLines := strings.Split(bigString, "\n")
	return allLines
}

func readSalaries(fileName string) []SalaryInfo {
	allLines := getFile(fileName)
	allJobs := make([]SalaryInfo, 0)
	for i, line := range allLines {
		if i > 0 {
			parts := strings.Split(line, ",")
			nextJob := SalaryInfo{}
			nextJob.jobTitle = parts[0]
			nextJob.experienceLevel = parts[1]
			nextJob.workModel = parts[3]
			workYear, _ := strconv.Atoi(parts[4])
			nextJob.workYear = workYear
			salaryUSD, _ := strconv.Atoi(parts[8])
			nextJob.salaryUSD = salaryUSD
			allJobs = append(allJobs, nextJob)
		}
	}
	return allJobs
}

func getSalaries(allJobs []SalaryInfo, year int, level string) []int {
	salaries := make([]int, 0)
	for _, job := range allJobs {
		if job.workYear == year && job.experienceLevel == level {
			salaries = append(salaries, job.salaryUSD)
		}
	}
	return salaries
}

func getWorkModelSalaries(allJobs []SalaryInfo, year int, level string, workModel string) []int {
	salaries := make([]int, 0)
	for _, job := range allJobs {
		if job.workYear == year && job.experienceLevel == level && job.workModel == workModel {
			salaries = append(salaries, job.salaryUSD)
		}
	}
	return salaries
}

func getAllLevelsWorkModelSalaries(allJobs []SalaryInfo, year int, workModel string) []int {
	salaries := make([]int, 0)
	for _, job := range allJobs {
		if job.workYear == year && job.workModel == workModel {
			salaries = append(salaries, job.salaryUSD)
		}
	}
	return salaries
}

func average(salaries []int) int {
	total := 0
	for _, salary := range salaries {
		total = total + salary
	}
	return total / len(salaries)
}

func median(salaries []int) int {
	slices.Sort(salaries)
	middle := len(salaries) / 2
	if len(salaries)%2 == 0 {
		return (salaries[middle-1] + salaries[middle]) / 2
	}
	return salaries[middle]
}

func partOne(allJobs []SalaryInfo) {
	fmt.Println("Part 1: Average entry level salary, 2024 vs 2023")
	average2024 := average(getSalaries(allJobs, 2024, "Entry-level"))
	average2023 := average(getSalaries(allJobs, 2023, "Entry-level"))
	if average2024 < average2023 {
		color.Red("2024 average entry level salary: $%d", average2024)
		color.Red("2023 average entry level salary: $%d", average2023)
		color.Red("The 2024 average is lower than 2023")
	} else {
		color.Green("2024 average entry level salary: $%d", average2024)
		color.Green("2023 average entry level salary: $%d", average2023)
		color.Green("The 2024 average is the same or higher than 2023")
	}
}

func partTwo(allJobs []SalaryInfo) {
	fmt.Println("Part 2: Median entry level salary, 2024 vs 2023")
	median2024 := median(getSalaries(allJobs, 2024, "Entry-level"))
	median2023 := median(getSalaries(allJobs, 2023, "Entry-level"))
	if median2023 > median2024 {
		color.Red("2024 median entry level salary: $%d", median2024)
		color.Red("2023 median entry level salary: $%d", median2023)
		color.Red("The median went down by $%d from 2023 to 2024", median2023-median2024)
	} else {
		color.Green("2024 median entry level salary: $%d", median2024)
		color.Green("2023 median entry level salary: $%d", median2023)
		color.Green("The median went up by $%d from 2023 to 2024", median2024-median2023)
	}

	average2024 := average(getSalaries(allJobs, 2024, "Entry-level"))
	average2023 := average(getSalaries(allJobs, 2023, "Entry-level"))
	fmt.Println("Average vs median for entry level jobs:")
	color.Cyan("2024: the average is $%d, that's $%d higher than the median of $%d", average2024, average2024-median2024, median2024)
	color.Cyan("2023: the average is $%d, that's $%d higher than the median of $%d", average2023, average2023-median2023, median2023)
}

func partThree(allJobs []SalaryInfo) {
	fmt.Println("Part 3: Entry level vs mid level, 2024")
	entryMedian := median(getSalaries(allJobs, 2024, "Entry-level"))
	midMedian := median(getSalaries(allJobs, 2024, "Mid-level"))
	color.Cyan("2024 median entry level salary: $%d", entryMedian)
	color.Cyan("2024 median mid level salary:   $%d", midMedian)
	color.Cyan("Moving from entry level to mid level could get you about a $%d raise", midMedian-entryMedian)
}

func partFour(allJobs []SalaryInfo) {
	fmt.Println("Part 4: Remote vs on-site medians, 2024")
	remoteEntry := median(getWorkModelSalaries(allJobs, 2024, "Entry-level", "Remote"))
	onsiteEntry := median(getWorkModelSalaries(allJobs, 2024, "Entry-level", "On-site"))
	remoteAll := median(getAllLevelsWorkModelSalaries(allJobs, 2024, "Remote"))
	onsiteAll := median(getAllLevelsWorkModelSalaries(allJobs, 2024, "On-site"))
	printRemoteVsOnsite("Entry level", remoteEntry, onsiteEntry)
	printRemoteVsOnsite("All levels", remoteAll, onsiteAll)
}

func printRemoteVsOnsite(label string, remoteMedian int, onsiteMedian int) {
	if onsiteMedian > remoteMedian {
		color.Red("%s - fully remote median: $%d", label, remoteMedian)
		color.Red("%s - on-site median:      $%d", label, onsiteMedian)
	} else {
		color.Green("%s - fully remote median: $%d", label, remoteMedian)
		color.Green("%s - on-site median:      $%d", label, onsiteMedian)
	}
}

func partFive(allJobs []SalaryInfo) {
	fmt.Println("Part 5: Highest paying entry level jobs in 2024, not counting Research jobs")
	titleSalaries := make(map[string][]int)
	for _, job := range allJobs {
		if job.workYear == 2024 && job.experienceLevel == "Entry-level" && !strings.Contains(job.jobTitle, "Research") {
			titleSalaries[job.jobTitle] = append(titleSalaries[job.jobTitle], job.salaryUSD)
		}
	}

	jobMedians := make([]JobMedian, 0)
	for title, salaries := range titleSalaries {
		nextJob := JobMedian{}
		nextJob.jobTitle = title
		nextJob.median = median(salaries)
		jobMedians = append(jobMedians, nextJob)
	}
	slices.SortFunc(jobMedians, func(a, b JobMedian) int {
		return b.median - a.median
	})

	color.Green("#1 highest paying entry level job: %s, median $%d", jobMedians[0].jobTitle, jobMedians[0].median)
	fmt.Println("The next 4 highest paying:")
	for i := 1; i < 5; i++ {
		color.Yellow("#%d %s, median $%d", i+1, jobMedians[i].jobTitle, jobMedians[i].median)
	}
}
