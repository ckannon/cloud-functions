package cloudfunctions

import (
	"context"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"log"
	"strings"
)

func ProcessStopMessage(ctx context.Context, m PubSubMessage) error {
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	splitData := strings.Split(ParseMessage(m.Data), ",")
	if len(splitData) < 3 {
		log.Println(fmt.Sprint("Bad message on bus: '%s'", m.Data))
		return nil
	}

	project := splitData[0]
	zone := splitData[1]
	instance := splitData[2]

	resp, err := computeService.Instances.Stop(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	println(resp)

	return nil
}

func ParseMessage(data string) string {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Println(fmt.Sprint("Failed to decode message: '%s'", data))
	}
	return string(decoded)
}

func ProcessStartMessage(ctx context.Context, m PubSubMessage) error {
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	splitData := strings.Split(ParseMessage(m.Data), ",")
	if len(splitData) < 3 {
		log.Println(fmt.Sprint("Bad message on bus: '%s'", m.Data))
		return nil
	}

	project := splitData[0]
	zone := splitData[1]
	instance := splitData[2]

	resp, err := computeService.Instances.Start(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	println(resp)

	return nil
}
