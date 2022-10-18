#import "WindowController.h"

@interface WindowController () <NSTouchBarDelegate>
@end

@implementation WindowController
static NSTouchBarItemIdentifier label = @"net.lbrunner.touchbar.label.";

- (NSTouchBar*)makeTouchBar {
  NSMutableArray *items = [[NSMutableArray alloc]init];
  [items addObject:label];
  [items addObject:NSTouchBarItemIdentifierOtherItemsProxy];
  NSTouchBar* bar = [[NSTouchBar alloc] init];
  [bar setDelegate:self];
  [bar setDefaultItemIdentifiers:items];
  return bar;
}

- (nullable NSTouchBarItem *)touchBar:(NSTouchBar *)touchBar makeItemForIdentifier:(NSTouchBarItemIdentifier)identifier {
  if ([identifier isEqualToString:label]) {
    NSCustomTouchBarItem* item = [[NSCustomTouchBarItem alloc] initWithIdentifier:label];
    [item autorelease];

    NSTextField* view = [NSTextField labelWithString:@"ABC"];
    [item setView:view];
    return item;
  }
  return nil;
}
@end
