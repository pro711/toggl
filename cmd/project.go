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

var projectCmd = &cobra.Command{
    Use: "project",
    Short: "Manage toggl projects",
}

var projectListCmd = &cobra.Command{
    Use: "list",
    Short: "List toggl projects",
    Run: func(cmd *cobra.Command, args []string) {
        var session = toggl.OpenSession(viper.GetString("token"))
        account, err := session.GetAccount()
        if err != nil {
            log.Fatal(err)
        }
        for _, project := range account.Data.Projects {
            if project.IsActive() {
                fmt.Printf(" - %s (%d)\n", project.Name, project.ID)
            }
        }
    },
}

func init() {
    TogglCommand.AddCommand(projectCmd)
    projectCmd.AddCommand(projectListCmd)
}
