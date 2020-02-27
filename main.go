package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gofiber/fiber"
	"github.com/gofiber/middlewares"
	"os"
)

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String("test-tracking"),
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(*result.TopicArn)

	input := &sns.PublishInput{
		Message:  aws.String("Hello world!"),
		TopicArn: aws.String(*result.TopicArn),
	}

	response, err := svc.Publish(input)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}

	fmt.Println(response)
	// Create new Fiber instance
	app := fiber.New()

	// Enable prefork
	app.Settings.Prefork = true

	// Enable Logger
	app.Use(middleware.Logger())

	// Create new GET route on path "/hello"
	app.Get("/hello", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	// Listen on port 3000
	app.Listen(3000)

	// Run the following command to see all processes sharing port 3000:
	// sudo lsof -i -P -n | grep LISTEN
}
