package main

import (
	"github.com/apex/log"
	publisherGralde "github.com/ted-vo/publisher-gradle/pkg/publisher"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
	"github.com/ted-vo/semantic-release/v3/pkg/publisher"
)

func main() {
	log.SetHandler(publisherGralde.NewLogHandler())
	plugin.Serve(&plugin.ServeOpts{
		Publisher: func() publisher.Publisher {
			return &publisherGralde.GradlePublisher{}
		},
	})
}
