// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

    certificates_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"

    "github.com/pkg/errors"
    "github.com/solo-io/skv2/pkg/events"
    "sigs.k8s.io/controller-runtime/pkg/manager"
    "sigs.k8s.io/controller-runtime/pkg/predicate"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

// Handle events for the IssuedCertificate Resource
// DEPRECATED: Prefer reconciler pattern.
type IssuedCertificateEventHandler interface {
    CreateIssuedCertificate(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error
    UpdateIssuedCertificate(old, new *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error
    DeleteIssuedCertificate(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error
    GenericIssuedCertificate(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error
}

type IssuedCertificateEventHandlerFuncs struct {
    OnCreate  func(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error
    OnUpdate  func(old, new *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error
    OnDelete  func(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error
    OnGeneric func(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error
}

func (f *IssuedCertificateEventHandlerFuncs) CreateIssuedCertificate(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *IssuedCertificateEventHandlerFuncs) DeleteIssuedCertificate(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *IssuedCertificateEventHandlerFuncs) UpdateIssuedCertificate(objOld, objNew *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *IssuedCertificateEventHandlerFuncs) GenericIssuedCertificate(obj *certificates_mesh_gloo_solo_io_v1.IssuedCertificate) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type IssuedCertificateEventWatcher interface {
    AddEventHandler(ctx context.Context, h IssuedCertificateEventHandler, predicates ...predicate.Predicate) error
}

type issuedCertificateEventWatcher struct {
    watcher events.EventWatcher
}

func NewIssuedCertificateEventWatcher(name string, mgr manager.Manager) IssuedCertificateEventWatcher {
    return &issuedCertificateEventWatcher{
        watcher: events.NewWatcher(name, mgr, &certificates_mesh_gloo_solo_io_v1.IssuedCertificate{}),
    }
}

func (c *issuedCertificateEventWatcher) AddEventHandler(ctx context.Context, h IssuedCertificateEventHandler, predicates ...predicate.Predicate) error {
	handler := genericIssuedCertificateHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericIssuedCertificateHandler implements a generic events.EventHandler
type genericIssuedCertificateHandler struct {
    handler IssuedCertificateEventHandler
}

func (h genericIssuedCertificateHandler) Create(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.IssuedCertificate)
    if !ok {
        return errors.Errorf("internal error: IssuedCertificate handler received event for %T", object)
    }
    return h.handler.CreateIssuedCertificate(obj)
}

func (h genericIssuedCertificateHandler) Delete(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.IssuedCertificate)
    if !ok {
        return errors.Errorf("internal error: IssuedCertificate handler received event for %T", object)
    }
    return h.handler.DeleteIssuedCertificate(obj)
}

func (h genericIssuedCertificateHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*certificates_mesh_gloo_solo_io_v1.IssuedCertificate)
    if !ok {
        return errors.Errorf("internal error: IssuedCertificate handler received event for %T", old)
    }
    objNew, ok := new.(*certificates_mesh_gloo_solo_io_v1.IssuedCertificate)
    if !ok {
        return errors.Errorf("internal error: IssuedCertificate handler received event for %T", new)
    }
    return h.handler.UpdateIssuedCertificate(objOld, objNew)
}

func (h genericIssuedCertificateHandler) Generic(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.IssuedCertificate)
    if !ok {
        return errors.Errorf("internal error: IssuedCertificate handler received event for %T", object)
    }
    return h.handler.GenericIssuedCertificate(obj)
}

// Handle events for the CertificateRequest Resource
// DEPRECATED: Prefer reconciler pattern.
type CertificateRequestEventHandler interface {
    CreateCertificateRequest(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error
    UpdateCertificateRequest(old, new *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error
    DeleteCertificateRequest(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error
    GenericCertificateRequest(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error
}

type CertificateRequestEventHandlerFuncs struct {
    OnCreate  func(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error
    OnUpdate  func(old, new *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error
    OnDelete  func(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error
    OnGeneric func(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error
}

func (f *CertificateRequestEventHandlerFuncs) CreateCertificateRequest(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *CertificateRequestEventHandlerFuncs) DeleteCertificateRequest(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *CertificateRequestEventHandlerFuncs) UpdateCertificateRequest(objOld, objNew *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *CertificateRequestEventHandlerFuncs) GenericCertificateRequest(obj *certificates_mesh_gloo_solo_io_v1.CertificateRequest) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type CertificateRequestEventWatcher interface {
    AddEventHandler(ctx context.Context, h CertificateRequestEventHandler, predicates ...predicate.Predicate) error
}

type certificateRequestEventWatcher struct {
    watcher events.EventWatcher
}

func NewCertificateRequestEventWatcher(name string, mgr manager.Manager) CertificateRequestEventWatcher {
    return &certificateRequestEventWatcher{
        watcher: events.NewWatcher(name, mgr, &certificates_mesh_gloo_solo_io_v1.CertificateRequest{}),
    }
}

func (c *certificateRequestEventWatcher) AddEventHandler(ctx context.Context, h CertificateRequestEventHandler, predicates ...predicate.Predicate) error {
	handler := genericCertificateRequestHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericCertificateRequestHandler implements a generic events.EventHandler
type genericCertificateRequestHandler struct {
    handler CertificateRequestEventHandler
}

func (h genericCertificateRequestHandler) Create(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.CertificateRequest)
    if !ok {
        return errors.Errorf("internal error: CertificateRequest handler received event for %T", object)
    }
    return h.handler.CreateCertificateRequest(obj)
}

func (h genericCertificateRequestHandler) Delete(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.CertificateRequest)
    if !ok {
        return errors.Errorf("internal error: CertificateRequest handler received event for %T", object)
    }
    return h.handler.DeleteCertificateRequest(obj)
}

func (h genericCertificateRequestHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*certificates_mesh_gloo_solo_io_v1.CertificateRequest)
    if !ok {
        return errors.Errorf("internal error: CertificateRequest handler received event for %T", old)
    }
    objNew, ok := new.(*certificates_mesh_gloo_solo_io_v1.CertificateRequest)
    if !ok {
        return errors.Errorf("internal error: CertificateRequest handler received event for %T", new)
    }
    return h.handler.UpdateCertificateRequest(objOld, objNew)
}

func (h genericCertificateRequestHandler) Generic(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.CertificateRequest)
    if !ok {
        return errors.Errorf("internal error: CertificateRequest handler received event for %T", object)
    }
    return h.handler.GenericCertificateRequest(obj)
}

// Handle events for the PodBounceDirective Resource
// DEPRECATED: Prefer reconciler pattern.
type PodBounceDirectiveEventHandler interface {
    CreatePodBounceDirective(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error
    UpdatePodBounceDirective(old, new *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error
    DeletePodBounceDirective(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error
    GenericPodBounceDirective(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error
}

type PodBounceDirectiveEventHandlerFuncs struct {
    OnCreate  func(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error
    OnUpdate  func(old, new *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error
    OnDelete  func(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error
    OnGeneric func(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error
}

func (f *PodBounceDirectiveEventHandlerFuncs) CreatePodBounceDirective(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error {
    if f.OnCreate == nil {
        return nil
    }
    return f.OnCreate(obj)
}

func (f *PodBounceDirectiveEventHandlerFuncs) DeletePodBounceDirective(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error {
    if f.OnDelete == nil {
        return nil
    }
    return f.OnDelete(obj)
}

func (f *PodBounceDirectiveEventHandlerFuncs) UpdatePodBounceDirective(objOld, objNew *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error {
    if f.OnUpdate == nil {
        return nil
    }
    return f.OnUpdate(objOld, objNew)
}

func (f *PodBounceDirectiveEventHandlerFuncs) GenericPodBounceDirective(obj *certificates_mesh_gloo_solo_io_v1.PodBounceDirective) error {
    if f.OnGeneric == nil {
        return nil
    }
    return f.OnGeneric(obj)
}

type PodBounceDirectiveEventWatcher interface {
    AddEventHandler(ctx context.Context, h PodBounceDirectiveEventHandler, predicates ...predicate.Predicate) error
}

type podBounceDirectiveEventWatcher struct {
    watcher events.EventWatcher
}

func NewPodBounceDirectiveEventWatcher(name string, mgr manager.Manager) PodBounceDirectiveEventWatcher {
    return &podBounceDirectiveEventWatcher{
        watcher: events.NewWatcher(name, mgr, &certificates_mesh_gloo_solo_io_v1.PodBounceDirective{}),
    }
}

func (c *podBounceDirectiveEventWatcher) AddEventHandler(ctx context.Context, h PodBounceDirectiveEventHandler, predicates ...predicate.Predicate) error {
	handler := genericPodBounceDirectiveHandler{handler: h}
    if err := c.watcher.Watch(ctx, handler, predicates...); err != nil{
        return err
    }
    return nil
}

// genericPodBounceDirectiveHandler implements a generic events.EventHandler
type genericPodBounceDirectiveHandler struct {
    handler PodBounceDirectiveEventHandler
}

func (h genericPodBounceDirectiveHandler) Create(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.PodBounceDirective)
    if !ok {
        return errors.Errorf("internal error: PodBounceDirective handler received event for %T", object)
    }
    return h.handler.CreatePodBounceDirective(obj)
}

func (h genericPodBounceDirectiveHandler) Delete(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.PodBounceDirective)
    if !ok {
        return errors.Errorf("internal error: PodBounceDirective handler received event for %T", object)
    }
    return h.handler.DeletePodBounceDirective(obj)
}

func (h genericPodBounceDirectiveHandler) Update(old, new client.Object) error {
    objOld, ok := old.(*certificates_mesh_gloo_solo_io_v1.PodBounceDirective)
    if !ok {
        return errors.Errorf("internal error: PodBounceDirective handler received event for %T", old)
    }
    objNew, ok := new.(*certificates_mesh_gloo_solo_io_v1.PodBounceDirective)
    if !ok {
        return errors.Errorf("internal error: PodBounceDirective handler received event for %T", new)
    }
    return h.handler.UpdatePodBounceDirective(objOld, objNew)
}

func (h genericPodBounceDirectiveHandler) Generic(object client.Object) error {
    obj, ok := object.(*certificates_mesh_gloo_solo_io_v1.PodBounceDirective)
    if !ok {
        return errors.Errorf("internal error: PodBounceDirective handler received event for %T", object)
    }
    return h.handler.GenericPodBounceDirective(obj)
}
