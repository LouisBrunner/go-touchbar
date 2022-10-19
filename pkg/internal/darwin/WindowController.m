#import "WindowController.h"

@interface WindowController () <NSTouchBarDelegate>
@property void (^handler)(char *);
@property NSString* data2;
@end

@implementation WindowController
static NSTouchBarItemIdentifier standardSpaceSmall = @"net.lbrunner.touchbar.small_space";
static NSTouchBarItemIdentifier standardSpaceLarge = @"net.lbrunner.touchbar.large_space";
static NSTouchBarItemIdentifier standardSpaceFlexible = @"net.lbrunner.touchbar.flexible_space";
static NSTouchBarItemIdentifier standardCandidateList = @"net.lbrunner.touchbar.candidates";
static NSTouchBarItemIdentifier standardCharacterPicker = @"net.lbrunner.touchbar.char_picker";
static NSTouchBarItemIdentifier standardTextFormat = @"net.lbrunner.touchbar.text_format";
static NSTouchBarItemIdentifier standardTextAlignment = @"net.lbrunner.touchbar.text_align";
static NSTouchBarItemIdentifier standardTextColorPicker = @"net.lbrunner.touchbar.text_color";
static NSTouchBarItemIdentifier standardTextList = @"net.lbrunner.touchbar.text_list";
static NSTouchBarItemIdentifier standardTextStyle = @"net.lbrunner.touchbar.text_style";

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
static NSTouchBarItemIdentifier prefixStepper = @"net.lbrunner.touchbar.stepper.";

- (id)initWithData:(const char *)data andHandler:(void (^)(char *))handler error:(NSError**)error {
  if ((self = [WindowController alloc]) == nil) {
    return nil;
  }
  _handler = handler;
  NSError* err = [self setData:data];
  if (err != nil) {
    if (error != nil) {
      *error = err;
    }
    return nil;
  }
  return self;
}

- (NSError*)updateWithData:(const char *)data {
  NSError* err = [self setData:data];
  if (err != nil) {
    return err;
  }
  // TODO: update code
  return nil;
}

- (NSError*)setData:(const char *)data {
  // TODO: parse data
  self.data2 = [NSString stringWithCString:data encoding:NSUTF8StringEncoding];
  return nil;
}

- (NSTouchBar*)makeTouchBar {
  NSMutableArray *items = [[NSMutableArray alloc]init];
  // TODO: use data
  [items addObject:prefixLabel];
  [items addObject:NSTouchBarItemIdentifierOtherItemsProxy];
  NSTouchBar* bar = [[NSTouchBar alloc] init];
  [bar setDelegate:self];
  [bar setDefaultItemIdentifiers:items];
  return bar;
}

- (nullable NSTouchBarItem *)touchBar:(NSTouchBar *)touchBar makeItemForIdentifier:(NSTouchBarItemIdentifier)identifier {
  // TODO: use data
  if ([identifier hasPrefix:prefixLabel]) {
    NSCustomTouchBarItem* item = [[[NSCustomTouchBarItem alloc] initWithIdentifier:identifier] autorelease];

    NSTextField* view = [NSTextField labelWithString:self.data2];
    [item setView:view];
    return item;
  }
  return nil;
}
@end
