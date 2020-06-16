// Package target provides the Chrome DevTools Protocol
// commands, types, and events for the Target domain.
//
// Supports additional targets discovery and allows to attach to them.
//
// Generated by the cdproto-gen command.
package target

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// ActivateTargetParams activates (focuses) the target.
type ActivateTargetParams struct {
	TargetID ID `json:"targetId"`
}

// ActivateTarget activates (focuses) the target.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-activateTarget
//
// parameters:
//   targetID
func ActivateTarget(targetID ID) *ActivateTargetParams {
	return &ActivateTargetParams{
		TargetID: targetID,
	}
}

// Do executes Target.activateTarget against the provided context.
func (p *ActivateTargetParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandActivateTarget, p, nil)
}

// AttachToTargetParams attaches to the target with given id.
type AttachToTargetParams struct {
	TargetID ID   `json:"targetId"`
	Flatten  bool `json:"flatten,omitempty"` // Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
}

// AttachToTarget attaches to the target with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-attachToTarget
//
// parameters:
//   targetID
func AttachToTarget(targetID ID) *AttachToTargetParams {
	return &AttachToTargetParams{
		TargetID: targetID,
	}
}

// WithFlatten enables "flat" access to the session via specifying sessionId
// attribute in the commands. We plan to make this the default, deprecate
// non-flattened mode, and eventually retire it. See crbug.com/991325.
func (p AttachToTargetParams) WithFlatten(flatten bool) *AttachToTargetParams {
	p.Flatten = flatten
	return &p
}

// AttachToTargetReturns return values.
type AttachToTargetReturns struct {
	SessionID SessionID `json:"sessionId,omitempty"` // Id assigned to the session.
}

// Do executes Target.attachToTarget against the provided context.
//
// returns:
//   sessionID - Id assigned to the session.
func (p *AttachToTargetParams) Do(ctx context.Context) (sessionID SessionID, err error) {
	// execute
	var res AttachToTargetReturns
	err = cdp.Execute(ctx, CommandAttachToTarget, p, &res)
	if err != nil {
		return "", err
	}

	return res.SessionID, nil
}

// AttachToBrowserTargetParams attaches to the browser target, only uses flat
// sessionId mode.
type AttachToBrowserTargetParams struct{}

// AttachToBrowserTarget attaches to the browser target, only uses flat
// sessionId mode.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-attachToBrowserTarget
func AttachToBrowserTarget() *AttachToBrowserTargetParams {
	return &AttachToBrowserTargetParams{}
}

// AttachToBrowserTargetReturns return values.
type AttachToBrowserTargetReturns struct {
	SessionID SessionID `json:"sessionId,omitempty"` // Id assigned to the session.
}

// Do executes Target.attachToBrowserTarget against the provided context.
//
// returns:
//   sessionID - Id assigned to the session.
func (p *AttachToBrowserTargetParams) Do(ctx context.Context) (sessionID SessionID, err error) {
	// execute
	var res AttachToBrowserTargetReturns
	err = cdp.Execute(ctx, CommandAttachToBrowserTarget, nil, &res)
	if err != nil {
		return "", err
	}

	return res.SessionID, nil
}

// CloseTargetParams closes the target. If the target is a page that gets
// closed too.
type CloseTargetParams struct {
	TargetID ID `json:"targetId"`
}

// CloseTarget closes the target. If the target is a page that gets closed
// too.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-closeTarget
//
// parameters:
//   targetID
func CloseTarget(targetID ID) *CloseTargetParams {
	return &CloseTargetParams{
		TargetID: targetID,
	}
}

// CloseTargetReturns return values.
type CloseTargetReturns struct {
	Success bool `json:"success,omitempty"`
}

// Do executes Target.closeTarget against the provided context.
//
// returns:
//   success
func (p *CloseTargetParams) Do(ctx context.Context) (success bool, err error) {
	// execute
	var res CloseTargetReturns
	err = cdp.Execute(ctx, CommandCloseTarget, p, &res)
	if err != nil {
		return false, err
	}

	return res.Success, nil
}

// ExposeDevToolsProtocolParams inject object to the target's main frame that
// provides a communication channel with browser target. Injected object will be
// available as window[bindingName]. The object has the follwing API: -
// binding.send(json) - a method to send messages over the remote debugging
// protocol - binding.onmessage = json => handleMessage(json) - a callback that
// will be called for the protocol notifications and command responses.
type ExposeDevToolsProtocolParams struct {
	TargetID    ID     `json:"targetId"`
	BindingName string `json:"bindingName,omitempty"` // Binding name, 'cdp' if not specified.
}

// ExposeDevToolsProtocol inject object to the target's main frame that
// provides a communication channel with browser target. Injected object will be
// available as window[bindingName]. The object has the follwing API: -
// binding.send(json) - a method to send messages over the remote debugging
// protocol - binding.onmessage = json => handleMessage(json) - a callback that
// will be called for the protocol notifications and command responses.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-exposeDevToolsProtocol
//
// parameters:
//   targetID
func ExposeDevToolsProtocol(targetID ID) *ExposeDevToolsProtocolParams {
	return &ExposeDevToolsProtocolParams{
		TargetID: targetID,
	}
}

// WithBindingName binding name, 'cdp' if not specified.
func (p ExposeDevToolsProtocolParams) WithBindingName(bindingName string) *ExposeDevToolsProtocolParams {
	p.BindingName = bindingName
	return &p
}

// Do executes Target.exposeDevToolsProtocol against the provided context.
func (p *ExposeDevToolsProtocolParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandExposeDevToolsProtocol, p, nil)
}

// CreateBrowserContextParams creates a new empty BrowserContext. Similar to
// an incognito profile but you can have more than one.
type CreateBrowserContextParams struct {
	DisposeOnDetach bool `json:"disposeOnDetach,omitempty"` // If specified, disposes this context when debugging session disconnects.
}

// CreateBrowserContext creates a new empty BrowserContext. Similar to an
// incognito profile but you can have more than one.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-createBrowserContext
//
// parameters:
func CreateBrowserContext() *CreateBrowserContextParams {
	return &CreateBrowserContextParams{}
}

// WithDisposeOnDetach if specified, disposes this context when debugging
// session disconnects.
func (p CreateBrowserContextParams) WithDisposeOnDetach(disposeOnDetach bool) *CreateBrowserContextParams {
	p.DisposeOnDetach = disposeOnDetach
	return &p
}

// CreateBrowserContextReturns return values.
type CreateBrowserContextReturns struct {
	BrowserContextID cdp.BrowserContextID `json:"browserContextId,omitempty"` // The id of the context created.
}

// Do executes Target.createBrowserContext against the provided context.
//
// returns:
//   browserContextID - The id of the context created.
func (p *CreateBrowserContextParams) Do(ctx context.Context) (browserContextID cdp.BrowserContextID, err error) {
	// execute
	var res CreateBrowserContextReturns
	err = cdp.Execute(ctx, CommandCreateBrowserContext, p, &res)
	if err != nil {
		return "", err
	}

	return res.BrowserContextID, nil
}

// GetBrowserContextsParams returns all browser contexts created with
// Target.createBrowserContext method.
type GetBrowserContextsParams struct{}

// GetBrowserContexts returns all browser contexts created with
// Target.createBrowserContext method.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-getBrowserContexts
func GetBrowserContexts() *GetBrowserContextsParams {
	return &GetBrowserContextsParams{}
}

// GetBrowserContextsReturns return values.
type GetBrowserContextsReturns struct {
	BrowserContextIds []cdp.BrowserContextID `json:"browserContextIds,omitempty"` // An array of browser context ids.
}

// Do executes Target.getBrowserContexts against the provided context.
//
// returns:
//   browserContextIds - An array of browser context ids.
func (p *GetBrowserContextsParams) Do(ctx context.Context) (browserContextIds []cdp.BrowserContextID, err error) {
	// execute
	var res GetBrowserContextsReturns
	err = cdp.Execute(ctx, CommandGetBrowserContexts, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.BrowserContextIds, nil
}

// CreateTargetParams creates a new page.
type CreateTargetParams struct {
	URL                     string               `json:"url"`                               // The initial URL the page will be navigated to.
	Width                   int64                `json:"width,omitempty"`                   // Frame width in DIP (headless chrome only).
	Height                  int64                `json:"height,omitempty"`                  // Frame height in DIP (headless chrome only).
	BrowserContextID        cdp.BrowserContextID `json:"browserContextId,omitempty"`        // The browser context to create the page in.
	EnableBeginFrameControl bool                 `json:"enableBeginFrameControl,omitempty"` // Whether BeginFrames for this target will be controlled via DevTools (headless chrome only, not supported on MacOS yet, false by default).
	NewWindow               bool                 `json:"newWindow,omitempty"`               // Whether to create a new Window or Tab (chrome-only, false by default).
	Background              bool                 `json:"background,omitempty"`              // Whether to create the target in background or foreground (chrome-only, false by default).
}

// CreateTarget creates a new page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-createTarget
//
// parameters:
//   url - The initial URL the page will be navigated to.
func CreateTarget(url string) *CreateTargetParams {
	return &CreateTargetParams{
		URL: url,
	}
}

// WithWidth frame width in DIP (headless chrome only).
func (p CreateTargetParams) WithWidth(width int64) *CreateTargetParams {
	p.Width = width
	return &p
}

// WithHeight frame height in DIP (headless chrome only).
func (p CreateTargetParams) WithHeight(height int64) *CreateTargetParams {
	p.Height = height
	return &p
}

// WithBrowserContextID the browser context to create the page in.
func (p CreateTargetParams) WithBrowserContextID(browserContextID cdp.BrowserContextID) *CreateTargetParams {
	p.BrowserContextID = browserContextID
	return &p
}

// WithEnableBeginFrameControl whether BeginFrames for this target will be
// controlled via DevTools (headless chrome only, not supported on MacOS yet,
// false by default).
func (p CreateTargetParams) WithEnableBeginFrameControl(enableBeginFrameControl bool) *CreateTargetParams {
	p.EnableBeginFrameControl = enableBeginFrameControl
	return &p
}

// WithNewWindow whether to create a new Window or Tab (chrome-only, false by
// default).
func (p CreateTargetParams) WithNewWindow(newWindow bool) *CreateTargetParams {
	p.NewWindow = newWindow
	return &p
}

// WithBackground whether to create the target in background or foreground
// (chrome-only, false by default).
func (p CreateTargetParams) WithBackground(background bool) *CreateTargetParams {
	p.Background = background
	return &p
}

// CreateTargetReturns return values.
type CreateTargetReturns struct {
	TargetID ID `json:"targetId,omitempty"` // The id of the page opened.
}

// Do executes Target.createTarget against the provided context.
//
// returns:
//   targetID - The id of the page opened.
func (p *CreateTargetParams) Do(ctx context.Context) (targetID ID, err error) {
	// execute
	var res CreateTargetReturns
	err = cdp.Execute(ctx, CommandCreateTarget, p, &res)
	if err != nil {
		return "", err
	}

	return res.TargetID, nil
}

// DetachFromTargetParams detaches session with given id.
type DetachFromTargetParams struct {
	SessionID SessionID `json:"sessionId,omitempty"` // Session to detach.
}

// DetachFromTarget detaches session with given id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-detachFromTarget
//
// parameters:
func DetachFromTarget() *DetachFromTargetParams {
	return &DetachFromTargetParams{}
}

// WithSessionID session to detach.
func (p DetachFromTargetParams) WithSessionID(sessionID SessionID) *DetachFromTargetParams {
	p.SessionID = sessionID
	return &p
}

// Do executes Target.detachFromTarget against the provided context.
func (p *DetachFromTargetParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDetachFromTarget, p, nil)
}

// DisposeBrowserContextParams deletes a BrowserContext. All the belonging
// pages will be closed without calling their beforeunload hooks.
type DisposeBrowserContextParams struct {
	BrowserContextID cdp.BrowserContextID `json:"browserContextId"`
}

// DisposeBrowserContext deletes a BrowserContext. All the belonging pages
// will be closed without calling their beforeunload hooks.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-disposeBrowserContext
//
// parameters:
//   browserContextID
func DisposeBrowserContext(browserContextID cdp.BrowserContextID) *DisposeBrowserContextParams {
	return &DisposeBrowserContextParams{
		BrowserContextID: browserContextID,
	}
}

// Do executes Target.disposeBrowserContext against the provided context.
func (p *DisposeBrowserContextParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisposeBrowserContext, p, nil)
}

// GetTargetInfoParams returns information about a target.
type GetTargetInfoParams struct {
	TargetID ID `json:"targetId,omitempty"`
}

// GetTargetInfo returns information about a target.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-getTargetInfo
//
// parameters:
func GetTargetInfo() *GetTargetInfoParams {
	return &GetTargetInfoParams{}
}

// WithTargetID [no description].
func (p GetTargetInfoParams) WithTargetID(targetID ID) *GetTargetInfoParams {
	p.TargetID = targetID
	return &p
}

// GetTargetInfoReturns return values.
type GetTargetInfoReturns struct {
	TargetInfo *Info `json:"targetInfo,omitempty"`
}

// Do executes Target.getTargetInfo against the provided context.
//
// returns:
//   targetInfo
func (p *GetTargetInfoParams) Do(ctx context.Context) (targetInfo *Info, err error) {
	// execute
	var res GetTargetInfoReturns
	err = cdp.Execute(ctx, CommandGetTargetInfo, p, &res)
	if err != nil {
		return nil, err
	}

	return res.TargetInfo, nil
}

// GetTargetsParams retrieves a list of available targets.
type GetTargetsParams struct{}

// GetTargets retrieves a list of available targets.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-getTargets
func GetTargets() *GetTargetsParams {
	return &GetTargetsParams{}
}

// GetTargetsReturns return values.
type GetTargetsReturns struct {
	TargetInfos []*Info `json:"targetInfos,omitempty"` // The list of targets.
}

// Do executes Target.getTargets against the provided context.
//
// returns:
//   targetInfos - The list of targets.
func (p *GetTargetsParams) Do(ctx context.Context) (targetInfos []*Info, err error) {
	// execute
	var res GetTargetsReturns
	err = cdp.Execute(ctx, CommandGetTargets, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.TargetInfos, nil
}

// SetAutoAttachParams controls whether to automatically attach to new
// targets which are considered to be related to this one. When turned on,
// attaches to all existing related targets as well. When turned off,
// automatically detaches from all currently attached targets.
type SetAutoAttachParams struct {
	AutoAttach             bool `json:"autoAttach"`             // Whether to auto-attach to related targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"` // Whether to pause new targets when attaching to them. Use Runtime.runIfWaitingForDebugger to run paused targets.
	Flatten                bool `json:"flatten,omitempty"`      // Enables "flat" access to the session via specifying sessionId attribute in the commands. We plan to make this the default, deprecate non-flattened mode, and eventually retire it. See crbug.com/991325.
}

// SetAutoAttach controls whether to automatically attach to new targets
// which are considered to be related to this one. When turned on, attaches to
// all existing related targets as well. When turned off, automatically detaches
// from all currently attached targets.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-setAutoAttach
//
// parameters:
//   autoAttach - Whether to auto-attach to related targets.
//   waitForDebuggerOnStart - Whether to pause new targets when attaching to them. Use Runtime.runIfWaitingForDebugger to run paused targets.
func SetAutoAttach(autoAttach bool, waitForDebuggerOnStart bool) *SetAutoAttachParams {
	return &SetAutoAttachParams{
		AutoAttach:             autoAttach,
		WaitForDebuggerOnStart: waitForDebuggerOnStart,
	}
}

// WithFlatten enables "flat" access to the session via specifying sessionId
// attribute in the commands. We plan to make this the default, deprecate
// non-flattened mode, and eventually retire it. See crbug.com/991325.
func (p SetAutoAttachParams) WithFlatten(flatten bool) *SetAutoAttachParams {
	p.Flatten = flatten
	return &p
}

// Do executes Target.setAutoAttach against the provided context.
func (p *SetAutoAttachParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetAutoAttach, p, nil)
}

// SetDiscoverTargetsParams controls whether to discover available targets
// and notify via targetCreated/targetInfoChanged/targetDestroyed events.
type SetDiscoverTargetsParams struct {
	Discover bool `json:"discover"` // Whether to discover available targets.
}

// SetDiscoverTargets controls whether to discover available targets and
// notify via targetCreated/targetInfoChanged/targetDestroyed events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-setDiscoverTargets
//
// parameters:
//   discover - Whether to discover available targets.
func SetDiscoverTargets(discover bool) *SetDiscoverTargetsParams {
	return &SetDiscoverTargetsParams{
		Discover: discover,
	}
}

// Do executes Target.setDiscoverTargets against the provided context.
func (p *SetDiscoverTargetsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetDiscoverTargets, p, nil)
}

// SetRemoteLocationsParams enables target discovery for the specified
// locations, when setDiscoverTargets was set to true.
type SetRemoteLocationsParams struct {
	Locations []*RemoteLocation `json:"locations"` // List of remote locations.
}

// SetRemoteLocations enables target discovery for the specified locations,
// when setDiscoverTargets was set to true.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Target#method-setRemoteLocations
//
// parameters:
//   locations - List of remote locations.
func SetRemoteLocations(locations []*RemoteLocation) *SetRemoteLocationsParams {
	return &SetRemoteLocationsParams{
		Locations: locations,
	}
}

// Do executes Target.setRemoteLocations against the provided context.
func (p *SetRemoteLocationsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetRemoteLocations, p, nil)
}

// Command names.
const (
	CommandActivateTarget         = "Target.activateTarget"
	CommandAttachToTarget         = "Target.attachToTarget"
	CommandAttachToBrowserTarget  = "Target.attachToBrowserTarget"
	CommandCloseTarget            = "Target.closeTarget"
	CommandExposeDevToolsProtocol = "Target.exposeDevToolsProtocol"
	CommandCreateBrowserContext   = "Target.createBrowserContext"
	CommandGetBrowserContexts     = "Target.getBrowserContexts"
	CommandCreateTarget           = "Target.createTarget"
	CommandDetachFromTarget       = "Target.detachFromTarget"
	CommandDisposeBrowserContext  = "Target.disposeBrowserContext"
	CommandGetTargetInfo          = "Target.getTargetInfo"
	CommandGetTargets             = "Target.getTargets"
	CommandSetAutoAttach          = "Target.setAutoAttach"
	CommandSetDiscoverTargets     = "Target.setDiscoverTargets"
	CommandSetRemoteLocations     = "Target.setRemoteLocations"
)
