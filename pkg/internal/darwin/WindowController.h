#ifndef GO_TOUCH_BAR_WINDOW_CONTROLLER_H
#define GO_TOUCH_BAR_WINDOW_CONTROLLER_H

#import <Cocoa/Cocoa.h>

@interface WindowController : NSWindowController {}
- (id)initWithData:(const char *)data;
- (void)setData:(const char *)data;
@end

#endif
