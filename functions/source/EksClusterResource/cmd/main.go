// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws-quickstart/quickstart-amazon-eks-cluster-resource-provider/cmd/resource"
)

// Handler is a container for the CRUDL actions exported by resources
type Handler struct{}

// Create wraps the related Create function exposed by the resource code
func (r *Handler) Create(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Create)
}

// Read wraps the related Read function exposed by the resource code
func (r *Handler) Read(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Read)
}

// Update wraps the related Update function exposed by the resource code
func (r *Handler) Update(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Update)
}

// Delete wraps the related Delete function exposed by the resource code
func (r *Handler) Delete(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Delete)
}

// List wraps the related List function exposed by the resource code
func (r *Handler) List(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.List)
}

// main is the entry point of the application.
func main() {
	fmt.Printf("Starting handler for EKS")
	cfn.Start(&Handler{})
	fmt.Printf("finished starting  handler for EKS")
}

type handlerFunc func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)

func wrap(req handler.Request, f handlerFunc) (response handler.ProgressEvent) {
	defer func() {
		// Catch any panics and return a failed ProgressEvent
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = errors.New(fmt.Sprint(r))
			}

		    fmt.Printf("Trapped error in handler: %v", err)

			response = handler.NewFailedEvent(err)
		}
	}()

	// Populate the previous model
	prevModel := &resource.Model{}
	if err := req.UnmarshalPrevious(prevModel); err != nil {
		fmt.Printf("Error unmarshaling prev model: %v", err)
		return handler.NewFailedEvent(err)
	}

	// Populate the current model
	currentModel := &resource.Model{}
	if err := req.Unmarshal(currentModel); err != nil {
		fmt.Printf("Error unmarshaling model: %v", err)
		return handler.NewFailedEvent(err)
	}

	response, err := f(req, prevModel, currentModel)
	if err != nil {
		fmt.Printf("Error returned from handler function: %v", err)
		return handler.NewFailedEvent(err)
	}

	return response
}
