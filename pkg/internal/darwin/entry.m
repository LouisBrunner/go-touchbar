#import "entry.h"
#import "WindowController.h"
#include <pthread.h>

// FIXME: I am 99% sure this code leaks heavily
// but I don't understand Objective-C enough to know where and how to fix it

static NSError* makeError(NSString* reason) {
  return [[[NSError alloc] initWithDomain:@"net.lbrunner.touchbar.go" code:1 userInfo:@{@"Error reason": reason}] autorelease];
}

static NSError* setWindowController(NSWindow* window, WindowController* controller) {
  if (window.windowController != nil) {
    return makeError(@"specified window already has a window controller");
  }
  window.windowController = controller;
  controller.window = window;
  return nil;
}

static NSWindow* makeDebugWindow(NSApplication *app) {
  NSWindow* window = [[[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, 200, 200) styleMask:NSWindowStyleMaskTitled backing:NSBackingStoreBuffered defer:NO] autorelease];
  [window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
  [window setTitle:@"go-touchbar tester"];
  [window makeKeyAndOrderFront:nil];

  id quitMenuItem = [[[NSMenuItem alloc] initWithTitle:@"Quit" action:@selector(terminate:) keyEquivalent:@"q"] autorelease];
  id appMenu = [[NSMenu new] autorelease];
  [appMenu addItem:quitMenuItem];
  id appMenuItem = [[NSMenuItem new] autorelease];
  [appMenuItem setSubmenu:appMenu];
  id menubar = [[NSMenu new] autorelease];
  [menubar addItem:appMenuItem];

  [app setMainMenu:menubar];
  [app activateIgnoringOtherApps:YES];

  return window;
}

typedef struct Context {
  AttachMode mode;
  NSWindow* window;
  WindowController* controller;
} Context;

InitResult initTouchBar(AttachMode mode, const char * data, void* me) {
  InitResult result;
  result.result = nil;
  result.err = nil;

  // TODO: pass a block with handleEvent and me together
  WindowController* controller = [[[WindowController alloc] initWithData:data] autorelease];

  NSApplication* app = [NSApplication sharedApplication];

  NSWindow* window;
  if (mode == kMainWindow) {
    window = app.mainWindow;
  } else if (mode == kDebug) {
    window = makeDebugWindow(app);
  } else {
    result.err = [[NSString stringWithFormat:@"Unknown mode %lu", mode] UTF8String];
    return result;
  }

  NSError* err = setWindowController(window, controller);
  if (err != nil) {
    result.err = [[err localizedDescription] UTF8String];
    return result;
  }

  Context* context = malloc(sizeof(Context));
  if (context == nil) {
    result.err = "could not allocate internal context";
    return result;
  }

  context->mode = mode;
  context->controller = controller;
  context->window = window;
  result.result = context;
  return result;
}

ErrorResult runDebug(void* ctx) {
  Context* context = (Context*) ctx;

  NSApplication* app = [NSApplication sharedApplication];
  [app run];

  ErrorResult result;
  result.err = nil;
  return result;
}

ErrorResult updateTouchBar(void* ctx, const char * data) {
  Context* context = (Context*) ctx;

  [context->controller setData:data];

  // Force the touchBar to be redrawn
  context->window.touchBar = nil;

  ErrorResult result;
  result.err = nil;
  return result;
}

ErrorResult destroyTouchBar(void* ctx) {
  Context* context = (Context*) ctx;
  if (context->mode == kDebug) {
    [[NSApplication sharedApplication] terminate:nil];
  }
  free(context);

  ErrorResult result;
  result.err = nil;
  return result;
}
