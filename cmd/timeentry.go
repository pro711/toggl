// Copyright (c) 2017 Tao Chen <pro711@gmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package cmd

import (
    "fmt"
    "log"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    toggl "github.com/jason0x43/go-toggl"
)

func getProjectID(account toggl.Account, proj string) (int, error) {
    for _, project := range account.Data.Projects {
        if project.Name == proj {
            return project.ID, nil
        }
    }
    return -1, fmt.Errorf("Project not found: %s", proj)
}

func getCurrentTimeEntry(account toggl.Account) (*toggl.TimeEntry, error) {
    for _, timeEntry := range account.Data.TimeEntries {
        if timeEntry.Stop == nil {
            return &timeEntry, nil
        }
    }
    return nil, fmt.Errorf("Current time entry not found")
}

func startTimeEntry(desc string, proj string) error {
    var session = toggl.OpenSession(viper.GetString("token"))

    if proj == "" {
        if _, err := session.StartTimeEntry(desc); err != nil {
            return err
        }
    } else {
        var projID int

        account, err := session.GetAccount()
        if err != nil {
            return err
        }
        if projID, err = getProjectID(account, proj); err != nil {
            return err
        }
        if _, err = session.StartTimeEntryForProject(desc, projID); err != nil {
            return err
        }
    }
    return nil
}

func stopCurrentTimeEntry() error {
    var session = toggl.OpenSession(viper.GetString("token"))
    account, err := session.GetAccount()
    if err != nil {
        return err
    }
    timeEntry, err := getCurrentTimeEntry(account)
    if timeEntry != nil {
        session.StopTimeEntry(*timeEntry)
    }
    return nil
}

var project string

var startTimeEntryCmd = &cobra.Command{
    Use: "start [flags] description",
    Short: "Start a new time entry",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) != 1 {
            log.Fatal("Error: toggl start takes exactly one argument.")
        }
        startTimeEntry(args[0], project)
    },
}

var stopTimeEntryCmd = &cobra.Command{
    Use: "stop",
    Short: "Stop the current time entry",
    Run: func(cmd *cobra.Command, args []string) {
        stopCurrentTimeEntry()
    },
}

var timeEntryCommands = [...]*cobra.Command{
    startTimeEntryCmd,
    stopTimeEntryCmd,
}

func init() {
    TogglCommand.AddCommand(startTimeEntryCmd)
    TogglCommand.AddCommand(stopTimeEntryCmd)
    for _, cmd := range timeEntryCommands {
        cmd.Flags().StringVarP(&project, "project", "P", "", "Project")
    }
}
