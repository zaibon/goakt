package actors

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tochemey/goakt/config"
)

func TestNewActorSystem(t *testing.T) {
	t.Run("With Defaults", func(t *testing.T) {
		cfg, err := config.New("testSys", "localhost:0")
		require.NoError(t, err)
		assert.NotNil(t, cfg)

		actorSys, err := NewActorSystem(cfg)
		require.NoError(t, err)
		require.NotNil(t, actorSys)
		var iface any = actorSys
		_, ok := iface.(ActorSystem)
		assert.True(t, ok)
		assert.Equal(t, "testSys", actorSys.Name())
		assert.Empty(t, actorSys.Actors())
		assert.Equal(t, "localhost:0", actorSys.NodeAddr())
	})

	t.Run("With Missing Config", func(t *testing.T) {
		sys, err := NewActorSystem(nil)
		assert.Error(t, err)
		assert.Nil(t, sys)
		assert.EqualError(t, err, ErrMissingConfig.Error())
	})

	t.Run("With Spawn an actor when not started", func(t *testing.T) {
		ctx := context.TODO()
		cfg, _ := config.New("testSys", "localhost:0")
		sys, _ := NewActorSystem(cfg)
		kind := "TestActor"
		id := "test-1"
		actor := NewTestActor(id)
		actorRef := sys.Spawn(ctx, kind, actor)
		assert.Nil(t, actorRef)
	})

	t.Run("With Spawn an actor when started", func(t *testing.T) {
		ctx := context.TODO()
		cfg, _ := config.New("testSys", "localhost:0")
		sys, _ := NewActorSystem(cfg)

		// start the actor system
		err := sys.Start(ctx)
		assert.NoError(t, err)

		kind := "TestActor"
		id := "test-1"
		actor := NewTestActor(id)
		actorRef := sys.Spawn(ctx, kind, actor)
		assert.NotNil(t, actorRef)

		assert.NoError(t, sys.Stop(ctx))
	})

	t.Run("With Spawn an actor already exist", func(t *testing.T) {
		ctx := context.TODO()
		cfg, _ := config.New("testSys", "localhost:0")
		sys, _ := NewActorSystem(cfg)

		// start the actor system
		err := sys.Start(ctx)
		assert.NoError(t, err)

		kind := "TestActor"
		id := "test-1"
		actor := NewTestActor(id)
		ref1 := sys.Spawn(ctx, kind, actor)
		assert.NotNil(t, ref1)

		ref2 := sys.Spawn(ctx, kind, actor)
		assert.NotNil(t, ref2)

		// point to the same memory address
		assert.True(t, ref1 == ref2)

		assert.NoError(t, sys.Stop(ctx))
	})

	t.Run("With housekeeping", func(t *testing.T) {
		ctx := context.TODO()

		cfg, _ := config.New("testSys", "localhost:0", config.WithExpireActorAfter(passivateAfter))
		sys, _ := NewActorSystem(cfg)

		// start the actor system
		err := sys.Start(ctx)
		assert.NoError(t, err)

		kind := "TestActor"
		id := "test-1"
		actor := NewTestActor(id)
		actorRef := sys.Spawn(ctx, kind, actor)
		assert.NotNil(t, actorRef)
		assert.EqualValues(t, 1, len(sys.Actors()))

		// let us sleep for some time to make the actor idle
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			time.Sleep(recvDelay)
			wg.Done()
		}()
		// block until timer is up
		wg.Wait()

		assert.Empty(t, sys.Actors())
		assert.NoError(t, sys.Stop(ctx))
	})
}