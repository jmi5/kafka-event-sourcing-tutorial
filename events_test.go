package main_test

import (
    // . "banku"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/satori/go.uuid"
)

type Event struct {
    AccId string
    Type  string
}

type CreateEvent struct {
    Event
    AccName string
}

func NewCreateAccountEvent(name string) CreateEvent {
    event := new(CreateEvent)
    event.Type = "CreateEvent"
    event.AccId = uuid.NewV4().String()
    event.AccName = name
    return *event
}

var _ = Describe("Event", func() {
    Describe("NewCreateAccountEvent", func() {
        It("can create a create account event", func() {
            name := "John Smith"

            event := NewCreateAccountEvent(name)

            Expect(event.AccName).To(Equal(name))
            Expect(event.AccId).NotTo(BeNil())
            Expect(event.Type).To(Equal("CreateEvent"))
        })
    })
})
