package gdextension

/*------------------------------------------------------------------------------
//   This code was generated by template globalconstants.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "globalconstants.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//revive:disable

// #include <godot/gdextension_interface.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	. "github.com/godot-go/godot-go/pkg/gdextensionffi"
)

// native structures
type AudioFrame struct {
	Left  float32
	Right float32
}
type CaretInfo struct {
	Leading_caret      Rect2
	Trailing_caret     Rect2
	Leading_direction  TextServerDirection
	Trailing_direction TextServerDirection
}
type Glyph struct {
	Start     int64
	End       int64
	Count     uint8
	Repeat    uint8
	Flags     uint16
	X_off     float32
	Y_off     float32
	Advance   float32
	Font_rid  RID
	Font_size int64
	Index     Int32T
}
type ObjectID struct {
	Id uint64
}
type PhysicsServer2DExtensionMotionResult struct {
	Travel                    Vector2
	Remainder                 Vector2
	Collision_point           Vector2
	Collision_normal          Vector2
	Collider_velocity         Vector2
	Collision_depth           float32
	Collision_safe_fraction   float32
	Collision_unsafe_fraction float32
	Collision_local_shape     int64
	Collider_id               ObjectID
	Collider                  RID
	Collider_shape            int64
}
type PhysicsServer2DExtensionRayResult struct {
	Position    Vector2
	Normal      Vector2
	Rid         RID
	Collider_id ObjectID
	Collider    Object
	Shape       int64
}
type PhysicsServer2DExtensionShapeRestInfo struct {
	Point           Vector2
	Normal          Vector2
	Rid             RID
	Collider_id     ObjectID
	Shape           int64
	Linear_velocity Vector2
}
type PhysicsServer2DExtensionShapeResult struct {
	Rid         RID
	Collider_id ObjectID
	Collider    Object
	Shape       int64
}
type PhysicsServer3DExtensionMotionCollision struct {
}
type PhysicsServer3DExtensionMotionResult struct {
}
type PhysicsServer3DExtensionRayResult struct {
}
type PhysicsServer3DExtensionShapeRestInfo struct {
}
type PhysicsServer3DExtensionShapeResult struct {
}
type ScriptLanguageExtensionProfilingInfo struct {
	Signature  StringName
	Call_count uint64
	Total_time uint64
	Self_time  uint64
}
