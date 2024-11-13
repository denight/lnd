package wtwire

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

// prefixWithMsgType takes []byte and adds a wire protocol prefix
// to make the []byte into an actual message to be used in fuzzing.
func prefixWithMsgType(data []byte, prefix MessageType) []byte {
	var prefixBytes [2]byte
	binary.BigEndian.PutUint16(prefixBytes[:], uint16(prefix))
	data = append(prefixBytes[:], data...)

	return data
}

// wireMsgHarness performs the actual fuzz testing of the appropriate wire
// message. This function will check that the passed-in message passes wire
// length checks, is a valid message once deserialized, and passes a sequence of
// serialization and deserialization checks.
func wireMsgHarness(t *testing.T, data []byte, emptyMsg Message) {
	t.Helper()

	// Make sure byte array length is less than max payload size for the
	// wire message.
	payloadLen := uint32(len(data))
	if payloadLen > emptyMsg.MaxPayloadLength(0) {
		// Ignore this input - max payload constraint violated.
		return
	}

	data = prefixWithMsgType(data, emptyMsg.MsgType())

	// Create a reader with the byte array.
	r := bytes.NewReader(data)

	msg, err := ReadMessage(r, 0)
	if err != nil {
		return
	}

	// We will serialize the message into a new bytes buffer.
	var b bytes.Buffer
	_, err = WriteMessage(&b, msg, 0)
	require.NoError(t, err)

	// Deserialize the message from the serialized bytes buffer, and then
	// assert that the original message is equal to the newly deserialized
	// message.
	newMsg, err := ReadMessage(&b, 0)
	require.NoError(t, err)
	require.Equal(t, msg, newMsg)
}

func FuzzCreateSessionReply(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create an empty message so that the FuzzHarness func can
		// check if the max payload constraint is violated.
		emptyMsg := CreateSessionReply{}

		wireMsgHarness(t, data, &emptyMsg)
	})
}

func FuzzCreateSession(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create an empty message so that the FuzzHarness func can
		// check if the max payload constraint is violated.
		emptyMsg := CreateSession{}

		wireMsgHarness(t, data, &emptyMsg)
	})
}

func FuzzDeleteSessionReply(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create an empty message so that the FuzzHarness func can
		// check if the max payload constraint is violated.
		emptyMsg := DeleteSessionReply{}

		wireMsgHarness(t, data, &emptyMsg)
	})
}

func FuzzDeleteSession(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create an empty message so that the FuzzHarness func can
		// check if the max payload constraint is violated.
		emptyMsg := DeleteSession{}

		wireMsgHarness(t, data, &emptyMsg)
	})
}

func FuzzError(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create an empty message so that the FuzzHarness func can
		// check if the max payload constraint is violated.
		emptyMsg := Error{}

		wireMsgHarness(t, data, &emptyMsg)
	})
}

func FuzzInit(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create an empty message so that the FuzzHarness func can
		// check if the max payload constraint is violated.
		emptyMsg := Init{}

		wireMsgHarness(t, data, &emptyMsg)
	})
}

func FuzzStateUpdateReply(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create an empty message so that the FuzzHarness func can
		// check if the max payload constraint is violated.
		emptyMsg := StateUpdateReply{}

		wireMsgHarness(t, data, &emptyMsg)
	})
}

func FuzzStateUpdate(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create an empty message so that the FuzzHarness func can
		// check if the max payload constraint is violated.
		emptyMsg := StateUpdate{}

		wireMsgHarness(t, data, &emptyMsg)
	})
}
