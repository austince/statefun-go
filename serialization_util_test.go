package statefun_go

import (
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
)

func TestUnmarshalNil(t *testing.T) {
	err := unmarshall(&anypb.Any{}, nil)
	if err == nil {
		assert.Fail(t, "Unmarshal should fail on nil receiver")
	}
}

func TestUnmarshalAny(t *testing.T) {
	value := anypb.Any{
		TypeUrl: "test/type",
		Value:   nil,
	}
	receiver := any.Any{}

	err := unmarshall(&value, &receiver)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, value, receiver)
}

func TestMarshalAny(t *testing.T) {
	value := &anypb.Any{
		TypeUrl: "test/type",
		Value:   nil,
	}

	marshalled, err := marshall(value)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, value, marshalled)
}

func TestMarshalNil(t *testing.T) {
	marshalled, err := marshall(nil)
	if err != nil {
		t.Error(err)
	}

	if marshalled != nil {
		assert.Fail(t, "Marhsalled nil should be nil")
	}
}
