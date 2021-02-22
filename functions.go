package cloudfunctions

import (
	"context"
	"encoding/base64"
	"errors"
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

	project, zone, instance, err := SplitFields(m.Data)
	if err != nil {
		log.Printf("Error processsing message: '%s'", m.Data)
	}

	resp, err := computeService.Instances.Stop(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	println(resp)

	return nil
}

func SplitFields(data string) (project string, zone string, instance string, error error) {
	splitData := strings.Split(ParseMessage(data), ",")
	if len(splitData) < 3 {
		return "", "", "", errors.New("bad message on PubSub topic")
	}

	project = splitData[0]
	zone = splitData[1]
	instance = splitData[2]

	return project, zone, instance, nil
}

func ParseMessage(data string) string {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Printf("Failed to decode message: '%s'", data)
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

	project, zone, instance, err := SplitFields(m.Data)
	if err != nil {
		log.Printf("Error processsing message: '%s'", m.Data)
	}

	resp, err := computeService.Instances.Start(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	println(resp)

	return nil
}
