#import "WindowController.h"

@interface WindowController () <NSTouchBarDelegate>
@property (copy) void (^handler)(char *);
@property NSDictionary* goData;
@property NSDictionary* identifierMapping;
@property NSDictionary* imageMapping;
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
  self.handler = handler;
  NSError* err = [self setData:data];
  if (err != nil) {
    if (error != nil) {
      *error = err;
    }
    return nil;
  }
  [self initMapping];
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

- (NSError*)setData:(const char *)rawData {
  NSData* data = [NSData dataWithBytes:(const void *)rawData length:sizeof(unsigned char) * strlen(rawData)];

  NSError* error = nil;
  self.goData = [NSJSONSerialization JSONObjectWithData:data options:0 error:&error];
  if (error != nil) {
    return error;
  }
  return nil;
}

- (NSTouchBar*)makeTouchBar {
  NSTouchBar* bar = [[NSTouchBar alloc] init];

  NSMutableArray* defaults = [[NSMutableArray alloc] init];
  [defaults setArray:[self.goData objectForKey:@"Default"]];
  if ([[self.goData objectForKey:@"OtherItemsProxy"] intValue] == 1) {
    [defaults addObject:NSTouchBarItemIdentifierOtherItemsProxy];
  }

  for (int i = 0; i < [defaults count]; ++i) {
    NSTouchBarItemIdentifier newIdentifier = [self transformIdentifier:[defaults objectAtIndex:i]];
    if (newIdentifier != nil) {
      [defaults replaceObjectAtIndex:i withObject:newIdentifier];
    }
  }

  [bar setDelegate:self];
  [bar setDefaultItemIdentifiers:defaults];
  [bar setPrincipalItemIdentifier:[self transformIdentifier:[self.goData objectForKey:@"Principal"]]];
  [bar setEscapeKeyReplacementItemIdentifier:[self transformIdentifier:[self.goData objectForKey:@"Escape"]]];
  return bar;
}

- (nullable NSTouchBarItem *)touchBar:(NSTouchBar *)touchBar makeItemForIdentifier:(NSTouchBarItemIdentifier)identifier {
  // TODO: finish
  NSDictionary* data = [[self.goData objectForKey:@"Items"] objectForKey:identifier];
  if ([identifier hasPrefix:prefixButton]) {
    NSButtonTouchBarItem* item = [[[NSButtonTouchBarItem alloc] initWithIdentifier:identifier] autorelease];
    [self updateWidgetButton:item withData:data];
    return item;
  } else if ([identifier hasPrefix:prefixLabel]) {
    NSCustomTouchBarItem* item = [[[NSCustomTouchBarItem alloc] initWithIdentifier:identifier] autorelease];
    [self updateWidgetLabel:item withData:data];
    return item;
  } else {
    NSLog(@"warning: unsupported identifier %@ with %@", identifier, data);
  }
  return nil;
}

- (void)updateWidgetCore:(NSTouchBarItem*)item withData:(NSDictionary*)data {
  item.visibilityPriority = [[data objectForKey:@"Priority"] floatValue] == 0;
}

- (void)updateWidgetButton:(NSButtonTouchBarItem*)item withData:(NSDictionary*)data {
  [self updateWidgetCore:item withData:data];

  item.title = [data valueForKeyPath:@"Title"];
  item.image = [self transformImage:[data valueForKeyPath:@"Image"]];
  item.target = self;
  item.action = @selector(buttonAction:);
  item.enabled = [[data valueForKeyPath:@"Disabled"] intValue] == 0;
  item.bezelColor = [self transformColor:[data valueForKeyPath:@"BezelColor"]];
}

- (void)buttonAction:(id)sender {
  NSString* identifier = ((NSButtonTouchBarItem*) sender).identifier;
  const char * event = [[NSString stringWithFormat:@"{\"Kind\":\"button\",\"Target\":\"%@\"}", identifier] UTF8String];
  self.handler((char*) event);
}

- (void)updateWidgetLabel:(NSCustomTouchBarItem*)item withData:(NSDictionary*)data {
  [self updateWidgetCore:item withData:data];

  NSString* text = [data valueForKeyPath:@"Content.Text"];
  NSString* image = [data valueForKeyPath:@"Content.Image"];
  if (text != nil) {
    NSTextField* view = [NSTextField labelWithString:text];
    view.textColor = [self transformColor:[data valueForKeyPath:@"Content.Color"]];;
    [item setView:view];
  } else if (image != nil) {
    NSImageView* view = [NSImageView imageViewWithImage: [self transformImage:image]];
    [item setView:view];
  } else {
    NSLog(@"warning: label with invalid data %@", data);
  }
}

- (NSTouchBarItemIdentifier) transformIdentifier:(NSString*) name {
  NSTouchBarItemIdentifier standard = [self.identifierMapping objectForKey:name];
  if (standard != nil) {
    return standard;
  }
  return name;
}

- (NSImage*) transformImage:(NSString*)name {
  if (name == nil || name == (id)[NSNull null]) {
    return nil;
  }
  NSImageName standard = [self.imageMapping objectForKey:name];
  if (standard != nil) {
    return [NSImage imageNamed:standard];
  }
  NSImage* sf = [NSImage imageWithSystemSymbolName:name accessibilityDescription:name];
  if (sf == nil) {
    NSLog(@"warning: could not find SF Symbols for %@", name);
  }
  return sf;
}

- (NSColor*) transformColor:(NSDictionary*)details {
  if (details == nil || details == (id)[NSNull null]) {
    return nil;
  }
  return [NSColor
    colorWithSRGBRed:[[details objectForKey:@"Red"] doubleValue]
    green:[[details objectForKey:@"Green"] doubleValue]
    blue:[[details objectForKey:@"Blue"] doubleValue]
    alpha:[[details objectForKey:@"Alpha"] doubleValue]
  ];
}

- (void) initMapping {
  self.identifierMapping = @{
    standardSpaceSmall: NSTouchBarItemIdentifierFixedSpaceSmall,
    standardSpaceLarge: NSTouchBarItemIdentifierFixedSpaceLarge,
    standardSpaceFlexible: NSTouchBarItemIdentifierFlexibleSpace,
    standardCandidateList: NSTouchBarItemIdentifierCandidateList,
    standardCharacterPicker: NSTouchBarItemIdentifierCharacterPicker,
    standardTextFormat: NSTouchBarItemIdentifierTextFormat,
    standardTextAlignment: NSTouchBarItemIdentifierTextAlignment,
    standardTextColorPicker: NSTouchBarItemIdentifierTextColorPicker,
    standardTextList: NSTouchBarItemIdentifierTextList,
    standardTextStyle: NSTouchBarItemIdentifierTextStyle,
  };
  self.imageMapping = @{
    @"TBAddDetailTemplate": NSImageNameTouchBarAddDetailTemplate,
    @"TBAddTemplate": NSImageNameTouchBarAddTemplate,
    @"TBAlarmTemplate": NSImageNameTouchBarAlarmTemplate,
    @"TBAudioInputMuteTemplate": NSImageNameTouchBarAudioInputMuteTemplate,
    @"TBAudioInputTemplate": NSImageNameTouchBarAudioInputTemplate,
    @"TBAudioOutputMuteTemplate": NSImageNameTouchBarAudioOutputMuteTemplate,
    @"TBAudioOutputVolumeHighTemplate": NSImageNameTouchBarAudioOutputVolumeHighTemplate,
    @"TBAudioOutputVolumeLowTemplate": NSImageNameTouchBarAudioOutputVolumeLowTemplate,
    @"TBAudioOutputVolumeMediumTemplate": NSImageNameTouchBarAudioOutputVolumeMediumTemplate,
    @"TBAudioOutputVolumeOffTemplate": NSImageNameTouchBarAudioOutputVolumeOffTemplate,
    @"TBBookmarksTemplate": NSImageNameTouchBarBookmarksTemplate,
    @"TBColorPickerFill": NSImageNameTouchBarColorPickerFill,
    @"TBColorPickerFont": NSImageNameTouchBarColorPickerFont,
    @"TBColorPickerStroke": NSImageNameTouchBarColorPickerStroke,
    @"TBCommunicationAudioTemplate": NSImageNameTouchBarCommunicationAudioTemplate,
    @"TBCommunicationVideoTemplate": NSImageNameTouchBarCommunicationVideoTemplate,
    @"TBComposeTemplate": NSImageNameTouchBarComposeTemplate,
    @"TBDeleteTemplate": NSImageNameTouchBarDeleteTemplate,
    @"TBDownloadTemplate": NSImageNameTouchBarDownloadTemplate,
    @"TBEnterFullScreenTemplate": NSImageNameTouchBarEnterFullScreenTemplate,
    @"TBExitFullScreenTemplate": NSImageNameTouchBarExitFullScreenTemplate,
    @"TBFastForwardTemplate": NSImageNameTouchBarFastForwardTemplate,
    @"TBFolderCopyToTemplate": NSImageNameTouchBarFolderCopyToTemplate,
    @"TBFolderMoveToTemplate": NSImageNameTouchBarFolderMoveToTemplate,
    @"TBFolderTemplate": NSImageNameTouchBarFolderTemplate,
    @"TBGetInfoTemplate": NSImageNameTouchBarGetInfoTemplate,
    @"TBGoBackTemplate": NSImageNameTouchBarGoBackTemplate,
    @"TBGoDownTemplate": NSImageNameTouchBarGoDownTemplate,
    @"TBGoForwardTemplate": NSImageNameTouchBarGoForwardTemplate,
    @"TBGoUpTemplate": NSImageNameTouchBarGoUpTemplate,
    @"TBHistoryTemplate": NSImageNameTouchBarHistoryTemplate,
    @"TBIconViewTemplate": NSImageNameTouchBarIconViewTemplate,
    @"TBListViewTemplate": NSImageNameTouchBarListViewTemplate,
    @"TBMailTemplate": NSImageNameTouchBarMailTemplate,
    @"TBNewFolderTemplate": NSImageNameTouchBarNewFolderTemplate,
    @"TBNewMessageTemplate": NSImageNameTouchBarNewMessageTemplate,
    @"TBOpenInBrowserTemplate": NSImageNameTouchBarOpenInBrowserTemplate,
    @"TBPauseTemplate": NSImageNameTouchBarPauseTemplate,
    @"TBPlayheadTemplate": NSImageNameTouchBarPlayheadTemplate,
    @"TBPlayPauseTemplate": NSImageNameTouchBarPlayPauseTemplate,
    @"TBPlayTemplate": NSImageNameTouchBarPlayTemplate,
    @"TBQuickLookTemplate": NSImageNameTouchBarQuickLookTemplate,
    @"TBRecordStartTemplate": NSImageNameTouchBarRecordStartTemplate,
    @"TBRecordStopTemplate": NSImageNameTouchBarRecordStopTemplate,
    @"TBRefreshTemplate": NSImageNameTouchBarRefreshTemplate,
    @"TBRewindTemplate": NSImageNameTouchBarRewindTemplate,
    @"TBRotateLeftTemplate": NSImageNameTouchBarRotateLeftTemplate,
    @"TBRotateRightTemplate": NSImageNameTouchBarRotateRightTemplate,
    @"TBSearchTemplate": NSImageNameTouchBarSearchTemplate,
    @"TBShareTemplate": NSImageNameTouchBarShareTemplate,
    @"TBSidebarTemplate": NSImageNameTouchBarSidebarTemplate,
    @"TBSkipAhead15SecondsTemplate": NSImageNameTouchBarSkipAhead15SecondsTemplate,
    @"TBSkipAhead30SecondsTemplate": NSImageNameTouchBarSkipAhead30SecondsTemplate,
    @"TBSkipAheadTemplate": NSImageNameTouchBarSkipAheadTemplate,
    @"TBSkipBack15SecondsTemplate": NSImageNameTouchBarSkipBack15SecondsTemplate,
    @"TBSkipBack30SecondsTemplate": NSImageNameTouchBarSkipBack30SecondsTemplate,
    @"TBSkipBackTemplate": NSImageNameTouchBarSkipBackTemplate,
    @"TBSkipToEndTemplate": NSImageNameTouchBarSkipToEndTemplate,
    @"TBSkipToStartTemplate": NSImageNameTouchBarSkipToStartTemplate,
    @"TBSlideshowTemplate": NSImageNameTouchBarSlideshowTemplate,
    @"TBTagIconTemplate": NSImageNameTouchBarTagIconTemplate,
    @"TBTextBoldTemplate": NSImageNameTouchBarTextBoldTemplate,
    @"TBTextBoxTemplate": NSImageNameTouchBarTextBoxTemplate,
    @"TBTextCenterAlignTemplate": NSImageNameTouchBarTextCenterAlignTemplate,
    @"TBTextItalicTemplate": NSImageNameTouchBarTextItalicTemplate,
    @"TBTextJustifiedAlignTemplate": NSImageNameTouchBarTextJustifiedAlignTemplate,
    @"TBTextLeftAlignTemplate": NSImageNameTouchBarTextLeftAlignTemplate,
    @"TBTextListTemplate": NSImageNameTouchBarTextListTemplate,
    @"TBTextRightAlignTemplate": NSImageNameTouchBarTextRightAlignTemplate,
    @"TBTextStrikethroughTemplate": NSImageNameTouchBarTextStrikethroughTemplate,
    @"TBTextUnderlineTemplate": NSImageNameTouchBarTextUnderlineTemplate,
    @"TBUserAddTemplate": NSImageNameTouchBarUserAddTemplate,
    @"TBUserGroupTemplate": NSImageNameTouchBarUserGroupTemplate,
    @"TBUserTemplate": NSImageNameTouchBarUserTemplate,
  };
}
@end
