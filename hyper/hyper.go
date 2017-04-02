package hyper

import (
	"../core/redis"
	"fmt"
	"github.com/hyperhq/hypercli/api/client"
	"github.com/hyperhq/hypercli/cli"
	"github.com/hyperhq/hypercli/pkg/term"
)

func StartBohrium(user string, id string, res string, browser string) {
	stdin, stdout, stderr := term.StdStreams()

	command := []string{
		"run",
		"--rm",
		"-i",
		"-e",
		"USER_ID=" + user,
		"-e",
		"SCEN_ID=" + id,
		"-e",
		"RES_ID=" + res,
		"-e",
		"Browser=" + browser,
		"--name",
		res,
		"fffilimonov/bohrium",
	}

	fmt.Printf("Command: %v\n", command)

	clientCli := client.NewDockerCli(stdin, stdout, stderr, clientFlags)

	c := cli.New(clientCli)
	err := c.Run(command...)
	errText := "nil"
	if err != nil {
		errText = err.Error()
	}
	fmt.Printf("Err: %v\n", err)

	redisConn := redis.Pconnect(3)
	defer redisConn.Close()
	redisConn.HappendValue(user, res, ", {\"Error\": \""+errText+"\"}")
}
