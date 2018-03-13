package mtp

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"strings"
)

// predicates crc
const (
	crc_boolFalse                                        = 0xbc799737
	crc_boolTrue                                         = 0x997275b5
	crc_error                                            = 0xc4b9f9bb
	crc_null                                             = 0x56730bcc
	crc_vector                                           = 0x1cb5c415
	crc_inputPeerEmpty                                   = 0x7f3b18ea
	crc_inputPeerSelf                                    = 0x7da07ec9
	crc_inputPeerChat                                    = 0x179be863
	crc_inputUserEmpty                                   = 0xb98886cf
	crc_inputUserSelf                                    = 0xf7c1b13f
	crc_inputPhoneContact                                = 0xf392b7f4
	crc_inputFile                                        = 0xf52ff27f
	crc_inputMediaEmpty                                  = 0x9664f57f
	crc_inputMediaUploadedPhoto                          = 0x2f37e231
	crc_inputMediaPhoto                                  = 0x81fa373a
	crc_inputMediaGeoPoint                               = 0xf9c44144
	crc_inputMediaContact                                = 0xa6e45987
	crc_inputChatPhotoEmpty                              = 0x1ca48f57
	crc_inputChatUploadedPhoto                           = 0x927c55b4
	crc_inputChatPhoto                                   = 0x8953ad37
	crc_inputGeoPointEmpty                               = 0xe4c123d6
	crc_inputGeoPoint                                    = 0xf3b7acc9
	crc_inputPhotoEmpty                                  = 0x1cd7bf0d
	crc_inputPhoto                                       = 0xfb95c6c4
	crc_inputFileLocation                                = 0x14637196
	crc_inputAppEvent                                    = 0x770656a8
	crc_peerUser                                         = 0x9db1bc6d
	crc_peerChat                                         = 0xbad0e5bb
	crc_storageFileUnknown                               = 0xaa963b05
	crc_storageFileJpeg                                  = 0x007efe0e
	crc_storageFileGif                                   = 0xcae1aadf
	crc_storageFilePng                                   = 0x0a4f63c0
	crc_storageFileMp3                                   = 0x528a0677
	crc_storageFileMov                                   = 0x4b09ebbc
	crc_storageFilePartial                               = 0x40bc6f52
	crc_storageFileMp4                                   = 0xb3cea0e4
	crc_storageFileWebp                                  = 0x1081464c
	crc_fileLocationUnavailable                          = 0x7c596b46
	crc_fileLocation                                     = 0x53d69076
	crc_userEmpty                                        = 0x200250ba
	crc_userProfilePhotoEmpty                            = 0x4f11bae1
	crc_userProfilePhoto                                 = 0xd559d8c8
	crc_userStatusEmpty                                  = 0x09d05049
	crc_userStatusOnline                                 = 0xedb93949
	crc_userStatusOffline                                = 0x008c703f
	crc_chatEmpty                                        = 0x9ba2d800
	crc_chat                                             = 0xd91cdd54
	crc_chatForbidden                                    = 0x07328bdb
	crc_chatFull                                         = 0x2e02a614
	crc_chatParticipant                                  = 0xc8d7493e
	crc_chatParticipantsForbidden                        = 0xfc900c2b
	crc_chatParticipants                                 = 0x3f460fed
	crc_chatPhotoEmpty                                   = 0x37c1011c
	crc_chatPhoto                                        = 0x6153276a
	crc_messageEmpty                                     = 0x83e5de54
	crc_message                                          = 0x90dddc11
	crc_messageService                                   = 0x9e19a1f6
	crc_messageMediaEmpty                                = 0x3ded6320
	crc_messageMediaPhoto                                = 0xb5223b0f
	crc_messageMediaGeo                                  = 0x56e0d474
	crc_messageMediaContact                              = 0x5e7d2f39
	crc_messageMediaUnsupported                          = 0x9f84f49e
	crc_messageActionEmpty                               = 0xb6aef7b0
	crc_messageActionChatCreate                          = 0xa6638b9a
	crc_messageActionChatEditTitle                       = 0xb5a1ce5a
	crc_messageActionChatEditPhoto                       = 0x7fcb13a8
	crc_messageActionChatDeletePhoto                     = 0x95e3fbef
	crc_messageActionChatAddUser                         = 0x488a7337
	crc_messageActionChatDeleteUser                      = 0xb2ae9b0c
	crc_dialog                                           = 0xe4def5db
	crc_photoEmpty                                       = 0x2331b22d
	crc_photo                                            = 0x9288dd29
	crc_photoSizeEmpty                                   = 0x0e17e23c
	crc_photoSize                                        = 0x77bfb61b
	crc_photoCachedSize                                  = 0xe9a734fa
	crc_geoPointEmpty                                    = 0x1117dd5f
	crc_geoPoint                                         = 0x2049d70c
	crc_authCheckedPhone                                 = 0x811ea28e
	crc_authSentCode                                     = 0x5e002502
	crc_authAuthorization                                = 0xcd050916
	crc_authExportedAuthorization                        = 0xdf969c2d
	crc_inputNotifyPeer                                  = 0xb8bc5b0c
	crc_inputNotifyUsers                                 = 0x193b4417
	crc_inputNotifyChats                                 = 0x4a95e84e
	crc_inputNotifyAll                                   = 0xa429b886
	crc_inputPeerNotifySettings                          = 0x38935eb2
	crc_peerNotifyEventsEmpty                            = 0xadd53cb3
	crc_peerNotifyEventsAll                              = 0x6d1ded88
	crc_peerNotifySettingsEmpty                          = 0x70a68512
	crc_peerNotifySettings                               = 0x9acda4c0
	crc_wallPaper                                        = 0xccb03657
	crc_userFull                                         = 0x0f220f3f
	crc_contact                                          = 0xf911c994
	crc_importedContact                                  = 0xd0028438
	crc_contactBlocked                                   = 0x561bc879
	crc_contactStatus                                    = 0xd3680c61
	crc_contactsLink                                     = 0x3ace484c
	crc_contactsContacts                                 = 0xeae87e42
	crc_contactsContactsNotModified                      = 0xb74ba9d2
	crc_contactsImportedContacts                         = 0x77d01c3b
	crc_contactsBlocked                                  = 0x1c138d15
	crc_contactsBlockedSlice                             = 0x900802a1
	crc_contactsFound                                    = 0x1aa1f784
	crc_messagesDialogs                                  = 0x15ba6c40
	crc_messagesDialogsSlice                             = 0x71e094f3
	crc_messagesMessages                                 = 0x8c718e87
	crc_messagesMessagesSlice                            = 0x0b446ae3
	crc_messagesChats                                    = 0x64ff9fd5
	crc_messagesChatFull                                 = 0xe5d7d19c
	crc_messagesAffectedHistory                          = 0xb45c69d1
	crc_inputMessagesFilterEmpty                         = 0x57e2f66c
	crc_inputMessagesFilterPhotos                        = 0x9609a51c
	crc_inputMessagesFilterVideo                         = 0x9fc00e65
	crc_inputMessagesFilterPhotoVideo                    = 0x56e9f0e4
	crc_updateNewMessage                                 = 0x1f2b0afd
	crc_updateMessageID                                  = 0x4e90bfd6
	crc_updateDeleteMessages                             = 0xa20db0e5
	crc_updateUserTyping                                 = 0x5c486927
	crc_updateChatUserTyping                             = 0x9a65ea1f
	crc_updateChatParticipants                           = 0x07761198
	crc_updateUserStatus                                 = 0x1bfbd823
	crc_updateUserName                                   = 0xa7332b73
	crc_updateUserPhoto                                  = 0x95313b0c
	crc_updateContactRegistered                          = 0x2575bbb9
	crc_updateContactLink                                = 0x9d2e67c5
	crc_updatesState                                     = 0xa56c2a3e
	crc_updatesDifferenceEmpty                           = 0x5d75a138
	crc_updatesDifference                                = 0x00f49ca0
	crc_updatesDifferenceSlice                           = 0xa8fb1981
	crc_updatesTooLong                                   = 0xe317af7e
	crc_updateShortMessage                               = 0x914fbf11
	crc_updateShortChatMessage                           = 0x16812688
	crc_updateShort                                      = 0x78d4dec1
	crc_updatesCombined                                  = 0x725b04c3
	crc_updates                                          = 0x74ae4240
	crc_photosPhoto                                      = 0x20212ca8
	crc_uploadFile                                       = 0x096a18d5
	crc_dcOption                                         = 0x05d8c6cc
	crc_config                                           = 0x8df376a4
	crc_nearestDc                                        = 0x8e1a1775
	crc_helpAppUpdate                                    = 0x8987f311
	crc_helpNoAppUpdate                                  = 0xc45a6536
	crc_helpInviteText                                   = 0x18cb9f78
	crc_inputPeerNotifyEventsEmpty                       = 0xf03064d8
	crc_inputPeerNotifyEventsAll                         = 0xe86a2c74
	crc_photosPhotos                                     = 0x8dca6aa5
	crc_photosPhotosSlice                                = 0x15051f54
	crc_wallPaperSolid                                   = 0x63117f24
	crc_updateNewEncryptedMessage                        = 0x12bcbd9a
	crc_updateEncryptedChatTyping                        = 0x1710f156
	crc_updateEncryption                                 = 0xb4a2e88d
	crc_updateEncryptedMessagesRead                      = 0x38fe25b7
	crc_encryptedChatEmpty                               = 0xab7ec0a0
	crc_encryptedChatWaiting                             = 0x3bf703dc
	crc_encryptedChatRequested                           = 0xc878527e
	crc_encryptedChat                                    = 0xfa56ce36
	crc_encryptedChatDiscarded                           = 0x13d6dd27
	crc_inputEncryptedChat                               = 0xf141b5e1
	crc_encryptedFileEmpty                               = 0xc21f497e
	crc_encryptedFile                                    = 0x4a70994c
	crc_inputEncryptedFileEmpty                          = 0x1837c364
	crc_inputEncryptedFileUploaded                       = 0x64bd0306
	crc_inputEncryptedFile                               = 0x5a17b5e5
	crc_inputEncryptedFileLocation                       = 0xf5235d55
	crc_encryptedMessage                                 = 0xed18c118
	crc_encryptedMessageService                          = 0x23734b06
	crc_messagesDhConfigNotModified                      = 0xc0e24635
	crc_messagesDhConfig                                 = 0x2c221edd
	crc_messagesSentEncryptedMessage                     = 0x560f8935
	crc_messagesSentEncryptedFile                        = 0x9493ff32
	crc_inputFileBig                                     = 0xfa4f0bb5
	crc_inputEncryptedFileBigUploaded                    = 0x2dc173c8
	crc_storageFilePdf                                   = 0xae1e508d
	crc_inputMessagesFilterDocument                      = 0x9eddf188
	crc_inputMessagesFilterPhotoVideoDocuments           = 0xd95e73bb
	crc_updateChatParticipantAdd                         = 0xea4b0e5c
	crc_updateChatParticipantDelete                      = 0x6e5f8c22
	crc_updateDcOptions                                  = 0x8e5e9873
	crc_inputMediaUploadedDocument                       = 0xe39621fd
	crc_inputMediaDocument                               = 0x5acb668e
	crc_messageMediaDocument                             = 0x7c4414d3
	crc_inputDocumentEmpty                               = 0x72f0eaae
	crc_inputDocument                                    = 0x18798952
	crc_inputDocumentFileLocation                        = 0x430f0724
	crc_documentEmpty                                    = 0x36f8c871
	crc_document                                         = 0x87232bc7
	crc_helpSupport                                      = 0x17c6b5f6
	crc_notifyAll                                        = 0x74d07c60
	crc_notifyChats                                      = 0xc007cec3
	crc_notifyPeer                                       = 0x9fd40bd8
	crc_notifyUsers                                      = 0xb4c83b4c
	crc_updateUserBlocked                                = 0x80ece81a
	crc_updateNotifySettings                             = 0xbec268ef
	crc_sendMessageTypingAction                          = 0x16bf744e
	crc_sendMessageCancelAction                          = 0xfd5ec8f5
	crc_sendMessageRecordVideoAction                     = 0xa187d66f
	crc_sendMessageUploadVideoAction                     = 0xe9763aec
	crc_sendMessageRecordAudioAction                     = 0xd52f73f7
	crc_sendMessageUploadAudioAction                     = 0xf351d7ab
	crc_sendMessageUploadPhotoAction                     = 0xd1d34a26
	crc_sendMessageUploadDocumentAction                  = 0xaa0cd9e4
	crc_sendMessageGeoLocationAction                     = 0x176f8ba1
	crc_sendMessageChooseContactAction                   = 0x628cbc6f
	crc_updateServiceNotification                        = 0xebe46819
	crc_userStatusRecently                               = 0xe26f42f1
	crc_userStatusLastWeek                               = 0x07bf09fc
	crc_userStatusLastMonth                              = 0x77ebc742
	crc_updatePrivacy                                    = 0xee3b272a
	crc_inputPrivacyKeyStatusTimestamp                   = 0x4f96cb18
	crc_privacyKeyStatusTimestamp                        = 0xbc2eab30
	crc_inputPrivacyValueAllowContacts                   = 0x0d09e07b
	crc_inputPrivacyValueAllowAll                        = 0x184b35ce
	crc_inputPrivacyValueAllowUsers                      = 0x131cc67f
	crc_inputPrivacyValueDisallowContacts                = 0x0ba52007
	crc_inputPrivacyValueDisallowAll                     = 0xd66b66c9
	crc_inputPrivacyValueDisallowUsers                   = 0x90110467
	crc_privacyValueAllowContacts                        = 0xfffe1bac
	crc_privacyValueAllowAll                             = 0x65427b82
	crc_privacyValueAllowUsers                           = 0x4d5bbe0c
	crc_privacyValueDisallowContacts                     = 0xf888fa1a
	crc_privacyValueDisallowAll                          = 0x8b73e763
	crc_privacyValueDisallowUsers                        = 0x0c7f49b7
	crc_accountPrivacyRules                              = 0x554abb6f
	crc_accountDaysTTL                                   = 0xb8d0afdf
	crc_updateUserPhone                                  = 0x12b9417b
	crc_disabledFeature                                  = 0xae636f24
	crc_documentAttributeImageSize                       = 0x6c37c15c
	crc_documentAttributeAnimated                        = 0x11b58939
	crc_documentAttributeSticker                         = 0x6319d612
	crc_documentAttributeVideo                           = 0x0ef02ce6
	crc_documentAttributeAudio                           = 0x9852f9c6
	crc_documentAttributeFilename                        = 0x15590068
	crc_messagesStickersNotModified                      = 0xf1749a22
	crc_messagesStickers                                 = 0x8a8ecd32
	crc_stickerPack                                      = 0x12b299d4
	crc_messagesAllStickersNotModified                   = 0xe86602c3
	crc_messagesAllStickers                              = 0xedfd405f
	crc_accountNoPassword                                = 0x96dabc18
	crc_accountPassword                                  = 0x7c18141c
	crc_updateReadHistoryInbox                           = 0x9961fd5c
	crc_updateReadHistoryOutbox                          = 0x2f2f21bf
	crc_messagesAffectedMessages                         = 0x84d19185
	crc_contactLinkUnknown                               = 0x5f4f9247
	crc_contactLinkNone                                  = 0xfeedd3ad
	crc_contactLinkHasPhone                              = 0x268f3f59
	crc_contactLinkContact                               = 0xd502c2d0
	crc_updateWebPage                                    = 0x7f891213
	crc_webPageEmpty                                     = 0xeb1477e8
	crc_webPagePending                                   = 0xc586da1c
	crc_webPage                                          = 0x5f07b4bc
	crc_messageMediaWebPage                              = 0xa32dd600
	crc_authorization                                    = 0x7bf2e6f6
	crc_accountAuthorizations                            = 0x1250abde
	crc_accountPasswordSettings                          = 0xb7b72ab3
	crc_accountPasswordInputSettings                     = 0x86916deb
	crc_authPasswordRecovery                             = 0x137948a5
	crc_inputMediaVenue                                  = 0x2827a81a
	crc_messageMediaVenue                                = 0x7912b71f
	crc_receivedNotifyMessage                            = 0xa384b779
	crc_chatInviteEmpty                                  = 0x69df3769
	crc_chatInviteExported                               = 0xfc2e05bc
	crc_chatInviteAlready                                = 0x5a686d7c
	crc_chatInvite                                       = 0xdb74f558
	crc_messageActionChatJoinedByLink                    = 0xf89cf5e8
	crc_updateReadMessagesContents                       = 0x68c13933
	crc_inputStickerSetEmpty                             = 0xffb62b95
	crc_inputStickerSetID                                = 0x9de7a269
	crc_inputStickerSetShortName                         = 0x861cc8a0
	crc_stickerSet                                       = 0xcd303b41
	crc_messagesStickerSet                               = 0xb60a24a6
	crc_user                                             = 0x2e13f4c3
	crc_botCommand                                       = 0xc27ac8c7
	crc_botInfo                                          = 0x98e81d3a
	crc_keyboardButton                                   = 0xa2fa4880
	crc_keyboardButtonRow                                = 0x77608b83
	crc_replyKeyboardHide                                = 0xa03e5b85
	crc_replyKeyboardForceReply                          = 0xf4108aa0
	crc_replyKeyboardMarkup                              = 0x3502758c
	crc_inputMessagesFilterUrl                           = 0x7ef0dd87
	crc_inputPeerUser                                    = 0x7b8e7de6
	crc_inputUser                                        = 0xd8292816
	crc_messageEntityUnknown                             = 0xbb92ba95
	crc_messageEntityMention                             = 0xfa04579d
	crc_messageEntityHashtag                             = 0x6f635b0d
	crc_messageEntityBotCommand                          = 0x6cef8ac7
	crc_messageEntityUrl                                 = 0x6ed02538
	crc_messageEntityEmail                               = 0x64e475c2
	crc_messageEntityBold                                = 0xbd610bc9
	crc_messageEntityItalic                              = 0x826f8b60
	crc_messageEntityCode                                = 0x28a20571
	crc_messageEntityPre                                 = 0x73924be0
	crc_messageEntityTextUrl                             = 0x76a6d327
	crc_updateShortSentMessage                           = 0x11f1331c
	crc_inputPeerChannel                                 = 0x20adaef8
	crc_peerChannel                                      = 0xbddde532
	crc_channel                                          = 0x0cb44b1c
	crc_channelForbidden                                 = 0x289da732
	crc_channelFull                                      = 0x17f45fcf
	crc_messageActionChannelCreate                       = 0x95d2ac92
	crc_messagesChannelMessages                          = 0x99262e37
	crc_updateChannelTooLong                             = 0xeb0467fb
	crc_updateChannel                                    = 0xb6d45656
	crc_updateNewChannelMessage                          = 0x62ba04d9
	crc_updateReadChannelInbox                           = 0x4214f37f
	crc_updateDeleteChannelMessages                      = 0xc37521c9
	crc_updateChannelMessageViews                        = 0x98a12b4b
	crc_inputChannelEmpty                                = 0xee8c1e86
	crc_inputChannel                                     = 0xafeb712e
	crc_contactsResolvedPeer                             = 0x7f077ad9
	crc_messageRange                                     = 0x0ae30253
	crc_updatesChannelDifferenceEmpty                    = 0x3e11affb
	crc_updatesChannelDifferenceTooLong                  = 0x6a9d7b35
	crc_updatesChannelDifference                         = 0x2064674e
	crc_channelMessagesFilterEmpty                       = 0x94d42ee7
	crc_channelMessagesFilter                            = 0xcd77d957
	crc_channelParticipant                               = 0x15ebac1d
	crc_channelParticipantSelf                           = 0xa3289a6d
	crc_channelParticipantCreator                        = 0xe3e2e1f9
	crc_channelParticipantsRecent                        = 0xde3f3c79
	crc_channelParticipantsAdmins                        = 0xb4608969
	crc_channelParticipantsKicked                        = 0xa3b54985
	crc_channelsChannelParticipants                      = 0xf56ee2a8
	crc_channelsChannelParticipant                       = 0xd0d9b163
	crc_true                                             = 0x3fedd339
	crc_chatParticipantCreator                           = 0xda13538a
	crc_chatParticipantAdmin                             = 0xe2d6e436
	crc_updateChatAdmins                                 = 0x6e947941
	crc_updateChatParticipantAdmin                       = 0xb6901959
	crc_messageActionChatMigrateTo                       = 0x51bdb021
	crc_messageActionChannelMigrateFrom                  = 0xb055eaee
	crc_channelParticipantsBots                          = 0xb0d1865b
	crc_inputReportReasonSpam                            = 0x58dbcab8
	crc_inputReportReasonViolence                        = 0x1e22c78d
	crc_inputReportReasonPornography                     = 0x2e59d922
	crc_inputReportReasonOther                           = 0xe1746d0a
	crc_updateNewStickerSet                              = 0x688a30aa
	crc_updateStickerSetsOrder                           = 0x0bb2d201
	crc_updateStickerSets                                = 0x43ae3dec
	crc_helpTermsOfService                               = 0xf1ee3e90
	crc_foundGif                                         = 0x162ecc1f
	crc_inputMediaGifExternal                            = 0x4843b0fd
	crc_messagesFoundGifs                                = 0x450a1c0a
	crc_inputMessagesFilterGif                           = 0xffc86587
	crc_updateSavedGifs                                  = 0x9375341e
	crc_updateBotInlineQuery                             = 0x54826690
	crc_foundGifCached                                   = 0x9c750409
	crc_messagesSavedGifsNotModified                     = 0xe8025ca2
	crc_messagesSavedGifs                                = 0x2e0709a5
	crc_inputBotInlineMessageMediaAuto                   = 0x292fed13
	crc_inputBotInlineMessageText                        = 0x3dcd7a87
	crc_inputBotInlineResult                             = 0x2cbbe15a
	crc_botInlineMessageMediaAuto                        = 0x0a74b15b
	crc_botInlineMessageText                             = 0x8c7f65e2
	crc_botInlineResult                                  = 0x9bebaeb9
	crc_messagesBotResults                               = 0xccd3563d
	crc_inputMessagesFilterVoice                         = 0x50f5c392
	crc_inputMessagesFilterMusic                         = 0x3751b49e
	crc_updateBotInlineSend                              = 0x0e48f964
	crc_inputPrivacyKeyChatInvite                        = 0xbdfb0426
	crc_privacyKeyChatInvite                             = 0x500e6dfa
	crc_updateEditChannelMessage                         = 0x1b3f4df7
	crc_exportedMessageLink                              = 0x1f486803
	crc_messageFwdHeader                                 = 0xfadff4ac
	crc_messageActionPinMessage                          = 0x94bd38ed
	crc_peerSettings                                     = 0x818426cd
	crc_updateChannelPinnedMessage                       = 0x98592475
	crc_keyboardButtonUrl                                = 0x258aff05
	crc_keyboardButtonCallback                           = 0x683a5e46
	crc_keyboardButtonRequestPhone                       = 0xb16a6c29
	crc_keyboardButtonRequestGeoLocation                 = 0xfc796b3f
	crc_authCodeTypeSms                                  = 0x72a3158c
	crc_authCodeTypeCall                                 = 0x741cd3e3
	crc_authCodeTypeFlashCall                            = 0x226ccefb
	crc_authSentCodeTypeApp                              = 0x3dbb5986
	crc_authSentCodeTypeSms                              = 0xc000bba2
	crc_authSentCodeTypeCall                             = 0x5353e5a7
	crc_authSentCodeTypeFlashCall                        = 0xab03c6d9
	crc_keyboardButtonSwitchInline                       = 0x0568a748
	crc_replyInlineMarkup                                = 0x48a30254
	crc_messagesBotCallbackAnswer                        = 0x36585ea4
	crc_updateBotCallbackQuery                           = 0xe73547e1
	crc_messagesMessageEditData                          = 0x26b5dde6
	crc_updateEditMessage                                = 0xe40370a3
	crc_inputBotInlineMessageMediaGeo                    = 0xf4a59de1
	crc_inputBotInlineMessageMediaVenue                  = 0xaaafadc8
	crc_inputBotInlineMessageMediaContact                = 0x2daf01a7
	crc_botInlineMessageMediaGeo                         = 0x3a8fd8b8
	crc_botInlineMessageMediaVenue                       = 0x4366232e
	crc_botInlineMessageMediaContact                     = 0x35edb4d4
	crc_inputBotInlineResultPhoto                        = 0xa8d864a7
	crc_inputBotInlineResultDocument                     = 0xfff8fdc4
	crc_botInlineMediaResult                             = 0x17db940b
	crc_inputBotInlineMessageID                          = 0x890c3d89
	crc_updateInlineBotCallbackQuery                     = 0xf9d27a5a
	crc_inlineBotSwitchPM                                = 0x3c20629f
	crc_messageEntityMentionName                         = 0x352dca58
	crc_inputMessageEntityMentionName                    = 0x208e68c9
	crc_messagesPeerDialogs                              = 0x3371c354
	crc_topPeer                                          = 0xedcdc05b
	crc_topPeerCategoryBotsPM                            = 0xab661b5b
	crc_topPeerCategoryBotsInline                        = 0x148677e2
	crc_topPeerCategoryCorrespondents                    = 0x0637b7ed
	crc_topPeerCategoryGroups                            = 0xbd17a14a
	crc_topPeerCategoryChannels                          = 0x161d9628
	crc_topPeerCategoryPeers                             = 0xfb834291
	crc_contactsTopPeersNotModified                      = 0xde266ef5
	crc_contactsTopPeers                                 = 0x70b772a8
	crc_inputMessagesFilterChatPhotos                    = 0x3a20ecb8
	crc_updateReadChannelOutbox                          = 0x25d6c9c7
	crc_updateDraftMessage                               = 0xee2bb969
	crc_draftMessageEmpty                                = 0xba4baec5
	crc_draftMessage                                     = 0xfd8e711f
	crc_messageActionHistoryClear                        = 0x9fbab604
	crc_updateReadFeaturedStickers                       = 0x571d2742
	crc_updateRecentStickers                             = 0x9a422c20
	crc_messagesFeaturedStickersNotModified              = 0x04ede3cf
	crc_messagesFeaturedStickers                         = 0xf89d88e5
	crc_messagesRecentStickersNotModified                = 0x0b17f890
	crc_messagesRecentStickers                           = 0x5ce20970
	crc_messagesArchivedStickers                         = 0x4fcba9c8
	crc_messagesStickerSetInstallResultSuccess           = 0x38641628
	crc_messagesStickerSetInstallResultArchive           = 0x35e410a8
	crc_stickerSetCovered                                = 0x6410a5d2
	crc_inputMediaPhotoExternal                          = 0x0922aec1
	crc_inputMediaDocumentExternal                       = 0xb6f74335
	crc_updateConfig                                     = 0xa229dd06
	crc_updatePtsChanged                                 = 0x3354678f
	crc_messageActionGameScore                           = 0x92a72876
	crc_documentAttributeHasStickers                     = 0x9801d2f7
	crc_keyboardButtonGame                               = 0x50f41ccf
	crc_stickerSetMultiCovered                           = 0x3407e51b
	crc_maskCoords                                       = 0xaed6dbb2
	crc_inputStickeredMediaPhoto                         = 0x4a992157
	crc_inputStickeredMediaDocument                      = 0x0438865b
	crc_inputMediaGame                                   = 0xd33f43f3
	crc_messageMediaGame                                 = 0xfdb19008
	crc_inputBotInlineMessageGame                        = 0x4b425864
	crc_inputBotInlineResultGame                         = 0x4fa417f2
	crc_game                                             = 0xbdf9653b
	crc_inputGameID                                      = 0x032c3e77
	crc_inputGameShortName                               = 0xc331e80a
	crc_highScore                                        = 0x58fffcd0
	crc_messagesHighScores                               = 0x9a3bfd99
	crc_messagesChatsSlice                               = 0x9cd81144
	crc_updateChannelWebPage                             = 0x40771900
	crc_updatesDifferenceTooLong                         = 0x4afe8f6d
	crc_sendMessageGamePlayAction                        = 0xdd6a8f48
	crc_webPageNotModified                               = 0x85849473
	crc_textEmpty                                        = 0xdc3d824f
	crc_textPlain                                        = 0x744694e0
	crc_textBold                                         = 0x6724abc4
	crc_textItalic                                       = 0xd912a59c
	crc_textUnderline                                    = 0xc12622c4
	crc_textStrike                                       = 0x9bf8bb95
	crc_textFixed                                        = 0x6c3f19b9
	crc_textUrl                                          = 0x3c2884c1
	crc_textEmail                                        = 0xde5a0dd6
	crc_textConcat                                       = 0x7e6260d7
	crc_pageBlockTitle                                   = 0x70abc3fd
	crc_pageBlockSubtitle                                = 0x8ffa9a1f
	crc_pageBlockAuthorDate                              = 0xbaafe5e0
	crc_pageBlockHeader                                  = 0xbfd064ec
	crc_pageBlockSubheader                               = 0xf12bb6e1
	crc_pageBlockParagraph                               = 0x467a0766
	crc_pageBlockPreformatted                            = 0xc070d93e
	crc_pageBlockFooter                                  = 0x48870999
	crc_pageBlockDivider                                 = 0xdb20b188
	crc_pageBlockList                                    = 0x3a58c7f4
	crc_pageBlockBlockquote                              = 0x263d7c26
	crc_pageBlockPullquote                               = 0x4f4456d3
	crc_pageBlockPhoto                                   = 0xe9c69982
	crc_pageBlockVideo                                   = 0xd9d71866
	crc_pageBlockCover                                   = 0x39f23300
	crc_pageBlockEmbed                                   = 0xcde200d1
	crc_pageBlockEmbedPost                               = 0x292c7be9
	crc_pageBlockSlideshow                               = 0x130c8963
	crc_pagePart                                         = 0x8e3f9ebe
	crc_pageFull                                         = 0x556ec7aa
	crc_updatePhoneCall                                  = 0xab0f6b1e
	crc_updateDialogPinned                               = 0xd711a2cc
	crc_updatePinnedDialogs                              = 0xd8caf68d
	crc_inputPrivacyKeyPhoneCall                         = 0xfabadc5f
	crc_privacyKeyPhoneCall                              = 0x3d662b7b
	crc_pageBlockUnsupported                             = 0x13567e8a
	crc_pageBlockAnchor                                  = 0xce0d37b0
	crc_pageBlockCollage                                 = 0x08b31c4f
	crc_inputPhoneCall                                   = 0x1e36fded
	crc_phoneCallEmpty                                   = 0x5366c915
	crc_phoneCallWaiting                                 = 0x1b8f4ad1
	crc_phoneCallRequested                               = 0x83761ce4
	crc_phoneCall                                        = 0xffe6ab67
	crc_phoneCallDiscarded                               = 0x50ca4de1
	crc_phoneConnection                                  = 0x9d4c17c0
	crc_phoneCallProtocol                                = 0xa2bb35cb
	crc_phonePhoneCall                                   = 0xec82e140
	crc_phoneCallDiscardReasonMissed                     = 0x85e42301
	crc_phoneCallDiscardReasonDisconnect                 = 0xe095c1a0
	crc_phoneCallDiscardReasonHangup                     = 0x57adc690
	crc_phoneCallDiscardReasonBusy                       = 0xfaf7e8c9
	crc_inputMessagesFilterPhoneCalls                    = 0x80c99768
	crc_messageActionPhoneCall                           = 0x80e11a7f
	crc_invoice                                          = 0xc30aa358
	crc_inputMediaInvoice                                = 0x92153685
	crc_messageActionPaymentSentMe                       = 0x8f31b327
	crc_messageMediaInvoice                              = 0x84551347
	crc_keyboardButtonBuy                                = 0xafd93fbb
	crc_messageActionPaymentSent                         = 0x40699cd0
	crc_paymentsPaymentForm                              = 0x3f56aea3
	crc_postAddress                                      = 0x1e8caaeb
	crc_paymentRequestedInfo                             = 0x909c3f94
	crc_updateBotWebhookJSON                             = 0x8317c0c3
	crc_updateBotWebhookJSONQuery                        = 0x9b9240a6
	crc_updateBotShippingQuery                           = 0xe0cdc940
	crc_updateBotPrecheckoutQuery                        = 0x5d2f3aa9
	crc_dataJSON                                         = 0x7d748d04
	crc_labeledPrice                                     = 0xcb296bf8
	crc_paymentCharge                                    = 0xea02c27e
	crc_paymentSavedCredentialsCard                      = 0xcdc27a1f
	crc_webDocument                                      = 0xc61acbd8
	crc_inputWebDocument                                 = 0x9bed434d
	crc_inputWebFileLocation                             = 0xc239d686
	crc_uploadWebFile                                    = 0x21e753bc
	crc_paymentsValidatedRequestedInfo                   = 0xd1451883
	crc_paymentsPaymentResult                            = 0x4e5f810d
	crc_paymentsPaymentVerficationNeeded                 = 0x6b56b921
	crc_paymentsPaymentReceipt                           = 0x500911e1
	crc_paymentsSavedInfo                                = 0xfb8fe43c
	crc_inputPaymentCredentialsSaved                     = 0xc10eb2cf
	crc_inputPaymentCredentials                          = 0x3417d728
	crc_accountTmpPassword                               = 0xdb64fd34
	crc_shippingOption                                   = 0xb6213cdf
	crc_phoneCallAccepted                                = 0x6d003d3f
	crc_inputMessagesFilterRoundVoice                    = 0x7a7c17a4
	crc_inputMessagesFilterRoundVideo                    = 0xb549da53
	crc_uploadFileCdnRedirect                            = 0xea52fe5a
	crc_sendMessageRecordRoundAction                     = 0x88f27fbc
	crc_sendMessageUploadRoundAction                     = 0x243e1c66
	crc_uploadCdnFileReuploadNeeded                      = 0xeea8e46e
	crc_uploadCdnFile                                    = 0xa99fca4f
	crc_cdnPublicKey                                     = 0xc982eaba
	crc_cdnConfig                                        = 0x5725e40a
	crc_updateLangPackTooLong                            = 0x10c2404b
	crc_updateLangPack                                   = 0x56022f4d
	crc_pageBlockChannel                                 = 0xef1751b5
	crc_inputStickerSetItem                              = 0xffa0a496
	crc_langPackString                                   = 0xcad181f6
	crc_langPackStringPluralized                         = 0x6c47ac9f
	crc_langPackStringDeleted                            = 0x2979eeb2
	crc_langPackDifference                               = 0xf385c1f6
	crc_langPackLanguage                                 = 0x117698f1
	crc_channelParticipantAdmin                          = 0xa82fa898
	crc_channelParticipantBanned                         = 0x222c1886
	crc_channelParticipantsBanned                        = 0x1427a5e1
	crc_channelParticipantsSearch                        = 0x0656ac4b
	crc_topPeerCategoryPhoneCalls                        = 0x1e76a78c
	crc_pageBlockAudio                                   = 0x31b81a7f
	crc_channelAdminRights                               = 0x5d7ceba5
	crc_channelBannedRights                              = 0x58cf4249
	crc_channelAdminLogEventActionChangeTitle            = 0xe6dfb825
	crc_channelAdminLogEventActionChangeAbout            = 0x55188a2e
	crc_channelAdminLogEventActionChangeUsername         = 0x6a4afc38
	crc_channelAdminLogEventActionChangePhoto            = 0xb82f55c3
	crc_channelAdminLogEventActionToggleInvites          = 0x1b7907ae
	crc_channelAdminLogEventActionToggleSignatures       = 0x26ae0971
	crc_channelAdminLogEventActionUpdatePinned           = 0xe9e82c18
	crc_channelAdminLogEventActionEditMessage            = 0x709b2405
	crc_channelAdminLogEventActionDeleteMessage          = 0x42e047bb
	crc_channelAdminLogEventActionParticipantJoin        = 0x183040d3
	crc_channelAdminLogEventActionParticipantLeave       = 0xf89777f2
	crc_channelAdminLogEventActionParticipantInvite      = 0xe31c34d8
	crc_channelAdminLogEventActionParticipantToggleBan   = 0xe6d83d7e
	crc_channelAdminLogEventActionParticipantToggleAdmin = 0xd5676710
	crc_channelAdminLogEvent                             = 0x3b5a3e40
	crc_channelsAdminLogResults                          = 0xed8af74d
	crc_channelAdminLogEventsFilter                      = 0xea107ae4
	crc_messageActionScreenshotTaken                     = 0x4792929b
	crc_popularContact                                   = 0x5ce14175
	crc_cdnFileHash                                      = 0x77eec38f
	crc_inputMessagesFilterMyMentions                    = 0xc1f8e69a
	crc_inputMessagesFilterMyMentionsUnread              = 0x46caf4a8
	crc_updateContactsReset                              = 0x7084a7be
	crc_channelAdminLogEventActionChangeStickerSet       = 0xb1c3caa7
	crc_updateFavedStickers                              = 0xe511996d
	crc_messagesFavedStickers                            = 0xf37f2f16
	crc_messagesFavedStickersNotModified                 = 0x9e8fa6d3
	crc_updateChannelReadMessagesContents                = 0x89893b45
)

// methods crc
const (
	crc_invokeAfterMsg                     = 0xcb9f372d
	crc_invokeAfterMsgs                    = 0x3dc4b4f0
	crc_authCheckPhone                     = 0x6fe51dfb
	crc_authSendCode                       = 0x86aef0ec
	crc_authSignUp                         = 0x1b067634
	crc_authSignIn                         = 0xbcd51581
	crc_authLogOut                         = 0x5717da40
	crc_authResetAuthorizations            = 0x9fab0d1a
	crc_authSendInvites                    = 0x771c1d97
	crc_authExportAuthorization            = 0xe5bfffcd
	crc_authImportAuthorization            = 0xe3ef9613
	crc_accountRegisterDevice              = 0x637ea878
	crc_accountUnregisterDevice            = 0x65c55b40
	crc_accountUpdateNotifySettings        = 0x84be5b93
	crc_accountGetNotifySettings           = 0x12b3ad31
	crc_accountResetNotifySettings         = 0xdb7e1747
	crc_accountUpdateProfile               = 0x78515775
	crc_accountUpdateStatus                = 0x6628562c
	crc_accountGetWallPapers               = 0xc04cfac2
	crc_usersGetUsers                      = 0x0d91a548
	crc_usersGetFullUser                   = 0xca30a5b1
	crc_contactsGetStatuses                = 0xc4a353ee
	crc_contactsGetContacts                = 0xc023849f
	crc_contactsImportContacts             = 0x2c800be5
	crc_contactsSearch                     = 0x11f812d8
	crc_contactsDeleteContact              = 0x8e953744
	crc_contactsDeleteContacts             = 0x59ab389e
	crc_contactsBlock                      = 0x332b49fc
	crc_contactsUnblock                    = 0xe54100bd
	crc_contactsGetBlocked                 = 0xf57c350f
	crc_messagesGetMessages                = 0x4222fa74
	crc_messagesGetDialogs                 = 0x191ba9c5
	crc_messagesGetHistory                 = 0xafa92846
	crc_messagesSearch                     = 0x039e9ea0
	crc_messagesReadHistory                = 0x0e306d3a
	crc_messagesDeleteHistory              = 0x1c015b09
	crc_messagesDeleteMessages             = 0xe58e95d2
	crc_messagesReceivedMessages           = 0x05a954c0
	crc_messagesSetTyping                  = 0xa3825e50
	crc_messagesSendMessage                = 0xfa88427a
	crc_messagesSendMedia                  = 0xc8f16791
	crc_messagesForwardMessages            = 0x708e0195
	crc_messagesGetChats                   = 0x3c6aa187
	crc_messagesGetFullChat                = 0x3b831c66
	crc_messagesEditChatTitle              = 0xdc452855
	crc_messagesEditChatPhoto              = 0xca4c79d8
	crc_messagesAddChatUser                = 0xf9a0aa09
	crc_messagesDeleteChatUser             = 0xe0611f16
	crc_messagesCreateChat                 = 0x09cb126e
	crc_updatesGetState                    = 0xedd4882a
	crc_updatesGetDifference               = 0x25939651
	crc_photosUpdateProfilePhoto           = 0xf0bb5152
	crc_photosUploadProfilePhoto           = 0x4f32c098
	crc_uploadSaveFilePart                 = 0xb304a621
	crc_uploadGetFile                      = 0xe3a6cfb5
	crc_helpGetConfig                      = 0xc4f9186b
	crc_helpGetNearestDc                   = 0x1fb33026
	crc_helpGetAppUpdate                   = 0xae2de196
	crc_helpSaveAppLog                     = 0x6f02f748
	crc_helpGetInviteText                  = 0x4d392343
	crc_photosDeletePhotos                 = 0x87cf7f2f
	crc_photosGetUserPhotos                = 0x91cd32a8
	crc_messagesForwardMessage             = 0x33963bf9
	crc_messagesGetDhConfig                = 0x26cf8950
	crc_messagesRequestEncryption          = 0xf64daf43
	crc_messagesAcceptEncryption           = 0x3dbc0415
	crc_messagesDiscardEncryption          = 0xedd923c5
	crc_messagesSetEncryptedTyping         = 0x791451ed
	crc_messagesReadEncryptedHistory       = 0x7f4b690a
	crc_messagesSendEncrypted              = 0xa9776773
	crc_messagesSendEncryptedFile          = 0x9a901b66
	crc_messagesSendEncryptedService       = 0x32d439a4
	crc_messagesReceivedQueue              = 0x55a5bb66
	crc_uploadSaveBigFilePart              = 0xde7b673d
	crc_initConnection                     = 0xc7481da6
	crc_helpGetSupport                     = 0x9cdf08cd
	crc_authBindTempAuthKey                = 0xcdd42a05
	crc_contactsExportCard                 = 0x84e53737
	crc_contactsImportCard                 = 0x4fe196fe
	crc_messagesReadMessageContents        = 0x36a73f77
	crc_accountCheckUsername               = 0x2714d86c
	crc_accountUpdateUsername              = 0x3e0bdd7c
	crc_accountGetPrivacy                  = 0xdadbc950
	crc_accountSetPrivacy                  = 0xc9f81ce8
	crc_accountDeleteAccount               = 0x418d4e0b
	crc_accountGetAccountTTL               = 0x08fc711d
	crc_accountSetAccountTTL               = 0x2442485e
	crc_invokeWithLayer                    = 0xda9b0d0d
	crc_contactsResolveUsername            = 0xf93ccba3
	crc_accountSendChangePhoneCode         = 0x08e57deb
	crc_accountChangePhone                 = 0x70c32edb
	crc_messagesGetAllStickers             = 0x1c9618b1
	crc_accountUpdateDeviceLocked          = 0x38df3532
	crc_accountGetPassword                 = 0x548a30f5
	crc_authCheckPassword                  = 0x0a63011e
	crc_messagesGetWebPagePreview          = 0x25223e24
	crc_accountGetAuthorizations           = 0xe320c158
	crc_accountResetAuthorization          = 0xdf77f3bc
	crc_accountGetPasswordSettings         = 0xbc8d11bb
	crc_accountUpdatePasswordSettings      = 0xfa7c4b86
	crc_authRequestPasswordRecovery        = 0xd897bc66
	crc_authRecoverPassword                = 0x4ea56e92
	crc_invokeWithoutUpdates               = 0xbf9459b7
	crc_messagesExportChatInvite           = 0x7d885289
	crc_messagesCheckChatInvite            = 0x3eadb1bb
	crc_messagesImportChatInvite           = 0x6c50051c
	crc_messagesGetStickerSet              = 0x2619a90e
	crc_messagesInstallStickerSet          = 0xc78fe460
	crc_messagesUninstallStickerSet        = 0xf96e55de
	crc_authImportBotAuthorization         = 0x67a3ff2c
	crc_messagesStartBot                   = 0xe6df7378
	crc_helpGetAppChangelog                = 0x9010ef6f
	crc_messagesReportSpam                 = 0xcf1592db
	crc_messagesGetMessagesViews           = 0xc4c8a55d
	crc_updatesGetChannelDifference        = 0x03173d78
	crc_channelsReadHistory                = 0xcc104937
	crc_channelsDeleteMessages             = 0x84c1fd4e
	crc_channelsDeleteUserHistory          = 0xd10dd71b
	crc_channelsReportSpam                 = 0xfe087810
	crc_channelsGetMessages                = 0x93d7b347
	crc_channelsGetParticipants            = 0x24d98f92
	crc_channelsGetParticipant             = 0x546dd7a6
	crc_channelsGetChannels                = 0x0a7f6bbb
	crc_channelsGetFullChannel             = 0x08736a09
	crc_channelsCreateChannel              = 0xf4893d7f
	crc_channelsEditAbout                  = 0x13e27f1e
	crc_channelsEditAdmin                  = 0x20b88214
	crc_channelsEditTitle                  = 0x566decd0
	crc_channelsEditPhoto                  = 0xf12e57c9
	crc_channelsCheckUsername              = 0x10e6bd2c
	crc_channelsUpdateUsername             = 0x3514b3de
	crc_channelsJoinChannel                = 0x24b524c5
	crc_channelsLeaveChannel               = 0xf836aa95
	crc_channelsInviteToChannel            = 0x199f3a6c
	crc_channelsExportInvite               = 0xc7560885
	crc_channelsDeleteChannel              = 0xc0111fe3
	crc_messagesToggleChatAdmins           = 0xec8bd9e1
	crc_messagesEditChatAdmin              = 0xa9e69f2e
	crc_messagesMigrateChat                = 0x15a3b8e3
	crc_messagesSearchGlobal               = 0x9e3cacb0
	crc_accountReportPeer                  = 0xae189d5f
	crc_messagesReorderStickerSets         = 0x78337739
	crc_helpGetTermsOfService              = 0x350170f3
	crc_messagesGetDocumentByHash          = 0x338e2464
	crc_messagesSearchGifs                 = 0xbf9a776b
	crc_messagesGetSavedGifs               = 0x83bf3d52
	crc_messagesSaveGif                    = 0x327a30cb
	crc_messagesGetInlineBotResults        = 0x514e999d
	crc_messagesSetInlineBotResults        = 0xeb5ea206
	crc_messagesSendInlineBotResult        = 0xb16e06fe
	crc_channelsToggleInvites              = 0x49609307
	crc_channelsExportMessageLink          = 0xc846d22d
	crc_channelsToggleSignatures           = 0x1f69b606
	crc_messagesHideReportSpam             = 0xa8f1709b
	crc_messagesGetPeerSettings            = 0x3672e09c
	crc_channelsUpdatePinnedMessage        = 0xa72ded52
	crc_authResendCode                     = 0x3ef1a9bf
	crc_authCancelCode                     = 0x1f040578
	crc_messagesGetMessageEditData         = 0xfda68d36
	crc_messagesEditMessage                = 0xce91e4ca
	crc_messagesEditInlineBotMessage       = 0x130c2c85
	crc_messagesGetBotCallbackAnswer       = 0x810a9fec
	crc_messagesSetBotCallbackAnswer       = 0xd58f130a
	crc_contactsGetTopPeers                = 0xd4982db5
	crc_contactsResetTopPeerRating         = 0x1ae373ac
	crc_messagesGetPeerDialogs             = 0x2d9776b9
	crc_messagesSaveDraft                  = 0xbc39e14b
	crc_messagesGetAllDrafts               = 0x6a3f8d65
	crc_accountSendConfirmPhoneCode        = 0x1516d7bd
	crc_accountConfirmPhone                = 0x5f2178c3
	crc_messagesGetFeaturedStickers        = 0x2dacca4f
	crc_messagesReadFeaturedStickers       = 0x5b118126
	crc_messagesGetRecentStickers          = 0x5ea192c9
	crc_messagesSaveRecentSticker          = 0x392718f8
	crc_messagesClearRecentStickers        = 0x8999602d
	crc_messagesGetArchivedStickers        = 0x57f17692
	crc_channelsGetAdminedPublicChannels   = 0x8d8d82d7
	crc_authDropTempAuthKeys               = 0x8e48a188
	crc_messagesSetGameScore               = 0x8ef8ecc0
	crc_messagesSetInlineGameScore         = 0x15ad9f64
	crc_messagesGetMaskStickers            = 0x65b8c79f
	crc_messagesGetAttachedStickers        = 0xcc5b67cc
	crc_messagesGetGameHighScores          = 0xe822649d
	crc_messagesGetInlineGameHighScores    = 0x0f635e1b
	crc_messagesGetCommonChats             = 0x0d0a48c4
	crc_messagesGetAllChats                = 0xeba80ff0
	crc_helpSetBotUpdatesStatus            = 0xec22cfcd
	crc_messagesGetWebPage                 = 0x32ca8f91
	crc_messagesToggleDialogPin            = 0x3289be6a
	crc_messagesReorderPinnedDialogs       = 0x959ff644
	crc_messagesGetPinnedDialogs           = 0xe254d64e
	crc_phoneRequestCall                   = 0x5b95b3d4
	crc_phoneAcceptCall                    = 0x3bd2b4a0
	crc_phoneDiscardCall                   = 0x78d413a6
	crc_phoneReceivedCall                  = 0x17d54f61
	crc_messagesReportEncryptedSpam        = 0x4b0c8c0f
	crc_paymentsGetPaymentForm             = 0x99f09745
	crc_paymentsSendPaymentForm            = 0x2b8879b3
	crc_accountGetTmpPassword              = 0x4a82327e
	crc_messagesSetBotShippingResults      = 0xe5f672fa
	crc_messagesSetBotPrecheckoutResults   = 0x09c2dd95
	crc_uploadGetWebFile                   = 0x24e6818d
	crc_botsSendCustomRequest              = 0xaa2769ed
	crc_botsAnswerWebhookJSONQuery         = 0xe6213f4d
	crc_paymentsGetPaymentReceipt          = 0xa092a980
	crc_paymentsValidateRequestedInfo      = 0x770a8e74
	crc_paymentsGetSavedInfo               = 0x227d824b
	crc_paymentsClearSavedInfo             = 0xd83d70c1
	crc_phoneGetCallConfig                 = 0x55451fa9
	crc_phoneConfirmCall                   = 0x2efe1722
	crc_phoneSetCallRating                 = 0x1c536a34
	crc_phoneSaveCallDebug                 = 0x277add7e
	crc_uploadGetCdnFile                   = 0x2000bcc3
	crc_uploadReuploadCdnFile              = 0x1af91c09
	crc_helpGetCdnConfig                   = 0x52029342
	crc_messagesUploadMedia                = 0x519bc2b1
	crc_stickersCreateStickerSet           = 0x9bd86e6a
	crc_langpackGetLangPack                = 0x9ab5c58e
	crc_langpackGetStrings                 = 0x2e1ee318
	crc_langpackGetDifference              = 0x0b2e4d7d
	crc_langpackGetLanguages               = 0x800fd57d
	crc_channelsEditBanned                 = 0xbfd915cd
	crc_channelsGetAdminLog                = 0x33ddf480
	crc_stickersRemoveStickerFromSet       = 0xf7760f51
	crc_stickersChangeStickerPosition      = 0xffb6d4ca
	crc_stickersAddStickerToSet            = 0x8653febe
	crc_messagesSendScreenshotNotification = 0xc97df020
	crc_uploadGetCdnFileHashes             = 0xf715c87b
	crc_messagesGetUnreadMentions          = 0x46578472
	crc_messagesFaveSticker                = 0xb9ffc55b
	crc_channelsSetStickers                = 0xea8ca4f9
	crc_contactsResetSaved                 = 0x879537f1
	crc_messagesGetFavedStickers           = 0x21ce0b0e
	crc_channelsReadMessageContents        = 0xeab5dc38
)

// Encode funcs for types
func (e *TypeBool) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeBool_BoolFalse:
		return x.BoolFalse.encode()
	case *TypeBool_BoolTrue:
		return x.BoolTrue.encode()
	}
	return nil
}
func (e *TypeError) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeNull) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputPeer) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputPeer_InputPeerEmpty:
		return x.InputPeerEmpty.encode()
	case *TypeInputPeer_InputPeerSelf:
		return x.InputPeerSelf.encode()
	case *TypeInputPeer_InputPeerChat:
		return x.InputPeerChat.encode()
	case *TypeInputPeer_InputPeerUser:
		return x.InputPeerUser.encode()
	case *TypeInputPeer_InputPeerChannel:
		return x.InputPeerChannel.encode()
	}
	return nil
}
func (e *TypeInputUser) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputUser_InputUserEmpty:
		return x.InputUserEmpty.encode()
	case *TypeInputUser_InputUserSelf:
		return x.InputUserSelf.encode()
	case *TypeInputUser_InputUser:
		return x.InputUser.encode()
	}
	return nil
}
func (e *TypeInputContact) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputFile) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputFile_InputFile:
		return x.InputFile.encode()
	case *TypeInputFile_InputFileBig:
		return x.InputFileBig.encode()
	}
	return nil
}
func (e *TypeInputMedia) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputMedia_InputMediaEmpty:
		return x.InputMediaEmpty.encode()
	case *TypeInputMedia_InputMediaUploadedPhoto:
		return x.InputMediaUploadedPhoto.encode()
	case *TypeInputMedia_InputMediaPhoto:
		return x.InputMediaPhoto.encode()
	case *TypeInputMedia_InputMediaGeoPoint:
		return x.InputMediaGeoPoint.encode()
	case *TypeInputMedia_InputMediaContact:
		return x.InputMediaContact.encode()
	case *TypeInputMedia_InputMediaUploadedDocument:
		return x.InputMediaUploadedDocument.encode()
	case *TypeInputMedia_InputMediaDocument:
		return x.InputMediaDocument.encode()
	case *TypeInputMedia_InputMediaVenue:
		return x.InputMediaVenue.encode()
	case *TypeInputMedia_InputMediaGifExternal:
		return x.InputMediaGifExternal.encode()
	case *TypeInputMedia_InputMediaPhotoExternal:
		return x.InputMediaPhotoExternal.encode()
	case *TypeInputMedia_InputMediaDocumentExternal:
		return x.InputMediaDocumentExternal.encode()
	case *TypeInputMedia_InputMediaGame:
		return x.InputMediaGame.encode()
	case *TypeInputMedia_InputMediaInvoice:
		return x.InputMediaInvoice.encode()
	}
	return nil
}
func (e *TypeInputChatPhoto) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputChatPhoto_InputChatPhotoEmpty:
		return x.InputChatPhotoEmpty.encode()
	case *TypeInputChatPhoto_InputChatUploadedPhoto:
		return x.InputChatUploadedPhoto.encode()
	case *TypeInputChatPhoto_InputChatPhoto:
		return x.InputChatPhoto.encode()
	}
	return nil
}
func (e *TypeInputGeoPoint) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputGeoPoint_InputGeoPointEmpty:
		return x.InputGeoPointEmpty.encode()
	case *TypeInputGeoPoint_InputGeoPoint:
		return x.InputGeoPoint.encode()
	}
	return nil
}
func (e *TypeInputPhoto) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputPhoto_InputPhotoEmpty:
		return x.InputPhotoEmpty.encode()
	case *TypeInputPhoto_InputPhoto:
		return x.InputPhoto.encode()
	}
	return nil
}
func (e *TypeInputFileLocation) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputFileLocation_InputFileLocation:
		return x.InputFileLocation.encode()
	case *TypeInputFileLocation_InputEncryptedFileLocation:
		return x.InputEncryptedFileLocation.encode()
	case *TypeInputFileLocation_InputDocumentFileLocation:
		return x.InputDocumentFileLocation.encode()
	}
	return nil
}
func (e *TypeInputAppEvent) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePeer) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePeer_PeerUser:
		return x.PeerUser.encode()
	case *TypePeer_PeerChat:
		return x.PeerChat.encode()
	case *TypePeer_PeerChannel:
		return x.PeerChannel.encode()
	}
	return nil
}
func (e *TypeStorageFileType) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeStorageFileType_StorageFileUnknown:
		return x.StorageFileUnknown.encode()
	case *TypeStorageFileType_StorageFileJpeg:
		return x.StorageFileJpeg.encode()
	case *TypeStorageFileType_StorageFileGif:
		return x.StorageFileGif.encode()
	case *TypeStorageFileType_StorageFilePng:
		return x.StorageFilePng.encode()
	case *TypeStorageFileType_StorageFileMp3:
		return x.StorageFileMp3.encode()
	case *TypeStorageFileType_StorageFileMov:
		return x.StorageFileMov.encode()
	case *TypeStorageFileType_StorageFilePartial:
		return x.StorageFilePartial.encode()
	case *TypeStorageFileType_StorageFileMp4:
		return x.StorageFileMp4.encode()
	case *TypeStorageFileType_StorageFileWebp:
		return x.StorageFileWebp.encode()
	case *TypeStorageFileType_StorageFilePdf:
		return x.StorageFilePdf.encode()
	}
	return nil
}
func (e *TypeFileLocation) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeFileLocation_FileLocationUnavailable:
		return x.FileLocationUnavailable.encode()
	case *TypeFileLocation_FileLocation:
		return x.FileLocation.encode()
	}
	return nil
}
func (e *TypeUser) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUser_UserEmpty:
		return x.UserEmpty.encode()
	case *TypeUser_User:
		return x.User.encode()
	}
	return nil
}
func (e *TypeUserProfilePhoto) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUserProfilePhoto_UserProfilePhotoEmpty:
		return x.UserProfilePhotoEmpty.encode()
	case *TypeUserProfilePhoto_UserProfilePhoto:
		return x.UserProfilePhoto.encode()
	}
	return nil
}
func (e *TypeUserStatus) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUserStatus_UserStatusEmpty:
		return x.UserStatusEmpty.encode()
	case *TypeUserStatus_UserStatusOnline:
		return x.UserStatusOnline.encode()
	case *TypeUserStatus_UserStatusOffline:
		return x.UserStatusOffline.encode()
	case *TypeUserStatus_UserStatusRecently:
		return x.UserStatusRecently.encode()
	case *TypeUserStatus_UserStatusLastWeek:
		return x.UserStatusLastWeek.encode()
	case *TypeUserStatus_UserStatusLastMonth:
		return x.UserStatusLastMonth.encode()
	}
	return nil
}
func (e *TypeChat) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChat_ChatEmpty:
		return x.ChatEmpty.encode()
	case *TypeChat_Chat:
		return x.Chat.encode()
	case *TypeChat_ChatForbidden:
		return x.ChatForbidden.encode()
	case *TypeChat_Channel:
		return x.Channel.encode()
	case *TypeChat_ChannelForbidden:
		return x.ChannelForbidden.encode()
	}
	return nil
}
func (e *TypeChatFull) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChatFull_ChatFull:
		return x.ChatFull.encode()
	case *TypeChatFull_ChannelFull:
		return x.ChannelFull.encode()
	}
	return nil
}
func (e *TypeChatParticipant) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChatParticipant_ChatParticipant:
		return x.ChatParticipant.encode()
	case *TypeChatParticipant_ChatParticipantCreator:
		return x.ChatParticipantCreator.encode()
	case *TypeChatParticipant_ChatParticipantAdmin:
		return x.ChatParticipantAdmin.encode()
	}
	return nil
}
func (e *TypeChatParticipants) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChatParticipants_ChatParticipantsForbidden:
		return x.ChatParticipantsForbidden.encode()
	case *TypeChatParticipants_ChatParticipants:
		return x.ChatParticipants.encode()
	}
	return nil
}
func (e *TypeChatPhoto) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChatPhoto_ChatPhotoEmpty:
		return x.ChatPhotoEmpty.encode()
	case *TypeChatPhoto_ChatPhoto:
		return x.ChatPhoto.encode()
	}
	return nil
}
func (e *TypeMessage) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessage_MessageEmpty:
		return x.MessageEmpty.encode()
	case *TypeMessage_Message:
		return x.Message.encode()
	case *TypeMessage_MessageService:
		return x.MessageService.encode()
	}
	return nil
}
func (e *TypeMessageMedia) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessageMedia_MessageMediaEmpty:
		return x.MessageMediaEmpty.encode()
	case *TypeMessageMedia_MessageMediaPhoto:
		return x.MessageMediaPhoto.encode()
	case *TypeMessageMedia_MessageMediaGeo:
		return x.MessageMediaGeo.encode()
	case *TypeMessageMedia_MessageMediaContact:
		return x.MessageMediaContact.encode()
	case *TypeMessageMedia_MessageMediaUnsupported:
		return x.MessageMediaUnsupported.encode()
	case *TypeMessageMedia_MessageMediaDocument:
		return x.MessageMediaDocument.encode()
	case *TypeMessageMedia_MessageMediaWebPage:
		return x.MessageMediaWebPage.encode()
	case *TypeMessageMedia_MessageMediaVenue:
		return x.MessageMediaVenue.encode()
	case *TypeMessageMedia_MessageMediaGame:
		return x.MessageMediaGame.encode()
	case *TypeMessageMedia_MessageMediaInvoice:
		return x.MessageMediaInvoice.encode()
	}
	return nil
}
func (e *TypeMessageAction) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessageAction_MessageActionEmpty:
		return x.MessageActionEmpty.encode()
	case *TypeMessageAction_MessageActionChatCreate:
		return x.MessageActionChatCreate.encode()
	case *TypeMessageAction_MessageActionChatEditTitle:
		return x.MessageActionChatEditTitle.encode()
	case *TypeMessageAction_MessageActionChatEditPhoto:
		return x.MessageActionChatEditPhoto.encode()
	case *TypeMessageAction_MessageActionChatDeletePhoto:
		return x.MessageActionChatDeletePhoto.encode()
	case *TypeMessageAction_MessageActionChatAddUser:
		return x.MessageActionChatAddUser.encode()
	case *TypeMessageAction_MessageActionChatDeleteUser:
		return x.MessageActionChatDeleteUser.encode()
	case *TypeMessageAction_MessageActionChatJoinedByLink:
		return x.MessageActionChatJoinedByLink.encode()
	case *TypeMessageAction_MessageActionChannelCreate:
		return x.MessageActionChannelCreate.encode()
	case *TypeMessageAction_MessageActionChatMigrateTo:
		return x.MessageActionChatMigrateTo.encode()
	case *TypeMessageAction_MessageActionChannelMigrateFrom:
		return x.MessageActionChannelMigrateFrom.encode()
	case *TypeMessageAction_MessageActionPinMessage:
		return x.MessageActionPinMessage.encode()
	case *TypeMessageAction_MessageActionHistoryClear:
		return x.MessageActionHistoryClear.encode()
	case *TypeMessageAction_MessageActionGameScore:
		return x.MessageActionGameScore.encode()
	case *TypeMessageAction_MessageActionPhoneCall:
		return x.MessageActionPhoneCall.encode()
	case *TypeMessageAction_MessageActionPaymentSentMe:
		return x.MessageActionPaymentSentMe.encode()
	case *TypeMessageAction_MessageActionPaymentSent:
		return x.MessageActionPaymentSent.encode()
	case *TypeMessageAction_MessageActionScreenshotTaken:
		return x.MessageActionScreenshotTaken.encode()
	}
	return nil
}
func (e *TypeDialog) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePhoto) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePhoto_PhotoEmpty:
		return x.PhotoEmpty.encode()
	case *TypePhoto_Photo:
		return x.Photo.encode()
	}
	return nil
}
func (e *TypePhotoSize) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePhotoSize_PhotoSizeEmpty:
		return x.PhotoSizeEmpty.encode()
	case *TypePhotoSize_PhotoSize:
		return x.PhotoSize.encode()
	case *TypePhotoSize_PhotoCachedSize:
		return x.PhotoCachedSize.encode()
	}
	return nil
}
func (e *TypeGeoPoint) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeGeoPoint_GeoPointEmpty:
		return x.GeoPointEmpty.encode()
	case *TypeGeoPoint_GeoPoint:
		return x.GeoPoint.encode()
	}
	return nil
}
func (e *TypeAuthCheckedPhone) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAuthSentCode) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAuthAuthorization) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAuthExportedAuthorization) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputNotifyPeer) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputNotifyPeer_InputNotifyPeer:
		return x.InputNotifyPeer.encode()
	case *TypeInputNotifyPeer_InputNotifyUsers:
		return x.InputNotifyUsers.encode()
	case *TypeInputNotifyPeer_InputNotifyChats:
		return x.InputNotifyChats.encode()
	case *TypeInputNotifyPeer_InputNotifyAll:
		return x.InputNotifyAll.encode()
	}
	return nil
}
func (e *TypeInputPeerNotifySettings) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePeerNotifyEvents) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePeerNotifyEvents_PeerNotifyEventsEmpty:
		return x.PeerNotifyEventsEmpty.encode()
	case *TypePeerNotifyEvents_PeerNotifyEventsAll:
		return x.PeerNotifyEventsAll.encode()
	}
	return nil
}
func (e *TypePeerNotifySettings) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePeerNotifySettings_PeerNotifySettingsEmpty:
		return x.PeerNotifySettingsEmpty.encode()
	case *TypePeerNotifySettings_PeerNotifySettings:
		return x.PeerNotifySettings.encode()
	}
	return nil
}
func (e *TypeWallPaper) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeWallPaper_WallPaper:
		return x.WallPaper.encode()
	case *TypeWallPaper_WallPaperSolid:
		return x.WallPaperSolid.encode()
	}
	return nil
}
func (e *TypeUserFull) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeContact) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeImportedContact) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeContactBlocked) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeContactStatus) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeContactsLink) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeContactsContacts) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeContactsContacts_ContactsContacts:
		return x.ContactsContacts.encode()
	case *TypeContactsContacts_ContactsContactsNotModified:
		return x.ContactsContactsNotModified.encode()
	}
	return nil
}
func (e *TypeContactsImportedContacts) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeContactsBlocked) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeContactsBlocked_ContactsBlocked:
		return x.ContactsBlocked.encode()
	case *TypeContactsBlocked_ContactsBlockedSlice:
		return x.ContactsBlockedSlice.encode()
	}
	return nil
}
func (e *TypeContactsFound) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesDialogs) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesDialogs_MessagesDialogs:
		return x.MessagesDialogs.encode()
	case *TypeMessagesDialogs_MessagesDialogsSlice:
		return x.MessagesDialogsSlice.encode()
	}
	return nil
}
func (e *TypeMessagesMessages) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesMessages_MessagesMessages:
		return x.MessagesMessages.encode()
	case *TypeMessagesMessages_MessagesMessagesSlice:
		return x.MessagesMessagesSlice.encode()
	case *TypeMessagesMessages_MessagesChannelMessages:
		return x.MessagesChannelMessages.encode()
	}
	return nil
}
func (e *TypeMessagesChats) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesChats_MessagesChats:
		return x.MessagesChats.encode()
	case *TypeMessagesChats_MessagesChatsSlice:
		return x.MessagesChatsSlice.encode()
	}
	return nil
}
func (e *TypeMessagesChatFull) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesAffectedHistory) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesFilter) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesFilter_InputMessagesFilterEmpty:
		return x.InputMessagesFilterEmpty.encode()
	case *TypeMessagesFilter_InputMessagesFilterPhotos:
		return x.InputMessagesFilterPhotos.encode()
	case *TypeMessagesFilter_InputMessagesFilterVideo:
		return x.InputMessagesFilterVideo.encode()
	case *TypeMessagesFilter_InputMessagesFilterPhotoVideo:
		return x.InputMessagesFilterPhotoVideo.encode()
	case *TypeMessagesFilter_InputMessagesFilterDocument:
		return x.InputMessagesFilterDocument.encode()
	case *TypeMessagesFilter_InputMessagesFilterPhotoVideoDocuments:
		return x.InputMessagesFilterPhotoVideoDocuments.encode()
	case *TypeMessagesFilter_InputMessagesFilterUrl:
		return x.InputMessagesFilterUrl.encode()
	case *TypeMessagesFilter_InputMessagesFilterGif:
		return x.InputMessagesFilterGif.encode()
	case *TypeMessagesFilter_InputMessagesFilterVoice:
		return x.InputMessagesFilterVoice.encode()
	case *TypeMessagesFilter_InputMessagesFilterMusic:
		return x.InputMessagesFilterMusic.encode()
	case *TypeMessagesFilter_InputMessagesFilterChatPhotos:
		return x.InputMessagesFilterChatPhotos.encode()
	case *TypeMessagesFilter_InputMessagesFilterPhoneCalls:
		return x.InputMessagesFilterPhoneCalls.encode()
	case *TypeMessagesFilter_InputMessagesFilterRoundVoice:
		return x.InputMessagesFilterRoundVoice.encode()
	case *TypeMessagesFilter_InputMessagesFilterRoundVideo:
		return x.InputMessagesFilterRoundVideo.encode()
	case *TypeMessagesFilter_InputMessagesFilterMyMentions:
		return x.InputMessagesFilterMyMentions.encode()
	case *TypeMessagesFilter_InputMessagesFilterMyMentionsUnread:
		return x.InputMessagesFilterMyMentionsUnread.encode()
	}
	return nil
}
func (e *TypeUpdate) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUpdate_UpdateNewMessage:
		return x.UpdateNewMessage.encode()
	case *TypeUpdate_UpdateMessageID:
		return x.UpdateMessageID.encode()
	case *TypeUpdate_UpdateDeleteMessages:
		return x.UpdateDeleteMessages.encode()
	case *TypeUpdate_UpdateUserTyping:
		return x.UpdateUserTyping.encode()
	case *TypeUpdate_UpdateChatUserTyping:
		return x.UpdateChatUserTyping.encode()
	case *TypeUpdate_UpdateChatParticipants:
		return x.UpdateChatParticipants.encode()
	case *TypeUpdate_UpdateUserStatus:
		return x.UpdateUserStatus.encode()
	case *TypeUpdate_UpdateUserName:
		return x.UpdateUserName.encode()
	case *TypeUpdate_UpdateUserPhoto:
		return x.UpdateUserPhoto.encode()
	case *TypeUpdate_UpdateContactRegistered:
		return x.UpdateContactRegistered.encode()
	case *TypeUpdate_UpdateContactLink:
		return x.UpdateContactLink.encode()
	case *TypeUpdate_UpdateNewEncryptedMessage:
		return x.UpdateNewEncryptedMessage.encode()
	case *TypeUpdate_UpdateEncryptedChatTyping:
		return x.UpdateEncryptedChatTyping.encode()
	case *TypeUpdate_UpdateEncryption:
		return x.UpdateEncryption.encode()
	case *TypeUpdate_UpdateEncryptedMessagesRead:
		return x.UpdateEncryptedMessagesRead.encode()
	case *TypeUpdate_UpdateChatParticipantAdd:
		return x.UpdateChatParticipantAdd.encode()
	case *TypeUpdate_UpdateChatParticipantDelete:
		return x.UpdateChatParticipantDelete.encode()
	case *TypeUpdate_UpdateDcOptions:
		return x.UpdateDcOptions.encode()
	case *TypeUpdate_UpdateUserBlocked:
		return x.UpdateUserBlocked.encode()
	case *TypeUpdate_UpdateNotifySettings:
		return x.UpdateNotifySettings.encode()
	case *TypeUpdate_UpdateServiceNotification:
		return x.UpdateServiceNotification.encode()
	case *TypeUpdate_UpdatePrivacy:
		return x.UpdatePrivacy.encode()
	case *TypeUpdate_UpdateUserPhone:
		return x.UpdateUserPhone.encode()
	case *TypeUpdate_UpdateReadHistoryInbox:
		return x.UpdateReadHistoryInbox.encode()
	case *TypeUpdate_UpdateReadHistoryOutbox:
		return x.UpdateReadHistoryOutbox.encode()
	case *TypeUpdate_UpdateWebPage:
		return x.UpdateWebPage.encode()
	case *TypeUpdate_UpdateReadMessagesContents:
		return x.UpdateReadMessagesContents.encode()
	case *TypeUpdate_UpdateChannelTooLong:
		return x.UpdateChannelTooLong.encode()
	case *TypeUpdate_UpdateChannel:
		return x.UpdateChannel.encode()
	case *TypeUpdate_UpdateNewChannelMessage:
		return x.UpdateNewChannelMessage.encode()
	case *TypeUpdate_UpdateReadChannelInbox:
		return x.UpdateReadChannelInbox.encode()
	case *TypeUpdate_UpdateDeleteChannelMessages:
		return x.UpdateDeleteChannelMessages.encode()
	case *TypeUpdate_UpdateChannelMessageViews:
		return x.UpdateChannelMessageViews.encode()
	case *TypeUpdate_UpdateChatAdmins:
		return x.UpdateChatAdmins.encode()
	case *TypeUpdate_UpdateChatParticipantAdmin:
		return x.UpdateChatParticipantAdmin.encode()
	case *TypeUpdate_UpdateNewStickerSet:
		return x.UpdateNewStickerSet.encode()
	case *TypeUpdate_UpdateStickerSetsOrder:
		return x.UpdateStickerSetsOrder.encode()
	case *TypeUpdate_UpdateStickerSets:
		return x.UpdateStickerSets.encode()
	case *TypeUpdate_UpdateSavedGifs:
		return x.UpdateSavedGifs.encode()
	case *TypeUpdate_UpdateBotInlineQuery:
		return x.UpdateBotInlineQuery.encode()
	case *TypeUpdate_UpdateBotInlineSend:
		return x.UpdateBotInlineSend.encode()
	case *TypeUpdate_UpdateEditChannelMessage:
		return x.UpdateEditChannelMessage.encode()
	case *TypeUpdate_UpdateChannelPinnedMessage:
		return x.UpdateChannelPinnedMessage.encode()
	case *TypeUpdate_UpdateBotCallbackQuery:
		return x.UpdateBotCallbackQuery.encode()
	case *TypeUpdate_UpdateEditMessage:
		return x.UpdateEditMessage.encode()
	case *TypeUpdate_UpdateInlineBotCallbackQuery:
		return x.UpdateInlineBotCallbackQuery.encode()
	case *TypeUpdate_UpdateReadChannelOutbox:
		return x.UpdateReadChannelOutbox.encode()
	case *TypeUpdate_UpdateDraftMessage:
		return x.UpdateDraftMessage.encode()
	case *TypeUpdate_UpdateReadFeaturedStickers:
		return x.UpdateReadFeaturedStickers.encode()
	case *TypeUpdate_UpdateRecentStickers:
		return x.UpdateRecentStickers.encode()
	case *TypeUpdate_UpdateConfig:
		return x.UpdateConfig.encode()
	case *TypeUpdate_UpdatePtsChanged:
		return x.UpdatePtsChanged.encode()
	case *TypeUpdate_UpdateChannelWebPage:
		return x.UpdateChannelWebPage.encode()
	case *TypeUpdate_UpdatePhoneCall:
		return x.UpdatePhoneCall.encode()
	case *TypeUpdate_UpdateDialogPinned:
		return x.UpdateDialogPinned.encode()
	case *TypeUpdate_UpdatePinnedDialogs:
		return x.UpdatePinnedDialogs.encode()
	case *TypeUpdate_UpdateBotWebhookJSON:
		return x.UpdateBotWebhookJSON.encode()
	case *TypeUpdate_UpdateBotWebhookJSONQuery:
		return x.UpdateBotWebhookJSONQuery.encode()
	case *TypeUpdate_UpdateBotShippingQuery:
		return x.UpdateBotShippingQuery.encode()
	case *TypeUpdate_UpdateBotPrecheckoutQuery:
		return x.UpdateBotPrecheckoutQuery.encode()
	case *TypeUpdate_UpdateLangPackTooLong:
		return x.UpdateLangPackTooLong.encode()
	case *TypeUpdate_UpdateLangPack:
		return x.UpdateLangPack.encode()
	case *TypeUpdate_UpdateContactsReset:
		return x.UpdateContactsReset.encode()
	case *TypeUpdate_UpdateFavedStickers:
		return x.UpdateFavedStickers.encode()
	case *TypeUpdate_UpdateChannelReadMessagesContents:
		return x.UpdateChannelReadMessagesContents.encode()
	}
	return nil
}
func (e *TypeUpdatesState) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeUpdatesDifference) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUpdatesDifference_UpdatesDifferenceEmpty:
		return x.UpdatesDifferenceEmpty.encode()
	case *TypeUpdatesDifference_UpdatesDifference:
		return x.UpdatesDifference.encode()
	case *TypeUpdatesDifference_UpdatesDifferenceSlice:
		return x.UpdatesDifferenceSlice.encode()
	case *TypeUpdatesDifference_UpdatesDifferenceTooLong:
		return x.UpdatesDifferenceTooLong.encode()
	}
	return nil
}
func (e *TypeUpdates) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUpdates_UpdatesTooLong:
		return x.UpdatesTooLong.encode()
	case *TypeUpdates_UpdateShortMessage:
		return x.UpdateShortMessage.encode()
	case *TypeUpdates_UpdateShortChatMessage:
		return x.UpdateShortChatMessage.encode()
	case *TypeUpdates_UpdateShort:
		return x.UpdateShort.encode()
	case *TypeUpdates_UpdatesCombined:
		return x.UpdatesCombined.encode()
	case *TypeUpdates_Updates:
		return x.Updates.encode()
	case *TypeUpdates_UpdateShortSentMessage:
		return x.UpdateShortSentMessage.encode()
	}
	return nil
}
func (e *TypePhotosPhoto) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeUploadFile) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUploadFile_UploadFile:
		return x.UploadFile.encode()
	case *TypeUploadFile_UploadFileCdnRedirect:
		return x.UploadFileCdnRedirect.encode()
	}
	return nil
}
func (e *TypeDcOption) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeConfig) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeNearestDc) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeHelpAppUpdate) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeHelpAppUpdate_HelpAppUpdate:
		return x.HelpAppUpdate.encode()
	case *TypeHelpAppUpdate_HelpNoAppUpdate:
		return x.HelpNoAppUpdate.encode()
	}
	return nil
}
func (e *TypeHelpInviteText) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputPeerNotifyEvents) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputPeerNotifyEvents_InputPeerNotifyEventsEmpty:
		return x.InputPeerNotifyEventsEmpty.encode()
	case *TypeInputPeerNotifyEvents_InputPeerNotifyEventsAll:
		return x.InputPeerNotifyEventsAll.encode()
	}
	return nil
}
func (e *TypePhotosPhotos) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePhotosPhotos_PhotosPhotos:
		return x.PhotosPhotos.encode()
	case *TypePhotosPhotos_PhotosPhotosSlice:
		return x.PhotosPhotosSlice.encode()
	}
	return nil
}
func (e *TypeEncryptedChat) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeEncryptedChat_EncryptedChatEmpty:
		return x.EncryptedChatEmpty.encode()
	case *TypeEncryptedChat_EncryptedChatWaiting:
		return x.EncryptedChatWaiting.encode()
	case *TypeEncryptedChat_EncryptedChatRequested:
		return x.EncryptedChatRequested.encode()
	case *TypeEncryptedChat_EncryptedChat:
		return x.EncryptedChat.encode()
	case *TypeEncryptedChat_EncryptedChatDiscarded:
		return x.EncryptedChatDiscarded.encode()
	}
	return nil
}
func (e *TypeInputEncryptedChat) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeEncryptedFile) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeEncryptedFile_EncryptedFileEmpty:
		return x.EncryptedFileEmpty.encode()
	case *TypeEncryptedFile_EncryptedFile:
		return x.EncryptedFile.encode()
	}
	return nil
}
func (e *TypeInputEncryptedFile) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputEncryptedFile_InputEncryptedFileEmpty:
		return x.InputEncryptedFileEmpty.encode()
	case *TypeInputEncryptedFile_InputEncryptedFileUploaded:
		return x.InputEncryptedFileUploaded.encode()
	case *TypeInputEncryptedFile_InputEncryptedFile:
		return x.InputEncryptedFile.encode()
	case *TypeInputEncryptedFile_InputEncryptedFileBigUploaded:
		return x.InputEncryptedFileBigUploaded.encode()
	}
	return nil
}
func (e *TypeEncryptedMessage) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeEncryptedMessage_EncryptedMessage:
		return x.EncryptedMessage.encode()
	case *TypeEncryptedMessage_EncryptedMessageService:
		return x.EncryptedMessageService.encode()
	}
	return nil
}
func (e *TypeMessagesDhConfig) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesDhConfig_MessagesDhConfigNotModified:
		return x.MessagesDhConfigNotModified.encode()
	case *TypeMessagesDhConfig_MessagesDhConfig:
		return x.MessagesDhConfig.encode()
	}
	return nil
}
func (e *TypeMessagesSentEncryptedMessage) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesSentEncryptedMessage_MessagesSentEncryptedMessage:
		return x.MessagesSentEncryptedMessage.encode()
	case *TypeMessagesSentEncryptedMessage_MessagesSentEncryptedFile:
		return x.MessagesSentEncryptedFile.encode()
	}
	return nil
}
func (e *TypeInputDocument) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputDocument_InputDocumentEmpty:
		return x.InputDocumentEmpty.encode()
	case *TypeInputDocument_InputDocument:
		return x.InputDocument.encode()
	}
	return nil
}
func (e *TypeDocument) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeDocument_DocumentEmpty:
		return x.DocumentEmpty.encode()
	case *TypeDocument_Document:
		return x.Document.encode()
	}
	return nil
}
func (e *TypeHelpSupport) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeNotifyPeer) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeNotifyPeer_NotifyAll:
		return x.NotifyAll.encode()
	case *TypeNotifyPeer_NotifyChats:
		return x.NotifyChats.encode()
	case *TypeNotifyPeer_NotifyPeer:
		return x.NotifyPeer.encode()
	case *TypeNotifyPeer_NotifyUsers:
		return x.NotifyUsers.encode()
	}
	return nil
}
func (e *TypeSendMessageAction) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeSendMessageAction_SendMessageTypingAction:
		return x.SendMessageTypingAction.encode()
	case *TypeSendMessageAction_SendMessageCancelAction:
		return x.SendMessageCancelAction.encode()
	case *TypeSendMessageAction_SendMessageRecordVideoAction:
		return x.SendMessageRecordVideoAction.encode()
	case *TypeSendMessageAction_SendMessageUploadVideoAction:
		return x.SendMessageUploadVideoAction.encode()
	case *TypeSendMessageAction_SendMessageRecordAudioAction:
		return x.SendMessageRecordAudioAction.encode()
	case *TypeSendMessageAction_SendMessageUploadAudioAction:
		return x.SendMessageUploadAudioAction.encode()
	case *TypeSendMessageAction_SendMessageUploadPhotoAction:
		return x.SendMessageUploadPhotoAction.encode()
	case *TypeSendMessageAction_SendMessageUploadDocumentAction:
		return x.SendMessageUploadDocumentAction.encode()
	case *TypeSendMessageAction_SendMessageGeoLocationAction:
		return x.SendMessageGeoLocationAction.encode()
	case *TypeSendMessageAction_SendMessageChooseContactAction:
		return x.SendMessageChooseContactAction.encode()
	case *TypeSendMessageAction_SendMessageGamePlayAction:
		return x.SendMessageGamePlayAction.encode()
	case *TypeSendMessageAction_SendMessageRecordRoundAction:
		return x.SendMessageRecordRoundAction.encode()
	case *TypeSendMessageAction_SendMessageUploadRoundAction:
		return x.SendMessageUploadRoundAction.encode()
	}
	return nil
}
func (e *TypeInputPrivacyKey) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputPrivacyKey_InputPrivacyKeyStatusTimestamp:
		return x.InputPrivacyKeyStatusTimestamp.encode()
	case *TypeInputPrivacyKey_InputPrivacyKeyChatInvite:
		return x.InputPrivacyKeyChatInvite.encode()
	case *TypeInputPrivacyKey_InputPrivacyKeyPhoneCall:
		return x.InputPrivacyKeyPhoneCall.encode()
	}
	return nil
}
func (e *TypePrivacyKey) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePrivacyKey_PrivacyKeyStatusTimestamp:
		return x.PrivacyKeyStatusTimestamp.encode()
	case *TypePrivacyKey_PrivacyKeyChatInvite:
		return x.PrivacyKeyChatInvite.encode()
	case *TypePrivacyKey_PrivacyKeyPhoneCall:
		return x.PrivacyKeyPhoneCall.encode()
	}
	return nil
}
func (e *TypeInputPrivacyRule) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputPrivacyRule_InputPrivacyValueAllowContacts:
		return x.InputPrivacyValueAllowContacts.encode()
	case *TypeInputPrivacyRule_InputPrivacyValueAllowAll:
		return x.InputPrivacyValueAllowAll.encode()
	case *TypeInputPrivacyRule_InputPrivacyValueAllowUsers:
		return x.InputPrivacyValueAllowUsers.encode()
	case *TypeInputPrivacyRule_InputPrivacyValueDisallowContacts:
		return x.InputPrivacyValueDisallowContacts.encode()
	case *TypeInputPrivacyRule_InputPrivacyValueDisallowAll:
		return x.InputPrivacyValueDisallowAll.encode()
	case *TypeInputPrivacyRule_InputPrivacyValueDisallowUsers:
		return x.InputPrivacyValueDisallowUsers.encode()
	}
	return nil
}
func (e *TypePrivacyRule) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePrivacyRule_PrivacyValueAllowContacts:
		return x.PrivacyValueAllowContacts.encode()
	case *TypePrivacyRule_PrivacyValueAllowAll:
		return x.PrivacyValueAllowAll.encode()
	case *TypePrivacyRule_PrivacyValueAllowUsers:
		return x.PrivacyValueAllowUsers.encode()
	case *TypePrivacyRule_PrivacyValueDisallowContacts:
		return x.PrivacyValueDisallowContacts.encode()
	case *TypePrivacyRule_PrivacyValueDisallowAll:
		return x.PrivacyValueDisallowAll.encode()
	case *TypePrivacyRule_PrivacyValueDisallowUsers:
		return x.PrivacyValueDisallowUsers.encode()
	}
	return nil
}
func (e *TypeAccountPrivacyRules) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAccountDaysTTL) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeDisabledFeature) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeDocumentAttribute) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeDocumentAttribute_DocumentAttributeImageSize:
		return x.DocumentAttributeImageSize.encode()
	case *TypeDocumentAttribute_DocumentAttributeAnimated:
		return x.DocumentAttributeAnimated.encode()
	case *TypeDocumentAttribute_DocumentAttributeSticker:
		return x.DocumentAttributeSticker.encode()
	case *TypeDocumentAttribute_DocumentAttributeVideo:
		return x.DocumentAttributeVideo.encode()
	case *TypeDocumentAttribute_DocumentAttributeAudio:
		return x.DocumentAttributeAudio.encode()
	case *TypeDocumentAttribute_DocumentAttributeFilename:
		return x.DocumentAttributeFilename.encode()
	case *TypeDocumentAttribute_DocumentAttributeHasStickers:
		return x.DocumentAttributeHasStickers.encode()
	}
	return nil
}
func (e *TypeMessagesStickers) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesStickers_MessagesStickersNotModified:
		return x.MessagesStickersNotModified.encode()
	case *TypeMessagesStickers_MessagesStickers:
		return x.MessagesStickers.encode()
	}
	return nil
}
func (e *TypeStickerPack) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesAllStickers) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesAllStickers_MessagesAllStickersNotModified:
		return x.MessagesAllStickersNotModified.encode()
	case *TypeMessagesAllStickers_MessagesAllStickers:
		return x.MessagesAllStickers.encode()
	}
	return nil
}
func (e *TypeAccountPassword) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeAccountPassword_AccountNoPassword:
		return x.AccountNoPassword.encode()
	case *TypeAccountPassword_AccountPassword:
		return x.AccountPassword.encode()
	}
	return nil
}
func (e *TypeMessagesAffectedMessages) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeContactLink) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeContactLink_ContactLinkUnknown:
		return x.ContactLinkUnknown.encode()
	case *TypeContactLink_ContactLinkNone:
		return x.ContactLinkNone.encode()
	case *TypeContactLink_ContactLinkHasPhone:
		return x.ContactLinkHasPhone.encode()
	case *TypeContactLink_ContactLinkContact:
		return x.ContactLinkContact.encode()
	}
	return nil
}
func (e *TypeWebPage) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeWebPage_WebPageEmpty:
		return x.WebPageEmpty.encode()
	case *TypeWebPage_WebPagePending:
		return x.WebPagePending.encode()
	case *TypeWebPage_WebPage:
		return x.WebPage.encode()
	case *TypeWebPage_WebPageNotModified:
		return x.WebPageNotModified.encode()
	}
	return nil
}
func (e *TypeAuthorization) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAccountAuthorizations) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAccountPasswordSettings) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAccountPasswordInputSettings) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAuthPasswordRecovery) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeReceivedNotifyMessage) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeExportedChatInvite) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeExportedChatInvite_ChatInviteEmpty:
		return x.ChatInviteEmpty.encode()
	case *TypeExportedChatInvite_ChatInviteExported:
		return x.ChatInviteExported.encode()
	}
	return nil
}
func (e *TypeChatInvite) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChatInvite_ChatInviteAlready:
		return x.ChatInviteAlready.encode()
	case *TypeChatInvite_ChatInvite:
		return x.ChatInvite.encode()
	}
	return nil
}
func (e *TypeInputStickerSet) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputStickerSet_InputStickerSetEmpty:
		return x.InputStickerSetEmpty.encode()
	case *TypeInputStickerSet_InputStickerSetID:
		return x.InputStickerSetID.encode()
	case *TypeInputStickerSet_InputStickerSetShortName:
		return x.InputStickerSetShortName.encode()
	}
	return nil
}
func (e *TypeStickerSet) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesStickerSet) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeBotCommand) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeBotInfo) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeKeyboardButton) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeKeyboardButton_KeyboardButton:
		return x.KeyboardButton.encode()
	case *TypeKeyboardButton_KeyboardButtonUrl:
		return x.KeyboardButtonUrl.encode()
	case *TypeKeyboardButton_KeyboardButtonCallback:
		return x.KeyboardButtonCallback.encode()
	case *TypeKeyboardButton_KeyboardButtonRequestPhone:
		return x.KeyboardButtonRequestPhone.encode()
	case *TypeKeyboardButton_KeyboardButtonRequestGeoLocation:
		return x.KeyboardButtonRequestGeoLocation.encode()
	case *TypeKeyboardButton_KeyboardButtonSwitchInline:
		return x.KeyboardButtonSwitchInline.encode()
	case *TypeKeyboardButton_KeyboardButtonGame:
		return x.KeyboardButtonGame.encode()
	case *TypeKeyboardButton_KeyboardButtonBuy:
		return x.KeyboardButtonBuy.encode()
	}
	return nil
}
func (e *TypeKeyboardButtonRow) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeReplyMarkup) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeReplyMarkup_ReplyKeyboardHide:
		return x.ReplyKeyboardHide.encode()
	case *TypeReplyMarkup_ReplyKeyboardForceReply:
		return x.ReplyKeyboardForceReply.encode()
	case *TypeReplyMarkup_ReplyKeyboardMarkup:
		return x.ReplyKeyboardMarkup.encode()
	case *TypeReplyMarkup_ReplyInlineMarkup:
		return x.ReplyInlineMarkup.encode()
	}
	return nil
}
func (e *TypeMessageEntity) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessageEntity_MessageEntityUnknown:
		return x.MessageEntityUnknown.encode()
	case *TypeMessageEntity_MessageEntityMention:
		return x.MessageEntityMention.encode()
	case *TypeMessageEntity_MessageEntityHashtag:
		return x.MessageEntityHashtag.encode()
	case *TypeMessageEntity_MessageEntityBotCommand:
		return x.MessageEntityBotCommand.encode()
	case *TypeMessageEntity_MessageEntityUrl:
		return x.MessageEntityUrl.encode()
	case *TypeMessageEntity_MessageEntityEmail:
		return x.MessageEntityEmail.encode()
	case *TypeMessageEntity_MessageEntityBold:
		return x.MessageEntityBold.encode()
	case *TypeMessageEntity_MessageEntityItalic:
		return x.MessageEntityItalic.encode()
	case *TypeMessageEntity_MessageEntityCode:
		return x.MessageEntityCode.encode()
	case *TypeMessageEntity_MessageEntityPre:
		return x.MessageEntityPre.encode()
	case *TypeMessageEntity_MessageEntityTextUrl:
		return x.MessageEntityTextUrl.encode()
	case *TypeMessageEntity_MessageEntityMentionName:
		return x.MessageEntityMentionName.encode()
	case *TypeMessageEntity_InputMessageEntityMentionName:
		return x.InputMessageEntityMentionName.encode()
	}
	return nil
}
func (e *TypeInputChannel) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputChannel_InputChannelEmpty:
		return x.InputChannelEmpty.encode()
	case *TypeInputChannel_InputChannel:
		return x.InputChannel.encode()
	}
	return nil
}
func (e *TypeContactsResolvedPeer) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessageRange) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeUpdatesChannelDifference) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUpdatesChannelDifference_UpdatesChannelDifferenceEmpty:
		return x.UpdatesChannelDifferenceEmpty.encode()
	case *TypeUpdatesChannelDifference_UpdatesChannelDifferenceTooLong:
		return x.UpdatesChannelDifferenceTooLong.encode()
	case *TypeUpdatesChannelDifference_UpdatesChannelDifference:
		return x.UpdatesChannelDifference.encode()
	}
	return nil
}
func (e *TypeChannelMessagesFilter) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChannelMessagesFilter_ChannelMessagesFilterEmpty:
		return x.ChannelMessagesFilterEmpty.encode()
	case *TypeChannelMessagesFilter_ChannelMessagesFilter:
		return x.ChannelMessagesFilter.encode()
	}
	return nil
}
func (e *TypeChannelParticipant) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChannelParticipant_ChannelParticipant:
		return x.ChannelParticipant.encode()
	case *TypeChannelParticipant_ChannelParticipantSelf:
		return x.ChannelParticipantSelf.encode()
	case *TypeChannelParticipant_ChannelParticipantCreator:
		return x.ChannelParticipantCreator.encode()
	case *TypeChannelParticipant_ChannelParticipantAdmin:
		return x.ChannelParticipantAdmin.encode()
	case *TypeChannelParticipant_ChannelParticipantBanned:
		return x.ChannelParticipantBanned.encode()
	}
	return nil
}
func (e *TypeChannelParticipantsFilter) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChannelParticipantsFilter_ChannelParticipantsRecent:
		return x.ChannelParticipantsRecent.encode()
	case *TypeChannelParticipantsFilter_ChannelParticipantsAdmins:
		return x.ChannelParticipantsAdmins.encode()
	case *TypeChannelParticipantsFilter_ChannelParticipantsKicked:
		return x.ChannelParticipantsKicked.encode()
	case *TypeChannelParticipantsFilter_ChannelParticipantsBots:
		return x.ChannelParticipantsBots.encode()
	case *TypeChannelParticipantsFilter_ChannelParticipantsBanned:
		return x.ChannelParticipantsBanned.encode()
	case *TypeChannelParticipantsFilter_ChannelParticipantsSearch:
		return x.ChannelParticipantsSearch.encode()
	}
	return nil
}
func (e *TypeChannelsChannelParticipants) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeChannelsChannelParticipant) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeTrue) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeReportReason) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeReportReason_InputReportReasonSpam:
		return x.InputReportReasonSpam.encode()
	case *TypeReportReason_InputReportReasonViolence:
		return x.InputReportReasonViolence.encode()
	case *TypeReportReason_InputReportReasonPornography:
		return x.InputReportReasonPornography.encode()
	case *TypeReportReason_InputReportReasonOther:
		return x.InputReportReasonOther.encode()
	}
	return nil
}
func (e *TypeHelpTermsOfService) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeFoundGif) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeFoundGif_FoundGif:
		return x.FoundGif.encode()
	case *TypeFoundGif_FoundGifCached:
		return x.FoundGifCached.encode()
	}
	return nil
}
func (e *TypeMessagesFoundGifs) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesSavedGifs) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesSavedGifs_MessagesSavedGifsNotModified:
		return x.MessagesSavedGifsNotModified.encode()
	case *TypeMessagesSavedGifs_MessagesSavedGifs:
		return x.MessagesSavedGifs.encode()
	}
	return nil
}
func (e *TypeInputBotInlineMessage) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputBotInlineMessage_InputBotInlineMessageMediaAuto:
		return x.InputBotInlineMessageMediaAuto.encode()
	case *TypeInputBotInlineMessage_InputBotInlineMessageText:
		return x.InputBotInlineMessageText.encode()
	case *TypeInputBotInlineMessage_InputBotInlineMessageMediaGeo:
		return x.InputBotInlineMessageMediaGeo.encode()
	case *TypeInputBotInlineMessage_InputBotInlineMessageMediaVenue:
		return x.InputBotInlineMessageMediaVenue.encode()
	case *TypeInputBotInlineMessage_InputBotInlineMessageMediaContact:
		return x.InputBotInlineMessageMediaContact.encode()
	case *TypeInputBotInlineMessage_InputBotInlineMessageGame:
		return x.InputBotInlineMessageGame.encode()
	}
	return nil
}
func (e *TypeInputBotInlineResult) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputBotInlineResult_InputBotInlineResult:
		return x.InputBotInlineResult.encode()
	case *TypeInputBotInlineResult_InputBotInlineResultPhoto:
		return x.InputBotInlineResultPhoto.encode()
	case *TypeInputBotInlineResult_InputBotInlineResultDocument:
		return x.InputBotInlineResultDocument.encode()
	case *TypeInputBotInlineResult_InputBotInlineResultGame:
		return x.InputBotInlineResultGame.encode()
	}
	return nil
}
func (e *TypeBotInlineMessage) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeBotInlineMessage_BotInlineMessageMediaAuto:
		return x.BotInlineMessageMediaAuto.encode()
	case *TypeBotInlineMessage_BotInlineMessageText:
		return x.BotInlineMessageText.encode()
	case *TypeBotInlineMessage_BotInlineMessageMediaGeo:
		return x.BotInlineMessageMediaGeo.encode()
	case *TypeBotInlineMessage_BotInlineMessageMediaVenue:
		return x.BotInlineMessageMediaVenue.encode()
	case *TypeBotInlineMessage_BotInlineMessageMediaContact:
		return x.BotInlineMessageMediaContact.encode()
	}
	return nil
}
func (e *TypeBotInlineResult) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeBotInlineResult_BotInlineResult:
		return x.BotInlineResult.encode()
	case *TypeBotInlineResult_BotInlineMediaResult:
		return x.BotInlineMediaResult.encode()
	}
	return nil
}
func (e *TypeMessagesBotResults) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeExportedMessageLink) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessageFwdHeader) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePeerSettings) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeAuthCodeType) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeAuthCodeType_AuthCodeTypeSms:
		return x.AuthCodeTypeSms.encode()
	case *TypeAuthCodeType_AuthCodeTypeCall:
		return x.AuthCodeTypeCall.encode()
	case *TypeAuthCodeType_AuthCodeTypeFlashCall:
		return x.AuthCodeTypeFlashCall.encode()
	}
	return nil
}
func (e *TypeAuthSentCodeType) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeAuthSentCodeType_AuthSentCodeTypeApp:
		return x.AuthSentCodeTypeApp.encode()
	case *TypeAuthSentCodeType_AuthSentCodeTypeSms:
		return x.AuthSentCodeTypeSms.encode()
	case *TypeAuthSentCodeType_AuthSentCodeTypeCall:
		return x.AuthSentCodeTypeCall.encode()
	case *TypeAuthSentCodeType_AuthSentCodeTypeFlashCall:
		return x.AuthSentCodeTypeFlashCall.encode()
	}
	return nil
}
func (e *TypeMessagesBotCallbackAnswer) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesMessageEditData) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputBotInlineMessageID) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInlineBotSwitchPM) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesPeerDialogs) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeTopPeer) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeTopPeerCategory) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeTopPeerCategory_TopPeerCategoryBotsPM:
		return x.TopPeerCategoryBotsPM.encode()
	case *TypeTopPeerCategory_TopPeerCategoryBotsInline:
		return x.TopPeerCategoryBotsInline.encode()
	case *TypeTopPeerCategory_TopPeerCategoryCorrespondents:
		return x.TopPeerCategoryCorrespondents.encode()
	case *TypeTopPeerCategory_TopPeerCategoryGroups:
		return x.TopPeerCategoryGroups.encode()
	case *TypeTopPeerCategory_TopPeerCategoryChannels:
		return x.TopPeerCategoryChannels.encode()
	case *TypeTopPeerCategory_TopPeerCategoryPhoneCalls:
		return x.TopPeerCategoryPhoneCalls.encode()
	}
	return nil
}
func (e *TypeTopPeerCategoryPeers) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeContactsTopPeers) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeContactsTopPeers_ContactsTopPeersNotModified:
		return x.ContactsTopPeersNotModified.encode()
	case *TypeContactsTopPeers_ContactsTopPeers:
		return x.ContactsTopPeers.encode()
	}
	return nil
}
func (e *TypeDraftMessage) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeDraftMessage_DraftMessageEmpty:
		return x.DraftMessageEmpty.encode()
	case *TypeDraftMessage_DraftMessage:
		return x.DraftMessage.encode()
	}
	return nil
}
func (e *TypeMessagesFeaturedStickers) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesFeaturedStickers_MessagesFeaturedStickersNotModified:
		return x.MessagesFeaturedStickersNotModified.encode()
	case *TypeMessagesFeaturedStickers_MessagesFeaturedStickers:
		return x.MessagesFeaturedStickers.encode()
	}
	return nil
}
func (e *TypeMessagesRecentStickers) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesRecentStickers_MessagesRecentStickersNotModified:
		return x.MessagesRecentStickersNotModified.encode()
	case *TypeMessagesRecentStickers_MessagesRecentStickers:
		return x.MessagesRecentStickers.encode()
	}
	return nil
}
func (e *TypeMessagesArchivedStickers) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesStickerSetInstallResult) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesStickerSetInstallResult_MessagesStickerSetInstallResultSuccess:
		return x.MessagesStickerSetInstallResultSuccess.encode()
	case *TypeMessagesStickerSetInstallResult_MessagesStickerSetInstallResultArchive:
		return x.MessagesStickerSetInstallResultArchive.encode()
	}
	return nil
}
func (e *TypeStickerSetCovered) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeStickerSetCovered_StickerSetCovered:
		return x.StickerSetCovered.encode()
	case *TypeStickerSetCovered_StickerSetMultiCovered:
		return x.StickerSetMultiCovered.encode()
	}
	return nil
}
func (e *TypeMaskCoords) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputStickeredMedia) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputStickeredMedia_InputStickeredMediaPhoto:
		return x.InputStickeredMediaPhoto.encode()
	case *TypeInputStickeredMedia_InputStickeredMediaDocument:
		return x.InputStickeredMediaDocument.encode()
	}
	return nil
}
func (e *TypeGame) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputGame) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputGame_InputGameID:
		return x.InputGameID.encode()
	case *TypeInputGame_InputGameShortName:
		return x.InputGameShortName.encode()
	}
	return nil
}
func (e *TypeHighScore) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesHighScores) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeRichText) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeRichText_TextEmpty:
		return x.TextEmpty.encode()
	case *TypeRichText_TextPlain:
		return x.TextPlain.encode()
	case *TypeRichText_TextBold:
		return x.TextBold.encode()
	case *TypeRichText_TextItalic:
		return x.TextItalic.encode()
	case *TypeRichText_TextUnderline:
		return x.TextUnderline.encode()
	case *TypeRichText_TextStrike:
		return x.TextStrike.encode()
	case *TypeRichText_TextFixed:
		return x.TextFixed.encode()
	case *TypeRichText_TextUrl:
		return x.TextUrl.encode()
	case *TypeRichText_TextEmail:
		return x.TextEmail.encode()
	case *TypeRichText_TextConcat:
		return x.TextConcat.encode()
	}
	return nil
}
func (e *TypePageBlock) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePageBlock_PageBlockTitle:
		return x.PageBlockTitle.encode()
	case *TypePageBlock_PageBlockSubtitle:
		return x.PageBlockSubtitle.encode()
	case *TypePageBlock_PageBlockAuthorDate:
		return x.PageBlockAuthorDate.encode()
	case *TypePageBlock_PageBlockHeader:
		return x.PageBlockHeader.encode()
	case *TypePageBlock_PageBlockSubheader:
		return x.PageBlockSubheader.encode()
	case *TypePageBlock_PageBlockParagraph:
		return x.PageBlockParagraph.encode()
	case *TypePageBlock_PageBlockPreformatted:
		return x.PageBlockPreformatted.encode()
	case *TypePageBlock_PageBlockFooter:
		return x.PageBlockFooter.encode()
	case *TypePageBlock_PageBlockDivider:
		return x.PageBlockDivider.encode()
	case *TypePageBlock_PageBlockList:
		return x.PageBlockList.encode()
	case *TypePageBlock_PageBlockBlockquote:
		return x.PageBlockBlockquote.encode()
	case *TypePageBlock_PageBlockPullquote:
		return x.PageBlockPullquote.encode()
	case *TypePageBlock_PageBlockPhoto:
		return x.PageBlockPhoto.encode()
	case *TypePageBlock_PageBlockVideo:
		return x.PageBlockVideo.encode()
	case *TypePageBlock_PageBlockCover:
		return x.PageBlockCover.encode()
	case *TypePageBlock_PageBlockEmbed:
		return x.PageBlockEmbed.encode()
	case *TypePageBlock_PageBlockEmbedPost:
		return x.PageBlockEmbedPost.encode()
	case *TypePageBlock_PageBlockSlideshow:
		return x.PageBlockSlideshow.encode()
	case *TypePageBlock_PageBlockUnsupported:
		return x.PageBlockUnsupported.encode()
	case *TypePageBlock_PageBlockAnchor:
		return x.PageBlockAnchor.encode()
	case *TypePageBlock_PageBlockCollage:
		return x.PageBlockCollage.encode()
	case *TypePageBlock_PageBlockChannel:
		return x.PageBlockChannel.encode()
	case *TypePageBlock_PageBlockAudio:
		return x.PageBlockAudio.encode()
	}
	return nil
}
func (e *TypePage) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePage_PagePart:
		return x.PagePart.encode()
	case *TypePage_PageFull:
		return x.PageFull.encode()
	}
	return nil
}
func (e *TypeInputPhoneCall) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePhoneCall) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePhoneCall_PhoneCallEmpty:
		return x.PhoneCallEmpty.encode()
	case *TypePhoneCall_PhoneCallWaiting:
		return x.PhoneCallWaiting.encode()
	case *TypePhoneCall_PhoneCallRequested:
		return x.PhoneCallRequested.encode()
	case *TypePhoneCall_PhoneCall:
		return x.PhoneCall.encode()
	case *TypePhoneCall_PhoneCallDiscarded:
		return x.PhoneCallDiscarded.encode()
	case *TypePhoneCall_PhoneCallAccepted:
		return x.PhoneCallAccepted.encode()
	}
	return nil
}
func (e *TypePhoneConnection) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePhoneCallProtocol) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePhonePhoneCall) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePhoneCallDiscardReason) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePhoneCallDiscardReason_PhoneCallDiscardReasonMissed:
		return x.PhoneCallDiscardReasonMissed.encode()
	case *TypePhoneCallDiscardReason_PhoneCallDiscardReasonDisconnect:
		return x.PhoneCallDiscardReasonDisconnect.encode()
	case *TypePhoneCallDiscardReason_PhoneCallDiscardReasonHangup:
		return x.PhoneCallDiscardReasonHangup.encode()
	case *TypePhoneCallDiscardReason_PhoneCallDiscardReasonBusy:
		return x.PhoneCallDiscardReasonBusy.encode()
	}
	return nil
}
func (e *TypeInvoice) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePaymentsPaymentForm) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePostAddress) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePaymentRequestedInfo) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeDataJSON) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeLabeledPrice) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePaymentCharge) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePaymentSavedCredentials) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeWebDocument) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputWebDocument) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputWebFileLocation) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeUploadWebFile) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePaymentsValidatedRequestedInfo) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePaymentsPaymentResult) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypePaymentsPaymentResult_PaymentsPaymentResult:
		return x.PaymentsPaymentResult.encode()
	case *TypePaymentsPaymentResult_PaymentsPaymentVerficationNeeded:
		return x.PaymentsPaymentVerficationNeeded.encode()
	}
	return nil
}
func (e *TypePaymentsPaymentReceipt) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePaymentsSavedInfo) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputPaymentCredentials) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeInputPaymentCredentials_InputPaymentCredentialsSaved:
		return x.InputPaymentCredentialsSaved.encode()
	case *TypeInputPaymentCredentials_InputPaymentCredentials:
		return x.InputPaymentCredentials.encode()
	}
	return nil
}
func (e *TypeAccountTmpPassword) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeShippingOption) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeUploadCdnFile) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeUploadCdnFile_UploadCdnFileReuploadNeeded:
		return x.UploadCdnFileReuploadNeeded.encode()
	case *TypeUploadCdnFile_UploadCdnFile:
		return x.UploadCdnFile.encode()
	}
	return nil
}
func (e *TypeCdnPublicKey) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeCdnConfig) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeInputStickerSetItem) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeLangPackString) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeLangPackString_LangPackString:
		return x.LangPackString.encode()
	case *TypeLangPackString_LangPackStringPluralized:
		return x.LangPackStringPluralized.encode()
	case *TypeLangPackString_LangPackStringDeleted:
		return x.LangPackStringDeleted.encode()
	}
	return nil
}
func (e *TypeLangPackDifference) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeLangPackLanguage) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeChannelAdminRights) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeChannelBannedRights) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeChannelAdminLogEventAction) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeTitle:
		return x.ChannelAdminLogEventActionChangeTitle.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeAbout:
		return x.ChannelAdminLogEventActionChangeAbout.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeUsername:
		return x.ChannelAdminLogEventActionChangeUsername.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangePhoto:
		return x.ChannelAdminLogEventActionChangePhoto.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionToggleInvites:
		return x.ChannelAdminLogEventActionToggleInvites.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionToggleSignatures:
		return x.ChannelAdminLogEventActionToggleSignatures.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionUpdatePinned:
		return x.ChannelAdminLogEventActionUpdatePinned.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionEditMessage:
		return x.ChannelAdminLogEventActionEditMessage.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionDeleteMessage:
		return x.ChannelAdminLogEventActionDeleteMessage.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantJoin:
		return x.ChannelAdminLogEventActionParticipantJoin.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantLeave:
		return x.ChannelAdminLogEventActionParticipantLeave.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantInvite:
		return x.ChannelAdminLogEventActionParticipantInvite.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleBan:
		return x.ChannelAdminLogEventActionParticipantToggleBan.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleAdmin:
		return x.ChannelAdminLogEventActionParticipantToggleAdmin.encode()
	case *TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeStickerSet:
		return x.ChannelAdminLogEventActionChangeStickerSet.encode()
	}
	return nil
}
func (e *TypeChannelAdminLogEvent) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeChannelsAdminLogResults) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeChannelAdminLogEventsFilter) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypePopularContact) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeCdnFileHash) encode() []byte {
	return e.GetValue().encode()
}
func (e *TypeMessagesFavedStickers) encode() []byte {
	switch x := e.GetValue().(type) {
	case *TypeMessagesFavedStickers_MessagesFavedStickers:
		return x.MessagesFavedStickers.encode()
	case *TypeMessagesFavedStickers_MessagesFavedStickersNotModified:
		return x.MessagesFavedStickersNotModified.encode()
	}
	return nil
}

// Encode funcs for predicates
func (e *PredBoolFalse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolFalse)
	return x.buf
}
func (e *PredBoolTrue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolTrue)
	return x.buf
}
func (e *PredError) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_error)
	x.Int(e.Code)
	x.String(e.Text)
	return x.buf
}
func (e *PredNull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_null)
	return x.buf
}
func (e *PredInputPeerEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerEmpty)
	return x.buf
}
func (e *PredInputPeerSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerSelf)
	return x.buf
}
func (e *PredInputPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChat)
	x.Int(e.ChatId)
	return x.buf
}
func (e *PredInputUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserEmpty)
	return x.buf
}
func (e *PredInputUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserSelf)
	return x.buf
}
func (e *PredInputPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneContact)
	x.Long(e.ClientId)
	x.String(e.Phone)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}
func (e *PredInputFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFile)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	x.String(e.Md5Checksum)
	return x.buf
}
func (e *PredInputMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaEmpty)
	return x.buf
}
func (e *PredInputMediaUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedPhoto)
	x.Int(e.Flags)
	x.Bytes(e.File.encode())
	x.String(e.Caption)
	x.FlaggedVector(e.Flags, 0, toTLslice(e.Stickers))
	x.FlaggedInt(e.Flags, 1, e.TtlSeconds)
	return x.buf
}
func (e *PredInputMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhoto)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.String(e.Caption)
	x.FlaggedInt(e.Flags, 0, e.TtlSeconds)
	return x.buf
}
func (e *PredInputMediaGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGeoPoint)
	x.Bytes(e.GeoPoint.encode())
	return x.buf
}
func (e *PredInputMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaContact)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}
func (e *PredInputChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhotoEmpty)
	return x.buf
}
func (e *PredInputChatUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatUploadedPhoto)
	x.Bytes(e.File.encode())
	return x.buf
}
func (e *PredInputChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *PredInputGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPointEmpty)
	return x.buf
}
func (e *PredInputGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPoint)
	x.Double(e.Lat)
	x.Double(e.Long)
	return x.buf
}
func (e *PredInputPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoEmpty)
	return x.buf
}
func (e *PredInputPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoto)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredInputFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileLocation)
	x.Long(e.VolumeId)
	x.Int(e.LocalId)
	x.Long(e.Secret)
	return x.buf
}
func (e *PredInputAppEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAppEvent)
	x.Double(e.Time)
	x.String(e.Type)
	x.Long(e.Peer)
	x.String(e.Data)
	return x.buf
}
func (e *PredPeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerUser)
	x.Int(e.UserId)
	return x.buf
}
func (e *PredPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChat)
	x.Int(e.ChatId)
	return x.buf
}
func (e *PredStorageFileUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFileUnknown)
	return x.buf
}
func (e *PredStorageFileJpeg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFileJpeg)
	return x.buf
}
func (e *PredStorageFileGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFileGif)
	return x.buf
}
func (e *PredStorageFilePng) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFilePng)
	return x.buf
}
func (e *PredStorageFileMp3) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFileMp3)
	return x.buf
}
func (e *PredStorageFileMov) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFileMov)
	return x.buf
}
func (e *PredStorageFilePartial) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFilePartial)
	return x.buf
}
func (e *PredStorageFileMp4) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFileMp4)
	return x.buf
}
func (e *PredStorageFileWebp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFileWebp)
	return x.buf
}
func (e *PredFileLocationUnavailable) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocationUnavailable)
	x.Long(e.VolumeId)
	x.Int(e.LocalId)
	x.Long(e.Secret)
	return x.buf
}
func (e *PredFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocation)
	x.Int(e.DcId)
	x.Long(e.VolumeId)
	x.Int(e.LocalId)
	x.Long(e.Secret)
	return x.buf
}
func (e *PredUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userEmpty)
	x.Int(e.Id)
	return x.buf
}
func (e *PredUserProfilePhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhotoEmpty)
	return x.buf
}
func (e *PredUserProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhoto)
	x.Long(e.PhotoId)
	x.Bytes(e.PhotoSmall.encode())
	x.Bytes(e.PhotoBig.encode())
	return x.buf
}
func (e *PredUserStatusEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusEmpty)
	return x.buf
}
func (e *PredUserStatusOnline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOnline)
	x.Int(e.Expires)
	return x.buf
}
func (e *PredUserStatusOffline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOffline)
	x.Int(e.WasOnline)
	return x.buf
}
func (e *PredChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatEmpty)
	x.Int(e.Id)
	return x.buf
}
func (e *PredChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chat)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.ParticipantsCount)
	x.Int(e.Date)
	x.Int(e.Version)
	x.FlaggedObject(e.Flags, 6, e.MigratedTo)
	return x.buf
}
func (e *PredChatForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatForbidden)
	x.Int(e.Id)
	x.String(e.Title)
	return x.buf
}
func (e *PredChatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatFull)
	x.Int(e.Id)
	x.Bytes(e.Participants.encode())
	x.Bytes(e.ChatPhoto.encode())
	x.Bytes(e.NotifySettings.encode())
	x.Bytes(e.ExportedInvite.encode())
	x.Vector(toTLslice(e.BotInfos))
	return x.buf
}
func (e *PredChatParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipant)
	x.Int(e.UserId)
	x.Int(e.InviterId)
	x.Int(e.Date)
	return x.buf
}
func (e *PredChatParticipantsForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantsForbidden)
	x.Int(e.Flags)
	x.Int(e.ChatId)
	x.FlaggedObject(e.Flags, 0, e.SelfParticipant)
	return x.buf
}
func (e *PredChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipants)
	x.Int(e.ChatId)
	x.Vector(toTLslice(e.Participants))
	x.Int(e.Version)
	return x.buf
}
func (e *PredChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhotoEmpty)
	return x.buf
}
func (e *PredChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhoto)
	x.Bytes(e.PhotoSmall.encode())
	x.Bytes(e.PhotoBig.encode())
	return x.buf
}
func (e *PredMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEmpty)
	x.Int(e.Id)
	return x.buf
}
func (e *PredMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_message)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.FlaggedInt(e.Flags, 8, e.FromId)
	x.Bytes(e.ToId.encode())
	x.FlaggedObject(e.Flags, 2, e.FwdFrom)
	x.FlaggedInt(e.Flags, 11, e.ViaBotId)
	x.FlaggedInt(e.Flags, 3, e.ReplyToMsgId)
	x.Int(e.Date)
	x.String(e.Message)
	x.FlaggedObject(e.Flags, 9, e.Media)
	x.FlaggedObject(e.Flags, 6, e.ReplyMarkup)
	x.FlaggedVector(e.Flags, 7, toTLslice(e.Entities))
	x.FlaggedInt(e.Flags, 10, e.Views)
	x.FlaggedInt(e.Flags, 15, e.EditDate)
	x.FlaggedString(e.Flags, 16, e.PostAuthor)
	return x.buf
}
func (e *PredMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageService)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.FlaggedInt(e.Flags, 8, e.FromId)
	x.Bytes(e.ToId.encode())
	x.FlaggedInt(e.Flags, 3, e.ReplyToMsgId)
	x.Int(e.Date)
	x.Bytes(e.Action.encode())
	return x.buf
}
func (e *PredMessageMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaEmpty)
	return x.buf
}
func (e *PredMessageMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaPhoto)
	x.Int(e.Flags)
	x.FlaggedObject(e.Flags, 0, e.Photo)
	x.FlaggedString(e.Flags, 1, e.Caption)
	x.FlaggedInt(e.Flags, 2, e.TtlSeconds)
	return x.buf
}
func (e *PredMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGeo)
	x.Bytes(e.Geo.encode())
	return x.buf
}
func (e *PredMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaContact)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.Int(e.UserId)
	return x.buf
}
func (e *PredMessageMediaUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaUnsupported)
	return x.buf
}
func (e *PredMessageActionEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionEmpty)
	return x.buf
}
func (e *PredMessageActionChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatCreate)
	x.String(e.Title)
	x.VectorInt(e.Users)
	return x.buf
}
func (e *PredMessageActionChatEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditTitle)
	x.String(e.Title)
	return x.buf
}
func (e *PredMessageActionChatEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditPhoto)
	x.Bytes(e.Photo.encode())
	return x.buf
}
func (e *PredMessageActionChatDeletePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeletePhoto)
	return x.buf
}
func (e *PredMessageActionChatAddUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatAddUser)
	x.VectorInt(e.Users)
	return x.buf
}
func (e *PredMessageActionChatDeleteUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeleteUser)
	x.Int(e.UserId)
	return x.buf
}
func (e *PredDialog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialog)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.TopMessage)
	x.Int(e.ReadInboxMaxId)
	x.Int(e.ReadOutboxMaxId)
	x.Int(e.UnreadCount)
	x.Int(e.UnreadMentionsCount)
	x.Bytes(e.NotifySettings.encode())
	x.FlaggedInt(e.Flags, 0, e.Pts)
	x.FlaggedObject(e.Flags, 1, e.Draft)
	return x.buf
}
func (e *PredPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoEmpty)
	x.Long(e.Id)
	return x.buf
}
func (e *PredPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photo)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Vector(toTLslice(e.Sizes))
	return x.buf
}
func (e *PredPhotoSizeEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSizeEmpty)
	x.String(e.Type)
	return x.buf
}
func (e *PredPhotoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSize)
	x.String(e.Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	return x.buf
}
func (e *PredPhotoCachedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoCachedSize)
	x.String(e.Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *PredGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPointEmpty)
	return x.buf
}
func (e *PredGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPoint)
	x.Double(e.Long)
	x.Double(e.Lat)
	return x.buf
}
func (e *PredAuthCheckedPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authCheckedPhone)
	x.Bytes(e.PhoneRegistered.encode())
	return x.buf
}
func (e *PredAuthSentCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSentCode)
	x.Int(e.Flags)
	x.Bytes(e.Type.encode())
	x.String(e.PhoneCodeHash)
	x.FlaggedObject(e.Flags, 1, e.NextType)
	x.FlaggedInt(e.Flags, 2, e.Timeout)
	return x.buf
}
func (e *PredAuthAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authAuthorization)
	x.Int(e.Flags)
	x.FlaggedInt(e.Flags, 0, e.TmpSessions)
	x.Bytes(e.User.encode())
	return x.buf
}
func (e *PredAuthExportedAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authExportedAuthorization)
	x.Int(e.Id)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *PredInputNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *PredInputNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyUsers)
	return x.buf
}
func (e *PredInputNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyChats)
	return x.buf
}
func (e *PredInputNotifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyAll)
	return x.buf
}
func (e *PredInputPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifySettings)
	x.Int(e.Flags)
	x.Int(e.MuteUntil)
	x.String(e.Sound)
	return x.buf
}
func (e *PredPeerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsEmpty)
	return x.buf
}
func (e *PredPeerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsAll)
	return x.buf
}
func (e *PredPeerNotifySettingsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettingsEmpty)
	return x.buf
}
func (e *PredPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettings)
	x.Int(e.Flags)
	x.Int(e.MuteUntil)
	x.String(e.Sound)
	return x.buf
}
func (e *PredWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaper)
	x.Int(e.Id)
	x.String(e.Title)
	x.Vector(toTLslice(e.Sizes))
	x.Int(e.Color)
	return x.buf
}
func (e *PredUserFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userFull)
	x.Int(e.Flags)
	x.Bytes(e.User.encode())
	x.FlaggedString(e.Flags, 1, e.About)
	x.Bytes(e.Link.encode())
	x.FlaggedObject(e.Flags, 2, e.ProfilePhoto)
	x.Bytes(e.NotifySettings.encode())
	x.FlaggedObject(e.Flags, 3, e.BotInfo)
	x.Int(e.CommonChatsCount)
	return x.buf
}
func (e *PredContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contact)
	x.Int(e.UserId)
	x.Bytes(e.Mutual.encode())
	return x.buf
}
func (e *PredImportedContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_importedContact)
	x.Int(e.UserId)
	x.Long(e.ClientId)
	return x.buf
}
func (e *PredContactBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactBlocked)
	x.Int(e.UserId)
	x.Int(e.Date)
	return x.buf
}
func (e *PredContactStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactStatus)
	x.Int(e.UserId)
	x.Bytes(e.Status.encode())
	return x.buf
}
func (e *PredContactsLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsLink)
	x.Bytes(e.MyLink.encode())
	x.Bytes(e.ForeignLink.encode())
	x.Bytes(e.User.encode())
	return x.buf
}
func (e *PredContactsContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsContacts)
	x.Vector(toTLslice(e.Contacts))
	x.Int(e.SavedCount)
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredContactsContactsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsContactsNotModified)
	return x.buf
}
func (e *PredContactsImportedContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsImportedContacts)
	x.Vector(toTLslice(e.Imported))
	x.Vector(toTLslice(e.PopularInvites))
	x.VectorLong(e.RetryContacts)
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredContactsBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsBlocked)
	x.Vector(toTLslice(e.Blocked))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredContactsBlockedSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsBlockedSlice)
	x.Int(e.Count)
	x.Vector(toTLslice(e.Blocked))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredContactsFound) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsFound)
	x.Vector(toTLslice(e.Results))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredMessagesDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesDialogs)
	x.Vector(toTLslice(e.Dialogs))
	x.Vector(toTLslice(e.Messages))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredMessagesDialogsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesDialogsSlice)
	x.Int(e.Count)
	x.Vector(toTLslice(e.Dialogs))
	x.Vector(toTLslice(e.Messages))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredMessagesMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesMessages)
	x.Vector(toTLslice(e.Messages))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredMessagesMessagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesMessagesSlice)
	x.Int(e.Count)
	x.Vector(toTLslice(e.Messages))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredMessagesChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesChats)
	x.Vector(toTLslice(e.Chats))
	return x.buf
}
func (e *PredMessagesChatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesChatFull)
	x.Bytes(e.FullChat.encode())
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredMessagesAffectedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesAffectedHistory)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Offset)
	return x.buf
}
func (e *PredInputMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterEmpty)
	return x.buf
}
func (e *PredInputMessagesFilterPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotos)
	return x.buf
}
func (e *PredInputMessagesFilterVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVideo)
	return x.buf
}
func (e *PredInputMessagesFilterPhotoVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideo)
	return x.buf
}
func (e *PredUpdateNewMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredUpdateMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateMessageID)
	x.Int(e.Id)
	x.Long(e.RandomId)
	return x.buf
}
func (e *PredUpdateDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteMessages)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredUpdateUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserTyping)
	x.Int(e.UserId)
	x.Bytes(e.Action.encode())
	return x.buf
}
func (e *PredUpdateChatUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatUserTyping)
	x.Int(e.ChatId)
	x.Int(e.UserId)
	x.Bytes(e.Action.encode())
	return x.buf
}
func (e *PredUpdateChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipants)
	x.Bytes(e.Participants.encode())
	return x.buf
}
func (e *PredUpdateUserStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserStatus)
	x.Int(e.UserId)
	x.Bytes(e.Status.encode())
	return x.buf
}
func (e *PredUpdateUserName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserName)
	x.Int(e.UserId)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	return x.buf
}
func (e *PredUpdateUserPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhoto)
	x.Int(e.UserId)
	x.Int(e.Date)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Previous.encode())
	return x.buf
}
func (e *PredUpdateContactRegistered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactRegistered)
	x.Int(e.UserId)
	x.Int(e.Date)
	return x.buf
}
func (e *PredUpdateContactLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactLink)
	x.Int(e.UserId)
	x.Bytes(e.MyLink.encode())
	x.Bytes(e.ForeignLink.encode())
	return x.buf
}
func (e *PredUpdatesState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesState)
	x.Int(e.Pts)
	x.Int(e.Qts)
	x.Int(e.Date)
	x.Int(e.Seq)
	x.Int(e.UnreadCount)
	return x.buf
}
func (e *PredUpdatesDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesDifferenceEmpty)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}
func (e *PredUpdatesDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesDifference)
	x.Vector(toTLslice(e.NewMessages))
	x.Vector(toTLslice(e.NewEncryptedMessages))
	x.Vector(toTLslice(e.OtherUpdates))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	x.Bytes(e.State.encode())
	return x.buf
}
func (e *PredUpdatesDifferenceSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesDifferenceSlice)
	x.Vector(toTLslice(e.NewMessages))
	x.Vector(toTLslice(e.NewEncryptedMessages))
	x.Vector(toTLslice(e.OtherUpdates))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	x.Bytes(e.IntermediateState.encode())
	return x.buf
}
func (e *PredUpdatesTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesTooLong)
	return x.buf
}
func (e *PredUpdateShortMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.UserId)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	x.FlaggedObject(e.Flags, 2, e.FwdFrom)
	x.FlaggedInt(e.Flags, 11, e.ViaBotId)
	x.FlaggedInt(e.Flags, 3, e.ReplyToMsgId)
	x.FlaggedVector(e.Flags, 7, toTLslice(e.Entities))
	return x.buf
}
func (e *PredUpdateShortChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortChatMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.FromId)
	x.Int(e.ChatId)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	x.FlaggedObject(e.Flags, 2, e.FwdFrom)
	x.FlaggedInt(e.Flags, 11, e.ViaBotId)
	x.FlaggedInt(e.Flags, 3, e.ReplyToMsgId)
	x.FlaggedVector(e.Flags, 7, toTLslice(e.Entities))
	return x.buf
}
func (e *PredUpdateShort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShort)
	x.Bytes(e.Update.encode())
	x.Int(e.Date)
	return x.buf
}
func (e *PredUpdatesCombined) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesCombined)
	x.Vector(toTLslice(e.Updates))
	x.Vector(toTLslice(e.Users))
	x.Vector(toTLslice(e.Chats))
	x.Int(e.Date)
	x.Int(e.SeqStart)
	x.Int(e.Seq)
	return x.buf
}
func (e *PredUpdates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates)
	x.Vector(toTLslice(e.Updates))
	x.Vector(toTLslice(e.Users))
	x.Vector(toTLslice(e.Chats))
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}
func (e *PredPhotosPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photosPhoto)
	x.Bytes(e.Photo.encode())
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredUploadFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadFile)
	x.Bytes(e.Type.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *PredDcOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dcOption)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.IpAddress)
	x.Int(e.Port)
	return x.buf
}
func (e *PredConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_config)
	x.Int(e.Flags)
	x.Int(e.Date)
	x.Int(e.Expires)
	x.Bytes(e.TestMode.encode())
	x.Int(e.ThisDc)
	x.Vector(toTLslice(e.DcOptions))
	x.Int(e.ChatSizeMax)
	x.Int(e.MegagroupSizeMax)
	x.Int(e.ForwardedCountMax)
	x.Int(e.OnlineUpdatePeriodMs)
	x.Int(e.OfflineBlurTimeoutMs)
	x.Int(e.OfflineIdleTimeoutMs)
	x.Int(e.OnlineCloudTimeoutMs)
	x.Int(e.NotifyCloudDelayMs)
	x.Int(e.NotifyDefaultDelayMs)
	x.Int(e.ChatBigSize)
	x.Int(e.PushChatPeriodMs)
	x.Int(e.PushChatLimit)
	x.Int(e.SavedGifsLimit)
	x.Int(e.EditTimeLimit)
	x.Int(e.RatingEDecay)
	x.Int(e.StickersRecentLimit)
	x.Int(e.StickersFavedLimit)
	x.FlaggedInt(e.Flags, 0, e.TmpSessions)
	x.Int(e.PinnedDialogsCountMax)
	x.Int(e.CallReceiveTimeoutMs)
	x.Int(e.CallRingTimeoutMs)
	x.Int(e.CallConnectTimeoutMs)
	x.Int(e.CallPacketTimeoutMs)
	x.String(e.MeUrlPrefix)
	x.FlaggedString(e.Flags, 2, e.SuggestedLangCode)
	x.FlaggedInt(e.Flags, 2, e.LangPackVersion)
	x.Vector(toTLslice(e.DisabledFeatures))
	return x.buf
}
func (e *PredNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_nearestDc)
	x.String(e.Country)
	x.Int(e.ThisDc)
	x.Int(e.NearestDc)
	return x.buf
}
func (e *PredHelpAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpAppUpdate)
	x.Int(e.Id)
	x.Bytes(e.Critical.encode())
	x.String(e.Url)
	x.String(e.Text)
	return x.buf
}
func (e *PredHelpNoAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpNoAppUpdate)
	return x.buf
}
func (e *PredHelpInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpInviteText)
	x.String(e.Message)
	return x.buf
}
func (e *PredInputPeerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsEmpty)
	return x.buf
}
func (e *PredInputPeerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsAll)
	return x.buf
}
func (e *PredPhotosPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photosPhotos)
	x.Vector(toTLslice(e.Photos))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredPhotosPhotosSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photosPhotosSlice)
	x.Int(e.Count)
	x.Vector(toTLslice(e.Photos))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredWallPaperSolid) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaperSolid)
	x.Int(e.Id)
	x.String(e.Title)
	x.Int(e.BgColor)
	x.Int(e.Color)
	return x.buf
}
func (e *PredUpdateNewEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewEncryptedMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Qts)
	return x.buf
}
func (e *PredUpdateEncryptedChatTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedChatTyping)
	x.Int(e.ChatId)
	return x.buf
}
func (e *PredUpdateEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryption)
	x.Bytes(e.Chat.encode())
	x.Int(e.Date)
	return x.buf
}
func (e *PredUpdateEncryptedMessagesRead) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedMessagesRead)
	x.Int(e.ChatId)
	x.Int(e.MaxDate)
	x.Int(e.Date)
	return x.buf
}
func (e *PredEncryptedChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatEmpty)
	x.Int(e.Id)
	return x.buf
}
func (e *PredEncryptedChatWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatWaiting)
	x.Int(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	return x.buf
}
func (e *PredEncryptedChatRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatRequested)
	x.Int(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	x.StringBytes(e.GA)
	return x.buf
}
func (e *PredEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChat)
	x.Int(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	x.StringBytes(e.GAOrB)
	x.Long(e.KeyFingerprint)
	return x.buf
}
func (e *PredEncryptedChatDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatDiscarded)
	x.Int(e.Id)
	return x.buf
}
func (e *PredInputEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedChat)
	x.Int(e.ChatId)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFileEmpty)
	return x.buf
}
func (e *PredEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFile)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.Int(e.DcId)
	x.Int(e.KeyFingerprint)
	return x.buf
}
func (e *PredInputEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileEmpty)
	return x.buf
}
func (e *PredInputEncryptedFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Md5Checksum)
	x.Int(e.KeyFingerprint)
	return x.buf
}
func (e *PredInputEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFile)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredInputEncryptedFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileLocation)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessage)
	x.Long(e.RandomId)
	x.Int(e.ChatId)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	x.Bytes(e.File.encode())
	return x.buf
}
func (e *PredEncryptedMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessageService)
	x.Long(e.RandomId)
	x.Int(e.ChatId)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *PredMessagesDhConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesDhConfigNotModified)
	x.StringBytes(e.Random)
	return x.buf
}
func (e *PredMessagesDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesDhConfig)
	x.Int(e.G)
	x.StringBytes(e.P)
	x.Int(e.Version)
	x.StringBytes(e.Random)
	return x.buf
}
func (e *PredMessagesSentEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSentEncryptedMessage)
	x.Int(e.Date)
	return x.buf
}
func (e *PredMessagesSentEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSentEncryptedFile)
	x.Int(e.Date)
	x.Bytes(e.File.encode())
	return x.buf
}
func (e *PredInputFileBig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileBig)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	return x.buf
}
func (e *PredInputEncryptedFileBigUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileBigUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.Int(e.KeyFingerprint)
	return x.buf
}
func (e *PredStorageFilePdf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storageFilePdf)
	return x.buf
}
func (e *PredInputMessagesFilterDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterDocument)
	return x.buf
}
func (e *PredInputMessagesFilterPhotoVideoDocuments) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideoDocuments)
	return x.buf
}
func (e *PredUpdateChatParticipantAdd) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdd)
	x.Int(e.ChatId)
	x.Int(e.UserId)
	x.Int(e.InviterId)
	x.Int(e.Date)
	x.Int(e.Version)
	return x.buf
}
func (e *PredUpdateChatParticipantDelete) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantDelete)
	x.Int(e.ChatId)
	x.Int(e.UserId)
	x.Int(e.Version)
	return x.buf
}
func (e *PredUpdateDcOptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDcOptions)
	x.Vector(toTLslice(e.DcOptions))
	return x.buf
}
func (e *PredInputMediaUploadedDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedDocument)
	x.Int(e.Flags)
	x.Bytes(e.File.encode())
	x.FlaggedObject(e.Flags, 2, e.Thumb)
	x.String(e.MimeType)
	x.Vector(toTLslice(e.Attributes))
	x.String(e.Caption)
	x.FlaggedVector(e.Flags, 0, toTLslice(e.Stickers))
	x.FlaggedInt(e.Flags, 1, e.TtlSeconds)
	return x.buf
}
func (e *PredInputMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocument)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.String(e.Caption)
	x.FlaggedInt(e.Flags, 0, e.TtlSeconds)
	return x.buf
}
func (e *PredMessageMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaDocument)
	x.Int(e.Flags)
	x.FlaggedObject(e.Flags, 0, e.Document)
	x.FlaggedString(e.Flags, 1, e.Caption)
	x.FlaggedInt(e.Flags, 2, e.TtlSeconds)
	return x.buf
}
func (e *PredInputDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentEmpty)
	return x.buf
}
func (e *PredInputDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocument)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredInputDocumentFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentFileLocation)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Version)
	return x.buf
}
func (e *PredDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentEmpty)
	x.Long(e.Id)
	return x.buf
}
func (e *PredDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_document)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.String(e.MimeType)
	x.Int(e.Size)
	x.Bytes(e.Thumb.encode())
	x.Int(e.DcId)
	x.Int(e.Version)
	x.Vector(toTLslice(e.Attributes))
	return x.buf
}
func (e *PredHelpSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpSupport)
	x.String(e.PhoneNumber)
	x.Bytes(e.User.encode())
	return x.buf
}
func (e *PredNotifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyAll)
	return x.buf
}
func (e *PredNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyChats)
	return x.buf
}
func (e *PredNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *PredNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyUsers)
	return x.buf
}
func (e *PredUpdateUserBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserBlocked)
	x.Int(e.UserId)
	x.Bytes(e.Blocked.encode())
	return x.buf
}
func (e *PredUpdateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.NotifySettings.encode())
	return x.buf
}
func (e *PredSendMessageTypingAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageTypingAction)
	return x.buf
}
func (e *PredSendMessageCancelAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageCancelAction)
	return x.buf
}
func (e *PredSendMessageRecordVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordVideoAction)
	return x.buf
}
func (e *PredSendMessageUploadVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadVideoAction)
	x.Int(e.Progress)
	return x.buf
}
func (e *PredSendMessageRecordAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordAudioAction)
	return x.buf
}
func (e *PredSendMessageUploadAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadAudioAction)
	x.Int(e.Progress)
	return x.buf
}
func (e *PredSendMessageUploadPhotoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadPhotoAction)
	x.Int(e.Progress)
	return x.buf
}
func (e *PredSendMessageUploadDocumentAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadDocumentAction)
	x.Int(e.Progress)
	return x.buf
}
func (e *PredSendMessageGeoLocationAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGeoLocationAction)
	return x.buf
}
func (e *PredSendMessageChooseContactAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageChooseContactAction)
	return x.buf
}
func (e *PredUpdateServiceNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateServiceNotification)
	x.Int(e.Flags)
	x.FlaggedInt(e.Flags, 1, e.InboxDate)
	x.String(e.Type)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	x.Vector(toTLslice(e.Entities))
	return x.buf
}
func (e *PredUserStatusRecently) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusRecently)
	return x.buf
}
func (e *PredUserStatusLastWeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastWeek)
	return x.buf
}
func (e *PredUserStatusLastMonth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastMonth)
	return x.buf
}
func (e *PredUpdatePrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(toTLslice(e.Rules))
	return x.buf
}
func (e *PredInputPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyStatusTimestamp)
	return x.buf
}
func (e *PredPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyStatusTimestamp)
	return x.buf
}
func (e *PredInputPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowContacts)
	return x.buf
}
func (e *PredInputPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowAll)
	return x.buf
}
func (e *PredInputPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowUsers)
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredInputPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowContacts)
	return x.buf
}
func (e *PredInputPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowAll)
	return x.buf
}
func (e *PredInputPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowUsers)
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowContacts)
	return x.buf
}
func (e *PredPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowAll)
	return x.buf
}
func (e *PredPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowUsers)
	x.VectorInt(e.Users)
	return x.buf
}
func (e *PredPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowContacts)
	return x.buf
}
func (e *PredPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowAll)
	return x.buf
}
func (e *PredPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowUsers)
	x.VectorInt(e.Users)
	return x.buf
}
func (e *PredAccountPrivacyRules) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountPrivacyRules)
	x.Vector(toTLslice(e.Rules))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredAccountDaysTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountDaysTTL)
	x.Int(e.Days)
	return x.buf
}
func (e *PredUpdateUserPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhone)
	x.Int(e.UserId)
	x.String(e.Phone)
	return x.buf
}
func (e *PredDisabledFeature) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_disabledFeature)
	x.String(e.Feature)
	x.String(e.Description)
	return x.buf
}
func (e *PredDocumentAttributeImageSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeImageSize)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}
func (e *PredDocumentAttributeAnimated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAnimated)
	return x.buf
}
func (e *PredDocumentAttributeSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeSticker)
	x.Int(e.Flags)
	x.String(e.Alt)
	x.Bytes(e.Stickerset.encode())
	x.FlaggedObject(e.Flags, 0, e.MaskCoords)
	return x.buf
}
func (e *PredDocumentAttributeVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeVideo)
	x.Int(e.Flags)
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}
func (e *PredDocumentAttributeAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAudio)
	x.Int(e.Flags)
	x.Int(e.Duration)
	x.FlaggedString(e.Flags, 0, e.Title)
	x.FlaggedString(e.Flags, 1, e.Performer)
	x.FlaggedStringBytes(e.Flags, 2, e.Waveform)
	return x.buf
}
func (e *PredDocumentAttributeFilename) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeFilename)
	x.String(e.FileName)
	return x.buf
}
func (e *PredMessagesStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesStickersNotModified)
	return x.buf
}
func (e *PredMessagesStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesStickers)
	x.String(e.Hash)
	x.Vector(toTLslice(e.Stickers))
	return x.buf
}
func (e *PredStickerPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerPack)
	x.String(e.Emoticon)
	x.VectorLong(e.Documents)
	return x.buf
}
func (e *PredMessagesAllStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesAllStickersNotModified)
	return x.buf
}
func (e *PredMessagesAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesAllStickers)
	x.Int(e.Hash)
	x.Vector(toTLslice(e.Sets))
	return x.buf
}
func (e *PredAccountNoPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountNoPassword)
	x.StringBytes(e.NewSalt)
	x.String(e.EmailUnconfirmedPattern)
	return x.buf
}
func (e *PredAccountPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountPassword)
	x.StringBytes(e.CurrentSalt)
	x.StringBytes(e.NewSalt)
	x.String(e.Hint)
	x.Bytes(e.HasRecovery.encode())
	x.String(e.EmailUnconfirmedPattern)
	return x.buf
}
func (e *PredUpdateReadHistoryInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryInbox)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxId)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredUpdateReadHistoryOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryOutbox)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxId)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredMessagesAffectedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesAffectedMessages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredContactLinkUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkUnknown)
	return x.buf
}
func (e *PredContactLinkNone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkNone)
	return x.buf
}
func (e *PredContactLinkHasPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkHasPhone)
	return x.buf
}
func (e *PredContactLinkContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkContact)
	return x.buf
}
func (e *PredUpdateWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateWebPage)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredWebPageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageEmpty)
	x.Long(e.Id)
	return x.buf
}
func (e *PredWebPagePending) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPagePending)
	x.Long(e.Id)
	x.Int(e.Date)
	return x.buf
}
func (e *PredWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPage)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.String(e.Url)
	x.String(e.DisplayUrl)
	x.Int(e.Hash)
	x.FlaggedString(e.Flags, 0, e.Type)
	x.FlaggedString(e.Flags, 1, e.SiteName)
	x.FlaggedString(e.Flags, 2, e.Title)
	x.FlaggedString(e.Flags, 3, e.Description)
	x.FlaggedObject(e.Flags, 4, e.Photo)
	x.FlaggedString(e.Flags, 5, e.EmbedUrl)
	x.FlaggedString(e.Flags, 5, e.EmbedType)
	x.FlaggedInt(e.Flags, 6, e.EmbedWidth)
	x.FlaggedInt(e.Flags, 6, e.EmbedHeight)
	x.FlaggedInt(e.Flags, 7, e.Duration)
	x.FlaggedString(e.Flags, 8, e.Author)
	x.FlaggedObject(e.Flags, 9, e.Document)
	x.FlaggedObject(e.Flags, 10, e.CachedPage)
	return x.buf
}
func (e *PredMessageMediaWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaWebPage)
	x.Bytes(e.Webpage.encode())
	return x.buf
}
func (e *PredAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authorization)
	x.Long(e.Hash)
	x.Int(e.Flags)
	x.String(e.DeviceModel)
	x.String(e.Platform)
	x.String(e.SystemVersion)
	x.Int(e.ApiId)
	x.String(e.AppName)
	x.String(e.AppVersion)
	x.Int(e.DateCreated)
	x.Int(e.DateActive)
	x.String(e.Ip)
	x.String(e.Country)
	x.String(e.Region)
	return x.buf
}
func (e *PredAccountAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountAuthorizations)
	x.Vector(toTLslice(e.Authorizations))
	return x.buf
}
func (e *PredAccountPasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountPasswordSettings)
	x.String(e.Email)
	return x.buf
}
func (e *PredAccountPasswordInputSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountPasswordInputSettings)
	x.Int(e.Flags)
	x.FlaggedStringBytes(e.Flags, 0, e.NewSalt)
	x.FlaggedStringBytes(e.Flags, 0, e.NewPasswordHash)
	x.FlaggedString(e.Flags, 0, e.Hint)
	x.FlaggedString(e.Flags, 1, e.Email)
	return x.buf
}
func (e *PredAuthPasswordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authPasswordRecovery)
	x.String(e.EmailPattern)
	return x.buf
}
func (e *PredInputMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaVenue)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueId)
	return x.buf
}
func (e *PredMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaVenue)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueId)
	return x.buf
}
func (e *PredReceivedNotifyMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_receivedNotifyMessage)
	x.Int(e.Id)
	x.Int(e.Flags)
	return x.buf
}
func (e *PredChatInviteEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteEmpty)
	return x.buf
}
func (e *PredChatInviteExported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteExported)
	x.String(e.Link)
	return x.buf
}
func (e *PredChatInviteAlready) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteAlready)
	x.Bytes(e.Chat.encode())
	return x.buf
}
func (e *PredChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInvite)
	x.Int(e.Flags)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.ParticipantsCount)
	x.FlaggedVector(e.Flags, 4, toTLslice(e.Participants))
	return x.buf
}
func (e *PredMessageActionChatJoinedByLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatJoinedByLink)
	x.Int(e.InviterId)
	return x.buf
}
func (e *PredUpdateReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadMessagesContents)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredInputStickerSetEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetEmpty)
	return x.buf
}
func (e *PredInputStickerSetID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetID)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredInputStickerSetShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetShortName)
	x.String(e.ShortName)
	return x.buf
}
func (e *PredStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSet)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.String(e.Title)
	x.String(e.ShortName)
	x.Int(e.Count)
	x.Int(e.Hash)
	return x.buf
}
func (e *PredMessagesStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesStickerSet)
	x.Bytes(e.Set.encode())
	x.Vector(toTLslice(e.Packs))
	x.Vector(toTLslice(e.Documents))
	return x.buf
}
func (e *PredUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_user)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.FlaggedLong(e.Flags, 0, e.AccessHash)
	x.FlaggedString(e.Flags, 1, e.FirstName)
	x.FlaggedString(e.Flags, 2, e.LastName)
	x.FlaggedString(e.Flags, 3, e.Username)
	x.FlaggedString(e.Flags, 4, e.Phone)
	x.FlaggedObject(e.Flags, 5, e.Photo)
	x.FlaggedObject(e.Flags, 6, e.Status)
	x.FlaggedInt(e.Flags, 14, e.BotInfoVersion)
	x.FlaggedString(e.Flags, 18, e.RestrictionReason)
	x.FlaggedString(e.Flags, 19, e.BotInlinePlaceholder)
	x.FlaggedString(e.Flags, 22, e.LangCode)
	return x.buf
}
func (e *PredBotCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botCommand)
	x.String(e.Command)
	x.String(e.Description)
	return x.buf
}
func (e *PredBotInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInfo)
	x.Int(e.UserId)
	x.String(e.Description)
	x.Vector(toTLslice(e.Commands))
	return x.buf
}
func (e *PredKeyboardButton) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButton)
	x.String(e.Text)
	return x.buf
}
func (e *PredKeyboardButtonRow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRow)
	x.Vector(toTLslice(e.Buttons))
	return x.buf
}
func (e *PredReplyKeyboardHide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardHide)
	x.Int(e.Flags)
	return x.buf
}
func (e *PredReplyKeyboardForceReply) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardForceReply)
	x.Int(e.Flags)
	return x.buf
}
func (e *PredReplyKeyboardMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardMarkup)
	x.Int(e.Flags)
	x.Vector(toTLslice(e.Rows))
	return x.buf
}
func (e *PredInputMessagesFilterUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterUrl)
	return x.buf
}
func (e *PredInputPeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerUser)
	x.Int(e.UserId)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredInputUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUser)
	x.Int(e.UserId)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredMessageEntityUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUnknown)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityMention) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMention)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityHashtag) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityHashtag)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityBotCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBotCommand)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityEmail)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBold)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityItalic)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityCode)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}
func (e *PredMessageEntityPre) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityPre)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Language)
	return x.buf
}
func (e *PredMessageEntityTextUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityTextUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Url)
	return x.buf
}
func (e *PredUpdateShortSentMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortSentMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	x.FlaggedObject(e.Flags, 9, e.Media)
	x.FlaggedVector(e.Flags, 7, toTLslice(e.Entities))
	return x.buf
}
func (e *PredInputPeerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChannel)
	x.Int(e.ChannelId)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredPeerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChannel)
	x.Int(e.ChannelId)
	return x.buf
}
func (e *PredChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channel)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.FlaggedLong(e.Flags, 13, e.AccessHash)
	x.String(e.Title)
	x.FlaggedString(e.Flags, 6, e.Username)
	x.Bytes(e.Photo.encode())
	x.Int(e.Date)
	x.Int(e.Version)
	x.FlaggedString(e.Flags, 9, e.RestrictionReason)
	x.FlaggedObject(e.Flags, 14, e.AdminRights)
	x.FlaggedObject(e.Flags, 15, e.BannedRights)
	return x.buf
}
func (e *PredChannelForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelForbidden)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Long(e.AccessHash)
	x.String(e.Title)
	x.FlaggedInt(e.Flags, 16, e.UntilDate)
	return x.buf
}
func (e *PredChannelFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelFull)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.About)
	x.FlaggedInt(e.Flags, 0, e.ParticipantsCount)
	x.FlaggedInt(e.Flags, 1, e.AdminsCount)
	x.FlaggedInt(e.Flags, 2, e.KickedCount)
	x.FlaggedInt(e.Flags, 2, e.BannedCount)
	x.Int(e.ReadInboxMaxId)
	x.Int(e.ReadOutboxMaxId)
	x.Int(e.UnreadCount)
	x.Bytes(e.ChatPhoto.encode())
	x.Bytes(e.NotifySettings.encode())
	x.Bytes(e.ExportedInvite.encode())
	x.Vector(toTLslice(e.BotInfos))
	x.FlaggedInt(e.Flags, 4, e.MigratedFromChatId)
	x.FlaggedInt(e.Flags, 4, e.MigratedFromMaxId)
	x.FlaggedInt(e.Flags, 5, e.PinnedMsgId)
	x.FlaggedObject(e.Flags, 8, e.Stickerset)
	return x.buf
}
func (e *PredMessageActionChannelCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelCreate)
	x.String(e.Title)
	return x.buf
}
func (e *PredMessagesChannelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesChannelMessages)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Count)
	x.Vector(toTLslice(e.Messages))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredUpdateChannelTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelTooLong)
	x.Int(e.Flags)
	x.Int(e.ChannelId)
	x.FlaggedInt(e.Flags, 0, e.Pts)
	return x.buf
}
func (e *PredUpdateChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannel)
	x.Int(e.ChannelId)
	return x.buf
}
func (e *PredUpdateNewChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredUpdateReadChannelInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelInbox)
	x.Int(e.ChannelId)
	x.Int(e.MaxId)
	return x.buf
}
func (e *PredUpdateDeleteChannelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteChannelMessages)
	x.Int(e.ChannelId)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredUpdateChannelMessageViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelMessageViews)
	x.Int(e.ChannelId)
	x.Int(e.Id)
	x.Int(e.Views)
	return x.buf
}
func (e *PredInputChannelEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannelEmpty)
	return x.buf
}
func (e *PredInputChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannel)
	x.Int(e.ChannelId)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredContactsResolvedPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsResolvedPeer)
	x.Bytes(e.Peer.encode())
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredMessageRange) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageRange)
	x.Int(e.MinId)
	x.Int(e.MaxId)
	return x.buf
}
func (e *PredUpdatesChannelDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesChannelDifferenceEmpty)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.FlaggedInt(e.Flags, 1, e.Timeout)
	return x.buf
}
func (e *PredUpdatesChannelDifferenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesChannelDifferenceTooLong)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.FlaggedInt(e.Flags, 1, e.Timeout)
	x.Int(e.TopMessage)
	x.Int(e.ReadInboxMaxId)
	x.Int(e.ReadOutboxMaxId)
	x.Int(e.UnreadCount)
	x.Int(e.UnreadMentionsCount)
	x.Vector(toTLslice(e.Messages))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredUpdatesChannelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesChannelDifference)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.FlaggedInt(e.Flags, 1, e.Timeout)
	x.Vector(toTLslice(e.NewMessages))
	x.Vector(toTLslice(e.OtherUpdates))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredChannelMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilterEmpty)
	return x.buf
}
func (e *PredChannelMessagesFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilter)
	x.Int(e.Flags)
	x.Vector(toTLslice(e.Ranges))
	return x.buf
}
func (e *PredChannelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipant)
	x.Int(e.UserId)
	x.Int(e.Date)
	return x.buf
}
func (e *PredChannelParticipantSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantSelf)
	x.Int(e.UserId)
	x.Int(e.InviterId)
	x.Int(e.Date)
	return x.buf
}
func (e *PredChannelParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantCreator)
	x.Int(e.UserId)
	return x.buf
}
func (e *PredChannelParticipantsRecent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsRecent)
	return x.buf
}
func (e *PredChannelParticipantsAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsAdmins)
	return x.buf
}
func (e *PredChannelParticipantsKicked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsKicked)
	x.String(e.Q)
	return x.buf
}
func (e *PredChannelsChannelParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsChannelParticipants)
	x.Int(e.Count)
	x.Vector(toTLslice(e.Participants))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredChannelsChannelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsChannelParticipant)
	x.Bytes(e.Participant.encode())
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredTrue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_true)
	return x.buf
}
func (e *PredChatParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantCreator)
	x.Int(e.UserId)
	return x.buf
}
func (e *PredChatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantAdmin)
	x.Int(e.UserId)
	x.Int(e.InviterId)
	x.Int(e.Date)
	return x.buf
}
func (e *PredUpdateChatAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatAdmins)
	x.Int(e.ChatId)
	x.Bytes(e.Enabled.encode())
	x.Int(e.Version)
	return x.buf
}
func (e *PredUpdateChatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdmin)
	x.Int(e.ChatId)
	x.Int(e.UserId)
	x.Bytes(e.IsAdmin.encode())
	x.Int(e.Version)
	return x.buf
}
func (e *PredMessageActionChatMigrateTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatMigrateTo)
	x.Int(e.ChannelId)
	return x.buf
}
func (e *PredMessageActionChannelMigrateFrom) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelMigrateFrom)
	x.String(e.Title)
	x.Int(e.ChatId)
	return x.buf
}
func (e *PredChannelParticipantsBots) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsBots)
	return x.buf
}
func (e *PredInputReportReasonSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonSpam)
	return x.buf
}
func (e *PredInputReportReasonViolence) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonViolence)
	return x.buf
}
func (e *PredInputReportReasonPornography) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonPornography)
	return x.buf
}
func (e *PredInputReportReasonOther) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonOther)
	x.String(e.Text)
	return x.buf
}
func (e *PredUpdateNewStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}
func (e *PredUpdateStickerSetsOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSetsOrder)
	x.Int(e.Flags)
	x.VectorLong(e.Order)
	return x.buf
}
func (e *PredUpdateStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSets)
	return x.buf
}
func (e *PredHelpTermsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpTermsOfService)
	x.String(e.Text)
	return x.buf
}
func (e *PredFoundGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGif)
	x.String(e.Url)
	x.String(e.ThumbUrl)
	x.String(e.ContentUrl)
	x.String(e.ContentType)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}
func (e *PredInputMediaGifExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGifExternal)
	x.String(e.Url)
	x.String(e.Q)
	return x.buf
}
func (e *PredMessagesFoundGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesFoundGifs)
	x.Int(e.NextOffset)
	x.Vector(toTLslice(e.Results))
	return x.buf
}
func (e *PredInputMessagesFilterGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterGif)
	return x.buf
}
func (e *PredUpdateSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateSavedGifs)
	return x.buf
}
func (e *PredUpdateBotInlineQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineQuery)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.Int(e.UserId)
	x.String(e.Query)
	x.FlaggedObject(e.Flags, 0, e.Geo)
	x.String(e.Offset)
	return x.buf
}
func (e *PredFoundGifCached) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGifCached)
	x.String(e.Url)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Document.encode())
	return x.buf
}
func (e *PredMessagesSavedGifsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSavedGifsNotModified)
	return x.buf
}
func (e *PredMessagesSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSavedGifs)
	x.Int(e.Hash)
	x.Vector(toTLslice(e.Gifs))
	return x.buf
}
func (e *PredInputBotInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Caption)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredInputBotInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageText)
	x.Int(e.Flags)
	x.String(e.Message)
	x.FlaggedVector(e.Flags, 1, toTLslice(e.Entities))
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredInputBotInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e.Type)
	x.FlaggedString(e.Flags, 1, e.Title)
	x.FlaggedString(e.Flags, 2, e.Description)
	x.FlaggedString(e.Flags, 3, e.Url)
	x.FlaggedString(e.Flags, 4, e.ThumbUrl)
	x.FlaggedString(e.Flags, 5, e.ContentUrl)
	x.FlaggedString(e.Flags, 5, e.ContentType)
	x.FlaggedInt(e.Flags, 6, e.W)
	x.FlaggedInt(e.Flags, 6, e.H)
	x.FlaggedInt(e.Flags, 7, e.Duration)
	x.Bytes(e.SendMessage.encode())
	return x.buf
}
func (e *PredBotInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Caption)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredBotInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageText)
	x.Int(e.Flags)
	x.String(e.Message)
	x.FlaggedVector(e.Flags, 1, toTLslice(e.Entities))
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredBotInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e.Type)
	x.FlaggedString(e.Flags, 1, e.Title)
	x.FlaggedString(e.Flags, 2, e.Description)
	x.FlaggedString(e.Flags, 3, e.Url)
	x.FlaggedString(e.Flags, 4, e.ThumbUrl)
	x.FlaggedString(e.Flags, 5, e.ContentUrl)
	x.FlaggedString(e.Flags, 5, e.ContentType)
	x.FlaggedInt(e.Flags, 6, e.W)
	x.FlaggedInt(e.Flags, 6, e.H)
	x.FlaggedInt(e.Flags, 7, e.Duration)
	x.Bytes(e.SendMessage.encode())
	return x.buf
}
func (e *PredMessagesBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesBotResults)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.FlaggedString(e.Flags, 1, e.NextOffset)
	x.FlaggedObject(e.Flags, 2, e.SwitchPm)
	x.Vector(toTLslice(e.Results))
	x.Int(e.CacheTime)
	return x.buf
}
func (e *PredInputMessagesFilterVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVoice)
	return x.buf
}
func (e *PredInputMessagesFilterMusic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMusic)
	return x.buf
}
func (e *PredUpdateBotInlineSend) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineSend)
	x.Int(e.Flags)
	x.Int(e.UserId)
	x.String(e.Query)
	x.FlaggedObject(e.Flags, 0, e.Geo)
	x.String(e.Id)
	x.FlaggedObject(e.Flags, 1, e.MsgId)
	return x.buf
}
func (e *PredInputPrivacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyChatInvite)
	return x.buf
}
func (e *PredPrivacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyChatInvite)
	return x.buf
}
func (e *PredUpdateEditChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredExportedMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_exportedMessageLink)
	x.String(e.Link)
	return x.buf
}
func (e *PredMessageFwdHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageFwdHeader)
	x.Int(e.Flags)
	x.FlaggedInt(e.Flags, 0, e.FromId)
	x.Int(e.Date)
	x.FlaggedInt(e.Flags, 1, e.ChannelId)
	x.FlaggedInt(e.Flags, 2, e.ChannelPost)
	x.FlaggedString(e.Flags, 3, e.PostAuthor)
	return x.buf
}
func (e *PredMessageActionPinMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPinMessage)
	return x.buf
}
func (e *PredPeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerSettings)
	x.Int(e.Flags)
	return x.buf
}
func (e *PredUpdateChannelPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelPinnedMessage)
	x.Int(e.ChannelId)
	x.Int(e.Id)
	return x.buf
}
func (e *PredKeyboardButtonUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonUrl)
	x.String(e.Text)
	x.String(e.Url)
	return x.buf
}
func (e *PredKeyboardButtonCallback) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonCallback)
	x.String(e.Text)
	x.StringBytes(e.Data)
	return x.buf
}
func (e *PredKeyboardButtonRequestPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestPhone)
	x.String(e.Text)
	return x.buf
}
func (e *PredKeyboardButtonRequestGeoLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestGeoLocation)
	x.String(e.Text)
	return x.buf
}
func (e *PredAuthCodeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authCodeTypeSms)
	return x.buf
}
func (e *PredAuthCodeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authCodeTypeCall)
	return x.buf
}
func (e *PredAuthCodeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authCodeTypeFlashCall)
	return x.buf
}
func (e *PredAuthSentCodeTypeApp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSentCodeTypeApp)
	x.Int(e.Length)
	return x.buf
}
func (e *PredAuthSentCodeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSentCodeTypeSms)
	x.Int(e.Length)
	return x.buf
}
func (e *PredAuthSentCodeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSentCodeTypeCall)
	x.Int(e.Length)
	return x.buf
}
func (e *PredAuthSentCodeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSentCodeTypeFlashCall)
	x.String(e.Pattern)
	return x.buf
}
func (e *PredKeyboardButtonSwitchInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonSwitchInline)
	x.Int(e.Flags)
	x.String(e.Text)
	x.String(e.Query)
	return x.buf
}
func (e *PredReplyInlineMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyInlineMarkup)
	x.Vector(toTLslice(e.Rows))
	return x.buf
}
func (e *PredMessagesBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesBotCallbackAnswer)
	x.Int(e.Flags)
	x.FlaggedString(e.Flags, 0, e.Message)
	x.FlaggedString(e.Flags, 2, e.Url)
	x.Int(e.CacheTime)
	return x.buf
}
func (e *PredUpdateBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.Int(e.UserId)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgId)
	x.Long(e.ChatInstance)
	x.FlaggedStringBytes(e.Flags, 0, e.Data)
	x.FlaggedString(e.Flags, 1, e.GameShortName)
	return x.buf
}
func (e *PredMessagesMessageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesMessageEditData)
	x.Int(e.Flags)
	return x.buf
}
func (e *PredUpdateEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredInputBotInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.GeoPoint.encode())
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredInputBotInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueId)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredInputBotInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredBotInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.Geo.encode())
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredBotInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueId)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredBotInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredInputBotInlineResultPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultPhoto)
	x.String(e.Id)
	x.String(e.Type)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.SendMessage.encode())
	return x.buf
}
func (e *PredInputBotInlineResultDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultDocument)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e.Type)
	x.FlaggedString(e.Flags, 1, e.Title)
	x.FlaggedString(e.Flags, 2, e.Description)
	x.Bytes(e.Document.encode())
	x.Bytes(e.SendMessage.encode())
	return x.buf
}
func (e *PredBotInlineMediaResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMediaResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e.Type)
	x.FlaggedObject(e.Flags, 0, e.Photo)
	x.FlaggedObject(e.Flags, 1, e.Document)
	x.FlaggedString(e.Flags, 2, e.Title)
	x.FlaggedString(e.Flags, 3, e.Description)
	x.Bytes(e.SendMessage.encode())
	return x.buf
}
func (e *PredInputBotInlineMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageID)
	x.Int(e.DcId)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredUpdateInlineBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateInlineBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.Int(e.UserId)
	x.Bytes(e.MsgId.encode())
	x.Long(e.ChatInstance)
	x.FlaggedStringBytes(e.Flags, 0, e.Data)
	x.FlaggedString(e.Flags, 1, e.GameShortName)
	return x.buf
}
func (e *PredInlineBotSwitchPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inlineBotSwitchPM)
	x.String(e.Text)
	x.String(e.StartParam)
	return x.buf
}
func (e *PredMessageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Int(e.UserId)
	return x.buf
}
func (e *PredInputMessageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Bytes(e.UserId.encode())
	return x.buf
}
func (e *PredMessagesPeerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesPeerDialogs)
	x.Vector(toTLslice(e.Dialogs))
	x.Vector(toTLslice(e.Messages))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	x.Bytes(e.State.encode())
	return x.buf
}
func (e *PredTopPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeer)
	x.Bytes(e.Peer.encode())
	x.Double(e.Rating)
	return x.buf
}
func (e *PredTopPeerCategoryBotsPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsPM)
	return x.buf
}
func (e *PredTopPeerCategoryBotsInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsInline)
	return x.buf
}
func (e *PredTopPeerCategoryCorrespondents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryCorrespondents)
	return x.buf
}
func (e *PredTopPeerCategoryGroups) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryGroups)
	return x.buf
}
func (e *PredTopPeerCategoryChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryChannels)
	return x.buf
}
func (e *PredTopPeerCategoryPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryPeers)
	x.Bytes(e.Category.encode())
	x.Int(e.Count)
	x.Vector(toTLslice(e.Peers))
	return x.buf
}
func (e *PredContactsTopPeersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsTopPeersNotModified)
	return x.buf
}
func (e *PredContactsTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsTopPeers)
	x.Vector(toTLslice(e.Categories))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredInputMessagesFilterChatPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterChatPhotos)
	return x.buf
}
func (e *PredUpdateReadChannelOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelOutbox)
	x.Int(e.ChannelId)
	x.Int(e.MaxId)
	return x.buf
}
func (e *PredUpdateDraftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDraftMessage)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Draft.encode())
	return x.buf
}
func (e *PredDraftMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessageEmpty)
	return x.buf
}
func (e *PredDraftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessage)
	x.Int(e.Flags)
	x.FlaggedInt(e.Flags, 0, e.ReplyToMsgId)
	x.String(e.Message)
	x.FlaggedVector(e.Flags, 3, toTLslice(e.Entities))
	x.Int(e.Date)
	return x.buf
}
func (e *PredMessageActionHistoryClear) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionHistoryClear)
	return x.buf
}
func (e *PredUpdateReadFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadFeaturedStickers)
	return x.buf
}
func (e *PredUpdateRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateRecentStickers)
	return x.buf
}
func (e *PredMessagesFeaturedStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesFeaturedStickersNotModified)
	return x.buf
}
func (e *PredMessagesFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesFeaturedStickers)
	x.Int(e.Hash)
	x.Vector(toTLslice(e.Sets))
	x.VectorLong(e.Unread)
	return x.buf
}
func (e *PredMessagesRecentStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesRecentStickersNotModified)
	return x.buf
}
func (e *PredMessagesRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesRecentStickers)
	x.Int(e.Hash)
	x.Vector(toTLslice(e.Stickers))
	return x.buf
}
func (e *PredMessagesArchivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesArchivedStickers)
	x.Int(e.Count)
	x.Vector(toTLslice(e.Sets))
	return x.buf
}
func (e *PredMessagesStickerSetInstallResultSuccess) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesStickerSetInstallResultSuccess)
	return x.buf
}
func (e *PredMessagesStickerSetInstallResultArchive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesStickerSetInstallResultArchive)
	x.Vector(toTLslice(e.Sets))
	return x.buf
}
func (e *PredStickerSetCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetCovered)
	x.Bytes(e.Set.encode())
	x.Bytes(e.Cover.encode())
	return x.buf
}
func (e *PredInputMediaPhotoExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhotoExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	x.String(e.Caption)
	x.FlaggedInt(e.Flags, 0, e.TtlSeconds)
	return x.buf
}
func (e *PredInputMediaDocumentExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocumentExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	x.String(e.Caption)
	x.FlaggedInt(e.Flags, 0, e.TtlSeconds)
	return x.buf
}
func (e *PredUpdateConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateConfig)
	return x.buf
}
func (e *PredUpdatePtsChanged) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePtsChanged)
	return x.buf
}
func (e *PredMessageActionGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionGameScore)
	x.Long(e.GameId)
	x.Int(e.Score)
	return x.buf
}
func (e *PredDocumentAttributeHasStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeHasStickers)
	return x.buf
}
func (e *PredKeyboardButtonGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonGame)
	x.String(e.Text)
	return x.buf
}
func (e *PredStickerSetMultiCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetMultiCovered)
	x.Bytes(e.Set.encode())
	x.Vector(toTLslice(e.Covers))
	return x.buf
}
func (e *PredMaskCoords) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_maskCoords)
	x.Int(e.N)
	x.Double(e.X)
	x.Double(e.Y)
	x.Double(e.Zoom)
	return x.buf
}
func (e *PredInputStickeredMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickeredMediaPhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *PredInputStickeredMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickeredMediaDocument)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *PredInputMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGame)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *PredMessageMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGame)
	x.Bytes(e.Game.encode())
	return x.buf
}
func (e *PredInputBotInlineMessageGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageGame)
	x.Int(e.Flags)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *PredInputBotInlineResultGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultGame)
	x.String(e.Id)
	x.String(e.ShortName)
	x.Bytes(e.SendMessage.encode())
	return x.buf
}
func (e *PredGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_game)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.String(e.ShortName)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.encode())
	x.FlaggedObject(e.Flags, 0, e.Document)
	return x.buf
}
func (e *PredInputGameID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGameID)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredInputGameShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGameShortName)
	x.Bytes(e.BotId.encode())
	x.String(e.ShortName)
	return x.buf
}
func (e *PredHighScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_highScore)
	x.Int(e.Pos)
	x.Int(e.UserId)
	x.Int(e.Score)
	return x.buf
}
func (e *PredMessagesHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesHighScores)
	x.Vector(toTLslice(e.Scores))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredMessagesChatsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesChatsSlice)
	x.Int(e.Count)
	x.Vector(toTLslice(e.Chats))
	return x.buf
}
func (e *PredUpdateChannelWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelWebPage)
	x.Int(e.ChannelId)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}
func (e *PredUpdatesDifferenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesDifferenceTooLong)
	x.Int(e.Pts)
	return x.buf
}
func (e *PredSendMessageGamePlayAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGamePlayAction)
	return x.buf
}
func (e *PredWebPageNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageNotModified)
	return x.buf
}
func (e *PredTextEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textEmpty)
	return x.buf
}
func (e *PredTextPlain) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textPlain)
	x.String(e.Text)
	return x.buf
}
func (e *PredTextBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textBold)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredTextItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textItalic)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredTextUnderline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textUnderline)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredTextStrike) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textStrike)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredTextFixed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textFixed)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredTextUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textUrl)
	x.Bytes(e.Text.encode())
	x.String(e.Url)
	x.Long(e.WebpageId)
	return x.buf
}
func (e *PredTextEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textEmail)
	x.Bytes(e.Text.encode())
	x.String(e.Email)
	return x.buf
}
func (e *PredTextConcat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textConcat)
	x.Vector(toTLslice(e.Texts))
	return x.buf
}
func (e *PredPageBlockTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockTitle)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredPageBlockSubtitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSubtitle)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredPageBlockAuthorDate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAuthorDate)
	x.Bytes(e.Author.encode())
	x.Int(e.PublishedDate)
	return x.buf
}
func (e *PredPageBlockHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockHeader)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredPageBlockSubheader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSubheader)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredPageBlockParagraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockParagraph)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredPageBlockPreformatted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPreformatted)
	x.Bytes(e.Text.encode())
	x.String(e.Language)
	return x.buf
}
func (e *PredPageBlockFooter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockFooter)
	x.Bytes(e.Text.encode())
	return x.buf
}
func (e *PredPageBlockDivider) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockDivider)
	return x.buf
}
func (e *PredPageBlockList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockList)
	x.Bytes(e.Ordered.encode())
	x.Vector(toTLslice(e.Items))
	return x.buf
}
func (e *PredPageBlockBlockquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockBlockquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredPageBlockPullquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPullquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredPageBlockPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPhoto)
	x.Long(e.PhotoId)
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredPageBlockVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockVideo)
	x.Int(e.Flags)
	x.Long(e.VideoId)
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredPageBlockCover) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockCover)
	x.Bytes(e.Cover.encode())
	return x.buf
}
func (e *PredPageBlockEmbed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockEmbed)
	x.Int(e.Flags)
	x.FlaggedString(e.Flags, 1, e.Url)
	x.FlaggedString(e.Flags, 2, e.Html)
	x.FlaggedLong(e.Flags, 4, e.PosterPhotoId)
	x.Int(e.W)
	x.Int(e.H)
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredPageBlockEmbedPost) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockEmbedPost)
	x.String(e.Url)
	x.Long(e.WebpageId)
	x.Long(e.AuthorPhotoId)
	x.String(e.Author)
	x.Int(e.Date)
	x.Vector(toTLslice(e.Blocks))
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredPageBlockSlideshow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSlideshow)
	x.Vector(toTLslice(e.Items))
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredPagePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pagePart)
	x.Vector(toTLslice(e.Blocks))
	x.Vector(toTLslice(e.Photos))
	x.Vector(toTLslice(e.Documents))
	return x.buf
}
func (e *PredPageFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageFull)
	x.Vector(toTLslice(e.Blocks))
	x.Vector(toTLslice(e.Photos))
	x.Vector(toTLslice(e.Documents))
	return x.buf
}
func (e *PredUpdatePhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePhoneCall)
	x.Bytes(e.PhoneCall.encode())
	return x.buf
}
func (e *PredUpdateDialogPinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDialogPinned)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *PredUpdatePinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePinnedDialogs)
	x.Int(e.Flags)
	x.FlaggedVector(e.Flags, 0, toTLslice(e.Order))
	return x.buf
}
func (e *PredInputPrivacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyPhoneCall)
	return x.buf
}
func (e *PredPrivacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyPhoneCall)
	return x.buf
}
func (e *PredPageBlockUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockUnsupported)
	return x.buf
}
func (e *PredPageBlockAnchor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAnchor)
	x.String(e.Name)
	return x.buf
}
func (e *PredPageBlockCollage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockCollage)
	x.Vector(toTLslice(e.Items))
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredInputPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneCall)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredPhoneCallEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallEmpty)
	x.Long(e.Id)
	return x.buf
}
func (e *PredPhoneCallWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallWaiting)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	x.Bytes(e.Protocol.encode())
	x.FlaggedInt(e.Flags, 0, e.ReceiveDate)
	return x.buf
}
func (e *PredPhoneCallRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallRequested)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	x.StringBytes(e.GAHash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}
func (e *PredPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCall)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	x.StringBytes(e.GAOrB)
	x.Long(e.KeyFingerprint)
	x.Bytes(e.Protocol.encode())
	x.Bytes(e.Connection.encode())
	x.Vector(toTLslice(e.AlternativeConnections))
	x.Int(e.StartDate)
	return x.buf
}
func (e *PredPhoneCallDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscarded)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.FlaggedObject(e.Flags, 0, e.Reason)
	x.FlaggedInt(e.Flags, 1, e.Duration)
	return x.buf
}
func (e *PredPhoneConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneConnection)
	x.Long(e.Id)
	x.String(e.Ip)
	x.String(e.Ipv6)
	x.Int(e.Port)
	x.StringBytes(e.PeerTag)
	return x.buf
}
func (e *PredPhoneCallProtocol) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallProtocol)
	x.Int(e.Flags)
	x.Int(e.MinLayer)
	x.Int(e.MaxLayer)
	return x.buf
}
func (e *PredPhonePhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phonePhoneCall)
	x.Bytes(e.PhoneCall.encode())
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredPhoneCallDiscardReasonMissed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonMissed)
	return x.buf
}
func (e *PredPhoneCallDiscardReasonDisconnect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonDisconnect)
	return x.buf
}
func (e *PredPhoneCallDiscardReasonHangup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonHangup)
	return x.buf
}
func (e *PredPhoneCallDiscardReasonBusy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonBusy)
	return x.buf
}
func (e *PredInputMessagesFilterPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhoneCalls)
	x.Int(e.Flags)
	return x.buf
}
func (e *PredMessageActionPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPhoneCall)
	x.Int(e.Flags)
	x.Long(e.CallId)
	x.FlaggedObject(e.Flags, 0, e.Reason)
	x.FlaggedInt(e.Flags, 1, e.Duration)
	return x.buf
}
func (e *PredInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invoice)
	x.Int(e.Flags)
	x.String(e.Currency)
	x.Vector(toTLslice(e.Prices))
	return x.buf
}
func (e *PredInputMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaInvoice)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.Description)
	x.FlaggedObject(e.Flags, 0, e.Photo)
	x.Bytes(e.Invoice.encode())
	x.StringBytes(e.Payload)
	x.String(e.Provider)
	x.String(e.StartParam)
	return x.buf
}
func (e *PredMessageActionPaymentSentMe) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPaymentSentMe)
	x.Int(e.Flags)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.StringBytes(e.Payload)
	x.FlaggedObject(e.Flags, 0, e.Info)
	x.FlaggedString(e.Flags, 1, e.ShippingOptionId)
	x.Bytes(e.Charge.encode())
	return x.buf
}
func (e *PredMessageMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaInvoice)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.Description)
	x.FlaggedObject(e.Flags, 0, e.Photo)
	x.FlaggedInt(e.Flags, 2, e.ReceiptMsgId)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.String(e.StartParam)
	return x.buf
}
func (e *PredKeyboardButtonBuy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonBuy)
	x.String(e.Text)
	return x.buf
}
func (e *PredMessageActionPaymentSent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPaymentSent)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	return x.buf
}
func (e *PredPaymentsPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsPaymentForm)
	x.Int(e.Flags)
	x.Int(e.BotId)
	x.Bytes(e.Invoice.encode())
	x.Int(e.ProviderId)
	x.String(e.Url)
	x.FlaggedString(e.Flags, 4, e.NativeProvider)
	x.FlaggedObject(e.Flags, 4, e.NativeParams)
	x.FlaggedObject(e.Flags, 0, e.SavedInfo)
	x.FlaggedObject(e.Flags, 1, e.SavedCredentials)
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredPostAddress) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_postAddress)
	x.String(e.StreetLine1)
	x.String(e.StreetLine2)
	x.String(e.City)
	x.String(e.State)
	x.String(e.CountryIso2)
	x.String(e.PostCode)
	return x.buf
}
func (e *PredPaymentRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentRequestedInfo)
	x.Int(e.Flags)
	x.FlaggedString(e.Flags, 0, e.Name)
	x.FlaggedString(e.Flags, 1, e.Phone)
	x.FlaggedString(e.Flags, 2, e.Email)
	x.FlaggedObject(e.Flags, 3, e.ShippingAddress)
	return x.buf
}
func (e *PredUpdateBotWebhookJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotWebhookJSON)
	x.Bytes(e.Data.encode())
	return x.buf
}
func (e *PredUpdateBotWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotWebhookJSONQuery)
	x.Long(e.QueryId)
	x.Bytes(e.Data.encode())
	x.Int(e.Timeout)
	return x.buf
}
func (e *PredUpdateBotShippingQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotShippingQuery)
	x.Long(e.QueryId)
	x.Int(e.UserId)
	x.StringBytes(e.Payload)
	x.Bytes(e.ShippingAddress.encode())
	return x.buf
}
func (e *PredUpdateBotPrecheckoutQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotPrecheckoutQuery)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.Int(e.UserId)
	x.StringBytes(e.Payload)
	x.FlaggedObject(e.Flags, 0, e.Info)
	x.FlaggedString(e.Flags, 1, e.ShippingOptionId)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	return x.buf
}
func (e *PredDataJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dataJSON)
	x.String(e.Data)
	return x.buf
}
func (e *PredLabeledPrice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_labeledPrice)
	x.String(e.Label)
	x.Long(e.Amount)
	return x.buf
}
func (e *PredPaymentCharge) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentCharge)
	x.String(e.Id)
	x.String(e.ProviderChargeId)
	return x.buf
}
func (e *PredPaymentSavedCredentialsCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentSavedCredentialsCard)
	x.String(e.Id)
	x.String(e.Title)
	return x.buf
}
func (e *PredWebDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webDocument)
	x.String(e.Url)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Vector(toTLslice(e.Attributes))
	x.Int(e.DcId)
	return x.buf
}
func (e *PredInputWebDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebDocument)
	x.String(e.Url)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Vector(toTLslice(e.Attributes))
	return x.buf
}
func (e *PredInputWebFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebFileLocation)
	x.String(e.Url)
	x.Long(e.AccessHash)
	return x.buf
}
func (e *PredUploadWebFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadWebFile)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Bytes(e.FileType.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *PredPaymentsValidatedRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsValidatedRequestedInfo)
	x.Int(e.Flags)
	x.FlaggedString(e.Flags, 0, e.Id)
	x.FlaggedVector(e.Flags, 1, toTLslice(e.ShippingOptions))
	return x.buf
}
func (e *PredPaymentsPaymentResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsPaymentResult)
	x.Bytes(e.Updates.encode())
	return x.buf
}
func (e *PredPaymentsPaymentVerficationNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsPaymentVerficationNeeded)
	x.String(e.Url)
	return x.buf
}
func (e *PredPaymentsPaymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsPaymentReceipt)
	x.Int(e.Flags)
	x.Int(e.Date)
	x.Int(e.BotId)
	x.Bytes(e.Invoice.encode())
	x.Int(e.ProviderId)
	x.FlaggedObject(e.Flags, 0, e.Info)
	x.FlaggedObject(e.Flags, 1, e.Shipping)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.String(e.CredentialsTitle)
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredPaymentsSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsSavedInfo)
	x.Int(e.Flags)
	x.FlaggedObject(e.Flags, 0, e.SavedInfo)
	return x.buf
}
func (e *PredInputPaymentCredentialsSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentialsSaved)
	x.String(e.Id)
	x.StringBytes(e.TmpPassword)
	return x.buf
}
func (e *PredInputPaymentCredentials) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentials)
	x.Int(e.Flags)
	x.Bytes(e.Data.encode())
	return x.buf
}
func (e *PredAccountTmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountTmpPassword)
	x.StringBytes(e.TmpPassword)
	x.Int(e.ValidUntil)
	return x.buf
}
func (e *PredShippingOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_shippingOption)
	x.String(e.Id)
	x.String(e.Title)
	x.Vector(toTLslice(e.Prices))
	return x.buf
}
func (e *PredPhoneCallAccepted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallAccepted)
	x.Long(e.Id)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminId)
	x.Int(e.ParticipantId)
	x.StringBytes(e.GB)
	x.Bytes(e.Protocol.encode())
	return x.buf
}
func (e *PredInputMessagesFilterRoundVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterRoundVoice)
	return x.buf
}
func (e *PredInputMessagesFilterRoundVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterRoundVideo)
	return x.buf
}
func (e *PredUploadFileCdnRedirect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadFileCdnRedirect)
	x.Int(e.DcId)
	x.StringBytes(e.FileToken)
	x.StringBytes(e.EncryptionKey)
	x.StringBytes(e.EncryptionIv)
	x.Vector(toTLslice(e.CdnFileHashes))
	return x.buf
}
func (e *PredSendMessageRecordRoundAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordRoundAction)
	return x.buf
}
func (e *PredSendMessageUploadRoundAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadRoundAction)
	x.Int(e.Progress)
	return x.buf
}
func (e *PredUploadCdnFileReuploadNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadCdnFileReuploadNeeded)
	x.StringBytes(e.RequestToken)
	return x.buf
}
func (e *PredUploadCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadCdnFile)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *PredCdnPublicKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_cdnPublicKey)
	x.Int(e.DcId)
	x.String(e.PublicKey)
	return x.buf
}
func (e *PredCdnConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_cdnConfig)
	x.Vector(toTLslice(e.PublicKeys))
	return x.buf
}
func (e *PredUpdateLangPackTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateLangPackTooLong)
	return x.buf
}
func (e *PredUpdateLangPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateLangPack)
	x.Bytes(e.Difference.encode())
	return x.buf
}
func (e *PredPageBlockChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}
func (e *PredInputStickerSetItem) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetItem)
	x.Int(e.Flags)
	x.Bytes(e.Document.encode())
	x.String(e.Emoji)
	x.FlaggedObject(e.Flags, 0, e.MaskCoords)
	return x.buf
}
func (e *PredLangPackString) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackString)
	x.String(e.Key)
	x.String(e.Value)
	return x.buf
}
func (e *PredLangPackStringPluralized) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackStringPluralized)
	x.Int(e.Flags)
	x.String(e.Key)
	x.FlaggedString(e.Flags, 0, e.ZeroValue)
	x.FlaggedString(e.Flags, 1, e.OneValue)
	x.FlaggedString(e.Flags, 2, e.TwoValue)
	x.FlaggedString(e.Flags, 3, e.FewValue)
	x.FlaggedString(e.Flags, 4, e.ManyValue)
	x.String(e.OtherValue)
	return x.buf
}
func (e *PredLangPackStringDeleted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackStringDeleted)
	x.String(e.Key)
	return x.buf
}
func (e *PredLangPackDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackDifference)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	x.Int(e.Version)
	x.Vector(toTLslice(e.Strings))
	return x.buf
}
func (e *PredLangPackLanguage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackLanguage)
	x.String(e.Name)
	x.String(e.NativeName)
	x.String(e.LangCode)
	return x.buf
}
func (e *PredChannelParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantAdmin)
	x.Int(e.Flags)
	x.Int(e.UserId)
	x.Int(e.InviterId)
	x.Int(e.PromotedBy)
	x.Int(e.Date)
	x.Bytes(e.AdminRights.encode())
	return x.buf
}
func (e *PredChannelParticipantBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantBanned)
	x.Int(e.Flags)
	x.Int(e.UserId)
	x.Int(e.KickedBy)
	x.Int(e.Date)
	x.Bytes(e.BannedRights.encode())
	return x.buf
}
func (e *PredChannelParticipantsBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsBanned)
	x.String(e.Q)
	return x.buf
}
func (e *PredChannelParticipantsSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsSearch)
	x.String(e.Q)
	return x.buf
}
func (e *PredTopPeerCategoryPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryPhoneCalls)
	return x.buf
}
func (e *PredPageBlockAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAudio)
	x.Long(e.AudioId)
	x.Bytes(e.Caption.encode())
	return x.buf
}
func (e *PredChannelAdminRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminRights)
	x.Int(e.Flags)
	return x.buf
}
func (e *PredChannelBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelBannedRights)
	x.Int(e.Flags)
	x.Int(e.UntilDate)
	return x.buf
}
func (e *PredChannelAdminLogEventActionChangeTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeTitle)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}
func (e *PredChannelAdminLogEventActionChangeAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeAbout)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}
func (e *PredChannelAdminLogEventActionChangeUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeUsername)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}
func (e *PredChannelAdminLogEventActionChangePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangePhoto)
	x.Bytes(e.PrevPhoto.encode())
	x.Bytes(e.NewPhoto.encode())
	return x.buf
}
func (e *PredChannelAdminLogEventActionToggleInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionToggleInvites)
	x.Bytes(e.NewValue.encode())
	return x.buf
}
func (e *PredChannelAdminLogEventActionToggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionToggleSignatures)
	x.Bytes(e.NewValue.encode())
	return x.buf
}
func (e *PredChannelAdminLogEventActionUpdatePinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionUpdatePinned)
	x.Bytes(e.Message.encode())
	return x.buf
}
func (e *PredChannelAdminLogEventActionEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionEditMessage)
	x.Bytes(e.PrevMessage.encode())
	x.Bytes(e.NewMessage.encode())
	return x.buf
}
func (e *PredChannelAdminLogEventActionDeleteMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionDeleteMessage)
	x.Bytes(e.Message.encode())
	return x.buf
}
func (e *PredChannelAdminLogEventActionParticipantJoin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantJoin)
	return x.buf
}
func (e *PredChannelAdminLogEventActionParticipantLeave) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantLeave)
	return x.buf
}
func (e *PredChannelAdminLogEventActionParticipantInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantInvite)
	x.Bytes(e.Participant.encode())
	return x.buf
}
func (e *PredChannelAdminLogEventActionParticipantToggleBan) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantToggleBan)
	x.Bytes(e.PrevParticipant.encode())
	x.Bytes(e.NewParticipant.encode())
	return x.buf
}
func (e *PredChannelAdminLogEventActionParticipantToggleAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantToggleAdmin)
	x.Bytes(e.PrevParticipant.encode())
	x.Bytes(e.NewParticipant.encode())
	return x.buf
}
func (e *PredChannelAdminLogEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEvent)
	x.Long(e.Id)
	x.Int(e.Date)
	x.Int(e.UserId)
	x.Bytes(e.Action.encode())
	return x.buf
}
func (e *PredChannelsAdminLogResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsAdminLogResults)
	x.Vector(toTLslice(e.Events))
	x.Vector(toTLslice(e.Chats))
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *PredChannelAdminLogEventsFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventsFilter)
	x.Int(e.Flags)
	return x.buf
}
func (e *PredMessageActionScreenshotTaken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionScreenshotTaken)
	return x.buf
}
func (e *PredPopularContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_popularContact)
	x.Long(e.ClientId)
	x.Int(e.Importers)
	return x.buf
}
func (e *PredCdnFileHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_cdnFileHash)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.StringBytes(e.Hash)
	return x.buf
}
func (e *PredInputMessagesFilterMyMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMyMentions)
	return x.buf
}
func (e *PredInputMessagesFilterMyMentionsUnread) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMyMentionsUnread)
	return x.buf
}
func (e *PredUpdateContactsReset) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactsReset)
	return x.buf
}
func (e *PredChannelAdminLogEventActionChangeStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeStickerSet)
	x.Bytes(e.PrevStickerset.encode())
	x.Bytes(e.NewStickerset.encode())
	return x.buf
}
func (e *PredUpdateFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateFavedStickers)
	return x.buf
}
func (e *PredMessagesFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesFavedStickers)
	x.Int(e.Hash)
	x.Vector(toTLslice(e.Packs))
	x.Vector(toTLslice(e.Stickers))
	return x.buf
}
func (e *PredMessagesFavedStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesFavedStickersNotModified)
	return x.buf
}
func (e *PredUpdateChannelReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelReadMessagesContents)
	x.Int(e.ChannelId)
	x.VectorInt(e.Messages)
	return x.buf
}

// Encode funcs for methods
func (e *ReqInvokeAfterMsg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsg)
	x.Long(e.MsgId)
	unpackedQuery := unpack(e.Query)
	if unpackedQuery != nil {
		x.Bytes(unpackedQuery.encode())
	}
	return x.buf
}
func (e *ReqInvokeAfterMsgs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsgs)
	x.VectorLong(e.MsgIds)
	unpackedQuery := unpack(e.Query)
	if unpackedQuery != nil {
		x.Bytes(unpackedQuery.encode())
	}
	return x.buf
}
func (e *ReqAuthCheckPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authCheckPhone)
	x.String(e.PhoneNumber)
	return x.buf
}
func (e *ReqAuthSendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSendCode)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.FlaggedObject(e.Flags, 0, e.CurrentNumber)
	x.Int(e.ApiId)
	x.String(e.ApiHash)
	return x.buf
}
func (e *ReqAuthSignUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSignUp)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}
func (e *ReqAuthSignIn) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSignIn)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}
func (e *ReqAuthLogOut) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authLogOut)
	return x.buf
}
func (e *ReqAuthResetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authResetAuthorizations)
	return x.buf
}
func (e *ReqAuthSendInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authSendInvites)
	x.VectorString(e.PhoneNumbers)
	x.String(e.Message)
	return x.buf
}
func (e *ReqAuthExportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authExportAuthorization)
	x.Int(e.DcId)
	return x.buf
}
func (e *ReqAuthImportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authImportAuthorization)
	x.Int(e.Id)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *ReqAccountRegisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountRegisterDevice)
	x.Int(e.TokenType)
	x.String(e.Token)
	return x.buf
}
func (e *ReqAccountUnregisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountUnregisterDevice)
	x.Int(e.TokenType)
	x.String(e.Token)
	return x.buf
}
func (e *ReqAccountUpdateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountUpdateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}
func (e *ReqAccountGetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountGetNotifySettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *ReqAccountResetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountResetNotifySettings)
	return x.buf
}
func (e *ReqAccountUpdateProfile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountUpdateProfile)
	x.Int(e.Flags)
	x.FlaggedString(e.Flags, 0, e.FirstName)
	x.FlaggedString(e.Flags, 1, e.LastName)
	x.FlaggedString(e.Flags, 2, e.About)
	return x.buf
}
func (e *ReqAccountUpdateStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountUpdateStatus)
	x.Bytes(e.Offline.encode())
	return x.buf
}
func (e *ReqAccountGetWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountGetWallPapers)
	return x.buf
}
func (e *ReqUsersGetUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_usersGetUsers)
	x.Vector(toTLslice(e.Id))
	return x.buf
}
func (e *ReqUsersGetFullUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_usersGetFullUser)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *ReqContactsGetStatuses) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsGetStatuses)
	return x.buf
}
func (e *ReqContactsGetContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsGetContacts)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqContactsImportContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsImportContacts)
	x.Vector(toTLslice(e.Contacts))
	return x.buf
}
func (e *ReqContactsSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsSearch)
	x.String(e.Q)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqContactsDeleteContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsDeleteContact)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *ReqContactsDeleteContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsDeleteContacts)
	x.Vector(toTLslice(e.Id))
	return x.buf
}
func (e *ReqContactsBlock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsBlock)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *ReqContactsUnblock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsUnblock)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *ReqContactsGetBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsGetBlocked)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqMessagesGetMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetMessages)
	x.VectorInt(e.Id)
	return x.buf
}
func (e *ReqMessagesGetDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetDialogs)
	x.Int(e.Flags)
	x.Int(e.OffsetDate)
	x.Int(e.OffsetId)
	x.Bytes(e.OffsetPeer.encode())
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqMessagesGetHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.OffsetId)
	x.Int(e.OffsetDate)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxId)
	x.Int(e.MinId)
	return x.buf
}
func (e *ReqMessagesSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSearch)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.String(e.Q)
	x.FlaggedObject(e.Flags, 0, e.FromId)
	x.Bytes(e.Filter.encode())
	x.Int(e.MinDate)
	x.Int(e.MaxDate)
	x.Int(e.OffsetId)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxId)
	x.Int(e.MinId)
	return x.buf
}
func (e *ReqMessagesReadHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReadHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxId)
	return x.buf
}
func (e *ReqMessagesDeleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesDeleteHistory)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxId)
	return x.buf
}
func (e *ReqMessagesDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesDeleteMessages)
	x.Int(e.Flags)
	x.VectorInt(e.Id)
	return x.buf
}
func (e *ReqMessagesReceivedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReceivedMessages)
	x.Int(e.MaxId)
	return x.buf
}
func (e *ReqMessagesSetTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSetTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Action.encode())
	return x.buf
}
func (e *ReqMessagesSendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSendMessage)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.FlaggedInt(e.Flags, 0, e.ReplyToMsgId)
	x.String(e.Message)
	x.Long(e.RandomId)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	x.FlaggedVector(e.Flags, 3, toTLslice(e.Entities))
	return x.buf
}
func (e *ReqMessagesSendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSendMedia)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.FlaggedInt(e.Flags, 0, e.ReplyToMsgId)
	x.Bytes(e.Media.encode())
	x.Long(e.RandomId)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	return x.buf
}
func (e *ReqMessagesForwardMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesForwardMessages)
	x.Int(e.Flags)
	x.Bytes(e.FromPeer.encode())
	x.VectorInt(e.Id)
	x.VectorLong(e.RandomId)
	x.Bytes(e.ToPeer.encode())
	return x.buf
}
func (e *ReqMessagesGetChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetChats)
	x.VectorInt(e.Id)
	return x.buf
}
func (e *ReqMessagesGetFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetFullChat)
	x.Int(e.ChatId)
	return x.buf
}
func (e *ReqMessagesEditChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesEditChatTitle)
	x.Int(e.ChatId)
	x.String(e.Title)
	return x.buf
}
func (e *ReqMessagesEditChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesEditChatPhoto)
	x.Int(e.ChatId)
	x.Bytes(e.Photo.encode())
	return x.buf
}
func (e *ReqMessagesAddChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesAddChatUser)
	x.Int(e.ChatId)
	x.Bytes(e.UserId.encode())
	x.Int(e.FwdLimit)
	return x.buf
}
func (e *ReqMessagesDeleteChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesDeleteChatUser)
	x.Int(e.ChatId)
	x.Bytes(e.UserId.encode())
	return x.buf
}
func (e *ReqMessagesCreateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesCreateChat)
	x.Vector(toTLslice(e.Users))
	x.String(e.Title)
	return x.buf
}
func (e *ReqUpdatesGetState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesGetState)
	return x.buf
}
func (e *ReqUpdatesGetDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesGetDifference)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.FlaggedInt(e.Flags, 0, e.PtsTotalLimit)
	x.Int(e.Date)
	x.Int(e.Qts)
	return x.buf
}
func (e *ReqPhotosUpdateProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photosUpdateProfilePhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}
func (e *ReqPhotosUploadProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photosUploadProfilePhoto)
	x.Bytes(e.File.encode())
	return x.buf
}
func (e *ReqUploadSaveFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadSaveFilePart)
	x.Long(e.FileId)
	x.Int(e.FilePart)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *ReqUploadGetFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadGetFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqHelpGetConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpGetConfig)
	return x.buf
}
func (e *ReqHelpGetNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpGetNearestDc)
	return x.buf
}
func (e *ReqHelpGetAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpGetAppUpdate)
	return x.buf
}
func (e *ReqHelpSaveAppLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpSaveAppLog)
	x.Vector(toTLslice(e.Events))
	return x.buf
}
func (e *ReqHelpGetInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpGetInviteText)
	return x.buf
}
func (e *ReqPhotosDeletePhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photosDeletePhotos)
	x.Vector(toTLslice(e.Id))
	return x.buf
}
func (e *ReqPhotosGetUserPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photosGetUserPhotos)
	x.Bytes(e.UserId.encode())
	x.Int(e.Offset)
	x.Long(e.MaxId)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqMessagesForwardMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesForwardMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Long(e.RandomId)
	return x.buf
}
func (e *ReqMessagesGetDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetDhConfig)
	x.Int(e.Version)
	x.Int(e.RandomLength)
	return x.buf
}
func (e *ReqMessagesRequestEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesRequestEncryption)
	x.Bytes(e.UserId.encode())
	x.Int(e.RandomId)
	x.StringBytes(e.GA)
	return x.buf
}
func (e *ReqMessagesAcceptEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesAcceptEncryption)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GB)
	x.Long(e.KeyFingerprint)
	return x.buf
}
func (e *ReqMessagesDiscardEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesDiscardEncryption)
	x.Int(e.ChatId)
	return x.buf
}
func (e *ReqMessagesSetEncryptedTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSetEncryptedTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Typing.encode())
	return x.buf
}
func (e *ReqMessagesReadEncryptedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReadEncryptedHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxDate)
	return x.buf
}
func (e *ReqMessagesSendEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSendEncrypted)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomId)
	x.StringBytes(e.Data)
	return x.buf
}
func (e *ReqMessagesSendEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSendEncryptedFile)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomId)
	x.StringBytes(e.Data)
	x.Bytes(e.File.encode())
	return x.buf
}
func (e *ReqMessagesSendEncryptedService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSendEncryptedService)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomId)
	x.StringBytes(e.Data)
	return x.buf
}
func (e *ReqMessagesReceivedQueue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReceivedQueue)
	x.Int(e.MaxQts)
	return x.buf
}
func (e *ReqUploadSaveBigFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadSaveBigFilePart)
	x.Long(e.FileId)
	x.Int(e.FilePart)
	x.Int(e.FileTotalParts)
	x.StringBytes(e.Bytes)
	return x.buf
}
func (e *ReqInitConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_initConnection)
	x.Int(e.ApiId)
	x.String(e.DeviceModel)
	x.String(e.SystemVersion)
	x.String(e.AppVersion)
	x.String(e.SystemLangCode)
	x.String(e.LangPack)
	x.String(e.LangCode)
	unpackedQuery := unpack(e.Query)
	if unpackedQuery != nil {
		x.Bytes(unpackedQuery.encode())
	}
	return x.buf
}
func (e *ReqHelpGetSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpGetSupport)
	return x.buf
}
func (e *ReqAuthBindTempAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authBindTempAuthKey)
	x.Long(e.PermAuthKeyId)
	x.Long(e.Nonce)
	x.Int(e.ExpiresAt)
	x.StringBytes(e.EncryptedMessage)
	return x.buf
}
func (e *ReqContactsExportCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsExportCard)
	return x.buf
}
func (e *ReqContactsImportCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsImportCard)
	x.VectorInt(e.ExportCard)
	return x.buf
}
func (e *ReqMessagesReadMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReadMessageContents)
	x.VectorInt(e.Id)
	return x.buf
}
func (e *ReqAccountCheckUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountCheckUsername)
	x.String(e.Username)
	return x.buf
}
func (e *ReqAccountUpdateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountUpdateUsername)
	x.String(e.Username)
	return x.buf
}
func (e *ReqAccountGetPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountGetPrivacy)
	x.Bytes(e.Key.encode())
	return x.buf
}
func (e *ReqAccountSetPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountSetPrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(toTLslice(e.Rules))
	return x.buf
}
func (e *ReqAccountDeleteAccount) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountDeleteAccount)
	x.String(e.Reason)
	return x.buf
}
func (e *ReqAccountGetAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountGetAccountTTL)
	return x.buf
}
func (e *ReqAccountSetAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountSetAccountTTL)
	x.Bytes(e.Ttl.encode())
	return x.buf
}
func (e *ReqInvokeWithLayer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithLayer)
	x.Int(e.Layer)
	unpackedQuery := unpack(e.Query)
	if unpackedQuery != nil {
		x.Bytes(unpackedQuery.encode())
	}
	return x.buf
}
func (e *ReqContactsResolveUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsResolveUsername)
	x.String(e.Username)
	return x.buf
}
func (e *ReqAccountSendChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountSendChangePhoneCode)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.FlaggedObject(e.Flags, 0, e.CurrentNumber)
	return x.buf
}
func (e *ReqAccountChangePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountChangePhone)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}
func (e *ReqMessagesGetAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetAllStickers)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqAccountUpdateDeviceLocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountUpdateDeviceLocked)
	x.Int(e.Period)
	return x.buf
}
func (e *ReqAccountGetPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountGetPassword)
	return x.buf
}
func (e *ReqAuthCheckPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authCheckPassword)
	x.StringBytes(e.PasswordHash)
	return x.buf
}
func (e *ReqMessagesGetWebPagePreview) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetWebPagePreview)
	x.String(e.Message)
	return x.buf
}
func (e *ReqAccountGetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountGetAuthorizations)
	return x.buf
}
func (e *ReqAccountResetAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountResetAuthorization)
	x.Long(e.Hash)
	return x.buf
}
func (e *ReqAccountGetPasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountGetPasswordSettings)
	x.StringBytes(e.CurrentPasswordHash)
	return x.buf
}
func (e *ReqAccountUpdatePasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountUpdatePasswordSettings)
	x.StringBytes(e.CurrentPasswordHash)
	x.Bytes(e.NewSettings.encode())
	return x.buf
}
func (e *ReqAuthRequestPasswordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authRequestPasswordRecovery)
	return x.buf
}
func (e *ReqAuthRecoverPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authRecoverPassword)
	x.String(e.Code)
	return x.buf
}
func (e *ReqInvokeWithoutUpdates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithoutUpdates)
	unpackedQuery := unpack(e.Query)
	if unpackedQuery != nil {
		x.Bytes(unpackedQuery.encode())
	}
	return x.buf
}
func (e *ReqMessagesExportChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesExportChatInvite)
	x.Int(e.ChatId)
	return x.buf
}
func (e *ReqMessagesCheckChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesCheckChatInvite)
	x.String(e.Hash)
	return x.buf
}
func (e *ReqMessagesImportChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesImportChatInvite)
	x.String(e.Hash)
	return x.buf
}
func (e *ReqMessagesGetStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}
func (e *ReqMessagesInstallStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesInstallStickerSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Archived.encode())
	return x.buf
}
func (e *ReqMessagesUninstallStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesUninstallStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}
func (e *ReqAuthImportBotAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authImportBotAuthorization)
	x.Int(e.Flags)
	x.Int(e.ApiId)
	x.String(e.ApiHash)
	x.String(e.BotAuthToken)
	return x.buf
}
func (e *ReqMessagesStartBot) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesStartBot)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomId)
	x.String(e.StartParam)
	return x.buf
}
func (e *ReqHelpGetAppChangelog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpGetAppChangelog)
	x.String(e.PrevAppVersion)
	return x.buf
}
func (e *ReqMessagesReportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReportSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *ReqMessagesGetMessagesViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetMessagesViews)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.Id)
	x.Bytes(e.Increment.encode())
	return x.buf
}
func (e *ReqUpdatesGetChannelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesGetChannelDifference)
	x.Int(e.Flags)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Pts)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqChannelsReadHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsReadHistory)
	x.Bytes(e.Channel.encode())
	x.Int(e.MaxId)
	return x.buf
}
func (e *ReqChannelsDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsDeleteMessages)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.Id)
	return x.buf
}
func (e *ReqChannelsDeleteUserHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsDeleteUserHistory)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserId.encode())
	return x.buf
}
func (e *ReqChannelsReportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsReportSpam)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserId.encode())
	x.VectorInt(e.Id)
	return x.buf
}
func (e *ReqChannelsGetMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsGetMessages)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.Id)
	return x.buf
}
func (e *ReqChannelsGetParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsGetParticipants)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqChannelsGetParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsGetParticipant)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserId.encode())
	return x.buf
}
func (e *ReqChannelsGetChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsGetChannels)
	x.Vector(toTLslice(e.Id))
	return x.buf
}
func (e *ReqChannelsGetFullChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsGetFullChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}
func (e *ReqChannelsCreateChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsCreateChannel)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.About)
	return x.buf
}
func (e *ReqChannelsEditAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsEditAbout)
	x.Bytes(e.Channel.encode())
	x.String(e.About)
	return x.buf
}
func (e *ReqChannelsEditAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsEditAdmin)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserId.encode())
	x.Bytes(e.AdminRights.encode())
	return x.buf
}
func (e *ReqChannelsEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsEditTitle)
	x.Bytes(e.Channel.encode())
	x.String(e.Title)
	return x.buf
}
func (e *ReqChannelsEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsEditPhoto)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Photo.encode())
	return x.buf
}
func (e *ReqChannelsCheckUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsCheckUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}
func (e *ReqChannelsUpdateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsUpdateUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}
func (e *ReqChannelsJoinChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsJoinChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}
func (e *ReqChannelsLeaveChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsLeaveChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}
func (e *ReqChannelsInviteToChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsInviteToChannel)
	x.Bytes(e.Channel.encode())
	x.Vector(toTLslice(e.Users))
	return x.buf
}
func (e *ReqChannelsExportInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsExportInvite)
	x.Bytes(e.Channel.encode())
	return x.buf
}
func (e *ReqChannelsDeleteChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsDeleteChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}
func (e *ReqMessagesToggleChatAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesToggleChatAdmins)
	x.Int(e.ChatId)
	x.Bytes(e.Enabled.encode())
	return x.buf
}
func (e *ReqMessagesEditChatAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesEditChatAdmin)
	x.Int(e.ChatId)
	x.Bytes(e.UserId.encode())
	x.Bytes(e.IsAdmin.encode())
	return x.buf
}
func (e *ReqMessagesMigrateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesMigrateChat)
	x.Int(e.ChatId)
	return x.buf
}
func (e *ReqMessagesSearchGlobal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSearchGlobal)
	x.String(e.Q)
	x.Int(e.OffsetDate)
	x.Bytes(e.OffsetPeer.encode())
	x.Int(e.OffsetId)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqAccountReportPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountReportPeer)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Reason.encode())
	return x.buf
}
func (e *ReqMessagesReorderStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReorderStickerSets)
	x.Int(e.Flags)
	x.VectorLong(e.Order)
	return x.buf
}
func (e *ReqHelpGetTermsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpGetTermsOfService)
	return x.buf
}
func (e *ReqMessagesGetDocumentByHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetDocumentByHash)
	x.StringBytes(e.Sha256)
	x.Int(e.Size)
	x.String(e.MimeType)
	return x.buf
}
func (e *ReqMessagesSearchGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSearchGifs)
	x.String(e.Q)
	x.Int(e.Offset)
	return x.buf
}
func (e *ReqMessagesGetSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetSavedGifs)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqMessagesSaveGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSaveGif)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}
func (e *ReqMessagesGetInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetInlineBotResults)
	x.Int(e.Flags)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	x.FlaggedObject(e.Flags, 0, e.GeoPoint)
	x.String(e.Query)
	x.String(e.Offset)
	return x.buf
}
func (e *ReqMessagesSetInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSetInlineBotResults)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.Vector(toTLslice(e.Results))
	x.Int(e.CacheTime)
	x.FlaggedString(e.Flags, 2, e.NextOffset)
	x.FlaggedObject(e.Flags, 3, e.SwitchPm)
	return x.buf
}
func (e *ReqMessagesSendInlineBotResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSendInlineBotResult)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.FlaggedInt(e.Flags, 0, e.ReplyToMsgId)
	x.Long(e.RandomId)
	x.Long(e.QueryId)
	x.String(e.Id)
	return x.buf
}
func (e *ReqChannelsToggleInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsToggleInvites)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}
func (e *ReqChannelsExportMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsExportMessageLink)
	x.Bytes(e.Channel.encode())
	x.Int(e.Id)
	return x.buf
}
func (e *ReqChannelsToggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsToggleSignatures)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}
func (e *ReqMessagesHideReportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesHideReportSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *ReqMessagesGetPeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetPeerSettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *ReqChannelsUpdatePinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsUpdatePinnedMessage)
	x.Int(e.Flags)
	x.Bytes(e.Channel.encode())
	x.Int(e.Id)
	return x.buf
}
func (e *ReqAuthResendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authResendCode)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	return x.buf
}
func (e *ReqAuthCancelCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authCancelCode)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	return x.buf
}
func (e *ReqMessagesGetMessageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetMessageEditData)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	return x.buf
}
func (e *ReqMessagesEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesEditMessage)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.FlaggedString(e.Flags, 11, e.Message)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	x.FlaggedVector(e.Flags, 3, toTLslice(e.Entities))
	return x.buf
}
func (e *ReqMessagesEditInlineBotMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesEditInlineBotMessage)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.FlaggedString(e.Flags, 11, e.Message)
	x.FlaggedObject(e.Flags, 2, e.ReplyMarkup)
	x.FlaggedVector(e.Flags, 3, toTLslice(e.Entities))
	return x.buf
}
func (e *ReqMessagesGetBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetBotCallbackAnswer)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgId)
	x.FlaggedStringBytes(e.Flags, 0, e.Data)
	return x.buf
}
func (e *ReqMessagesSetBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSetBotCallbackAnswer)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.FlaggedString(e.Flags, 0, e.Message)
	x.FlaggedString(e.Flags, 2, e.Url)
	x.Int(e.CacheTime)
	return x.buf
}
func (e *ReqContactsGetTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsGetTopPeers)
	x.Int(e.Flags)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqContactsResetTopPeerRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsResetTopPeerRating)
	x.Bytes(e.Category.encode())
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *ReqMessagesGetPeerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetPeerDialogs)
	x.Vector(toTLslice(e.Peers))
	return x.buf
}
func (e *ReqMessagesSaveDraft) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSaveDraft)
	x.Int(e.Flags)
	x.FlaggedInt(e.Flags, 0, e.ReplyToMsgId)
	x.Bytes(e.Peer.encode())
	x.String(e.Message)
	x.FlaggedVector(e.Flags, 3, toTLslice(e.Entities))
	return x.buf
}
func (e *ReqMessagesGetAllDrafts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetAllDrafts)
	return x.buf
}
func (e *ReqAccountSendConfirmPhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountSendConfirmPhoneCode)
	x.Int(e.Flags)
	x.String(e.Hash)
	x.FlaggedObject(e.Flags, 0, e.CurrentNumber)
	return x.buf
}
func (e *ReqAccountConfirmPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountConfirmPhone)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}
func (e *ReqMessagesGetFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetFeaturedStickers)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqMessagesReadFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReadFeaturedStickers)
	x.VectorLong(e.Id)
	return x.buf
}
func (e *ReqMessagesGetRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetRecentStickers)
	x.Int(e.Flags)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqMessagesSaveRecentSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSaveRecentSticker)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}
func (e *ReqMessagesClearRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesClearRecentStickers)
	x.Int(e.Flags)
	return x.buf
}
func (e *ReqMessagesGetArchivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetArchivedStickers)
	x.Int(e.Flags)
	x.Long(e.OffsetId)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqChannelsGetAdminedPublicChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsGetAdminedPublicChannels)
	return x.buf
}
func (e *ReqAuthDropTempAuthKeys) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authDropTempAuthKeys)
	x.VectorLong(e.ExceptAuthKeys)
	return x.buf
}
func (e *ReqMessagesSetGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSetGameScore)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Bytes(e.UserId.encode())
	x.Int(e.Score)
	return x.buf
}
func (e *ReqMessagesSetInlineGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSetInlineGameScore)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.Bytes(e.UserId.encode())
	x.Int(e.Score)
	return x.buf
}
func (e *ReqMessagesGetMaskStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetMaskStickers)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqMessagesGetAttachedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetAttachedStickers)
	x.Bytes(e.Media.encode())
	return x.buf
}
func (e *ReqMessagesGetGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetGameHighScores)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Bytes(e.UserId.encode())
	return x.buf
}
func (e *ReqMessagesGetInlineGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetInlineGameHighScores)
	x.Bytes(e.Id.encode())
	x.Bytes(e.UserId.encode())
	return x.buf
}
func (e *ReqMessagesGetCommonChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetCommonChats)
	x.Bytes(e.UserId.encode())
	x.Int(e.MaxId)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqMessagesGetAllChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetAllChats)
	x.VectorInt(e.ExceptIds)
	return x.buf
}
func (e *ReqHelpSetBotUpdatesStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpSetBotUpdatesStatus)
	x.Int(e.PendingUpdatesCount)
	x.String(e.Message)
	return x.buf
}
func (e *ReqMessagesGetWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetWebPage)
	x.String(e.Url)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqMessagesToggleDialogPin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesToggleDialogPin)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *ReqMessagesReorderPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReorderPinnedDialogs)
	x.Int(e.Flags)
	x.Vector(toTLslice(e.Order))
	return x.buf
}
func (e *ReqMessagesGetPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetPinnedDialogs)
	return x.buf
}
func (e *ReqPhoneRequestCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneRequestCall)
	x.Bytes(e.UserId.encode())
	x.Int(e.RandomId)
	x.StringBytes(e.GAHash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}
func (e *ReqPhoneAcceptCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneAcceptCall)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GB)
	x.Bytes(e.Protocol.encode())
	return x.buf
}
func (e *ReqPhoneDiscardCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneDiscardCall)
	x.Bytes(e.Peer.encode())
	x.Int(e.Duration)
	x.Bytes(e.Reason.encode())
	x.Long(e.ConnectionId)
	return x.buf
}
func (e *ReqPhoneReceivedCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneReceivedCall)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *ReqMessagesReportEncryptedSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesReportEncryptedSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}
func (e *ReqPaymentsGetPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsGetPaymentForm)
	x.Int(e.MsgId)
	return x.buf
}
func (e *ReqPaymentsSendPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsSendPaymentForm)
	x.Int(e.Flags)
	x.Int(e.MsgId)
	x.FlaggedString(e.Flags, 0, e.RequestedInfoId)
	x.FlaggedString(e.Flags, 1, e.ShippingOptionId)
	x.Bytes(e.Credentials.encode())
	return x.buf
}
func (e *ReqAccountGetTmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountGetTmpPassword)
	x.StringBytes(e.PasswordHash)
	x.Int(e.Period)
	return x.buf
}
func (e *ReqMessagesSetBotShippingResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSetBotShippingResults)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.FlaggedString(e.Flags, 0, e.Error)
	x.FlaggedVector(e.Flags, 1, toTLslice(e.ShippingOptions))
	return x.buf
}
func (e *ReqMessagesSetBotPrecheckoutResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSetBotPrecheckoutResults)
	x.Int(e.Flags)
	x.Long(e.QueryId)
	x.FlaggedString(e.Flags, 0, e.Error)
	return x.buf
}
func (e *ReqUploadGetWebFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadGetWebFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqBotsSendCustomRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botsSendCustomRequest)
	x.String(e.CustomMethod)
	x.Bytes(e.Params.encode())
	return x.buf
}
func (e *ReqBotsAnswerWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botsAnswerWebhookJSONQuery)
	x.Long(e.QueryId)
	x.Bytes(e.Data.encode())
	return x.buf
}
func (e *ReqPaymentsGetPaymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsGetPaymentReceipt)
	x.Int(e.MsgId)
	return x.buf
}
func (e *ReqPaymentsValidateRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsValidateRequestedInfo)
	x.Int(e.Flags)
	x.Int(e.MsgId)
	x.Bytes(e.Info.encode())
	return x.buf
}
func (e *ReqPaymentsGetSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsGetSavedInfo)
	return x.buf
}
func (e *ReqPaymentsClearSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentsClearSavedInfo)
	x.Int(e.Flags)
	return x.buf
}
func (e *ReqPhoneGetCallConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneGetCallConfig)
	return x.buf
}
func (e *ReqPhoneConfirmCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneConfirmCall)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GA)
	x.Long(e.KeyFingerprint)
	x.Bytes(e.Protocol.encode())
	return x.buf
}
func (e *ReqPhoneSetCallRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneSetCallRating)
	x.Bytes(e.Peer.encode())
	x.Int(e.Rating)
	x.String(e.Comment)
	return x.buf
}
func (e *ReqPhoneSaveCallDebug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneSaveCallDebug)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Debug.encode())
	return x.buf
}
func (e *ReqUploadGetCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadGetCdnFile)
	x.StringBytes(e.FileToken)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqUploadReuploadCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadReuploadCdnFile)
	x.StringBytes(e.FileToken)
	x.StringBytes(e.RequestToken)
	return x.buf
}
func (e *ReqHelpGetCdnConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_helpGetCdnConfig)
	return x.buf
}
func (e *ReqMessagesUploadMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesUploadMedia)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Media.encode())
	return x.buf
}
func (e *ReqStickersCreateStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickersCreateStickerSet)
	x.Int(e.Flags)
	x.Bytes(e.UserId.encode())
	x.String(e.Title)
	x.String(e.ShortName)
	x.Vector(toTLslice(e.Stickers))
	return x.buf
}
func (e *ReqLangpackGetLangPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpackGetLangPack)
	x.String(e.LangCode)
	return x.buf
}
func (e *ReqLangpackGetStrings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpackGetStrings)
	x.String(e.LangCode)
	x.VectorString(e.Keys)
	return x.buf
}
func (e *ReqLangpackGetDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpackGetDifference)
	x.Int(e.FromVersion)
	return x.buf
}
func (e *ReqLangpackGetLanguages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpackGetLanguages)
	return x.buf
}
func (e *ReqChannelsEditBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsEditBanned)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserId.encode())
	x.Bytes(e.BannedRights.encode())
	return x.buf
}
func (e *ReqChannelsGetAdminLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsGetAdminLog)
	x.Int(e.Flags)
	x.Bytes(e.Channel.encode())
	x.String(e.Q)
	x.FlaggedObject(e.Flags, 0, e.EventsFilter)
	x.FlaggedVector(e.Flags, 1, toTLslice(e.Admins))
	x.Long(e.MaxId)
	x.Long(e.MinId)
	x.Int(e.Limit)
	return x.buf
}
func (e *ReqStickersRemoveStickerFromSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickersRemoveStickerFromSet)
	x.Bytes(e.Sticker.encode())
	return x.buf
}
func (e *ReqStickersChangeStickerPosition) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickersChangeStickerPosition)
	x.Bytes(e.Sticker.encode())
	x.Int(e.Position)
	return x.buf
}
func (e *ReqStickersAddStickerToSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickersAddStickerToSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Sticker.encode())
	return x.buf
}
func (e *ReqMessagesSendScreenshotNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesSendScreenshotNotification)
	x.Bytes(e.Peer.encode())
	x.Int(e.ReplyToMsgId)
	x.Long(e.RandomId)
	return x.buf
}
func (e *ReqUploadGetCdnFileHashes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_uploadGetCdnFileHashes)
	x.StringBytes(e.FileToken)
	x.Int(e.Offset)
	return x.buf
}
func (e *ReqMessagesGetUnreadMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetUnreadMentions)
	x.Bytes(e.Peer.encode())
	x.Int(e.OffsetId)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxId)
	x.Int(e.MinId)
	return x.buf
}
func (e *ReqMessagesFaveSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesFaveSticker)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Unfave.encode())
	return x.buf
}
func (e *ReqChannelsSetStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsSetStickers)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Stickerset.encode())
	return x.buf
}
func (e *ReqContactsResetSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactsResetSaved)
	return x.buf
}
func (e *ReqMessagesGetFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messagesGetFavedStickers)
	x.Int(e.Hash)
	return x.buf
}
func (e *ReqChannelsReadMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelsReadMessageContents)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.Id)
	return x.buf
}

// TL converters to a Type
func toTypeBool(tl TL) *TypeBool {
	switch x := tl.(type) {
	case *PredBoolFalse:
		return &TypeBool{&TypeBool_BoolFalse{x}}
	case *PredBoolTrue:
		return &TypeBool{&TypeBool_BoolTrue{x}}
	}
	return nil
}
func toTypeError(tl TL) *TypeError {
	switch x := tl.(type) {
	case *PredError:
		return &TypeError{x}
	}
	return nil
}
func toTypeNull(tl TL) *TypeNull {
	switch x := tl.(type) {
	case *PredNull:
		return &TypeNull{x}
	}
	return nil
}
func toTypeInputPeer(tl TL) *TypeInputPeer {
	switch x := tl.(type) {
	case *PredInputPeerEmpty:
		return &TypeInputPeer{&TypeInputPeer_InputPeerEmpty{x}}
	case *PredInputPeerSelf:
		return &TypeInputPeer{&TypeInputPeer_InputPeerSelf{x}}
	case *PredInputPeerChat:
		return &TypeInputPeer{&TypeInputPeer_InputPeerChat{x}}
	case *PredInputPeerUser:
		return &TypeInputPeer{&TypeInputPeer_InputPeerUser{x}}
	case *PredInputPeerChannel:
		return &TypeInputPeer{&TypeInputPeer_InputPeerChannel{x}}
	}
	return nil
}
func toTypeInputUser(tl TL) *TypeInputUser {
	switch x := tl.(type) {
	case *PredInputUserEmpty:
		return &TypeInputUser{&TypeInputUser_InputUserEmpty{x}}
	case *PredInputUserSelf:
		return &TypeInputUser{&TypeInputUser_InputUserSelf{x}}
	case *PredInputUser:
		return &TypeInputUser{&TypeInputUser_InputUser{x}}
	}
	return nil
}
func toTypeInputContact(tl TL) *TypeInputContact {
	switch x := tl.(type) {
	case *PredInputPhoneContact:
		return &TypeInputContact{x}
	}
	return nil
}
func toTypeInputFile(tl TL) *TypeInputFile {
	switch x := tl.(type) {
	case *PredInputFile:
		return &TypeInputFile{&TypeInputFile_InputFile{x}}
	case *PredInputFileBig:
		return &TypeInputFile{&TypeInputFile_InputFileBig{x}}
	}
	return nil
}
func toTypeInputMedia(tl TL) *TypeInputMedia {
	switch x := tl.(type) {
	case *PredInputMediaEmpty:
		return &TypeInputMedia{&TypeInputMedia_InputMediaEmpty{x}}
	case *PredInputMediaUploadedPhoto:
		return &TypeInputMedia{&TypeInputMedia_InputMediaUploadedPhoto{x}}
	case *PredInputMediaPhoto:
		return &TypeInputMedia{&TypeInputMedia_InputMediaPhoto{x}}
	case *PredInputMediaGeoPoint:
		return &TypeInputMedia{&TypeInputMedia_InputMediaGeoPoint{x}}
	case *PredInputMediaContact:
		return &TypeInputMedia{&TypeInputMedia_InputMediaContact{x}}
	case *PredInputMediaUploadedDocument:
		return &TypeInputMedia{&TypeInputMedia_InputMediaUploadedDocument{x}}
	case *PredInputMediaDocument:
		return &TypeInputMedia{&TypeInputMedia_InputMediaDocument{x}}
	case *PredInputMediaVenue:
		return &TypeInputMedia{&TypeInputMedia_InputMediaVenue{x}}
	case *PredInputMediaGifExternal:
		return &TypeInputMedia{&TypeInputMedia_InputMediaGifExternal{x}}
	case *PredInputMediaPhotoExternal:
		return &TypeInputMedia{&TypeInputMedia_InputMediaPhotoExternal{x}}
	case *PredInputMediaDocumentExternal:
		return &TypeInputMedia{&TypeInputMedia_InputMediaDocumentExternal{x}}
	case *PredInputMediaGame:
		return &TypeInputMedia{&TypeInputMedia_InputMediaGame{x}}
	case *PredInputMediaInvoice:
		return &TypeInputMedia{&TypeInputMedia_InputMediaInvoice{x}}
	}
	return nil
}
func toTypeInputChatPhoto(tl TL) *TypeInputChatPhoto {
	switch x := tl.(type) {
	case *PredInputChatPhotoEmpty:
		return &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatPhotoEmpty{x}}
	case *PredInputChatUploadedPhoto:
		return &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatUploadedPhoto{x}}
	case *PredInputChatPhoto:
		return &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatPhoto{x}}
	}
	return nil
}
func toTypeInputGeoPoint(tl TL) *TypeInputGeoPoint {
	switch x := tl.(type) {
	case *PredInputGeoPointEmpty:
		return &TypeInputGeoPoint{&TypeInputGeoPoint_InputGeoPointEmpty{x}}
	case *PredInputGeoPoint:
		return &TypeInputGeoPoint{&TypeInputGeoPoint_InputGeoPoint{x}}
	}
	return nil
}
func toTypeInputPhoto(tl TL) *TypeInputPhoto {
	switch x := tl.(type) {
	case *PredInputPhotoEmpty:
		return &TypeInputPhoto{&TypeInputPhoto_InputPhotoEmpty{x}}
	case *PredInputPhoto:
		return &TypeInputPhoto{&TypeInputPhoto_InputPhoto{x}}
	}
	return nil
}
func toTypeInputFileLocation(tl TL) *TypeInputFileLocation {
	switch x := tl.(type) {
	case *PredInputFileLocation:
		return &TypeInputFileLocation{&TypeInputFileLocation_InputFileLocation{x}}
	case *PredInputEncryptedFileLocation:
		return &TypeInputFileLocation{&TypeInputFileLocation_InputEncryptedFileLocation{x}}
	case *PredInputDocumentFileLocation:
		return &TypeInputFileLocation{&TypeInputFileLocation_InputDocumentFileLocation{x}}
	}
	return nil
}
func toTypeInputAppEvent(tl TL) *TypeInputAppEvent {
	switch x := tl.(type) {
	case *PredInputAppEvent:
		return &TypeInputAppEvent{x}
	}
	return nil
}
func toTypePeer(tl TL) *TypePeer {
	switch x := tl.(type) {
	case *PredPeerUser:
		return &TypePeer{&TypePeer_PeerUser{x}}
	case *PredPeerChat:
		return &TypePeer{&TypePeer_PeerChat{x}}
	case *PredPeerChannel:
		return &TypePeer{&TypePeer_PeerChannel{x}}
	}
	return nil
}
func toTypeStorageFileType(tl TL) *TypeStorageFileType {
	switch x := tl.(type) {
	case *PredStorageFileUnknown:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFileUnknown{x}}
	case *PredStorageFileJpeg:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFileJpeg{x}}
	case *PredStorageFileGif:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFileGif{x}}
	case *PredStorageFilePng:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFilePng{x}}
	case *PredStorageFileMp3:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFileMp3{x}}
	case *PredStorageFileMov:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFileMov{x}}
	case *PredStorageFilePartial:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFilePartial{x}}
	case *PredStorageFileMp4:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFileMp4{x}}
	case *PredStorageFileWebp:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFileWebp{x}}
	case *PredStorageFilePdf:
		return &TypeStorageFileType{&TypeStorageFileType_StorageFilePdf{x}}
	}
	return nil
}
func toTypeFileLocation(tl TL) *TypeFileLocation {
	switch x := tl.(type) {
	case *PredFileLocationUnavailable:
		return &TypeFileLocation{&TypeFileLocation_FileLocationUnavailable{x}}
	case *PredFileLocation:
		return &TypeFileLocation{&TypeFileLocation_FileLocation{x}}
	}
	return nil
}
func toTypeUser(tl TL) *TypeUser {
	switch x := tl.(type) {
	case *PredUserEmpty:
		return &TypeUser{&TypeUser_UserEmpty{x}}
	case *PredUser:
		return &TypeUser{&TypeUser_User{x}}
	}
	return nil
}
func toTypeUserProfilePhoto(tl TL) *TypeUserProfilePhoto {
	switch x := tl.(type) {
	case *PredUserProfilePhotoEmpty:
		return &TypeUserProfilePhoto{&TypeUserProfilePhoto_UserProfilePhotoEmpty{x}}
	case *PredUserProfilePhoto:
		return &TypeUserProfilePhoto{&TypeUserProfilePhoto_UserProfilePhoto{x}}
	}
	return nil
}
func toTypeUserStatus(tl TL) *TypeUserStatus {
	switch x := tl.(type) {
	case *PredUserStatusEmpty:
		return &TypeUserStatus{&TypeUserStatus_UserStatusEmpty{x}}
	case *PredUserStatusOnline:
		return &TypeUserStatus{&TypeUserStatus_UserStatusOnline{x}}
	case *PredUserStatusOffline:
		return &TypeUserStatus{&TypeUserStatus_UserStatusOffline{x}}
	case *PredUserStatusRecently:
		return &TypeUserStatus{&TypeUserStatus_UserStatusRecently{x}}
	case *PredUserStatusLastWeek:
		return &TypeUserStatus{&TypeUserStatus_UserStatusLastWeek{x}}
	case *PredUserStatusLastMonth:
		return &TypeUserStatus{&TypeUserStatus_UserStatusLastMonth{x}}
	}
	return nil
}
func toTypeChat(tl TL) *TypeChat {
	switch x := tl.(type) {
	case *PredChatEmpty:
		return &TypeChat{&TypeChat_ChatEmpty{x}}
	case *PredChat:
		return &TypeChat{&TypeChat_Chat{x}}
	case *PredChatForbidden:
		return &TypeChat{&TypeChat_ChatForbidden{x}}
	case *PredChannel:
		return &TypeChat{&TypeChat_Channel{x}}
	case *PredChannelForbidden:
		return &TypeChat{&TypeChat_ChannelForbidden{x}}
	}
	return nil
}
func toTypeChatFull(tl TL) *TypeChatFull {
	switch x := tl.(type) {
	case *PredChatFull:
		return &TypeChatFull{&TypeChatFull_ChatFull{x}}
	case *PredChannelFull:
		return &TypeChatFull{&TypeChatFull_ChannelFull{x}}
	}
	return nil
}
func toTypeChatParticipant(tl TL) *TypeChatParticipant {
	switch x := tl.(type) {
	case *PredChatParticipant:
		return &TypeChatParticipant{&TypeChatParticipant_ChatParticipant{x}}
	case *PredChatParticipantCreator:
		return &TypeChatParticipant{&TypeChatParticipant_ChatParticipantCreator{x}}
	case *PredChatParticipantAdmin:
		return &TypeChatParticipant{&TypeChatParticipant_ChatParticipantAdmin{x}}
	}
	return nil
}
func toTypeChatParticipants(tl TL) *TypeChatParticipants {
	switch x := tl.(type) {
	case *PredChatParticipantsForbidden:
		return &TypeChatParticipants{&TypeChatParticipants_ChatParticipantsForbidden{x}}
	case *PredChatParticipants:
		return &TypeChatParticipants{&TypeChatParticipants_ChatParticipants{x}}
	}
	return nil
}
func toTypeChatPhoto(tl TL) *TypeChatPhoto {
	switch x := tl.(type) {
	case *PredChatPhotoEmpty:
		return &TypeChatPhoto{&TypeChatPhoto_ChatPhotoEmpty{x}}
	case *PredChatPhoto:
		return &TypeChatPhoto{&TypeChatPhoto_ChatPhoto{x}}
	}
	return nil
}
func toTypeMessage(tl TL) *TypeMessage {
	switch x := tl.(type) {
	case *PredMessageEmpty:
		return &TypeMessage{&TypeMessage_MessageEmpty{x}}
	case *PredMessage:
		return &TypeMessage{&TypeMessage_Message{x}}
	case *PredMessageService:
		return &TypeMessage{&TypeMessage_MessageService{x}}
	}
	return nil
}
func toTypeMessageMedia(tl TL) *TypeMessageMedia {
	switch x := tl.(type) {
	case *PredMessageMediaEmpty:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaEmpty{x}}
	case *PredMessageMediaPhoto:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaPhoto{x}}
	case *PredMessageMediaGeo:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaGeo{x}}
	case *PredMessageMediaContact:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaContact{x}}
	case *PredMessageMediaUnsupported:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaUnsupported{x}}
	case *PredMessageMediaDocument:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaDocument{x}}
	case *PredMessageMediaWebPage:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaWebPage{x}}
	case *PredMessageMediaVenue:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaVenue{x}}
	case *PredMessageMediaGame:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaGame{x}}
	case *PredMessageMediaInvoice:
		return &TypeMessageMedia{&TypeMessageMedia_MessageMediaInvoice{x}}
	}
	return nil
}
func toTypeMessageAction(tl TL) *TypeMessageAction {
	switch x := tl.(type) {
	case *PredMessageActionEmpty:
		return &TypeMessageAction{&TypeMessageAction_MessageActionEmpty{x}}
	case *PredMessageActionChatCreate:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChatCreate{x}}
	case *PredMessageActionChatEditTitle:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChatEditTitle{x}}
	case *PredMessageActionChatEditPhoto:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChatEditPhoto{x}}
	case *PredMessageActionChatDeletePhoto:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChatDeletePhoto{x}}
	case *PredMessageActionChatAddUser:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChatAddUser{x}}
	case *PredMessageActionChatDeleteUser:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChatDeleteUser{x}}
	case *PredMessageActionChatJoinedByLink:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChatJoinedByLink{x}}
	case *PredMessageActionChannelCreate:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChannelCreate{x}}
	case *PredMessageActionChatMigrateTo:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChatMigrateTo{x}}
	case *PredMessageActionChannelMigrateFrom:
		return &TypeMessageAction{&TypeMessageAction_MessageActionChannelMigrateFrom{x}}
	case *PredMessageActionPinMessage:
		return &TypeMessageAction{&TypeMessageAction_MessageActionPinMessage{x}}
	case *PredMessageActionHistoryClear:
		return &TypeMessageAction{&TypeMessageAction_MessageActionHistoryClear{x}}
	case *PredMessageActionGameScore:
		return &TypeMessageAction{&TypeMessageAction_MessageActionGameScore{x}}
	case *PredMessageActionPhoneCall:
		return &TypeMessageAction{&TypeMessageAction_MessageActionPhoneCall{x}}
	case *PredMessageActionPaymentSentMe:
		return &TypeMessageAction{&TypeMessageAction_MessageActionPaymentSentMe{x}}
	case *PredMessageActionPaymentSent:
		return &TypeMessageAction{&TypeMessageAction_MessageActionPaymentSent{x}}
	case *PredMessageActionScreenshotTaken:
		return &TypeMessageAction{&TypeMessageAction_MessageActionScreenshotTaken{x}}
	}
	return nil
}
func toTypeDialog(tl TL) *TypeDialog {
	switch x := tl.(type) {
	case *PredDialog:
		return &TypeDialog{x}
	}
	return nil
}
func toTypePhoto(tl TL) *TypePhoto {
	switch x := tl.(type) {
	case *PredPhotoEmpty:
		return &TypePhoto{&TypePhoto_PhotoEmpty{x}}
	case *PredPhoto:
		return &TypePhoto{&TypePhoto_Photo{x}}
	}
	return nil
}
func toTypePhotoSize(tl TL) *TypePhotoSize {
	switch x := tl.(type) {
	case *PredPhotoSizeEmpty:
		return &TypePhotoSize{&TypePhotoSize_PhotoSizeEmpty{x}}
	case *PredPhotoSize:
		return &TypePhotoSize{&TypePhotoSize_PhotoSize{x}}
	case *PredPhotoCachedSize:
		return &TypePhotoSize{&TypePhotoSize_PhotoCachedSize{x}}
	}
	return nil
}
func toTypeGeoPoint(tl TL) *TypeGeoPoint {
	switch x := tl.(type) {
	case *PredGeoPointEmpty:
		return &TypeGeoPoint{&TypeGeoPoint_GeoPointEmpty{x}}
	case *PredGeoPoint:
		return &TypeGeoPoint{&TypeGeoPoint_GeoPoint{x}}
	}
	return nil
}
func toTypeAuthCheckedPhone(tl TL) *TypeAuthCheckedPhone {
	switch x := tl.(type) {
	case *PredAuthCheckedPhone:
		return &TypeAuthCheckedPhone{x}
	}
	return nil
}
func toTypeAuthSentCode(tl TL) *TypeAuthSentCode {
	switch x := tl.(type) {
	case *PredAuthSentCode:
		return &TypeAuthSentCode{x}
	}
	return nil
}
func toTypeAuthAuthorization(tl TL) *TypeAuthAuthorization {
	switch x := tl.(type) {
	case *PredAuthAuthorization:
		return &TypeAuthAuthorization{x}
	}
	return nil
}
func toTypeAuthExportedAuthorization(tl TL) *TypeAuthExportedAuthorization {
	switch x := tl.(type) {
	case *PredAuthExportedAuthorization:
		return &TypeAuthExportedAuthorization{x}
	}
	return nil
}
func toTypeInputNotifyPeer(tl TL) *TypeInputNotifyPeer {
	switch x := tl.(type) {
	case *PredInputNotifyPeer:
		return &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyPeer{x}}
	case *PredInputNotifyUsers:
		return &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyUsers{x}}
	case *PredInputNotifyChats:
		return &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyChats{x}}
	case *PredInputNotifyAll:
		return &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyAll{x}}
	}
	return nil
}
func toTypeInputPeerNotifySettings(tl TL) *TypeInputPeerNotifySettings {
	switch x := tl.(type) {
	case *PredInputPeerNotifySettings:
		return &TypeInputPeerNotifySettings{x}
	}
	return nil
}
func toTypePeerNotifyEvents(tl TL) *TypePeerNotifyEvents {
	switch x := tl.(type) {
	case *PredPeerNotifyEventsEmpty:
		return &TypePeerNotifyEvents{&TypePeerNotifyEvents_PeerNotifyEventsEmpty{x}}
	case *PredPeerNotifyEventsAll:
		return &TypePeerNotifyEvents{&TypePeerNotifyEvents_PeerNotifyEventsAll{x}}
	}
	return nil
}
func toTypePeerNotifySettings(tl TL) *TypePeerNotifySettings {
	switch x := tl.(type) {
	case *PredPeerNotifySettingsEmpty:
		return &TypePeerNotifySettings{&TypePeerNotifySettings_PeerNotifySettingsEmpty{x}}
	case *PredPeerNotifySettings:
		return &TypePeerNotifySettings{&TypePeerNotifySettings_PeerNotifySettings{x}}
	}
	return nil
}
func toTypeWallPaper(tl TL) *TypeWallPaper {
	switch x := tl.(type) {
	case *PredWallPaper:
		return &TypeWallPaper{&TypeWallPaper_WallPaper{x}}
	case *PredWallPaperSolid:
		return &TypeWallPaper{&TypeWallPaper_WallPaperSolid{x}}
	}
	return nil
}
func toTypeUserFull(tl TL) *TypeUserFull {
	switch x := tl.(type) {
	case *PredUserFull:
		return &TypeUserFull{x}
	}
	return nil
}
func toTypeContact(tl TL) *TypeContact {
	switch x := tl.(type) {
	case *PredContact:
		return &TypeContact{x}
	}
	return nil
}
func toTypeImportedContact(tl TL) *TypeImportedContact {
	switch x := tl.(type) {
	case *PredImportedContact:
		return &TypeImportedContact{x}
	}
	return nil
}
func toTypeContactBlocked(tl TL) *TypeContactBlocked {
	switch x := tl.(type) {
	case *PredContactBlocked:
		return &TypeContactBlocked{x}
	}
	return nil
}
func toTypeContactStatus(tl TL) *TypeContactStatus {
	switch x := tl.(type) {
	case *PredContactStatus:
		return &TypeContactStatus{x}
	}
	return nil
}
func toTypeContactsLink(tl TL) *TypeContactsLink {
	switch x := tl.(type) {
	case *PredContactsLink:
		return &TypeContactsLink{x}
	}
	return nil
}
func toTypeContactsContacts(tl TL) *TypeContactsContacts {
	switch x := tl.(type) {
	case *PredContactsContacts:
		return &TypeContactsContacts{&TypeContactsContacts_ContactsContacts{x}}
	case *PredContactsContactsNotModified:
		return &TypeContactsContacts{&TypeContactsContacts_ContactsContactsNotModified{x}}
	}
	return nil
}
func toTypeContactsImportedContacts(tl TL) *TypeContactsImportedContacts {
	switch x := tl.(type) {
	case *PredContactsImportedContacts:
		return &TypeContactsImportedContacts{x}
	}
	return nil
}
func toTypeContactsBlocked(tl TL) *TypeContactsBlocked {
	switch x := tl.(type) {
	case *PredContactsBlocked:
		return &TypeContactsBlocked{&TypeContactsBlocked_ContactsBlocked{x}}
	case *PredContactsBlockedSlice:
		return &TypeContactsBlocked{&TypeContactsBlocked_ContactsBlockedSlice{x}}
	}
	return nil
}
func toTypeContactsFound(tl TL) *TypeContactsFound {
	switch x := tl.(type) {
	case *PredContactsFound:
		return &TypeContactsFound{x}
	}
	return nil
}
func toTypeMessagesDialogs(tl TL) *TypeMessagesDialogs {
	switch x := tl.(type) {
	case *PredMessagesDialogs:
		return &TypeMessagesDialogs{&TypeMessagesDialogs_MessagesDialogs{x}}
	case *PredMessagesDialogsSlice:
		return &TypeMessagesDialogs{&TypeMessagesDialogs_MessagesDialogsSlice{x}}
	}
	return nil
}
func toTypeMessagesMessages(tl TL) *TypeMessagesMessages {
	switch x := tl.(type) {
	case *PredMessagesMessages:
		return &TypeMessagesMessages{&TypeMessagesMessages_MessagesMessages{x}}
	case *PredMessagesMessagesSlice:
		return &TypeMessagesMessages{&TypeMessagesMessages_MessagesMessagesSlice{x}}
	case *PredMessagesChannelMessages:
		return &TypeMessagesMessages{&TypeMessagesMessages_MessagesChannelMessages{x}}
	}
	return nil
}
func toTypeMessagesChats(tl TL) *TypeMessagesChats {
	switch x := tl.(type) {
	case *PredMessagesChats:
		return &TypeMessagesChats{&TypeMessagesChats_MessagesChats{x}}
	case *PredMessagesChatsSlice:
		return &TypeMessagesChats{&TypeMessagesChats_MessagesChatsSlice{x}}
	}
	return nil
}
func toTypeMessagesChatFull(tl TL) *TypeMessagesChatFull {
	switch x := tl.(type) {
	case *PredMessagesChatFull:
		return &TypeMessagesChatFull{x}
	}
	return nil
}
func toTypeMessagesAffectedHistory(tl TL) *TypeMessagesAffectedHistory {
	switch x := tl.(type) {
	case *PredMessagesAffectedHistory:
		return &TypeMessagesAffectedHistory{x}
	}
	return nil
}
func toTypeMessagesFilter(tl TL) *TypeMessagesFilter {
	switch x := tl.(type) {
	case *PredInputMessagesFilterEmpty:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterEmpty{x}}
	case *PredInputMessagesFilterPhotos:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotos{x}}
	case *PredInputMessagesFilterVideo:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterVideo{x}}
	case *PredInputMessagesFilterPhotoVideo:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotoVideo{x}}
	case *PredInputMessagesFilterDocument:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterDocument{x}}
	case *PredInputMessagesFilterPhotoVideoDocuments:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotoVideoDocuments{x}}
	case *PredInputMessagesFilterUrl:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterUrl{x}}
	case *PredInputMessagesFilterGif:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterGif{x}}
	case *PredInputMessagesFilterVoice:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterVoice{x}}
	case *PredInputMessagesFilterMusic:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMusic{x}}
	case *PredInputMessagesFilterChatPhotos:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterChatPhotos{x}}
	case *PredInputMessagesFilterPhoneCalls:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhoneCalls{x}}
	case *PredInputMessagesFilterRoundVoice:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterRoundVoice{x}}
	case *PredInputMessagesFilterRoundVideo:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterRoundVideo{x}}
	case *PredInputMessagesFilterMyMentions:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMyMentions{x}}
	case *PredInputMessagesFilterMyMentionsUnread:
		return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMyMentionsUnread{x}}
	}
	return nil
}
func toTypeUpdate(tl TL) *TypeUpdate {
	switch x := tl.(type) {
	case *PredUpdateNewMessage:
		return &TypeUpdate{&TypeUpdate_UpdateNewMessage{x}}
	case *PredUpdateMessageID:
		return &TypeUpdate{&TypeUpdate_UpdateMessageID{x}}
	case *PredUpdateDeleteMessages:
		return &TypeUpdate{&TypeUpdate_UpdateDeleteMessages{x}}
	case *PredUpdateUserTyping:
		return &TypeUpdate{&TypeUpdate_UpdateUserTyping{x}}
	case *PredUpdateChatUserTyping:
		return &TypeUpdate{&TypeUpdate_UpdateChatUserTyping{x}}
	case *PredUpdateChatParticipants:
		return &TypeUpdate{&TypeUpdate_UpdateChatParticipants{x}}
	case *PredUpdateUserStatus:
		return &TypeUpdate{&TypeUpdate_UpdateUserStatus{x}}
	case *PredUpdateUserName:
		return &TypeUpdate{&TypeUpdate_UpdateUserName{x}}
	case *PredUpdateUserPhoto:
		return &TypeUpdate{&TypeUpdate_UpdateUserPhoto{x}}
	case *PredUpdateContactRegistered:
		return &TypeUpdate{&TypeUpdate_UpdateContactRegistered{x}}
	case *PredUpdateContactLink:
		return &TypeUpdate{&TypeUpdate_UpdateContactLink{x}}
	case *PredUpdateNewEncryptedMessage:
		return &TypeUpdate{&TypeUpdate_UpdateNewEncryptedMessage{x}}
	case *PredUpdateEncryptedChatTyping:
		return &TypeUpdate{&TypeUpdate_UpdateEncryptedChatTyping{x}}
	case *PredUpdateEncryption:
		return &TypeUpdate{&TypeUpdate_UpdateEncryption{x}}
	case *PredUpdateEncryptedMessagesRead:
		return &TypeUpdate{&TypeUpdate_UpdateEncryptedMessagesRead{x}}
	case *PredUpdateChatParticipantAdd:
		return &TypeUpdate{&TypeUpdate_UpdateChatParticipantAdd{x}}
	case *PredUpdateChatParticipantDelete:
		return &TypeUpdate{&TypeUpdate_UpdateChatParticipantDelete{x}}
	case *PredUpdateDcOptions:
		return &TypeUpdate{&TypeUpdate_UpdateDcOptions{x}}
	case *PredUpdateUserBlocked:
		return &TypeUpdate{&TypeUpdate_UpdateUserBlocked{x}}
	case *PredUpdateNotifySettings:
		return &TypeUpdate{&TypeUpdate_UpdateNotifySettings{x}}
	case *PredUpdateServiceNotification:
		return &TypeUpdate{&TypeUpdate_UpdateServiceNotification{x}}
	case *PredUpdatePrivacy:
		return &TypeUpdate{&TypeUpdate_UpdatePrivacy{x}}
	case *PredUpdateUserPhone:
		return &TypeUpdate{&TypeUpdate_UpdateUserPhone{x}}
	case *PredUpdateReadHistoryInbox:
		return &TypeUpdate{&TypeUpdate_UpdateReadHistoryInbox{x}}
	case *PredUpdateReadHistoryOutbox:
		return &TypeUpdate{&TypeUpdate_UpdateReadHistoryOutbox{x}}
	case *PredUpdateWebPage:
		return &TypeUpdate{&TypeUpdate_UpdateWebPage{x}}
	case *PredUpdateReadMessagesContents:
		return &TypeUpdate{&TypeUpdate_UpdateReadMessagesContents{x}}
	case *PredUpdateChannelTooLong:
		return &TypeUpdate{&TypeUpdate_UpdateChannelTooLong{x}}
	case *PredUpdateChannel:
		return &TypeUpdate{&TypeUpdate_UpdateChannel{x}}
	case *PredUpdateNewChannelMessage:
		return &TypeUpdate{&TypeUpdate_UpdateNewChannelMessage{x}}
	case *PredUpdateReadChannelInbox:
		return &TypeUpdate{&TypeUpdate_UpdateReadChannelInbox{x}}
	case *PredUpdateDeleteChannelMessages:
		return &TypeUpdate{&TypeUpdate_UpdateDeleteChannelMessages{x}}
	case *PredUpdateChannelMessageViews:
		return &TypeUpdate{&TypeUpdate_UpdateChannelMessageViews{x}}
	case *PredUpdateChatAdmins:
		return &TypeUpdate{&TypeUpdate_UpdateChatAdmins{x}}
	case *PredUpdateChatParticipantAdmin:
		return &TypeUpdate{&TypeUpdate_UpdateChatParticipantAdmin{x}}
	case *PredUpdateNewStickerSet:
		return &TypeUpdate{&TypeUpdate_UpdateNewStickerSet{x}}
	case *PredUpdateStickerSetsOrder:
		return &TypeUpdate{&TypeUpdate_UpdateStickerSetsOrder{x}}
	case *PredUpdateStickerSets:
		return &TypeUpdate{&TypeUpdate_UpdateStickerSets{x}}
	case *PredUpdateSavedGifs:
		return &TypeUpdate{&TypeUpdate_UpdateSavedGifs{x}}
	case *PredUpdateBotInlineQuery:
		return &TypeUpdate{&TypeUpdate_UpdateBotInlineQuery{x}}
	case *PredUpdateBotInlineSend:
		return &TypeUpdate{&TypeUpdate_UpdateBotInlineSend{x}}
	case *PredUpdateEditChannelMessage:
		return &TypeUpdate{&TypeUpdate_UpdateEditChannelMessage{x}}
	case *PredUpdateChannelPinnedMessage:
		return &TypeUpdate{&TypeUpdate_UpdateChannelPinnedMessage{x}}
	case *PredUpdateBotCallbackQuery:
		return &TypeUpdate{&TypeUpdate_UpdateBotCallbackQuery{x}}
	case *PredUpdateEditMessage:
		return &TypeUpdate{&TypeUpdate_UpdateEditMessage{x}}
	case *PredUpdateInlineBotCallbackQuery:
		return &TypeUpdate{&TypeUpdate_UpdateInlineBotCallbackQuery{x}}
	case *PredUpdateReadChannelOutbox:
		return &TypeUpdate{&TypeUpdate_UpdateReadChannelOutbox{x}}
	case *PredUpdateDraftMessage:
		return &TypeUpdate{&TypeUpdate_UpdateDraftMessage{x}}
	case *PredUpdateReadFeaturedStickers:
		return &TypeUpdate{&TypeUpdate_UpdateReadFeaturedStickers{x}}
	case *PredUpdateRecentStickers:
		return &TypeUpdate{&TypeUpdate_UpdateRecentStickers{x}}
	case *PredUpdateConfig:
		return &TypeUpdate{&TypeUpdate_UpdateConfig{x}}
	case *PredUpdatePtsChanged:
		return &TypeUpdate{&TypeUpdate_UpdatePtsChanged{x}}
	case *PredUpdateChannelWebPage:
		return &TypeUpdate{&TypeUpdate_UpdateChannelWebPage{x}}
	case *PredUpdatePhoneCall:
		return &TypeUpdate{&TypeUpdate_UpdatePhoneCall{x}}
	case *PredUpdateDialogPinned:
		return &TypeUpdate{&TypeUpdate_UpdateDialogPinned{x}}
	case *PredUpdatePinnedDialogs:
		return &TypeUpdate{&TypeUpdate_UpdatePinnedDialogs{x}}
	case *PredUpdateBotWebhookJSON:
		return &TypeUpdate{&TypeUpdate_UpdateBotWebhookJSON{x}}
	case *PredUpdateBotWebhookJSONQuery:
		return &TypeUpdate{&TypeUpdate_UpdateBotWebhookJSONQuery{x}}
	case *PredUpdateBotShippingQuery:
		return &TypeUpdate{&TypeUpdate_UpdateBotShippingQuery{x}}
	case *PredUpdateBotPrecheckoutQuery:
		return &TypeUpdate{&TypeUpdate_UpdateBotPrecheckoutQuery{x}}
	case *PredUpdateLangPackTooLong:
		return &TypeUpdate{&TypeUpdate_UpdateLangPackTooLong{x}}
	case *PredUpdateLangPack:
		return &TypeUpdate{&TypeUpdate_UpdateLangPack{x}}
	case *PredUpdateContactsReset:
		return &TypeUpdate{&TypeUpdate_UpdateContactsReset{x}}
	case *PredUpdateFavedStickers:
		return &TypeUpdate{&TypeUpdate_UpdateFavedStickers{x}}
	case *PredUpdateChannelReadMessagesContents:
		return &TypeUpdate{&TypeUpdate_UpdateChannelReadMessagesContents{x}}
	}
	return nil
}
func toTypeUpdatesState(tl TL) *TypeUpdatesState {
	switch x := tl.(type) {
	case *PredUpdatesState:
		return &TypeUpdatesState{x}
	}
	return nil
}
func toTypeUpdatesDifference(tl TL) *TypeUpdatesDifference {
	switch x := tl.(type) {
	case *PredUpdatesDifferenceEmpty:
		return &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceEmpty{x}}
	case *PredUpdatesDifference:
		return &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifference{x}}
	case *PredUpdatesDifferenceSlice:
		return &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceSlice{x}}
	case *PredUpdatesDifferenceTooLong:
		return &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceTooLong{x}}
	}
	return nil
}
func toTypeUpdates(tl TL) *TypeUpdates {
	switch x := tl.(type) {
	case *PredUpdatesTooLong:
		return &TypeUpdates{&TypeUpdates_UpdatesTooLong{x}}
	case *PredUpdateShortMessage:
		return &TypeUpdates{&TypeUpdates_UpdateShortMessage{x}}
	case *PredUpdateShortChatMessage:
		return &TypeUpdates{&TypeUpdates_UpdateShortChatMessage{x}}
	case *PredUpdateShort:
		return &TypeUpdates{&TypeUpdates_UpdateShort{x}}
	case *PredUpdatesCombined:
		return &TypeUpdates{&TypeUpdates_UpdatesCombined{x}}
	case *PredUpdates:
		return &TypeUpdates{&TypeUpdates_Updates{x}}
	case *PredUpdateShortSentMessage:
		return &TypeUpdates{&TypeUpdates_UpdateShortSentMessage{x}}
	}
	return nil
}
func toTypePhotosPhoto(tl TL) *TypePhotosPhoto {
	switch x := tl.(type) {
	case *PredPhotosPhoto:
		return &TypePhotosPhoto{x}
	}
	return nil
}
func toTypeUploadFile(tl TL) *TypeUploadFile {
	switch x := tl.(type) {
	case *PredUploadFile:
		return &TypeUploadFile{&TypeUploadFile_UploadFile{x}}
	case *PredUploadFileCdnRedirect:
		return &TypeUploadFile{&TypeUploadFile_UploadFileCdnRedirect{x}}
	}
	return nil
}
func toTypeDcOption(tl TL) *TypeDcOption {
	switch x := tl.(type) {
	case *PredDcOption:
		return &TypeDcOption{x}
	}
	return nil
}
func toTypeConfig(tl TL) *TypeConfig {
	switch x := tl.(type) {
	case *PredConfig:
		return &TypeConfig{x}
	}
	return nil
}
func toTypeNearestDc(tl TL) *TypeNearestDc {
	switch x := tl.(type) {
	case *PredNearestDc:
		return &TypeNearestDc{x}
	}
	return nil
}
func toTypeHelpAppUpdate(tl TL) *TypeHelpAppUpdate {
	switch x := tl.(type) {
	case *PredHelpAppUpdate:
		return &TypeHelpAppUpdate{&TypeHelpAppUpdate_HelpAppUpdate{x}}
	case *PredHelpNoAppUpdate:
		return &TypeHelpAppUpdate{&TypeHelpAppUpdate_HelpNoAppUpdate{x}}
	}
	return nil
}
func toTypeHelpInviteText(tl TL) *TypeHelpInviteText {
	switch x := tl.(type) {
	case *PredHelpInviteText:
		return &TypeHelpInviteText{x}
	}
	return nil
}
func toTypeInputPeerNotifyEvents(tl TL) *TypeInputPeerNotifyEvents {
	switch x := tl.(type) {
	case *PredInputPeerNotifyEventsEmpty:
		return &TypeInputPeerNotifyEvents{&TypeInputPeerNotifyEvents_InputPeerNotifyEventsEmpty{x}}
	case *PredInputPeerNotifyEventsAll:
		return &TypeInputPeerNotifyEvents{&TypeInputPeerNotifyEvents_InputPeerNotifyEventsAll{x}}
	}
	return nil
}
func toTypePhotosPhotos(tl TL) *TypePhotosPhotos {
	switch x := tl.(type) {
	case *PredPhotosPhotos:
		return &TypePhotosPhotos{&TypePhotosPhotos_PhotosPhotos{x}}
	case *PredPhotosPhotosSlice:
		return &TypePhotosPhotos{&TypePhotosPhotos_PhotosPhotosSlice{x}}
	}
	return nil
}
func toTypeEncryptedChat(tl TL) *TypeEncryptedChat {
	switch x := tl.(type) {
	case *PredEncryptedChatEmpty:
		return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatEmpty{x}}
	case *PredEncryptedChatWaiting:
		return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatWaiting{x}}
	case *PredEncryptedChatRequested:
		return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatRequested{x}}
	case *PredEncryptedChat:
		return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChat{x}}
	case *PredEncryptedChatDiscarded:
		return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatDiscarded{x}}
	}
	return nil
}
func toTypeInputEncryptedChat(tl TL) *TypeInputEncryptedChat {
	switch x := tl.(type) {
	case *PredInputEncryptedChat:
		return &TypeInputEncryptedChat{x}
	}
	return nil
}
func toTypeEncryptedFile(tl TL) *TypeEncryptedFile {
	switch x := tl.(type) {
	case *PredEncryptedFileEmpty:
		return &TypeEncryptedFile{&TypeEncryptedFile_EncryptedFileEmpty{x}}
	case *PredEncryptedFile:
		return &TypeEncryptedFile{&TypeEncryptedFile_EncryptedFile{x}}
	}
	return nil
}
func toTypeInputEncryptedFile(tl TL) *TypeInputEncryptedFile {
	switch x := tl.(type) {
	case *PredInputEncryptedFileEmpty:
		return &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileEmpty{x}}
	case *PredInputEncryptedFileUploaded:
		return &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileUploaded{x}}
	case *PredInputEncryptedFile:
		return &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFile{x}}
	case *PredInputEncryptedFileBigUploaded:
		return &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileBigUploaded{x}}
	}
	return nil
}
func toTypeEncryptedMessage(tl TL) *TypeEncryptedMessage {
	switch x := tl.(type) {
	case *PredEncryptedMessage:
		return &TypeEncryptedMessage{&TypeEncryptedMessage_EncryptedMessage{x}}
	case *PredEncryptedMessageService:
		return &TypeEncryptedMessage{&TypeEncryptedMessage_EncryptedMessageService{x}}
	}
	return nil
}
func toTypeMessagesDhConfig(tl TL) *TypeMessagesDhConfig {
	switch x := tl.(type) {
	case *PredMessagesDhConfigNotModified:
		return &TypeMessagesDhConfig{&TypeMessagesDhConfig_MessagesDhConfigNotModified{x}}
	case *PredMessagesDhConfig:
		return &TypeMessagesDhConfig{&TypeMessagesDhConfig_MessagesDhConfig{x}}
	}
	return nil
}
func toTypeMessagesSentEncryptedMessage(tl TL) *TypeMessagesSentEncryptedMessage {
	switch x := tl.(type) {
	case *PredMessagesSentEncryptedMessage:
		return &TypeMessagesSentEncryptedMessage{&TypeMessagesSentEncryptedMessage_MessagesSentEncryptedMessage{x}}
	case *PredMessagesSentEncryptedFile:
		return &TypeMessagesSentEncryptedMessage{&TypeMessagesSentEncryptedMessage_MessagesSentEncryptedFile{x}}
	}
	return nil
}
func toTypeInputDocument(tl TL) *TypeInputDocument {
	switch x := tl.(type) {
	case *PredInputDocumentEmpty:
		return &TypeInputDocument{&TypeInputDocument_InputDocumentEmpty{x}}
	case *PredInputDocument:
		return &TypeInputDocument{&TypeInputDocument_InputDocument{x}}
	}
	return nil
}
func toTypeDocument(tl TL) *TypeDocument {
	switch x := tl.(type) {
	case *PredDocumentEmpty:
		return &TypeDocument{&TypeDocument_DocumentEmpty{x}}
	case *PredDocument:
		return &TypeDocument{&TypeDocument_Document{x}}
	}
	return nil
}
func toTypeHelpSupport(tl TL) *TypeHelpSupport {
	switch x := tl.(type) {
	case *PredHelpSupport:
		return &TypeHelpSupport{x}
	}
	return nil
}
func toTypeNotifyPeer(tl TL) *TypeNotifyPeer {
	switch x := tl.(type) {
	case *PredNotifyAll:
		return &TypeNotifyPeer{&TypeNotifyPeer_NotifyAll{x}}
	case *PredNotifyChats:
		return &TypeNotifyPeer{&TypeNotifyPeer_NotifyChats{x}}
	case *PredNotifyPeer:
		return &TypeNotifyPeer{&TypeNotifyPeer_NotifyPeer{x}}
	case *PredNotifyUsers:
		return &TypeNotifyPeer{&TypeNotifyPeer_NotifyUsers{x}}
	}
	return nil
}
func toTypeSendMessageAction(tl TL) *TypeSendMessageAction {
	switch x := tl.(type) {
	case *PredSendMessageTypingAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageTypingAction{x}}
	case *PredSendMessageCancelAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageCancelAction{x}}
	case *PredSendMessageRecordVideoAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordVideoAction{x}}
	case *PredSendMessageUploadVideoAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadVideoAction{x}}
	case *PredSendMessageRecordAudioAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordAudioAction{x}}
	case *PredSendMessageUploadAudioAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadAudioAction{x}}
	case *PredSendMessageUploadPhotoAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadPhotoAction{x}}
	case *PredSendMessageUploadDocumentAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadDocumentAction{x}}
	case *PredSendMessageGeoLocationAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageGeoLocationAction{x}}
	case *PredSendMessageChooseContactAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageChooseContactAction{x}}
	case *PredSendMessageGamePlayAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageGamePlayAction{x}}
	case *PredSendMessageRecordRoundAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordRoundAction{x}}
	case *PredSendMessageUploadRoundAction:
		return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadRoundAction{x}}
	}
	return nil
}
func toTypeInputPrivacyKey(tl TL) *TypeInputPrivacyKey {
	switch x := tl.(type) {
	case *PredInputPrivacyKeyStatusTimestamp:
		return &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyStatusTimestamp{x}}
	case *PredInputPrivacyKeyChatInvite:
		return &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyChatInvite{x}}
	case *PredInputPrivacyKeyPhoneCall:
		return &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyPhoneCall{x}}
	}
	return nil
}
func toTypePrivacyKey(tl TL) *TypePrivacyKey {
	switch x := tl.(type) {
	case *PredPrivacyKeyStatusTimestamp:
		return &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyStatusTimestamp{x}}
	case *PredPrivacyKeyChatInvite:
		return &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyChatInvite{x}}
	case *PredPrivacyKeyPhoneCall:
		return &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyPhoneCall{x}}
	}
	return nil
}
func toTypeInputPrivacyRule(tl TL) *TypeInputPrivacyRule {
	switch x := tl.(type) {
	case *PredInputPrivacyValueAllowContacts:
		return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowContacts{x}}
	case *PredInputPrivacyValueAllowAll:
		return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowAll{x}}
	case *PredInputPrivacyValueAllowUsers:
		return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowUsers{x}}
	case *PredInputPrivacyValueDisallowContacts:
		return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowContacts{x}}
	case *PredInputPrivacyValueDisallowAll:
		return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowAll{x}}
	case *PredInputPrivacyValueDisallowUsers:
		return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowUsers{x}}
	}
	return nil
}
func toTypePrivacyRule(tl TL) *TypePrivacyRule {
	switch x := tl.(type) {
	case *PredPrivacyValueAllowContacts:
		return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowContacts{x}}
	case *PredPrivacyValueAllowAll:
		return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowAll{x}}
	case *PredPrivacyValueAllowUsers:
		return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowUsers{x}}
	case *PredPrivacyValueDisallowContacts:
		return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowContacts{x}}
	case *PredPrivacyValueDisallowAll:
		return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowAll{x}}
	case *PredPrivacyValueDisallowUsers:
		return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowUsers{x}}
	}
	return nil
}
func toTypeAccountPrivacyRules(tl TL) *TypeAccountPrivacyRules {
	switch x := tl.(type) {
	case *PredAccountPrivacyRules:
		return &TypeAccountPrivacyRules{x}
	}
	return nil
}
func toTypeAccountDaysTTL(tl TL) *TypeAccountDaysTTL {
	switch x := tl.(type) {
	case *PredAccountDaysTTL:
		return &TypeAccountDaysTTL{x}
	}
	return nil
}
func toTypeDisabledFeature(tl TL) *TypeDisabledFeature {
	switch x := tl.(type) {
	case *PredDisabledFeature:
		return &TypeDisabledFeature{x}
	}
	return nil
}
func toTypeDocumentAttribute(tl TL) *TypeDocumentAttribute {
	switch x := tl.(type) {
	case *PredDocumentAttributeImageSize:
		return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeImageSize{x}}
	case *PredDocumentAttributeAnimated:
		return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeAnimated{x}}
	case *PredDocumentAttributeSticker:
		return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeSticker{x}}
	case *PredDocumentAttributeVideo:
		return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeVideo{x}}
	case *PredDocumentAttributeAudio:
		return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeAudio{x}}
	case *PredDocumentAttributeFilename:
		return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeFilename{x}}
	case *PredDocumentAttributeHasStickers:
		return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeHasStickers{x}}
	}
	return nil
}
func toTypeMessagesStickers(tl TL) *TypeMessagesStickers {
	switch x := tl.(type) {
	case *PredMessagesStickersNotModified:
		return &TypeMessagesStickers{&TypeMessagesStickers_MessagesStickersNotModified{x}}
	case *PredMessagesStickers:
		return &TypeMessagesStickers{&TypeMessagesStickers_MessagesStickers{x}}
	}
	return nil
}
func toTypeStickerPack(tl TL) *TypeStickerPack {
	switch x := tl.(type) {
	case *PredStickerPack:
		return &TypeStickerPack{x}
	}
	return nil
}
func toTypeMessagesAllStickers(tl TL) *TypeMessagesAllStickers {
	switch x := tl.(type) {
	case *PredMessagesAllStickersNotModified:
		return &TypeMessagesAllStickers{&TypeMessagesAllStickers_MessagesAllStickersNotModified{x}}
	case *PredMessagesAllStickers:
		return &TypeMessagesAllStickers{&TypeMessagesAllStickers_MessagesAllStickers{x}}
	}
	return nil
}
func toTypeAccountPassword(tl TL) *TypeAccountPassword {
	switch x := tl.(type) {
	case *PredAccountNoPassword:
		return &TypeAccountPassword{&TypeAccountPassword_AccountNoPassword{x}}
	case *PredAccountPassword:
		return &TypeAccountPassword{&TypeAccountPassword_AccountPassword{x}}
	}
	return nil
}
func toTypeMessagesAffectedMessages(tl TL) *TypeMessagesAffectedMessages {
	switch x := tl.(type) {
	case *PredMessagesAffectedMessages:
		return &TypeMessagesAffectedMessages{x}
	}
	return nil
}
func toTypeContactLink(tl TL) *TypeContactLink {
	switch x := tl.(type) {
	case *PredContactLinkUnknown:
		return &TypeContactLink{&TypeContactLink_ContactLinkUnknown{x}}
	case *PredContactLinkNone:
		return &TypeContactLink{&TypeContactLink_ContactLinkNone{x}}
	case *PredContactLinkHasPhone:
		return &TypeContactLink{&TypeContactLink_ContactLinkHasPhone{x}}
	case *PredContactLinkContact:
		return &TypeContactLink{&TypeContactLink_ContactLinkContact{x}}
	}
	return nil
}
func toTypeWebPage(tl TL) *TypeWebPage {
	switch x := tl.(type) {
	case *PredWebPageEmpty:
		return &TypeWebPage{&TypeWebPage_WebPageEmpty{x}}
	case *PredWebPagePending:
		return &TypeWebPage{&TypeWebPage_WebPagePending{x}}
	case *PredWebPage:
		return &TypeWebPage{&TypeWebPage_WebPage{x}}
	case *PredWebPageNotModified:
		return &TypeWebPage{&TypeWebPage_WebPageNotModified{x}}
	}
	return nil
}
func toTypeAuthorization(tl TL) *TypeAuthorization {
	switch x := tl.(type) {
	case *PredAuthorization:
		return &TypeAuthorization{x}
	}
	return nil
}
func toTypeAccountAuthorizations(tl TL) *TypeAccountAuthorizations {
	switch x := tl.(type) {
	case *PredAccountAuthorizations:
		return &TypeAccountAuthorizations{x}
	}
	return nil
}
func toTypeAccountPasswordSettings(tl TL) *TypeAccountPasswordSettings {
	switch x := tl.(type) {
	case *PredAccountPasswordSettings:
		return &TypeAccountPasswordSettings{x}
	}
	return nil
}
func toTypeAccountPasswordInputSettings(tl TL) *TypeAccountPasswordInputSettings {
	switch x := tl.(type) {
	case *PredAccountPasswordInputSettings:
		return &TypeAccountPasswordInputSettings{x}
	}
	return nil
}
func toTypeAuthPasswordRecovery(tl TL) *TypeAuthPasswordRecovery {
	switch x := tl.(type) {
	case *PredAuthPasswordRecovery:
		return &TypeAuthPasswordRecovery{x}
	}
	return nil
}
func toTypeReceivedNotifyMessage(tl TL) *TypeReceivedNotifyMessage {
	switch x := tl.(type) {
	case *PredReceivedNotifyMessage:
		return &TypeReceivedNotifyMessage{x}
	}
	return nil
}
func toTypeExportedChatInvite(tl TL) *TypeExportedChatInvite {
	switch x := tl.(type) {
	case *PredChatInviteEmpty:
		return &TypeExportedChatInvite{&TypeExportedChatInvite_ChatInviteEmpty{x}}
	case *PredChatInviteExported:
		return &TypeExportedChatInvite{&TypeExportedChatInvite_ChatInviteExported{x}}
	}
	return nil
}
func toTypeChatInvite(tl TL) *TypeChatInvite {
	switch x := tl.(type) {
	case *PredChatInviteAlready:
		return &TypeChatInvite{&TypeChatInvite_ChatInviteAlready{x}}
	case *PredChatInvite:
		return &TypeChatInvite{&TypeChatInvite_ChatInvite{x}}
	}
	return nil
}
func toTypeInputStickerSet(tl TL) *TypeInputStickerSet {
	switch x := tl.(type) {
	case *PredInputStickerSetEmpty:
		return &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetEmpty{x}}
	case *PredInputStickerSetID:
		return &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetID{x}}
	case *PredInputStickerSetShortName:
		return &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetShortName{x}}
	}
	return nil
}
func toTypeStickerSet(tl TL) *TypeStickerSet {
	switch x := tl.(type) {
	case *PredStickerSet:
		return &TypeStickerSet{x}
	}
	return nil
}
func toTypeMessagesStickerSet(tl TL) *TypeMessagesStickerSet {
	switch x := tl.(type) {
	case *PredMessagesStickerSet:
		return &TypeMessagesStickerSet{x}
	}
	return nil
}
func toTypeBotCommand(tl TL) *TypeBotCommand {
	switch x := tl.(type) {
	case *PredBotCommand:
		return &TypeBotCommand{x}
	}
	return nil
}
func toTypeBotInfo(tl TL) *TypeBotInfo {
	switch x := tl.(type) {
	case *PredBotInfo:
		return &TypeBotInfo{x}
	}
	return nil
}
func toTypeKeyboardButton(tl TL) *TypeKeyboardButton {
	switch x := tl.(type) {
	case *PredKeyboardButton:
		return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButton{x}}
	case *PredKeyboardButtonUrl:
		return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonUrl{x}}
	case *PredKeyboardButtonCallback:
		return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonCallback{x}}
	case *PredKeyboardButtonRequestPhone:
		return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonRequestPhone{x}}
	case *PredKeyboardButtonRequestGeoLocation:
		return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonRequestGeoLocation{x}}
	case *PredKeyboardButtonSwitchInline:
		return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonSwitchInline{x}}
	case *PredKeyboardButtonGame:
		return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonGame{x}}
	case *PredKeyboardButtonBuy:
		return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonBuy{x}}
	}
	return nil
}
func toTypeKeyboardButtonRow(tl TL) *TypeKeyboardButtonRow {
	switch x := tl.(type) {
	case *PredKeyboardButtonRow:
		return &TypeKeyboardButtonRow{x}
	}
	return nil
}
func toTypeReplyMarkup(tl TL) *TypeReplyMarkup {
	switch x := tl.(type) {
	case *PredReplyKeyboardHide:
		return &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardHide{x}}
	case *PredReplyKeyboardForceReply:
		return &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardForceReply{x}}
	case *PredReplyKeyboardMarkup:
		return &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardMarkup{x}}
	case *PredReplyInlineMarkup:
		return &TypeReplyMarkup{&TypeReplyMarkup_ReplyInlineMarkup{x}}
	}
	return nil
}
func toTypeMessageEntity(tl TL) *TypeMessageEntity {
	switch x := tl.(type) {
	case *PredMessageEntityUnknown:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityUnknown{x}}
	case *PredMessageEntityMention:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityMention{x}}
	case *PredMessageEntityHashtag:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityHashtag{x}}
	case *PredMessageEntityBotCommand:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityBotCommand{x}}
	case *PredMessageEntityUrl:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityUrl{x}}
	case *PredMessageEntityEmail:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityEmail{x}}
	case *PredMessageEntityBold:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityBold{x}}
	case *PredMessageEntityItalic:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityItalic{x}}
	case *PredMessageEntityCode:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityCode{x}}
	case *PredMessageEntityPre:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityPre{x}}
	case *PredMessageEntityTextUrl:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityTextUrl{x}}
	case *PredMessageEntityMentionName:
		return &TypeMessageEntity{&TypeMessageEntity_MessageEntityMentionName{x}}
	case *PredInputMessageEntityMentionName:
		return &TypeMessageEntity{&TypeMessageEntity_InputMessageEntityMentionName{x}}
	}
	return nil
}
func toTypeInputChannel(tl TL) *TypeInputChannel {
	switch x := tl.(type) {
	case *PredInputChannelEmpty:
		return &TypeInputChannel{&TypeInputChannel_InputChannelEmpty{x}}
	case *PredInputChannel:
		return &TypeInputChannel{&TypeInputChannel_InputChannel{x}}
	}
	return nil
}
func toTypeContactsResolvedPeer(tl TL) *TypeContactsResolvedPeer {
	switch x := tl.(type) {
	case *PredContactsResolvedPeer:
		return &TypeContactsResolvedPeer{x}
	}
	return nil
}
func toTypeMessageRange(tl TL) *TypeMessageRange {
	switch x := tl.(type) {
	case *PredMessageRange:
		return &TypeMessageRange{x}
	}
	return nil
}
func toTypeUpdatesChannelDifference(tl TL) *TypeUpdatesChannelDifference {
	switch x := tl.(type) {
	case *PredUpdatesChannelDifferenceEmpty:
		return &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifferenceEmpty{x}}
	case *PredUpdatesChannelDifferenceTooLong:
		return &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifferenceTooLong{x}}
	case *PredUpdatesChannelDifference:
		return &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifference{x}}
	}
	return nil
}
func toTypeChannelMessagesFilter(tl TL) *TypeChannelMessagesFilter {
	switch x := tl.(type) {
	case *PredChannelMessagesFilterEmpty:
		return &TypeChannelMessagesFilter{&TypeChannelMessagesFilter_ChannelMessagesFilterEmpty{x}}
	case *PredChannelMessagesFilter:
		return &TypeChannelMessagesFilter{&TypeChannelMessagesFilter_ChannelMessagesFilter{x}}
	}
	return nil
}
func toTypeChannelParticipant(tl TL) *TypeChannelParticipant {
	switch x := tl.(type) {
	case *PredChannelParticipant:
		return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipant{x}}
	case *PredChannelParticipantSelf:
		return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantSelf{x}}
	case *PredChannelParticipantCreator:
		return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantCreator{x}}
	case *PredChannelParticipantAdmin:
		return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantAdmin{x}}
	case *PredChannelParticipantBanned:
		return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantBanned{x}}
	}
	return nil
}
func toTypeChannelParticipantsFilter(tl TL) *TypeChannelParticipantsFilter {
	switch x := tl.(type) {
	case *PredChannelParticipantsRecent:
		return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsRecent{x}}
	case *PredChannelParticipantsAdmins:
		return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsAdmins{x}}
	case *PredChannelParticipantsKicked:
		return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsKicked{x}}
	case *PredChannelParticipantsBots:
		return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsBots{x}}
	case *PredChannelParticipantsBanned:
		return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsBanned{x}}
	case *PredChannelParticipantsSearch:
		return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsSearch{x}}
	}
	return nil
}
func toTypeChannelsChannelParticipants(tl TL) *TypeChannelsChannelParticipants {
	switch x := tl.(type) {
	case *PredChannelsChannelParticipants:
		return &TypeChannelsChannelParticipants{x}
	}
	return nil
}
func toTypeChannelsChannelParticipant(tl TL) *TypeChannelsChannelParticipant {
	switch x := tl.(type) {
	case *PredChannelsChannelParticipant:
		return &TypeChannelsChannelParticipant{x}
	}
	return nil
}
func toTypeTrue(tl TL) *TypeTrue {
	switch x := tl.(type) {
	case *PredTrue:
		return &TypeTrue{x}
	}
	return nil
}
func toTypeReportReason(tl TL) *TypeReportReason {
	switch x := tl.(type) {
	case *PredInputReportReasonSpam:
		return &TypeReportReason{&TypeReportReason_InputReportReasonSpam{x}}
	case *PredInputReportReasonViolence:
		return &TypeReportReason{&TypeReportReason_InputReportReasonViolence{x}}
	case *PredInputReportReasonPornography:
		return &TypeReportReason{&TypeReportReason_InputReportReasonPornography{x}}
	case *PredInputReportReasonOther:
		return &TypeReportReason{&TypeReportReason_InputReportReasonOther{x}}
	}
	return nil
}
func toTypeHelpTermsOfService(tl TL) *TypeHelpTermsOfService {
	switch x := tl.(type) {
	case *PredHelpTermsOfService:
		return &TypeHelpTermsOfService{x}
	}
	return nil
}
func toTypeFoundGif(tl TL) *TypeFoundGif {
	switch x := tl.(type) {
	case *PredFoundGif:
		return &TypeFoundGif{&TypeFoundGif_FoundGif{x}}
	case *PredFoundGifCached:
		return &TypeFoundGif{&TypeFoundGif_FoundGifCached{x}}
	}
	return nil
}
func toTypeMessagesFoundGifs(tl TL) *TypeMessagesFoundGifs {
	switch x := tl.(type) {
	case *PredMessagesFoundGifs:
		return &TypeMessagesFoundGifs{x}
	}
	return nil
}
func toTypeMessagesSavedGifs(tl TL) *TypeMessagesSavedGifs {
	switch x := tl.(type) {
	case *PredMessagesSavedGifsNotModified:
		return &TypeMessagesSavedGifs{&TypeMessagesSavedGifs_MessagesSavedGifsNotModified{x}}
	case *PredMessagesSavedGifs:
		return &TypeMessagesSavedGifs{&TypeMessagesSavedGifs_MessagesSavedGifs{x}}
	}
	return nil
}
func toTypeInputBotInlineMessage(tl TL) *TypeInputBotInlineMessage {
	switch x := tl.(type) {
	case *PredInputBotInlineMessageMediaAuto:
		return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaAuto{x}}
	case *PredInputBotInlineMessageText:
		return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageText{x}}
	case *PredInputBotInlineMessageMediaGeo:
		return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaGeo{x}}
	case *PredInputBotInlineMessageMediaVenue:
		return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaVenue{x}}
	case *PredInputBotInlineMessageMediaContact:
		return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaContact{x}}
	case *PredInputBotInlineMessageGame:
		return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageGame{x}}
	}
	return nil
}
func toTypeInputBotInlineResult(tl TL) *TypeInputBotInlineResult {
	switch x := tl.(type) {
	case *PredInputBotInlineResult:
		return &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResult{x}}
	case *PredInputBotInlineResultPhoto:
		return &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultPhoto{x}}
	case *PredInputBotInlineResultDocument:
		return &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultDocument{x}}
	case *PredInputBotInlineResultGame:
		return &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultGame{x}}
	}
	return nil
}
func toTypeBotInlineMessage(tl TL) *TypeBotInlineMessage {
	switch x := tl.(type) {
	case *PredBotInlineMessageMediaAuto:
		return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaAuto{x}}
	case *PredBotInlineMessageText:
		return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageText{x}}
	case *PredBotInlineMessageMediaGeo:
		return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaGeo{x}}
	case *PredBotInlineMessageMediaVenue:
		return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaVenue{x}}
	case *PredBotInlineMessageMediaContact:
		return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaContact{x}}
	}
	return nil
}
func toTypeBotInlineResult(tl TL) *TypeBotInlineResult {
	switch x := tl.(type) {
	case *PredBotInlineResult:
		return &TypeBotInlineResult{&TypeBotInlineResult_BotInlineResult{x}}
	case *PredBotInlineMediaResult:
		return &TypeBotInlineResult{&TypeBotInlineResult_BotInlineMediaResult{x}}
	}
	return nil
}
func toTypeMessagesBotResults(tl TL) *TypeMessagesBotResults {
	switch x := tl.(type) {
	case *PredMessagesBotResults:
		return &TypeMessagesBotResults{x}
	}
	return nil
}
func toTypeExportedMessageLink(tl TL) *TypeExportedMessageLink {
	switch x := tl.(type) {
	case *PredExportedMessageLink:
		return &TypeExportedMessageLink{x}
	}
	return nil
}
func toTypeMessageFwdHeader(tl TL) *TypeMessageFwdHeader {
	switch x := tl.(type) {
	case *PredMessageFwdHeader:
		return &TypeMessageFwdHeader{x}
	}
	return nil
}
func toTypePeerSettings(tl TL) *TypePeerSettings {
	switch x := tl.(type) {
	case *PredPeerSettings:
		return &TypePeerSettings{x}
	}
	return nil
}
func toTypeAuthCodeType(tl TL) *TypeAuthCodeType {
	switch x := tl.(type) {
	case *PredAuthCodeTypeSms:
		return &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeSms{x}}
	case *PredAuthCodeTypeCall:
		return &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeCall{x}}
	case *PredAuthCodeTypeFlashCall:
		return &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeFlashCall{x}}
	}
	return nil
}
func toTypeAuthSentCodeType(tl TL) *TypeAuthSentCodeType {
	switch x := tl.(type) {
	case *PredAuthSentCodeTypeApp:
		return &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeApp{x}}
	case *PredAuthSentCodeTypeSms:
		return &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeSms{x}}
	case *PredAuthSentCodeTypeCall:
		return &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeCall{x}}
	case *PredAuthSentCodeTypeFlashCall:
		return &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeFlashCall{x}}
	}
	return nil
}
func toTypeMessagesBotCallbackAnswer(tl TL) *TypeMessagesBotCallbackAnswer {
	switch x := tl.(type) {
	case *PredMessagesBotCallbackAnswer:
		return &TypeMessagesBotCallbackAnswer{x}
	}
	return nil
}
func toTypeMessagesMessageEditData(tl TL) *TypeMessagesMessageEditData {
	switch x := tl.(type) {
	case *PredMessagesMessageEditData:
		return &TypeMessagesMessageEditData{x}
	}
	return nil
}
func toTypeInputBotInlineMessageID(tl TL) *TypeInputBotInlineMessageID {
	switch x := tl.(type) {
	case *PredInputBotInlineMessageID:
		return &TypeInputBotInlineMessageID{x}
	}
	return nil
}
func toTypeInlineBotSwitchPM(tl TL) *TypeInlineBotSwitchPM {
	switch x := tl.(type) {
	case *PredInlineBotSwitchPM:
		return &TypeInlineBotSwitchPM{x}
	}
	return nil
}
func toTypeMessagesPeerDialogs(tl TL) *TypeMessagesPeerDialogs {
	switch x := tl.(type) {
	case *PredMessagesPeerDialogs:
		return &TypeMessagesPeerDialogs{x}
	}
	return nil
}
func toTypeTopPeer(tl TL) *TypeTopPeer {
	switch x := tl.(type) {
	case *PredTopPeer:
		return &TypeTopPeer{x}
	}
	return nil
}
func toTypeTopPeerCategory(tl TL) *TypeTopPeerCategory {
	switch x := tl.(type) {
	case *PredTopPeerCategoryBotsPM:
		return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryBotsPM{x}}
	case *PredTopPeerCategoryBotsInline:
		return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryBotsInline{x}}
	case *PredTopPeerCategoryCorrespondents:
		return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryCorrespondents{x}}
	case *PredTopPeerCategoryGroups:
		return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryGroups{x}}
	case *PredTopPeerCategoryChannels:
		return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryChannels{x}}
	case *PredTopPeerCategoryPhoneCalls:
		return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryPhoneCalls{x}}
	}
	return nil
}
func toTypeTopPeerCategoryPeers(tl TL) *TypeTopPeerCategoryPeers {
	switch x := tl.(type) {
	case *PredTopPeerCategoryPeers:
		return &TypeTopPeerCategoryPeers{x}
	}
	return nil
}
func toTypeContactsTopPeers(tl TL) *TypeContactsTopPeers {
	switch x := tl.(type) {
	case *PredContactsTopPeersNotModified:
		return &TypeContactsTopPeers{&TypeContactsTopPeers_ContactsTopPeersNotModified{x}}
	case *PredContactsTopPeers:
		return &TypeContactsTopPeers{&TypeContactsTopPeers_ContactsTopPeers{x}}
	}
	return nil
}
func toTypeDraftMessage(tl TL) *TypeDraftMessage {
	switch x := tl.(type) {
	case *PredDraftMessageEmpty:
		return &TypeDraftMessage{&TypeDraftMessage_DraftMessageEmpty{x}}
	case *PredDraftMessage:
		return &TypeDraftMessage{&TypeDraftMessage_DraftMessage{x}}
	}
	return nil
}
func toTypeMessagesFeaturedStickers(tl TL) *TypeMessagesFeaturedStickers {
	switch x := tl.(type) {
	case *PredMessagesFeaturedStickersNotModified:
		return &TypeMessagesFeaturedStickers{&TypeMessagesFeaturedStickers_MessagesFeaturedStickersNotModified{x}}
	case *PredMessagesFeaturedStickers:
		return &TypeMessagesFeaturedStickers{&TypeMessagesFeaturedStickers_MessagesFeaturedStickers{x}}
	}
	return nil
}
func toTypeMessagesRecentStickers(tl TL) *TypeMessagesRecentStickers {
	switch x := tl.(type) {
	case *PredMessagesRecentStickersNotModified:
		return &TypeMessagesRecentStickers{&TypeMessagesRecentStickers_MessagesRecentStickersNotModified{x}}
	case *PredMessagesRecentStickers:
		return &TypeMessagesRecentStickers{&TypeMessagesRecentStickers_MessagesRecentStickers{x}}
	}
	return nil
}
func toTypeMessagesArchivedStickers(tl TL) *TypeMessagesArchivedStickers {
	switch x := tl.(type) {
	case *PredMessagesArchivedStickers:
		return &TypeMessagesArchivedStickers{x}
	}
	return nil
}
func toTypeMessagesStickerSetInstallResult(tl TL) *TypeMessagesStickerSetInstallResult {
	switch x := tl.(type) {
	case *PredMessagesStickerSetInstallResultSuccess:
		return &TypeMessagesStickerSetInstallResult{&TypeMessagesStickerSetInstallResult_MessagesStickerSetInstallResultSuccess{x}}
	case *PredMessagesStickerSetInstallResultArchive:
		return &TypeMessagesStickerSetInstallResult{&TypeMessagesStickerSetInstallResult_MessagesStickerSetInstallResultArchive{x}}
	}
	return nil
}
func toTypeStickerSetCovered(tl TL) *TypeStickerSetCovered {
	switch x := tl.(type) {
	case *PredStickerSetCovered:
		return &TypeStickerSetCovered{&TypeStickerSetCovered_StickerSetCovered{x}}
	case *PredStickerSetMultiCovered:
		return &TypeStickerSetCovered{&TypeStickerSetCovered_StickerSetMultiCovered{x}}
	}
	return nil
}
func toTypeMaskCoords(tl TL) *TypeMaskCoords {
	switch x := tl.(type) {
	case *PredMaskCoords:
		return &TypeMaskCoords{x}
	}
	return nil
}
func toTypeInputStickeredMedia(tl TL) *TypeInputStickeredMedia {
	switch x := tl.(type) {
	case *PredInputStickeredMediaPhoto:
		return &TypeInputStickeredMedia{&TypeInputStickeredMedia_InputStickeredMediaPhoto{x}}
	case *PredInputStickeredMediaDocument:
		return &TypeInputStickeredMedia{&TypeInputStickeredMedia_InputStickeredMediaDocument{x}}
	}
	return nil
}
func toTypeGame(tl TL) *TypeGame {
	switch x := tl.(type) {
	case *PredGame:
		return &TypeGame{x}
	}
	return nil
}
func toTypeInputGame(tl TL) *TypeInputGame {
	switch x := tl.(type) {
	case *PredInputGameID:
		return &TypeInputGame{&TypeInputGame_InputGameID{x}}
	case *PredInputGameShortName:
		return &TypeInputGame{&TypeInputGame_InputGameShortName{x}}
	}
	return nil
}
func toTypeHighScore(tl TL) *TypeHighScore {
	switch x := tl.(type) {
	case *PredHighScore:
		return &TypeHighScore{x}
	}
	return nil
}
func toTypeMessagesHighScores(tl TL) *TypeMessagesHighScores {
	switch x := tl.(type) {
	case *PredMessagesHighScores:
		return &TypeMessagesHighScores{x}
	}
	return nil
}
func toTypeRichText(tl TL) *TypeRichText {
	switch x := tl.(type) {
	case *PredTextEmpty:
		return &TypeRichText{&TypeRichText_TextEmpty{x}}
	case *PredTextPlain:
		return &TypeRichText{&TypeRichText_TextPlain{x}}
	case *PredTextBold:
		return &TypeRichText{&TypeRichText_TextBold{x}}
	case *PredTextItalic:
		return &TypeRichText{&TypeRichText_TextItalic{x}}
	case *PredTextUnderline:
		return &TypeRichText{&TypeRichText_TextUnderline{x}}
	case *PredTextStrike:
		return &TypeRichText{&TypeRichText_TextStrike{x}}
	case *PredTextFixed:
		return &TypeRichText{&TypeRichText_TextFixed{x}}
	case *PredTextUrl:
		return &TypeRichText{&TypeRichText_TextUrl{x}}
	case *PredTextEmail:
		return &TypeRichText{&TypeRichText_TextEmail{x}}
	case *PredTextConcat:
		return &TypeRichText{&TypeRichText_TextConcat{x}}
	}
	return nil
}
func toTypePageBlock(tl TL) *TypePageBlock {
	switch x := tl.(type) {
	case *PredPageBlockTitle:
		return &TypePageBlock{&TypePageBlock_PageBlockTitle{x}}
	case *PredPageBlockSubtitle:
		return &TypePageBlock{&TypePageBlock_PageBlockSubtitle{x}}
	case *PredPageBlockAuthorDate:
		return &TypePageBlock{&TypePageBlock_PageBlockAuthorDate{x}}
	case *PredPageBlockHeader:
		return &TypePageBlock{&TypePageBlock_PageBlockHeader{x}}
	case *PredPageBlockSubheader:
		return &TypePageBlock{&TypePageBlock_PageBlockSubheader{x}}
	case *PredPageBlockParagraph:
		return &TypePageBlock{&TypePageBlock_PageBlockParagraph{x}}
	case *PredPageBlockPreformatted:
		return &TypePageBlock{&TypePageBlock_PageBlockPreformatted{x}}
	case *PredPageBlockFooter:
		return &TypePageBlock{&TypePageBlock_PageBlockFooter{x}}
	case *PredPageBlockDivider:
		return &TypePageBlock{&TypePageBlock_PageBlockDivider{x}}
	case *PredPageBlockList:
		return &TypePageBlock{&TypePageBlock_PageBlockList{x}}
	case *PredPageBlockBlockquote:
		return &TypePageBlock{&TypePageBlock_PageBlockBlockquote{x}}
	case *PredPageBlockPullquote:
		return &TypePageBlock{&TypePageBlock_PageBlockPullquote{x}}
	case *PredPageBlockPhoto:
		return &TypePageBlock{&TypePageBlock_PageBlockPhoto{x}}
	case *PredPageBlockVideo:
		return &TypePageBlock{&TypePageBlock_PageBlockVideo{x}}
	case *PredPageBlockCover:
		return &TypePageBlock{&TypePageBlock_PageBlockCover{x}}
	case *PredPageBlockEmbed:
		return &TypePageBlock{&TypePageBlock_PageBlockEmbed{x}}
	case *PredPageBlockEmbedPost:
		return &TypePageBlock{&TypePageBlock_PageBlockEmbedPost{x}}
	case *PredPageBlockSlideshow:
		return &TypePageBlock{&TypePageBlock_PageBlockSlideshow{x}}
	case *PredPageBlockUnsupported:
		return &TypePageBlock{&TypePageBlock_PageBlockUnsupported{x}}
	case *PredPageBlockAnchor:
		return &TypePageBlock{&TypePageBlock_PageBlockAnchor{x}}
	case *PredPageBlockCollage:
		return &TypePageBlock{&TypePageBlock_PageBlockCollage{x}}
	case *PredPageBlockChannel:
		return &TypePageBlock{&TypePageBlock_PageBlockChannel{x}}
	case *PredPageBlockAudio:
		return &TypePageBlock{&TypePageBlock_PageBlockAudio{x}}
	}
	return nil
}
func toTypePage(tl TL) *TypePage {
	switch x := tl.(type) {
	case *PredPagePart:
		return &TypePage{&TypePage_PagePart{x}}
	case *PredPageFull:
		return &TypePage{&TypePage_PageFull{x}}
	}
	return nil
}
func toTypeInputPhoneCall(tl TL) *TypeInputPhoneCall {
	switch x := tl.(type) {
	case *PredInputPhoneCall:
		return &TypeInputPhoneCall{x}
	}
	return nil
}
func toTypePhoneCall(tl TL) *TypePhoneCall {
	switch x := tl.(type) {
	case *PredPhoneCallEmpty:
		return &TypePhoneCall{&TypePhoneCall_PhoneCallEmpty{x}}
	case *PredPhoneCallWaiting:
		return &TypePhoneCall{&TypePhoneCall_PhoneCallWaiting{x}}
	case *PredPhoneCallRequested:
		return &TypePhoneCall{&TypePhoneCall_PhoneCallRequested{x}}
	case *PredPhoneCall:
		return &TypePhoneCall{&TypePhoneCall_PhoneCall{x}}
	case *PredPhoneCallDiscarded:
		return &TypePhoneCall{&TypePhoneCall_PhoneCallDiscarded{x}}
	case *PredPhoneCallAccepted:
		return &TypePhoneCall{&TypePhoneCall_PhoneCallAccepted{x}}
	}
	return nil
}
func toTypePhoneConnection(tl TL) *TypePhoneConnection {
	switch x := tl.(type) {
	case *PredPhoneConnection:
		return &TypePhoneConnection{x}
	}
	return nil
}
func toTypePhoneCallProtocol(tl TL) *TypePhoneCallProtocol {
	switch x := tl.(type) {
	case *PredPhoneCallProtocol:
		return &TypePhoneCallProtocol{x}
	}
	return nil
}
func toTypePhonePhoneCall(tl TL) *TypePhonePhoneCall {
	switch x := tl.(type) {
	case *PredPhonePhoneCall:
		return &TypePhonePhoneCall{x}
	}
	return nil
}
func toTypePhoneCallDiscardReason(tl TL) *TypePhoneCallDiscardReason {
	switch x := tl.(type) {
	case *PredPhoneCallDiscardReasonMissed:
		return &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonMissed{x}}
	case *PredPhoneCallDiscardReasonDisconnect:
		return &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonDisconnect{x}}
	case *PredPhoneCallDiscardReasonHangup:
		return &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonHangup{x}}
	case *PredPhoneCallDiscardReasonBusy:
		return &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonBusy{x}}
	}
	return nil
}
func toTypeInvoice(tl TL) *TypeInvoice {
	switch x := tl.(type) {
	case *PredInvoice:
		return &TypeInvoice{x}
	}
	return nil
}
func toTypePaymentsPaymentForm(tl TL) *TypePaymentsPaymentForm {
	switch x := tl.(type) {
	case *PredPaymentsPaymentForm:
		return &TypePaymentsPaymentForm{x}
	}
	return nil
}
func toTypePostAddress(tl TL) *TypePostAddress {
	switch x := tl.(type) {
	case *PredPostAddress:
		return &TypePostAddress{x}
	}
	return nil
}
func toTypePaymentRequestedInfo(tl TL) *TypePaymentRequestedInfo {
	switch x := tl.(type) {
	case *PredPaymentRequestedInfo:
		return &TypePaymentRequestedInfo{x}
	}
	return nil
}
func toTypeDataJSON(tl TL) *TypeDataJSON {
	switch x := tl.(type) {
	case *PredDataJSON:
		return &TypeDataJSON{x}
	}
	return nil
}
func toTypeLabeledPrice(tl TL) *TypeLabeledPrice {
	switch x := tl.(type) {
	case *PredLabeledPrice:
		return &TypeLabeledPrice{x}
	}
	return nil
}
func toTypePaymentCharge(tl TL) *TypePaymentCharge {
	switch x := tl.(type) {
	case *PredPaymentCharge:
		return &TypePaymentCharge{x}
	}
	return nil
}
func toTypePaymentSavedCredentials(tl TL) *TypePaymentSavedCredentials {
	switch x := tl.(type) {
	case *PredPaymentSavedCredentialsCard:
		return &TypePaymentSavedCredentials{x}
	}
	return nil
}
func toTypeWebDocument(tl TL) *TypeWebDocument {
	switch x := tl.(type) {
	case *PredWebDocument:
		return &TypeWebDocument{x}
	}
	return nil
}
func toTypeInputWebDocument(tl TL) *TypeInputWebDocument {
	switch x := tl.(type) {
	case *PredInputWebDocument:
		return &TypeInputWebDocument{x}
	}
	return nil
}
func toTypeInputWebFileLocation(tl TL) *TypeInputWebFileLocation {
	switch x := tl.(type) {
	case *PredInputWebFileLocation:
		return &TypeInputWebFileLocation{x}
	}
	return nil
}
func toTypeUploadWebFile(tl TL) *TypeUploadWebFile {
	switch x := tl.(type) {
	case *PredUploadWebFile:
		return &TypeUploadWebFile{x}
	}
	return nil
}
func toTypePaymentsValidatedRequestedInfo(tl TL) *TypePaymentsValidatedRequestedInfo {
	switch x := tl.(type) {
	case *PredPaymentsValidatedRequestedInfo:
		return &TypePaymentsValidatedRequestedInfo{x}
	}
	return nil
}
func toTypePaymentsPaymentResult(tl TL) *TypePaymentsPaymentResult {
	switch x := tl.(type) {
	case *PredPaymentsPaymentResult:
		return &TypePaymentsPaymentResult{&TypePaymentsPaymentResult_PaymentsPaymentResult{x}}
	case *PredPaymentsPaymentVerficationNeeded:
		return &TypePaymentsPaymentResult{&TypePaymentsPaymentResult_PaymentsPaymentVerficationNeeded{x}}
	}
	return nil
}
func toTypePaymentsPaymentReceipt(tl TL) *TypePaymentsPaymentReceipt {
	switch x := tl.(type) {
	case *PredPaymentsPaymentReceipt:
		return &TypePaymentsPaymentReceipt{x}
	}
	return nil
}
func toTypePaymentsSavedInfo(tl TL) *TypePaymentsSavedInfo {
	switch x := tl.(type) {
	case *PredPaymentsSavedInfo:
		return &TypePaymentsSavedInfo{x}
	}
	return nil
}
func toTypeInputPaymentCredentials(tl TL) *TypeInputPaymentCredentials {
	switch x := tl.(type) {
	case *PredInputPaymentCredentialsSaved:
		return &TypeInputPaymentCredentials{&TypeInputPaymentCredentials_InputPaymentCredentialsSaved{x}}
	case *PredInputPaymentCredentials:
		return &TypeInputPaymentCredentials{&TypeInputPaymentCredentials_InputPaymentCredentials{x}}
	}
	return nil
}
func toTypeAccountTmpPassword(tl TL) *TypeAccountTmpPassword {
	switch x := tl.(type) {
	case *PredAccountTmpPassword:
		return &TypeAccountTmpPassword{x}
	}
	return nil
}
func toTypeShippingOption(tl TL) *TypeShippingOption {
	switch x := tl.(type) {
	case *PredShippingOption:
		return &TypeShippingOption{x}
	}
	return nil
}
func toTypeUploadCdnFile(tl TL) *TypeUploadCdnFile {
	switch x := tl.(type) {
	case *PredUploadCdnFileReuploadNeeded:
		return &TypeUploadCdnFile{&TypeUploadCdnFile_UploadCdnFileReuploadNeeded{x}}
	case *PredUploadCdnFile:
		return &TypeUploadCdnFile{&TypeUploadCdnFile_UploadCdnFile{x}}
	}
	return nil
}
func toTypeCdnPublicKey(tl TL) *TypeCdnPublicKey {
	switch x := tl.(type) {
	case *PredCdnPublicKey:
		return &TypeCdnPublicKey{x}
	}
	return nil
}
func toTypeCdnConfig(tl TL) *TypeCdnConfig {
	switch x := tl.(type) {
	case *PredCdnConfig:
		return &TypeCdnConfig{x}
	}
	return nil
}
func toTypeInputStickerSetItem(tl TL) *TypeInputStickerSetItem {
	switch x := tl.(type) {
	case *PredInputStickerSetItem:
		return &TypeInputStickerSetItem{x}
	}
	return nil
}
func toTypeLangPackString(tl TL) *TypeLangPackString {
	switch x := tl.(type) {
	case *PredLangPackString:
		return &TypeLangPackString{&TypeLangPackString_LangPackString{x}}
	case *PredLangPackStringPluralized:
		return &TypeLangPackString{&TypeLangPackString_LangPackStringPluralized{x}}
	case *PredLangPackStringDeleted:
		return &TypeLangPackString{&TypeLangPackString_LangPackStringDeleted{x}}
	}
	return nil
}
func toTypeLangPackDifference(tl TL) *TypeLangPackDifference {
	switch x := tl.(type) {
	case *PredLangPackDifference:
		return &TypeLangPackDifference{x}
	}
	return nil
}
func toTypeLangPackLanguage(tl TL) *TypeLangPackLanguage {
	switch x := tl.(type) {
	case *PredLangPackLanguage:
		return &TypeLangPackLanguage{x}
	}
	return nil
}
func toTypeChannelAdminRights(tl TL) *TypeChannelAdminRights {
	switch x := tl.(type) {
	case *PredChannelAdminRights:
		return &TypeChannelAdminRights{x}
	}
	return nil
}
func toTypeChannelBannedRights(tl TL) *TypeChannelBannedRights {
	switch x := tl.(type) {
	case *PredChannelBannedRights:
		return &TypeChannelBannedRights{x}
	}
	return nil
}
func toTypeChannelAdminLogEventAction(tl TL) *TypeChannelAdminLogEventAction {
	switch x := tl.(type) {
	case *PredChannelAdminLogEventActionChangeTitle:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeTitle{x}}
	case *PredChannelAdminLogEventActionChangeAbout:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeAbout{x}}
	case *PredChannelAdminLogEventActionChangeUsername:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeUsername{x}}
	case *PredChannelAdminLogEventActionChangePhoto:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangePhoto{x}}
	case *PredChannelAdminLogEventActionToggleInvites:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionToggleInvites{x}}
	case *PredChannelAdminLogEventActionToggleSignatures:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionToggleSignatures{x}}
	case *PredChannelAdminLogEventActionUpdatePinned:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionUpdatePinned{x}}
	case *PredChannelAdminLogEventActionEditMessage:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionEditMessage{x}}
	case *PredChannelAdminLogEventActionDeleteMessage:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionDeleteMessage{x}}
	case *PredChannelAdminLogEventActionParticipantJoin:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantJoin{x}}
	case *PredChannelAdminLogEventActionParticipantLeave:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantLeave{x}}
	case *PredChannelAdminLogEventActionParticipantInvite:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantInvite{x}}
	case *PredChannelAdminLogEventActionParticipantToggleBan:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleBan{x}}
	case *PredChannelAdminLogEventActionParticipantToggleAdmin:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleAdmin{x}}
	case *PredChannelAdminLogEventActionChangeStickerSet:
		return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeStickerSet{x}}
	}
	return nil
}
func toTypeChannelAdminLogEvent(tl TL) *TypeChannelAdminLogEvent {
	switch x := tl.(type) {
	case *PredChannelAdminLogEvent:
		return &TypeChannelAdminLogEvent{x}
	}
	return nil
}
func toTypeChannelsAdminLogResults(tl TL) *TypeChannelsAdminLogResults {
	switch x := tl.(type) {
	case *PredChannelsAdminLogResults:
		return &TypeChannelsAdminLogResults{x}
	}
	return nil
}
func toTypeChannelAdminLogEventsFilter(tl TL) *TypeChannelAdminLogEventsFilter {
	switch x := tl.(type) {
	case *PredChannelAdminLogEventsFilter:
		return &TypeChannelAdminLogEventsFilter{x}
	}
	return nil
}
func toTypePopularContact(tl TL) *TypePopularContact {
	switch x := tl.(type) {
	case *PredPopularContact:
		return &TypePopularContact{x}
	}
	return nil
}
func toTypeCdnFileHash(tl TL) *TypeCdnFileHash {
	switch x := tl.(type) {
	case *PredCdnFileHash:
		return &TypeCdnFileHash{x}
	}
	return nil
}
func toTypeMessagesFavedStickers(tl TL) *TypeMessagesFavedStickers {
	switch x := tl.(type) {
	case *PredMessagesFavedStickers:
		return &TypeMessagesFavedStickers{&TypeMessagesFavedStickers_MessagesFavedStickers{x}}
	case *PredMessagesFavedStickersNotModified:
		return &TypeMessagesFavedStickers{&TypeMessagesFavedStickers_MessagesFavedStickersNotModified{x}}
	}
	return nil
}

// TL slice converters to a Type slice
func toTypeBoolSlice(tlslice []TL) (converted []*TypeBool) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredBoolFalse:
			converted = append(converted, &TypeBool{&TypeBool_BoolFalse{x}})
		case *PredBoolTrue:
			converted = append(converted, &TypeBool{&TypeBool_BoolTrue{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeErrorSlice(tlslice []TL) (converted []*TypeError) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredError:
			converted = append(converted, &TypeError{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeNullSlice(tlslice []TL) (converted []*TypeNull) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredNull:
			converted = append(converted, &TypeNull{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputPeerSlice(tlslice []TL) (converted []*TypeInputPeer) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPeerEmpty:
			converted = append(converted, &TypeInputPeer{&TypeInputPeer_InputPeerEmpty{x}})
		case *PredInputPeerSelf:
			converted = append(converted, &TypeInputPeer{&TypeInputPeer_InputPeerSelf{x}})
		case *PredInputPeerChat:
			converted = append(converted, &TypeInputPeer{&TypeInputPeer_InputPeerChat{x}})
		case *PredInputPeerUser:
			converted = append(converted, &TypeInputPeer{&TypeInputPeer_InputPeerUser{x}})
		case *PredInputPeerChannel:
			converted = append(converted, &TypeInputPeer{&TypeInputPeer_InputPeerChannel{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputUserSlice(tlslice []TL) (converted []*TypeInputUser) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputUserEmpty:
			converted = append(converted, &TypeInputUser{&TypeInputUser_InputUserEmpty{x}})
		case *PredInputUserSelf:
			converted = append(converted, &TypeInputUser{&TypeInputUser_InputUserSelf{x}})
		case *PredInputUser:
			converted = append(converted, &TypeInputUser{&TypeInputUser_InputUser{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputContactSlice(tlslice []TL) (converted []*TypeInputContact) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPhoneContact:
			converted = append(converted, &TypeInputContact{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputFileSlice(tlslice []TL) (converted []*TypeInputFile) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputFile:
			converted = append(converted, &TypeInputFile{&TypeInputFile_InputFile{x}})
		case *PredInputFileBig:
			converted = append(converted, &TypeInputFile{&TypeInputFile_InputFileBig{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputMediaSlice(tlslice []TL) (converted []*TypeInputMedia) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputMediaEmpty:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaEmpty{x}})
		case *PredInputMediaUploadedPhoto:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaUploadedPhoto{x}})
		case *PredInputMediaPhoto:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaPhoto{x}})
		case *PredInputMediaGeoPoint:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaGeoPoint{x}})
		case *PredInputMediaContact:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaContact{x}})
		case *PredInputMediaUploadedDocument:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaUploadedDocument{x}})
		case *PredInputMediaDocument:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaDocument{x}})
		case *PredInputMediaVenue:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaVenue{x}})
		case *PredInputMediaGifExternal:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaGifExternal{x}})
		case *PredInputMediaPhotoExternal:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaPhotoExternal{x}})
		case *PredInputMediaDocumentExternal:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaDocumentExternal{x}})
		case *PredInputMediaGame:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaGame{x}})
		case *PredInputMediaInvoice:
			converted = append(converted, &TypeInputMedia{&TypeInputMedia_InputMediaInvoice{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputChatPhotoSlice(tlslice []TL) (converted []*TypeInputChatPhoto) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputChatPhotoEmpty:
			converted = append(converted, &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatPhotoEmpty{x}})
		case *PredInputChatUploadedPhoto:
			converted = append(converted, &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatUploadedPhoto{x}})
		case *PredInputChatPhoto:
			converted = append(converted, &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatPhoto{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputGeoPointSlice(tlslice []TL) (converted []*TypeInputGeoPoint) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputGeoPointEmpty:
			converted = append(converted, &TypeInputGeoPoint{&TypeInputGeoPoint_InputGeoPointEmpty{x}})
		case *PredInputGeoPoint:
			converted = append(converted, &TypeInputGeoPoint{&TypeInputGeoPoint_InputGeoPoint{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputPhotoSlice(tlslice []TL) (converted []*TypeInputPhoto) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPhotoEmpty:
			converted = append(converted, &TypeInputPhoto{&TypeInputPhoto_InputPhotoEmpty{x}})
		case *PredInputPhoto:
			converted = append(converted, &TypeInputPhoto{&TypeInputPhoto_InputPhoto{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputFileLocationSlice(tlslice []TL) (converted []*TypeInputFileLocation) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputFileLocation:
			converted = append(converted, &TypeInputFileLocation{&TypeInputFileLocation_InputFileLocation{x}})
		case *PredInputEncryptedFileLocation:
			converted = append(converted, &TypeInputFileLocation{&TypeInputFileLocation_InputEncryptedFileLocation{x}})
		case *PredInputDocumentFileLocation:
			converted = append(converted, &TypeInputFileLocation{&TypeInputFileLocation_InputDocumentFileLocation{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputAppEventSlice(tlslice []TL) (converted []*TypeInputAppEvent) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputAppEvent:
			converted = append(converted, &TypeInputAppEvent{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePeerSlice(tlslice []TL) (converted []*TypePeer) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPeerUser:
			converted = append(converted, &TypePeer{&TypePeer_PeerUser{x}})
		case *PredPeerChat:
			converted = append(converted, &TypePeer{&TypePeer_PeerChat{x}})
		case *PredPeerChannel:
			converted = append(converted, &TypePeer{&TypePeer_PeerChannel{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeStorageFileTypeSlice(tlslice []TL) (converted []*TypeStorageFileType) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredStorageFileUnknown:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFileUnknown{x}})
		case *PredStorageFileJpeg:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFileJpeg{x}})
		case *PredStorageFileGif:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFileGif{x}})
		case *PredStorageFilePng:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFilePng{x}})
		case *PredStorageFileMp3:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFileMp3{x}})
		case *PredStorageFileMov:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFileMov{x}})
		case *PredStorageFilePartial:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFilePartial{x}})
		case *PredStorageFileMp4:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFileMp4{x}})
		case *PredStorageFileWebp:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFileWebp{x}})
		case *PredStorageFilePdf:
			converted = append(converted, &TypeStorageFileType{&TypeStorageFileType_StorageFilePdf{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeFileLocationSlice(tlslice []TL) (converted []*TypeFileLocation) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredFileLocationUnavailable:
			converted = append(converted, &TypeFileLocation{&TypeFileLocation_FileLocationUnavailable{x}})
		case *PredFileLocation:
			converted = append(converted, &TypeFileLocation{&TypeFileLocation_FileLocation{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUserSlice(tlslice []TL) (converted []*TypeUser) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUserEmpty:
			converted = append(converted, &TypeUser{&TypeUser_UserEmpty{x}})
		case *PredUser:
			converted = append(converted, &TypeUser{&TypeUser_User{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUserProfilePhotoSlice(tlslice []TL) (converted []*TypeUserProfilePhoto) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUserProfilePhotoEmpty:
			converted = append(converted, &TypeUserProfilePhoto{&TypeUserProfilePhoto_UserProfilePhotoEmpty{x}})
		case *PredUserProfilePhoto:
			converted = append(converted, &TypeUserProfilePhoto{&TypeUserProfilePhoto_UserProfilePhoto{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUserStatusSlice(tlslice []TL) (converted []*TypeUserStatus) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUserStatusEmpty:
			converted = append(converted, &TypeUserStatus{&TypeUserStatus_UserStatusEmpty{x}})
		case *PredUserStatusOnline:
			converted = append(converted, &TypeUserStatus{&TypeUserStatus_UserStatusOnline{x}})
		case *PredUserStatusOffline:
			converted = append(converted, &TypeUserStatus{&TypeUserStatus_UserStatusOffline{x}})
		case *PredUserStatusRecently:
			converted = append(converted, &TypeUserStatus{&TypeUserStatus_UserStatusRecently{x}})
		case *PredUserStatusLastWeek:
			converted = append(converted, &TypeUserStatus{&TypeUserStatus_UserStatusLastWeek{x}})
		case *PredUserStatusLastMonth:
			converted = append(converted, &TypeUserStatus{&TypeUserStatus_UserStatusLastMonth{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChatSlice(tlslice []TL) (converted []*TypeChat) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChatEmpty:
			converted = append(converted, &TypeChat{&TypeChat_ChatEmpty{x}})
		case *PredChat:
			converted = append(converted, &TypeChat{&TypeChat_Chat{x}})
		case *PredChatForbidden:
			converted = append(converted, &TypeChat{&TypeChat_ChatForbidden{x}})
		case *PredChannel:
			converted = append(converted, &TypeChat{&TypeChat_Channel{x}})
		case *PredChannelForbidden:
			converted = append(converted, &TypeChat{&TypeChat_ChannelForbidden{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChatFullSlice(tlslice []TL) (converted []*TypeChatFull) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChatFull:
			converted = append(converted, &TypeChatFull{&TypeChatFull_ChatFull{x}})
		case *PredChannelFull:
			converted = append(converted, &TypeChatFull{&TypeChatFull_ChannelFull{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChatParticipantSlice(tlslice []TL) (converted []*TypeChatParticipant) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChatParticipant:
			converted = append(converted, &TypeChatParticipant{&TypeChatParticipant_ChatParticipant{x}})
		case *PredChatParticipantCreator:
			converted = append(converted, &TypeChatParticipant{&TypeChatParticipant_ChatParticipantCreator{x}})
		case *PredChatParticipantAdmin:
			converted = append(converted, &TypeChatParticipant{&TypeChatParticipant_ChatParticipantAdmin{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChatParticipantsSlice(tlslice []TL) (converted []*TypeChatParticipants) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChatParticipantsForbidden:
			converted = append(converted, &TypeChatParticipants{&TypeChatParticipants_ChatParticipantsForbidden{x}})
		case *PredChatParticipants:
			converted = append(converted, &TypeChatParticipants{&TypeChatParticipants_ChatParticipants{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChatPhotoSlice(tlslice []TL) (converted []*TypeChatPhoto) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChatPhotoEmpty:
			converted = append(converted, &TypeChatPhoto{&TypeChatPhoto_ChatPhotoEmpty{x}})
		case *PredChatPhoto:
			converted = append(converted, &TypeChatPhoto{&TypeChatPhoto_ChatPhoto{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessageSlice(tlslice []TL) (converted []*TypeMessage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessageEmpty:
			converted = append(converted, &TypeMessage{&TypeMessage_MessageEmpty{x}})
		case *PredMessage:
			converted = append(converted, &TypeMessage{&TypeMessage_Message{x}})
		case *PredMessageService:
			converted = append(converted, &TypeMessage{&TypeMessage_MessageService{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessageMediaSlice(tlslice []TL) (converted []*TypeMessageMedia) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessageMediaEmpty:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaEmpty{x}})
		case *PredMessageMediaPhoto:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaPhoto{x}})
		case *PredMessageMediaGeo:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaGeo{x}})
		case *PredMessageMediaContact:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaContact{x}})
		case *PredMessageMediaUnsupported:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaUnsupported{x}})
		case *PredMessageMediaDocument:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaDocument{x}})
		case *PredMessageMediaWebPage:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaWebPage{x}})
		case *PredMessageMediaVenue:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaVenue{x}})
		case *PredMessageMediaGame:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaGame{x}})
		case *PredMessageMediaInvoice:
			converted = append(converted, &TypeMessageMedia{&TypeMessageMedia_MessageMediaInvoice{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessageActionSlice(tlslice []TL) (converted []*TypeMessageAction) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessageActionEmpty:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionEmpty{x}})
		case *PredMessageActionChatCreate:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChatCreate{x}})
		case *PredMessageActionChatEditTitle:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChatEditTitle{x}})
		case *PredMessageActionChatEditPhoto:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChatEditPhoto{x}})
		case *PredMessageActionChatDeletePhoto:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChatDeletePhoto{x}})
		case *PredMessageActionChatAddUser:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChatAddUser{x}})
		case *PredMessageActionChatDeleteUser:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChatDeleteUser{x}})
		case *PredMessageActionChatJoinedByLink:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChatJoinedByLink{x}})
		case *PredMessageActionChannelCreate:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChannelCreate{x}})
		case *PredMessageActionChatMigrateTo:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChatMigrateTo{x}})
		case *PredMessageActionChannelMigrateFrom:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionChannelMigrateFrom{x}})
		case *PredMessageActionPinMessage:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionPinMessage{x}})
		case *PredMessageActionHistoryClear:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionHistoryClear{x}})
		case *PredMessageActionGameScore:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionGameScore{x}})
		case *PredMessageActionPhoneCall:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionPhoneCall{x}})
		case *PredMessageActionPaymentSentMe:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionPaymentSentMe{x}})
		case *PredMessageActionPaymentSent:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionPaymentSent{x}})
		case *PredMessageActionScreenshotTaken:
			converted = append(converted, &TypeMessageAction{&TypeMessageAction_MessageActionScreenshotTaken{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeDialogSlice(tlslice []TL) (converted []*TypeDialog) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredDialog:
			converted = append(converted, &TypeDialog{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhotoSlice(tlslice []TL) (converted []*TypePhoto) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhotoEmpty:
			converted = append(converted, &TypePhoto{&TypePhoto_PhotoEmpty{x}})
		case *PredPhoto:
			converted = append(converted, &TypePhoto{&TypePhoto_Photo{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhotoSizeSlice(tlslice []TL) (converted []*TypePhotoSize) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhotoSizeEmpty:
			converted = append(converted, &TypePhotoSize{&TypePhotoSize_PhotoSizeEmpty{x}})
		case *PredPhotoSize:
			converted = append(converted, &TypePhotoSize{&TypePhotoSize_PhotoSize{x}})
		case *PredPhotoCachedSize:
			converted = append(converted, &TypePhotoSize{&TypePhotoSize_PhotoCachedSize{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeGeoPointSlice(tlslice []TL) (converted []*TypeGeoPoint) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredGeoPointEmpty:
			converted = append(converted, &TypeGeoPoint{&TypeGeoPoint_GeoPointEmpty{x}})
		case *PredGeoPoint:
			converted = append(converted, &TypeGeoPoint{&TypeGeoPoint_GeoPoint{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAuthCheckedPhoneSlice(tlslice []TL) (converted []*TypeAuthCheckedPhone) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAuthCheckedPhone:
			converted = append(converted, &TypeAuthCheckedPhone{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAuthSentCodeSlice(tlslice []TL) (converted []*TypeAuthSentCode) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAuthSentCode:
			converted = append(converted, &TypeAuthSentCode{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAuthAuthorizationSlice(tlslice []TL) (converted []*TypeAuthAuthorization) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAuthAuthorization:
			converted = append(converted, &TypeAuthAuthorization{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAuthExportedAuthorizationSlice(tlslice []TL) (converted []*TypeAuthExportedAuthorization) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAuthExportedAuthorization:
			converted = append(converted, &TypeAuthExportedAuthorization{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputNotifyPeerSlice(tlslice []TL) (converted []*TypeInputNotifyPeer) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputNotifyPeer:
			converted = append(converted, &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyPeer{x}})
		case *PredInputNotifyUsers:
			converted = append(converted, &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyUsers{x}})
		case *PredInputNotifyChats:
			converted = append(converted, &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyChats{x}})
		case *PredInputNotifyAll:
			converted = append(converted, &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyAll{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputPeerNotifySettingsSlice(tlslice []TL) (converted []*TypeInputPeerNotifySettings) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPeerNotifySettings:
			converted = append(converted, &TypeInputPeerNotifySettings{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePeerNotifyEventsSlice(tlslice []TL) (converted []*TypePeerNotifyEvents) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPeerNotifyEventsEmpty:
			converted = append(converted, &TypePeerNotifyEvents{&TypePeerNotifyEvents_PeerNotifyEventsEmpty{x}})
		case *PredPeerNotifyEventsAll:
			converted = append(converted, &TypePeerNotifyEvents{&TypePeerNotifyEvents_PeerNotifyEventsAll{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePeerNotifySettingsSlice(tlslice []TL) (converted []*TypePeerNotifySettings) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPeerNotifySettingsEmpty:
			converted = append(converted, &TypePeerNotifySettings{&TypePeerNotifySettings_PeerNotifySettingsEmpty{x}})
		case *PredPeerNotifySettings:
			converted = append(converted, &TypePeerNotifySettings{&TypePeerNotifySettings_PeerNotifySettings{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeWallPaperSlice(tlslice []TL) (converted []*TypeWallPaper) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredWallPaper:
			converted = append(converted, &TypeWallPaper{&TypeWallPaper_WallPaper{x}})
		case *PredWallPaperSolid:
			converted = append(converted, &TypeWallPaper{&TypeWallPaper_WallPaperSolid{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUserFullSlice(tlslice []TL) (converted []*TypeUserFull) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUserFull:
			converted = append(converted, &TypeUserFull{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactSlice(tlslice []TL) (converted []*TypeContact) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContact:
			converted = append(converted, &TypeContact{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeImportedContactSlice(tlslice []TL) (converted []*TypeImportedContact) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredImportedContact:
			converted = append(converted, &TypeImportedContact{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactBlockedSlice(tlslice []TL) (converted []*TypeContactBlocked) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactBlocked:
			converted = append(converted, &TypeContactBlocked{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactStatusSlice(tlslice []TL) (converted []*TypeContactStatus) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactStatus:
			converted = append(converted, &TypeContactStatus{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactsLinkSlice(tlslice []TL) (converted []*TypeContactsLink) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactsLink:
			converted = append(converted, &TypeContactsLink{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactsContactsSlice(tlslice []TL) (converted []*TypeContactsContacts) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactsContacts:
			converted = append(converted, &TypeContactsContacts{&TypeContactsContacts_ContactsContacts{x}})
		case *PredContactsContactsNotModified:
			converted = append(converted, &TypeContactsContacts{&TypeContactsContacts_ContactsContactsNotModified{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactsImportedContactsSlice(tlslice []TL) (converted []*TypeContactsImportedContacts) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactsImportedContacts:
			converted = append(converted, &TypeContactsImportedContacts{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactsBlockedSlice(tlslice []TL) (converted []*TypeContactsBlocked) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactsBlocked:
			converted = append(converted, &TypeContactsBlocked{&TypeContactsBlocked_ContactsBlocked{x}})
		case *PredContactsBlockedSlice:
			converted = append(converted, &TypeContactsBlocked{&TypeContactsBlocked_ContactsBlockedSlice{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactsFoundSlice(tlslice []TL) (converted []*TypeContactsFound) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactsFound:
			converted = append(converted, &TypeContactsFound{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesDialogsSlice(tlslice []TL) (converted []*TypeMessagesDialogs) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesDialogs:
			converted = append(converted, &TypeMessagesDialogs{&TypeMessagesDialogs_MessagesDialogs{x}})
		case *PredMessagesDialogsSlice:
			converted = append(converted, &TypeMessagesDialogs{&TypeMessagesDialogs_MessagesDialogsSlice{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesMessagesSlice(tlslice []TL) (converted []*TypeMessagesMessages) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesMessages:
			converted = append(converted, &TypeMessagesMessages{&TypeMessagesMessages_MessagesMessages{x}})
		case *PredMessagesMessagesSlice:
			converted = append(converted, &TypeMessagesMessages{&TypeMessagesMessages_MessagesMessagesSlice{x}})
		case *PredMessagesChannelMessages:
			converted = append(converted, &TypeMessagesMessages{&TypeMessagesMessages_MessagesChannelMessages{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesChatsSlice(tlslice []TL) (converted []*TypeMessagesChats) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesChats:
			converted = append(converted, &TypeMessagesChats{&TypeMessagesChats_MessagesChats{x}})
		case *PredMessagesChatsSlice:
			converted = append(converted, &TypeMessagesChats{&TypeMessagesChats_MessagesChatsSlice{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesChatFullSlice(tlslice []TL) (converted []*TypeMessagesChatFull) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesChatFull:
			converted = append(converted, &TypeMessagesChatFull{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesAffectedHistorySlice(tlslice []TL) (converted []*TypeMessagesAffectedHistory) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesAffectedHistory:
			converted = append(converted, &TypeMessagesAffectedHistory{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesFilterSlice(tlslice []TL) (converted []*TypeMessagesFilter) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputMessagesFilterEmpty:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterEmpty{x}})
		case *PredInputMessagesFilterPhotos:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotos{x}})
		case *PredInputMessagesFilterVideo:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterVideo{x}})
		case *PredInputMessagesFilterPhotoVideo:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotoVideo{x}})
		case *PredInputMessagesFilterDocument:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterDocument{x}})
		case *PredInputMessagesFilterPhotoVideoDocuments:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotoVideoDocuments{x}})
		case *PredInputMessagesFilterUrl:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterUrl{x}})
		case *PredInputMessagesFilterGif:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterGif{x}})
		case *PredInputMessagesFilterVoice:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterVoice{x}})
		case *PredInputMessagesFilterMusic:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMusic{x}})
		case *PredInputMessagesFilterChatPhotos:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterChatPhotos{x}})
		case *PredInputMessagesFilterPhoneCalls:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhoneCalls{x}})
		case *PredInputMessagesFilterRoundVoice:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterRoundVoice{x}})
		case *PredInputMessagesFilterRoundVideo:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterRoundVideo{x}})
		case *PredInputMessagesFilterMyMentions:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMyMentions{x}})
		case *PredInputMessagesFilterMyMentionsUnread:
			converted = append(converted, &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMyMentionsUnread{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUpdateSlice(tlslice []TL) (converted []*TypeUpdate) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUpdateNewMessage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateNewMessage{x}})
		case *PredUpdateMessageID:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateMessageID{x}})
		case *PredUpdateDeleteMessages:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateDeleteMessages{x}})
		case *PredUpdateUserTyping:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateUserTyping{x}})
		case *PredUpdateChatUserTyping:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChatUserTyping{x}})
		case *PredUpdateChatParticipants:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChatParticipants{x}})
		case *PredUpdateUserStatus:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateUserStatus{x}})
		case *PredUpdateUserName:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateUserName{x}})
		case *PredUpdateUserPhoto:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateUserPhoto{x}})
		case *PredUpdateContactRegistered:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateContactRegistered{x}})
		case *PredUpdateContactLink:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateContactLink{x}})
		case *PredUpdateNewEncryptedMessage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateNewEncryptedMessage{x}})
		case *PredUpdateEncryptedChatTyping:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateEncryptedChatTyping{x}})
		case *PredUpdateEncryption:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateEncryption{x}})
		case *PredUpdateEncryptedMessagesRead:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateEncryptedMessagesRead{x}})
		case *PredUpdateChatParticipantAdd:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChatParticipantAdd{x}})
		case *PredUpdateChatParticipantDelete:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChatParticipantDelete{x}})
		case *PredUpdateDcOptions:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateDcOptions{x}})
		case *PredUpdateUserBlocked:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateUserBlocked{x}})
		case *PredUpdateNotifySettings:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateNotifySettings{x}})
		case *PredUpdateServiceNotification:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateServiceNotification{x}})
		case *PredUpdatePrivacy:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdatePrivacy{x}})
		case *PredUpdateUserPhone:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateUserPhone{x}})
		case *PredUpdateReadHistoryInbox:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateReadHistoryInbox{x}})
		case *PredUpdateReadHistoryOutbox:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateReadHistoryOutbox{x}})
		case *PredUpdateWebPage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateWebPage{x}})
		case *PredUpdateReadMessagesContents:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateReadMessagesContents{x}})
		case *PredUpdateChannelTooLong:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChannelTooLong{x}})
		case *PredUpdateChannel:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChannel{x}})
		case *PredUpdateNewChannelMessage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateNewChannelMessage{x}})
		case *PredUpdateReadChannelInbox:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateReadChannelInbox{x}})
		case *PredUpdateDeleteChannelMessages:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateDeleteChannelMessages{x}})
		case *PredUpdateChannelMessageViews:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChannelMessageViews{x}})
		case *PredUpdateChatAdmins:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChatAdmins{x}})
		case *PredUpdateChatParticipantAdmin:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChatParticipantAdmin{x}})
		case *PredUpdateNewStickerSet:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateNewStickerSet{x}})
		case *PredUpdateStickerSetsOrder:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateStickerSetsOrder{x}})
		case *PredUpdateStickerSets:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateStickerSets{x}})
		case *PredUpdateSavedGifs:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateSavedGifs{x}})
		case *PredUpdateBotInlineQuery:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateBotInlineQuery{x}})
		case *PredUpdateBotInlineSend:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateBotInlineSend{x}})
		case *PredUpdateEditChannelMessage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateEditChannelMessage{x}})
		case *PredUpdateChannelPinnedMessage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChannelPinnedMessage{x}})
		case *PredUpdateBotCallbackQuery:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateBotCallbackQuery{x}})
		case *PredUpdateEditMessage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateEditMessage{x}})
		case *PredUpdateInlineBotCallbackQuery:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateInlineBotCallbackQuery{x}})
		case *PredUpdateReadChannelOutbox:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateReadChannelOutbox{x}})
		case *PredUpdateDraftMessage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateDraftMessage{x}})
		case *PredUpdateReadFeaturedStickers:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateReadFeaturedStickers{x}})
		case *PredUpdateRecentStickers:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateRecentStickers{x}})
		case *PredUpdateConfig:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateConfig{x}})
		case *PredUpdatePtsChanged:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdatePtsChanged{x}})
		case *PredUpdateChannelWebPage:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChannelWebPage{x}})
		case *PredUpdatePhoneCall:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdatePhoneCall{x}})
		case *PredUpdateDialogPinned:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateDialogPinned{x}})
		case *PredUpdatePinnedDialogs:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdatePinnedDialogs{x}})
		case *PredUpdateBotWebhookJSON:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateBotWebhookJSON{x}})
		case *PredUpdateBotWebhookJSONQuery:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateBotWebhookJSONQuery{x}})
		case *PredUpdateBotShippingQuery:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateBotShippingQuery{x}})
		case *PredUpdateBotPrecheckoutQuery:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateBotPrecheckoutQuery{x}})
		case *PredUpdateLangPackTooLong:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateLangPackTooLong{x}})
		case *PredUpdateLangPack:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateLangPack{x}})
		case *PredUpdateContactsReset:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateContactsReset{x}})
		case *PredUpdateFavedStickers:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateFavedStickers{x}})
		case *PredUpdateChannelReadMessagesContents:
			converted = append(converted, &TypeUpdate{&TypeUpdate_UpdateChannelReadMessagesContents{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUpdatesStateSlice(tlslice []TL) (converted []*TypeUpdatesState) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUpdatesState:
			converted = append(converted, &TypeUpdatesState{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUpdatesDifferenceSlice(tlslice []TL) (converted []*TypeUpdatesDifference) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUpdatesDifferenceEmpty:
			converted = append(converted, &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceEmpty{x}})
		case *PredUpdatesDifference:
			converted = append(converted, &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifference{x}})
		case *PredUpdatesDifferenceSlice:
			converted = append(converted, &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceSlice{x}})
		case *PredUpdatesDifferenceTooLong:
			converted = append(converted, &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceTooLong{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUpdatesSlice(tlslice []TL) (converted []*TypeUpdates) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUpdatesTooLong:
			converted = append(converted, &TypeUpdates{&TypeUpdates_UpdatesTooLong{x}})
		case *PredUpdateShortMessage:
			converted = append(converted, &TypeUpdates{&TypeUpdates_UpdateShortMessage{x}})
		case *PredUpdateShortChatMessage:
			converted = append(converted, &TypeUpdates{&TypeUpdates_UpdateShortChatMessage{x}})
		case *PredUpdateShort:
			converted = append(converted, &TypeUpdates{&TypeUpdates_UpdateShort{x}})
		case *PredUpdatesCombined:
			converted = append(converted, &TypeUpdates{&TypeUpdates_UpdatesCombined{x}})
		case *PredUpdates:
			converted = append(converted, &TypeUpdates{&TypeUpdates_Updates{x}})
		case *PredUpdateShortSentMessage:
			converted = append(converted, &TypeUpdates{&TypeUpdates_UpdateShortSentMessage{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhotosPhotoSlice(tlslice []TL) (converted []*TypePhotosPhoto) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhotosPhoto:
			converted = append(converted, &TypePhotosPhoto{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUploadFileSlice(tlslice []TL) (converted []*TypeUploadFile) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUploadFile:
			converted = append(converted, &TypeUploadFile{&TypeUploadFile_UploadFile{x}})
		case *PredUploadFileCdnRedirect:
			converted = append(converted, &TypeUploadFile{&TypeUploadFile_UploadFileCdnRedirect{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeDcOptionSlice(tlslice []TL) (converted []*TypeDcOption) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredDcOption:
			converted = append(converted, &TypeDcOption{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeConfigSlice(tlslice []TL) (converted []*TypeConfig) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredConfig:
			converted = append(converted, &TypeConfig{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeNearestDcSlice(tlslice []TL) (converted []*TypeNearestDc) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredNearestDc:
			converted = append(converted, &TypeNearestDc{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeHelpAppUpdateSlice(tlslice []TL) (converted []*TypeHelpAppUpdate) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredHelpAppUpdate:
			converted = append(converted, &TypeHelpAppUpdate{&TypeHelpAppUpdate_HelpAppUpdate{x}})
		case *PredHelpNoAppUpdate:
			converted = append(converted, &TypeHelpAppUpdate{&TypeHelpAppUpdate_HelpNoAppUpdate{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeHelpInviteTextSlice(tlslice []TL) (converted []*TypeHelpInviteText) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredHelpInviteText:
			converted = append(converted, &TypeHelpInviteText{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputPeerNotifyEventsSlice(tlslice []TL) (converted []*TypeInputPeerNotifyEvents) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPeerNotifyEventsEmpty:
			converted = append(converted, &TypeInputPeerNotifyEvents{&TypeInputPeerNotifyEvents_InputPeerNotifyEventsEmpty{x}})
		case *PredInputPeerNotifyEventsAll:
			converted = append(converted, &TypeInputPeerNotifyEvents{&TypeInputPeerNotifyEvents_InputPeerNotifyEventsAll{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhotosPhotosSlice(tlslice []TL) (converted []*TypePhotosPhotos) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhotosPhotos:
			converted = append(converted, &TypePhotosPhotos{&TypePhotosPhotos_PhotosPhotos{x}})
		case *PredPhotosPhotosSlice:
			converted = append(converted, &TypePhotosPhotos{&TypePhotosPhotos_PhotosPhotosSlice{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeEncryptedChatSlice(tlslice []TL) (converted []*TypeEncryptedChat) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredEncryptedChatEmpty:
			converted = append(converted, &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatEmpty{x}})
		case *PredEncryptedChatWaiting:
			converted = append(converted, &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatWaiting{x}})
		case *PredEncryptedChatRequested:
			converted = append(converted, &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatRequested{x}})
		case *PredEncryptedChat:
			converted = append(converted, &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChat{x}})
		case *PredEncryptedChatDiscarded:
			converted = append(converted, &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatDiscarded{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputEncryptedChatSlice(tlslice []TL) (converted []*TypeInputEncryptedChat) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputEncryptedChat:
			converted = append(converted, &TypeInputEncryptedChat{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeEncryptedFileSlice(tlslice []TL) (converted []*TypeEncryptedFile) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredEncryptedFileEmpty:
			converted = append(converted, &TypeEncryptedFile{&TypeEncryptedFile_EncryptedFileEmpty{x}})
		case *PredEncryptedFile:
			converted = append(converted, &TypeEncryptedFile{&TypeEncryptedFile_EncryptedFile{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputEncryptedFileSlice(tlslice []TL) (converted []*TypeInputEncryptedFile) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputEncryptedFileEmpty:
			converted = append(converted, &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileEmpty{x}})
		case *PredInputEncryptedFileUploaded:
			converted = append(converted, &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileUploaded{x}})
		case *PredInputEncryptedFile:
			converted = append(converted, &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFile{x}})
		case *PredInputEncryptedFileBigUploaded:
			converted = append(converted, &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileBigUploaded{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeEncryptedMessageSlice(tlslice []TL) (converted []*TypeEncryptedMessage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredEncryptedMessage:
			converted = append(converted, &TypeEncryptedMessage{&TypeEncryptedMessage_EncryptedMessage{x}})
		case *PredEncryptedMessageService:
			converted = append(converted, &TypeEncryptedMessage{&TypeEncryptedMessage_EncryptedMessageService{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesDhConfigSlice(tlslice []TL) (converted []*TypeMessagesDhConfig) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesDhConfigNotModified:
			converted = append(converted, &TypeMessagesDhConfig{&TypeMessagesDhConfig_MessagesDhConfigNotModified{x}})
		case *PredMessagesDhConfig:
			converted = append(converted, &TypeMessagesDhConfig{&TypeMessagesDhConfig_MessagesDhConfig{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesSentEncryptedMessageSlice(tlslice []TL) (converted []*TypeMessagesSentEncryptedMessage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesSentEncryptedMessage:
			converted = append(converted, &TypeMessagesSentEncryptedMessage{&TypeMessagesSentEncryptedMessage_MessagesSentEncryptedMessage{x}})
		case *PredMessagesSentEncryptedFile:
			converted = append(converted, &TypeMessagesSentEncryptedMessage{&TypeMessagesSentEncryptedMessage_MessagesSentEncryptedFile{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputDocumentSlice(tlslice []TL) (converted []*TypeInputDocument) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputDocumentEmpty:
			converted = append(converted, &TypeInputDocument{&TypeInputDocument_InputDocumentEmpty{x}})
		case *PredInputDocument:
			converted = append(converted, &TypeInputDocument{&TypeInputDocument_InputDocument{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeDocumentSlice(tlslice []TL) (converted []*TypeDocument) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredDocumentEmpty:
			converted = append(converted, &TypeDocument{&TypeDocument_DocumentEmpty{x}})
		case *PredDocument:
			converted = append(converted, &TypeDocument{&TypeDocument_Document{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeHelpSupportSlice(tlslice []TL) (converted []*TypeHelpSupport) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredHelpSupport:
			converted = append(converted, &TypeHelpSupport{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeNotifyPeerSlice(tlslice []TL) (converted []*TypeNotifyPeer) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredNotifyAll:
			converted = append(converted, &TypeNotifyPeer{&TypeNotifyPeer_NotifyAll{x}})
		case *PredNotifyChats:
			converted = append(converted, &TypeNotifyPeer{&TypeNotifyPeer_NotifyChats{x}})
		case *PredNotifyPeer:
			converted = append(converted, &TypeNotifyPeer{&TypeNotifyPeer_NotifyPeer{x}})
		case *PredNotifyUsers:
			converted = append(converted, &TypeNotifyPeer{&TypeNotifyPeer_NotifyUsers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeSendMessageActionSlice(tlslice []TL) (converted []*TypeSendMessageAction) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredSendMessageTypingAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageTypingAction{x}})
		case *PredSendMessageCancelAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageCancelAction{x}})
		case *PredSendMessageRecordVideoAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordVideoAction{x}})
		case *PredSendMessageUploadVideoAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadVideoAction{x}})
		case *PredSendMessageRecordAudioAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordAudioAction{x}})
		case *PredSendMessageUploadAudioAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadAudioAction{x}})
		case *PredSendMessageUploadPhotoAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadPhotoAction{x}})
		case *PredSendMessageUploadDocumentAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadDocumentAction{x}})
		case *PredSendMessageGeoLocationAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageGeoLocationAction{x}})
		case *PredSendMessageChooseContactAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageChooseContactAction{x}})
		case *PredSendMessageGamePlayAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageGamePlayAction{x}})
		case *PredSendMessageRecordRoundAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordRoundAction{x}})
		case *PredSendMessageUploadRoundAction:
			converted = append(converted, &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadRoundAction{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputPrivacyKeySlice(tlslice []TL) (converted []*TypeInputPrivacyKey) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPrivacyKeyStatusTimestamp:
			converted = append(converted, &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyStatusTimestamp{x}})
		case *PredInputPrivacyKeyChatInvite:
			converted = append(converted, &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyChatInvite{x}})
		case *PredInputPrivacyKeyPhoneCall:
			converted = append(converted, &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyPhoneCall{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePrivacyKeySlice(tlslice []TL) (converted []*TypePrivacyKey) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPrivacyKeyStatusTimestamp:
			converted = append(converted, &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyStatusTimestamp{x}})
		case *PredPrivacyKeyChatInvite:
			converted = append(converted, &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyChatInvite{x}})
		case *PredPrivacyKeyPhoneCall:
			converted = append(converted, &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyPhoneCall{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputPrivacyRuleSlice(tlslice []TL) (converted []*TypeInputPrivacyRule) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPrivacyValueAllowContacts:
			converted = append(converted, &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowContacts{x}})
		case *PredInputPrivacyValueAllowAll:
			converted = append(converted, &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowAll{x}})
		case *PredInputPrivacyValueAllowUsers:
			converted = append(converted, &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowUsers{x}})
		case *PredInputPrivacyValueDisallowContacts:
			converted = append(converted, &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowContacts{x}})
		case *PredInputPrivacyValueDisallowAll:
			converted = append(converted, &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowAll{x}})
		case *PredInputPrivacyValueDisallowUsers:
			converted = append(converted, &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowUsers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePrivacyRuleSlice(tlslice []TL) (converted []*TypePrivacyRule) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPrivacyValueAllowContacts:
			converted = append(converted, &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowContacts{x}})
		case *PredPrivacyValueAllowAll:
			converted = append(converted, &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowAll{x}})
		case *PredPrivacyValueAllowUsers:
			converted = append(converted, &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowUsers{x}})
		case *PredPrivacyValueDisallowContacts:
			converted = append(converted, &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowContacts{x}})
		case *PredPrivacyValueDisallowAll:
			converted = append(converted, &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowAll{x}})
		case *PredPrivacyValueDisallowUsers:
			converted = append(converted, &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowUsers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAccountPrivacyRulesSlice(tlslice []TL) (converted []*TypeAccountPrivacyRules) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAccountPrivacyRules:
			converted = append(converted, &TypeAccountPrivacyRules{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAccountDaysTTLSlice(tlslice []TL) (converted []*TypeAccountDaysTTL) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAccountDaysTTL:
			converted = append(converted, &TypeAccountDaysTTL{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeDisabledFeatureSlice(tlslice []TL) (converted []*TypeDisabledFeature) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredDisabledFeature:
			converted = append(converted, &TypeDisabledFeature{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeDocumentAttributeSlice(tlslice []TL) (converted []*TypeDocumentAttribute) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredDocumentAttributeImageSize:
			converted = append(converted, &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeImageSize{x}})
		case *PredDocumentAttributeAnimated:
			converted = append(converted, &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeAnimated{x}})
		case *PredDocumentAttributeSticker:
			converted = append(converted, &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeSticker{x}})
		case *PredDocumentAttributeVideo:
			converted = append(converted, &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeVideo{x}})
		case *PredDocumentAttributeAudio:
			converted = append(converted, &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeAudio{x}})
		case *PredDocumentAttributeFilename:
			converted = append(converted, &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeFilename{x}})
		case *PredDocumentAttributeHasStickers:
			converted = append(converted, &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeHasStickers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesStickersSlice(tlslice []TL) (converted []*TypeMessagesStickers) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesStickersNotModified:
			converted = append(converted, &TypeMessagesStickers{&TypeMessagesStickers_MessagesStickersNotModified{x}})
		case *PredMessagesStickers:
			converted = append(converted, &TypeMessagesStickers{&TypeMessagesStickers_MessagesStickers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeStickerPackSlice(tlslice []TL) (converted []*TypeStickerPack) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredStickerPack:
			converted = append(converted, &TypeStickerPack{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesAllStickersSlice(tlslice []TL) (converted []*TypeMessagesAllStickers) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesAllStickersNotModified:
			converted = append(converted, &TypeMessagesAllStickers{&TypeMessagesAllStickers_MessagesAllStickersNotModified{x}})
		case *PredMessagesAllStickers:
			converted = append(converted, &TypeMessagesAllStickers{&TypeMessagesAllStickers_MessagesAllStickers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAccountPasswordSlice(tlslice []TL) (converted []*TypeAccountPassword) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAccountNoPassword:
			converted = append(converted, &TypeAccountPassword{&TypeAccountPassword_AccountNoPassword{x}})
		case *PredAccountPassword:
			converted = append(converted, &TypeAccountPassword{&TypeAccountPassword_AccountPassword{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesAffectedMessagesSlice(tlslice []TL) (converted []*TypeMessagesAffectedMessages) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesAffectedMessages:
			converted = append(converted, &TypeMessagesAffectedMessages{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactLinkSlice(tlslice []TL) (converted []*TypeContactLink) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactLinkUnknown:
			converted = append(converted, &TypeContactLink{&TypeContactLink_ContactLinkUnknown{x}})
		case *PredContactLinkNone:
			converted = append(converted, &TypeContactLink{&TypeContactLink_ContactLinkNone{x}})
		case *PredContactLinkHasPhone:
			converted = append(converted, &TypeContactLink{&TypeContactLink_ContactLinkHasPhone{x}})
		case *PredContactLinkContact:
			converted = append(converted, &TypeContactLink{&TypeContactLink_ContactLinkContact{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeWebPageSlice(tlslice []TL) (converted []*TypeWebPage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredWebPageEmpty:
			converted = append(converted, &TypeWebPage{&TypeWebPage_WebPageEmpty{x}})
		case *PredWebPagePending:
			converted = append(converted, &TypeWebPage{&TypeWebPage_WebPagePending{x}})
		case *PredWebPage:
			converted = append(converted, &TypeWebPage{&TypeWebPage_WebPage{x}})
		case *PredWebPageNotModified:
			converted = append(converted, &TypeWebPage{&TypeWebPage_WebPageNotModified{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAuthorizationSlice(tlslice []TL) (converted []*TypeAuthorization) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAuthorization:
			converted = append(converted, &TypeAuthorization{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAccountAuthorizationsSlice(tlslice []TL) (converted []*TypeAccountAuthorizations) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAccountAuthorizations:
			converted = append(converted, &TypeAccountAuthorizations{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAccountPasswordSettingsSlice(tlslice []TL) (converted []*TypeAccountPasswordSettings) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAccountPasswordSettings:
			converted = append(converted, &TypeAccountPasswordSettings{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAccountPasswordInputSettingsSlice(tlslice []TL) (converted []*TypeAccountPasswordInputSettings) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAccountPasswordInputSettings:
			converted = append(converted, &TypeAccountPasswordInputSettings{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAuthPasswordRecoverySlice(tlslice []TL) (converted []*TypeAuthPasswordRecovery) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAuthPasswordRecovery:
			converted = append(converted, &TypeAuthPasswordRecovery{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeReceivedNotifyMessageSlice(tlslice []TL) (converted []*TypeReceivedNotifyMessage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredReceivedNotifyMessage:
			converted = append(converted, &TypeReceivedNotifyMessage{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeExportedChatInviteSlice(tlslice []TL) (converted []*TypeExportedChatInvite) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChatInviteEmpty:
			converted = append(converted, &TypeExportedChatInvite{&TypeExportedChatInvite_ChatInviteEmpty{x}})
		case *PredChatInviteExported:
			converted = append(converted, &TypeExportedChatInvite{&TypeExportedChatInvite_ChatInviteExported{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChatInviteSlice(tlslice []TL) (converted []*TypeChatInvite) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChatInviteAlready:
			converted = append(converted, &TypeChatInvite{&TypeChatInvite_ChatInviteAlready{x}})
		case *PredChatInvite:
			converted = append(converted, &TypeChatInvite{&TypeChatInvite_ChatInvite{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputStickerSetSlice(tlslice []TL) (converted []*TypeInputStickerSet) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputStickerSetEmpty:
			converted = append(converted, &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetEmpty{x}})
		case *PredInputStickerSetID:
			converted = append(converted, &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetID{x}})
		case *PredInputStickerSetShortName:
			converted = append(converted, &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetShortName{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeStickerSetSlice(tlslice []TL) (converted []*TypeStickerSet) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredStickerSet:
			converted = append(converted, &TypeStickerSet{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesStickerSetSlice(tlslice []TL) (converted []*TypeMessagesStickerSet) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesStickerSet:
			converted = append(converted, &TypeMessagesStickerSet{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeBotCommandSlice(tlslice []TL) (converted []*TypeBotCommand) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredBotCommand:
			converted = append(converted, &TypeBotCommand{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeBotInfoSlice(tlslice []TL) (converted []*TypeBotInfo) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredBotInfo:
			converted = append(converted, &TypeBotInfo{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeKeyboardButtonSlice(tlslice []TL) (converted []*TypeKeyboardButton) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredKeyboardButton:
			converted = append(converted, &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButton{x}})
		case *PredKeyboardButtonUrl:
			converted = append(converted, &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonUrl{x}})
		case *PredKeyboardButtonCallback:
			converted = append(converted, &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonCallback{x}})
		case *PredKeyboardButtonRequestPhone:
			converted = append(converted, &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonRequestPhone{x}})
		case *PredKeyboardButtonRequestGeoLocation:
			converted = append(converted, &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonRequestGeoLocation{x}})
		case *PredKeyboardButtonSwitchInline:
			converted = append(converted, &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonSwitchInline{x}})
		case *PredKeyboardButtonGame:
			converted = append(converted, &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonGame{x}})
		case *PredKeyboardButtonBuy:
			converted = append(converted, &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonBuy{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeKeyboardButtonRowSlice(tlslice []TL) (converted []*TypeKeyboardButtonRow) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredKeyboardButtonRow:
			converted = append(converted, &TypeKeyboardButtonRow{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeReplyMarkupSlice(tlslice []TL) (converted []*TypeReplyMarkup) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredReplyKeyboardHide:
			converted = append(converted, &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardHide{x}})
		case *PredReplyKeyboardForceReply:
			converted = append(converted, &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardForceReply{x}})
		case *PredReplyKeyboardMarkup:
			converted = append(converted, &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardMarkup{x}})
		case *PredReplyInlineMarkup:
			converted = append(converted, &TypeReplyMarkup{&TypeReplyMarkup_ReplyInlineMarkup{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessageEntitySlice(tlslice []TL) (converted []*TypeMessageEntity) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessageEntityUnknown:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityUnknown{x}})
		case *PredMessageEntityMention:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityMention{x}})
		case *PredMessageEntityHashtag:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityHashtag{x}})
		case *PredMessageEntityBotCommand:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityBotCommand{x}})
		case *PredMessageEntityUrl:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityUrl{x}})
		case *PredMessageEntityEmail:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityEmail{x}})
		case *PredMessageEntityBold:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityBold{x}})
		case *PredMessageEntityItalic:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityItalic{x}})
		case *PredMessageEntityCode:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityCode{x}})
		case *PredMessageEntityPre:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityPre{x}})
		case *PredMessageEntityTextUrl:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityTextUrl{x}})
		case *PredMessageEntityMentionName:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_MessageEntityMentionName{x}})
		case *PredInputMessageEntityMentionName:
			converted = append(converted, &TypeMessageEntity{&TypeMessageEntity_InputMessageEntityMentionName{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputChannelSlice(tlslice []TL) (converted []*TypeInputChannel) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputChannelEmpty:
			converted = append(converted, &TypeInputChannel{&TypeInputChannel_InputChannelEmpty{x}})
		case *PredInputChannel:
			converted = append(converted, &TypeInputChannel{&TypeInputChannel_InputChannel{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactsResolvedPeerSlice(tlslice []TL) (converted []*TypeContactsResolvedPeer) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactsResolvedPeer:
			converted = append(converted, &TypeContactsResolvedPeer{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessageRangeSlice(tlslice []TL) (converted []*TypeMessageRange) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessageRange:
			converted = append(converted, &TypeMessageRange{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUpdatesChannelDifferenceSlice(tlslice []TL) (converted []*TypeUpdatesChannelDifference) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUpdatesChannelDifferenceEmpty:
			converted = append(converted, &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifferenceEmpty{x}})
		case *PredUpdatesChannelDifferenceTooLong:
			converted = append(converted, &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifferenceTooLong{x}})
		case *PredUpdatesChannelDifference:
			converted = append(converted, &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifference{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelMessagesFilterSlice(tlslice []TL) (converted []*TypeChannelMessagesFilter) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelMessagesFilterEmpty:
			converted = append(converted, &TypeChannelMessagesFilter{&TypeChannelMessagesFilter_ChannelMessagesFilterEmpty{x}})
		case *PredChannelMessagesFilter:
			converted = append(converted, &TypeChannelMessagesFilter{&TypeChannelMessagesFilter_ChannelMessagesFilter{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelParticipantSlice(tlslice []TL) (converted []*TypeChannelParticipant) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelParticipant:
			converted = append(converted, &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipant{x}})
		case *PredChannelParticipantSelf:
			converted = append(converted, &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantSelf{x}})
		case *PredChannelParticipantCreator:
			converted = append(converted, &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantCreator{x}})
		case *PredChannelParticipantAdmin:
			converted = append(converted, &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantAdmin{x}})
		case *PredChannelParticipantBanned:
			converted = append(converted, &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantBanned{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelParticipantsFilterSlice(tlslice []TL) (converted []*TypeChannelParticipantsFilter) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelParticipantsRecent:
			converted = append(converted, &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsRecent{x}})
		case *PredChannelParticipantsAdmins:
			converted = append(converted, &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsAdmins{x}})
		case *PredChannelParticipantsKicked:
			converted = append(converted, &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsKicked{x}})
		case *PredChannelParticipantsBots:
			converted = append(converted, &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsBots{x}})
		case *PredChannelParticipantsBanned:
			converted = append(converted, &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsBanned{x}})
		case *PredChannelParticipantsSearch:
			converted = append(converted, &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsSearch{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelsChannelParticipantsSlice(tlslice []TL) (converted []*TypeChannelsChannelParticipants) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelsChannelParticipants:
			converted = append(converted, &TypeChannelsChannelParticipants{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelsChannelParticipantSlice(tlslice []TL) (converted []*TypeChannelsChannelParticipant) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelsChannelParticipant:
			converted = append(converted, &TypeChannelsChannelParticipant{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeTrueSlice(tlslice []TL) (converted []*TypeTrue) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredTrue:
			converted = append(converted, &TypeTrue{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeReportReasonSlice(tlslice []TL) (converted []*TypeReportReason) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputReportReasonSpam:
			converted = append(converted, &TypeReportReason{&TypeReportReason_InputReportReasonSpam{x}})
		case *PredInputReportReasonViolence:
			converted = append(converted, &TypeReportReason{&TypeReportReason_InputReportReasonViolence{x}})
		case *PredInputReportReasonPornography:
			converted = append(converted, &TypeReportReason{&TypeReportReason_InputReportReasonPornography{x}})
		case *PredInputReportReasonOther:
			converted = append(converted, &TypeReportReason{&TypeReportReason_InputReportReasonOther{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeHelpTermsOfServiceSlice(tlslice []TL) (converted []*TypeHelpTermsOfService) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredHelpTermsOfService:
			converted = append(converted, &TypeHelpTermsOfService{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeFoundGifSlice(tlslice []TL) (converted []*TypeFoundGif) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredFoundGif:
			converted = append(converted, &TypeFoundGif{&TypeFoundGif_FoundGif{x}})
		case *PredFoundGifCached:
			converted = append(converted, &TypeFoundGif{&TypeFoundGif_FoundGifCached{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesFoundGifsSlice(tlslice []TL) (converted []*TypeMessagesFoundGifs) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesFoundGifs:
			converted = append(converted, &TypeMessagesFoundGifs{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesSavedGifsSlice(tlslice []TL) (converted []*TypeMessagesSavedGifs) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesSavedGifsNotModified:
			converted = append(converted, &TypeMessagesSavedGifs{&TypeMessagesSavedGifs_MessagesSavedGifsNotModified{x}})
		case *PredMessagesSavedGifs:
			converted = append(converted, &TypeMessagesSavedGifs{&TypeMessagesSavedGifs_MessagesSavedGifs{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputBotInlineMessageSlice(tlslice []TL) (converted []*TypeInputBotInlineMessage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputBotInlineMessageMediaAuto:
			converted = append(converted, &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaAuto{x}})
		case *PredInputBotInlineMessageText:
			converted = append(converted, &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageText{x}})
		case *PredInputBotInlineMessageMediaGeo:
			converted = append(converted, &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaGeo{x}})
		case *PredInputBotInlineMessageMediaVenue:
			converted = append(converted, &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaVenue{x}})
		case *PredInputBotInlineMessageMediaContact:
			converted = append(converted, &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaContact{x}})
		case *PredInputBotInlineMessageGame:
			converted = append(converted, &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageGame{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputBotInlineResultSlice(tlslice []TL) (converted []*TypeInputBotInlineResult) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputBotInlineResult:
			converted = append(converted, &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResult{x}})
		case *PredInputBotInlineResultPhoto:
			converted = append(converted, &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultPhoto{x}})
		case *PredInputBotInlineResultDocument:
			converted = append(converted, &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultDocument{x}})
		case *PredInputBotInlineResultGame:
			converted = append(converted, &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultGame{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeBotInlineMessageSlice(tlslice []TL) (converted []*TypeBotInlineMessage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredBotInlineMessageMediaAuto:
			converted = append(converted, &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaAuto{x}})
		case *PredBotInlineMessageText:
			converted = append(converted, &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageText{x}})
		case *PredBotInlineMessageMediaGeo:
			converted = append(converted, &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaGeo{x}})
		case *PredBotInlineMessageMediaVenue:
			converted = append(converted, &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaVenue{x}})
		case *PredBotInlineMessageMediaContact:
			converted = append(converted, &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaContact{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeBotInlineResultSlice(tlslice []TL) (converted []*TypeBotInlineResult) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredBotInlineResult:
			converted = append(converted, &TypeBotInlineResult{&TypeBotInlineResult_BotInlineResult{x}})
		case *PredBotInlineMediaResult:
			converted = append(converted, &TypeBotInlineResult{&TypeBotInlineResult_BotInlineMediaResult{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesBotResultsSlice(tlslice []TL) (converted []*TypeMessagesBotResults) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesBotResults:
			converted = append(converted, &TypeMessagesBotResults{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeExportedMessageLinkSlice(tlslice []TL) (converted []*TypeExportedMessageLink) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredExportedMessageLink:
			converted = append(converted, &TypeExportedMessageLink{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessageFwdHeaderSlice(tlslice []TL) (converted []*TypeMessageFwdHeader) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessageFwdHeader:
			converted = append(converted, &TypeMessageFwdHeader{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePeerSettingsSlice(tlslice []TL) (converted []*TypePeerSettings) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPeerSettings:
			converted = append(converted, &TypePeerSettings{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAuthCodeTypeSlice(tlslice []TL) (converted []*TypeAuthCodeType) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAuthCodeTypeSms:
			converted = append(converted, &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeSms{x}})
		case *PredAuthCodeTypeCall:
			converted = append(converted, &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeCall{x}})
		case *PredAuthCodeTypeFlashCall:
			converted = append(converted, &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeFlashCall{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAuthSentCodeTypeSlice(tlslice []TL) (converted []*TypeAuthSentCodeType) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAuthSentCodeTypeApp:
			converted = append(converted, &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeApp{x}})
		case *PredAuthSentCodeTypeSms:
			converted = append(converted, &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeSms{x}})
		case *PredAuthSentCodeTypeCall:
			converted = append(converted, &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeCall{x}})
		case *PredAuthSentCodeTypeFlashCall:
			converted = append(converted, &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeFlashCall{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesBotCallbackAnswerSlice(tlslice []TL) (converted []*TypeMessagesBotCallbackAnswer) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesBotCallbackAnswer:
			converted = append(converted, &TypeMessagesBotCallbackAnswer{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesMessageEditDataSlice(tlslice []TL) (converted []*TypeMessagesMessageEditData) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesMessageEditData:
			converted = append(converted, &TypeMessagesMessageEditData{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputBotInlineMessageIDSlice(tlslice []TL) (converted []*TypeInputBotInlineMessageID) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputBotInlineMessageID:
			converted = append(converted, &TypeInputBotInlineMessageID{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInlineBotSwitchPMSlice(tlslice []TL) (converted []*TypeInlineBotSwitchPM) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInlineBotSwitchPM:
			converted = append(converted, &TypeInlineBotSwitchPM{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesPeerDialogsSlice(tlslice []TL) (converted []*TypeMessagesPeerDialogs) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesPeerDialogs:
			converted = append(converted, &TypeMessagesPeerDialogs{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeTopPeerSlice(tlslice []TL) (converted []*TypeTopPeer) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredTopPeer:
			converted = append(converted, &TypeTopPeer{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeTopPeerCategorySlice(tlslice []TL) (converted []*TypeTopPeerCategory) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredTopPeerCategoryBotsPM:
			converted = append(converted, &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryBotsPM{x}})
		case *PredTopPeerCategoryBotsInline:
			converted = append(converted, &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryBotsInline{x}})
		case *PredTopPeerCategoryCorrespondents:
			converted = append(converted, &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryCorrespondents{x}})
		case *PredTopPeerCategoryGroups:
			converted = append(converted, &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryGroups{x}})
		case *PredTopPeerCategoryChannels:
			converted = append(converted, &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryChannels{x}})
		case *PredTopPeerCategoryPhoneCalls:
			converted = append(converted, &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryPhoneCalls{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeTopPeerCategoryPeersSlice(tlslice []TL) (converted []*TypeTopPeerCategoryPeers) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredTopPeerCategoryPeers:
			converted = append(converted, &TypeTopPeerCategoryPeers{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeContactsTopPeersSlice(tlslice []TL) (converted []*TypeContactsTopPeers) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredContactsTopPeersNotModified:
			converted = append(converted, &TypeContactsTopPeers{&TypeContactsTopPeers_ContactsTopPeersNotModified{x}})
		case *PredContactsTopPeers:
			converted = append(converted, &TypeContactsTopPeers{&TypeContactsTopPeers_ContactsTopPeers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeDraftMessageSlice(tlslice []TL) (converted []*TypeDraftMessage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredDraftMessageEmpty:
			converted = append(converted, &TypeDraftMessage{&TypeDraftMessage_DraftMessageEmpty{x}})
		case *PredDraftMessage:
			converted = append(converted, &TypeDraftMessage{&TypeDraftMessage_DraftMessage{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesFeaturedStickersSlice(tlslice []TL) (converted []*TypeMessagesFeaturedStickers) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesFeaturedStickersNotModified:
			converted = append(converted, &TypeMessagesFeaturedStickers{&TypeMessagesFeaturedStickers_MessagesFeaturedStickersNotModified{x}})
		case *PredMessagesFeaturedStickers:
			converted = append(converted, &TypeMessagesFeaturedStickers{&TypeMessagesFeaturedStickers_MessagesFeaturedStickers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesRecentStickersSlice(tlslice []TL) (converted []*TypeMessagesRecentStickers) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesRecentStickersNotModified:
			converted = append(converted, &TypeMessagesRecentStickers{&TypeMessagesRecentStickers_MessagesRecentStickersNotModified{x}})
		case *PredMessagesRecentStickers:
			converted = append(converted, &TypeMessagesRecentStickers{&TypeMessagesRecentStickers_MessagesRecentStickers{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesArchivedStickersSlice(tlslice []TL) (converted []*TypeMessagesArchivedStickers) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesArchivedStickers:
			converted = append(converted, &TypeMessagesArchivedStickers{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesStickerSetInstallResultSlice(tlslice []TL) (converted []*TypeMessagesStickerSetInstallResult) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesStickerSetInstallResultSuccess:
			converted = append(converted, &TypeMessagesStickerSetInstallResult{&TypeMessagesStickerSetInstallResult_MessagesStickerSetInstallResultSuccess{x}})
		case *PredMessagesStickerSetInstallResultArchive:
			converted = append(converted, &TypeMessagesStickerSetInstallResult{&TypeMessagesStickerSetInstallResult_MessagesStickerSetInstallResultArchive{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeStickerSetCoveredSlice(tlslice []TL) (converted []*TypeStickerSetCovered) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredStickerSetCovered:
			converted = append(converted, &TypeStickerSetCovered{&TypeStickerSetCovered_StickerSetCovered{x}})
		case *PredStickerSetMultiCovered:
			converted = append(converted, &TypeStickerSetCovered{&TypeStickerSetCovered_StickerSetMultiCovered{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMaskCoordsSlice(tlslice []TL) (converted []*TypeMaskCoords) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMaskCoords:
			converted = append(converted, &TypeMaskCoords{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputStickeredMediaSlice(tlslice []TL) (converted []*TypeInputStickeredMedia) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputStickeredMediaPhoto:
			converted = append(converted, &TypeInputStickeredMedia{&TypeInputStickeredMedia_InputStickeredMediaPhoto{x}})
		case *PredInputStickeredMediaDocument:
			converted = append(converted, &TypeInputStickeredMedia{&TypeInputStickeredMedia_InputStickeredMediaDocument{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeGameSlice(tlslice []TL) (converted []*TypeGame) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredGame:
			converted = append(converted, &TypeGame{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputGameSlice(tlslice []TL) (converted []*TypeInputGame) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputGameID:
			converted = append(converted, &TypeInputGame{&TypeInputGame_InputGameID{x}})
		case *PredInputGameShortName:
			converted = append(converted, &TypeInputGame{&TypeInputGame_InputGameShortName{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeHighScoreSlice(tlslice []TL) (converted []*TypeHighScore) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredHighScore:
			converted = append(converted, &TypeHighScore{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesHighScoresSlice(tlslice []TL) (converted []*TypeMessagesHighScores) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesHighScores:
			converted = append(converted, &TypeMessagesHighScores{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeRichTextSlice(tlslice []TL) (converted []*TypeRichText) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredTextEmpty:
			converted = append(converted, &TypeRichText{&TypeRichText_TextEmpty{x}})
		case *PredTextPlain:
			converted = append(converted, &TypeRichText{&TypeRichText_TextPlain{x}})
		case *PredTextBold:
			converted = append(converted, &TypeRichText{&TypeRichText_TextBold{x}})
		case *PredTextItalic:
			converted = append(converted, &TypeRichText{&TypeRichText_TextItalic{x}})
		case *PredTextUnderline:
			converted = append(converted, &TypeRichText{&TypeRichText_TextUnderline{x}})
		case *PredTextStrike:
			converted = append(converted, &TypeRichText{&TypeRichText_TextStrike{x}})
		case *PredTextFixed:
			converted = append(converted, &TypeRichText{&TypeRichText_TextFixed{x}})
		case *PredTextUrl:
			converted = append(converted, &TypeRichText{&TypeRichText_TextUrl{x}})
		case *PredTextEmail:
			converted = append(converted, &TypeRichText{&TypeRichText_TextEmail{x}})
		case *PredTextConcat:
			converted = append(converted, &TypeRichText{&TypeRichText_TextConcat{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePageBlockSlice(tlslice []TL) (converted []*TypePageBlock) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPageBlockTitle:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockTitle{x}})
		case *PredPageBlockSubtitle:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockSubtitle{x}})
		case *PredPageBlockAuthorDate:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockAuthorDate{x}})
		case *PredPageBlockHeader:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockHeader{x}})
		case *PredPageBlockSubheader:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockSubheader{x}})
		case *PredPageBlockParagraph:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockParagraph{x}})
		case *PredPageBlockPreformatted:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockPreformatted{x}})
		case *PredPageBlockFooter:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockFooter{x}})
		case *PredPageBlockDivider:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockDivider{x}})
		case *PredPageBlockList:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockList{x}})
		case *PredPageBlockBlockquote:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockBlockquote{x}})
		case *PredPageBlockPullquote:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockPullquote{x}})
		case *PredPageBlockPhoto:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockPhoto{x}})
		case *PredPageBlockVideo:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockVideo{x}})
		case *PredPageBlockCover:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockCover{x}})
		case *PredPageBlockEmbed:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockEmbed{x}})
		case *PredPageBlockEmbedPost:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockEmbedPost{x}})
		case *PredPageBlockSlideshow:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockSlideshow{x}})
		case *PredPageBlockUnsupported:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockUnsupported{x}})
		case *PredPageBlockAnchor:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockAnchor{x}})
		case *PredPageBlockCollage:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockCollage{x}})
		case *PredPageBlockChannel:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockChannel{x}})
		case *PredPageBlockAudio:
			converted = append(converted, &TypePageBlock{&TypePageBlock_PageBlockAudio{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePageSlice(tlslice []TL) (converted []*TypePage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPagePart:
			converted = append(converted, &TypePage{&TypePage_PagePart{x}})
		case *PredPageFull:
			converted = append(converted, &TypePage{&TypePage_PageFull{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputPhoneCallSlice(tlslice []TL) (converted []*TypeInputPhoneCall) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPhoneCall:
			converted = append(converted, &TypeInputPhoneCall{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhoneCallSlice(tlslice []TL) (converted []*TypePhoneCall) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhoneCallEmpty:
			converted = append(converted, &TypePhoneCall{&TypePhoneCall_PhoneCallEmpty{x}})
		case *PredPhoneCallWaiting:
			converted = append(converted, &TypePhoneCall{&TypePhoneCall_PhoneCallWaiting{x}})
		case *PredPhoneCallRequested:
			converted = append(converted, &TypePhoneCall{&TypePhoneCall_PhoneCallRequested{x}})
		case *PredPhoneCall:
			converted = append(converted, &TypePhoneCall{&TypePhoneCall_PhoneCall{x}})
		case *PredPhoneCallDiscarded:
			converted = append(converted, &TypePhoneCall{&TypePhoneCall_PhoneCallDiscarded{x}})
		case *PredPhoneCallAccepted:
			converted = append(converted, &TypePhoneCall{&TypePhoneCall_PhoneCallAccepted{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhoneConnectionSlice(tlslice []TL) (converted []*TypePhoneConnection) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhoneConnection:
			converted = append(converted, &TypePhoneConnection{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhoneCallProtocolSlice(tlslice []TL) (converted []*TypePhoneCallProtocol) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhoneCallProtocol:
			converted = append(converted, &TypePhoneCallProtocol{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhonePhoneCallSlice(tlslice []TL) (converted []*TypePhonePhoneCall) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhonePhoneCall:
			converted = append(converted, &TypePhonePhoneCall{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePhoneCallDiscardReasonSlice(tlslice []TL) (converted []*TypePhoneCallDiscardReason) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPhoneCallDiscardReasonMissed:
			converted = append(converted, &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonMissed{x}})
		case *PredPhoneCallDiscardReasonDisconnect:
			converted = append(converted, &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonDisconnect{x}})
		case *PredPhoneCallDiscardReasonHangup:
			converted = append(converted, &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonHangup{x}})
		case *PredPhoneCallDiscardReasonBusy:
			converted = append(converted, &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonBusy{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInvoiceSlice(tlslice []TL) (converted []*TypeInvoice) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInvoice:
			converted = append(converted, &TypeInvoice{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePaymentsPaymentFormSlice(tlslice []TL) (converted []*TypePaymentsPaymentForm) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPaymentsPaymentForm:
			converted = append(converted, &TypePaymentsPaymentForm{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePostAddressSlice(tlslice []TL) (converted []*TypePostAddress) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPostAddress:
			converted = append(converted, &TypePostAddress{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePaymentRequestedInfoSlice(tlslice []TL) (converted []*TypePaymentRequestedInfo) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPaymentRequestedInfo:
			converted = append(converted, &TypePaymentRequestedInfo{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeDataJSONSlice(tlslice []TL) (converted []*TypeDataJSON) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredDataJSON:
			converted = append(converted, &TypeDataJSON{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeLabeledPriceSlice(tlslice []TL) (converted []*TypeLabeledPrice) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredLabeledPrice:
			converted = append(converted, &TypeLabeledPrice{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePaymentChargeSlice(tlslice []TL) (converted []*TypePaymentCharge) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPaymentCharge:
			converted = append(converted, &TypePaymentCharge{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePaymentSavedCredentialsSlice(tlslice []TL) (converted []*TypePaymentSavedCredentials) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPaymentSavedCredentialsCard:
			converted = append(converted, &TypePaymentSavedCredentials{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeWebDocumentSlice(tlslice []TL) (converted []*TypeWebDocument) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredWebDocument:
			converted = append(converted, &TypeWebDocument{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputWebDocumentSlice(tlslice []TL) (converted []*TypeInputWebDocument) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputWebDocument:
			converted = append(converted, &TypeInputWebDocument{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputWebFileLocationSlice(tlslice []TL) (converted []*TypeInputWebFileLocation) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputWebFileLocation:
			converted = append(converted, &TypeInputWebFileLocation{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUploadWebFileSlice(tlslice []TL) (converted []*TypeUploadWebFile) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUploadWebFile:
			converted = append(converted, &TypeUploadWebFile{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePaymentsValidatedRequestedInfoSlice(tlslice []TL) (converted []*TypePaymentsValidatedRequestedInfo) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPaymentsValidatedRequestedInfo:
			converted = append(converted, &TypePaymentsValidatedRequestedInfo{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePaymentsPaymentResultSlice(tlslice []TL) (converted []*TypePaymentsPaymentResult) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPaymentsPaymentResult:
			converted = append(converted, &TypePaymentsPaymentResult{&TypePaymentsPaymentResult_PaymentsPaymentResult{x}})
		case *PredPaymentsPaymentVerficationNeeded:
			converted = append(converted, &TypePaymentsPaymentResult{&TypePaymentsPaymentResult_PaymentsPaymentVerficationNeeded{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePaymentsPaymentReceiptSlice(tlslice []TL) (converted []*TypePaymentsPaymentReceipt) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPaymentsPaymentReceipt:
			converted = append(converted, &TypePaymentsPaymentReceipt{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePaymentsSavedInfoSlice(tlslice []TL) (converted []*TypePaymentsSavedInfo) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPaymentsSavedInfo:
			converted = append(converted, &TypePaymentsSavedInfo{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputPaymentCredentialsSlice(tlslice []TL) (converted []*TypeInputPaymentCredentials) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputPaymentCredentialsSaved:
			converted = append(converted, &TypeInputPaymentCredentials{&TypeInputPaymentCredentials_InputPaymentCredentialsSaved{x}})
		case *PredInputPaymentCredentials:
			converted = append(converted, &TypeInputPaymentCredentials{&TypeInputPaymentCredentials_InputPaymentCredentials{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeAccountTmpPasswordSlice(tlslice []TL) (converted []*TypeAccountTmpPassword) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredAccountTmpPassword:
			converted = append(converted, &TypeAccountTmpPassword{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeShippingOptionSlice(tlslice []TL) (converted []*TypeShippingOption) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredShippingOption:
			converted = append(converted, &TypeShippingOption{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeUploadCdnFileSlice(tlslice []TL) (converted []*TypeUploadCdnFile) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredUploadCdnFileReuploadNeeded:
			converted = append(converted, &TypeUploadCdnFile{&TypeUploadCdnFile_UploadCdnFileReuploadNeeded{x}})
		case *PredUploadCdnFile:
			converted = append(converted, &TypeUploadCdnFile{&TypeUploadCdnFile_UploadCdnFile{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeCdnPublicKeySlice(tlslice []TL) (converted []*TypeCdnPublicKey) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredCdnPublicKey:
			converted = append(converted, &TypeCdnPublicKey{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeCdnConfigSlice(tlslice []TL) (converted []*TypeCdnConfig) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredCdnConfig:
			converted = append(converted, &TypeCdnConfig{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeInputStickerSetItemSlice(tlslice []TL) (converted []*TypeInputStickerSetItem) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredInputStickerSetItem:
			converted = append(converted, &TypeInputStickerSetItem{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeLangPackStringSlice(tlslice []TL) (converted []*TypeLangPackString) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredLangPackString:
			converted = append(converted, &TypeLangPackString{&TypeLangPackString_LangPackString{x}})
		case *PredLangPackStringPluralized:
			converted = append(converted, &TypeLangPackString{&TypeLangPackString_LangPackStringPluralized{x}})
		case *PredLangPackStringDeleted:
			converted = append(converted, &TypeLangPackString{&TypeLangPackString_LangPackStringDeleted{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeLangPackDifferenceSlice(tlslice []TL) (converted []*TypeLangPackDifference) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredLangPackDifference:
			converted = append(converted, &TypeLangPackDifference{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeLangPackLanguageSlice(tlslice []TL) (converted []*TypeLangPackLanguage) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredLangPackLanguage:
			converted = append(converted, &TypeLangPackLanguage{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelAdminRightsSlice(tlslice []TL) (converted []*TypeChannelAdminRights) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelAdminRights:
			converted = append(converted, &TypeChannelAdminRights{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelBannedRightsSlice(tlslice []TL) (converted []*TypeChannelBannedRights) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelBannedRights:
			converted = append(converted, &TypeChannelBannedRights{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelAdminLogEventActionSlice(tlslice []TL) (converted []*TypeChannelAdminLogEventAction) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelAdminLogEventActionChangeTitle:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeTitle{x}})
		case *PredChannelAdminLogEventActionChangeAbout:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeAbout{x}})
		case *PredChannelAdminLogEventActionChangeUsername:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeUsername{x}})
		case *PredChannelAdminLogEventActionChangePhoto:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangePhoto{x}})
		case *PredChannelAdminLogEventActionToggleInvites:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionToggleInvites{x}})
		case *PredChannelAdminLogEventActionToggleSignatures:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionToggleSignatures{x}})
		case *PredChannelAdminLogEventActionUpdatePinned:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionUpdatePinned{x}})
		case *PredChannelAdminLogEventActionEditMessage:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionEditMessage{x}})
		case *PredChannelAdminLogEventActionDeleteMessage:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionDeleteMessage{x}})
		case *PredChannelAdminLogEventActionParticipantJoin:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantJoin{x}})
		case *PredChannelAdminLogEventActionParticipantLeave:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantLeave{x}})
		case *PredChannelAdminLogEventActionParticipantInvite:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantInvite{x}})
		case *PredChannelAdminLogEventActionParticipantToggleBan:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleBan{x}})
		case *PredChannelAdminLogEventActionParticipantToggleAdmin:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleAdmin{x}})
		case *PredChannelAdminLogEventActionChangeStickerSet:
			converted = append(converted, &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeStickerSet{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelAdminLogEventSlice(tlslice []TL) (converted []*TypeChannelAdminLogEvent) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelAdminLogEvent:
			converted = append(converted, &TypeChannelAdminLogEvent{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelsAdminLogResultsSlice(tlslice []TL) (converted []*TypeChannelsAdminLogResults) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelsAdminLogResults:
			converted = append(converted, &TypeChannelsAdminLogResults{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeChannelAdminLogEventsFilterSlice(tlslice []TL) (converted []*TypeChannelAdminLogEventsFilter) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredChannelAdminLogEventsFilter:
			converted = append(converted, &TypeChannelAdminLogEventsFilter{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypePopularContactSlice(tlslice []TL) (converted []*TypePopularContact) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredPopularContact:
			converted = append(converted, &TypePopularContact{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeCdnFileHashSlice(tlslice []TL) (converted []*TypeCdnFileHash) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredCdnFileHash:
			converted = append(converted, &TypeCdnFileHash{x})
		default:
			// invalid predicate
		}
	}
	return converted
}
func toTypeMessagesFavedStickersSlice(tlslice []TL) (converted []*TypeMessagesFavedStickers) {
	for _, tl := range tlslice {
		switch x := tl.(type) {
		case *PredMessagesFavedStickers:
			converted = append(converted, &TypeMessagesFavedStickers{&TypeMessagesFavedStickers_MessagesFavedStickers{x}})
		case *PredMessagesFavedStickersNotModified:
			converted = append(converted, &TypeMessagesFavedStickers{&TypeMessagesFavedStickers_MessagesFavedStickersNotModified{x}})
		default:
			// invalid predicate
		}
	}
	return converted
}

// predicate converters to a Type
func (p *PredBoolFalse) ToType() TL {
	return &TypeBool{&TypeBool_BoolFalse{p}}
}
func (p *PredBoolTrue) ToType() TL {
	return &TypeBool{&TypeBool_BoolTrue{p}}
}
func (p *PredError) ToType() TL {
	return &TypeError{p}
}
func (p *PredNull) ToType() TL {
	return &TypeNull{p}
}
func (p *PredInputPeerEmpty) ToType() TL {
	return &TypeInputPeer{&TypeInputPeer_InputPeerEmpty{p}}
}
func (p *PredInputPeerSelf) ToType() TL {
	return &TypeInputPeer{&TypeInputPeer_InputPeerSelf{p}}
}
func (p *PredInputPeerChat) ToType() TL {
	return &TypeInputPeer{&TypeInputPeer_InputPeerChat{p}}
}
func (p *PredInputUserEmpty) ToType() TL {
	return &TypeInputUser{&TypeInputUser_InputUserEmpty{p}}
}
func (p *PredInputUserSelf) ToType() TL {
	return &TypeInputUser{&TypeInputUser_InputUserSelf{p}}
}
func (p *PredInputPhoneContact) ToType() TL {
	return &TypeInputContact{p}
}
func (p *PredInputFile) ToType() TL {
	return &TypeInputFile{&TypeInputFile_InputFile{p}}
}
func (p *PredInputMediaEmpty) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaEmpty{p}}
}
func (p *PredInputMediaUploadedPhoto) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaUploadedPhoto{p}}
}
func (p *PredInputMediaPhoto) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaPhoto{p}}
}
func (p *PredInputMediaGeoPoint) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaGeoPoint{p}}
}
func (p *PredInputMediaContact) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaContact{p}}
}
func (p *PredInputChatPhotoEmpty) ToType() TL {
	return &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatPhotoEmpty{p}}
}
func (p *PredInputChatUploadedPhoto) ToType() TL {
	return &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatUploadedPhoto{p}}
}
func (p *PredInputChatPhoto) ToType() TL {
	return &TypeInputChatPhoto{&TypeInputChatPhoto_InputChatPhoto{p}}
}
func (p *PredInputGeoPointEmpty) ToType() TL {
	return &TypeInputGeoPoint{&TypeInputGeoPoint_InputGeoPointEmpty{p}}
}
func (p *PredInputGeoPoint) ToType() TL {
	return &TypeInputGeoPoint{&TypeInputGeoPoint_InputGeoPoint{p}}
}
func (p *PredInputPhotoEmpty) ToType() TL {
	return &TypeInputPhoto{&TypeInputPhoto_InputPhotoEmpty{p}}
}
func (p *PredInputPhoto) ToType() TL {
	return &TypeInputPhoto{&TypeInputPhoto_InputPhoto{p}}
}
func (p *PredInputFileLocation) ToType() TL {
	return &TypeInputFileLocation{&TypeInputFileLocation_InputFileLocation{p}}
}
func (p *PredInputAppEvent) ToType() TL {
	return &TypeInputAppEvent{p}
}
func (p *PredPeerUser) ToType() TL {
	return &TypePeer{&TypePeer_PeerUser{p}}
}
func (p *PredPeerChat) ToType() TL {
	return &TypePeer{&TypePeer_PeerChat{p}}
}
func (p *PredStorageFileUnknown) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFileUnknown{p}}
}
func (p *PredStorageFileJpeg) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFileJpeg{p}}
}
func (p *PredStorageFileGif) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFileGif{p}}
}
func (p *PredStorageFilePng) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFilePng{p}}
}
func (p *PredStorageFileMp3) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFileMp3{p}}
}
func (p *PredStorageFileMov) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFileMov{p}}
}
func (p *PredStorageFilePartial) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFilePartial{p}}
}
func (p *PredStorageFileMp4) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFileMp4{p}}
}
func (p *PredStorageFileWebp) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFileWebp{p}}
}
func (p *PredFileLocationUnavailable) ToType() TL {
	return &TypeFileLocation{&TypeFileLocation_FileLocationUnavailable{p}}
}
func (p *PredFileLocation) ToType() TL {
	return &TypeFileLocation{&TypeFileLocation_FileLocation{p}}
}
func (p *PredUserEmpty) ToType() TL {
	return &TypeUser{&TypeUser_UserEmpty{p}}
}
func (p *PredUserProfilePhotoEmpty) ToType() TL {
	return &TypeUserProfilePhoto{&TypeUserProfilePhoto_UserProfilePhotoEmpty{p}}
}
func (p *PredUserProfilePhoto) ToType() TL {
	return &TypeUserProfilePhoto{&TypeUserProfilePhoto_UserProfilePhoto{p}}
}
func (p *PredUserStatusEmpty) ToType() TL {
	return &TypeUserStatus{&TypeUserStatus_UserStatusEmpty{p}}
}
func (p *PredUserStatusOnline) ToType() TL {
	return &TypeUserStatus{&TypeUserStatus_UserStatusOnline{p}}
}
func (p *PredUserStatusOffline) ToType() TL {
	return &TypeUserStatus{&TypeUserStatus_UserStatusOffline{p}}
}
func (p *PredChatEmpty) ToType() TL {
	return &TypeChat{&TypeChat_ChatEmpty{p}}
}
func (p *PredChat) ToType() TL {
	return &TypeChat{&TypeChat_Chat{p}}
}
func (p *PredChatForbidden) ToType() TL {
	return &TypeChat{&TypeChat_ChatForbidden{p}}
}
func (p *PredChatFull) ToType() TL {
	return &TypeChatFull{&TypeChatFull_ChatFull{p}}
}
func (p *PredChatParticipant) ToType() TL {
	return &TypeChatParticipant{&TypeChatParticipant_ChatParticipant{p}}
}
func (p *PredChatParticipantsForbidden) ToType() TL {
	return &TypeChatParticipants{&TypeChatParticipants_ChatParticipantsForbidden{p}}
}
func (p *PredChatParticipants) ToType() TL {
	return &TypeChatParticipants{&TypeChatParticipants_ChatParticipants{p}}
}
func (p *PredChatPhotoEmpty) ToType() TL {
	return &TypeChatPhoto{&TypeChatPhoto_ChatPhotoEmpty{p}}
}
func (p *PredChatPhoto) ToType() TL {
	return &TypeChatPhoto{&TypeChatPhoto_ChatPhoto{p}}
}
func (p *PredMessageEmpty) ToType() TL {
	return &TypeMessage{&TypeMessage_MessageEmpty{p}}
}
func (p *PredMessage) ToType() TL {
	return &TypeMessage{&TypeMessage_Message{p}}
}
func (p *PredMessageService) ToType() TL {
	return &TypeMessage{&TypeMessage_MessageService{p}}
}
func (p *PredMessageMediaEmpty) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaEmpty{p}}
}
func (p *PredMessageMediaPhoto) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaPhoto{p}}
}
func (p *PredMessageMediaGeo) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaGeo{p}}
}
func (p *PredMessageMediaContact) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaContact{p}}
}
func (p *PredMessageMediaUnsupported) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaUnsupported{p}}
}
func (p *PredMessageActionEmpty) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionEmpty{p}}
}
func (p *PredMessageActionChatCreate) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChatCreate{p}}
}
func (p *PredMessageActionChatEditTitle) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChatEditTitle{p}}
}
func (p *PredMessageActionChatEditPhoto) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChatEditPhoto{p}}
}
func (p *PredMessageActionChatDeletePhoto) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChatDeletePhoto{p}}
}
func (p *PredMessageActionChatAddUser) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChatAddUser{p}}
}
func (p *PredMessageActionChatDeleteUser) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChatDeleteUser{p}}
}
func (p *PredDialog) ToType() TL {
	return &TypeDialog{p}
}
func (p *PredPhotoEmpty) ToType() TL {
	return &TypePhoto{&TypePhoto_PhotoEmpty{p}}
}
func (p *PredPhoto) ToType() TL {
	return &TypePhoto{&TypePhoto_Photo{p}}
}
func (p *PredPhotoSizeEmpty) ToType() TL {
	return &TypePhotoSize{&TypePhotoSize_PhotoSizeEmpty{p}}
}
func (p *PredPhotoSize) ToType() TL {
	return &TypePhotoSize{&TypePhotoSize_PhotoSize{p}}
}
func (p *PredPhotoCachedSize) ToType() TL {
	return &TypePhotoSize{&TypePhotoSize_PhotoCachedSize{p}}
}
func (p *PredGeoPointEmpty) ToType() TL {
	return &TypeGeoPoint{&TypeGeoPoint_GeoPointEmpty{p}}
}
func (p *PredGeoPoint) ToType() TL {
	return &TypeGeoPoint{&TypeGeoPoint_GeoPoint{p}}
}
func (p *PredAuthCheckedPhone) ToType() TL {
	return &TypeAuthCheckedPhone{p}
}
func (p *PredAuthSentCode) ToType() TL {
	return &TypeAuthSentCode{p}
}
func (p *PredAuthAuthorization) ToType() TL {
	return &TypeAuthAuthorization{p}
}
func (p *PredAuthExportedAuthorization) ToType() TL {
	return &TypeAuthExportedAuthorization{p}
}
func (p *PredInputNotifyPeer) ToType() TL {
	return &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyPeer{p}}
}
func (p *PredInputNotifyUsers) ToType() TL {
	return &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyUsers{p}}
}
func (p *PredInputNotifyChats) ToType() TL {
	return &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyChats{p}}
}
func (p *PredInputNotifyAll) ToType() TL {
	return &TypeInputNotifyPeer{&TypeInputNotifyPeer_InputNotifyAll{p}}
}
func (p *PredInputPeerNotifySettings) ToType() TL {
	return &TypeInputPeerNotifySettings{p}
}
func (p *PredPeerNotifyEventsEmpty) ToType() TL {
	return &TypePeerNotifyEvents{&TypePeerNotifyEvents_PeerNotifyEventsEmpty{p}}
}
func (p *PredPeerNotifyEventsAll) ToType() TL {
	return &TypePeerNotifyEvents{&TypePeerNotifyEvents_PeerNotifyEventsAll{p}}
}
func (p *PredPeerNotifySettingsEmpty) ToType() TL {
	return &TypePeerNotifySettings{&TypePeerNotifySettings_PeerNotifySettingsEmpty{p}}
}
func (p *PredPeerNotifySettings) ToType() TL {
	return &TypePeerNotifySettings{&TypePeerNotifySettings_PeerNotifySettings{p}}
}
func (p *PredWallPaper) ToType() TL {
	return &TypeWallPaper{&TypeWallPaper_WallPaper{p}}
}
func (p *PredUserFull) ToType() TL {
	return &TypeUserFull{p}
}
func (p *PredContact) ToType() TL {
	return &TypeContact{p}
}
func (p *PredImportedContact) ToType() TL {
	return &TypeImportedContact{p}
}
func (p *PredContactBlocked) ToType() TL {
	return &TypeContactBlocked{p}
}
func (p *PredContactStatus) ToType() TL {
	return &TypeContactStatus{p}
}
func (p *PredContactsLink) ToType() TL {
	return &TypeContactsLink{p}
}
func (p *PredContactsContacts) ToType() TL {
	return &TypeContactsContacts{&TypeContactsContacts_ContactsContacts{p}}
}
func (p *PredContactsContactsNotModified) ToType() TL {
	return &TypeContactsContacts{&TypeContactsContacts_ContactsContactsNotModified{p}}
}
func (p *PredContactsImportedContacts) ToType() TL {
	return &TypeContactsImportedContacts{p}
}
func (p *PredContactsBlocked) ToType() TL {
	return &TypeContactsBlocked{&TypeContactsBlocked_ContactsBlocked{p}}
}
func (p *PredContactsBlockedSlice) ToType() TL {
	return &TypeContactsBlocked{&TypeContactsBlocked_ContactsBlockedSlice{p}}
}
func (p *PredContactsFound) ToType() TL {
	return &TypeContactsFound{p}
}
func (p *PredMessagesDialogs) ToType() TL {
	return &TypeMessagesDialogs{&TypeMessagesDialogs_MessagesDialogs{p}}
}
func (p *PredMessagesDialogsSlice) ToType() TL {
	return &TypeMessagesDialogs{&TypeMessagesDialogs_MessagesDialogsSlice{p}}
}
func (p *PredMessagesMessages) ToType() TL {
	return &TypeMessagesMessages{&TypeMessagesMessages_MessagesMessages{p}}
}
func (p *PredMessagesMessagesSlice) ToType() TL {
	return &TypeMessagesMessages{&TypeMessagesMessages_MessagesMessagesSlice{p}}
}
func (p *PredMessagesChats) ToType() TL {
	return &TypeMessagesChats{&TypeMessagesChats_MessagesChats{p}}
}
func (p *PredMessagesChatFull) ToType() TL {
	return &TypeMessagesChatFull{p}
}
func (p *PredMessagesAffectedHistory) ToType() TL {
	return &TypeMessagesAffectedHistory{p}
}
func (p *PredInputMessagesFilterEmpty) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterEmpty{p}}
}
func (p *PredInputMessagesFilterPhotos) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotos{p}}
}
func (p *PredInputMessagesFilterVideo) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterVideo{p}}
}
func (p *PredInputMessagesFilterPhotoVideo) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotoVideo{p}}
}
func (p *PredUpdateNewMessage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateNewMessage{p}}
}
func (p *PredUpdateMessageID) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateMessageID{p}}
}
func (p *PredUpdateDeleteMessages) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateDeleteMessages{p}}
}
func (p *PredUpdateUserTyping) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateUserTyping{p}}
}
func (p *PredUpdateChatUserTyping) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChatUserTyping{p}}
}
func (p *PredUpdateChatParticipants) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChatParticipants{p}}
}
func (p *PredUpdateUserStatus) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateUserStatus{p}}
}
func (p *PredUpdateUserName) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateUserName{p}}
}
func (p *PredUpdateUserPhoto) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateUserPhoto{p}}
}
func (p *PredUpdateContactRegistered) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateContactRegistered{p}}
}
func (p *PredUpdateContactLink) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateContactLink{p}}
}
func (p *PredUpdatesState) ToType() TL {
	return &TypeUpdatesState{p}
}
func (p *PredUpdatesDifferenceEmpty) ToType() TL {
	return &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceEmpty{p}}
}
func (p *PredUpdatesDifference) ToType() TL {
	return &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifference{p}}
}
func (p *PredUpdatesDifferenceSlice) ToType() TL {
	return &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceSlice{p}}
}
func (p *PredUpdatesTooLong) ToType() TL {
	return &TypeUpdates{&TypeUpdates_UpdatesTooLong{p}}
}
func (p *PredUpdateShortMessage) ToType() TL {
	return &TypeUpdates{&TypeUpdates_UpdateShortMessage{p}}
}
func (p *PredUpdateShortChatMessage) ToType() TL {
	return &TypeUpdates{&TypeUpdates_UpdateShortChatMessage{p}}
}
func (p *PredUpdateShort) ToType() TL {
	return &TypeUpdates{&TypeUpdates_UpdateShort{p}}
}
func (p *PredUpdatesCombined) ToType() TL {
	return &TypeUpdates{&TypeUpdates_UpdatesCombined{p}}
}
func (p *PredUpdates) ToType() TL {
	return &TypeUpdates{&TypeUpdates_Updates{p}}
}
func (p *PredPhotosPhoto) ToType() TL {
	return &TypePhotosPhoto{p}
}
func (p *PredUploadFile) ToType() TL {
	return &TypeUploadFile{&TypeUploadFile_UploadFile{p}}
}
func (p *PredDcOption) ToType() TL {
	return &TypeDcOption{p}
}
func (p *PredConfig) ToType() TL {
	return &TypeConfig{p}
}
func (p *PredNearestDc) ToType() TL {
	return &TypeNearestDc{p}
}
func (p *PredHelpAppUpdate) ToType() TL {
	return &TypeHelpAppUpdate{&TypeHelpAppUpdate_HelpAppUpdate{p}}
}
func (p *PredHelpNoAppUpdate) ToType() TL {
	return &TypeHelpAppUpdate{&TypeHelpAppUpdate_HelpNoAppUpdate{p}}
}
func (p *PredHelpInviteText) ToType() TL {
	return &TypeHelpInviteText{p}
}
func (p *PredInputPeerNotifyEventsEmpty) ToType() TL {
	return &TypeInputPeerNotifyEvents{&TypeInputPeerNotifyEvents_InputPeerNotifyEventsEmpty{p}}
}
func (p *PredInputPeerNotifyEventsAll) ToType() TL {
	return &TypeInputPeerNotifyEvents{&TypeInputPeerNotifyEvents_InputPeerNotifyEventsAll{p}}
}
func (p *PredPhotosPhotos) ToType() TL {
	return &TypePhotosPhotos{&TypePhotosPhotos_PhotosPhotos{p}}
}
func (p *PredPhotosPhotosSlice) ToType() TL {
	return &TypePhotosPhotos{&TypePhotosPhotos_PhotosPhotosSlice{p}}
}
func (p *PredWallPaperSolid) ToType() TL {
	return &TypeWallPaper{&TypeWallPaper_WallPaperSolid{p}}
}
func (p *PredUpdateNewEncryptedMessage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateNewEncryptedMessage{p}}
}
func (p *PredUpdateEncryptedChatTyping) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateEncryptedChatTyping{p}}
}
func (p *PredUpdateEncryption) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateEncryption{p}}
}
func (p *PredUpdateEncryptedMessagesRead) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateEncryptedMessagesRead{p}}
}
func (p *PredEncryptedChatEmpty) ToType() TL {
	return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatEmpty{p}}
}
func (p *PredEncryptedChatWaiting) ToType() TL {
	return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatWaiting{p}}
}
func (p *PredEncryptedChatRequested) ToType() TL {
	return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatRequested{p}}
}
func (p *PredEncryptedChat) ToType() TL {
	return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChat{p}}
}
func (p *PredEncryptedChatDiscarded) ToType() TL {
	return &TypeEncryptedChat{&TypeEncryptedChat_EncryptedChatDiscarded{p}}
}
func (p *PredInputEncryptedChat) ToType() TL {
	return &TypeInputEncryptedChat{p}
}
func (p *PredEncryptedFileEmpty) ToType() TL {
	return &TypeEncryptedFile{&TypeEncryptedFile_EncryptedFileEmpty{p}}
}
func (p *PredEncryptedFile) ToType() TL {
	return &TypeEncryptedFile{&TypeEncryptedFile_EncryptedFile{p}}
}
func (p *PredInputEncryptedFileEmpty) ToType() TL {
	return &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileEmpty{p}}
}
func (p *PredInputEncryptedFileUploaded) ToType() TL {
	return &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileUploaded{p}}
}
func (p *PredInputEncryptedFile) ToType() TL {
	return &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFile{p}}
}
func (p *PredInputEncryptedFileLocation) ToType() TL {
	return &TypeInputFileLocation{&TypeInputFileLocation_InputEncryptedFileLocation{p}}
}
func (p *PredEncryptedMessage) ToType() TL {
	return &TypeEncryptedMessage{&TypeEncryptedMessage_EncryptedMessage{p}}
}
func (p *PredEncryptedMessageService) ToType() TL {
	return &TypeEncryptedMessage{&TypeEncryptedMessage_EncryptedMessageService{p}}
}
func (p *PredMessagesDhConfigNotModified) ToType() TL {
	return &TypeMessagesDhConfig{&TypeMessagesDhConfig_MessagesDhConfigNotModified{p}}
}
func (p *PredMessagesDhConfig) ToType() TL {
	return &TypeMessagesDhConfig{&TypeMessagesDhConfig_MessagesDhConfig{p}}
}
func (p *PredMessagesSentEncryptedMessage) ToType() TL {
	return &TypeMessagesSentEncryptedMessage{&TypeMessagesSentEncryptedMessage_MessagesSentEncryptedMessage{p}}
}
func (p *PredMessagesSentEncryptedFile) ToType() TL {
	return &TypeMessagesSentEncryptedMessage{&TypeMessagesSentEncryptedMessage_MessagesSentEncryptedFile{p}}
}
func (p *PredInputFileBig) ToType() TL {
	return &TypeInputFile{&TypeInputFile_InputFileBig{p}}
}
func (p *PredInputEncryptedFileBigUploaded) ToType() TL {
	return &TypeInputEncryptedFile{&TypeInputEncryptedFile_InputEncryptedFileBigUploaded{p}}
}
func (p *PredStorageFilePdf) ToType() TL {
	return &TypeStorageFileType{&TypeStorageFileType_StorageFilePdf{p}}
}
func (p *PredInputMessagesFilterDocument) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterDocument{p}}
}
func (p *PredInputMessagesFilterPhotoVideoDocuments) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhotoVideoDocuments{p}}
}
func (p *PredUpdateChatParticipantAdd) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChatParticipantAdd{p}}
}
func (p *PredUpdateChatParticipantDelete) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChatParticipantDelete{p}}
}
func (p *PredUpdateDcOptions) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateDcOptions{p}}
}
func (p *PredInputMediaUploadedDocument) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaUploadedDocument{p}}
}
func (p *PredInputMediaDocument) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaDocument{p}}
}
func (p *PredMessageMediaDocument) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaDocument{p}}
}
func (p *PredInputDocumentEmpty) ToType() TL {
	return &TypeInputDocument{&TypeInputDocument_InputDocumentEmpty{p}}
}
func (p *PredInputDocument) ToType() TL {
	return &TypeInputDocument{&TypeInputDocument_InputDocument{p}}
}
func (p *PredInputDocumentFileLocation) ToType() TL {
	return &TypeInputFileLocation{&TypeInputFileLocation_InputDocumentFileLocation{p}}
}
func (p *PredDocumentEmpty) ToType() TL {
	return &TypeDocument{&TypeDocument_DocumentEmpty{p}}
}
func (p *PredDocument) ToType() TL {
	return &TypeDocument{&TypeDocument_Document{p}}
}
func (p *PredHelpSupport) ToType() TL {
	return &TypeHelpSupport{p}
}
func (p *PredNotifyAll) ToType() TL {
	return &TypeNotifyPeer{&TypeNotifyPeer_NotifyAll{p}}
}
func (p *PredNotifyChats) ToType() TL {
	return &TypeNotifyPeer{&TypeNotifyPeer_NotifyChats{p}}
}
func (p *PredNotifyPeer) ToType() TL {
	return &TypeNotifyPeer{&TypeNotifyPeer_NotifyPeer{p}}
}
func (p *PredNotifyUsers) ToType() TL {
	return &TypeNotifyPeer{&TypeNotifyPeer_NotifyUsers{p}}
}
func (p *PredUpdateUserBlocked) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateUserBlocked{p}}
}
func (p *PredUpdateNotifySettings) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateNotifySettings{p}}
}
func (p *PredSendMessageTypingAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageTypingAction{p}}
}
func (p *PredSendMessageCancelAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageCancelAction{p}}
}
func (p *PredSendMessageRecordVideoAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordVideoAction{p}}
}
func (p *PredSendMessageUploadVideoAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadVideoAction{p}}
}
func (p *PredSendMessageRecordAudioAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordAudioAction{p}}
}
func (p *PredSendMessageUploadAudioAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadAudioAction{p}}
}
func (p *PredSendMessageUploadPhotoAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadPhotoAction{p}}
}
func (p *PredSendMessageUploadDocumentAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadDocumentAction{p}}
}
func (p *PredSendMessageGeoLocationAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageGeoLocationAction{p}}
}
func (p *PredSendMessageChooseContactAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageChooseContactAction{p}}
}
func (p *PredUpdateServiceNotification) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateServiceNotification{p}}
}
func (p *PredUserStatusRecently) ToType() TL {
	return &TypeUserStatus{&TypeUserStatus_UserStatusRecently{p}}
}
func (p *PredUserStatusLastWeek) ToType() TL {
	return &TypeUserStatus{&TypeUserStatus_UserStatusLastWeek{p}}
}
func (p *PredUserStatusLastMonth) ToType() TL {
	return &TypeUserStatus{&TypeUserStatus_UserStatusLastMonth{p}}
}
func (p *PredUpdatePrivacy) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdatePrivacy{p}}
}
func (p *PredInputPrivacyKeyStatusTimestamp) ToType() TL {
	return &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyStatusTimestamp{p}}
}
func (p *PredPrivacyKeyStatusTimestamp) ToType() TL {
	return &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyStatusTimestamp{p}}
}
func (p *PredInputPrivacyValueAllowContacts) ToType() TL {
	return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowContacts{p}}
}
func (p *PredInputPrivacyValueAllowAll) ToType() TL {
	return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowAll{p}}
}
func (p *PredInputPrivacyValueAllowUsers) ToType() TL {
	return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueAllowUsers{p}}
}
func (p *PredInputPrivacyValueDisallowContacts) ToType() TL {
	return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowContacts{p}}
}
func (p *PredInputPrivacyValueDisallowAll) ToType() TL {
	return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowAll{p}}
}
func (p *PredInputPrivacyValueDisallowUsers) ToType() TL {
	return &TypeInputPrivacyRule{&TypeInputPrivacyRule_InputPrivacyValueDisallowUsers{p}}
}
func (p *PredPrivacyValueAllowContacts) ToType() TL {
	return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowContacts{p}}
}
func (p *PredPrivacyValueAllowAll) ToType() TL {
	return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowAll{p}}
}
func (p *PredPrivacyValueAllowUsers) ToType() TL {
	return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueAllowUsers{p}}
}
func (p *PredPrivacyValueDisallowContacts) ToType() TL {
	return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowContacts{p}}
}
func (p *PredPrivacyValueDisallowAll) ToType() TL {
	return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowAll{p}}
}
func (p *PredPrivacyValueDisallowUsers) ToType() TL {
	return &TypePrivacyRule{&TypePrivacyRule_PrivacyValueDisallowUsers{p}}
}
func (p *PredAccountPrivacyRules) ToType() TL {
	return &TypeAccountPrivacyRules{p}
}
func (p *PredAccountDaysTTL) ToType() TL {
	return &TypeAccountDaysTTL{p}
}
func (p *PredUpdateUserPhone) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateUserPhone{p}}
}
func (p *PredDisabledFeature) ToType() TL {
	return &TypeDisabledFeature{p}
}
func (p *PredDocumentAttributeImageSize) ToType() TL {
	return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeImageSize{p}}
}
func (p *PredDocumentAttributeAnimated) ToType() TL {
	return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeAnimated{p}}
}
func (p *PredDocumentAttributeSticker) ToType() TL {
	return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeSticker{p}}
}
func (p *PredDocumentAttributeVideo) ToType() TL {
	return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeVideo{p}}
}
func (p *PredDocumentAttributeAudio) ToType() TL {
	return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeAudio{p}}
}
func (p *PredDocumentAttributeFilename) ToType() TL {
	return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeFilename{p}}
}
func (p *PredMessagesStickersNotModified) ToType() TL {
	return &TypeMessagesStickers{&TypeMessagesStickers_MessagesStickersNotModified{p}}
}
func (p *PredMessagesStickers) ToType() TL {
	return &TypeMessagesStickers{&TypeMessagesStickers_MessagesStickers{p}}
}
func (p *PredStickerPack) ToType() TL {
	return &TypeStickerPack{p}
}
func (p *PredMessagesAllStickersNotModified) ToType() TL {
	return &TypeMessagesAllStickers{&TypeMessagesAllStickers_MessagesAllStickersNotModified{p}}
}
func (p *PredMessagesAllStickers) ToType() TL {
	return &TypeMessagesAllStickers{&TypeMessagesAllStickers_MessagesAllStickers{p}}
}
func (p *PredAccountNoPassword) ToType() TL {
	return &TypeAccountPassword{&TypeAccountPassword_AccountNoPassword{p}}
}
func (p *PredAccountPassword) ToType() TL {
	return &TypeAccountPassword{&TypeAccountPassword_AccountPassword{p}}
}
func (p *PredUpdateReadHistoryInbox) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateReadHistoryInbox{p}}
}
func (p *PredUpdateReadHistoryOutbox) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateReadHistoryOutbox{p}}
}
func (p *PredMessagesAffectedMessages) ToType() TL {
	return &TypeMessagesAffectedMessages{p}
}
func (p *PredContactLinkUnknown) ToType() TL {
	return &TypeContactLink{&TypeContactLink_ContactLinkUnknown{p}}
}
func (p *PredContactLinkNone) ToType() TL {
	return &TypeContactLink{&TypeContactLink_ContactLinkNone{p}}
}
func (p *PredContactLinkHasPhone) ToType() TL {
	return &TypeContactLink{&TypeContactLink_ContactLinkHasPhone{p}}
}
func (p *PredContactLinkContact) ToType() TL {
	return &TypeContactLink{&TypeContactLink_ContactLinkContact{p}}
}
func (p *PredUpdateWebPage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateWebPage{p}}
}
func (p *PredWebPageEmpty) ToType() TL {
	return &TypeWebPage{&TypeWebPage_WebPageEmpty{p}}
}
func (p *PredWebPagePending) ToType() TL {
	return &TypeWebPage{&TypeWebPage_WebPagePending{p}}
}
func (p *PredWebPage) ToType() TL {
	return &TypeWebPage{&TypeWebPage_WebPage{p}}
}
func (p *PredMessageMediaWebPage) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaWebPage{p}}
}
func (p *PredAuthorization) ToType() TL {
	return &TypeAuthorization{p}
}
func (p *PredAccountAuthorizations) ToType() TL {
	return &TypeAccountAuthorizations{p}
}
func (p *PredAccountPasswordSettings) ToType() TL {
	return &TypeAccountPasswordSettings{p}
}
func (p *PredAccountPasswordInputSettings) ToType() TL {
	return &TypeAccountPasswordInputSettings{p}
}
func (p *PredAuthPasswordRecovery) ToType() TL {
	return &TypeAuthPasswordRecovery{p}
}
func (p *PredInputMediaVenue) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaVenue{p}}
}
func (p *PredMessageMediaVenue) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaVenue{p}}
}
func (p *PredReceivedNotifyMessage) ToType() TL {
	return &TypeReceivedNotifyMessage{p}
}
func (p *PredChatInviteEmpty) ToType() TL {
	return &TypeExportedChatInvite{&TypeExportedChatInvite_ChatInviteEmpty{p}}
}
func (p *PredChatInviteExported) ToType() TL {
	return &TypeExportedChatInvite{&TypeExportedChatInvite_ChatInviteExported{p}}
}
func (p *PredChatInviteAlready) ToType() TL {
	return &TypeChatInvite{&TypeChatInvite_ChatInviteAlready{p}}
}
func (p *PredChatInvite) ToType() TL {
	return &TypeChatInvite{&TypeChatInvite_ChatInvite{p}}
}
func (p *PredMessageActionChatJoinedByLink) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChatJoinedByLink{p}}
}
func (p *PredUpdateReadMessagesContents) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateReadMessagesContents{p}}
}
func (p *PredInputStickerSetEmpty) ToType() TL {
	return &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetEmpty{p}}
}
func (p *PredInputStickerSetID) ToType() TL {
	return &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetID{p}}
}
func (p *PredInputStickerSetShortName) ToType() TL {
	return &TypeInputStickerSet{&TypeInputStickerSet_InputStickerSetShortName{p}}
}
func (p *PredStickerSet) ToType() TL {
	return &TypeStickerSet{p}
}
func (p *PredMessagesStickerSet) ToType() TL {
	return &TypeMessagesStickerSet{p}
}
func (p *PredUser) ToType() TL {
	return &TypeUser{&TypeUser_User{p}}
}
func (p *PredBotCommand) ToType() TL {
	return &TypeBotCommand{p}
}
func (p *PredBotInfo) ToType() TL {
	return &TypeBotInfo{p}
}
func (p *PredKeyboardButton) ToType() TL {
	return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButton{p}}
}
func (p *PredKeyboardButtonRow) ToType() TL {
	return &TypeKeyboardButtonRow{p}
}
func (p *PredReplyKeyboardHide) ToType() TL {
	return &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardHide{p}}
}
func (p *PredReplyKeyboardForceReply) ToType() TL {
	return &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardForceReply{p}}
}
func (p *PredReplyKeyboardMarkup) ToType() TL {
	return &TypeReplyMarkup{&TypeReplyMarkup_ReplyKeyboardMarkup{p}}
}
func (p *PredInputMessagesFilterUrl) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterUrl{p}}
}
func (p *PredInputPeerUser) ToType() TL {
	return &TypeInputPeer{&TypeInputPeer_InputPeerUser{p}}
}
func (p *PredInputUser) ToType() TL {
	return &TypeInputUser{&TypeInputUser_InputUser{p}}
}
func (p *PredMessageEntityUnknown) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityUnknown{p}}
}
func (p *PredMessageEntityMention) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityMention{p}}
}
func (p *PredMessageEntityHashtag) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityHashtag{p}}
}
func (p *PredMessageEntityBotCommand) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityBotCommand{p}}
}
func (p *PredMessageEntityUrl) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityUrl{p}}
}
func (p *PredMessageEntityEmail) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityEmail{p}}
}
func (p *PredMessageEntityBold) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityBold{p}}
}
func (p *PredMessageEntityItalic) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityItalic{p}}
}
func (p *PredMessageEntityCode) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityCode{p}}
}
func (p *PredMessageEntityPre) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityPre{p}}
}
func (p *PredMessageEntityTextUrl) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityTextUrl{p}}
}
func (p *PredUpdateShortSentMessage) ToType() TL {
	return &TypeUpdates{&TypeUpdates_UpdateShortSentMessage{p}}
}
func (p *PredInputPeerChannel) ToType() TL {
	return &TypeInputPeer{&TypeInputPeer_InputPeerChannel{p}}
}
func (p *PredPeerChannel) ToType() TL {
	return &TypePeer{&TypePeer_PeerChannel{p}}
}
func (p *PredChannel) ToType() TL {
	return &TypeChat{&TypeChat_Channel{p}}
}
func (p *PredChannelForbidden) ToType() TL {
	return &TypeChat{&TypeChat_ChannelForbidden{p}}
}
func (p *PredChannelFull) ToType() TL {
	return &TypeChatFull{&TypeChatFull_ChannelFull{p}}
}
func (p *PredMessageActionChannelCreate) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChannelCreate{p}}
}
func (p *PredMessagesChannelMessages) ToType() TL {
	return &TypeMessagesMessages{&TypeMessagesMessages_MessagesChannelMessages{p}}
}
func (p *PredUpdateChannelTooLong) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChannelTooLong{p}}
}
func (p *PredUpdateChannel) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChannel{p}}
}
func (p *PredUpdateNewChannelMessage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateNewChannelMessage{p}}
}
func (p *PredUpdateReadChannelInbox) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateReadChannelInbox{p}}
}
func (p *PredUpdateDeleteChannelMessages) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateDeleteChannelMessages{p}}
}
func (p *PredUpdateChannelMessageViews) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChannelMessageViews{p}}
}
func (p *PredInputChannelEmpty) ToType() TL {
	return &TypeInputChannel{&TypeInputChannel_InputChannelEmpty{p}}
}
func (p *PredInputChannel) ToType() TL {
	return &TypeInputChannel{&TypeInputChannel_InputChannel{p}}
}
func (p *PredContactsResolvedPeer) ToType() TL {
	return &TypeContactsResolvedPeer{p}
}
func (p *PredMessageRange) ToType() TL {
	return &TypeMessageRange{p}
}
func (p *PredUpdatesChannelDifferenceEmpty) ToType() TL {
	return &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifferenceEmpty{p}}
}
func (p *PredUpdatesChannelDifferenceTooLong) ToType() TL {
	return &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifferenceTooLong{p}}
}
func (p *PredUpdatesChannelDifference) ToType() TL {
	return &TypeUpdatesChannelDifference{&TypeUpdatesChannelDifference_UpdatesChannelDifference{p}}
}
func (p *PredChannelMessagesFilterEmpty) ToType() TL {
	return &TypeChannelMessagesFilter{&TypeChannelMessagesFilter_ChannelMessagesFilterEmpty{p}}
}
func (p *PredChannelMessagesFilter) ToType() TL {
	return &TypeChannelMessagesFilter{&TypeChannelMessagesFilter_ChannelMessagesFilter{p}}
}
func (p *PredChannelParticipant) ToType() TL {
	return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipant{p}}
}
func (p *PredChannelParticipantSelf) ToType() TL {
	return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantSelf{p}}
}
func (p *PredChannelParticipantCreator) ToType() TL {
	return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantCreator{p}}
}
func (p *PredChannelParticipantsRecent) ToType() TL {
	return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsRecent{p}}
}
func (p *PredChannelParticipantsAdmins) ToType() TL {
	return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsAdmins{p}}
}
func (p *PredChannelParticipantsKicked) ToType() TL {
	return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsKicked{p}}
}
func (p *PredChannelsChannelParticipants) ToType() TL {
	return &TypeChannelsChannelParticipants{p}
}
func (p *PredChannelsChannelParticipant) ToType() TL {
	return &TypeChannelsChannelParticipant{p}
}
func (p *PredTrue) ToType() TL {
	return &TypeTrue{p}
}
func (p *PredChatParticipantCreator) ToType() TL {
	return &TypeChatParticipant{&TypeChatParticipant_ChatParticipantCreator{p}}
}
func (p *PredChatParticipantAdmin) ToType() TL {
	return &TypeChatParticipant{&TypeChatParticipant_ChatParticipantAdmin{p}}
}
func (p *PredUpdateChatAdmins) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChatAdmins{p}}
}
func (p *PredUpdateChatParticipantAdmin) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChatParticipantAdmin{p}}
}
func (p *PredMessageActionChatMigrateTo) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChatMigrateTo{p}}
}
func (p *PredMessageActionChannelMigrateFrom) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionChannelMigrateFrom{p}}
}
func (p *PredChannelParticipantsBots) ToType() TL {
	return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsBots{p}}
}
func (p *PredInputReportReasonSpam) ToType() TL {
	return &TypeReportReason{&TypeReportReason_InputReportReasonSpam{p}}
}
func (p *PredInputReportReasonViolence) ToType() TL {
	return &TypeReportReason{&TypeReportReason_InputReportReasonViolence{p}}
}
func (p *PredInputReportReasonPornography) ToType() TL {
	return &TypeReportReason{&TypeReportReason_InputReportReasonPornography{p}}
}
func (p *PredInputReportReasonOther) ToType() TL {
	return &TypeReportReason{&TypeReportReason_InputReportReasonOther{p}}
}
func (p *PredUpdateNewStickerSet) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateNewStickerSet{p}}
}
func (p *PredUpdateStickerSetsOrder) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateStickerSetsOrder{p}}
}
func (p *PredUpdateStickerSets) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateStickerSets{p}}
}
func (p *PredHelpTermsOfService) ToType() TL {
	return &TypeHelpTermsOfService{p}
}
func (p *PredFoundGif) ToType() TL {
	return &TypeFoundGif{&TypeFoundGif_FoundGif{p}}
}
func (p *PredInputMediaGifExternal) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaGifExternal{p}}
}
func (p *PredMessagesFoundGifs) ToType() TL {
	return &TypeMessagesFoundGifs{p}
}
func (p *PredInputMessagesFilterGif) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterGif{p}}
}
func (p *PredUpdateSavedGifs) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateSavedGifs{p}}
}
func (p *PredUpdateBotInlineQuery) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateBotInlineQuery{p}}
}
func (p *PredFoundGifCached) ToType() TL {
	return &TypeFoundGif{&TypeFoundGif_FoundGifCached{p}}
}
func (p *PredMessagesSavedGifsNotModified) ToType() TL {
	return &TypeMessagesSavedGifs{&TypeMessagesSavedGifs_MessagesSavedGifsNotModified{p}}
}
func (p *PredMessagesSavedGifs) ToType() TL {
	return &TypeMessagesSavedGifs{&TypeMessagesSavedGifs_MessagesSavedGifs{p}}
}
func (p *PredInputBotInlineMessageMediaAuto) ToType() TL {
	return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaAuto{p}}
}
func (p *PredInputBotInlineMessageText) ToType() TL {
	return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageText{p}}
}
func (p *PredInputBotInlineResult) ToType() TL {
	return &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResult{p}}
}
func (p *PredBotInlineMessageMediaAuto) ToType() TL {
	return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaAuto{p}}
}
func (p *PredBotInlineMessageText) ToType() TL {
	return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageText{p}}
}
func (p *PredBotInlineResult) ToType() TL {
	return &TypeBotInlineResult{&TypeBotInlineResult_BotInlineResult{p}}
}
func (p *PredMessagesBotResults) ToType() TL {
	return &TypeMessagesBotResults{p}
}
func (p *PredInputMessagesFilterVoice) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterVoice{p}}
}
func (p *PredInputMessagesFilterMusic) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMusic{p}}
}
func (p *PredUpdateBotInlineSend) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateBotInlineSend{p}}
}
func (p *PredInputPrivacyKeyChatInvite) ToType() TL {
	return &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyChatInvite{p}}
}
func (p *PredPrivacyKeyChatInvite) ToType() TL {
	return &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyChatInvite{p}}
}
func (p *PredUpdateEditChannelMessage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateEditChannelMessage{p}}
}
func (p *PredExportedMessageLink) ToType() TL {
	return &TypeExportedMessageLink{p}
}
func (p *PredMessageFwdHeader) ToType() TL {
	return &TypeMessageFwdHeader{p}
}
func (p *PredMessageActionPinMessage) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionPinMessage{p}}
}
func (p *PredPeerSettings) ToType() TL {
	return &TypePeerSettings{p}
}
func (p *PredUpdateChannelPinnedMessage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChannelPinnedMessage{p}}
}
func (p *PredKeyboardButtonUrl) ToType() TL {
	return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonUrl{p}}
}
func (p *PredKeyboardButtonCallback) ToType() TL {
	return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonCallback{p}}
}
func (p *PredKeyboardButtonRequestPhone) ToType() TL {
	return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonRequestPhone{p}}
}
func (p *PredKeyboardButtonRequestGeoLocation) ToType() TL {
	return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonRequestGeoLocation{p}}
}
func (p *PredAuthCodeTypeSms) ToType() TL {
	return &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeSms{p}}
}
func (p *PredAuthCodeTypeCall) ToType() TL {
	return &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeCall{p}}
}
func (p *PredAuthCodeTypeFlashCall) ToType() TL {
	return &TypeAuthCodeType{&TypeAuthCodeType_AuthCodeTypeFlashCall{p}}
}
func (p *PredAuthSentCodeTypeApp) ToType() TL {
	return &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeApp{p}}
}
func (p *PredAuthSentCodeTypeSms) ToType() TL {
	return &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeSms{p}}
}
func (p *PredAuthSentCodeTypeCall) ToType() TL {
	return &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeCall{p}}
}
func (p *PredAuthSentCodeTypeFlashCall) ToType() TL {
	return &TypeAuthSentCodeType{&TypeAuthSentCodeType_AuthSentCodeTypeFlashCall{p}}
}
func (p *PredKeyboardButtonSwitchInline) ToType() TL {
	return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonSwitchInline{p}}
}
func (p *PredReplyInlineMarkup) ToType() TL {
	return &TypeReplyMarkup{&TypeReplyMarkup_ReplyInlineMarkup{p}}
}
func (p *PredMessagesBotCallbackAnswer) ToType() TL {
	return &TypeMessagesBotCallbackAnswer{p}
}
func (p *PredUpdateBotCallbackQuery) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateBotCallbackQuery{p}}
}
func (p *PredMessagesMessageEditData) ToType() TL {
	return &TypeMessagesMessageEditData{p}
}
func (p *PredUpdateEditMessage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateEditMessage{p}}
}
func (p *PredInputBotInlineMessageMediaGeo) ToType() TL {
	return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaGeo{p}}
}
func (p *PredInputBotInlineMessageMediaVenue) ToType() TL {
	return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaVenue{p}}
}
func (p *PredInputBotInlineMessageMediaContact) ToType() TL {
	return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageMediaContact{p}}
}
func (p *PredBotInlineMessageMediaGeo) ToType() TL {
	return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaGeo{p}}
}
func (p *PredBotInlineMessageMediaVenue) ToType() TL {
	return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaVenue{p}}
}
func (p *PredBotInlineMessageMediaContact) ToType() TL {
	return &TypeBotInlineMessage{&TypeBotInlineMessage_BotInlineMessageMediaContact{p}}
}
func (p *PredInputBotInlineResultPhoto) ToType() TL {
	return &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultPhoto{p}}
}
func (p *PredInputBotInlineResultDocument) ToType() TL {
	return &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultDocument{p}}
}
func (p *PredBotInlineMediaResult) ToType() TL {
	return &TypeBotInlineResult{&TypeBotInlineResult_BotInlineMediaResult{p}}
}
func (p *PredInputBotInlineMessageID) ToType() TL {
	return &TypeInputBotInlineMessageID{p}
}
func (p *PredUpdateInlineBotCallbackQuery) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateInlineBotCallbackQuery{p}}
}
func (p *PredInlineBotSwitchPM) ToType() TL {
	return &TypeInlineBotSwitchPM{p}
}
func (p *PredMessageEntityMentionName) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_MessageEntityMentionName{p}}
}
func (p *PredInputMessageEntityMentionName) ToType() TL {
	return &TypeMessageEntity{&TypeMessageEntity_InputMessageEntityMentionName{p}}
}
func (p *PredMessagesPeerDialogs) ToType() TL {
	return &TypeMessagesPeerDialogs{p}
}
func (p *PredTopPeer) ToType() TL {
	return &TypeTopPeer{p}
}
func (p *PredTopPeerCategoryBotsPM) ToType() TL {
	return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryBotsPM{p}}
}
func (p *PredTopPeerCategoryBotsInline) ToType() TL {
	return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryBotsInline{p}}
}
func (p *PredTopPeerCategoryCorrespondents) ToType() TL {
	return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryCorrespondents{p}}
}
func (p *PredTopPeerCategoryGroups) ToType() TL {
	return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryGroups{p}}
}
func (p *PredTopPeerCategoryChannels) ToType() TL {
	return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryChannels{p}}
}
func (p *PredTopPeerCategoryPeers) ToType() TL {
	return &TypeTopPeerCategoryPeers{p}
}
func (p *PredContactsTopPeersNotModified) ToType() TL {
	return &TypeContactsTopPeers{&TypeContactsTopPeers_ContactsTopPeersNotModified{p}}
}
func (p *PredContactsTopPeers) ToType() TL {
	return &TypeContactsTopPeers{&TypeContactsTopPeers_ContactsTopPeers{p}}
}
func (p *PredInputMessagesFilterChatPhotos) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterChatPhotos{p}}
}
func (p *PredUpdateReadChannelOutbox) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateReadChannelOutbox{p}}
}
func (p *PredUpdateDraftMessage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateDraftMessage{p}}
}
func (p *PredDraftMessageEmpty) ToType() TL {
	return &TypeDraftMessage{&TypeDraftMessage_DraftMessageEmpty{p}}
}
func (p *PredDraftMessage) ToType() TL {
	return &TypeDraftMessage{&TypeDraftMessage_DraftMessage{p}}
}
func (p *PredMessageActionHistoryClear) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionHistoryClear{p}}
}
func (p *PredUpdateReadFeaturedStickers) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateReadFeaturedStickers{p}}
}
func (p *PredUpdateRecentStickers) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateRecentStickers{p}}
}
func (p *PredMessagesFeaturedStickersNotModified) ToType() TL {
	return &TypeMessagesFeaturedStickers{&TypeMessagesFeaturedStickers_MessagesFeaturedStickersNotModified{p}}
}
func (p *PredMessagesFeaturedStickers) ToType() TL {
	return &TypeMessagesFeaturedStickers{&TypeMessagesFeaturedStickers_MessagesFeaturedStickers{p}}
}
func (p *PredMessagesRecentStickersNotModified) ToType() TL {
	return &TypeMessagesRecentStickers{&TypeMessagesRecentStickers_MessagesRecentStickersNotModified{p}}
}
func (p *PredMessagesRecentStickers) ToType() TL {
	return &TypeMessagesRecentStickers{&TypeMessagesRecentStickers_MessagesRecentStickers{p}}
}
func (p *PredMessagesArchivedStickers) ToType() TL {
	return &TypeMessagesArchivedStickers{p}
}
func (p *PredMessagesStickerSetInstallResultSuccess) ToType() TL {
	return &TypeMessagesStickerSetInstallResult{&TypeMessagesStickerSetInstallResult_MessagesStickerSetInstallResultSuccess{p}}
}
func (p *PredMessagesStickerSetInstallResultArchive) ToType() TL {
	return &TypeMessagesStickerSetInstallResult{&TypeMessagesStickerSetInstallResult_MessagesStickerSetInstallResultArchive{p}}
}
func (p *PredStickerSetCovered) ToType() TL {
	return &TypeStickerSetCovered{&TypeStickerSetCovered_StickerSetCovered{p}}
}
func (p *PredInputMediaPhotoExternal) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaPhotoExternal{p}}
}
func (p *PredInputMediaDocumentExternal) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaDocumentExternal{p}}
}
func (p *PredUpdateConfig) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateConfig{p}}
}
func (p *PredUpdatePtsChanged) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdatePtsChanged{p}}
}
func (p *PredMessageActionGameScore) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionGameScore{p}}
}
func (p *PredDocumentAttributeHasStickers) ToType() TL {
	return &TypeDocumentAttribute{&TypeDocumentAttribute_DocumentAttributeHasStickers{p}}
}
func (p *PredKeyboardButtonGame) ToType() TL {
	return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonGame{p}}
}
func (p *PredStickerSetMultiCovered) ToType() TL {
	return &TypeStickerSetCovered{&TypeStickerSetCovered_StickerSetMultiCovered{p}}
}
func (p *PredMaskCoords) ToType() TL {
	return &TypeMaskCoords{p}
}
func (p *PredInputStickeredMediaPhoto) ToType() TL {
	return &TypeInputStickeredMedia{&TypeInputStickeredMedia_InputStickeredMediaPhoto{p}}
}
func (p *PredInputStickeredMediaDocument) ToType() TL {
	return &TypeInputStickeredMedia{&TypeInputStickeredMedia_InputStickeredMediaDocument{p}}
}
func (p *PredInputMediaGame) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaGame{p}}
}
func (p *PredMessageMediaGame) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaGame{p}}
}
func (p *PredInputBotInlineMessageGame) ToType() TL {
	return &TypeInputBotInlineMessage{&TypeInputBotInlineMessage_InputBotInlineMessageGame{p}}
}
func (p *PredInputBotInlineResultGame) ToType() TL {
	return &TypeInputBotInlineResult{&TypeInputBotInlineResult_InputBotInlineResultGame{p}}
}
func (p *PredGame) ToType() TL {
	return &TypeGame{p}
}
func (p *PredInputGameID) ToType() TL {
	return &TypeInputGame{&TypeInputGame_InputGameID{p}}
}
func (p *PredInputGameShortName) ToType() TL {
	return &TypeInputGame{&TypeInputGame_InputGameShortName{p}}
}
func (p *PredHighScore) ToType() TL {
	return &TypeHighScore{p}
}
func (p *PredMessagesHighScores) ToType() TL {
	return &TypeMessagesHighScores{p}
}
func (p *PredMessagesChatsSlice) ToType() TL {
	return &TypeMessagesChats{&TypeMessagesChats_MessagesChatsSlice{p}}
}
func (p *PredUpdateChannelWebPage) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChannelWebPage{p}}
}
func (p *PredUpdatesDifferenceTooLong) ToType() TL {
	return &TypeUpdatesDifference{&TypeUpdatesDifference_UpdatesDifferenceTooLong{p}}
}
func (p *PredSendMessageGamePlayAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageGamePlayAction{p}}
}
func (p *PredWebPageNotModified) ToType() TL {
	return &TypeWebPage{&TypeWebPage_WebPageNotModified{p}}
}
func (p *PredTextEmpty) ToType() TL {
	return &TypeRichText{&TypeRichText_TextEmpty{p}}
}
func (p *PredTextPlain) ToType() TL {
	return &TypeRichText{&TypeRichText_TextPlain{p}}
}
func (p *PredTextBold) ToType() TL {
	return &TypeRichText{&TypeRichText_TextBold{p}}
}
func (p *PredTextItalic) ToType() TL {
	return &TypeRichText{&TypeRichText_TextItalic{p}}
}
func (p *PredTextUnderline) ToType() TL {
	return &TypeRichText{&TypeRichText_TextUnderline{p}}
}
func (p *PredTextStrike) ToType() TL {
	return &TypeRichText{&TypeRichText_TextStrike{p}}
}
func (p *PredTextFixed) ToType() TL {
	return &TypeRichText{&TypeRichText_TextFixed{p}}
}
func (p *PredTextUrl) ToType() TL {
	return &TypeRichText{&TypeRichText_TextUrl{p}}
}
func (p *PredTextEmail) ToType() TL {
	return &TypeRichText{&TypeRichText_TextEmail{p}}
}
func (p *PredTextConcat) ToType() TL {
	return &TypeRichText{&TypeRichText_TextConcat{p}}
}
func (p *PredPageBlockTitle) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockTitle{p}}
}
func (p *PredPageBlockSubtitle) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockSubtitle{p}}
}
func (p *PredPageBlockAuthorDate) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockAuthorDate{p}}
}
func (p *PredPageBlockHeader) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockHeader{p}}
}
func (p *PredPageBlockSubheader) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockSubheader{p}}
}
func (p *PredPageBlockParagraph) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockParagraph{p}}
}
func (p *PredPageBlockPreformatted) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockPreformatted{p}}
}
func (p *PredPageBlockFooter) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockFooter{p}}
}
func (p *PredPageBlockDivider) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockDivider{p}}
}
func (p *PredPageBlockList) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockList{p}}
}
func (p *PredPageBlockBlockquote) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockBlockquote{p}}
}
func (p *PredPageBlockPullquote) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockPullquote{p}}
}
func (p *PredPageBlockPhoto) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockPhoto{p}}
}
func (p *PredPageBlockVideo) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockVideo{p}}
}
func (p *PredPageBlockCover) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockCover{p}}
}
func (p *PredPageBlockEmbed) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockEmbed{p}}
}
func (p *PredPageBlockEmbedPost) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockEmbedPost{p}}
}
func (p *PredPageBlockSlideshow) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockSlideshow{p}}
}
func (p *PredPagePart) ToType() TL {
	return &TypePage{&TypePage_PagePart{p}}
}
func (p *PredPageFull) ToType() TL {
	return &TypePage{&TypePage_PageFull{p}}
}
func (p *PredUpdatePhoneCall) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdatePhoneCall{p}}
}
func (p *PredUpdateDialogPinned) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateDialogPinned{p}}
}
func (p *PredUpdatePinnedDialogs) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdatePinnedDialogs{p}}
}
func (p *PredInputPrivacyKeyPhoneCall) ToType() TL {
	return &TypeInputPrivacyKey{&TypeInputPrivacyKey_InputPrivacyKeyPhoneCall{p}}
}
func (p *PredPrivacyKeyPhoneCall) ToType() TL {
	return &TypePrivacyKey{&TypePrivacyKey_PrivacyKeyPhoneCall{p}}
}
func (p *PredPageBlockUnsupported) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockUnsupported{p}}
}
func (p *PredPageBlockAnchor) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockAnchor{p}}
}
func (p *PredPageBlockCollage) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockCollage{p}}
}
func (p *PredInputPhoneCall) ToType() TL {
	return &TypeInputPhoneCall{p}
}
func (p *PredPhoneCallEmpty) ToType() TL {
	return &TypePhoneCall{&TypePhoneCall_PhoneCallEmpty{p}}
}
func (p *PredPhoneCallWaiting) ToType() TL {
	return &TypePhoneCall{&TypePhoneCall_PhoneCallWaiting{p}}
}
func (p *PredPhoneCallRequested) ToType() TL {
	return &TypePhoneCall{&TypePhoneCall_PhoneCallRequested{p}}
}
func (p *PredPhoneCall) ToType() TL {
	return &TypePhoneCall{&TypePhoneCall_PhoneCall{p}}
}
func (p *PredPhoneCallDiscarded) ToType() TL {
	return &TypePhoneCall{&TypePhoneCall_PhoneCallDiscarded{p}}
}
func (p *PredPhoneConnection) ToType() TL {
	return &TypePhoneConnection{p}
}
func (p *PredPhoneCallProtocol) ToType() TL {
	return &TypePhoneCallProtocol{p}
}
func (p *PredPhonePhoneCall) ToType() TL {
	return &TypePhonePhoneCall{p}
}
func (p *PredPhoneCallDiscardReasonMissed) ToType() TL {
	return &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonMissed{p}}
}
func (p *PredPhoneCallDiscardReasonDisconnect) ToType() TL {
	return &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonDisconnect{p}}
}
func (p *PredPhoneCallDiscardReasonHangup) ToType() TL {
	return &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonHangup{p}}
}
func (p *PredPhoneCallDiscardReasonBusy) ToType() TL {
	return &TypePhoneCallDiscardReason{&TypePhoneCallDiscardReason_PhoneCallDiscardReasonBusy{p}}
}
func (p *PredInputMessagesFilterPhoneCalls) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterPhoneCalls{p}}
}
func (p *PredMessageActionPhoneCall) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionPhoneCall{p}}
}
func (p *PredInvoice) ToType() TL {
	return &TypeInvoice{p}
}
func (p *PredInputMediaInvoice) ToType() TL {
	return &TypeInputMedia{&TypeInputMedia_InputMediaInvoice{p}}
}
func (p *PredMessageActionPaymentSentMe) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionPaymentSentMe{p}}
}
func (p *PredMessageMediaInvoice) ToType() TL {
	return &TypeMessageMedia{&TypeMessageMedia_MessageMediaInvoice{p}}
}
func (p *PredKeyboardButtonBuy) ToType() TL {
	return &TypeKeyboardButton{&TypeKeyboardButton_KeyboardButtonBuy{p}}
}
func (p *PredMessageActionPaymentSent) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionPaymentSent{p}}
}
func (p *PredPaymentsPaymentForm) ToType() TL {
	return &TypePaymentsPaymentForm{p}
}
func (p *PredPostAddress) ToType() TL {
	return &TypePostAddress{p}
}
func (p *PredPaymentRequestedInfo) ToType() TL {
	return &TypePaymentRequestedInfo{p}
}
func (p *PredUpdateBotWebhookJSON) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateBotWebhookJSON{p}}
}
func (p *PredUpdateBotWebhookJSONQuery) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateBotWebhookJSONQuery{p}}
}
func (p *PredUpdateBotShippingQuery) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateBotShippingQuery{p}}
}
func (p *PredUpdateBotPrecheckoutQuery) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateBotPrecheckoutQuery{p}}
}
func (p *PredDataJSON) ToType() TL {
	return &TypeDataJSON{p}
}
func (p *PredLabeledPrice) ToType() TL {
	return &TypeLabeledPrice{p}
}
func (p *PredPaymentCharge) ToType() TL {
	return &TypePaymentCharge{p}
}
func (p *PredPaymentSavedCredentialsCard) ToType() TL {
	return &TypePaymentSavedCredentials{p}
}
func (p *PredWebDocument) ToType() TL {
	return &TypeWebDocument{p}
}
func (p *PredInputWebDocument) ToType() TL {
	return &TypeInputWebDocument{p}
}
func (p *PredInputWebFileLocation) ToType() TL {
	return &TypeInputWebFileLocation{p}
}
func (p *PredUploadWebFile) ToType() TL {
	return &TypeUploadWebFile{p}
}
func (p *PredPaymentsValidatedRequestedInfo) ToType() TL {
	return &TypePaymentsValidatedRequestedInfo{p}
}
func (p *PredPaymentsPaymentResult) ToType() TL {
	return &TypePaymentsPaymentResult{&TypePaymentsPaymentResult_PaymentsPaymentResult{p}}
}
func (p *PredPaymentsPaymentVerficationNeeded) ToType() TL {
	return &TypePaymentsPaymentResult{&TypePaymentsPaymentResult_PaymentsPaymentVerficationNeeded{p}}
}
func (p *PredPaymentsPaymentReceipt) ToType() TL {
	return &TypePaymentsPaymentReceipt{p}
}
func (p *PredPaymentsSavedInfo) ToType() TL {
	return &TypePaymentsSavedInfo{p}
}
func (p *PredInputPaymentCredentialsSaved) ToType() TL {
	return &TypeInputPaymentCredentials{&TypeInputPaymentCredentials_InputPaymentCredentialsSaved{p}}
}
func (p *PredInputPaymentCredentials) ToType() TL {
	return &TypeInputPaymentCredentials{&TypeInputPaymentCredentials_InputPaymentCredentials{p}}
}
func (p *PredAccountTmpPassword) ToType() TL {
	return &TypeAccountTmpPassword{p}
}
func (p *PredShippingOption) ToType() TL {
	return &TypeShippingOption{p}
}
func (p *PredPhoneCallAccepted) ToType() TL {
	return &TypePhoneCall{&TypePhoneCall_PhoneCallAccepted{p}}
}
func (p *PredInputMessagesFilterRoundVoice) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterRoundVoice{p}}
}
func (p *PredInputMessagesFilterRoundVideo) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterRoundVideo{p}}
}
func (p *PredUploadFileCdnRedirect) ToType() TL {
	return &TypeUploadFile{&TypeUploadFile_UploadFileCdnRedirect{p}}
}
func (p *PredSendMessageRecordRoundAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageRecordRoundAction{p}}
}
func (p *PredSendMessageUploadRoundAction) ToType() TL {
	return &TypeSendMessageAction{&TypeSendMessageAction_SendMessageUploadRoundAction{p}}
}
func (p *PredUploadCdnFileReuploadNeeded) ToType() TL {
	return &TypeUploadCdnFile{&TypeUploadCdnFile_UploadCdnFileReuploadNeeded{p}}
}
func (p *PredUploadCdnFile) ToType() TL {
	return &TypeUploadCdnFile{&TypeUploadCdnFile_UploadCdnFile{p}}
}
func (p *PredCdnPublicKey) ToType() TL {
	return &TypeCdnPublicKey{p}
}
func (p *PredCdnConfig) ToType() TL {
	return &TypeCdnConfig{p}
}
func (p *PredUpdateLangPackTooLong) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateLangPackTooLong{p}}
}
func (p *PredUpdateLangPack) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateLangPack{p}}
}
func (p *PredPageBlockChannel) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockChannel{p}}
}
func (p *PredInputStickerSetItem) ToType() TL {
	return &TypeInputStickerSetItem{p}
}
func (p *PredLangPackString) ToType() TL {
	return &TypeLangPackString{&TypeLangPackString_LangPackString{p}}
}
func (p *PredLangPackStringPluralized) ToType() TL {
	return &TypeLangPackString{&TypeLangPackString_LangPackStringPluralized{p}}
}
func (p *PredLangPackStringDeleted) ToType() TL {
	return &TypeLangPackString{&TypeLangPackString_LangPackStringDeleted{p}}
}
func (p *PredLangPackDifference) ToType() TL {
	return &TypeLangPackDifference{p}
}
func (p *PredLangPackLanguage) ToType() TL {
	return &TypeLangPackLanguage{p}
}
func (p *PredChannelParticipantAdmin) ToType() TL {
	return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantAdmin{p}}
}
func (p *PredChannelParticipantBanned) ToType() TL {
	return &TypeChannelParticipant{&TypeChannelParticipant_ChannelParticipantBanned{p}}
}
func (p *PredChannelParticipantsBanned) ToType() TL {
	return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsBanned{p}}
}
func (p *PredChannelParticipantsSearch) ToType() TL {
	return &TypeChannelParticipantsFilter{&TypeChannelParticipantsFilter_ChannelParticipantsSearch{p}}
}
func (p *PredTopPeerCategoryPhoneCalls) ToType() TL {
	return &TypeTopPeerCategory{&TypeTopPeerCategory_TopPeerCategoryPhoneCalls{p}}
}
func (p *PredPageBlockAudio) ToType() TL {
	return &TypePageBlock{&TypePageBlock_PageBlockAudio{p}}
}
func (p *PredChannelAdminRights) ToType() TL {
	return &TypeChannelAdminRights{p}
}
func (p *PredChannelBannedRights) ToType() TL {
	return &TypeChannelBannedRights{p}
}
func (p *PredChannelAdminLogEventActionChangeTitle) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeTitle{p}}
}
func (p *PredChannelAdminLogEventActionChangeAbout) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeAbout{p}}
}
func (p *PredChannelAdminLogEventActionChangeUsername) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeUsername{p}}
}
func (p *PredChannelAdminLogEventActionChangePhoto) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangePhoto{p}}
}
func (p *PredChannelAdminLogEventActionToggleInvites) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionToggleInvites{p}}
}
func (p *PredChannelAdminLogEventActionToggleSignatures) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionToggleSignatures{p}}
}
func (p *PredChannelAdminLogEventActionUpdatePinned) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionUpdatePinned{p}}
}
func (p *PredChannelAdminLogEventActionEditMessage) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionEditMessage{p}}
}
func (p *PredChannelAdminLogEventActionDeleteMessage) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionDeleteMessage{p}}
}
func (p *PredChannelAdminLogEventActionParticipantJoin) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantJoin{p}}
}
func (p *PredChannelAdminLogEventActionParticipantLeave) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantLeave{p}}
}
func (p *PredChannelAdminLogEventActionParticipantInvite) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantInvite{p}}
}
func (p *PredChannelAdminLogEventActionParticipantToggleBan) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleBan{p}}
}
func (p *PredChannelAdminLogEventActionParticipantToggleAdmin) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionParticipantToggleAdmin{p}}
}
func (p *PredChannelAdminLogEvent) ToType() TL {
	return &TypeChannelAdminLogEvent{p}
}
func (p *PredChannelsAdminLogResults) ToType() TL {
	return &TypeChannelsAdminLogResults{p}
}
func (p *PredChannelAdminLogEventsFilter) ToType() TL {
	return &TypeChannelAdminLogEventsFilter{p}
}
func (p *PredMessageActionScreenshotTaken) ToType() TL {
	return &TypeMessageAction{&TypeMessageAction_MessageActionScreenshotTaken{p}}
}
func (p *PredPopularContact) ToType() TL {
	return &TypePopularContact{p}
}
func (p *PredCdnFileHash) ToType() TL {
	return &TypeCdnFileHash{p}
}
func (p *PredInputMessagesFilterMyMentions) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMyMentions{p}}
}
func (p *PredInputMessagesFilterMyMentionsUnread) ToType() TL {
	return &TypeMessagesFilter{&TypeMessagesFilter_InputMessagesFilterMyMentionsUnread{p}}
}
func (p *PredUpdateContactsReset) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateContactsReset{p}}
}
func (p *PredChannelAdminLogEventActionChangeStickerSet) ToType() TL {
	return &TypeChannelAdminLogEventAction{&TypeChannelAdminLogEventAction_ChannelAdminLogEventActionChangeStickerSet{p}}
}
func (p *PredUpdateFavedStickers) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateFavedStickers{p}}
}
func (p *PredMessagesFavedStickers) ToType() TL {
	return &TypeMessagesFavedStickers{&TypeMessagesFavedStickers_MessagesFavedStickers{p}}
}
func (p *PredMessagesFavedStickersNotModified) ToType() TL {
	return &TypeMessagesFavedStickers{&TypeMessagesFavedStickers_MessagesFavedStickersNotModified{p}}
}
func (p *PredUpdateChannelReadMessagesContents) ToType() TL {
	return &TypeUpdate{&TypeUpdate_UpdateChannelReadMessagesContents{p}}
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case crc_boolFalse:
		r = &PredBoolFalse{}

	case crc_boolTrue:
		r = &PredBoolTrue{}

	case crc_error:
		r = &PredError{
			m.Int(),
			m.String(),
		}

	case crc_null:
		r = &PredNull{}

	case crc_inputPeerEmpty:
		r = &PredInputPeerEmpty{}

	case crc_inputPeerSelf:
		r = &PredInputPeerSelf{}

	case crc_inputPeerChat:
		r = &PredInputPeerChat{
			m.Int(),
		}

	case crc_inputUserEmpty:
		r = &PredInputUserEmpty{}

	case crc_inputUserSelf:
		r = &PredInputUserSelf{}

	case crc_inputPhoneContact:
		r = &PredInputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputFile:
		r = &PredInputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_inputMediaEmpty:
		r = &PredInputMediaEmpty{}

	case crc_inputMediaUploadedPhoto:
		flags := m.Flags()
		_ = flags
		r = &PredInputMediaUploadedPhoto{
			flags,
			toTypeInputFile(m.Object()),
			m.String(),
			toTypeInputDocumentSlice(m.FlaggedVector(flags, 0)),
			m.FlaggedInt(flags, 1),
		}

	case crc_inputMediaPhoto:
		flags := m.Flags()
		_ = flags
		r = &PredInputMediaPhoto{
			flags,
			toTypeInputPhoto(m.Object()),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_inputMediaGeoPoint:
		r = &PredInputMediaGeoPoint{
			toTypeInputGeoPoint(m.Object()),
		}

	case crc_inputMediaContact:
		r = &PredInputMediaContact{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputChatPhotoEmpty:
		r = &PredInputChatPhotoEmpty{}

	case crc_inputChatUploadedPhoto:
		r = &PredInputChatUploadedPhoto{
			toTypeInputFile(m.Object()),
		}

	case crc_inputChatPhoto:
		r = &PredInputChatPhoto{
			toTypeInputPhoto(m.Object()),
		}

	case crc_inputGeoPointEmpty:
		r = &PredInputGeoPointEmpty{}

	case crc_inputGeoPoint:
		r = &PredInputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_inputPhotoEmpty:
		r = &PredInputPhotoEmpty{}

	case crc_inputPhoto:
		r = &PredInputPhoto{
			m.Long(),
			m.Long(),
		}

	case crc_inputFileLocation:
		r = &PredInputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_inputAppEvent:
		r = &PredInputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_peerUser:
		r = &PredPeerUser{
			m.Int(),
		}

	case crc_peerChat:
		r = &PredPeerChat{
			m.Int(),
		}

	case crc_storageFileUnknown:
		r = &PredStorageFileUnknown{}

	case crc_storageFileJpeg:
		r = &PredStorageFileJpeg{}

	case crc_storageFileGif:
		r = &PredStorageFileGif{}

	case crc_storageFilePng:
		r = &PredStorageFilePng{}

	case crc_storageFileMp3:
		r = &PredStorageFileMp3{}

	case crc_storageFileMov:
		r = &PredStorageFileMov{}

	case crc_storageFilePartial:
		r = &PredStorageFilePartial{}

	case crc_storageFileMp4:
		r = &PredStorageFileMp4{}

	case crc_storageFileWebp:
		r = &PredStorageFileWebp{}

	case crc_fileLocationUnavailable:
		r = &PredFileLocationUnavailable{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_fileLocation:
		r = &PredFileLocation{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_userEmpty:
		r = &PredUserEmpty{
			m.Int(),
		}

	case crc_userProfilePhotoEmpty:
		r = &PredUserProfilePhotoEmpty{}

	case crc_userProfilePhoto:
		r = &PredUserProfilePhoto{
			m.Long(),
			toTypeFileLocation(m.Object()),
			toTypeFileLocation(m.Object()),
		}

	case crc_userStatusEmpty:
		r = &PredUserStatusEmpty{}

	case crc_userStatusOnline:
		r = &PredUserStatusOnline{
			m.Int(),
		}

	case crc_userStatusOffline:
		r = &PredUserStatusOffline{
			m.Int(),
		}

	case crc_chatEmpty:
		r = &PredChatEmpty{
			m.Int(),
		}

	case crc_chat:
		flags := m.Flags()
		_ = flags
		r = &PredChat{
			flags,
			m.Int(),
			m.String(),
			toTypeChatPhoto(m.Object()),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypeInputChannel(m.FlaggedObject(flags, 6)),
		}

	case crc_chatForbidden:
		r = &PredChatForbidden{
			m.Int(),
			m.String(),
		}

	case crc_chatFull:
		r = &PredChatFull{
			m.Int(),
			toTypeChatParticipants(m.Object()),
			toTypePhoto(m.Object()),
			toTypePeerNotifySettings(m.Object()),
			toTypeExportedChatInvite(m.Object()),
			toTypeBotInfoSlice(m.Vector()),
		}

	case crc_chatParticipant:
		r = &PredChatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_chatParticipantsForbidden:
		flags := m.Flags()
		_ = flags
		r = &PredChatParticipantsForbidden{
			flags,
			m.Int(),
			toTypeChatParticipant(m.FlaggedObject(flags, 0)),
		}

	case crc_chatParticipants:
		r = &PredChatParticipants{
			m.Int(),
			toTypeChatParticipantSlice(m.Vector()),
			m.Int(),
		}

	case crc_chatPhotoEmpty:
		r = &PredChatPhotoEmpty{}

	case crc_chatPhoto:
		r = &PredChatPhoto{
			toTypeFileLocation(m.Object()),
			toTypeFileLocation(m.Object()),
		}

	case crc_messageEmpty:
		r = &PredMessageEmpty{
			m.Int(),
		}

	case crc_message:
		flags := m.Flags()
		_ = flags
		r = &PredMessage{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 8),
			toTypePeer(m.Object()),
			toTypeMessageFwdHeader(m.FlaggedObject(flags, 2)),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.Int(),
			m.String(),
			toTypeMessageMedia(m.FlaggedObject(flags, 9)),
			toTypeReplyMarkup(m.FlaggedObject(flags, 6)),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 7)),
			m.FlaggedInt(flags, 10),
			m.FlaggedInt(flags, 15),
			m.FlaggedString(flags, 16),
		}

	case crc_messageService:
		flags := m.Flags()
		_ = flags
		r = &PredMessageService{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 8),
			toTypePeer(m.Object()),
			m.FlaggedInt(flags, 3),
			m.Int(),
			toTypeMessageAction(m.Object()),
		}

	case crc_messageMediaEmpty:
		r = &PredMessageMediaEmpty{}

	case crc_messageMediaPhoto:
		flags := m.Flags()
		_ = flags
		r = &PredMessageMediaPhoto{
			flags,
			toTypePhoto(m.FlaggedObject(flags, 0)),
			m.FlaggedString(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crc_messageMediaGeo:
		r = &PredMessageMediaGeo{
			toTypeGeoPoint(m.Object()),
		}

	case crc_messageMediaContact:
		r = &PredMessageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_messageMediaUnsupported:
		r = &PredMessageMediaUnsupported{}

	case crc_messageActionEmpty:
		r = &PredMessageActionEmpty{}

	case crc_messageActionChatCreate:
		r = &PredMessageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case crc_messageActionChatEditTitle:
		r = &PredMessageActionChatEditTitle{
			m.String(),
		}

	case crc_messageActionChatEditPhoto:
		r = &PredMessageActionChatEditPhoto{
			toTypePhoto(m.Object()),
		}

	case crc_messageActionChatDeletePhoto:
		r = &PredMessageActionChatDeletePhoto{}

	case crc_messageActionChatAddUser:
		r = &PredMessageActionChatAddUser{
			m.VectorInt(),
		}

	case crc_messageActionChatDeleteUser:
		r = &PredMessageActionChatDeleteUser{
			m.Int(),
		}

	case crc_dialog:
		flags := m.Flags()
		_ = flags
		r = &PredDialog{
			flags,
			toTypePeer(m.Object()),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypePeerNotifySettings(m.Object()),
			m.FlaggedInt(flags, 0),
			toTypeDraftMessage(m.FlaggedObject(flags, 1)),
		}

	case crc_photoEmpty:
		r = &PredPhotoEmpty{
			m.Long(),
		}

	case crc_photo:
		flags := m.Flags()
		_ = flags
		r = &PredPhoto{
			flags,
			m.Long(),
			m.Long(),
			m.Int(),
			toTypePhotoSizeSlice(m.Vector()),
		}

	case crc_photoSizeEmpty:
		r = &PredPhotoSizeEmpty{
			m.String(),
		}

	case crc_photoSize:
		r = &PredPhotoSize{
			m.String(),
			toTypeFileLocation(m.Object()),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_photoCachedSize:
		r = &PredPhotoCachedSize{
			m.String(),
			toTypeFileLocation(m.Object()),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_geoPointEmpty:
		r = &PredGeoPointEmpty{}

	case crc_geoPoint:
		r = &PredGeoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_authCheckedPhone:
		r = &PredAuthCheckedPhone{
			toTypeBool(m.Object()),
		}

	case crc_authSentCode:
		flags := m.Flags()
		_ = flags
		r = &PredAuthSentCode{
			flags,
			toTypeAuthSentCodeType(m.Object()),
			m.String(),
			toTypeAuthCodeType(m.FlaggedObject(flags, 1)),
			m.FlaggedInt(flags, 2),
		}

	case crc_authAuthorization:
		flags := m.Flags()
		_ = flags
		r = &PredAuthAuthorization{
			flags,
			m.FlaggedInt(flags, 0),
			toTypeUser(m.Object()),
		}

	case crc_authExportedAuthorization:
		r = &PredAuthExportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputNotifyPeer:
		r = &PredInputNotifyPeer{
			toTypeInputPeer(m.Object()),
		}

	case crc_inputNotifyUsers:
		r = &PredInputNotifyUsers{}

	case crc_inputNotifyChats:
		r = &PredInputNotifyChats{}

	case crc_inputNotifyAll:
		r = &PredInputNotifyAll{}

	case crc_inputPeerNotifySettings:
		flags := m.Flags()
		_ = flags
		r = &PredInputPeerNotifySettings{
			flags,
			m.Int(),
			m.String(),
		}

	case crc_peerNotifyEventsEmpty:
		r = &PredPeerNotifyEventsEmpty{}

	case crc_peerNotifyEventsAll:
		r = &PredPeerNotifyEventsAll{}

	case crc_peerNotifySettingsEmpty:
		r = &PredPeerNotifySettingsEmpty{}

	case crc_peerNotifySettings:
		flags := m.Flags()
		_ = flags
		r = &PredPeerNotifySettings{
			flags,
			m.Int(),
			m.String(),
		}

	case crc_wallPaper:
		r = &PredWallPaper{
			m.Int(),
			m.String(),
			toTypePhotoSizeSlice(m.Vector()),
			m.Int(),
		}

	case crc_userFull:
		flags := m.Flags()
		_ = flags
		r = &PredUserFull{
			flags,
			toTypeUser(m.Object()),
			m.FlaggedString(flags, 1),
			toTypeContactsLink(m.Object()),
			toTypePhoto(m.FlaggedObject(flags, 2)),
			toTypePeerNotifySettings(m.Object()),
			toTypeBotInfo(m.FlaggedObject(flags, 3)),
			m.Int(),
		}

	case crc_contact:
		r = &PredContact{
			m.Int(),
			toTypeBool(m.Object()),
		}

	case crc_importedContact:
		r = &PredImportedContact{
			m.Int(),
			m.Long(),
		}

	case crc_contactBlocked:
		r = &PredContactBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_contactStatus:
		r = &PredContactStatus{
			m.Int(),
			toTypeUserStatus(m.Object()),
		}

	case crc_contactsLink:
		r = &PredContactsLink{
			toTypeContactLink(m.Object()),
			toTypeContactLink(m.Object()),
			toTypeUser(m.Object()),
		}

	case crc_contactsContacts:
		r = &PredContactsContacts{
			toTypeContactSlice(m.Vector()),
			m.Int(),
			toTypeUserSlice(m.Vector()),
		}

	case crc_contactsContactsNotModified:
		r = &PredContactsContactsNotModified{}

	case crc_contactsImportedContacts:
		r = &PredContactsImportedContacts{
			toTypeImportedContactSlice(m.Vector()),
			toTypePopularContactSlice(m.Vector()),
			m.VectorLong(),
			toTypeUserSlice(m.Vector()),
		}

	case crc_contactsBlocked:
		r = &PredContactsBlocked{
			toTypeContactBlockedSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_contactsBlockedSlice:
		r = &PredContactsBlockedSlice{
			m.Int(),
			toTypeContactBlockedSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_contactsFound:
		r = &PredContactsFound{
			toTypePeerSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_messagesDialogs:
		r = &PredMessagesDialogs{
			toTypeDialogSlice(m.Vector()),
			toTypeMessageSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_messagesDialogsSlice:
		r = &PredMessagesDialogsSlice{
			m.Int(),
			toTypeDialogSlice(m.Vector()),
			toTypeMessageSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_messagesMessages:
		r = &PredMessagesMessages{
			toTypeMessageSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_messagesMessagesSlice:
		r = &PredMessagesMessagesSlice{
			m.Int(),
			toTypeMessageSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_messagesChats:
		r = &PredMessagesChats{
			toTypeChatSlice(m.Vector()),
		}

	case crc_messagesChatFull:
		r = &PredMessagesChatFull{
			toTypeChatFull(m.Object()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_messagesAffectedHistory:
		r = &PredMessagesAffectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessagesFilterEmpty:
		r = &PredInputMessagesFilterEmpty{}

	case crc_inputMessagesFilterPhotos:
		r = &PredInputMessagesFilterPhotos{}

	case crc_inputMessagesFilterVideo:
		r = &PredInputMessagesFilterVideo{}

	case crc_inputMessagesFilterPhotoVideo:
		r = &PredInputMessagesFilterPhotoVideo{}

	case crc_updateNewMessage:
		r = &PredUpdateNewMessage{
			toTypeMessage(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_updateMessageID:
		r = &PredUpdateMessageID{
			m.Int(),
			m.Long(),
		}

	case crc_updateDeleteMessages:
		r = &PredUpdateDeleteMessages{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateUserTyping:
		r = &PredUpdateUserTyping{
			m.Int(),
			toTypeSendMessageAction(m.Object()),
		}

	case crc_updateChatUserTyping:
		r = &PredUpdateChatUserTyping{
			m.Int(),
			m.Int(),
			toTypeSendMessageAction(m.Object()),
		}

	case crc_updateChatParticipants:
		r = &PredUpdateChatParticipants{
			toTypeChatParticipants(m.Object()),
		}

	case crc_updateUserStatus:
		r = &PredUpdateUserStatus{
			m.Int(),
			toTypeUserStatus(m.Object()),
		}

	case crc_updateUserName:
		r = &PredUpdateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_updateUserPhoto:
		r = &PredUpdateUserPhoto{
			m.Int(),
			m.Int(),
			toTypeUserProfilePhoto(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_updateContactRegistered:
		r = &PredUpdateContactRegistered{
			m.Int(),
			m.Int(),
		}

	case crc_updateContactLink:
		r = &PredUpdateContactLink{
			m.Int(),
			toTypeContactLink(m.Object()),
			toTypeContactLink(m.Object()),
		}

	case crc_updatesState:
		r = &PredUpdatesState{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updatesDifferenceEmpty:
		r = &PredUpdatesDifferenceEmpty{
			m.Int(),
			m.Int(),
		}

	case crc_updatesDifference:
		r = &PredUpdatesDifference{
			toTypeMessageSlice(m.Vector()),
			toTypeEncryptedMessageSlice(m.Vector()),
			toTypeUpdateSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
			toTypeUpdatesState(m.Object()),
		}

	case crc_updatesDifferenceSlice:
		r = &PredUpdatesDifferenceSlice{
			toTypeMessageSlice(m.Vector()),
			toTypeEncryptedMessageSlice(m.Vector()),
			toTypeUpdateSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
			toTypeUpdatesState(m.Object()),
		}

	case crc_updatesTooLong:
		r = &PredUpdatesTooLong{}

	case crc_updateShortMessage:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateShortMessage{
			flags,
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypeMessageFwdHeader(m.FlaggedObject(flags, 2)),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 7)),
		}

	case crc_updateShortChatMessage:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateShortChatMessage{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypeMessageFwdHeader(m.FlaggedObject(flags, 2)),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 7)),
		}

	case crc_updateShort:
		r = &PredUpdateShort{
			toTypeUpdate(m.Object()),
			m.Int(),
		}

	case crc_updatesCombined:
		r = &PredUpdatesCombined{
			toTypeUpdateSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates:
		r = &PredUpdates{
			toTypeUpdateSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			m.Int(),
			m.Int(),
		}

	case crc_photosPhoto:
		r = &PredPhotosPhoto{
			toTypePhoto(m.Object()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_uploadFile:
		r = &PredUploadFile{
			toTypeStorageFileType(m.Object()),
			m.Int(),
			m.StringBytes(),
		}

	case crc_dcOption:
		flags := m.Flags()
		_ = flags
		r = &PredDcOption{
			flags,
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_config:
		flags := m.Flags()
		_ = flags
		r = &PredConfig{
			flags,
			m.Int(),
			m.Int(),
			toTypeBool(m.Object()),
			m.Int(),
			toTypeDcOptionSlice(m.Vector()),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 2),
			m.FlaggedInt(flags, 2),
			toTypeDisabledFeatureSlice(m.Vector()),
		}

	case crc_nearestDc:
		r = &PredNearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_helpAppUpdate:
		r = &PredHelpAppUpdate{
			m.Int(),
			toTypeBool(m.Object()),
			m.String(),
			m.String(),
		}

	case crc_helpNoAppUpdate:
		r = &PredHelpNoAppUpdate{}

	case crc_helpInviteText:
		r = &PredHelpInviteText{
			m.String(),
		}

	case crc_inputPeerNotifyEventsEmpty:
		r = &PredInputPeerNotifyEventsEmpty{}

	case crc_inputPeerNotifyEventsAll:
		r = &PredInputPeerNotifyEventsAll{}

	case crc_photosPhotos:
		r = &PredPhotosPhotos{
			toTypePhotoSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_photosPhotosSlice:
		r = &PredPhotosPhotosSlice{
			m.Int(),
			toTypePhotoSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_wallPaperSolid:
		r = &PredWallPaperSolid{
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_updateNewEncryptedMessage:
		r = &PredUpdateNewEncryptedMessage{
			toTypeEncryptedMessage(m.Object()),
			m.Int(),
		}

	case crc_updateEncryptedChatTyping:
		r = &PredUpdateEncryptedChatTyping{
			m.Int(),
		}

	case crc_updateEncryption:
		r = &PredUpdateEncryption{
			toTypeEncryptedChat(m.Object()),
			m.Int(),
		}

	case crc_updateEncryptedMessagesRead:
		r = &PredUpdateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatEmpty:
		r = &PredEncryptedChatEmpty{
			m.Int(),
		}

	case crc_encryptedChatWaiting:
		r = &PredEncryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatRequested:
		r = &PredEncryptedChatRequested{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_encryptedChat:
		r = &PredEncryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_encryptedChatDiscarded:
		r = &PredEncryptedChatDiscarded{
			m.Int(),
		}

	case crc_inputEncryptedChat:
		r = &PredInputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case crc_encryptedFileEmpty:
		r = &PredEncryptedFileEmpty{}

	case crc_encryptedFile:
		r = &PredEncryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputEncryptedFileEmpty:
		r = &PredInputEncryptedFileEmpty{}

	case crc_inputEncryptedFileUploaded:
		r = &PredInputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_inputEncryptedFile:
		r = &PredInputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case crc_inputEncryptedFileLocation:
		r = &PredInputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_encryptedMessage:
		r = &PredEncryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			toTypeEncryptedFile(m.Object()),
		}

	case crc_encryptedMessageService:
		r = &PredEncryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messagesDhConfigNotModified:
		r = &PredMessagesDhConfigNotModified{
			m.StringBytes(),
		}

	case crc_messagesDhConfig:
		r = &PredMessagesDhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messagesSentEncryptedMessage:
		r = &PredMessagesSentEncryptedMessage{
			m.Int(),
		}

	case crc_messagesSentEncryptedFile:
		r = &PredMessagesSentEncryptedFile{
			m.Int(),
			toTypeEncryptedFile(m.Object()),
		}

	case crc_inputFileBig:
		r = &PredInputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case crc_inputEncryptedFileBigUploaded:
		r = &PredInputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crc_storageFilePdf:
		r = &PredStorageFilePdf{}

	case crc_inputMessagesFilterDocument:
		r = &PredInputMessagesFilterDocument{}

	case crc_inputMessagesFilterPhotoVideoDocuments:
		r = &PredInputMessagesFilterPhotoVideoDocuments{}

	case crc_updateChatParticipantAdd:
		r = &PredUpdateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantDelete:
		r = &PredUpdateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateDcOptions:
		r = &PredUpdateDcOptions{
			toTypeDcOptionSlice(m.Vector()),
		}

	case crc_inputMediaUploadedDocument:
		flags := m.Flags()
		_ = flags
		r = &PredInputMediaUploadedDocument{
			flags,
			toTypeInputFile(m.Object()),
			toTypeInputFile(m.FlaggedObject(flags, 2)),
			m.String(),
			toTypeDocumentAttributeSlice(m.Vector()),
			m.String(),
			toTypeInputDocumentSlice(m.FlaggedVector(flags, 0)),
			m.FlaggedInt(flags, 1),
		}

	case crc_inputMediaDocument:
		flags := m.Flags()
		_ = flags
		r = &PredInputMediaDocument{
			flags,
			toTypeInputDocument(m.Object()),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_messageMediaDocument:
		flags := m.Flags()
		_ = flags
		r = &PredMessageMediaDocument{
			flags,
			toTypeDocument(m.FlaggedObject(flags, 0)),
			m.FlaggedString(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crc_inputDocumentEmpty:
		r = &PredInputDocumentEmpty{}

	case crc_inputDocument:
		r = &PredInputDocument{
			m.Long(),
			m.Long(),
		}

	case crc_inputDocumentFileLocation:
		r = &PredInputDocumentFileLocation{
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crc_documentEmpty:
		r = &PredDocumentEmpty{
			m.Long(),
		}

	case crc_document:
		r = &PredDocument{
			m.Long(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
			toTypePhotoSize(m.Object()),
			m.Int(),
			m.Int(),
			toTypeDocumentAttributeSlice(m.Vector()),
		}

	case crc_helpSupport:
		r = &PredHelpSupport{
			m.String(),
			toTypeUser(m.Object()),
		}

	case crc_notifyAll:
		r = &PredNotifyAll{}

	case crc_notifyChats:
		r = &PredNotifyChats{}

	case crc_notifyPeer:
		r = &PredNotifyPeer{
			toTypePeer(m.Object()),
		}

	case crc_notifyUsers:
		r = &PredNotifyUsers{}

	case crc_updateUserBlocked:
		r = &PredUpdateUserBlocked{
			m.Int(),
			toTypeBool(m.Object()),
		}

	case crc_updateNotifySettings:
		r = &PredUpdateNotifySettings{
			toTypeNotifyPeer(m.Object()),
			toTypePeerNotifySettings(m.Object()),
		}

	case crc_sendMessageTypingAction:
		r = &PredSendMessageTypingAction{}

	case crc_sendMessageCancelAction:
		r = &PredSendMessageCancelAction{}

	case crc_sendMessageRecordVideoAction:
		r = &PredSendMessageRecordVideoAction{}

	case crc_sendMessageUploadVideoAction:
		r = &PredSendMessageUploadVideoAction{
			m.Int(),
		}

	case crc_sendMessageRecordAudioAction:
		r = &PredSendMessageRecordAudioAction{}

	case crc_sendMessageUploadAudioAction:
		r = &PredSendMessageUploadAudioAction{
			m.Int(),
		}

	case crc_sendMessageUploadPhotoAction:
		r = &PredSendMessageUploadPhotoAction{
			m.Int(),
		}

	case crc_sendMessageUploadDocumentAction:
		r = &PredSendMessageUploadDocumentAction{
			m.Int(),
		}

	case crc_sendMessageGeoLocationAction:
		r = &PredSendMessageGeoLocationAction{}

	case crc_sendMessageChooseContactAction:
		r = &PredSendMessageChooseContactAction{}

	case crc_updateServiceNotification:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateServiceNotification{
			flags,
			m.FlaggedInt(flags, 1),
			m.String(),
			m.String(),
			toTypeMessageMedia(m.Object()),
			toTypeMessageEntitySlice(m.Vector()),
		}

	case crc_userStatusRecently:
		r = &PredUserStatusRecently{}

	case crc_userStatusLastWeek:
		r = &PredUserStatusLastWeek{}

	case crc_userStatusLastMonth:
		r = &PredUserStatusLastMonth{}

	case crc_updatePrivacy:
		r = &PredUpdatePrivacy{
			toTypePrivacyKey(m.Object()),
			toTypePrivacyRuleSlice(m.Vector()),
		}

	case crc_inputPrivacyKeyStatusTimestamp:
		r = &PredInputPrivacyKeyStatusTimestamp{}

	case crc_privacyKeyStatusTimestamp:
		r = &PredPrivacyKeyStatusTimestamp{}

	case crc_inputPrivacyValueAllowContacts:
		r = &PredInputPrivacyValueAllowContacts{}

	case crc_inputPrivacyValueAllowAll:
		r = &PredInputPrivacyValueAllowAll{}

	case crc_inputPrivacyValueAllowUsers:
		r = &PredInputPrivacyValueAllowUsers{
			toTypeInputUserSlice(m.Vector()),
		}

	case crc_inputPrivacyValueDisallowContacts:
		r = &PredInputPrivacyValueDisallowContacts{}

	case crc_inputPrivacyValueDisallowAll:
		r = &PredInputPrivacyValueDisallowAll{}

	case crc_inputPrivacyValueDisallowUsers:
		r = &PredInputPrivacyValueDisallowUsers{
			toTypeInputUserSlice(m.Vector()),
		}

	case crc_privacyValueAllowContacts:
		r = &PredPrivacyValueAllowContacts{}

	case crc_privacyValueAllowAll:
		r = &PredPrivacyValueAllowAll{}

	case crc_privacyValueAllowUsers:
		r = &PredPrivacyValueAllowUsers{
			m.VectorInt(),
		}

	case crc_privacyValueDisallowContacts:
		r = &PredPrivacyValueDisallowContacts{}

	case crc_privacyValueDisallowAll:
		r = &PredPrivacyValueDisallowAll{}

	case crc_privacyValueDisallowUsers:
		r = &PredPrivacyValueDisallowUsers{
			m.VectorInt(),
		}

	case crc_accountPrivacyRules:
		r = &PredAccountPrivacyRules{
			toTypePrivacyRuleSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_accountDaysTTL:
		r = &PredAccountDaysTTL{
			m.Int(),
		}

	case crc_updateUserPhone:
		r = &PredUpdateUserPhone{
			m.Int(),
			m.String(),
		}

	case crc_disabledFeature:
		r = &PredDisabledFeature{
			m.String(),
			m.String(),
		}

	case crc_documentAttributeImageSize:
		r = &PredDocumentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAnimated:
		r = &PredDocumentAttributeAnimated{}

	case crc_documentAttributeSticker:
		flags := m.Flags()
		_ = flags
		r = &PredDocumentAttributeSticker{
			flags,
			m.String(),
			toTypeInputStickerSet(m.Object()),
			toTypeMaskCoords(m.FlaggedObject(flags, 0)),
		}

	case crc_documentAttributeVideo:
		flags := m.Flags()
		_ = flags
		r = &PredDocumentAttributeVideo{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAudio:
		flags := m.Flags()
		_ = flags
		r = &PredDocumentAttributeAudio{
			flags,
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedStringBytes(flags, 2),
		}

	case crc_documentAttributeFilename:
		r = &PredDocumentAttributeFilename{
			m.String(),
		}

	case crc_messagesStickersNotModified:
		r = &PredMessagesStickersNotModified{}

	case crc_messagesStickers:
		r = &PredMessagesStickers{
			m.String(),
			toTypeDocumentSlice(m.Vector()),
		}

	case crc_stickerPack:
		r = &PredStickerPack{
			m.String(),
			m.VectorLong(),
		}

	case crc_messagesAllStickersNotModified:
		r = &PredMessagesAllStickersNotModified{}

	case crc_messagesAllStickers:
		r = &PredMessagesAllStickers{
			m.Int(),
			toTypeStickerSetSlice(m.Vector()),
		}

	case crc_accountNoPassword:
		r = &PredAccountNoPassword{
			m.StringBytes(),
			m.String(),
		}

	case crc_accountPassword:
		r = &PredAccountPassword{
			m.StringBytes(),
			m.StringBytes(),
			m.String(),
			toTypeBool(m.Object()),
			m.String(),
		}

	case crc_updateReadHistoryInbox:
		r = &PredUpdateReadHistoryInbox{
			toTypePeer(m.Object()),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadHistoryOutbox:
		r = &PredUpdateReadHistoryOutbox{
			toTypePeer(m.Object()),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messagesAffectedMessages:
		r = &PredMessagesAffectedMessages{
			m.Int(),
			m.Int(),
		}

	case crc_contactLinkUnknown:
		r = &PredContactLinkUnknown{}

	case crc_contactLinkNone:
		r = &PredContactLinkNone{}

	case crc_contactLinkHasPhone:
		r = &PredContactLinkHasPhone{}

	case crc_contactLinkContact:
		r = &PredContactLinkContact{}

	case crc_updateWebPage:
		r = &PredUpdateWebPage{
			toTypeWebPage(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_webPageEmpty:
		r = &PredWebPageEmpty{
			m.Long(),
		}

	case crc_webPagePending:
		r = &PredWebPagePending{
			m.Long(),
			m.Int(),
		}

	case crc_webPage:
		flags := m.Flags()
		_ = flags
		r = &PredWebPage{
			flags,
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			toTypePhoto(m.FlaggedObject(flags, 4)),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			m.FlaggedString(flags, 8),
			toTypeDocument(m.FlaggedObject(flags, 9)),
			toTypePage(m.FlaggedObject(flags, 10)),
		}

	case crc_messageMediaWebPage:
		r = &PredMessageMediaWebPage{
			toTypeWebPage(m.Object()),
		}

	case crc_authorization:
		r = &PredAuthorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_accountAuthorizations:
		r = &PredAccountAuthorizations{
			toTypeAuthorizationSlice(m.Vector()),
		}

	case crc_accountPasswordSettings:
		r = &PredAccountPasswordSettings{
			m.String(),
		}

	case crc_accountPasswordInputSettings:
		flags := m.Flags()
		_ = flags
		r = &PredAccountPasswordInputSettings{
			flags,
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_authPasswordRecovery:
		r = &PredAuthPasswordRecovery{
			m.String(),
		}

	case crc_inputMediaVenue:
		r = &PredInputMediaVenue{
			toTypeInputGeoPoint(m.Object()),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messageMediaVenue:
		r = &PredMessageMediaVenue{
			toTypeGeoPoint(m.Object()),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_receivedNotifyMessage:
		r = &PredReceivedNotifyMessage{
			m.Int(),
			m.Int(),
		}

	case crc_chatInviteEmpty:
		r = &PredChatInviteEmpty{}

	case crc_chatInviteExported:
		r = &PredChatInviteExported{
			m.String(),
		}

	case crc_chatInviteAlready:
		r = &PredChatInviteAlready{
			toTypeChat(m.Object()),
		}

	case crc_chatInvite:
		flags := m.Flags()
		_ = flags
		r = &PredChatInvite{
			flags,
			m.String(),
			toTypeChatPhoto(m.Object()),
			m.Int(),
			toTypeUserSlice(m.FlaggedVector(flags, 4)),
		}

	case crc_messageActionChatJoinedByLink:
		r = &PredMessageActionChatJoinedByLink{
			m.Int(),
		}

	case crc_updateReadMessagesContents:
		r = &PredUpdateReadMessagesContents{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_inputStickerSetEmpty:
		r = &PredInputStickerSetEmpty{}

	case crc_inputStickerSetID:
		r = &PredInputStickerSetID{
			m.Long(),
			m.Long(),
		}

	case crc_inputStickerSetShortName:
		r = &PredInputStickerSetShortName{
			m.String(),
		}

	case crc_stickerSet:
		flags := m.Flags()
		_ = flags
		r = &PredStickerSet{
			flags,
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_messagesStickerSet:
		r = &PredMessagesStickerSet{
			toTypeStickerSet(m.Object()),
			toTypeStickerPackSlice(m.Vector()),
			toTypeDocumentSlice(m.Vector()),
		}

	case crc_user:
		flags := m.Flags()
		_ = flags
		r = &PredUser{
			flags,
			m.Int(),
			m.FlaggedLong(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			toTypeUserProfilePhoto(m.FlaggedObject(flags, 5)),
			toTypeUserStatus(m.FlaggedObject(flags, 6)),
			m.FlaggedInt(flags, 14),
			m.FlaggedString(flags, 18),
			m.FlaggedString(flags, 19),
			m.FlaggedString(flags, 22),
		}

	case crc_botCommand:
		r = &PredBotCommand{
			m.String(),
			m.String(),
		}

	case crc_botInfo:
		r = &PredBotInfo{
			m.Int(),
			m.String(),
			toTypeBotCommandSlice(m.Vector()),
		}

	case crc_keyboardButton:
		r = &PredKeyboardButton{
			m.String(),
		}

	case crc_keyboardButtonRow:
		r = &PredKeyboardButtonRow{
			toTypeKeyboardButtonSlice(m.Vector()),
		}

	case crc_replyKeyboardHide:
		flags := m.Flags()
		_ = flags
		r = &PredReplyKeyboardHide{
			flags,
		}

	case crc_replyKeyboardForceReply:
		flags := m.Flags()
		_ = flags
		r = &PredReplyKeyboardForceReply{
			flags,
		}

	case crc_replyKeyboardMarkup:
		flags := m.Flags()
		_ = flags
		r = &PredReplyKeyboardMarkup{
			flags,
			toTypeKeyboardButtonRowSlice(m.Vector()),
		}

	case crc_inputMessagesFilterUrl:
		r = &PredInputMessagesFilterUrl{}

	case crc_inputPeerUser:
		r = &PredInputPeerUser{
			m.Int(),
			m.Long(),
		}

	case crc_inputUser:
		r = &PredInputUser{
			m.Int(),
			m.Long(),
		}

	case crc_messageEntityUnknown:
		r = &PredMessageEntityUnknown{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityMention:
		r = &PredMessageEntityMention{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityHashtag:
		r = &PredMessageEntityHashtag{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBotCommand:
		r = &PredMessageEntityBotCommand{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityUrl:
		r = &PredMessageEntityUrl{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityEmail:
		r = &PredMessageEntityEmail{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBold:
		r = &PredMessageEntityBold{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityItalic:
		r = &PredMessageEntityItalic{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityCode:
		r = &PredMessageEntityCode{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityPre:
		r = &PredMessageEntityPre{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_messageEntityTextUrl:
		r = &PredMessageEntityTextUrl{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_updateShortSentMessage:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateShortSentMessage{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypeMessageMedia(m.FlaggedObject(flags, 9)),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 7)),
		}

	case crc_inputPeerChannel:
		r = &PredInputPeerChannel{
			m.Int(),
			m.Long(),
		}

	case crc_peerChannel:
		r = &PredPeerChannel{
			m.Int(),
		}

	case crc_channel:
		flags := m.Flags()
		_ = flags
		r = &PredChannel{
			flags,
			m.Int(),
			m.FlaggedLong(flags, 13),
			m.String(),
			m.FlaggedString(flags, 6),
			toTypeChatPhoto(m.Object()),
			m.Int(),
			m.Int(),
			m.FlaggedString(flags, 9),
			toTypeChannelAdminRights(m.FlaggedObject(flags, 14)),
			toTypeChannelBannedRights(m.FlaggedObject(flags, 15)),
		}

	case crc_channelForbidden:
		flags := m.Flags()
		_ = flags
		r = &PredChannelForbidden{
			flags,
			m.Int(),
			m.Long(),
			m.String(),
			m.FlaggedInt(flags, 16),
		}

	case crc_channelFull:
		flags := m.Flags()
		_ = flags
		r = &PredChannelFull{
			flags,
			m.Int(),
			m.String(),
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedInt(flags, 2),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypePhoto(m.Object()),
			toTypePeerNotifySettings(m.Object()),
			toTypeExportedChatInvite(m.Object()),
			toTypeBotInfoSlice(m.Vector()),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 5),
			toTypeStickerSet(m.FlaggedObject(flags, 8)),
		}

	case crc_messageActionChannelCreate:
		r = &PredMessageActionChannelCreate{
			m.String(),
		}

	case crc_messagesChannelMessages:
		flags := m.Flags()
		_ = flags
		r = &PredMessagesChannelMessages{
			flags,
			m.Int(),
			m.Int(),
			toTypeMessageSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_updateChannelTooLong:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateChannelTooLong{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 0),
		}

	case crc_updateChannel:
		r = &PredUpdateChannel{
			m.Int(),
		}

	case crc_updateNewChannelMessage:
		r = &PredUpdateNewChannelMessage{
			toTypeMessage(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadChannelInbox:
		r = &PredUpdateReadChannelInbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDeleteChannelMessages:
		r = &PredUpdateDeleteChannelMessages{
			m.Int(),
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChannelMessageViews:
		r = &PredUpdateChannelMessageViews{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputChannelEmpty:
		r = &PredInputChannelEmpty{}

	case crc_inputChannel:
		r = &PredInputChannel{
			m.Int(),
			m.Long(),
		}

	case crc_contactsResolvedPeer:
		r = &PredContactsResolvedPeer{
			toTypePeer(m.Object()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_messageRange:
		r = &PredMessageRange{
			m.Int(),
			m.Int(),
		}

	case crc_updatesChannelDifferenceEmpty:
		flags := m.Flags()
		_ = flags
		r = &PredUpdatesChannelDifferenceEmpty{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
		}

	case crc_updatesChannelDifferenceTooLong:
		flags := m.Flags()
		_ = flags
		r = &PredUpdatesChannelDifferenceTooLong{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypeMessageSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_updatesChannelDifference:
		flags := m.Flags()
		_ = flags
		r = &PredUpdatesChannelDifference{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
			toTypeMessageSlice(m.Vector()),
			toTypeUpdateSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_channelMessagesFilterEmpty:
		r = &PredChannelMessagesFilterEmpty{}

	case crc_channelMessagesFilter:
		flags := m.Flags()
		_ = flags
		r = &PredChannelMessagesFilter{
			flags,
			toTypeMessageRangeSlice(m.Vector()),
		}

	case crc_channelParticipant:
		r = &PredChannelParticipant{
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantSelf:
		r = &PredChannelParticipantSelf{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantCreator:
		r = &PredChannelParticipantCreator{
			m.Int(),
		}

	case crc_channelParticipantsRecent:
		r = &PredChannelParticipantsRecent{}

	case crc_channelParticipantsAdmins:
		r = &PredChannelParticipantsAdmins{}

	case crc_channelParticipantsKicked:
		r = &PredChannelParticipantsKicked{
			m.String(),
		}

	case crc_channelsChannelParticipants:
		r = &PredChannelsChannelParticipants{
			m.Int(),
			toTypeChannelParticipantSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_channelsChannelParticipant:
		r = &PredChannelsChannelParticipant{
			toTypeChannelParticipant(m.Object()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_true:
		r = &PredTrue{}

	case crc_chatParticipantCreator:
		r = &PredChatParticipantCreator{
			m.Int(),
		}

	case crc_chatParticipantAdmin:
		r = &PredChatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatAdmins:
		r = &PredUpdateChatAdmins{
			m.Int(),
			toTypeBool(m.Object()),
			m.Int(),
		}

	case crc_updateChatParticipantAdmin:
		r = &PredUpdateChatParticipantAdmin{
			m.Int(),
			m.Int(),
			toTypeBool(m.Object()),
			m.Int(),
		}

	case crc_messageActionChatMigrateTo:
		r = &PredMessageActionChatMigrateTo{
			m.Int(),
		}

	case crc_messageActionChannelMigrateFrom:
		r = &PredMessageActionChannelMigrateFrom{
			m.String(),
			m.Int(),
		}

	case crc_channelParticipantsBots:
		r = &PredChannelParticipantsBots{}

	case crc_inputReportReasonSpam:
		r = &PredInputReportReasonSpam{}

	case crc_inputReportReasonViolence:
		r = &PredInputReportReasonViolence{}

	case crc_inputReportReasonPornography:
		r = &PredInputReportReasonPornography{}

	case crc_inputReportReasonOther:
		r = &PredInputReportReasonOther{
			m.String(),
		}

	case crc_updateNewStickerSet:
		r = &PredUpdateNewStickerSet{
			toTypeMessagesStickerSet(m.Object()),
		}

	case crc_updateStickerSetsOrder:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateStickerSetsOrder{
			flags,
			m.VectorLong(),
		}

	case crc_updateStickerSets:
		r = &PredUpdateStickerSets{}

	case crc_helpTermsOfService:
		r = &PredHelpTermsOfService{
			m.String(),
		}

	case crc_foundGif:
		r = &PredFoundGif{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMediaGifExternal:
		r = &PredInputMediaGifExternal{
			m.String(),
			m.String(),
		}

	case crc_messagesFoundGifs:
		r = &PredMessagesFoundGifs{
			m.Int(),
			toTypeFoundGifSlice(m.Vector()),
		}

	case crc_inputMessagesFilterGif:
		r = &PredInputMessagesFilterGif{}

	case crc_updateSavedGifs:
		r = &PredUpdateSavedGifs{}

	case crc_updateBotInlineQuery:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateBotInlineQuery{
			flags,
			m.Long(),
			m.Int(),
			m.String(),
			toTypeGeoPoint(m.FlaggedObject(flags, 0)),
			m.String(),
		}

	case crc_foundGifCached:
		r = &PredFoundGifCached{
			m.String(),
			toTypePhoto(m.Object()),
			toTypeDocument(m.Object()),
		}

	case crc_messagesSavedGifsNotModified:
		r = &PredMessagesSavedGifsNotModified{}

	case crc_messagesSavedGifs:
		r = &PredMessagesSavedGifs{
			m.Int(),
			toTypeDocumentSlice(m.Vector()),
		}

	case crc_inputBotInlineMessageMediaAuto:
		flags := m.Flags()
		_ = flags
		r = &PredInputBotInlineMessageMediaAuto{
			flags,
			m.String(),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_inputBotInlineMessageText:
		flags := m.Flags()
		_ = flags
		r = &PredInputBotInlineMessageText{
			flags,
			m.String(),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 1)),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_inputBotInlineResult:
		flags := m.Flags()
		_ = flags
		r = &PredInputBotInlineResult{
			flags,
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			toTypeInputBotInlineMessage(m.Object()),
		}

	case crc_botInlineMessageMediaAuto:
		flags := m.Flags()
		_ = flags
		r = &PredBotInlineMessageMediaAuto{
			flags,
			m.String(),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_botInlineMessageText:
		flags := m.Flags()
		_ = flags
		r = &PredBotInlineMessageText{
			flags,
			m.String(),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 1)),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_botInlineResult:
		flags := m.Flags()
		_ = flags
		r = &PredBotInlineResult{
			flags,
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			toTypeBotInlineMessage(m.Object()),
		}

	case crc_messagesBotResults:
		flags := m.Flags()
		_ = flags
		r = &PredMessagesBotResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 1),
			toTypeInlineBotSwitchPM(m.FlaggedObject(flags, 2)),
			toTypeBotInlineResultSlice(m.Vector()),
			m.Int(),
		}

	case crc_inputMessagesFilterVoice:
		r = &PredInputMessagesFilterVoice{}

	case crc_inputMessagesFilterMusic:
		r = &PredInputMessagesFilterMusic{}

	case crc_updateBotInlineSend:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateBotInlineSend{
			flags,
			m.Int(),
			m.String(),
			toTypeGeoPoint(m.FlaggedObject(flags, 0)),
			m.String(),
			toTypeInputBotInlineMessageID(m.FlaggedObject(flags, 1)),
		}

	case crc_inputPrivacyKeyChatInvite:
		r = &PredInputPrivacyKeyChatInvite{}

	case crc_privacyKeyChatInvite:
		r = &PredPrivacyKeyChatInvite{}

	case crc_updateEditChannelMessage:
		r = &PredUpdateEditChannelMessage{
			toTypeMessage(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_exportedMessageLink:
		r = &PredExportedMessageLink{
			m.String(),
		}

	case crc_messageFwdHeader:
		flags := m.Flags()
		_ = flags
		r = &PredMessageFwdHeader{
			flags,
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
		}

	case crc_messageActionPinMessage:
		r = &PredMessageActionPinMessage{}

	case crc_peerSettings:
		flags := m.Flags()
		_ = flags
		r = &PredPeerSettings{
			flags,
		}

	case crc_updateChannelPinnedMessage:
		r = &PredUpdateChannelPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case crc_keyboardButtonUrl:
		r = &PredKeyboardButtonUrl{
			m.String(),
			m.String(),
		}

	case crc_keyboardButtonCallback:
		r = &PredKeyboardButtonCallback{
			m.String(),
			m.StringBytes(),
		}

	case crc_keyboardButtonRequestPhone:
		r = &PredKeyboardButtonRequestPhone{
			m.String(),
		}

	case crc_keyboardButtonRequestGeoLocation:
		r = &PredKeyboardButtonRequestGeoLocation{
			m.String(),
		}

	case crc_authCodeTypeSms:
		r = &PredAuthCodeTypeSms{}

	case crc_authCodeTypeCall:
		r = &PredAuthCodeTypeCall{}

	case crc_authCodeTypeFlashCall:
		r = &PredAuthCodeTypeFlashCall{}

	case crc_authSentCodeTypeApp:
		r = &PredAuthSentCodeTypeApp{
			m.Int(),
		}

	case crc_authSentCodeTypeSms:
		r = &PredAuthSentCodeTypeSms{
			m.Int(),
		}

	case crc_authSentCodeTypeCall:
		r = &PredAuthSentCodeTypeCall{
			m.Int(),
		}

	case crc_authSentCodeTypeFlashCall:
		r = &PredAuthSentCodeTypeFlashCall{
			m.String(),
		}

	case crc_keyboardButtonSwitchInline:
		flags := m.Flags()
		_ = flags
		r = &PredKeyboardButtonSwitchInline{
			flags,
			m.String(),
			m.String(),
		}

	case crc_replyInlineMarkup:
		r = &PredReplyInlineMarkup{
			toTypeKeyboardButtonRowSlice(m.Vector()),
		}

	case crc_messagesBotCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = &PredMessagesBotCallbackAnswer{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case crc_updateBotCallbackQuery:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateBotCallbackQuery{
			flags,
			m.Long(),
			m.Int(),
			toTypePeer(m.Object()),
			m.Int(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_messagesMessageEditData:
		flags := m.Flags()
		_ = flags
		r = &PredMessagesMessageEditData{
			flags,
		}

	case crc_updateEditMessage:
		r = &PredUpdateEditMessage{
			toTypeMessage(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_inputBotInlineMessageMediaGeo:
		flags := m.Flags()
		_ = flags
		r = &PredInputBotInlineMessageMediaGeo{
			flags,
			toTypeInputGeoPoint(m.Object()),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_inputBotInlineMessageMediaVenue:
		flags := m.Flags()
		_ = flags
		r = &PredInputBotInlineMessageMediaVenue{
			flags,
			toTypeInputGeoPoint(m.Object()),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_inputBotInlineMessageMediaContact:
		flags := m.Flags()
		_ = flags
		r = &PredInputBotInlineMessageMediaContact{
			flags,
			m.String(),
			m.String(),
			m.String(),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_botInlineMessageMediaGeo:
		flags := m.Flags()
		_ = flags
		r = &PredBotInlineMessageMediaGeo{
			flags,
			toTypeGeoPoint(m.Object()),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_botInlineMessageMediaVenue:
		flags := m.Flags()
		_ = flags
		r = &PredBotInlineMessageMediaVenue{
			flags,
			toTypeGeoPoint(m.Object()),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_botInlineMessageMediaContact:
		flags := m.Flags()
		_ = flags
		r = &PredBotInlineMessageMediaContact{
			flags,
			m.String(),
			m.String(),
			m.String(),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_inputBotInlineResultPhoto:
		r = &PredInputBotInlineResultPhoto{
			m.String(),
			m.String(),
			toTypeInputPhoto(m.Object()),
			toTypeInputBotInlineMessage(m.Object()),
		}

	case crc_inputBotInlineResultDocument:
		flags := m.Flags()
		_ = flags
		r = &PredInputBotInlineResultDocument{
			flags,
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			toTypeInputDocument(m.Object()),
			toTypeInputBotInlineMessage(m.Object()),
		}

	case crc_botInlineMediaResult:
		flags := m.Flags()
		_ = flags
		r = &PredBotInlineMediaResult{
			flags,
			m.String(),
			m.String(),
			toTypePhoto(m.FlaggedObject(flags, 0)),
			toTypeDocument(m.FlaggedObject(flags, 1)),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			toTypeBotInlineMessage(m.Object()),
		}

	case crc_inputBotInlineMessageID:
		r = &PredInputBotInlineMessageID{
			m.Int(),
			m.Long(),
			m.Long(),
		}

	case crc_updateInlineBotCallbackQuery:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateInlineBotCallbackQuery{
			flags,
			m.Long(),
			m.Int(),
			toTypeInputBotInlineMessageID(m.Object()),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_inlineBotSwitchPM:
		r = &PredInlineBotSwitchPM{
			m.String(),
			m.String(),
		}

	case crc_messageEntityMentionName:
		r = &PredMessageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessageEntityMentionName:
		r = &PredInputMessageEntityMentionName{
			m.Int(),
			m.Int(),
			toTypeInputUser(m.Object()),
		}

	case crc_messagesPeerDialogs:
		r = &PredMessagesPeerDialogs{
			toTypeDialogSlice(m.Vector()),
			toTypeMessageSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
			toTypeUpdatesState(m.Object()),
		}

	case crc_topPeer:
		r = &PredTopPeer{
			toTypePeer(m.Object()),
			m.Double(),
		}

	case crc_topPeerCategoryBotsPM:
		r = &PredTopPeerCategoryBotsPM{}

	case crc_topPeerCategoryBotsInline:
		r = &PredTopPeerCategoryBotsInline{}

	case crc_topPeerCategoryCorrespondents:
		r = &PredTopPeerCategoryCorrespondents{}

	case crc_topPeerCategoryGroups:
		r = &PredTopPeerCategoryGroups{}

	case crc_topPeerCategoryChannels:
		r = &PredTopPeerCategoryChannels{}

	case crc_topPeerCategoryPeers:
		r = &PredTopPeerCategoryPeers{
			toTypeTopPeerCategory(m.Object()),
			m.Int(),
			toTypeTopPeerSlice(m.Vector()),
		}

	case crc_contactsTopPeersNotModified:
		r = &PredContactsTopPeersNotModified{}

	case crc_contactsTopPeers:
		r = &PredContactsTopPeers{
			toTypeTopPeerCategoryPeersSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_inputMessagesFilterChatPhotos:
		r = &PredInputMessagesFilterChatPhotos{}

	case crc_updateReadChannelOutbox:
		r = &PredUpdateReadChannelOutbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDraftMessage:
		r = &PredUpdateDraftMessage{
			toTypePeer(m.Object()),
			toTypeDraftMessage(m.Object()),
		}

	case crc_draftMessageEmpty:
		r = &PredDraftMessageEmpty{}

	case crc_draftMessage:
		flags := m.Flags()
		_ = flags
		r = &PredDraftMessage{
			flags,
			m.FlaggedInt(flags, 0),
			m.String(),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 3)),
			m.Int(),
		}

	case crc_messageActionHistoryClear:
		r = &PredMessageActionHistoryClear{}

	case crc_updateReadFeaturedStickers:
		r = &PredUpdateReadFeaturedStickers{}

	case crc_updateRecentStickers:
		r = &PredUpdateRecentStickers{}

	case crc_messagesFeaturedStickersNotModified:
		r = &PredMessagesFeaturedStickersNotModified{}

	case crc_messagesFeaturedStickers:
		r = &PredMessagesFeaturedStickers{
			m.Int(),
			toTypeStickerSetCoveredSlice(m.Vector()),
			m.VectorLong(),
		}

	case crc_messagesRecentStickersNotModified:
		r = &PredMessagesRecentStickersNotModified{}

	case crc_messagesRecentStickers:
		r = &PredMessagesRecentStickers{
			m.Int(),
			toTypeDocumentSlice(m.Vector()),
		}

	case crc_messagesArchivedStickers:
		r = &PredMessagesArchivedStickers{
			m.Int(),
			toTypeStickerSetCoveredSlice(m.Vector()),
		}

	case crc_messagesStickerSetInstallResultSuccess:
		r = &PredMessagesStickerSetInstallResultSuccess{}

	case crc_messagesStickerSetInstallResultArchive:
		r = &PredMessagesStickerSetInstallResultArchive{
			toTypeStickerSetCoveredSlice(m.Vector()),
		}

	case crc_stickerSetCovered:
		r = &PredStickerSetCovered{
			toTypeStickerSet(m.Object()),
			toTypeDocument(m.Object()),
		}

	case crc_inputMediaPhotoExternal:
		flags := m.Flags()
		_ = flags
		r = &PredInputMediaPhotoExternal{
			flags,
			m.String(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_inputMediaDocumentExternal:
		flags := m.Flags()
		_ = flags
		r = &PredInputMediaDocumentExternal{
			flags,
			m.String(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_updateConfig:
		r = &PredUpdateConfig{}

	case crc_updatePtsChanged:
		r = &PredUpdatePtsChanged{}

	case crc_messageActionGameScore:
		r = &PredMessageActionGameScore{
			m.Long(),
			m.Int(),
		}

	case crc_documentAttributeHasStickers:
		r = &PredDocumentAttributeHasStickers{}

	case crc_keyboardButtonGame:
		r = &PredKeyboardButtonGame{
			m.String(),
		}

	case crc_stickerSetMultiCovered:
		r = &PredStickerSetMultiCovered{
			toTypeStickerSet(m.Object()),
			toTypeDocumentSlice(m.Vector()),
		}

	case crc_maskCoords:
		r = &PredMaskCoords{
			m.Int(),
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case crc_inputStickeredMediaPhoto:
		r = &PredInputStickeredMediaPhoto{
			toTypeInputPhoto(m.Object()),
		}

	case crc_inputStickeredMediaDocument:
		r = &PredInputStickeredMediaDocument{
			toTypeInputDocument(m.Object()),
		}

	case crc_inputMediaGame:
		r = &PredInputMediaGame{
			toTypeInputGame(m.Object()),
		}

	case crc_messageMediaGame:
		r = &PredMessageMediaGame{
			toTypeGame(m.Object()),
		}

	case crc_inputBotInlineMessageGame:
		flags := m.Flags()
		_ = flags
		r = &PredInputBotInlineMessageGame{
			flags,
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_inputBotInlineResultGame:
		r = &PredInputBotInlineResultGame{
			m.String(),
			m.String(),
			toTypeInputBotInlineMessage(m.Object()),
		}

	case crc_game:
		flags := m.Flags()
		_ = flags
		r = &PredGame{
			flags,
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			toTypePhoto(m.Object()),
			toTypeDocument(m.FlaggedObject(flags, 0)),
		}

	case crc_inputGameID:
		r = &PredInputGameID{
			m.Long(),
			m.Long(),
		}

	case crc_inputGameShortName:
		r = &PredInputGameShortName{
			toTypeInputUser(m.Object()),
			m.String(),
		}

	case crc_highScore:
		r = &PredHighScore{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messagesHighScores:
		r = &PredMessagesHighScores{
			toTypeHighScoreSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_messagesChatsSlice:
		r = &PredMessagesChatsSlice{
			m.Int(),
			toTypeChatSlice(m.Vector()),
		}

	case crc_updateChannelWebPage:
		r = &PredUpdateChannelWebPage{
			m.Int(),
			toTypeWebPage(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_updatesDifferenceTooLong:
		r = &PredUpdatesDifferenceTooLong{
			m.Int(),
		}

	case crc_sendMessageGamePlayAction:
		r = &PredSendMessageGamePlayAction{}

	case crc_webPageNotModified:
		r = &PredWebPageNotModified{}

	case crc_textEmpty:
		r = &PredTextEmpty{}

	case crc_textPlain:
		r = &PredTextPlain{
			m.String(),
		}

	case crc_textBold:
		r = &PredTextBold{
			toTypeRichText(m.Object()),
		}

	case crc_textItalic:
		r = &PredTextItalic{
			toTypeRichText(m.Object()),
		}

	case crc_textUnderline:
		r = &PredTextUnderline{
			toTypeRichText(m.Object()),
		}

	case crc_textStrike:
		r = &PredTextStrike{
			toTypeRichText(m.Object()),
		}

	case crc_textFixed:
		r = &PredTextFixed{
			toTypeRichText(m.Object()),
		}

	case crc_textUrl:
		r = &PredTextUrl{
			toTypeRichText(m.Object()),
			m.String(),
			m.Long(),
		}

	case crc_textEmail:
		r = &PredTextEmail{
			toTypeRichText(m.Object()),
			m.String(),
		}

	case crc_textConcat:
		r = &PredTextConcat{
			toTypeRichTextSlice(m.Vector()),
		}

	case crc_pageBlockTitle:
		r = &PredPageBlockTitle{
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockSubtitle:
		r = &PredPageBlockSubtitle{
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockAuthorDate:
		r = &PredPageBlockAuthorDate{
			toTypeRichText(m.Object()),
			m.Int(),
		}

	case crc_pageBlockHeader:
		r = &PredPageBlockHeader{
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockSubheader:
		r = &PredPageBlockSubheader{
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockParagraph:
		r = &PredPageBlockParagraph{
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockPreformatted:
		r = &PredPageBlockPreformatted{
			toTypeRichText(m.Object()),
			m.String(),
		}

	case crc_pageBlockFooter:
		r = &PredPageBlockFooter{
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockDivider:
		r = &PredPageBlockDivider{}

	case crc_pageBlockList:
		r = &PredPageBlockList{
			toTypeBool(m.Object()),
			toTypeRichTextSlice(m.Vector()),
		}

	case crc_pageBlockBlockquote:
		r = &PredPageBlockBlockquote{
			toTypeRichText(m.Object()),
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockPullquote:
		r = &PredPageBlockPullquote{
			toTypeRichText(m.Object()),
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockPhoto:
		r = &PredPageBlockPhoto{
			m.Long(),
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockVideo:
		flags := m.Flags()
		_ = flags
		r = &PredPageBlockVideo{
			flags,
			m.Long(),
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockCover:
		r = &PredPageBlockCover{
			toTypePageBlock(m.Object()),
		}

	case crc_pageBlockEmbed:
		flags := m.Flags()
		_ = flags
		r = &PredPageBlockEmbed{
			flags,
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedLong(flags, 4),
			m.Int(),
			m.Int(),
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockEmbedPost:
		r = &PredPageBlockEmbedPost{
			m.String(),
			m.Long(),
			m.Long(),
			m.String(),
			m.Int(),
			toTypePageBlockSlice(m.Vector()),
			toTypeRichText(m.Object()),
		}

	case crc_pageBlockSlideshow:
		r = &PredPageBlockSlideshow{
			toTypePageBlockSlice(m.Vector()),
			toTypeRichText(m.Object()),
		}

	case crc_pagePart:
		r = &PredPagePart{
			toTypePageBlockSlice(m.Vector()),
			toTypePhotoSlice(m.Vector()),
			toTypeDocumentSlice(m.Vector()),
		}

	case crc_pageFull:
		r = &PredPageFull{
			toTypePageBlockSlice(m.Vector()),
			toTypePhotoSlice(m.Vector()),
			toTypeDocumentSlice(m.Vector()),
		}

	case crc_updatePhoneCall:
		r = &PredUpdatePhoneCall{
			toTypePhoneCall(m.Object()),
		}

	case crc_updateDialogPinned:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateDialogPinned{
			flags,
			toTypePeer(m.Object()),
		}

	case crc_updatePinnedDialogs:
		flags := m.Flags()
		_ = flags
		r = &PredUpdatePinnedDialogs{
			flags,
			toTypePeerSlice(m.FlaggedVector(flags, 0)),
		}

	case crc_inputPrivacyKeyPhoneCall:
		r = &PredInputPrivacyKeyPhoneCall{}

	case crc_privacyKeyPhoneCall:
		r = &PredPrivacyKeyPhoneCall{}

	case crc_pageBlockUnsupported:
		r = &PredPageBlockUnsupported{}

	case crc_pageBlockAnchor:
		r = &PredPageBlockAnchor{
			m.String(),
		}

	case crc_pageBlockCollage:
		r = &PredPageBlockCollage{
			toTypePageBlockSlice(m.Vector()),
			toTypeRichText(m.Object()),
		}

	case crc_inputPhoneCall:
		r = &PredInputPhoneCall{
			m.Long(),
			m.Long(),
		}

	case crc_phoneCallEmpty:
		r = &PredPhoneCallEmpty{
			m.Long(),
		}

	case crc_phoneCallWaiting:
		flags := m.Flags()
		_ = flags
		r = &PredPhoneCallWaiting{
			flags,
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypePhoneCallProtocol(m.Object()),
			m.FlaggedInt(flags, 0),
		}

	case crc_phoneCallRequested:
		r = &PredPhoneCallRequested{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			toTypePhoneCallProtocol(m.Object()),
		}

	case crc_phoneCall:
		r = &PredPhoneCall{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
			toTypePhoneCallProtocol(m.Object()),
			toTypePhoneConnection(m.Object()),
			toTypePhoneConnectionSlice(m.Vector()),
			m.Int(),
		}

	case crc_phoneCallDiscarded:
		flags := m.Flags()
		_ = flags
		r = &PredPhoneCallDiscarded{
			flags,
			m.Long(),
			toTypePhoneCallDiscardReason(m.FlaggedObject(flags, 0)),
			m.FlaggedInt(flags, 1),
		}

	case crc_phoneConnection:
		r = &PredPhoneConnection{
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_phoneCallProtocol:
		flags := m.Flags()
		_ = flags
		r = &PredPhoneCallProtocol{
			flags,
			m.Int(),
			m.Int(),
		}

	case crc_phonePhoneCall:
		r = &PredPhonePhoneCall{
			toTypePhoneCall(m.Object()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_phoneCallDiscardReasonMissed:
		r = &PredPhoneCallDiscardReasonMissed{}

	case crc_phoneCallDiscardReasonDisconnect:
		r = &PredPhoneCallDiscardReasonDisconnect{}

	case crc_phoneCallDiscardReasonHangup:
		r = &PredPhoneCallDiscardReasonHangup{}

	case crc_phoneCallDiscardReasonBusy:
		r = &PredPhoneCallDiscardReasonBusy{}

	case crc_inputMessagesFilterPhoneCalls:
		flags := m.Flags()
		_ = flags
		r = &PredInputMessagesFilterPhoneCalls{
			flags,
		}

	case crc_messageActionPhoneCall:
		flags := m.Flags()
		_ = flags
		r = &PredMessageActionPhoneCall{
			flags,
			m.Long(),
			toTypePhoneCallDiscardReason(m.FlaggedObject(flags, 0)),
			m.FlaggedInt(flags, 1),
		}

	case crc_invoice:
		flags := m.Flags()
		_ = flags
		r = &PredInvoice{
			flags,
			m.String(),
			toTypeLabeledPriceSlice(m.Vector()),
		}

	case crc_inputMediaInvoice:
		flags := m.Flags()
		_ = flags
		r = &PredInputMediaInvoice{
			flags,
			m.String(),
			m.String(),
			toTypeInputWebDocument(m.FlaggedObject(flags, 0)),
			toTypeInvoice(m.Object()),
			m.StringBytes(),
			m.String(),
			m.String(),
		}

	case crc_messageActionPaymentSentMe:
		flags := m.Flags()
		_ = flags
		r = &PredMessageActionPaymentSentMe{
			flags,
			m.String(),
			m.Long(),
			m.StringBytes(),
			toTypePaymentRequestedInfo(m.FlaggedObject(flags, 0)),
			m.FlaggedString(flags, 1),
			toTypePaymentCharge(m.Object()),
		}

	case crc_messageMediaInvoice:
		flags := m.Flags()
		_ = flags
		r = &PredMessageMediaInvoice{
			flags,
			m.String(),
			m.String(),
			toTypeWebDocument(m.FlaggedObject(flags, 0)),
			m.FlaggedInt(flags, 2),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_keyboardButtonBuy:
		r = &PredKeyboardButtonBuy{
			m.String(),
		}

	case crc_messageActionPaymentSent:
		r = &PredMessageActionPaymentSent{
			m.String(),
			m.Long(),
		}

	case crc_paymentsPaymentForm:
		flags := m.Flags()
		_ = flags
		r = &PredPaymentsPaymentForm{
			flags,
			m.Int(),
			toTypeInvoice(m.Object()),
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 4),
			toTypeDataJSON(m.FlaggedObject(flags, 4)),
			toTypePaymentRequestedInfo(m.FlaggedObject(flags, 0)),
			toTypePaymentSavedCredentials(m.FlaggedObject(flags, 1)),
			toTypeUserSlice(m.Vector()),
		}

	case crc_postAddress:
		r = &PredPostAddress{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_paymentRequestedInfo:
		flags := m.Flags()
		_ = flags
		r = &PredPaymentRequestedInfo{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			toTypePostAddress(m.FlaggedObject(flags, 3)),
		}

	case crc_updateBotWebhookJSON:
		r = &PredUpdateBotWebhookJSON{
			toTypeDataJSON(m.Object()),
		}

	case crc_updateBotWebhookJSONQuery:
		r = &PredUpdateBotWebhookJSONQuery{
			m.Long(),
			toTypeDataJSON(m.Object()),
			m.Int(),
		}

	case crc_updateBotShippingQuery:
		r = &PredUpdateBotShippingQuery{
			m.Long(),
			m.Int(),
			m.StringBytes(),
			toTypePostAddress(m.Object()),
		}

	case crc_updateBotPrecheckoutQuery:
		flags := m.Flags()
		_ = flags
		r = &PredUpdateBotPrecheckoutQuery{
			flags,
			m.Long(),
			m.Int(),
			m.StringBytes(),
			toTypePaymentRequestedInfo(m.FlaggedObject(flags, 0)),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Long(),
		}

	case crc_dataJSON:
		r = &PredDataJSON{
			m.String(),
		}

	case crc_labeledPrice:
		r = &PredLabeledPrice{
			m.String(),
			m.Long(),
		}

	case crc_paymentCharge:
		r = &PredPaymentCharge{
			m.String(),
			m.String(),
		}

	case crc_paymentSavedCredentialsCard:
		r = &PredPaymentSavedCredentialsCard{
			m.String(),
			m.String(),
		}

	case crc_webDocument:
		r = &PredWebDocument{
			m.String(),
			m.Long(),
			m.Int(),
			m.String(),
			toTypeDocumentAttributeSlice(m.Vector()),
			m.Int(),
		}

	case crc_inputWebDocument:
		r = &PredInputWebDocument{
			m.String(),
			m.Int(),
			m.String(),
			toTypeDocumentAttributeSlice(m.Vector()),
		}

	case crc_inputWebFileLocation:
		r = &PredInputWebFileLocation{
			m.String(),
			m.Long(),
		}

	case crc_uploadWebFile:
		r = &PredUploadWebFile{
			m.Int(),
			m.String(),
			toTypeStorageFileType(m.Object()),
			m.Int(),
			m.StringBytes(),
		}

	case crc_paymentsValidatedRequestedInfo:
		flags := m.Flags()
		_ = flags
		r = &PredPaymentsValidatedRequestedInfo{
			flags,
			m.FlaggedString(flags, 0),
			toTypeShippingOptionSlice(m.FlaggedVector(flags, 1)),
		}

	case crc_paymentsPaymentResult:
		r = &PredPaymentsPaymentResult{
			toTypeUpdates(m.Object()),
		}

	case crc_paymentsPaymentVerficationNeeded:
		r = &PredPaymentsPaymentVerficationNeeded{
			m.String(),
		}

	case crc_paymentsPaymentReceipt:
		flags := m.Flags()
		_ = flags
		r = &PredPaymentsPaymentReceipt{
			flags,
			m.Int(),
			m.Int(),
			toTypeInvoice(m.Object()),
			m.Int(),
			toTypePaymentRequestedInfo(m.FlaggedObject(flags, 0)),
			toTypeShippingOption(m.FlaggedObject(flags, 1)),
			m.String(),
			m.Long(),
			m.String(),
			toTypeUserSlice(m.Vector()),
		}

	case crc_paymentsSavedInfo:
		flags := m.Flags()
		_ = flags
		r = &PredPaymentsSavedInfo{
			flags,
			toTypePaymentRequestedInfo(m.FlaggedObject(flags, 0)),
		}

	case crc_inputPaymentCredentialsSaved:
		r = &PredInputPaymentCredentialsSaved{
			m.String(),
			m.StringBytes(),
		}

	case crc_inputPaymentCredentials:
		flags := m.Flags()
		_ = flags
		r = &PredInputPaymentCredentials{
			flags,
			toTypeDataJSON(m.Object()),
		}

	case crc_accountTmpPassword:
		r = &PredAccountTmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case crc_shippingOption:
		r = &PredShippingOption{
			m.String(),
			m.String(),
			toTypeLabeledPriceSlice(m.Vector()),
		}

	case crc_phoneCallAccepted:
		r = &PredPhoneCallAccepted{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			toTypePhoneCallProtocol(m.Object()),
		}

	case crc_inputMessagesFilterRoundVoice:
		r = &PredInputMessagesFilterRoundVoice{}

	case crc_inputMessagesFilterRoundVideo:
		r = &PredInputMessagesFilterRoundVideo{}

	case crc_uploadFileCdnRedirect:
		r = &PredUploadFileCdnRedirect{
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
			toTypeCdnFileHashSlice(m.Vector()),
		}

	case crc_sendMessageRecordRoundAction:
		r = &PredSendMessageRecordRoundAction{}

	case crc_sendMessageUploadRoundAction:
		r = &PredSendMessageUploadRoundAction{
			m.Int(),
		}

	case crc_uploadCdnFileReuploadNeeded:
		r = &PredUploadCdnFileReuploadNeeded{
			m.StringBytes(),
		}

	case crc_uploadCdnFile:
		r = &PredUploadCdnFile{
			m.StringBytes(),
		}

	case crc_cdnPublicKey:
		r = &PredCdnPublicKey{
			m.Int(),
			m.String(),
		}

	case crc_cdnConfig:
		r = &PredCdnConfig{
			toTypeCdnPublicKeySlice(m.Vector()),
		}

	case crc_updateLangPackTooLong:
		r = &PredUpdateLangPackTooLong{}

	case crc_updateLangPack:
		r = &PredUpdateLangPack{
			toTypeLangPackDifference(m.Object()),
		}

	case crc_pageBlockChannel:
		r = &PredPageBlockChannel{
			toTypeChat(m.Object()),
		}

	case crc_inputStickerSetItem:
		flags := m.Flags()
		_ = flags
		r = &PredInputStickerSetItem{
			flags,
			toTypeInputDocument(m.Object()),
			m.String(),
			toTypeMaskCoords(m.FlaggedObject(flags, 0)),
		}

	case crc_langPackString:
		r = &PredLangPackString{
			m.String(),
			m.String(),
		}

	case crc_langPackStringPluralized:
		flags := m.Flags()
		_ = flags
		r = &PredLangPackStringPluralized{
			flags,
			m.String(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.String(),
		}

	case crc_langPackStringDeleted:
		r = &PredLangPackStringDeleted{
			m.String(),
		}

	case crc_langPackDifference:
		r = &PredLangPackDifference{
			m.String(),
			m.Int(),
			m.Int(),
			toTypeLangPackStringSlice(m.Vector()),
		}

	case crc_langPackLanguage:
		r = &PredLangPackLanguage{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_channelParticipantAdmin:
		flags := m.Flags()
		_ = flags
		r = &PredChannelParticipantAdmin{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			toTypeChannelAdminRights(m.Object()),
		}

	case crc_channelParticipantBanned:
		flags := m.Flags()
		_ = flags
		r = &PredChannelParticipantBanned{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			toTypeChannelBannedRights(m.Object()),
		}

	case crc_channelParticipantsBanned:
		r = &PredChannelParticipantsBanned{
			m.String(),
		}

	case crc_channelParticipantsSearch:
		r = &PredChannelParticipantsSearch{
			m.String(),
		}

	case crc_topPeerCategoryPhoneCalls:
		r = &PredTopPeerCategoryPhoneCalls{}

	case crc_pageBlockAudio:
		r = &PredPageBlockAudio{
			m.Long(),
			toTypeRichText(m.Object()),
		}

	case crc_channelAdminRights:
		flags := m.Flags()
		_ = flags
		r = &PredChannelAdminRights{
			flags,
		}

	case crc_channelBannedRights:
		flags := m.Flags()
		_ = flags
		r = &PredChannelBannedRights{
			flags,
			m.Int(),
		}

	case crc_channelAdminLogEventActionChangeTitle:
		r = &PredChannelAdminLogEventActionChangeTitle{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeAbout:
		r = &PredChannelAdminLogEventActionChangeAbout{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeUsername:
		r = &PredChannelAdminLogEventActionChangeUsername{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangePhoto:
		r = &PredChannelAdminLogEventActionChangePhoto{
			toTypeChatPhoto(m.Object()),
			toTypeChatPhoto(m.Object()),
		}

	case crc_channelAdminLogEventActionToggleInvites:
		r = &PredChannelAdminLogEventActionToggleInvites{
			toTypeBool(m.Object()),
		}

	case crc_channelAdminLogEventActionToggleSignatures:
		r = &PredChannelAdminLogEventActionToggleSignatures{
			toTypeBool(m.Object()),
		}

	case crc_channelAdminLogEventActionUpdatePinned:
		r = &PredChannelAdminLogEventActionUpdatePinned{
			toTypeMessage(m.Object()),
		}

	case crc_channelAdminLogEventActionEditMessage:
		r = &PredChannelAdminLogEventActionEditMessage{
			toTypeMessage(m.Object()),
			toTypeMessage(m.Object()),
		}

	case crc_channelAdminLogEventActionDeleteMessage:
		r = &PredChannelAdminLogEventActionDeleteMessage{
			toTypeMessage(m.Object()),
		}

	case crc_channelAdminLogEventActionParticipantJoin:
		r = &PredChannelAdminLogEventActionParticipantJoin{}

	case crc_channelAdminLogEventActionParticipantLeave:
		r = &PredChannelAdminLogEventActionParticipantLeave{}

	case crc_channelAdminLogEventActionParticipantInvite:
		r = &PredChannelAdminLogEventActionParticipantInvite{
			toTypeChannelParticipant(m.Object()),
		}

	case crc_channelAdminLogEventActionParticipantToggleBan:
		r = &PredChannelAdminLogEventActionParticipantToggleBan{
			toTypeChannelParticipant(m.Object()),
			toTypeChannelParticipant(m.Object()),
		}

	case crc_channelAdminLogEventActionParticipantToggleAdmin:
		r = &PredChannelAdminLogEventActionParticipantToggleAdmin{
			toTypeChannelParticipant(m.Object()),
			toTypeChannelParticipant(m.Object()),
		}

	case crc_channelAdminLogEvent:
		r = &PredChannelAdminLogEvent{
			m.Long(),
			m.Int(),
			m.Int(),
			toTypeChannelAdminLogEventAction(m.Object()),
		}

	case crc_channelsAdminLogResults:
		r = &PredChannelsAdminLogResults{
			toTypeChannelAdminLogEventSlice(m.Vector()),
			toTypeChatSlice(m.Vector()),
			toTypeUserSlice(m.Vector()),
		}

	case crc_channelAdminLogEventsFilter:
		flags := m.Flags()
		_ = flags
		r = &PredChannelAdminLogEventsFilter{
			flags,
		}

	case crc_messageActionScreenshotTaken:
		r = &PredMessageActionScreenshotTaken{}

	case crc_popularContact:
		r = &PredPopularContact{
			m.Long(),
			m.Int(),
		}

	case crc_cdnFileHash:
		r = &PredCdnFileHash{
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputMessagesFilterMyMentions:
		r = &PredInputMessagesFilterMyMentions{}

	case crc_inputMessagesFilterMyMentionsUnread:
		r = &PredInputMessagesFilterMyMentionsUnread{}

	case crc_updateContactsReset:
		r = &PredUpdateContactsReset{}

	case crc_channelAdminLogEventActionChangeStickerSet:
		r = &PredChannelAdminLogEventActionChangeStickerSet{
			toTypeInputStickerSet(m.Object()),
			toTypeInputStickerSet(m.Object()),
		}

	case crc_updateFavedStickers:
		r = &PredUpdateFavedStickers{}

	case crc_messagesFavedStickers:
		r = &PredMessagesFavedStickers{
			m.Int(),
			toTypeStickerPackSlice(m.Vector()),
			toTypeDocumentSlice(m.Vector()),
		}

	case crc_messagesFavedStickersNotModified:
		r = &PredMessagesFavedStickersNotModified{}

	case crc_updateChannelReadMessagesContents:
		r = &PredUpdateChannelReadMessagesContents{
			m.Int(),
			m.VectorInt(),
		}

	case crc_invokeAfterMsg:
		r = &ReqInvokeAfterMsg{
			m.Long(),
			Pack(m.Object()),
		}

	case crc_invokeAfterMsgs:
		r = &ReqInvokeAfterMsgs{
			m.VectorLong(),
			Pack(m.Object()),
		}

	case crc_authCheckPhone:
		r = &ReqAuthCheckPhone{
			m.String(),
		}

	case crc_authSendCode:
		flags := m.Flags()
		_ = flags
		r = &ReqAuthSendCode{
			flags,
			m.String(),
			toTypeBool(m.FlaggedObject(flags, 0)),
			m.Int(),
			m.String(),
		}

	case crc_authSignUp:
		r = &ReqAuthSignUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_authSignIn:
		r = &ReqAuthSignIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_authLogOut:
		r = &ReqAuthLogOut{}

	case crc_authResetAuthorizations:
		r = &ReqAuthResetAuthorizations{}

	case crc_authSendInvites:
		r = &ReqAuthSendInvites{
			m.VectorString(),
			m.String(),
		}

	case crc_authExportAuthorization:
		r = &ReqAuthExportAuthorization{
			m.Int(),
		}

	case crc_authImportAuthorization:
		r = &ReqAuthImportAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_accountRegisterDevice:
		r = &ReqAccountRegisterDevice{
			m.Int(),
			m.String(),
		}

	case crc_accountUnregisterDevice:
		r = &ReqAccountUnregisterDevice{
			m.Int(),
			m.String(),
		}

	case crc_accountUpdateNotifySettings:
		r = &ReqAccountUpdateNotifySettings{
			toTypeInputNotifyPeer(m.Object()),
			toTypeInputPeerNotifySettings(m.Object()),
		}

	case crc_accountGetNotifySettings:
		r = &ReqAccountGetNotifySettings{
			toTypeInputNotifyPeer(m.Object()),
		}

	case crc_accountResetNotifySettings:
		r = &ReqAccountResetNotifySettings{}

	case crc_accountUpdateProfile:
		flags := m.Flags()
		_ = flags
		r = &ReqAccountUpdateProfile{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case crc_accountUpdateStatus:
		r = &ReqAccountUpdateStatus{
			toTypeBool(m.Object()),
		}

	case crc_accountGetWallPapers:
		r = &ReqAccountGetWallPapers{}

	case crc_usersGetUsers:
		r = &ReqUsersGetUsers{
			toTypeInputUserSlice(m.Vector()),
		}

	case crc_usersGetFullUser:
		r = &ReqUsersGetFullUser{
			toTypeInputUser(m.Object()),
		}

	case crc_contactsGetStatuses:
		r = &ReqContactsGetStatuses{}

	case crc_contactsGetContacts:
		r = &ReqContactsGetContacts{
			m.Int(),
		}

	case crc_contactsImportContacts:
		r = &ReqContactsImportContacts{
			toTypeInputContactSlice(m.Vector()),
		}

	case crc_contactsSearch:
		r = &ReqContactsSearch{
			m.String(),
			m.Int(),
		}

	case crc_contactsDeleteContact:
		r = &ReqContactsDeleteContact{
			toTypeInputUser(m.Object()),
		}

	case crc_contactsDeleteContacts:
		r = &ReqContactsDeleteContacts{
			toTypeInputUserSlice(m.Vector()),
		}

	case crc_contactsBlock:
		r = &ReqContactsBlock{
			toTypeInputUser(m.Object()),
		}

	case crc_contactsUnblock:
		r = &ReqContactsUnblock{
			toTypeInputUser(m.Object()),
		}

	case crc_contactsGetBlocked:
		r = &ReqContactsGetBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_messagesGetMessages:
		r = &ReqMessagesGetMessages{
			m.VectorInt(),
		}

	case crc_messagesGetDialogs:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesGetDialogs{
			flags,
			m.Int(),
			m.Int(),
			toTypeInputPeer(m.Object()),
			m.Int(),
		}

	case crc_messagesGetHistory:
		r = &ReqMessagesGetHistory{
			toTypeInputPeer(m.Object()),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messagesSearch:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSearch{
			flags,
			toTypeInputPeer(m.Object()),
			m.String(),
			toTypeInputUser(m.FlaggedObject(flags, 0)),
			toTypeMessagesFilter(m.Object()),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messagesReadHistory:
		r = &ReqMessagesReadHistory{
			toTypeInputPeer(m.Object()),
			m.Int(),
		}

	case crc_messagesDeleteHistory:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesDeleteHistory{
			flags,
			toTypeInputPeer(m.Object()),
			m.Int(),
		}

	case crc_messagesDeleteMessages:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesDeleteMessages{
			flags,
			m.VectorInt(),
		}

	case crc_messagesReceivedMessages:
		r = &ReqMessagesReceivedMessages{
			m.Int(),
		}

	case crc_messagesSetTyping:
		r = &ReqMessagesSetTyping{
			toTypeInputPeer(m.Object()),
			toTypeSendMessageAction(m.Object()),
		}

	case crc_messagesSendMessage:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSendMessage{
			flags,
			toTypeInputPeer(m.Object()),
			m.FlaggedInt(flags, 0),
			m.String(),
			m.Long(),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 3)),
		}

	case crc_messagesSendMedia:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSendMedia{
			flags,
			toTypeInputPeer(m.Object()),
			m.FlaggedInt(flags, 0),
			toTypeInputMedia(m.Object()),
			m.Long(),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
		}

	case crc_messagesForwardMessages:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesForwardMessages{
			flags,
			toTypeInputPeer(m.Object()),
			m.VectorInt(),
			m.VectorLong(),
			toTypeInputPeer(m.Object()),
		}

	case crc_messagesGetChats:
		r = &ReqMessagesGetChats{
			m.VectorInt(),
		}

	case crc_messagesGetFullChat:
		r = &ReqMessagesGetFullChat{
			m.Int(),
		}

	case crc_messagesEditChatTitle:
		r = &ReqMessagesEditChatTitle{
			m.Int(),
			m.String(),
		}

	case crc_messagesEditChatPhoto:
		r = &ReqMessagesEditChatPhoto{
			m.Int(),
			toTypeInputChatPhoto(m.Object()),
		}

	case crc_messagesAddChatUser:
		r = &ReqMessagesAddChatUser{
			m.Int(),
			toTypeInputUser(m.Object()),
			m.Int(),
		}

	case crc_messagesDeleteChatUser:
		r = &ReqMessagesDeleteChatUser{
			m.Int(),
			toTypeInputUser(m.Object()),
		}

	case crc_messagesCreateChat:
		r = &ReqMessagesCreateChat{
			toTypeInputUserSlice(m.Vector()),
			m.String(),
		}

	case crc_updatesGetState:
		r = &ReqUpdatesGetState{}

	case crc_updatesGetDifference:
		flags := m.Flags()
		_ = flags
		r = &ReqUpdatesGetDifference{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
		}

	case crc_photosUpdateProfilePhoto:
		r = &ReqPhotosUpdateProfilePhoto{
			toTypeInputPhoto(m.Object()),
		}

	case crc_photosUploadProfilePhoto:
		r = &ReqPhotosUploadProfilePhoto{
			toTypeInputFile(m.Object()),
		}

	case crc_uploadSaveFilePart:
		r = &ReqUploadSaveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_uploadGetFile:
		r = &ReqUploadGetFile{
			toTypeInputFileLocation(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_helpGetConfig:
		r = &ReqHelpGetConfig{}

	case crc_helpGetNearestDc:
		r = &ReqHelpGetNearestDc{}

	case crc_helpGetAppUpdate:
		r = &ReqHelpGetAppUpdate{}

	case crc_helpSaveAppLog:
		r = &ReqHelpSaveAppLog{
			toTypeInputAppEventSlice(m.Vector()),
		}

	case crc_helpGetInviteText:
		r = &ReqHelpGetInviteText{}

	case crc_photosDeletePhotos:
		r = &ReqPhotosDeletePhotos{
			toTypeInputPhotoSlice(m.Vector()),
		}

	case crc_photosGetUserPhotos:
		r = &ReqPhotosGetUserPhotos{
			toTypeInputUser(m.Object()),
			m.Int(),
			m.Long(),
			m.Int(),
		}

	case crc_messagesForwardMessage:
		r = &ReqMessagesForwardMessage{
			toTypeInputPeer(m.Object()),
			m.Int(),
			m.Long(),
		}

	case crc_messagesGetDhConfig:
		r = &ReqMessagesGetDhConfig{
			m.Int(),
			m.Int(),
		}

	case crc_messagesRequestEncryption:
		r = &ReqMessagesRequestEncryption{
			toTypeInputUser(m.Object()),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messagesAcceptEncryption:
		r = &ReqMessagesAcceptEncryption{
			toTypeInputEncryptedChat(m.Object()),
			m.StringBytes(),
			m.Long(),
		}

	case crc_messagesDiscardEncryption:
		r = &ReqMessagesDiscardEncryption{
			m.Int(),
		}

	case crc_messagesSetEncryptedTyping:
		r = &ReqMessagesSetEncryptedTyping{
			toTypeInputEncryptedChat(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_messagesReadEncryptedHistory:
		r = &ReqMessagesReadEncryptedHistory{
			toTypeInputEncryptedChat(m.Object()),
			m.Int(),
		}

	case crc_messagesSendEncrypted:
		r = &ReqMessagesSendEncrypted{
			toTypeInputEncryptedChat(m.Object()),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messagesSendEncryptedFile:
		r = &ReqMessagesSendEncryptedFile{
			toTypeInputEncryptedChat(m.Object()),
			m.Long(),
			m.StringBytes(),
			toTypeInputEncryptedFile(m.Object()),
		}

	case crc_messagesSendEncryptedService:
		r = &ReqMessagesSendEncryptedService{
			toTypeInputEncryptedChat(m.Object()),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messagesReceivedQueue:
		r = &ReqMessagesReceivedQueue{
			m.Int(),
		}

	case crc_uploadSaveBigFilePart:
		r = &ReqUploadSaveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_initConnection:
		r = &ReqInitConnection{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			Pack(m.Object()),
		}

	case crc_helpGetSupport:
		r = &ReqHelpGetSupport{}

	case crc_authBindTempAuthKey:
		r = &ReqAuthBindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_contactsExportCard:
		r = &ReqContactsExportCard{}

	case crc_contactsImportCard:
		r = &ReqContactsImportCard{
			m.VectorInt(),
		}

	case crc_messagesReadMessageContents:
		r = &ReqMessagesReadMessageContents{
			m.VectorInt(),
		}

	case crc_accountCheckUsername:
		r = &ReqAccountCheckUsername{
			m.String(),
		}

	case crc_accountUpdateUsername:
		r = &ReqAccountUpdateUsername{
			m.String(),
		}

	case crc_accountGetPrivacy:
		r = &ReqAccountGetPrivacy{
			toTypeInputPrivacyKey(m.Object()),
		}

	case crc_accountSetPrivacy:
		r = &ReqAccountSetPrivacy{
			toTypeInputPrivacyKey(m.Object()),
			toTypeInputPrivacyRuleSlice(m.Vector()),
		}

	case crc_accountDeleteAccount:
		r = &ReqAccountDeleteAccount{
			m.String(),
		}

	case crc_accountGetAccountTTL:
		r = &ReqAccountGetAccountTTL{}

	case crc_accountSetAccountTTL:
		r = &ReqAccountSetAccountTTL{
			toTypeAccountDaysTTL(m.Object()),
		}

	case crc_invokeWithLayer:
		r = &ReqInvokeWithLayer{
			m.Int(),
			Pack(m.Object()),
		}

	case crc_contactsResolveUsername:
		r = &ReqContactsResolveUsername{
			m.String(),
		}

	case crc_accountSendChangePhoneCode:
		flags := m.Flags()
		_ = flags
		r = &ReqAccountSendChangePhoneCode{
			flags,
			m.String(),
			toTypeBool(m.FlaggedObject(flags, 0)),
		}

	case crc_accountChangePhone:
		r = &ReqAccountChangePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messagesGetAllStickers:
		r = &ReqMessagesGetAllStickers{
			m.Int(),
		}

	case crc_accountUpdateDeviceLocked:
		r = &ReqAccountUpdateDeviceLocked{
			m.Int(),
		}

	case crc_accountGetPassword:
		r = &ReqAccountGetPassword{}

	case crc_authCheckPassword:
		r = &ReqAuthCheckPassword{
			m.StringBytes(),
		}

	case crc_messagesGetWebPagePreview:
		r = &ReqMessagesGetWebPagePreview{
			m.String(),
		}

	case crc_accountGetAuthorizations:
		r = &ReqAccountGetAuthorizations{}

	case crc_accountResetAuthorization:
		r = &ReqAccountResetAuthorization{
			m.Long(),
		}

	case crc_accountGetPasswordSettings:
		r = &ReqAccountGetPasswordSettings{
			m.StringBytes(),
		}

	case crc_accountUpdatePasswordSettings:
		r = &ReqAccountUpdatePasswordSettings{
			m.StringBytes(),
			toTypeAccountPasswordInputSettings(m.Object()),
		}

	case crc_authRequestPasswordRecovery:
		r = &ReqAuthRequestPasswordRecovery{}

	case crc_authRecoverPassword:
		r = &ReqAuthRecoverPassword{
			m.String(),
		}

	case crc_invokeWithoutUpdates:
		r = &ReqInvokeWithoutUpdates{
			Pack(m.Object()),
		}

	case crc_messagesExportChatInvite:
		r = &ReqMessagesExportChatInvite{
			m.Int(),
		}

	case crc_messagesCheckChatInvite:
		r = &ReqMessagesCheckChatInvite{
			m.String(),
		}

	case crc_messagesImportChatInvite:
		r = &ReqMessagesImportChatInvite{
			m.String(),
		}

	case crc_messagesGetStickerSet:
		r = &ReqMessagesGetStickerSet{
			toTypeInputStickerSet(m.Object()),
		}

	case crc_messagesInstallStickerSet:
		r = &ReqMessagesInstallStickerSet{
			toTypeInputStickerSet(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_messagesUninstallStickerSet:
		r = &ReqMessagesUninstallStickerSet{
			toTypeInputStickerSet(m.Object()),
		}

	case crc_authImportBotAuthorization:
		r = &ReqAuthImportBotAuthorization{
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_messagesStartBot:
		r = &ReqMessagesStartBot{
			toTypeInputUser(m.Object()),
			toTypeInputPeer(m.Object()),
			m.Long(),
			m.String(),
		}

	case crc_helpGetAppChangelog:
		r = &ReqHelpGetAppChangelog{
			m.String(),
		}

	case crc_messagesReportSpam:
		r = &ReqMessagesReportSpam{
			toTypeInputPeer(m.Object()),
		}

	case crc_messagesGetMessagesViews:
		r = &ReqMessagesGetMessagesViews{
			toTypeInputPeer(m.Object()),
			m.VectorInt(),
			toTypeBool(m.Object()),
		}

	case crc_updatesGetChannelDifference:
		flags := m.Flags()
		_ = flags
		r = &ReqUpdatesGetChannelDifference{
			flags,
			toTypeInputChannel(m.Object()),
			toTypeChannelMessagesFilter(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_channelsReadHistory:
		r = &ReqChannelsReadHistory{
			toTypeInputChannel(m.Object()),
			m.Int(),
		}

	case crc_channelsDeleteMessages:
		r = &ReqChannelsDeleteMessages{
			toTypeInputChannel(m.Object()),
			m.VectorInt(),
		}

	case crc_channelsDeleteUserHistory:
		r = &ReqChannelsDeleteUserHistory{
			toTypeInputChannel(m.Object()),
			toTypeInputUser(m.Object()),
		}

	case crc_channelsReportSpam:
		r = &ReqChannelsReportSpam{
			toTypeInputChannel(m.Object()),
			toTypeInputUser(m.Object()),
			m.VectorInt(),
		}

	case crc_channelsGetMessages:
		r = &ReqChannelsGetMessages{
			toTypeInputChannel(m.Object()),
			m.VectorInt(),
		}

	case crc_channelsGetParticipants:
		r = &ReqChannelsGetParticipants{
			toTypeInputChannel(m.Object()),
			toTypeChannelParticipantsFilter(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_channelsGetParticipant:
		r = &ReqChannelsGetParticipant{
			toTypeInputChannel(m.Object()),
			toTypeInputUser(m.Object()),
		}

	case crc_channelsGetChannels:
		r = &ReqChannelsGetChannels{
			toTypeInputChannelSlice(m.Vector()),
		}

	case crc_channelsGetFullChannel:
		r = &ReqChannelsGetFullChannel{
			toTypeInputChannel(m.Object()),
		}

	case crc_channelsCreateChannel:
		flags := m.Flags()
		_ = flags
		r = &ReqChannelsCreateChannel{
			flags,
			m.String(),
			m.String(),
		}

	case crc_channelsEditAbout:
		r = &ReqChannelsEditAbout{
			toTypeInputChannel(m.Object()),
			m.String(),
		}

	case crc_channelsEditAdmin:
		r = &ReqChannelsEditAdmin{
			toTypeInputChannel(m.Object()),
			toTypeInputUser(m.Object()),
			toTypeChannelAdminRights(m.Object()),
		}

	case crc_channelsEditTitle:
		r = &ReqChannelsEditTitle{
			toTypeInputChannel(m.Object()),
			m.String(),
		}

	case crc_channelsEditPhoto:
		r = &ReqChannelsEditPhoto{
			toTypeInputChannel(m.Object()),
			toTypeInputChatPhoto(m.Object()),
		}

	case crc_channelsCheckUsername:
		r = &ReqChannelsCheckUsername{
			toTypeInputChannel(m.Object()),
			m.String(),
		}

	case crc_channelsUpdateUsername:
		r = &ReqChannelsUpdateUsername{
			toTypeInputChannel(m.Object()),
			m.String(),
		}

	case crc_channelsJoinChannel:
		r = &ReqChannelsJoinChannel{
			toTypeInputChannel(m.Object()),
		}

	case crc_channelsLeaveChannel:
		r = &ReqChannelsLeaveChannel{
			toTypeInputChannel(m.Object()),
		}

	case crc_channelsInviteToChannel:
		r = &ReqChannelsInviteToChannel{
			toTypeInputChannel(m.Object()),
			toTypeInputUserSlice(m.Vector()),
		}

	case crc_channelsExportInvite:
		r = &ReqChannelsExportInvite{
			toTypeInputChannel(m.Object()),
		}

	case crc_channelsDeleteChannel:
		r = &ReqChannelsDeleteChannel{
			toTypeInputChannel(m.Object()),
		}

	case crc_messagesToggleChatAdmins:
		r = &ReqMessagesToggleChatAdmins{
			m.Int(),
			toTypeBool(m.Object()),
		}

	case crc_messagesEditChatAdmin:
		r = &ReqMessagesEditChatAdmin{
			m.Int(),
			toTypeInputUser(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_messagesMigrateChat:
		r = &ReqMessagesMigrateChat{
			m.Int(),
		}

	case crc_messagesSearchGlobal:
		r = &ReqMessagesSearchGlobal{
			m.String(),
			m.Int(),
			toTypeInputPeer(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_accountReportPeer:
		r = &ReqAccountReportPeer{
			toTypeInputPeer(m.Object()),
			toTypeReportReason(m.Object()),
		}

	case crc_messagesReorderStickerSets:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesReorderStickerSets{
			flags,
			m.VectorLong(),
		}

	case crc_helpGetTermsOfService:
		r = &ReqHelpGetTermsOfService{}

	case crc_messagesGetDocumentByHash:
		r = &ReqMessagesGetDocumentByHash{
			m.StringBytes(),
			m.Int(),
			m.String(),
		}

	case crc_messagesSearchGifs:
		r = &ReqMessagesSearchGifs{
			m.String(),
			m.Int(),
		}

	case crc_messagesGetSavedGifs:
		r = &ReqMessagesGetSavedGifs{
			m.Int(),
		}

	case crc_messagesSaveGif:
		r = &ReqMessagesSaveGif{
			toTypeInputDocument(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_messagesGetInlineBotResults:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesGetInlineBotResults{
			flags,
			toTypeInputUser(m.Object()),
			toTypeInputPeer(m.Object()),
			toTypeInputGeoPoint(m.FlaggedObject(flags, 0)),
			m.String(),
			m.String(),
		}

	case crc_messagesSetInlineBotResults:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSetInlineBotResults{
			flags,
			m.Long(),
			toTypeInputBotInlineResultSlice(m.Vector()),
			m.Int(),
			m.FlaggedString(flags, 2),
			toTypeInlineBotSwitchPM(m.FlaggedObject(flags, 3)),
		}

	case crc_messagesSendInlineBotResult:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSendInlineBotResult{
			flags,
			toTypeInputPeer(m.Object()),
			m.FlaggedInt(flags, 0),
			m.Long(),
			m.Long(),
			m.String(),
		}

	case crc_channelsToggleInvites:
		r = &ReqChannelsToggleInvites{
			toTypeInputChannel(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_channelsExportMessageLink:
		r = &ReqChannelsExportMessageLink{
			toTypeInputChannel(m.Object()),
			m.Int(),
		}

	case crc_channelsToggleSignatures:
		r = &ReqChannelsToggleSignatures{
			toTypeInputChannel(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_messagesHideReportSpam:
		r = &ReqMessagesHideReportSpam{
			toTypeInputPeer(m.Object()),
		}

	case crc_messagesGetPeerSettings:
		r = &ReqMessagesGetPeerSettings{
			toTypeInputPeer(m.Object()),
		}

	case crc_channelsUpdatePinnedMessage:
		flags := m.Flags()
		_ = flags
		r = &ReqChannelsUpdatePinnedMessage{
			flags,
			toTypeInputChannel(m.Object()),
			m.Int(),
		}

	case crc_authResendCode:
		r = &ReqAuthResendCode{
			m.String(),
			m.String(),
		}

	case crc_authCancelCode:
		r = &ReqAuthCancelCode{
			m.String(),
			m.String(),
		}

	case crc_messagesGetMessageEditData:
		r = &ReqMessagesGetMessageEditData{
			toTypeInputPeer(m.Object()),
			m.Int(),
		}

	case crc_messagesEditMessage:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesEditMessage{
			flags,
			toTypeInputPeer(m.Object()),
			m.Int(),
			m.FlaggedString(flags, 11),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 3)),
		}

	case crc_messagesEditInlineBotMessage:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesEditInlineBotMessage{
			flags,
			toTypeInputBotInlineMessageID(m.Object()),
			m.FlaggedString(flags, 11),
			toTypeReplyMarkup(m.FlaggedObject(flags, 2)),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 3)),
		}

	case crc_messagesGetBotCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesGetBotCallbackAnswer{
			flags,
			toTypeInputPeer(m.Object()),
			m.Int(),
			m.FlaggedStringBytes(flags, 0),
		}

	case crc_messagesSetBotCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSetBotCallbackAnswer{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case crc_contactsGetTopPeers:
		flags := m.Flags()
		_ = flags
		r = &ReqContactsGetTopPeers{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_contactsResetTopPeerRating:
		r = &ReqContactsResetTopPeerRating{
			toTypeTopPeerCategory(m.Object()),
			toTypeInputPeer(m.Object()),
		}

	case crc_messagesGetPeerDialogs:
		r = &ReqMessagesGetPeerDialogs{
			toTypeInputPeerSlice(m.Vector()),
		}

	case crc_messagesSaveDraft:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSaveDraft{
			flags,
			m.FlaggedInt(flags, 0),
			toTypeInputPeer(m.Object()),
			m.String(),
			toTypeMessageEntitySlice(m.FlaggedVector(flags, 3)),
		}

	case crc_messagesGetAllDrafts:
		r = &ReqMessagesGetAllDrafts{}

	case crc_accountSendConfirmPhoneCode:
		flags := m.Flags()
		_ = flags
		r = &ReqAccountSendConfirmPhoneCode{
			flags,
			m.String(),
			toTypeBool(m.FlaggedObject(flags, 0)),
		}

	case crc_accountConfirmPhone:
		r = &ReqAccountConfirmPhone{
			m.String(),
			m.String(),
		}

	case crc_messagesGetFeaturedStickers:
		r = &ReqMessagesGetFeaturedStickers{
			m.Int(),
		}

	case crc_messagesReadFeaturedStickers:
		r = &ReqMessagesReadFeaturedStickers{
			m.VectorLong(),
		}

	case crc_messagesGetRecentStickers:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesGetRecentStickers{
			flags,
			m.Int(),
		}

	case crc_messagesSaveRecentSticker:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSaveRecentSticker{
			flags,
			toTypeInputDocument(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_messagesClearRecentStickers:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesClearRecentStickers{
			flags,
		}

	case crc_messagesGetArchivedStickers:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesGetArchivedStickers{
			flags,
			m.Long(),
			m.Int(),
		}

	case crc_channelsGetAdminedPublicChannels:
		r = &ReqChannelsGetAdminedPublicChannels{}

	case crc_authDropTempAuthKeys:
		r = &ReqAuthDropTempAuthKeys{
			m.VectorLong(),
		}

	case crc_messagesSetGameScore:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSetGameScore{
			flags,
			toTypeInputPeer(m.Object()),
			m.Int(),
			toTypeInputUser(m.Object()),
			m.Int(),
		}

	case crc_messagesSetInlineGameScore:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSetInlineGameScore{
			flags,
			toTypeInputBotInlineMessageID(m.Object()),
			toTypeInputUser(m.Object()),
			m.Int(),
		}

	case crc_messagesGetMaskStickers:
		r = &ReqMessagesGetMaskStickers{
			m.Int(),
		}

	case crc_messagesGetAttachedStickers:
		r = &ReqMessagesGetAttachedStickers{
			toTypeInputStickeredMedia(m.Object()),
		}

	case crc_messagesGetGameHighScores:
		r = &ReqMessagesGetGameHighScores{
			toTypeInputPeer(m.Object()),
			m.Int(),
			toTypeInputUser(m.Object()),
		}

	case crc_messagesGetInlineGameHighScores:
		r = &ReqMessagesGetInlineGameHighScores{
			toTypeInputBotInlineMessageID(m.Object()),
			toTypeInputUser(m.Object()),
		}

	case crc_messagesGetCommonChats:
		r = &ReqMessagesGetCommonChats{
			toTypeInputUser(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_messagesGetAllChats:
		r = &ReqMessagesGetAllChats{
			m.VectorInt(),
		}

	case crc_helpSetBotUpdatesStatus:
		r = &ReqHelpSetBotUpdatesStatus{
			m.Int(),
			m.String(),
		}

	case crc_messagesGetWebPage:
		r = &ReqMessagesGetWebPage{
			m.String(),
			m.Int(),
		}

	case crc_messagesToggleDialogPin:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesToggleDialogPin{
			flags,
			toTypeInputPeer(m.Object()),
		}

	case crc_messagesReorderPinnedDialogs:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesReorderPinnedDialogs{
			flags,
			toTypeInputPeerSlice(m.Vector()),
		}

	case crc_messagesGetPinnedDialogs:
		r = &ReqMessagesGetPinnedDialogs{}

	case crc_phoneRequestCall:
		r = &ReqPhoneRequestCall{
			toTypeInputUser(m.Object()),
			m.Int(),
			m.StringBytes(),
			toTypePhoneCallProtocol(m.Object()),
		}

	case crc_phoneAcceptCall:
		r = &ReqPhoneAcceptCall{
			toTypeInputPhoneCall(m.Object()),
			m.StringBytes(),
			toTypePhoneCallProtocol(m.Object()),
		}

	case crc_phoneDiscardCall:
		r = &ReqPhoneDiscardCall{
			toTypeInputPhoneCall(m.Object()),
			m.Int(),
			toTypePhoneCallDiscardReason(m.Object()),
			m.Long(),
		}

	case crc_phoneReceivedCall:
		r = &ReqPhoneReceivedCall{
			toTypeInputPhoneCall(m.Object()),
		}

	case crc_messagesReportEncryptedSpam:
		r = &ReqMessagesReportEncryptedSpam{
			toTypeInputEncryptedChat(m.Object()),
		}

	case crc_paymentsGetPaymentForm:
		r = &ReqPaymentsGetPaymentForm{
			m.Int(),
		}

	case crc_paymentsSendPaymentForm:
		flags := m.Flags()
		_ = flags
		r = &ReqPaymentsSendPaymentForm{
			flags,
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			toTypeInputPaymentCredentials(m.Object()),
		}

	case crc_accountGetTmpPassword:
		r = &ReqAccountGetTmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case crc_messagesSetBotShippingResults:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSetBotShippingResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
			toTypeShippingOptionSlice(m.FlaggedVector(flags, 1)),
		}

	case crc_messagesSetBotPrecheckoutResults:
		flags := m.Flags()
		_ = flags
		r = &ReqMessagesSetBotPrecheckoutResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
		}

	case crc_uploadGetWebFile:
		r = &ReqUploadGetWebFile{
			toTypeInputWebFileLocation(m.Object()),
			m.Int(),
			m.Int(),
		}

	case crc_botsSendCustomRequest:
		r = &ReqBotsSendCustomRequest{
			m.String(),
			toTypeDataJSON(m.Object()),
		}

	case crc_botsAnswerWebhookJSONQuery:
		r = &ReqBotsAnswerWebhookJSONQuery{
			m.Long(),
			toTypeDataJSON(m.Object()),
		}

	case crc_paymentsGetPaymentReceipt:
		r = &ReqPaymentsGetPaymentReceipt{
			m.Int(),
		}

	case crc_paymentsValidateRequestedInfo:
		flags := m.Flags()
		_ = flags
		r = &ReqPaymentsValidateRequestedInfo{
			flags,
			m.Int(),
			toTypePaymentRequestedInfo(m.Object()),
		}

	case crc_paymentsGetSavedInfo:
		r = &ReqPaymentsGetSavedInfo{}

	case crc_paymentsClearSavedInfo:
		flags := m.Flags()
		_ = flags
		r = &ReqPaymentsClearSavedInfo{
			flags,
		}

	case crc_phoneGetCallConfig:
		r = &ReqPhoneGetCallConfig{}

	case crc_phoneConfirmCall:
		r = &ReqPhoneConfirmCall{
			toTypeInputPhoneCall(m.Object()),
			m.StringBytes(),
			m.Long(),
			toTypePhoneCallProtocol(m.Object()),
		}

	case crc_phoneSetCallRating:
		r = &ReqPhoneSetCallRating{
			toTypeInputPhoneCall(m.Object()),
			m.Int(),
			m.String(),
		}

	case crc_phoneSaveCallDebug:
		r = &ReqPhoneSaveCallDebug{
			toTypeInputPhoneCall(m.Object()),
			toTypeDataJSON(m.Object()),
		}

	case crc_uploadGetCdnFile:
		r = &ReqUploadGetCdnFile{
			m.StringBytes(),
			m.Int(),
			m.Int(),
		}

	case crc_uploadReuploadCdnFile:
		r = &ReqUploadReuploadCdnFile{
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_helpGetCdnConfig:
		r = &ReqHelpGetCdnConfig{}

	case crc_messagesUploadMedia:
		r = &ReqMessagesUploadMedia{
			toTypeInputPeer(m.Object()),
			toTypeInputMedia(m.Object()),
		}

	case crc_stickersCreateStickerSet:
		flags := m.Flags()
		_ = flags
		r = &ReqStickersCreateStickerSet{
			flags,
			toTypeInputUser(m.Object()),
			m.String(),
			m.String(),
			toTypeInputStickerSetItemSlice(m.Vector()),
		}

	case crc_langpackGetLangPack:
		r = &ReqLangpackGetLangPack{
			m.String(),
		}

	case crc_langpackGetStrings:
		r = &ReqLangpackGetStrings{
			m.String(),
			m.VectorString(),
		}

	case crc_langpackGetDifference:
		r = &ReqLangpackGetDifference{
			m.Int(),
		}

	case crc_langpackGetLanguages:
		r = &ReqLangpackGetLanguages{}

	case crc_channelsEditBanned:
		r = &ReqChannelsEditBanned{
			toTypeInputChannel(m.Object()),
			toTypeInputUser(m.Object()),
			toTypeChannelBannedRights(m.Object()),
		}

	case crc_channelsGetAdminLog:
		flags := m.Flags()
		_ = flags
		r = &ReqChannelsGetAdminLog{
			flags,
			toTypeInputChannel(m.Object()),
			m.String(),
			toTypeChannelAdminLogEventsFilter(m.FlaggedObject(flags, 0)),
			toTypeInputUserSlice(m.FlaggedVector(flags, 1)),
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crc_stickersRemoveStickerFromSet:
		r = &ReqStickersRemoveStickerFromSet{
			toTypeInputDocument(m.Object()),
		}

	case crc_stickersChangeStickerPosition:
		r = &ReqStickersChangeStickerPosition{
			toTypeInputDocument(m.Object()),
			m.Int(),
		}

	case crc_stickersAddStickerToSet:
		r = &ReqStickersAddStickerToSet{
			toTypeInputStickerSet(m.Object()),
			toTypeInputStickerSetItem(m.Object()),
		}

	case crc_messagesSendScreenshotNotification:
		r = &ReqMessagesSendScreenshotNotification{
			toTypeInputPeer(m.Object()),
			m.Int(),
			m.Long(),
		}

	case crc_uploadGetCdnFileHashes:
		r = &ReqUploadGetCdnFileHashes{
			m.StringBytes(),
			m.Int(),
		}

	case crc_messagesGetUnreadMentions:
		r = &ReqMessagesGetUnreadMentions{
			toTypeInputPeer(m.Object()),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messagesFaveSticker:
		r = &ReqMessagesFaveSticker{
			toTypeInputDocument(m.Object()),
			toTypeBool(m.Object()),
		}

	case crc_channelsSetStickers:
		r = &ReqChannelsSetStickers{
			toTypeInputChannel(m.Object()),
			toTypeInputStickerSet(m.Object()),
		}

	case crc_contactsResetSaved:
		r = &ReqContactsResetSaved{}

	case crc_messagesGetFavedStickers:
		r = &ReqMessagesGetFavedStickers{
			m.Int(),
		}

	case crc_channelsReadMessageContents:
		r = &ReqChannelsReadMessageContents{
			toTypeInputChannel(m.Object()),
			m.VectorInt(),
		}

	default:
		m.err = fmt.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}

// Packer from TypeXXX to gRPC Any
func Pack(tl TL) *any.Any {
	var marshaled *any.Any
	var err error
	if tl == nil {
		return nil
	}
	switch x := tl.(type) {
	// types
	case *TypeBool:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeError:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeNull:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputMedia:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputChatPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputGeoPoint:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputFileLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputAppEvent:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeStorageFileType:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeFileLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUserProfilePhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUserStatus:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChatFull:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChatParticipant:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChatParticipants:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChatPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessageMedia:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessageAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeDialog:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhotoSize:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeGeoPoint:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAuthCheckedPhone:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAuthSentCode:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAuthAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAuthExportedAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputNotifyPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputPeerNotifySettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePeerNotifyEvents:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePeerNotifySettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeWallPaper:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUserFull:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeImportedContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactBlocked:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactStatus:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactsLink:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactsContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactsImportedContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactsBlocked:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactsFound:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesChats:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesChatFull:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesAffectedHistory:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesFilter:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUpdate:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUpdatesState:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUpdatesDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUpdates:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhotosPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUploadFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeDcOption:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeNearestDc:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeHelpAppUpdate:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeHelpInviteText:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputPeerNotifyEvents:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhotosPhotos:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeEncryptedChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputEncryptedChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeEncryptedFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputEncryptedFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeEncryptedMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesDhConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesSentEncryptedMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeHelpSupport:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeNotifyPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeSendMessageAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputPrivacyKey:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePrivacyKey:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputPrivacyRule:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePrivacyRule:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAccountPrivacyRules:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAccountDaysTTL:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeDisabledFeature:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeDocumentAttribute:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeStickerPack:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesAllStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAccountPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesAffectedMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactLink:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeWebPage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAccountAuthorizations:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAccountPasswordSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAccountPasswordInputSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAuthPasswordRecovery:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeReceivedNotifyMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeExportedChatInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChatInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeBotCommand:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeBotInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeKeyboardButton:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeKeyboardButtonRow:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeReplyMarkup:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessageEntity:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactsResolvedPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessageRange:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUpdatesChannelDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelMessagesFilter:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelParticipant:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelParticipantsFilter:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelsChannelParticipants:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelsChannelParticipant:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeTrue:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeReportReason:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeHelpTermsOfService:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeFoundGif:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesFoundGifs:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesSavedGifs:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputBotInlineMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputBotInlineResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeBotInlineMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeBotInlineResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesBotResults:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeExportedMessageLink:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessageFwdHeader:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePeerSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAuthCodeType:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAuthSentCodeType:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesBotCallbackAnswer:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesMessageEditData:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputBotInlineMessageID:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInlineBotSwitchPM:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesPeerDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeTopPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeTopPeerCategory:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeTopPeerCategoryPeers:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeContactsTopPeers:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeDraftMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesFeaturedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesRecentStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesArchivedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesStickerSetInstallResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeStickerSetCovered:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMaskCoords:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputStickeredMedia:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeGame:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputGame:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeHighScore:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesHighScores:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeRichText:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePageBlock:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputPhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhoneConnection:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhoneCallProtocol:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhonePhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePhoneCallDiscardReason:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInvoice:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePaymentsPaymentForm:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePostAddress:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePaymentRequestedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeDataJSON:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeLabeledPrice:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePaymentCharge:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePaymentSavedCredentials:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeWebDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputWebDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputWebFileLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUploadWebFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePaymentsValidatedRequestedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePaymentsPaymentResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePaymentsPaymentReceipt:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePaymentsSavedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputPaymentCredentials:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeAccountTmpPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeShippingOption:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeUploadCdnFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeCdnPublicKey:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeCdnConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeInputStickerSetItem:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeLangPackString:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeLangPackDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeLangPackLanguage:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelAdminRights:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelBannedRights:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelAdminLogEventAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelAdminLogEvent:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelsAdminLogResults:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeChannelAdminLogEventsFilter:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypePopularContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeCdnFileHash:
		marshaled, err = ptypes.MarshalAny(x)
	case *TypeMessagesFavedStickers:
		marshaled, err = ptypes.MarshalAny(x)

	// methods
	case *ReqInvokeAfterMsg:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqInvokeAfterMsgs:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthCheckPhone:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthSendCode:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthSignUp:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthSignIn:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthLogOut:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthResetAuthorizations:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthSendInvites:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthExportAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthImportAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountRegisterDevice:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountUnregisterDevice:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountUpdateNotifySettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountGetNotifySettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountResetNotifySettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountUpdateProfile:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountUpdateStatus:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountGetWallPapers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUsersGetUsers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUsersGetFullUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsGetStatuses:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsGetContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsImportContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsSearch:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsDeleteContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsDeleteContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsBlock:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsUnblock:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsGetBlocked:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetHistory:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSearch:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReadHistory:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesDeleteHistory:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesDeleteMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReceivedMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSetTyping:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSendMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSendMedia:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesForwardMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetChats:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetFullChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesEditChatTitle:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesEditChatPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesAddChatUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesDeleteChatUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesCreateChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUpdatesGetState:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUpdatesGetDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhotosUpdateProfilePhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhotosUploadProfilePhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUploadSaveFilePart:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUploadGetFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpGetConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpGetNearestDc:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpGetAppUpdate:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpSaveAppLog:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpGetInviteText:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhotosDeletePhotos:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhotosGetUserPhotos:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesForwardMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetDhConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesRequestEncryption:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesAcceptEncryption:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesDiscardEncryption:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSetEncryptedTyping:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReadEncryptedHistory:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSendEncrypted:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSendEncryptedFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSendEncryptedService:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReceivedQueue:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUploadSaveBigFilePart:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqInitConnection:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpGetSupport:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthBindTempAuthKey:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsExportCard:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsImportCard:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReadMessageContents:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountCheckUsername:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountUpdateUsername:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountGetPrivacy:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountSetPrivacy:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountDeleteAccount:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountGetAccountTTL:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountSetAccountTTL:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqInvokeWithLayer:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsResolveUsername:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountSendChangePhoneCode:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountChangePhone:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetAllStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountUpdateDeviceLocked:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountGetPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthCheckPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetWebPagePreview:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountGetAuthorizations:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountResetAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountGetPasswordSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountUpdatePasswordSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthRequestPasswordRecovery:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthRecoverPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqInvokeWithoutUpdates:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesExportChatInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesCheckChatInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesImportChatInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesInstallStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesUninstallStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthImportBotAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesStartBot:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpGetAppChangelog:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReportSpam:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetMessagesViews:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUpdatesGetChannelDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsReadHistory:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsDeleteMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsDeleteUserHistory:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsReportSpam:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsGetMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsGetParticipants:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsGetParticipant:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsGetChannels:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsGetFullChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsCreateChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsEditAbout:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsEditAdmin:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsEditTitle:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsEditPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsCheckUsername:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsUpdateUsername:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsJoinChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsLeaveChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsInviteToChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsExportInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsDeleteChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesToggleChatAdmins:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesEditChatAdmin:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesMigrateChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSearchGlobal:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountReportPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReorderStickerSets:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpGetTermsOfService:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetDocumentByHash:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSearchGifs:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetSavedGifs:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSaveGif:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetInlineBotResults:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSetInlineBotResults:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSendInlineBotResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsToggleInvites:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsExportMessageLink:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsToggleSignatures:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesHideReportSpam:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetPeerSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsUpdatePinnedMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthResendCode:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthCancelCode:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetMessageEditData:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesEditMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesEditInlineBotMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetBotCallbackAnswer:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSetBotCallbackAnswer:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsGetTopPeers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsResetTopPeerRating:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetPeerDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSaveDraft:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetAllDrafts:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountSendConfirmPhoneCode:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountConfirmPhone:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetFeaturedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReadFeaturedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetRecentStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSaveRecentSticker:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesClearRecentStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetArchivedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsGetAdminedPublicChannels:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAuthDropTempAuthKeys:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSetGameScore:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSetInlineGameScore:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetMaskStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetAttachedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetGameHighScores:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetInlineGameHighScores:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetCommonChats:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetAllChats:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpSetBotUpdatesStatus:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetWebPage:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesToggleDialogPin:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReorderPinnedDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetPinnedDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhoneRequestCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhoneAcceptCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhoneDiscardCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhoneReceivedCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesReportEncryptedSpam:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPaymentsGetPaymentForm:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPaymentsSendPaymentForm:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqAccountGetTmpPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSetBotShippingResults:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSetBotPrecheckoutResults:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUploadGetWebFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqBotsSendCustomRequest:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqBotsAnswerWebhookJSONQuery:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPaymentsGetPaymentReceipt:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPaymentsValidateRequestedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPaymentsGetSavedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPaymentsClearSavedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhoneGetCallConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhoneConfirmCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhoneSetCallRating:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqPhoneSaveCallDebug:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUploadGetCdnFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUploadReuploadCdnFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqHelpGetCdnConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesUploadMedia:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqStickersCreateStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqLangpackGetLangPack:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqLangpackGetStrings:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqLangpackGetDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqLangpackGetLanguages:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsEditBanned:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsGetAdminLog:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqStickersRemoveStickerFromSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqStickersChangeStickerPosition:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqStickersAddStickerToSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesSendScreenshotNotification:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqUploadGetCdnFileHashes:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetUnreadMentions:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesFaveSticker:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsSetStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqContactsResetSaved:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqMessagesGetFavedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *ReqChannelsReadMessageContents:
		marshaled, err = ptypes.MarshalAny(x)

	// predicates
	case *PredBoolFalse:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBoolTrue:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredError:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredNull:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPeerEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPeerSelf:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPeerChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputUserEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputUserSelf:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPhoneContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaUploadedPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaGeoPoint:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputChatPhotoEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputChatUploadedPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputChatPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputGeoPointEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputGeoPoint:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPhotoEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputFileLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputAppEvent:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPeerUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPeerChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFileUnknown:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFileJpeg:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFileGif:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFilePng:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFileMp3:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFileMov:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFilePartial:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFileMp4:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFileWebp:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredFileLocationUnavailable:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredFileLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserProfilePhotoEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserProfilePhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserStatusEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserStatusOnline:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserStatusOffline:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatForbidden:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatFull:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatParticipant:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatParticipantsForbidden:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatParticipants:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatPhotoEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageService:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaGeo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaUnsupported:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChatCreate:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChatEditTitle:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChatEditPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChatDeletePhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChatAddUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChatDeleteUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDialog:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhotoEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhotoSizeEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhotoSize:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhotoCachedSize:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredGeoPointEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredGeoPoint:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthCheckedPhone:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthSentCode:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthExportedAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputNotifyPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputNotifyUsers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputNotifyChats:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputNotifyAll:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPeerNotifySettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPeerNotifyEventsEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPeerNotifyEventsAll:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPeerNotifySettingsEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPeerNotifySettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredWallPaper:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserFull:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredImportedContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactBlocked:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactStatus:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsLink:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsContactsNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsImportedContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsBlocked:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsBlockedSlice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsFound:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesDialogsSlice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesMessagesSlice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesChats:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesChatFull:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesAffectedHistory:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterPhotos:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterVideo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterPhotoVideo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateNewMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateMessageID:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateDeleteMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateUserTyping:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChatUserTyping:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChatParticipants:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateUserStatus:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateUserName:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateUserPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateContactRegistered:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateContactLink:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesState:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesDifferenceEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesDifferenceSlice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesTooLong:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateShortMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateShortChatMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateShort:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesCombined:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdates:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhotosPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUploadFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDcOption:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredNearestDc:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredHelpAppUpdate:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredHelpNoAppUpdate:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredHelpInviteText:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPeerNotifyEventsEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPeerNotifyEventsAll:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhotosPhotos:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhotosPhotosSlice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredWallPaperSolid:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateNewEncryptedMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateEncryptedChatTyping:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateEncryption:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateEncryptedMessagesRead:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedChatEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedChatWaiting:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedChatRequested:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedChatDiscarded:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputEncryptedChat:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedFileEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputEncryptedFileEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputEncryptedFileUploaded:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputEncryptedFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputEncryptedFileLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredEncryptedMessageService:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesDhConfigNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesDhConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesSentEncryptedMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesSentEncryptedFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputFileBig:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputEncryptedFileBigUploaded:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStorageFilePdf:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterPhotoVideoDocuments:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChatParticipantAdd:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChatParticipantDelete:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateDcOptions:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaUploadedDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputDocumentEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputDocumentFileLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocumentEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredHelpSupport:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredNotifyAll:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredNotifyChats:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredNotifyPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredNotifyUsers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateUserBlocked:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateNotifySettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageTypingAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageCancelAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageRecordVideoAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageUploadVideoAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageRecordAudioAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageUploadAudioAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageUploadPhotoAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageUploadDocumentAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageGeoLocationAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageChooseContactAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateServiceNotification:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserStatusRecently:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserStatusLastWeek:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUserStatusLastMonth:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatePrivacy:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyKeyStatusTimestamp:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyKeyStatusTimestamp:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyValueAllowContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyValueAllowAll:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyValueAllowUsers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyValueDisallowContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyValueDisallowAll:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyValueDisallowUsers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyValueAllowContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyValueAllowAll:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyValueAllowUsers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyValueDisallowContacts:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyValueDisallowAll:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyValueDisallowUsers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAccountPrivacyRules:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAccountDaysTTL:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateUserPhone:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDisabledFeature:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocumentAttributeImageSize:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocumentAttributeAnimated:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocumentAttributeSticker:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocumentAttributeVideo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocumentAttributeAudio:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocumentAttributeFilename:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesStickersNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStickerPack:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesAllStickersNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesAllStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAccountNoPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAccountPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateReadHistoryInbox:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateReadHistoryOutbox:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesAffectedMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactLinkUnknown:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactLinkNone:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactLinkHasPhone:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactLinkContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateWebPage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredWebPageEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredWebPagePending:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredWebPage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaWebPage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthorization:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAccountAuthorizations:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAccountPasswordSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAccountPasswordInputSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthPasswordRecovery:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaVenue:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaVenue:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredReceivedNotifyMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatInviteEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatInviteExported:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatInviteAlready:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChatJoinedByLink:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateReadMessagesContents:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputStickerSetEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputStickerSetID:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputStickerSetShortName:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotCommand:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButton:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButtonRow:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredReplyKeyboardHide:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredReplyKeyboardForceReply:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredReplyKeyboardMarkup:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterUrl:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPeerUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputUser:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityUnknown:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityMention:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityHashtag:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityBotCommand:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityUrl:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityEmail:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityBold:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityItalic:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityCode:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityPre:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityTextUrl:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateShortSentMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPeerChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPeerChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelForbidden:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelFull:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChannelCreate:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesChannelMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChannelTooLong:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateNewChannelMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateReadChannelInbox:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateDeleteChannelMessages:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChannelMessageViews:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputChannelEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsResolvedPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageRange:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesChannelDifferenceEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesChannelDifferenceTooLong:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesChannelDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelMessagesFilterEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelMessagesFilter:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipant:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantSelf:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantCreator:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantsRecent:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantsAdmins:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantsKicked:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelsChannelParticipants:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelsChannelParticipant:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTrue:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatParticipantCreator:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChatParticipantAdmin:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChatAdmins:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChatParticipantAdmin:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChatMigrateTo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionChannelMigrateFrom:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantsBots:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputReportReasonSpam:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputReportReasonViolence:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputReportReasonPornography:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputReportReasonOther:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateNewStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateStickerSetsOrder:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateStickerSets:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredHelpTermsOfService:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredFoundGif:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaGifExternal:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesFoundGifs:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterGif:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateSavedGifs:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateBotInlineQuery:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredFoundGifCached:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesSavedGifsNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesSavedGifs:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineMessageMediaAuto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineMessageText:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotInlineMessageMediaAuto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotInlineMessageText:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotInlineResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesBotResults:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterVoice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterMusic:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateBotInlineSend:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyKeyChatInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyKeyChatInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateEditChannelMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredExportedMessageLink:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageFwdHeader:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionPinMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPeerSettings:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChannelPinnedMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButtonUrl:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButtonCallback:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButtonRequestPhone:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButtonRequestGeoLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthCodeTypeSms:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthCodeTypeCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthCodeTypeFlashCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthSentCodeTypeApp:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthSentCodeTypeSms:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthSentCodeTypeCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAuthSentCodeTypeFlashCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButtonSwitchInline:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredReplyInlineMarkup:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesBotCallbackAnswer:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateBotCallbackQuery:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesMessageEditData:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateEditMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineMessageMediaGeo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineMessageMediaVenue:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineMessageMediaContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotInlineMessageMediaGeo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotInlineMessageMediaVenue:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotInlineMessageMediaContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineResultPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineResultDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredBotInlineMediaResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineMessageID:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateInlineBotCallbackQuery:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInlineBotSwitchPM:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageEntityMentionName:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessageEntityMentionName:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesPeerDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTopPeer:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTopPeerCategoryBotsPM:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTopPeerCategoryBotsInline:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTopPeerCategoryCorrespondents:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTopPeerCategoryGroups:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTopPeerCategoryChannels:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTopPeerCategoryPeers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsTopPeersNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredContactsTopPeers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterChatPhotos:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateReadChannelOutbox:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateDraftMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDraftMessageEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDraftMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionHistoryClear:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateReadFeaturedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateRecentStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesFeaturedStickersNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesFeaturedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesRecentStickersNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesRecentStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesArchivedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesStickerSetInstallResultSuccess:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesStickerSetInstallResultArchive:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStickerSetCovered:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaPhotoExternal:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaDocumentExternal:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatePtsChanged:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionGameScore:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDocumentAttributeHasStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButtonGame:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredStickerSetMultiCovered:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMaskCoords:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputStickeredMediaPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputStickeredMediaDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaGame:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaGame:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineMessageGame:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputBotInlineResultGame:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredGame:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputGameID:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputGameShortName:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredHighScore:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesHighScores:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesChatsSlice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChannelWebPage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatesDifferenceTooLong:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageGamePlayAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredWebPageNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextPlain:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextBold:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextItalic:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextUnderline:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextStrike:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextFixed:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextUrl:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextEmail:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTextConcat:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockTitle:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockSubtitle:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockAuthorDate:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockHeader:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockSubheader:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockParagraph:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockPreformatted:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockFooter:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockDivider:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockList:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockBlockquote:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockPullquote:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockPhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockVideo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockCover:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockEmbed:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockEmbedPost:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockSlideshow:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPagePart:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageFull:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatePhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateDialogPinned:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdatePinnedDialogs:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPrivacyKeyPhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPrivacyKeyPhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockUnsupported:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockAnchor:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockCollage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallEmpty:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallWaiting:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallRequested:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallDiscarded:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneConnection:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallProtocol:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhonePhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallDiscardReasonMissed:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallDiscardReasonDisconnect:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallDiscardReasonHangup:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallDiscardReasonBusy:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterPhoneCalls:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionPhoneCall:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInvoice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMediaInvoice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionPaymentSentMe:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageMediaInvoice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredKeyboardButtonBuy:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionPaymentSent:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentsPaymentForm:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPostAddress:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentRequestedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateBotWebhookJSON:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateBotWebhookJSONQuery:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateBotShippingQuery:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateBotPrecheckoutQuery:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredDataJSON:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredLabeledPrice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentCharge:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentSavedCredentialsCard:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredWebDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputWebDocument:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputWebFileLocation:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUploadWebFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentsValidatedRequestedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentsPaymentResult:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentsPaymentVerficationNeeded:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentsPaymentReceipt:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPaymentsSavedInfo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPaymentCredentialsSaved:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputPaymentCredentials:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredAccountTmpPassword:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredShippingOption:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPhoneCallAccepted:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterRoundVoice:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterRoundVideo:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUploadFileCdnRedirect:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageRecordRoundAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredSendMessageUploadRoundAction:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUploadCdnFileReuploadNeeded:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUploadCdnFile:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredCdnPublicKey:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredCdnConfig:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateLangPackTooLong:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateLangPack:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockChannel:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputStickerSetItem:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredLangPackString:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredLangPackStringPluralized:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredLangPackStringDeleted:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredLangPackDifference:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredLangPackLanguage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantAdmin:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantBanned:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantsBanned:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelParticipantsSearch:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredTopPeerCategoryPhoneCalls:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPageBlockAudio:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminRights:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelBannedRights:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionChangeTitle:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionChangeAbout:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionChangeUsername:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionChangePhoto:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionToggleInvites:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionToggleSignatures:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionUpdatePinned:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionEditMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionDeleteMessage:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionParticipantJoin:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionParticipantLeave:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionParticipantInvite:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionParticipantToggleBan:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionParticipantToggleAdmin:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEvent:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelsAdminLogResults:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventsFilter:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessageActionScreenshotTaken:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredPopularContact:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredCdnFileHash:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterMyMentions:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredInputMessagesFilterMyMentionsUnread:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateContactsReset:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredChannelAdminLogEventActionChangeStickerSet:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateFavedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesFavedStickers:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredMessagesFavedStickersNotModified:
		marshaled, err = ptypes.MarshalAny(x)
	case *PredUpdateChannelReadMessagesContents:
		marshaled, err = ptypes.MarshalAny(x)
	}
	if err != nil {
		return nil
	}
	return marshaled
}

// Unpacker from gRPC Any to ReqXXX
func unpack(x *any.Any) TL {
	if x == nil {
		return nil
	}
	splits := strings.Split(x.TypeUrl, ".")
	if len(splits) < 1 {
		return nil
	}
	typeString := splits[len(splits)-1]
	switch typeString {
	case "ReqInvokeAfterMsg":
		u := &ReqInvokeAfterMsg{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqInvokeAfterMsgs":
		u := &ReqInvokeAfterMsgs{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthCheckPhone":
		u := &ReqAuthCheckPhone{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthSendCode":
		u := &ReqAuthSendCode{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthSignUp":
		u := &ReqAuthSignUp{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthSignIn":
		u := &ReqAuthSignIn{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthLogOut":
		u := &ReqAuthLogOut{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthResetAuthorizations":
		u := &ReqAuthResetAuthorizations{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthSendInvites":
		u := &ReqAuthSendInvites{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthExportAuthorization":
		u := &ReqAuthExportAuthorization{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthImportAuthorization":
		u := &ReqAuthImportAuthorization{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountRegisterDevice":
		u := &ReqAccountRegisterDevice{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountUnregisterDevice":
		u := &ReqAccountUnregisterDevice{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountUpdateNotifySettings":
		u := &ReqAccountUpdateNotifySettings{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountGetNotifySettings":
		u := &ReqAccountGetNotifySettings{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountResetNotifySettings":
		u := &ReqAccountResetNotifySettings{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountUpdateProfile":
		u := &ReqAccountUpdateProfile{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountUpdateStatus":
		u := &ReqAccountUpdateStatus{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountGetWallPapers":
		u := &ReqAccountGetWallPapers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUsersGetUsers":
		u := &ReqUsersGetUsers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUsersGetFullUser":
		u := &ReqUsersGetFullUser{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsGetStatuses":
		u := &ReqContactsGetStatuses{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsGetContacts":
		u := &ReqContactsGetContacts{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsImportContacts":
		u := &ReqContactsImportContacts{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsSearch":
		u := &ReqContactsSearch{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsDeleteContact":
		u := &ReqContactsDeleteContact{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsDeleteContacts":
		u := &ReqContactsDeleteContacts{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsBlock":
		u := &ReqContactsBlock{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsUnblock":
		u := &ReqContactsUnblock{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsGetBlocked":
		u := &ReqContactsGetBlocked{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetMessages":
		u := &ReqMessagesGetMessages{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetDialogs":
		u := &ReqMessagesGetDialogs{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetHistory":
		u := &ReqMessagesGetHistory{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSearch":
		u := &ReqMessagesSearch{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReadHistory":
		u := &ReqMessagesReadHistory{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesDeleteHistory":
		u := &ReqMessagesDeleteHistory{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesDeleteMessages":
		u := &ReqMessagesDeleteMessages{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReceivedMessages":
		u := &ReqMessagesReceivedMessages{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSetTyping":
		u := &ReqMessagesSetTyping{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSendMessage":
		u := &ReqMessagesSendMessage{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSendMedia":
		u := &ReqMessagesSendMedia{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesForwardMessages":
		u := &ReqMessagesForwardMessages{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetChats":
		u := &ReqMessagesGetChats{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetFullChat":
		u := &ReqMessagesGetFullChat{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesEditChatTitle":
		u := &ReqMessagesEditChatTitle{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesEditChatPhoto":
		u := &ReqMessagesEditChatPhoto{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesAddChatUser":
		u := &ReqMessagesAddChatUser{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesDeleteChatUser":
		u := &ReqMessagesDeleteChatUser{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesCreateChat":
		u := &ReqMessagesCreateChat{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUpdatesGetState":
		u := &ReqUpdatesGetState{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUpdatesGetDifference":
		u := &ReqUpdatesGetDifference{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhotosUpdateProfilePhoto":
		u := &ReqPhotosUpdateProfilePhoto{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhotosUploadProfilePhoto":
		u := &ReqPhotosUploadProfilePhoto{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUploadSaveFilePart":
		u := &ReqUploadSaveFilePart{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUploadGetFile":
		u := &ReqUploadGetFile{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpGetConfig":
		u := &ReqHelpGetConfig{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpGetNearestDc":
		u := &ReqHelpGetNearestDc{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpGetAppUpdate":
		u := &ReqHelpGetAppUpdate{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpSaveAppLog":
		u := &ReqHelpSaveAppLog{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpGetInviteText":
		u := &ReqHelpGetInviteText{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhotosDeletePhotos":
		u := &ReqPhotosDeletePhotos{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhotosGetUserPhotos":
		u := &ReqPhotosGetUserPhotos{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesForwardMessage":
		u := &ReqMessagesForwardMessage{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetDhConfig":
		u := &ReqMessagesGetDhConfig{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesRequestEncryption":
		u := &ReqMessagesRequestEncryption{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesAcceptEncryption":
		u := &ReqMessagesAcceptEncryption{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesDiscardEncryption":
		u := &ReqMessagesDiscardEncryption{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSetEncryptedTyping":
		u := &ReqMessagesSetEncryptedTyping{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReadEncryptedHistory":
		u := &ReqMessagesReadEncryptedHistory{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSendEncrypted":
		u := &ReqMessagesSendEncrypted{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSendEncryptedFile":
		u := &ReqMessagesSendEncryptedFile{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSendEncryptedService":
		u := &ReqMessagesSendEncryptedService{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReceivedQueue":
		u := &ReqMessagesReceivedQueue{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUploadSaveBigFilePart":
		u := &ReqUploadSaveBigFilePart{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqInitConnection":
		u := &ReqInitConnection{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpGetSupport":
		u := &ReqHelpGetSupport{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthBindTempAuthKey":
		u := &ReqAuthBindTempAuthKey{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsExportCard":
		u := &ReqContactsExportCard{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsImportCard":
		u := &ReqContactsImportCard{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReadMessageContents":
		u := &ReqMessagesReadMessageContents{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountCheckUsername":
		u := &ReqAccountCheckUsername{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountUpdateUsername":
		u := &ReqAccountUpdateUsername{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountGetPrivacy":
		u := &ReqAccountGetPrivacy{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountSetPrivacy":
		u := &ReqAccountSetPrivacy{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountDeleteAccount":
		u := &ReqAccountDeleteAccount{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountGetAccountTTL":
		u := &ReqAccountGetAccountTTL{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountSetAccountTTL":
		u := &ReqAccountSetAccountTTL{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqInvokeWithLayer":
		u := &ReqInvokeWithLayer{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsResolveUsername":
		u := &ReqContactsResolveUsername{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountSendChangePhoneCode":
		u := &ReqAccountSendChangePhoneCode{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountChangePhone":
		u := &ReqAccountChangePhone{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetAllStickers":
		u := &ReqMessagesGetAllStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountUpdateDeviceLocked":
		u := &ReqAccountUpdateDeviceLocked{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountGetPassword":
		u := &ReqAccountGetPassword{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthCheckPassword":
		u := &ReqAuthCheckPassword{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetWebPagePreview":
		u := &ReqMessagesGetWebPagePreview{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountGetAuthorizations":
		u := &ReqAccountGetAuthorizations{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountResetAuthorization":
		u := &ReqAccountResetAuthorization{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountGetPasswordSettings":
		u := &ReqAccountGetPasswordSettings{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountUpdatePasswordSettings":
		u := &ReqAccountUpdatePasswordSettings{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthRequestPasswordRecovery":
		u := &ReqAuthRequestPasswordRecovery{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthRecoverPassword":
		u := &ReqAuthRecoverPassword{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqInvokeWithoutUpdates":
		u := &ReqInvokeWithoutUpdates{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesExportChatInvite":
		u := &ReqMessagesExportChatInvite{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesCheckChatInvite":
		u := &ReqMessagesCheckChatInvite{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesImportChatInvite":
		u := &ReqMessagesImportChatInvite{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetStickerSet":
		u := &ReqMessagesGetStickerSet{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesInstallStickerSet":
		u := &ReqMessagesInstallStickerSet{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesUninstallStickerSet":
		u := &ReqMessagesUninstallStickerSet{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthImportBotAuthorization":
		u := &ReqAuthImportBotAuthorization{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesStartBot":
		u := &ReqMessagesStartBot{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpGetAppChangelog":
		u := &ReqHelpGetAppChangelog{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReportSpam":
		u := &ReqMessagesReportSpam{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetMessagesViews":
		u := &ReqMessagesGetMessagesViews{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUpdatesGetChannelDifference":
		u := &ReqUpdatesGetChannelDifference{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsReadHistory":
		u := &ReqChannelsReadHistory{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsDeleteMessages":
		u := &ReqChannelsDeleteMessages{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsDeleteUserHistory":
		u := &ReqChannelsDeleteUserHistory{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsReportSpam":
		u := &ReqChannelsReportSpam{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsGetMessages":
		u := &ReqChannelsGetMessages{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsGetParticipants":
		u := &ReqChannelsGetParticipants{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsGetParticipant":
		u := &ReqChannelsGetParticipant{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsGetChannels":
		u := &ReqChannelsGetChannels{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsGetFullChannel":
		u := &ReqChannelsGetFullChannel{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsCreateChannel":
		u := &ReqChannelsCreateChannel{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsEditAbout":
		u := &ReqChannelsEditAbout{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsEditAdmin":
		u := &ReqChannelsEditAdmin{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsEditTitle":
		u := &ReqChannelsEditTitle{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsEditPhoto":
		u := &ReqChannelsEditPhoto{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsCheckUsername":
		u := &ReqChannelsCheckUsername{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsUpdateUsername":
		u := &ReqChannelsUpdateUsername{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsJoinChannel":
		u := &ReqChannelsJoinChannel{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsLeaveChannel":
		u := &ReqChannelsLeaveChannel{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsInviteToChannel":
		u := &ReqChannelsInviteToChannel{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsExportInvite":
		u := &ReqChannelsExportInvite{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsDeleteChannel":
		u := &ReqChannelsDeleteChannel{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesToggleChatAdmins":
		u := &ReqMessagesToggleChatAdmins{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesEditChatAdmin":
		u := &ReqMessagesEditChatAdmin{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesMigrateChat":
		u := &ReqMessagesMigrateChat{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSearchGlobal":
		u := &ReqMessagesSearchGlobal{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountReportPeer":
		u := &ReqAccountReportPeer{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReorderStickerSets":
		u := &ReqMessagesReorderStickerSets{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpGetTermsOfService":
		u := &ReqHelpGetTermsOfService{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetDocumentByHash":
		u := &ReqMessagesGetDocumentByHash{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSearchGifs":
		u := &ReqMessagesSearchGifs{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetSavedGifs":
		u := &ReqMessagesGetSavedGifs{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSaveGif":
		u := &ReqMessagesSaveGif{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetInlineBotResults":
		u := &ReqMessagesGetInlineBotResults{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSetInlineBotResults":
		u := &ReqMessagesSetInlineBotResults{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSendInlineBotResult":
		u := &ReqMessagesSendInlineBotResult{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsToggleInvites":
		u := &ReqChannelsToggleInvites{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsExportMessageLink":
		u := &ReqChannelsExportMessageLink{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsToggleSignatures":
		u := &ReqChannelsToggleSignatures{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesHideReportSpam":
		u := &ReqMessagesHideReportSpam{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetPeerSettings":
		u := &ReqMessagesGetPeerSettings{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsUpdatePinnedMessage":
		u := &ReqChannelsUpdatePinnedMessage{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthResendCode":
		u := &ReqAuthResendCode{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthCancelCode":
		u := &ReqAuthCancelCode{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetMessageEditData":
		u := &ReqMessagesGetMessageEditData{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesEditMessage":
		u := &ReqMessagesEditMessage{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesEditInlineBotMessage":
		u := &ReqMessagesEditInlineBotMessage{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetBotCallbackAnswer":
		u := &ReqMessagesGetBotCallbackAnswer{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSetBotCallbackAnswer":
		u := &ReqMessagesSetBotCallbackAnswer{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsGetTopPeers":
		u := &ReqContactsGetTopPeers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsResetTopPeerRating":
		u := &ReqContactsResetTopPeerRating{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetPeerDialogs":
		u := &ReqMessagesGetPeerDialogs{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSaveDraft":
		u := &ReqMessagesSaveDraft{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetAllDrafts":
		u := &ReqMessagesGetAllDrafts{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountSendConfirmPhoneCode":
		u := &ReqAccountSendConfirmPhoneCode{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountConfirmPhone":
		u := &ReqAccountConfirmPhone{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetFeaturedStickers":
		u := &ReqMessagesGetFeaturedStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReadFeaturedStickers":
		u := &ReqMessagesReadFeaturedStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetRecentStickers":
		u := &ReqMessagesGetRecentStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSaveRecentSticker":
		u := &ReqMessagesSaveRecentSticker{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesClearRecentStickers":
		u := &ReqMessagesClearRecentStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetArchivedStickers":
		u := &ReqMessagesGetArchivedStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsGetAdminedPublicChannels":
		u := &ReqChannelsGetAdminedPublicChannels{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAuthDropTempAuthKeys":
		u := &ReqAuthDropTempAuthKeys{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSetGameScore":
		u := &ReqMessagesSetGameScore{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSetInlineGameScore":
		u := &ReqMessagesSetInlineGameScore{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetMaskStickers":
		u := &ReqMessagesGetMaskStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetAttachedStickers":
		u := &ReqMessagesGetAttachedStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetGameHighScores":
		u := &ReqMessagesGetGameHighScores{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetInlineGameHighScores":
		u := &ReqMessagesGetInlineGameHighScores{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetCommonChats":
		u := &ReqMessagesGetCommonChats{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetAllChats":
		u := &ReqMessagesGetAllChats{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpSetBotUpdatesStatus":
		u := &ReqHelpSetBotUpdatesStatus{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetWebPage":
		u := &ReqMessagesGetWebPage{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesToggleDialogPin":
		u := &ReqMessagesToggleDialogPin{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReorderPinnedDialogs":
		u := &ReqMessagesReorderPinnedDialogs{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetPinnedDialogs":
		u := &ReqMessagesGetPinnedDialogs{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhoneRequestCall":
		u := &ReqPhoneRequestCall{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhoneAcceptCall":
		u := &ReqPhoneAcceptCall{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhoneDiscardCall":
		u := &ReqPhoneDiscardCall{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhoneReceivedCall":
		u := &ReqPhoneReceivedCall{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesReportEncryptedSpam":
		u := &ReqMessagesReportEncryptedSpam{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPaymentsGetPaymentForm":
		u := &ReqPaymentsGetPaymentForm{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPaymentsSendPaymentForm":
		u := &ReqPaymentsSendPaymentForm{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqAccountGetTmpPassword":
		u := &ReqAccountGetTmpPassword{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSetBotShippingResults":
		u := &ReqMessagesSetBotShippingResults{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSetBotPrecheckoutResults":
		u := &ReqMessagesSetBotPrecheckoutResults{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUploadGetWebFile":
		u := &ReqUploadGetWebFile{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqBotsSendCustomRequest":
		u := &ReqBotsSendCustomRequest{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqBotsAnswerWebhookJSONQuery":
		u := &ReqBotsAnswerWebhookJSONQuery{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPaymentsGetPaymentReceipt":
		u := &ReqPaymentsGetPaymentReceipt{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPaymentsValidateRequestedInfo":
		u := &ReqPaymentsValidateRequestedInfo{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPaymentsGetSavedInfo":
		u := &ReqPaymentsGetSavedInfo{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPaymentsClearSavedInfo":
		u := &ReqPaymentsClearSavedInfo{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhoneGetCallConfig":
		u := &ReqPhoneGetCallConfig{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhoneConfirmCall":
		u := &ReqPhoneConfirmCall{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhoneSetCallRating":
		u := &ReqPhoneSetCallRating{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqPhoneSaveCallDebug":
		u := &ReqPhoneSaveCallDebug{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUploadGetCdnFile":
		u := &ReqUploadGetCdnFile{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUploadReuploadCdnFile":
		u := &ReqUploadReuploadCdnFile{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqHelpGetCdnConfig":
		u := &ReqHelpGetCdnConfig{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesUploadMedia":
		u := &ReqMessagesUploadMedia{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqStickersCreateStickerSet":
		u := &ReqStickersCreateStickerSet{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqLangpackGetLangPack":
		u := &ReqLangpackGetLangPack{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqLangpackGetStrings":
		u := &ReqLangpackGetStrings{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqLangpackGetDifference":
		u := &ReqLangpackGetDifference{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqLangpackGetLanguages":
		u := &ReqLangpackGetLanguages{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsEditBanned":
		u := &ReqChannelsEditBanned{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsGetAdminLog":
		u := &ReqChannelsGetAdminLog{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqStickersRemoveStickerFromSet":
		u := &ReqStickersRemoveStickerFromSet{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqStickersChangeStickerPosition":
		u := &ReqStickersChangeStickerPosition{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqStickersAddStickerToSet":
		u := &ReqStickersAddStickerToSet{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesSendScreenshotNotification":
		u := &ReqMessagesSendScreenshotNotification{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqUploadGetCdnFileHashes":
		u := &ReqUploadGetCdnFileHashes{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetUnreadMentions":
		u := &ReqMessagesGetUnreadMentions{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesFaveSticker":
		u := &ReqMessagesFaveSticker{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsSetStickers":
		u := &ReqChannelsSetStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqContactsResetSaved":
		u := &ReqContactsResetSaved{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqMessagesGetFavedStickers":
		u := &ReqMessagesGetFavedStickers{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	case "ReqChannelsReadMessageContents":
		u := &ReqChannelsReadMessageContents{}
		err := ptypes.UnmarshalAny(x, u)
		if err != nil {
			return nil
		}
		return u
	}
	return nil
}
