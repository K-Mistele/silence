package main

import (
	"errors"
	"fmt"
	"github.com/k-mistele/silence/silence/transport"
	"math"
	"sync"
)
// Database TRACKS SESSIONS
type Database struct {
	sessions 			[math.MaxUint16] *Session
	lock 				sync.RWMutex
}

// NewDatabase CREATES A NEW DATABASE TYPE FOR SESSIONS
func NewDatabase() *Database {
	return &Database{
		sessions: 	[math.MaxUint16] *Session{},
		lock: 		sync.RWMutex{},
	}
}

// Add IS A CONCURRENCY-SAFE WAY TO ADD A SESSION TO A DATABASE
func (d *Database) Add(s *Session) error {

	// ACQUIRE THE LOCK
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.sessions[s.EndpointID] != nil {
		err := fmt.Sprintf("unable to add session with ID %d - there is already an entry there", s.EndpointID)
		return errors.New(err)
	}
	d.sessions[s.EndpointID] = s

	return nil
}

// Exists IS A CONCURRENCY-SAFE WAY TO CHECK IF A SESSION EXISTS
func (d *Database) Exists(endpointID uint16) bool {

	// DON'T NEED TO WORRY ABOUT THE INDEX NOT EXISTING SINCE THE LENGTH IS THE MAX FOR A UINT16
	d.lock.RLock()
	defer d.lock.RUnlock()
	return  d.sessions[endpointID] == nil


}

// Delete IS A CONCURRENCY-SAFE WAY TO DELETE A SESSION
func (d *Database) Delete(endpointID uint16) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.sessions[endpointID] = nil
}

// Get IS A CONCURRENCY-SAFE WAY TO RETRIEVE A RECORD IF IT EXISTS, IT IS READ-ONLY AND RETURNS A COPY
func (d *Database) Get(endpointID uint16) (s Session, exists bool) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	session := d.sessions[endpointID]
	if session == nil {
		return Session{}, false
	}
	return *session, exists
}

// Update IS NON-CONCURRENCY SAFE WAY TO UPDATE SOMETHING - YOU **MUST** MANUALLY UNLOCK WHEN YOU ARE FINISHED
func (d *Database) Update(endpointID uint16) (finishUpdate func() , s *Session, err error) {

	// LOCK FOR READING
	d.lock.RLock()
	defer d.lock.RUnlock()
	s = d.sessions[endpointID]
	if s == nil {
		return func(){}, nil, errors.New("no such session exists")
	}

	// LOCK IT, AND RETURN A CALLBACK TO CALL WHEN YOU'RE DONE THAT'LL UNLOCK
	d.lock.Lock()
	return func(){d.lock.Unlock()}, s, nil

}


// Session REPRESENTS A CONNECTION BETWEEN THE LISTENER AND A HOST.
type Session struct {
	EndpointID		uint16
	ListenerAddress	string
	CurSequenceNumber uint32
	CurAckNumber		uint32
	packetTimeout uint8
	CommandHistory 		[]interface{}				// A LIST OF COMMANDS BEEN RUN
	Fragments			[]transport.Datagram		// A SLICE OF DATAGRAMS THAT ARE WAITING TO BE ASSEMBLED
	Message				chan []byte 				// CHANNEL TO PUSH FINISHED MESSAGES INTO

}
