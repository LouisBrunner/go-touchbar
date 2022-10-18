#ifndef GO_TOUCH_BAR_ENTRY_H
#define GO_TOUCH_BAR_ENTRY_H

#import <Cocoa/Cocoa.h>

typedef enum AttachMode : NSUInteger {
    kMainWindow,
    kDebug
} AttachMode;

typedef struct TouchBar {

} TouchBar;

typedef struct InitResult {
  void* result;
  const char * err;
} InitResult;

typedef struct ErrorResult2 {
  const char * err;
} ErrorResult;

InitResult initTouchBar(AttachMode mode, TouchBar data);
ErrorResult runDebug(void* context);
ErrorResult updateTouchBar(void* context, TouchBar data);
ErrorResult destroyTouchBar(void* context);

#endif
