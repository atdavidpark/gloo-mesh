// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

    networking_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"

    "github.com/pkg/errors"
    "github.com/solo-io/skv2/pkg/events"
    "sigs.k8s.io/controller-runtime/pkg/manager"
    "sigs.k8s.io/controller-runtime/pkg/predicate"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

// Handle events for the TrafficPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type TrafficPolicyEventHandler interface {
    CreateTrafficPolicy(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error
    UpdateTrafficPolicy(old, new *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error
    DeleteTrafficPolicy(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error
    GenericTrafficPolicy(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error
}

type TrafficPolicyEventHandlerFuncs struct {
    OnCreate  func(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error
    OnUpdate  func(old, new *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error
    OnDelete  func(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error
    OnGeneric func(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error
}

func (f *TrafficPolicyEventHandlerFuncs) CreateTrafficPolicy(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *TrafficPolicyEventHandlerFuncs) DeleteTrafficPolicy(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *TrafficPolicyEventHandlerFuncs) UpdateTrafficPolicy(objOld, objNew *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *TrafficPolicyEventHandlerFuncs) GenericTrafficPolicy(obj *networking_mesh_gloo_solo_io_v1.TrafficPolicy) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type TrafficPolicyEventWatcher interface {
    AddEventHandler(ctx context.Context, h TrafficPolicyEventHandler, predicates ...predicate.Predicate) error
}

type trafficPolicyEventWatcher struct {
    watcher events.EventWatcher
}

func NewTrafficPolicyEventWatcher(name string, mgr manager.Manager) TrafficPolicyEventWatcher {
    return &trafficPolicyEventWatcher{
        watcher: events.NewWatcher(name, mgr, &networking_mesh_gloo_solo_io_v1.TrafficPolicy{}),
    }
}

func (c *trafficPolicyEventWatcher) AddEventHandler(ctx context.Context, h TrafficPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericTrafficPolicyHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericTrafficPolicyHandler implements a generic events.EventHandler
type genericTrafficPolicyHandler struct {
    handler TrafficPolicyEventHandler
}

func (h genericTrafficPolicyHandler) Create(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy)
    if !ok {
        return errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
    }
    return h.handler.CreateTrafficPolicy(obj)
}

func (h genericTrafficPolicyHandler) Delete(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy)
    if !ok {
        return errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
    }
    return h.handler.DeleteTrafficPolicy(obj)
}

func (h genericTrafficPolicyHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy)
    if !ok {
        return errors.Errorf("internal error: TrafficPolicy handler received event for %T", old)
    }
    objNew, ok := new.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy)
    if !ok {
        return errors.Errorf("internal error: TrafficPolicy handler received event for %T", new)
    }
    return h.handler.UpdateTrafficPolicy(objOld, objNew)
}

func (h genericTrafficPolicyHandler) Generic(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.TrafficPolicy)
    if !ok {
        return errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
    }
    return h.handler.GenericTrafficPolicy(obj)
}

// Handle events for the AccessPolicy Resource
// DEPRECATED: Prefer reconciler pattern.
type AccessPolicyEventHandler interface {
    CreateAccessPolicy(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error
    UpdateAccessPolicy(old, new *networking_mesh_gloo_solo_io_v1.AccessPolicy) error
    DeleteAccessPolicy(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error
    GenericAccessPolicy(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error
}

type AccessPolicyEventHandlerFuncs struct {
    OnCreate  func(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error
    OnUpdate  func(old, new *networking_mesh_gloo_solo_io_v1.AccessPolicy) error
    OnDelete  func(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error
    OnGeneric func(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error
}

func (f *AccessPolicyEventHandlerFuncs) CreateAccessPolicy(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *AccessPolicyEventHandlerFuncs) DeleteAccessPolicy(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *AccessPolicyEventHandlerFuncs) UpdateAccessPolicy(objOld, objNew *networking_mesh_gloo_solo_io_v1.AccessPolicy) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *AccessPolicyEventHandlerFuncs) GenericAccessPolicy(obj *networking_mesh_gloo_solo_io_v1.AccessPolicy) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type AccessPolicyEventWatcher interface {
    AddEventHandler(ctx context.Context, h AccessPolicyEventHandler, predicates ...predicate.Predicate) error
}

type accessPolicyEventWatcher struct {
    watcher events.EventWatcher
}

func NewAccessPolicyEventWatcher(name string, mgr manager.Manager) AccessPolicyEventWatcher {
    return &accessPolicyEventWatcher{
        watcher: events.NewWatcher(name, mgr, &networking_mesh_gloo_solo_io_v1.AccessPolicy{}),
    }
}

func (c *accessPolicyEventWatcher) AddEventHandler(ctx context.Context, h AccessPolicyEventHandler, predicates ...predicate.Predicate) error {
	handler := genericAccessPolicyHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericAccessPolicyHandler implements a generic events.EventHandler
type genericAccessPolicyHandler struct {
    handler AccessPolicyEventHandler
}

func (h genericAccessPolicyHandler) Create(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.AccessPolicy)
    if !ok {
        return errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
    }
    return h.handler.CreateAccessPolicy(obj)
}

func (h genericAccessPolicyHandler) Delete(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.AccessPolicy)
    if !ok {
        return errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
    }
    return h.handler.DeleteAccessPolicy(obj)
}

func (h genericAccessPolicyHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*networking_mesh_gloo_solo_io_v1.AccessPolicy)
    if !ok {
        return errors.Errorf("internal error: AccessPolicy handler received event for %T", old)
    }
    objNew, ok := new.(*networking_mesh_gloo_solo_io_v1.AccessPolicy)
    if !ok {
        return errors.Errorf("internal error: AccessPolicy handler received event for %T", new)
    }
    return h.handler.UpdateAccessPolicy(objOld, objNew)
}

func (h genericAccessPolicyHandler) Generic(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.AccessPolicy)
    if !ok {
        return errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
    }
    return h.handler.GenericAccessPolicy(obj)
}

// Handle events for the VirtualMesh Resource
// DEPRECATED: Prefer reconciler pattern.
type VirtualMeshEventHandler interface {
    CreateVirtualMesh(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error
    UpdateVirtualMesh(old, new *networking_mesh_gloo_solo_io_v1.VirtualMesh) error
    DeleteVirtualMesh(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error
    GenericVirtualMesh(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error
}

type VirtualMeshEventHandlerFuncs struct {
    OnCreate  func(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error
    OnUpdate  func(old, new *networking_mesh_gloo_solo_io_v1.VirtualMesh) error
    OnDelete  func(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error
    OnGeneric func(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error
}

func (f *VirtualMeshEventHandlerFuncs) CreateVirtualMesh(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *VirtualMeshEventHandlerFuncs) DeleteVirtualMesh(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *VirtualMeshEventHandlerFuncs) UpdateVirtualMesh(objOld, objNew *networking_mesh_gloo_solo_io_v1.VirtualMesh) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *VirtualMeshEventHandlerFuncs) GenericVirtualMesh(obj *networking_mesh_gloo_solo_io_v1.VirtualMesh) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type VirtualMeshEventWatcher interface {
    AddEventHandler(ctx context.Context, h VirtualMeshEventHandler, predicates ...predicate.Predicate) error
}

type virtualMeshEventWatcher struct {
    watcher events.EventWatcher
}

func NewVirtualMeshEventWatcher(name string, mgr manager.Manager) VirtualMeshEventWatcher {
    return &virtualMeshEventWatcher{
        watcher: events.NewWatcher(name, mgr, &networking_mesh_gloo_solo_io_v1.VirtualMesh{}),
    }
}

func (c *virtualMeshEventWatcher) AddEventHandler(ctx context.Context, h VirtualMeshEventHandler, predicates ...predicate.Predicate) error {
	handler := genericVirtualMeshHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericVirtualMeshHandler implements a generic events.EventHandler
type genericVirtualMeshHandler struct {
    handler VirtualMeshEventHandler
}

func (h genericVirtualMeshHandler) Create(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.VirtualMesh)
    if !ok {
        return errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
    }
    return h.handler.CreateVirtualMesh(obj)
}

func (h genericVirtualMeshHandler) Delete(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.VirtualMesh)
    if !ok {
        return errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
    }
    return h.handler.DeleteVirtualMesh(obj)
}

func (h genericVirtualMeshHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*networking_mesh_gloo_solo_io_v1.VirtualMesh)
    if !ok {
        return errors.Errorf("internal error: VirtualMesh handler received event for %T", old)
    }
    objNew, ok := new.(*networking_mesh_gloo_solo_io_v1.VirtualMesh)
    if !ok {
        return errors.Errorf("internal error: VirtualMesh handler received event for %T", new)
    }
    return h.handler.UpdateVirtualMesh(objOld, objNew)
}

func (h genericVirtualMeshHandler) Generic(object client.Object) error {
    obj, ok := object.(*networking_mesh_gloo_solo_io_v1.VirtualMesh)
    if !ok {
        return errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
    }
    return h.handler.GenericVirtualMesh(obj)
}
