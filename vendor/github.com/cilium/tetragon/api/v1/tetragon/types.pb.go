// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// Code generated by protoc-gen-go-tetragon. DO NOT EDIT

package tetragon

// IsGetEventsResponse_Event encapulates isGetEventsResponse_Event
type IsGetEventsResponse_Event = isGetEventsResponse_Event

// Event represents a Tetragon event
type Event interface {
	Encapsulate() IsGetEventsResponse_Event
}

// ProcessEvent represents a Tetragon event that has a Process field
type ProcessEvent interface {
	Event
	SetProcess(p *Process)
}

// ParentEvent represents a Tetragon event that has a Parent field
type ParentEvent interface {
	Event
	SetParent(p *Process)
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessExec) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessExec{
		ProcessExec: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessExec) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessExec) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessExit) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessExit{
		ProcessExit: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessExit) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessExit) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessKprobe) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessKprobe{
		ProcessKprobe: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessKprobe) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessKprobe) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *ProcessTracepoint) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_ProcessTracepoint{
		ProcessTracepoint: event,
	}
}

// SetProcess implements the ProcessEvent interface.
// Sets the Process field of an event.
func (event *ProcessTracepoint) SetProcess(p *Process) {
	event.Process = p
}

// SetParent implements the ParentEvent interface.
// Sets the Parent field of an event.
func (event *ProcessTracepoint) SetParent(p *Process) {
	event.Parent = p
}

// Encapsulate implements the Event interface.
// Returns the event wrapped by its GetEventsResponse_* type.
func (event *Test) Encapsulate() IsGetEventsResponse_Event {
	return &GetEventsResponse_Test{
		Test: event,
	}
}
