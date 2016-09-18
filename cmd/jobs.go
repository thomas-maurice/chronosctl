package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thomas-maurice/chronosctl/client"
	"github.com/thomas-maurice/chronosctl/types"
	"strings"
)

var (
	CPUs               float64
	RAM                float64
	Disk               float64
	Command            string
	Schedule           string
	Epsilon            string
	Owner              string
	OwnerName          string
	Description        string
	ContainerType      string
	ContainerImage     string
	ContainerNetwork   string
	ContainerForcePull bool
	Parents            string
	Async              bool
	Environment        string
)

var JobCmd = &cobra.Command{
	Use:   "job",
	Short: "Perform actions on the jobs",
	Long:  ``,
}

var JobListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the jobs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cli := client.NewClient(DebugMode)
		var jobs []types.ChronosJob
		_, err := cli.Get("/scheduler/jobs", &jobs, 200)
		if err != nil {
			logrus.Fatalf("Could not get job list: %s", err)
		}
		statuses, err := cli.GetJobsStatus()
		if err != nil {
			logrus.Fatalf("Could not get job statuses: %s", err)
		}

		for _, job := range jobs {
			for _, status := range statuses {
				if job.Name == status.Name {
					fmt.Printf("%-25s | %-40s | CPU: %5f | RAM: %5f | %5d fail | %5d ok | %10s | %10s\n", job.Name, job.Schedule, job.CPUs, job.Mem, job.ErrorCount, job.SuccessCount, status.Status, status.LastOutcome)
					continue
				}
			}
		}
	},
}

var JobRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logrus.Fatalf("Please provide at least a job name")
		}
		cli := client.NewClient(DebugMode)
		for _, job := range args {
			_, err := cli.Put("/scheduler/job/"+job, nil, nil, nil, 204)
			if err != nil {
				logrus.Errorf("Could not run job %s: %s", job, err)
			} else {
				logrus.Infof("Running job %s", job)
			}
		}
	},
}

var JobKillCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kills all the tasks for a job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logrus.Fatalf("Please provide at least a job name")
		}
		cli := client.NewClient(DebugMode)
		for _, job := range args {
			_, err := cli.Delete("/scheduler/task/kill/"+job, nil, 204)
			if err != nil {
				logrus.Errorf("Could not kill job %s: %s", job, err)
			} else {
				logrus.Infof("Killed tasks for job %s", job)
			}
		}
	},
}

var JobDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logrus.Fatalf("Please provide at least a job name")
		}
		cli := client.NewClient(DebugMode)
		for _, job := range args {
			_, err := cli.Delete("/scheduler/job/"+job, nil, 204)
			if err != nil {
				logrus.Errorf("Could not delete job %s: %s", job, err)
			} else {
				logrus.Infof("Deleted job %s", job)
			}
		}
	},
}

var JobShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows a job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logrus.Fatalf("Please provide a job name")
		}
		cli := client.NewClient(DebugMode)
		var jobs []types.ChronosJob
		_, err := cli.Get("/scheduler/jobs", &jobs, 200)
		if err != nil {
			logrus.Fatalf("Could not get job list: %s", err)
		}
		for _, job := range jobs {
			if job.Name == args[0] {
				b, err := json.MarshalIndent(job, "", "  ")
				if err != nil {
					logrus.Fatalf("Could not marshal structure: %s", err)
				}
				fmt.Println(string(b))
				return
			}
		}
		logrus.Fatalf("Could not find job: %s", args[0])
	},
}

var JobCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the job",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logrus.Fatalf("Please provide a job name")
		}

		if Schedule != "" && Parents != "" {
			logrus.Fatal("Please specify only one of --parents and --schedule")
		}

		if Schedule == "" && Parents == "" {
			logrus.Fatal("Please specify one of --parents or --schedule")
		}

		split := strings.Split(Environment, ",")
		var environ []types.EnvironmentVariable
		for _, env := range split {
			split := strings.Split(env, "=")
			if len(split) == 1 && split[0] != "" {
				environ = append(environ, types.EnvironmentVariable{Name: split[0], Value: ""})
			} else if len(split) == 2 {
				environ = append(environ, types.EnvironmentVariable{Name: split[0], Value: split[1]})
			}
		}

		var environment *[]types.EnvironmentVariable
		if len(environ) == 0 {
			environment = nil
		} else {
			environment = &environ
		}

		var deps *[]string
		arr := strings.Split(Parents, ",")
		if Parents != "" {
			deps = &arr
		} else {
			deps = nil
		}

		job := types.NewChronosJob{
			CPUs:                 CPUs,
			Disk:                 Disk,
			Mem:                  RAM,
			Async:                Async,
			Command:              Command,
			Schedule:             Schedule,
			Epsilon:              Epsilon,
			Owner:                Owner,
			OwnerName:            OwnerName,
			Description:          Description,
			Parents:              deps,
			Name:                 args[0],
			EnvironmentVariables: environment,
		}

		if ContainerType != "" {
			if ContainerImage == "" {
				logrus.Fatalf("Image cannot be an empty string")
			}
			job.Container = &types.ChronosContainerOptions{
				Type:           ContainerType,
				Image:          ContainerImage,
				Network:        ContainerNetwork,
				ForcePullImage: ContainerForcePull,
			}
		}

		cli := client.NewClient(DebugMode)
		if Schedule != "" {
			_, err := cli.Post("/scheduler/iso8601", job, nil, nil, 204)
			if err != nil {
				logrus.Fatalf("Could not create job: %s", err)
			}
		} else if Parents != "" {
			_, err := cli.Post("/scheduler/dependency", job, nil, nil, 204)
			if err != nil {
				logrus.Fatalf("Could not create job: %s", err)
			}
		} else {
			logrus.Fatalf("--schedule and --parents were empty, I don't know what to do :(")
		}
	},
}

func InitJobCmd() {
	JobCmd.AddCommand(JobListCmd)
	JobCmd.AddCommand(JobShowCmd)
	JobCmd.AddCommand(JobRunCmd)
	JobCmd.AddCommand(JobKillCmd)
	JobCmd.AddCommand(JobDeleteCmd)
	JobCmd.AddCommand(JobCreateCmd)

	JobCreateCmd.Flags().Float64VarP(&CPUs, "cpus", "c", 0.1, "Allocated CPUs for the task")
	JobCreateCmd.Flags().BoolVarP(&Async, "async", "a", false, "Is the task async ?")
	JobCreateCmd.Flags().Float64VarP(&RAM, "ram", "r", 64, "Allocated RAM for the task")
	JobCreateCmd.Flags().StringVarP(&Environment, "environment", "", "", "Environment variables")
	JobCreateCmd.Flags().Float64VarP(&Disk, "disk", "", 0, "Allocated disk for the task")
	JobCreateCmd.Flags().StringVarP(&Command, "command", "", "", "Command to launch")
	JobCreateCmd.Flags().StringVarP(&Schedule, "schedule", "s", "", "Schedule for the task")
	JobCreateCmd.Flags().StringVarP(&Epsilon, "epsilon", "e", "", "Epsilon for the task")
	JobCreateCmd.Flags().StringVarP(&Owner, "owner", "o", "", "Owner's email for the task")
	JobCreateCmd.Flags().StringVarP(&Description, "description", "", "", "Task description")
	JobCreateCmd.Flags().StringVarP(&OwnerName, "owner-name", "", "", "Owner's name for the task")
	JobCreateCmd.Flags().StringVarP(&Parents, "parents", "p", "", "Comma separated list of jobs this job depends on")
	JobCreateCmd.Flags().StringVarP(&ContainerType, "container", "", "", "Container type to use")
	JobCreateCmd.Flags().StringVarP(&ContainerImage, "container-image", "", "", "Container image to use")
	JobCreateCmd.Flags().StringVarP(&ContainerNetwork, "container-network", "", "BRIDGE", "Container network mode to use")
	JobCreateCmd.Flags().BoolVarP(&ContainerForcePull, "container-force-pull", "", false, "Force pull the image")

}
