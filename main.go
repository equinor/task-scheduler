package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "scheduler",
		Usage: "A simple command scheduler",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "command",
				Usage:    "Command to run",
				Required: true,
			},
			&cli.StringFlag{
				Name: "schedule",
				Usage: `Cron schedule. 
Example (every 5 seconds): 0/5 * * * * *

See more https://crontab.cronhub.io/ or https://docs.oracle.com/cd/E12058_01/doc/doc.1014/e12030/cron_expressions.htm`,
				Required: true,
			},
			&cli.StringFlag{
				Name:     "next-command",
				Usage:    "Next command to run, while scheduler is running",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "verbose",
				Usage:    "Verbose output",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			command := c.String("command")
			schedule := c.String("schedule")
			nextCommand := c.String("next-command")
			verbose := c.Bool("verbose")

			cron := cron.New(cron.WithSeconds())
			_, err := cron.AddFunc(schedule, func() {
				runCommand(command, verbose)
			})

			if err != nil {
				return err
			}

			cron.Start()

			if nextCommand != "" {
				runCommand(nextCommand, verbose)
			}
			select {} // Block forever

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(command string, verbose bool) {
	if verbose {
		log.Printf("Running command: %s\n", command)
	}
	commandElements := strings.Split(command, " ")
	if len(commandElements) == 0 {
		log.Printf("Missing command")
		return
	}
	var args []string
	for _, arg := range commandElements[1:] {
		arg = strings.Trim(arg, "\"")
		args = append(args, strings.Trim(arg, "'"))
	}
	cmd := exec.Command(commandElements[0], args...)
	output, err := cmd.Output()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			log.Printf("Command exited with code %d and error: %s\n", exitErr.ExitCode(), exitErr.Stderr)
			return
		}
		log.Printf("Failed to execute command: %v, error: %v\n", command, err)
		return
	}

	log.Printf("%s\n", output)
}
