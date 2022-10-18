#import "WindowController.h"

@interface WindowController () <NSTouchBarDelegate>
@property NSString* data2;
@end

@implementation WindowController
static NSTouchBarItemIdentifier prefixButton = @"net.lbrunner.touchbar.button.";
static NSTouchBarItemIdentifier prefixCandidates = @"net.lbrunner.touchbar.candidates.";
static NSTouchBarItemIdentifier prefixColorpicker = @"net.lbrunner.touchbar.colorpicker.";
static NSTouchBarItemIdentifier prefixCustom = @"net.lbrunner.touchbar.custom.";
static NSTouchBarItemIdentifier prefixGroup = @"net.lbrunner.touchbar.group.";
static NSTouchBarItemIdentifier prefixLabel = @"net.lbrunner.touchbar.label.";
static NSTouchBarItemIdentifier prefixPicker = @"net.lbrunner.touchbar.picker.";
static NSTouchBarItemIdentifier prefixPopover = @"net.lbrunner.touchbar.popover.";
static NSTouchBarItemIdentifier prefixScrubber = @"net.lbrunner.touchbar.scrubber.";
static NSTouchBarItemIdentifier prefixSegmented = @"net.lbrunner.touchbar.segmented.";
static NSTouchBarItemIdentifier prefixSharer = @"net.lbrunner.touchbar.sharer.";
static NSTouchBarItemIdentifier prefixSlider = @"net.lbrunner.touchbar.slider.";
static NSTouchBarItemIdentifier prefixSpacer = @"net.lbrunner.touchbar.spacer.";
static NSTouchBarItemIdentifier prefixStepper = @"net.lbrunner.touchbar.stepper.";

- (id)initWithData:(const char *)data {
  if ((self = [WindowController alloc]) == nil) {
    return nil;
  }
  [self setData:data];
  return self;
}

- (void)setData:(const char *)data {
  self.data2 = [NSString stringWithCString:data encoding:NSUTF8StringEncoding];;
}

- (NSTouchBar*)makeTouchBar {
  NSMutableArray *items = [[NSMutableArray alloc]init];
  [items addObject:prefixLabel];
  [items addObject:NSTouchBarItemIdentifierOtherItemsProxy];
  NSTouchBar* bar = [[NSTouchBar alloc] init];
  [bar setDelegate:self];
  [bar setDefaultItemIdentifiers:items];
  return bar;
}

- (nullable NSTouchBarItem *)touchBar:(NSTouchBar *)touchBar makeItemForIdentifier:(NSTouchBarItemIdentifier)identifier {
  if ([identifier hasPrefix:prefixLabel]) {
    NSCustomTouchBarItem* item = [[[NSCustomTouchBarItem alloc] initWithIdentifier:identifier] autorelease];

    NSTextField* view = [NSTextField labelWithString:self.data2];
    [item setView:view];
    return item;
  }
  return nil;
}
@end
