package hooks

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/apex/log"
)

var NAME = "Gradle Publisher"
var PVERSION = "dev"

type GradlePublisher struct {
	CMD string
}

func (gp *GradlePublisher) Init(m map[string]string) error {
	log.Infof("Init %v", m)
	gp.CMD = m["cmd"]

	return nil
}

func (gp *GradlePublisher) Name() string {
	return NAME
}

func (gp *GradlePublisher) Version() string {
	return PVERSION
}

func (gp *GradlePublisher) Publish(newRelease string) error {
	log.Infof("new version: " + newRelease)
	if err := gp.gradlePublish(); err != nil {
		return err
	}
	return nil
}

func (gp *GradlePublisher) gradlePublish() error {
	log.Infof("Start gradle publish...")

	cmd := gp.CMD
	cmdArgs := strings.Fields(cmd)

	if len(gp.CMD) == 0 {
		cmdArgs = append(cmdArgs, "./gradlew", "publish")
	}

	cmdPipe := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	log.Infof("Command: %s %s", cmdArgs[0], strings.Join(cmdArgs[1:], " "))

	stdout, err := cmdPipe.StdoutPipe()
	if err != nil {
		log.Infof("error oucring when publishing. Detail: %s", err.Error())
		return err
	}
	if err := cmdPipe.Start(); err != nil {
		log.Infof("error oucring when publishing. Detail: %s", err.Error())
		return err
	}

	// print the output of the subprocess
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		log.Infof(m)
	}

	if err := cmdPipe.Wait(); err != nil {
		log.Infof("error oucring when publishing. Detail: %s", err.Error())
		return err
	}
	return nil
}
