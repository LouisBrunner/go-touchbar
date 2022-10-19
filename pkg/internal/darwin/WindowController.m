#import "WindowController.h"

@interface WindowController () <NSTouchBarDelegate>
@property void (^handler)(char *);
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
  _handler = handler;
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
  if ([identifier hasPrefix:prefixLabel]) {
    NSDictionary* data = [[self.goData objectForKey:@"Items"] objectForKey:identifier];
    NSCustomTouchBarItem* item = [[[NSCustomTouchBarItem alloc] initWithIdentifier:identifier] autorelease];
    [self updateWidgetLabel:item withData:data];
    return item;
  }
  return nil;
}

- (void)updateWidgetCore:(NSCustomTouchBarItem*)item withData:(NSDictionary*)data {
  item.visibilityPriority = [[data objectForKey:@"Priority"] floatValue] == 0;
}

- (void)updateWidgetLabel:(NSCustomTouchBarItem*)item withData:(NSDictionary*)data {
  [self updateWidgetCore:item withData:data];

  NSString* text = [data valueForKeyPath:@"Content.Text"];
  NSString* image = [data valueForKeyPath:@"Content.Image"];
  if (text != nil) {
    NSTextField* view = [NSTextField labelWithString:text];
    [item setView:view];
  } else if (image != nil) {
    NSImageView* view = [[NSImageView alloc] init];
    NSImageName imageName = [self transformImage:image];
    if (imageName != nil) {
      view.image = [NSImage imageNamed:imageName];
    } else {
      view.image = [NSImage imageWithSystemSymbolName:image accessibilityDescription:image];
    }
    [item setView:view];
  }
}

- (NSTouchBarItemIdentifier) transformIdentifier:(NSString*) name {
  NSTouchBarItemIdentifier standard = [self.identifierMapping objectForKey:name];
  if (standard != nil) {
    return standard;
  }
  return name;
}

- (nullable NSImageName) transformImage:(NSString*)name {
  NSImageName standard = [self.imageMapping objectForKey:name];
  if (standard == nil) {
    return nil;
  }
  return standard;
}

- (void) initMapping {
  _identifierMapping = @{
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
  _imageMapping = @{
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
