package async

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/example/tracers"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracing_log "github.com/opentracing/opentracing-go/log"
)

func startServer() (*machinery.Server, error) {
	cnf, err := config.NewFromYaml("./config.yml", false)
	if err != nil {
		log.ERROR.Println("config failed", err)
	}
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	tasks := map[string]interface{}{
		"sum": Sum,
	}
	return server, server.RegisterTasks(tasks)
}

func Worker() error {
	consumerTag := "machineryDemo"

	cleanup, err := tracers.SetupTracer(consumerTag)
	if err != nil {
		log.FATAL.Fatalln(err)
	}
	defer cleanup()
	server, err := startServer()
	if err != nil {
		return err
	}

	worker := server.NewWorker(consumerTag, 1)
	errorhandler := func(err error) {
		log.ERROR.Println("error handler:", err)
	}
	pretaskhandler := func(signature *tasks.Signature) {
		log.INFO.Println("task handler for:", signature.Name)
	}

	posttaskhandler := func(signature *tasks.Signature) {
		log.INFO.Println("task end handler for:", signature.Name)
	}
	worker.SetPostTaskHandler(posttaskhandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(pretaskhandler)
	return worker.Launch()
}

func Send() error {
	cleanup, err := tracers.SetupTracer("sender")
	if err != nil {
		log.FATAL.Fatalln(err)
	}

	defer cleanup()

	server, err := startServer()

	if err != nil {
		return err
	}

	var (
		addTask tasks.Signature
	)

	var initTasks = func() {
		addTask = tasks.Signature{
			Name: "sum",
			Args: []tasks.Arg{
				{
					Type:  "[]int64",
					Value: []int64{1, 2, 3, 4, 5, 6},
				},
			},
		}
	}
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "send")
	defer span.Finish()

	batchId := uuid.New().String()
	span.SetBaggageItem("batch.id", batchId)
	span.LogFields(opentracing_log.String("batch.id", batchId))

	log.INFO.Println("starting batch:", batchId)
	initTasks()
	asyncResult, err := server.SendTaskWithContext(ctx, &addTask)
	if err != nil {
		return fmt.Errorf("not tasks", err)
	}
	log.INFO.Println(asyncResult)
	return nil
}
