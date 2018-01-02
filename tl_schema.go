package mtproto

import "fmt"

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
	crc_storage_fileUnknown                              = 0xaa963b05
	crc_storage_fileJpeg                                 = 0x007efe0e
	crc_storage_fileGif                                  = 0xcae1aadf
	crc_storage_filePng                                  = 0x0a4f63c0
	crc_storage_fileMp3                                  = 0x528a0677
	crc_storage_fileMov                                  = 0x4b09ebbc
	crc_storage_filePartial                              = 0x40bc6f52
	crc_storage_fileMp4                                  = 0xb3cea0e4
	crc_storage_fileWebp                                 = 0x1081464c
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
	crc_auth_checkedPhone                                = 0x811ea28e
	crc_auth_sentCode                                    = 0x5e002502
	crc_auth_authorization                               = 0xcd050916
	crc_auth_exportedAuthorization                       = 0xdf969c2d
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
	crc_contacts_link                                    = 0x3ace484c
	crc_contacts_contacts                                = 0xeae87e42
	crc_contacts_contactsNotModified                     = 0xb74ba9d2
	crc_contacts_importedContacts                        = 0x77d01c3b
	crc_contacts_blocked                                 = 0x1c138d15
	crc_contacts_blockedSlice                            = 0x900802a1
	crc_contacts_found                                   = 0x1aa1f784
	crc_messages_dialogs                                 = 0x15ba6c40
	crc_messages_dialogsSlice                            = 0x71e094f3
	crc_messages_messages                                = 0x8c718e87
	crc_messages_messagesSlice                           = 0x0b446ae3
	crc_messages_chats                                   = 0x64ff9fd5
	crc_messages_chatFull                                = 0xe5d7d19c
	crc_messages_affectedHistory                         = 0xb45c69d1
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
	crc_updates_state                                    = 0xa56c2a3e
	crc_updates_differenceEmpty                          = 0x5d75a138
	crc_updates_difference                               = 0x00f49ca0
	crc_updates_differenceSlice                          = 0xa8fb1981
	crc_updatesTooLong                                   = 0xe317af7e
	crc_updateShortMessage                               = 0x914fbf11
	crc_updateShortChatMessage                           = 0x16812688
	crc_updateShort                                      = 0x78d4dec1
	crc_updatesCombined                                  = 0x725b04c3
	crc_updates                                          = 0x74ae4240
	crc_photos_photo                                     = 0x20212ca8
	crc_upload_file                                      = 0x096a18d5
	crc_dcOption                                         = 0x05d8c6cc
	crc_config                                           = 0x8df376a4
	crc_nearestDc                                        = 0x8e1a1775
	crc_help_appUpdate                                   = 0x8987f311
	crc_help_noAppUpdate                                 = 0xc45a6536
	crc_help_inviteText                                  = 0x18cb9f78
	crc_inputPeerNotifyEventsEmpty                       = 0xf03064d8
	crc_inputPeerNotifyEventsAll                         = 0xe86a2c74
	crc_photos_photos                                    = 0x8dca6aa5
	crc_photos_photosSlice                               = 0x15051f54
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
	crc_messages_dhConfigNotModified                     = 0xc0e24635
	crc_messages_dhConfig                                = 0x2c221edd
	crc_messages_sentEncryptedMessage                    = 0x560f8935
	crc_messages_sentEncryptedFile                       = 0x9493ff32
	crc_inputFileBig                                     = 0xfa4f0bb5
	crc_inputEncryptedFileBigUploaded                    = 0x2dc173c8
	crc_storage_filePdf                                  = 0xae1e508d
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
	crc_help_support                                     = 0x17c6b5f6
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
	crc_account_privacyRules                             = 0x554abb6f
	crc_accountDaysTTL                                   = 0xb8d0afdf
	crc_updateUserPhone                                  = 0x12b9417b
	crc_disabledFeature                                  = 0xae636f24
	crc_documentAttributeImageSize                       = 0x6c37c15c
	crc_documentAttributeAnimated                        = 0x11b58939
	crc_documentAttributeSticker                         = 0x6319d612
	crc_documentAttributeVideo                           = 0x0ef02ce6
	crc_documentAttributeAudio                           = 0x9852f9c6
	crc_documentAttributeFilename                        = 0x15590068
	crc_messages_stickersNotModified                     = 0xf1749a22
	crc_messages_stickers                                = 0x8a8ecd32
	crc_stickerPack                                      = 0x12b299d4
	crc_messages_allStickersNotModified                  = 0xe86602c3
	crc_messages_allStickers                             = 0xedfd405f
	crc_account_noPassword                               = 0x96dabc18
	crc_account_password                                 = 0x7c18141c
	crc_updateReadHistoryInbox                           = 0x9961fd5c
	crc_updateReadHistoryOutbox                          = 0x2f2f21bf
	crc_messages_affectedMessages                        = 0x84d19185
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
	crc_account_authorizations                           = 0x1250abde
	crc_account_passwordSettings                         = 0xb7b72ab3
	crc_account_passwordInputSettings                    = 0x86916deb
	crc_auth_passwordRecovery                            = 0x137948a5
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
	crc_messages_stickerSet                              = 0xb60a24a6
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
	crc_messages_channelMessages                         = 0x99262e37
	crc_updateChannelTooLong                             = 0xeb0467fb
	crc_updateChannel                                    = 0xb6d45656
	crc_updateNewChannelMessage                          = 0x62ba04d9
	crc_updateReadChannelInbox                           = 0x4214f37f
	crc_updateDeleteChannelMessages                      = 0xc37521c9
	crc_updateChannelMessageViews                        = 0x98a12b4b
	crc_inputChannelEmpty                                = 0xee8c1e86
	crc_inputChannel                                     = 0xafeb712e
	crc_contacts_resolvedPeer                            = 0x7f077ad9
	crc_messageRange                                     = 0x0ae30253
	crc_updates_channelDifferenceEmpty                   = 0x3e11affb
	crc_updates_channelDifferenceTooLong                 = 0x6a9d7b35
	crc_updates_channelDifference                        = 0x2064674e
	crc_channelMessagesFilterEmpty                       = 0x94d42ee7
	crc_channelMessagesFilter                            = 0xcd77d957
	crc_channelParticipant                               = 0x15ebac1d
	crc_channelParticipantSelf                           = 0xa3289a6d
	crc_channelParticipantCreator                        = 0xe3e2e1f9
	crc_channelParticipantsRecent                        = 0xde3f3c79
	crc_channelParticipantsAdmins                        = 0xb4608969
	crc_channelParticipantsKicked                        = 0xa3b54985
	crc_channels_channelParticipants                     = 0xf56ee2a8
	crc_channels_channelParticipant                      = 0xd0d9b163
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
	crc_help_termsOfService                              = 0xf1ee3e90
	crc_foundGif                                         = 0x162ecc1f
	crc_inputMediaGifExternal                            = 0x4843b0fd
	crc_messages_foundGifs                               = 0x450a1c0a
	crc_inputMessagesFilterGif                           = 0xffc86587
	crc_updateSavedGifs                                  = 0x9375341e
	crc_updateBotInlineQuery                             = 0x54826690
	crc_foundGifCached                                   = 0x9c750409
	crc_messages_savedGifsNotModified                    = 0xe8025ca2
	crc_messages_savedGifs                               = 0x2e0709a5
	crc_inputBotInlineMessageMediaAuto                   = 0x292fed13
	crc_inputBotInlineMessageText                        = 0x3dcd7a87
	crc_inputBotInlineResult                             = 0x2cbbe15a
	crc_botInlineMessageMediaAuto                        = 0x0a74b15b
	crc_botInlineMessageText                             = 0x8c7f65e2
	crc_botInlineResult                                  = 0x9bebaeb9
	crc_messages_botResults                              = 0xccd3563d
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
	crc_auth_codeTypeSms                                 = 0x72a3158c
	crc_auth_codeTypeCall                                = 0x741cd3e3
	crc_auth_codeTypeFlashCall                           = 0x226ccefb
	crc_auth_sentCodeTypeApp                             = 0x3dbb5986
	crc_auth_sentCodeTypeSms                             = 0xc000bba2
	crc_auth_sentCodeTypeCall                            = 0x5353e5a7
	crc_auth_sentCodeTypeFlashCall                       = 0xab03c6d9
	crc_keyboardButtonSwitchInline                       = 0x0568a748
	crc_replyInlineMarkup                                = 0x48a30254
	crc_messages_botCallbackAnswer                       = 0x36585ea4
	crc_updateBotCallbackQuery                           = 0xe73547e1
	crc_messages_messageEditData                         = 0x26b5dde6
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
	crc_messages_peerDialogs                             = 0x3371c354
	crc_topPeer                                          = 0xedcdc05b
	crc_topPeerCategoryBotsPM                            = 0xab661b5b
	crc_topPeerCategoryBotsInline                        = 0x148677e2
	crc_topPeerCategoryCorrespondents                    = 0x0637b7ed
	crc_topPeerCategoryGroups                            = 0xbd17a14a
	crc_topPeerCategoryChannels                          = 0x161d9628
	crc_topPeerCategoryPeers                             = 0xfb834291
	crc_contacts_topPeersNotModified                     = 0xde266ef5
	crc_contacts_topPeers                                = 0x70b772a8
	crc_inputMessagesFilterChatPhotos                    = 0x3a20ecb8
	crc_updateReadChannelOutbox                          = 0x25d6c9c7
	crc_updateDraftMessage                               = 0xee2bb969
	crc_draftMessageEmpty                                = 0xba4baec5
	crc_draftMessage                                     = 0xfd8e711f
	crc_messageActionHistoryClear                        = 0x9fbab604
	crc_updateReadFeaturedStickers                       = 0x571d2742
	crc_updateRecentStickers                             = 0x9a422c20
	crc_messages_featuredStickersNotModified             = 0x04ede3cf
	crc_messages_featuredStickers                        = 0xf89d88e5
	crc_messages_recentStickersNotModified               = 0x0b17f890
	crc_messages_recentStickers                          = 0x5ce20970
	crc_messages_archivedStickers                        = 0x4fcba9c8
	crc_messages_stickerSetInstallResultSuccess          = 0x38641628
	crc_messages_stickerSetInstallResultArchive          = 0x35e410a8
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
	crc_messages_highScores                              = 0x9a3bfd99
	crc_messages_chatsSlice                              = 0x9cd81144
	crc_updateChannelWebPage                             = 0x40771900
	crc_updates_differenceTooLong                        = 0x4afe8f6d
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
	crc_phone_phoneCall                                  = 0xec82e140
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
	crc_payments_paymentForm                             = 0x3f56aea3
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
	crc_upload_webFile                                   = 0x21e753bc
	crc_payments_validatedRequestedInfo                  = 0xd1451883
	crc_payments_paymentResult                           = 0x4e5f810d
	crc_payments_paymentVerficationNeeded                = 0x6b56b921
	crc_payments_paymentReceipt                          = 0x500911e1
	crc_payments_savedInfo                               = 0xfb8fe43c
	crc_inputPaymentCredentialsSaved                     = 0xc10eb2cf
	crc_inputPaymentCredentials                          = 0x3417d728
	crc_account_tmpPassword                              = 0xdb64fd34
	crc_shippingOption                                   = 0xb6213cdf
	crc_phoneCallAccepted                                = 0x6d003d3f
	crc_inputMessagesFilterRoundVoice                    = 0x7a7c17a4
	crc_inputMessagesFilterRoundVideo                    = 0xb549da53
	crc_upload_fileCdnRedirect                           = 0xea52fe5a
	crc_sendMessageRecordRoundAction                     = 0x88f27fbc
	crc_sendMessageUploadRoundAction                     = 0x243e1c66
	crc_upload_cdnFileReuploadNeeded                     = 0xeea8e46e
	crc_upload_cdnFile                                   = 0xa99fca4f
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
	crc_channels_adminLogResults                         = 0xed8af74d
	crc_channelAdminLogEventsFilter                      = 0xea107ae4
	crc_messageActionScreenshotTaken                     = 0x4792929b
	crc_popularContact                                   = 0x5ce14175
	crc_cdnFileHash                                      = 0x77eec38f
	crc_inputMessagesFilterMyMentions                    = 0xc1f8e69a
	crc_inputMessagesFilterMyMentionsUnread              = 0x46caf4a8
	crc_updateContactsReset                              = 0x7084a7be
	crc_channelAdminLogEventActionChangeStickerSet       = 0xb1c3caa7
	crc_updateFavedStickers                              = 0xe511996d
	crc_messages_favedStickers                           = 0xf37f2f16
	crc_messages_favedStickersNotModified                = 0x9e8fa6d3
	crc_updateChannelReadMessagesContents                = 0x89893b45
	crc_invokeAfterMsg                                   = 0xcb9f372d
	crc_invokeAfterMsgs                                  = 0x3dc4b4f0
	crc_auth_checkPhone                                  = 0x6fe51dfb
	crc_auth_sendCode                                    = 0x86aef0ec
	crc_auth_signUp                                      = 0x1b067634
	crc_auth_signIn                                      = 0xbcd51581
	crc_auth_logOut                                      = 0x5717da40
	crc_auth_resetAuthorizations                         = 0x9fab0d1a
	crc_auth_sendInvites                                 = 0x771c1d97
	crc_auth_exportAuthorization                         = 0xe5bfffcd
	crc_auth_importAuthorization                         = 0xe3ef9613
	crc_account_registerDevice                           = 0x637ea878
	crc_account_unregisterDevice                         = 0x65c55b40
	crc_account_updateNotifySettings                     = 0x84be5b93
	crc_account_getNotifySettings                        = 0x12b3ad31
	crc_account_resetNotifySettings                      = 0xdb7e1747
	crc_account_updateProfile                            = 0x78515775
	crc_account_updateStatus                             = 0x6628562c
	crc_account_getWallPapers                            = 0xc04cfac2
	crc_users_getUsers                                   = 0x0d91a548
	crc_users_getFullUser                                = 0xca30a5b1
	crc_contacts_getStatuses                             = 0xc4a353ee
	crc_contacts_getContacts                             = 0xc023849f
	crc_contacts_importContacts                          = 0x2c800be5
	crc_contacts_search                                  = 0x11f812d8
	crc_contacts_deleteContact                           = 0x8e953744
	crc_contacts_deleteContacts                          = 0x59ab389e
	crc_contacts_block                                   = 0x332b49fc
	crc_contacts_unblock                                 = 0xe54100bd
	crc_contacts_getBlocked                              = 0xf57c350f
	crc_messages_getMessages                             = 0x4222fa74
	crc_messages_getDialogs                              = 0x191ba9c5
	crc_messages_getHistory                              = 0xafa92846
	crc_messages_search                                  = 0x039e9ea0
	crc_messages_readHistory                             = 0x0e306d3a
	crc_messages_deleteHistory                           = 0x1c015b09
	crc_messages_deleteMessages                          = 0xe58e95d2
	crc_messages_receivedMessages                        = 0x05a954c0
	crc_messages_setTyping                               = 0xa3825e50
	crc_messages_sendMessage                             = 0xfa88427a
	crc_messages_sendMedia                               = 0xc8f16791
	crc_messages_forwardMessages                         = 0x708e0195
	crc_messages_getChats                                = 0x3c6aa187
	crc_messages_getFullChat                             = 0x3b831c66
	crc_messages_editChatTitle                           = 0xdc452855
	crc_messages_editChatPhoto                           = 0xca4c79d8
	crc_messages_addChatUser                             = 0xf9a0aa09
	crc_messages_deleteChatUser                          = 0xe0611f16
	crc_messages_createChat                              = 0x09cb126e
	crc_updates_getState                                 = 0xedd4882a
	crc_updates_getDifference                            = 0x25939651
	crc_photos_updateProfilePhoto                        = 0xf0bb5152
	crc_photos_uploadProfilePhoto                        = 0x4f32c098
	crc_upload_saveFilePart                              = 0xb304a621
	crc_upload_getFile                                   = 0xe3a6cfb5
	crc_help_getConfig                                   = 0xc4f9186b
	crc_help_getNearestDc                                = 0x1fb33026
	crc_help_getAppUpdate                                = 0xae2de196
	crc_help_saveAppLog                                  = 0x6f02f748
	crc_help_getInviteText                               = 0x4d392343
	crc_photos_deletePhotos                              = 0x87cf7f2f
	crc_photos_getUserPhotos                             = 0x91cd32a8
	crc_messages_forwardMessage                          = 0x33963bf9
	crc_messages_getDhConfig                             = 0x26cf8950
	crc_messages_requestEncryption                       = 0xf64daf43
	crc_messages_acceptEncryption                        = 0x3dbc0415
	crc_messages_discardEncryption                       = 0xedd923c5
	crc_messages_setEncryptedTyping                      = 0x791451ed
	crc_messages_readEncryptedHistory                    = 0x7f4b690a
	crc_messages_sendEncrypted                           = 0xa9776773
	crc_messages_sendEncryptedFile                       = 0x9a901b66
	crc_messages_sendEncryptedService                    = 0x32d439a4
	crc_messages_receivedQueue                           = 0x55a5bb66
	crc_upload_saveBigFilePart                           = 0xde7b673d
	crc_initConnection                                   = 0xc7481da6
	crc_help_getSupport                                  = 0x9cdf08cd
	crc_auth_bindTempAuthKey                             = 0xcdd42a05
	crc_contacts_exportCard                              = 0x84e53737
	crc_contacts_importCard                              = 0x4fe196fe
	crc_messages_readMessageContents                     = 0x36a73f77
	crc_account_checkUsername                            = 0x2714d86c
	crc_account_updateUsername                           = 0x3e0bdd7c
	crc_account_getPrivacy                               = 0xdadbc950
	crc_account_setPrivacy                               = 0xc9f81ce8
	crc_account_deleteAccount                            = 0x418d4e0b
	crc_account_getAccountTTL                            = 0x08fc711d
	crc_account_setAccountTTL                            = 0x2442485e
	crc_invokeWithLayer                                  = 0xda9b0d0d
	crc_contacts_resolveUsername                         = 0xf93ccba3
	crc_account_sendChangePhoneCode                      = 0x08e57deb
	crc_account_changePhone                              = 0x70c32edb
	crc_messages_getAllStickers                          = 0x1c9618b1
	crc_account_updateDeviceLocked                       = 0x38df3532
	crc_account_getPassword                              = 0x548a30f5
	crc_auth_checkPassword                               = 0x0a63011e
	crc_messages_getWebPagePreview                       = 0x25223e24
	crc_account_getAuthorizations                        = 0xe320c158
	crc_account_resetAuthorization                       = 0xdf77f3bc
	crc_account_getPasswordSettings                      = 0xbc8d11bb
	crc_account_updatePasswordSettings                   = 0xfa7c4b86
	crc_auth_requestPasswordRecovery                     = 0xd897bc66
	crc_auth_recoverPassword                             = 0x4ea56e92
	crc_invokeWithoutUpdates                             = 0xbf9459b7
	crc_messages_exportChatInvite                        = 0x7d885289
	crc_messages_checkChatInvite                         = 0x3eadb1bb
	crc_messages_importChatInvite                        = 0x6c50051c
	crc_messages_getStickerSet                           = 0x2619a90e
	crc_messages_installStickerSet                       = 0xc78fe460
	crc_messages_uninstallStickerSet                     = 0xf96e55de
	crc_auth_importBotAuthorization                      = 0x67a3ff2c
	crc_messages_startBot                                = 0xe6df7378
	crc_help_getAppChangelog                             = 0x9010ef6f
	crc_messages_reportSpam                              = 0xcf1592db
	crc_messages_getMessagesViews                        = 0xc4c8a55d
	crc_updates_getChannelDifference                     = 0x03173d78
	crc_channels_readHistory                             = 0xcc104937
	crc_channels_deleteMessages                          = 0x84c1fd4e
	crc_channels_deleteUserHistory                       = 0xd10dd71b
	crc_channels_reportSpam                              = 0xfe087810
	crc_channels_getMessages                             = 0x93d7b347
	crc_channels_getParticipants                         = 0x24d98f92
	crc_channels_getParticipant                          = 0x546dd7a6
	crc_channels_getChannels                             = 0x0a7f6bbb
	crc_channels_getFullChannel                          = 0x08736a09
	crc_channels_createChannel                           = 0xf4893d7f
	crc_channels_editAbout                               = 0x13e27f1e
	crc_channels_editAdmin                               = 0x20b88214
	crc_channels_editTitle                               = 0x566decd0
	crc_channels_editPhoto                               = 0xf12e57c9
	crc_channels_checkUsername                           = 0x10e6bd2c
	crc_channels_updateUsername                          = 0x3514b3de
	crc_channels_joinChannel                             = 0x24b524c5
	crc_channels_leaveChannel                            = 0xf836aa95
	crc_channels_inviteToChannel                         = 0x199f3a6c
	crc_channels_exportInvite                            = 0xc7560885
	crc_channels_deleteChannel                           = 0xc0111fe3
	crc_messages_toggleChatAdmins                        = 0xec8bd9e1
	crc_messages_editChatAdmin                           = 0xa9e69f2e
	crc_messages_migrateChat                             = 0x15a3b8e3
	crc_messages_searchGlobal                            = 0x9e3cacb0
	crc_account_reportPeer                               = 0xae189d5f
	crc_messages_reorderStickerSets                      = 0x78337739
	crc_help_getTermsOfService                           = 0x350170f3
	crc_messages_getDocumentByHash                       = 0x338e2464
	crc_messages_searchGifs                              = 0xbf9a776b
	crc_messages_getSavedGifs                            = 0x83bf3d52
	crc_messages_saveGif                                 = 0x327a30cb
	crc_messages_getInlineBotResults                     = 0x514e999d
	crc_messages_setInlineBotResults                     = 0xeb5ea206
	crc_messages_sendInlineBotResult                     = 0xb16e06fe
	crc_channels_toggleInvites                           = 0x49609307
	crc_channels_exportMessageLink                       = 0xc846d22d
	crc_channels_toggleSignatures                        = 0x1f69b606
	crc_messages_hideReportSpam                          = 0xa8f1709b
	crc_messages_getPeerSettings                         = 0x3672e09c
	crc_channels_updatePinnedMessage                     = 0xa72ded52
	crc_auth_resendCode                                  = 0x3ef1a9bf
	crc_auth_cancelCode                                  = 0x1f040578
	crc_messages_getMessageEditData                      = 0xfda68d36
	crc_messages_editMessage                             = 0xce91e4ca
	crc_messages_editInlineBotMessage                    = 0x130c2c85
	crc_messages_getBotCallbackAnswer                    = 0x810a9fec
	crc_messages_setBotCallbackAnswer                    = 0xd58f130a
	crc_contacts_getTopPeers                             = 0xd4982db5
	crc_contacts_resetTopPeerRating                      = 0x1ae373ac
	crc_messages_getPeerDialogs                          = 0x2d9776b9
	crc_messages_saveDraft                               = 0xbc39e14b
	crc_messages_getAllDrafts                            = 0x6a3f8d65
	crc_account_sendConfirmPhoneCode                     = 0x1516d7bd
	crc_account_confirmPhone                             = 0x5f2178c3
	crc_messages_getFeaturedStickers                     = 0x2dacca4f
	crc_messages_readFeaturedStickers                    = 0x5b118126
	crc_messages_getRecentStickers                       = 0x5ea192c9
	crc_messages_saveRecentSticker                       = 0x392718f8
	crc_messages_clearRecentStickers                     = 0x8999602d
	crc_messages_getArchivedStickers                     = 0x57f17692
	crc_channels_getAdminedPublicChannels                = 0x8d8d82d7
	crc_auth_dropTempAuthKeys                            = 0x8e48a188
	crc_messages_setGameScore                            = 0x8ef8ecc0
	crc_messages_setInlineGameScore                      = 0x15ad9f64
	crc_messages_getMaskStickers                         = 0x65b8c79f
	crc_messages_getAttachedStickers                     = 0xcc5b67cc
	crc_messages_getGameHighScores                       = 0xe822649d
	crc_messages_getInlineGameHighScores                 = 0x0f635e1b
	crc_messages_getCommonChats                          = 0x0d0a48c4
	crc_messages_getAllChats                             = 0xeba80ff0
	crc_help_setBotUpdatesStatus                         = 0xec22cfcd
	crc_messages_getWebPage                              = 0x32ca8f91
	crc_messages_toggleDialogPin                         = 0x3289be6a
	crc_messages_reorderPinnedDialogs                    = 0x959ff644
	crc_messages_getPinnedDialogs                        = 0xe254d64e
	crc_phone_requestCall                                = 0x5b95b3d4
	crc_phone_acceptCall                                 = 0x3bd2b4a0
	crc_phone_discardCall                                = 0x78d413a6
	crc_phone_receivedCall                               = 0x17d54f61
	crc_messages_reportEncryptedSpam                     = 0x4b0c8c0f
	crc_payments_getPaymentForm                          = 0x99f09745
	crc_payments_sendPaymentForm                         = 0x2b8879b3
	crc_account_getTmpPassword                           = 0x4a82327e
	crc_messages_setBotShippingResults                   = 0xe5f672fa
	crc_messages_setBotPrecheckoutResults                = 0x09c2dd95
	crc_upload_getWebFile                                = 0x24e6818d
	crc_bots_sendCustomRequest                           = 0xaa2769ed
	crc_bots_answerWebhookJSONQuery                      = 0xe6213f4d
	crc_payments_getPaymentReceipt                       = 0xa092a980
	crc_payments_validateRequestedInfo                   = 0x770a8e74
	crc_payments_getSavedInfo                            = 0x227d824b
	crc_payments_clearSavedInfo                          = 0xd83d70c1
	crc_phone_getCallConfig                              = 0x55451fa9
	crc_phone_confirmCall                                = 0x2efe1722
	crc_phone_setCallRating                              = 0x1c536a34
	crc_phone_saveCallDebug                              = 0x277add7e
	crc_upload_getCdnFile                                = 0x2000bcc3
	crc_upload_reuploadCdnFile                           = 0x1af91c09
	crc_help_getCdnConfig                                = 0x52029342
	crc_messages_uploadMedia                             = 0x519bc2b1
	crc_stickers_createStickerSet                        = 0x9bd86e6a
	crc_langpack_getLangPack                             = 0x9ab5c58e
	crc_langpack_getStrings                              = 0x2e1ee318
	crc_langpack_getDifference                           = 0x0b2e4d7d
	crc_langpack_getLanguages                            = 0x800fd57d
	crc_channels_editBanned                              = 0xbfd915cd
	crc_channels_getAdminLog                             = 0x33ddf480
	crc_stickers_removeStickerFromSet                    = 0xf7760f51
	crc_stickers_changeStickerPosition                   = 0xffb6d4ca
	crc_stickers_addStickerToSet                         = 0x8653febe
	crc_messages_sendScreenshotNotification              = 0xc97df020
	crc_upload_getCdnFileHashes                          = 0xf715c87b
	crc_messages_getUnreadMentions                       = 0x46578472
	crc_messages_faveSticker                             = 0xb9ffc55b
	crc_channels_setStickers                             = 0xea8ca4f9
	crc_contacts_resetSaved                              = 0x879537f1
	crc_messages_getFavedStickers                        = 0x21ce0b0e
	crc_channels_readMessageContents                     = 0xeab5dc38
)

type TL_boolFalse struct {
}

type TL_boolTrue struct {
}

type TL_error struct {
	Code int32
	Text string
}

type TL_null struct {
}

type TL_vector struct {
	T []TL // t
}

type TL_inputPeerEmpty struct {
}

type TL_inputPeerSelf struct {
}

type TL_inputPeerChat struct {
	Chat_id int32
}

type TL_inputUserEmpty struct {
}

type TL_inputUserSelf struct {
}

type TL_inputPhoneContact struct {
	Client_id  int64
	Phone      string
	First_name string
	Last_name  string
}

type TL_inputFile struct {
	Id           int64
	Parts        int32
	Name         string
	Md5_checksum string
}

type TL_inputMediaEmpty struct {
}

type TL_inputMediaUploadedPhoto struct {
	Flags       int32
	File        TL // InputFile
	Caption     string
	Stickers    []TL // InputDocument
	Ttl_seconds int32
}

type TL_inputMediaPhoto struct {
	Flags       int32
	Id          TL // InputPhoto
	Caption     string
	Ttl_seconds int32
}

type TL_inputMediaGeoPoint struct {
	Geo_point TL // InputGeoPoint
}

type TL_inputMediaContact struct {
	Phone_number string
	First_name   string
	Last_name    string
}

type TL_inputChatPhotoEmpty struct {
}

type TL_inputChatUploadedPhoto struct {
	File TL // InputFile
}

type TL_inputChatPhoto struct {
	Id TL // InputPhoto
}

type TL_inputGeoPointEmpty struct {
}

type TL_inputGeoPoint struct {
	Lat  float64
	Long float64
}

type TL_inputPhotoEmpty struct {
}

type TL_inputPhoto struct {
	Id          int64
	Access_hash int64
}

type TL_inputFileLocation struct {
	Volume_id int64
	Local_id  int32
	Secret    int64
}

type TL_inputAppEvent struct {
	Time  float64
	_Type string
	Peer  int64
	Data  string
}

type TL_peerUser struct {
	User_id int32
}

type TL_peerChat struct {
	Chat_id int32
}

type TL_storage_fileUnknown struct {
}

type TL_storage_fileJpeg struct {
}

type TL_storage_fileGif struct {
}

type TL_storage_filePng struct {
}

type TL_storage_fileMp3 struct {
}

type TL_storage_fileMov struct {
}

type TL_storage_filePartial struct {
}

type TL_storage_fileMp4 struct {
}

type TL_storage_fileWebp struct {
}

type TL_fileLocationUnavailable struct {
	Volume_id int64
	Local_id  int32
	Secret    int64
}

type TL_fileLocation struct {
	Dc_id     int32
	Volume_id int64
	Local_id  int32
	Secret    int64
}

type TL_userEmpty struct {
	Id int32
}

type TL_userProfilePhotoEmpty struct {
}

type TL_userProfilePhoto struct {
	Photo_id    int64
	Photo_small TL // FileLocation
	Photo_big   TL // FileLocation
}

type TL_userStatusEmpty struct {
}

type TL_userStatusOnline struct {
	Expires int32
}

type TL_userStatusOffline struct {
	Was_online int32
}

type TL_chatEmpty struct {
	Id int32
}

type TL_chat struct {
	Flags int32
	// Creator	bool // flags_0?true
	// Kicked	bool // flags_1?true
	// Left	bool // flags_2?true
	// Admins_enabled	bool // flags_3?true
	// Admin	bool // flags_4?true
	// Deactivated	bool // flags_5?true
	Id                 int32
	Title              string
	Photo              TL // ChatPhoto
	Participants_count int32
	Date               int32
	Version            int32
	Migrated_to        TL // flags_6?InputChannel
}

type TL_chatForbidden struct {
	Id    int32
	Title string
}

type TL_chatFull struct {
	Id              int32
	Participants    TL   // ChatParticipants
	Chat_photo      TL   // Photo
	Notify_settings TL   // PeerNotifySettings
	Exported_invite TL   // ExportedChatInvite
	Bot_info        []TL // BotInfo
}

type TL_chatParticipant struct {
	User_id    int32
	Inviter_id int32
	Date       int32
}

type TL_chatParticipantsForbidden struct {
	Flags            int32
	Chat_id          int32
	Self_participant TL // flags_0?ChatParticipant
}

type TL_chatParticipants struct {
	Chat_id      int32
	Participants []TL // ChatParticipant
	Version      int32
}

type TL_chatPhotoEmpty struct {
}

type TL_chatPhoto struct {
	Photo_small TL // FileLocation
	Photo_big   TL // FileLocation
}

type TL_messageEmpty struct {
	Id int32
}

type TL_message struct {
	Flags int32
	// Out	bool // flags_1?true
	// Mentioned	bool // flags_4?true
	// Media_unread	bool // flags_5?true
	// Silent	bool // flags_13?true
	// Post	bool // flags_14?true
	Id              int32
	From_id         int32
	To_id           TL // Peer
	Fwd_from        TL // flags_2?MessageFwdHeader
	Via_bot_id      int32
	Reply_to_msg_id int32
	Date            int32
	Message         string
	Media           TL   // flags_9?MessageMedia
	Reply_markup    TL   // flags_6?ReplyMarkup
	Entities        []TL // MessageEntity
	Views           int32
	Edit_date       int32
	Post_author     string
}

type TL_messageService struct {
	Flags int32
	// Out	bool // flags_1?true
	// Mentioned	bool // flags_4?true
	// Media_unread	bool // flags_5?true
	// Silent	bool // flags_13?true
	// Post	bool // flags_14?true
	Id              int32
	From_id         int32
	To_id           TL // Peer
	Reply_to_msg_id int32
	Date            int32
	Action          TL // MessageAction
}

type TL_messageMediaEmpty struct {
}

type TL_messageMediaPhoto struct {
	Flags       int32
	Photo       TL // flags_0?Photo
	Caption     string
	Ttl_seconds int32
}

type TL_messageMediaGeo struct {
	Geo TL // GeoPoint
}

type TL_messageMediaContact struct {
	Phone_number string
	First_name   string
	Last_name    string
	User_id      int32
}

type TL_messageMediaUnsupported struct {
}

type TL_messageActionEmpty struct {
}

type TL_messageActionChatCreate struct {
	Title string
	Users []int32
}

type TL_messageActionChatEditTitle struct {
	Title string
}

type TL_messageActionChatEditPhoto struct {
	Photo TL // Photo
}

type TL_messageActionChatDeletePhoto struct {
}

type TL_messageActionChatAddUser struct {
	Users []int32
}

type TL_messageActionChatDeleteUser struct {
	User_id int32
}

type TL_dialog struct {
	Flags int32
	// Pinned	bool // flags_2?true
	Peer                  TL // Peer
	Top_message           int32
	Read_inbox_max_id     int32
	Read_outbox_max_id    int32
	Unread_count          int32
	Unread_mentions_count int32
	Notify_settings       TL // PeerNotifySettings
	Pts                   int32
	Draft                 TL // flags_1?DraftMessage
}

type TL_photoEmpty struct {
	Id int64
}

type TL_photo struct {
	Flags int32
	// Has_stickers	bool // flags_0?true
	Id          int64
	Access_hash int64
	Date        int32
	Sizes       []TL // PhotoSize
}

type TL_photoSizeEmpty struct {
	_Type string
}

type TL_photoSize struct {
	_Type    string
	Location TL // FileLocation
	W        int32
	H        int32
	Size     int32
}

type TL_photoCachedSize struct {
	_Type    string
	Location TL // FileLocation
	W        int32
	H        int32
	Bytes    []byte
}

type TL_geoPointEmpty struct {
}

type TL_geoPoint struct {
	Long float64
	Lat  float64
}

type TL_auth_checkedPhone struct {
	Phone_registered TL // Bool
}

type TL_auth_sentCode struct {
	Flags int32
	// Phone_registered	bool // flags_0?true
	_Type           TL // auth_SentCodeType
	Phone_code_hash string
	Next_type       TL // flags_1?auth_CodeType
	Timeout         int32
}

type TL_auth_authorization struct {
	Flags        int32
	Tmp_sessions int32
	User         TL // User
}

type TL_auth_exportedAuthorization struct {
	Id    int32
	Bytes []byte
}

type TL_inputNotifyPeer struct {
	Peer TL // InputPeer
}

type TL_inputNotifyUsers struct {
}

type TL_inputNotifyChats struct {
}

type TL_inputNotifyAll struct {
}

type TL_inputPeerNotifySettings struct {
	Flags int32
	// Show_previews	bool // flags_0?true
	// Silent	bool // flags_1?true
	Mute_until int32
	Sound      string
}

type TL_peerNotifyEventsEmpty struct {
}

type TL_peerNotifyEventsAll struct {
}

type TL_peerNotifySettingsEmpty struct {
}

type TL_peerNotifySettings struct {
	Flags int32
	// Show_previews	bool // flags_0?true
	// Silent	bool // flags_1?true
	Mute_until int32
	Sound      string
}

type TL_wallPaper struct {
	Id    int32
	Title string
	Sizes []TL // PhotoSize
	Color int32
}

type TL_userFull struct {
	Flags int32
	// Blocked	bool // flags_0?true
	// Phone_calls_available	bool // flags_4?true
	// Phone_calls_private	bool // flags_5?true
	User               TL // User
	About              string
	Link               TL // contacts_Link
	Profile_photo      TL // flags_2?Photo
	Notify_settings    TL // PeerNotifySettings
	Bot_info           TL // flags_3?BotInfo
	Common_chats_count int32
}

type TL_contact struct {
	User_id int32
	Mutual  TL // Bool
}

type TL_importedContact struct {
	User_id   int32
	Client_id int64
}

type TL_contactBlocked struct {
	User_id int32
	Date    int32
}

type TL_contactStatus struct {
	User_id int32
	Status  TL // UserStatus
}

type TL_contacts_link struct {
	My_link      TL // ContactLink
	Foreign_link TL // ContactLink
	User         TL // User
}

type TL_contacts_contacts struct {
	Contacts    []TL // Contact
	Saved_count int32
	Users       []TL // User
}

type TL_contacts_contactsNotModified struct {
}

type TL_contacts_importedContacts struct {
	Imported        []TL // ImportedContact
	Popular_invites []TL // PopularContact
	Retry_contacts  []int64
	Users           []TL // User
}

type TL_contacts_blocked struct {
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type TL_contacts_blockedSlice struct {
	Count   int32
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type TL_contacts_found struct {
	Results []TL // Peer
	Chats   []TL // Chat
	Users   []TL // User
}

type TL_messages_dialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_dialogsSlice struct {
	Count    int32
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_messages struct {
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_messagesSlice struct {
	Count    int32
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_chats struct {
	Chats []TL // Chat
}

type TL_messages_chatFull struct {
	Full_chat TL   // ChatFull
	Chats     []TL // Chat
	Users     []TL // User
}

type TL_messages_affectedHistory struct {
	Pts       int32
	Pts_count int32
	Offset    int32
}

type TL_inputMessagesFilterEmpty struct {
}

type TL_inputMessagesFilterPhotos struct {
}

type TL_inputMessagesFilterVideo struct {
}

type TL_inputMessagesFilterPhotoVideo struct {
}

type TL_updateNewMessage struct {
	Message   TL // Message
	Pts       int32
	Pts_count int32
}

type TL_updateMessageID struct {
	Id        int32
	Random_id int64
}

type TL_updateDeleteMessages struct {
	Messages  []int32
	Pts       int32
	Pts_count int32
}

type TL_updateUserTyping struct {
	User_id int32
	Action  TL // SendMessageAction
}

type TL_updateChatUserTyping struct {
	Chat_id int32
	User_id int32
	Action  TL // SendMessageAction
}

type TL_updateChatParticipants struct {
	Participants TL // ChatParticipants
}

type TL_updateUserStatus struct {
	User_id int32
	Status  TL // UserStatus
}

type TL_updateUserName struct {
	User_id    int32
	First_name string
	Last_name  string
	Username   string
}

type TL_updateUserPhoto struct {
	User_id  int32
	Date     int32
	Photo    TL // UserProfilePhoto
	Previous TL // Bool
}

type TL_updateContactRegistered struct {
	User_id int32
	Date    int32
}

type TL_updateContactLink struct {
	User_id      int32
	My_link      TL // ContactLink
	Foreign_link TL // ContactLink
}

type TL_updates_state struct {
	Pts          int32
	Qts          int32
	Date         int32
	Seq          int32
	Unread_count int32
}

type TL_updates_differenceEmpty struct {
	Date int32
	Seq  int32
}

type TL_updates_difference struct {
	New_messages           []TL // Message
	New_encrypted_messages []TL // EncryptedMessage
	Other_updates          []TL // Update
	Chats                  []TL // Chat
	Users                  []TL // User
	State                  TL   // updates_State
}

type TL_updates_differenceSlice struct {
	New_messages           []TL // Message
	New_encrypted_messages []TL // EncryptedMessage
	Other_updates          []TL // Update
	Chats                  []TL // Chat
	Users                  []TL // User
	Intermediate_state     TL   // updates_State
}

type TL_updatesTooLong struct {
}

type TL_updateShortMessage struct {
	Flags int32
	// Out	bool // flags_1?true
	// Mentioned	bool // flags_4?true
	// Media_unread	bool // flags_5?true
	// Silent	bool // flags_13?true
	Id              int32
	User_id         int32
	Message         string
	Pts             int32
	Pts_count       int32
	Date            int32
	Fwd_from        TL // flags_2?MessageFwdHeader
	Via_bot_id      int32
	Reply_to_msg_id int32
	Entities        []TL // MessageEntity
}

type TL_updateShortChatMessage struct {
	Flags int32
	// Out	bool // flags_1?true
	// Mentioned	bool // flags_4?true
	// Media_unread	bool // flags_5?true
	// Silent	bool // flags_13?true
	Id              int32
	From_id         int32
	Chat_id         int32
	Message         string
	Pts             int32
	Pts_count       int32
	Date            int32
	Fwd_from        TL // flags_2?MessageFwdHeader
	Via_bot_id      int32
	Reply_to_msg_id int32
	Entities        []TL // MessageEntity
}

type TL_updateShort struct {
	Update TL // Update
	Date   int32
}

type TL_updatesCombined struct {
	Updates   []TL // Update
	Users     []TL // User
	Chats     []TL // Chat
	Date      int32
	Seq_start int32
	Seq       int32
}

type TL_updates struct {
	Updates []TL // Update
	Users   []TL // User
	Chats   []TL // Chat
	Date    int32
	Seq     int32
}

type TL_photos_photo struct {
	Photo TL   // Photo
	Users []TL // User
}

type TL_upload_file struct {
	_Type TL // storage_FileType
	Mtime int32
	Bytes []byte
}

type TL_dcOption struct {
	Flags int32
	// Ipv6	bool // flags_0?true
	// Media_only	bool // flags_1?true
	// Tcpo_only	bool // flags_2?true
	// Cdn	bool // flags_3?true
	// Static	bool // flags_4?true
	Id         int32
	Ip_address string
	Port       int32
}

type TL_config struct {
	Flags int32
	// Phonecalls_enabled	bool // flags_1?true
	Date                     int32
	Expires                  int32
	Test_mode                TL // Bool
	This_dc                  int32
	Dc_options               []TL // DcOption
	Chat_size_max            int32
	Megagroup_size_max       int32
	Forwarded_count_max      int32
	Online_update_period_ms  int32
	Offline_blur_timeout_ms  int32
	Offline_idle_timeout_ms  int32
	Online_cloud_timeout_ms  int32
	Notify_cloud_delay_ms    int32
	Notify_default_delay_ms  int32
	Chat_big_size            int32
	Push_chat_period_ms      int32
	Push_chat_limit          int32
	Saved_gifs_limit         int32
	Edit_time_limit          int32
	Rating_e_decay           int32
	Stickers_recent_limit    int32
	Stickers_faved_limit     int32
	Tmp_sessions             int32
	Pinned_dialogs_count_max int32
	Call_receive_timeout_ms  int32
	Call_ring_timeout_ms     int32
	Call_connect_timeout_ms  int32
	Call_packet_timeout_ms   int32
	Me_url_prefix            string
	Suggested_lang_code      string
	Lang_pack_version        int32
	Disabled_features        []TL // DisabledFeature
}

type TL_nearestDc struct {
	Country    string
	This_dc    int32
	Nearest_dc int32
}

type TL_help_appUpdate struct {
	Id       int32
	Critical TL // Bool
	Url      string
	Text     string
}

type TL_help_noAppUpdate struct {
}

type TL_help_inviteText struct {
	Message string
}

type TL_inputPeerNotifyEventsEmpty struct {
}

type TL_inputPeerNotifyEventsAll struct {
}

type TL_photos_photos struct {
	Photos []TL // Photo
	Users  []TL // User
}

type TL_photos_photosSlice struct {
	Count  int32
	Photos []TL // Photo
	Users  []TL // User
}

type TL_wallPaperSolid struct {
	Id       int32
	Title    string
	Bg_color int32
	Color    int32
}

type TL_updateNewEncryptedMessage struct {
	Message TL // EncryptedMessage
	Qts     int32
}

type TL_updateEncryptedChatTyping struct {
	Chat_id int32
}

type TL_updateEncryption struct {
	Chat TL // EncryptedChat
	Date int32
}

type TL_updateEncryptedMessagesRead struct {
	Chat_id  int32
	Max_date int32
	Date     int32
}

type TL_encryptedChatEmpty struct {
	Id int32
}

type TL_encryptedChatWaiting struct {
	Id             int32
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
}

type TL_encryptedChatRequested struct {
	Id             int32
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
	G_a            []byte
}

type TL_encryptedChat struct {
	Id              int32
	Access_hash     int64
	Date            int32
	Admin_id        int32
	Participant_id  int32
	G_a_or_b        []byte
	Key_fingerprint int64
}

type TL_encryptedChatDiscarded struct {
	Id int32
}

type TL_inputEncryptedChat struct {
	Chat_id     int32
	Access_hash int64
}

type TL_encryptedFileEmpty struct {
}

type TL_encryptedFile struct {
	Id              int64
	Access_hash     int64
	Size            int32
	Dc_id           int32
	Key_fingerprint int32
}

type TL_inputEncryptedFileEmpty struct {
}

type TL_inputEncryptedFileUploaded struct {
	Id              int64
	Parts           int32
	Md5_checksum    string
	Key_fingerprint int32
}

type TL_inputEncryptedFile struct {
	Id          int64
	Access_hash int64
}

type TL_inputEncryptedFileLocation struct {
	Id          int64
	Access_hash int64
}

type TL_encryptedMessage struct {
	Random_id int64
	Chat_id   int32
	Date      int32
	Bytes     []byte
	File      TL // EncryptedFile
}

type TL_encryptedMessageService struct {
	Random_id int64
	Chat_id   int32
	Date      int32
	Bytes     []byte
}

type TL_messages_dhConfigNotModified struct {
	Random []byte
}

type TL_messages_dhConfig struct {
	G       int32
	P       []byte
	Version int32
	Random  []byte
}

type TL_messages_sentEncryptedMessage struct {
	Date int32
}

type TL_messages_sentEncryptedFile struct {
	Date int32
	File TL // EncryptedFile
}

type TL_inputFileBig struct {
	Id    int64
	Parts int32
	Name  string
}

type TL_inputEncryptedFileBigUploaded struct {
	Id              int64
	Parts           int32
	Key_fingerprint int32
}

type TL_storage_filePdf struct {
}

type TL_inputMessagesFilterDocument struct {
}

type TL_inputMessagesFilterPhotoVideoDocuments struct {
}

type TL_updateChatParticipantAdd struct {
	Chat_id    int32
	User_id    int32
	Inviter_id int32
	Date       int32
	Version    int32
}

type TL_updateChatParticipantDelete struct {
	Chat_id int32
	User_id int32
	Version int32
}

type TL_updateDcOptions struct {
	Dc_options []TL // DcOption
}

type TL_inputMediaUploadedDocument struct {
	Flags       int32
	File        TL // InputFile
	Thumb       TL // flags_2?InputFile
	Mime_type   string
	Attributes  []TL // DocumentAttribute
	Caption     string
	Stickers    []TL // InputDocument
	Ttl_seconds int32
}

type TL_inputMediaDocument struct {
	Flags       int32
	Id          TL // InputDocument
	Caption     string
	Ttl_seconds int32
}

type TL_messageMediaDocument struct {
	Flags       int32
	Document    TL // flags_0?Document
	Caption     string
	Ttl_seconds int32
}

type TL_inputDocumentEmpty struct {
}

type TL_inputDocument struct {
	Id          int64
	Access_hash int64
}

type TL_inputDocumentFileLocation struct {
	Id          int64
	Access_hash int64
	Version     int32
}

type TL_documentEmpty struct {
	Id int64
}

type TL_document struct {
	Id          int64
	Access_hash int64
	Date        int32
	Mime_type   string
	Size        int32
	Thumb       TL // PhotoSize
	Dc_id       int32
	Version     int32
	Attributes  []TL // DocumentAttribute
}

type TL_help_support struct {
	Phone_number string
	User         TL // User
}

type TL_notifyAll struct {
}

type TL_notifyChats struct {
}

type TL_notifyPeer struct {
	Peer TL // Peer
}

type TL_notifyUsers struct {
}

type TL_updateUserBlocked struct {
	User_id int32
	Blocked TL // Bool
}

type TL_updateNotifySettings struct {
	Peer            TL // NotifyPeer
	Notify_settings TL // PeerNotifySettings
}

type TL_sendMessageTypingAction struct {
}

type TL_sendMessageCancelAction struct {
}

type TL_sendMessageRecordVideoAction struct {
}

type TL_sendMessageUploadVideoAction struct {
	Progress int32
}

type TL_sendMessageRecordAudioAction struct {
}

type TL_sendMessageUploadAudioAction struct {
	Progress int32
}

type TL_sendMessageUploadPhotoAction struct {
	Progress int32
}

type TL_sendMessageUploadDocumentAction struct {
	Progress int32
}

type TL_sendMessageGeoLocationAction struct {
}

type TL_sendMessageChooseContactAction struct {
}

type TL_updateServiceNotification struct {
	Flags int32
	// Popup	bool // flags_0?true
	Inbox_date int32
	_Type      string
	Message    string
	Media      TL   // MessageMedia
	Entities   []TL // MessageEntity
}

type TL_userStatusRecently struct {
}

type TL_userStatusLastWeek struct {
}

type TL_userStatusLastMonth struct {
}

type TL_updatePrivacy struct {
	Key   TL   // PrivacyKey
	Rules []TL // PrivacyRule
}

type TL_inputPrivacyKeyStatusTimestamp struct {
}

type TL_privacyKeyStatusTimestamp struct {
}

type TL_inputPrivacyValueAllowContacts struct {
}

type TL_inputPrivacyValueAllowAll struct {
}

type TL_inputPrivacyValueAllowUsers struct {
	Users []TL // InputUser
}

type TL_inputPrivacyValueDisallowContacts struct {
}

type TL_inputPrivacyValueDisallowAll struct {
}

type TL_inputPrivacyValueDisallowUsers struct {
	Users []TL // InputUser
}

type TL_privacyValueAllowContacts struct {
}

type TL_privacyValueAllowAll struct {
}

type TL_privacyValueAllowUsers struct {
	Users []int32
}

type TL_privacyValueDisallowContacts struct {
}

type TL_privacyValueDisallowAll struct {
}

type TL_privacyValueDisallowUsers struct {
	Users []int32
}

type TL_account_privacyRules struct {
	Rules []TL // PrivacyRule
	Users []TL // User
}

type TL_accountDaysTTL struct {
	Days int32
}

type TL_updateUserPhone struct {
	User_id int32
	Phone   string
}

type TL_disabledFeature struct {
	Feature     string
	Description string
}

type TL_documentAttributeImageSize struct {
	W int32
	H int32
}

type TL_documentAttributeAnimated struct {
}

type TL_documentAttributeSticker struct {
	Flags int32
	// Mask	bool // flags_1?true
	Alt         string
	Stickerset  TL // InputStickerSet
	Mask_coords TL // flags_0?MaskCoords
}

type TL_documentAttributeVideo struct {
	Flags int32
	// Round_message	bool // flags_0?true
	Duration int32
	W        int32
	H        int32
}

type TL_documentAttributeAudio struct {
	Flags int32
	// Voice	bool // flags_10?true
	Duration  int32
	Title     string
	Performer string
	Waveform  []byte
}

type TL_documentAttributeFilename struct {
	File_name string
}

type TL_messages_stickersNotModified struct {
}

type TL_messages_stickers struct {
	Hash     string
	Stickers []TL // Document
}

type TL_stickerPack struct {
	Emoticon  string
	Documents []int64
}

type TL_messages_allStickersNotModified struct {
}

type TL_messages_allStickers struct {
	Hash int32
	Sets []TL // StickerSet
}

type TL_account_noPassword struct {
	New_salt                  []byte
	Email_unconfirmed_pattern string
}

type TL_account_password struct {
	Current_salt              []byte
	New_salt                  []byte
	Hint                      string
	Has_recovery              TL // Bool
	Email_unconfirmed_pattern string
}

type TL_updateReadHistoryInbox struct {
	Peer      TL // Peer
	Max_id    int32
	Pts       int32
	Pts_count int32
}

type TL_updateReadHistoryOutbox struct {
	Peer      TL // Peer
	Max_id    int32
	Pts       int32
	Pts_count int32
}

type TL_messages_affectedMessages struct {
	Pts       int32
	Pts_count int32
}

type TL_contactLinkUnknown struct {
}

type TL_contactLinkNone struct {
}

type TL_contactLinkHasPhone struct {
}

type TL_contactLinkContact struct {
}

type TL_updateWebPage struct {
	Webpage   TL // WebPage
	Pts       int32
	Pts_count int32
}

type TL_webPageEmpty struct {
	Id int64
}

type TL_webPagePending struct {
	Id   int64
	Date int32
}

type TL_webPage struct {
	Flags        int32
	Id           int64
	Url          string
	Display_url  string
	Hash         int32
	_Type        string
	Site_name    string
	Title        string
	Description  string
	Photo        TL // flags_4?Photo
	Embed_url    string
	Embed_type   string
	Embed_width  int32
	Embed_height int32
	Duration     int32
	Author       string
	Document     TL // flags_9?Document
	Cached_page  TL // flags_10?Page
}

type TL_messageMediaWebPage struct {
	Webpage TL // WebPage
}

type TL_authorization struct {
	Hash           int64
	Flags          int32
	Device_model   string
	Platform       string
	System_version string
	Api_id         int32
	App_name       string
	App_version    string
	Date_created   int32
	Date_active    int32
	Ip             string
	Country        string
	Region         string
}

type TL_account_authorizations struct {
	Authorizations []TL // Authorization
}

type TL_account_passwordSettings struct {
	Email string
}

type TL_account_passwordInputSettings struct {
	Flags             int32
	New_salt          []byte
	New_password_hash []byte
	Hint              string
	Email             string
}

type TL_auth_passwordRecovery struct {
	Email_pattern string
}

type TL_inputMediaVenue struct {
	Geo_point TL // InputGeoPoint
	Title     string
	Address   string
	Provider  string
	Venue_id  string
}

type TL_messageMediaVenue struct {
	Geo      TL // GeoPoint
	Title    string
	Address  string
	Provider string
	Venue_id string
}

type TL_receivedNotifyMessage struct {
	Id    int32
	Flags int32
}

type TL_chatInviteEmpty struct {
}

type TL_chatInviteExported struct {
	Link string
}

type TL_chatInviteAlready struct {
	Chat TL // Chat
}

type TL_chatInvite struct {
	Flags int32
	// Channel	bool // flags_0?true
	// Broadcast	bool // flags_1?true
	// Public	bool // flags_2?true
	// Megagroup	bool // flags_3?true
	Title              string
	Photo              TL // ChatPhoto
	Participants_count int32
	Participants       []TL // User
}

type TL_messageActionChatJoinedByLink struct {
	Inviter_id int32
}

type TL_updateReadMessagesContents struct {
	Messages  []int32
	Pts       int32
	Pts_count int32
}

type TL_inputStickerSetEmpty struct {
}

type TL_inputStickerSetID struct {
	Id          int64
	Access_hash int64
}

type TL_inputStickerSetShortName struct {
	Short_name string
}

type TL_stickerSet struct {
	Flags int32
	// Installed	bool // flags_0?true
	// Archived	bool // flags_1?true
	// Official	bool // flags_2?true
	// Masks	bool // flags_3?true
	Id          int64
	Access_hash int64
	Title       string
	Short_name  string
	Count       int32
	Hash        int32
}

type TL_messages_stickerSet struct {
	Set       TL   // StickerSet
	Packs     []TL // StickerPack
	Documents []TL // Document
}

type TL_user struct {
	Flags int32
	// Self	bool // flags_10?true
	// Contact	bool // flags_11?true
	// Mutual_contact	bool // flags_12?true
	// Deleted	bool // flags_13?true
	// Bot	bool // flags_14?true
	// Bot_chat_history	bool // flags_15?true
	// Bot_nochats	bool // flags_16?true
	// Verified	bool // flags_17?true
	// Restricted	bool // flags_18?true
	// Min	bool // flags_20?true
	// Bot_inline_geo	bool // flags_21?true
	Id                     int32
	Access_hash            int64
	First_name             string
	Last_name              string
	Username               string
	Phone                  string
	Photo                  TL // flags_5?UserProfilePhoto
	Status                 TL // flags_6?UserStatus
	Bot_info_version       int32
	Restriction_reason     string
	Bot_inline_placeholder string
	Lang_code              string
}

type TL_botCommand struct {
	Command     string
	Description string
}

type TL_botInfo struct {
	User_id     int32
	Description string
	Commands    []TL // BotCommand
}

type TL_keyboardButton struct {
	Text string
}

type TL_keyboardButtonRow struct {
	Buttons []TL // KeyboardButton
}

type TL_replyKeyboardHide struct {
	Flags int32
	// Selective	bool // flags_2?true
}

type TL_replyKeyboardForceReply struct {
	Flags int32
	// Single_use	bool // flags_1?true
	// Selective	bool // flags_2?true
}

type TL_replyKeyboardMarkup struct {
	Flags int32
	// Resize	bool // flags_0?true
	// Single_use	bool // flags_1?true
	// Selective	bool // flags_2?true
	Rows []TL // KeyboardButtonRow
}

type TL_inputMessagesFilterUrl struct {
}

type TL_inputPeerUser struct {
	User_id     int32
	Access_hash int64
}

type TL_inputUser struct {
	User_id     int32
	Access_hash int64
}

type TL_messageEntityUnknown struct {
	Offset int32
	Length int32
}

type TL_messageEntityMention struct {
	Offset int32
	Length int32
}

type TL_messageEntityHashtag struct {
	Offset int32
	Length int32
}

type TL_messageEntityBotCommand struct {
	Offset int32
	Length int32
}

type TL_messageEntityUrl struct {
	Offset int32
	Length int32
}

type TL_messageEntityEmail struct {
	Offset int32
	Length int32
}

type TL_messageEntityBold struct {
	Offset int32
	Length int32
}

type TL_messageEntityItalic struct {
	Offset int32
	Length int32
}

type TL_messageEntityCode struct {
	Offset int32
	Length int32
}

type TL_messageEntityPre struct {
	Offset   int32
	Length   int32
	Language string
}

type TL_messageEntityTextUrl struct {
	Offset int32
	Length int32
	Url    string
}

type TL_updateShortSentMessage struct {
	Flags int32
	// Out	bool // flags_1?true
	Id        int32
	Pts       int32
	Pts_count int32
	Date      int32
	Media     TL   // flags_9?MessageMedia
	Entities  []TL // MessageEntity
}

type TL_inputPeerChannel struct {
	Channel_id  int32
	Access_hash int64
}

type TL_peerChannel struct {
	Channel_id int32
}

type TL_channel struct {
	Flags int32
	// Creator	bool // flags_0?true
	// Left	bool // flags_2?true
	// Editor	bool // flags_3?true
	// Broadcast	bool // flags_5?true
	// Verified	bool // flags_7?true
	// Megagroup	bool // flags_8?true
	// Restricted	bool // flags_9?true
	// Democracy	bool // flags_10?true
	// Signatures	bool // flags_11?true
	// Min	bool // flags_12?true
	Id                 int32
	Access_hash        int64
	Title              string
	Username           string
	Photo              TL // ChatPhoto
	Date               int32
	Version            int32
	Restriction_reason string
	Admin_rights       TL // flags_14?ChannelAdminRights
	Banned_rights      TL // flags_15?ChannelBannedRights
}

type TL_channelForbidden struct {
	Flags int32
	// Broadcast	bool // flags_5?true
	// Megagroup	bool // flags_8?true
	Id          int32
	Access_hash int64
	Title       string
	Until_date  int32
}

type TL_channelFull struct {
	Flags int32
	// Can_view_participants	bool // flags_3?true
	// Can_set_username	bool // flags_6?true
	// Can_set_stickers	bool // flags_7?true
	Id                    int32
	About                 string
	Participants_count    int32
	Admins_count          int32
	Kicked_count          int32
	Banned_count          int32
	Read_inbox_max_id     int32
	Read_outbox_max_id    int32
	Unread_count          int32
	Chat_photo            TL   // Photo
	Notify_settings       TL   // PeerNotifySettings
	Exported_invite       TL   // ExportedChatInvite
	Bot_info              []TL // BotInfo
	Migrated_from_chat_id int32
	Migrated_from_max_id  int32
	Pinned_msg_id         int32
	Stickerset            TL // flags_8?StickerSet
}

type TL_messageActionChannelCreate struct {
	Title string
}

type TL_messages_channelMessages struct {
	Flags    int32
	Pts      int32
	Count    int32
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_updateChannelTooLong struct {
	Flags      int32
	Channel_id int32
	Pts        int32
}

type TL_updateChannel struct {
	Channel_id int32
}

type TL_updateNewChannelMessage struct {
	Message   TL // Message
	Pts       int32
	Pts_count int32
}

type TL_updateReadChannelInbox struct {
	Channel_id int32
	Max_id     int32
}

type TL_updateDeleteChannelMessages struct {
	Channel_id int32
	Messages   []int32
	Pts        int32
	Pts_count  int32
}

type TL_updateChannelMessageViews struct {
	Channel_id int32
	Id         int32
	Views      int32
}

type TL_inputChannelEmpty struct {
}

type TL_inputChannel struct {
	Channel_id  int32
	Access_hash int64
}

type TL_contacts_resolvedPeer struct {
	Peer  TL   // Peer
	Chats []TL // Chat
	Users []TL // User
}

type TL_messageRange struct {
	Min_id int32
	Max_id int32
}

type TL_updates_channelDifferenceEmpty struct {
	Flags int32
	// Final	bool // flags_0?true
	Pts     int32
	Timeout int32
}

type TL_updates_channelDifferenceTooLong struct {
	Flags int32
	// Final	bool // flags_0?true
	Pts                   int32
	Timeout               int32
	Top_message           int32
	Read_inbox_max_id     int32
	Read_outbox_max_id    int32
	Unread_count          int32
	Unread_mentions_count int32
	Messages              []TL // Message
	Chats                 []TL // Chat
	Users                 []TL // User
}

type TL_updates_channelDifference struct {
	Flags int32
	// Final	bool // flags_0?true
	Pts           int32
	Timeout       int32
	New_messages  []TL // Message
	Other_updates []TL // Update
	Chats         []TL // Chat
	Users         []TL // User
}

type TL_channelMessagesFilterEmpty struct {
}

type TL_channelMessagesFilter struct {
	Flags int32
	// Exclude_new_messages	bool // flags_1?true
	Ranges []TL // MessageRange
}

type TL_channelParticipant struct {
	User_id int32
	Date    int32
}

type TL_channelParticipantSelf struct {
	User_id    int32
	Inviter_id int32
	Date       int32
}

type TL_channelParticipantCreator struct {
	User_id int32
}

type TL_channelParticipantsRecent struct {
}

type TL_channelParticipantsAdmins struct {
}

type TL_channelParticipantsKicked struct {
	Q string
}

type TL_channels_channelParticipants struct {
	Count        int32
	Participants []TL // ChannelParticipant
	Users        []TL // User
}

type TL_channels_channelParticipant struct {
	Participant TL   // ChannelParticipant
	Users       []TL // User
}

type TL_true struct {
}

type TL_chatParticipantCreator struct {
	User_id int32
}

type TL_chatParticipantAdmin struct {
	User_id    int32
	Inviter_id int32
	Date       int32
}

type TL_updateChatAdmins struct {
	Chat_id int32
	Enabled TL // Bool
	Version int32
}

type TL_updateChatParticipantAdmin struct {
	Chat_id  int32
	User_id  int32
	Is_admin TL // Bool
	Version  int32
}

type TL_messageActionChatMigrateTo struct {
	Channel_id int32
}

type TL_messageActionChannelMigrateFrom struct {
	Title   string
	Chat_id int32
}

type TL_channelParticipantsBots struct {
}

type TL_inputReportReasonSpam struct {
}

type TL_inputReportReasonViolence struct {
}

type TL_inputReportReasonPornography struct {
}

type TL_inputReportReasonOther struct {
	Text string
}

type TL_updateNewStickerSet struct {
	Stickerset TL // messages_StickerSet
}

type TL_updateStickerSetsOrder struct {
	Flags int32
	// Masks	bool // flags_0?true
	Order []int64
}

type TL_updateStickerSets struct {
}

type TL_help_termsOfService struct {
	Text string
}

type TL_foundGif struct {
	Url          string
	Thumb_url    string
	Content_url  string
	Content_type string
	W            int32
	H            int32
}

type TL_inputMediaGifExternal struct {
	Url string
	Q   string
}

type TL_messages_foundGifs struct {
	Next_offset int32
	Results     []TL // FoundGif
}

type TL_inputMessagesFilterGif struct {
}

type TL_updateSavedGifs struct {
}

type TL_updateBotInlineQuery struct {
	Flags    int32
	Query_id int64
	User_id  int32
	Query    string
	Geo      TL // flags_0?GeoPoint
	Offset   string
}

type TL_foundGifCached struct {
	Url      string
	Photo    TL // Photo
	Document TL // Document
}

type TL_messages_savedGifsNotModified struct {
}

type TL_messages_savedGifs struct {
	Hash int32
	Gifs []TL // Document
}

type TL_inputBotInlineMessageMediaAuto struct {
	Flags        int32
	Caption      string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageText struct {
	Flags int32
	// No_webpage	bool // flags_0?true
	Message      string
	Entities     []TL // MessageEntity
	Reply_markup TL   // flags_2?ReplyMarkup
}

type TL_inputBotInlineResult struct {
	Flags        int32
	Id           string
	_Type        string
	Title        string
	Description  string
	Url          string
	Thumb_url    string
	Content_url  string
	Content_type string
	W            int32
	H            int32
	Duration     int32
	Send_message TL // InputBotInlineMessage
}

type TL_botInlineMessageMediaAuto struct {
	Flags        int32
	Caption      string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageText struct {
	Flags int32
	// No_webpage	bool // flags_0?true
	Message      string
	Entities     []TL // MessageEntity
	Reply_markup TL   // flags_2?ReplyMarkup
}

type TL_botInlineResult struct {
	Flags        int32
	Id           string
	_Type        string
	Title        string
	Description  string
	Url          string
	Thumb_url    string
	Content_url  string
	Content_type string
	W            int32
	H            int32
	Duration     int32
	Send_message TL // BotInlineMessage
}

type TL_messages_botResults struct {
	Flags int32
	// Gallery	bool // flags_0?true
	Query_id    int64
	Next_offset string
	Switch_pm   TL   // flags_2?InlineBotSwitchPM
	Results     []TL // BotInlineResult
	Cache_time  int32
}

type TL_inputMessagesFilterVoice struct {
}

type TL_inputMessagesFilterMusic struct {
}

type TL_updateBotInlineSend struct {
	Flags   int32
	User_id int32
	Query   string
	Geo     TL // flags_0?GeoPoint
	Id      string
	Msg_id  TL // flags_1?InputBotInlineMessageID
}

type TL_inputPrivacyKeyChatInvite struct {
}

type TL_privacyKeyChatInvite struct {
}

type TL_updateEditChannelMessage struct {
	Message   TL // Message
	Pts       int32
	Pts_count int32
}

type TL_exportedMessageLink struct {
	Link string
}

type TL_messageFwdHeader struct {
	Flags        int32
	From_id      int32
	Date         int32
	Channel_id   int32
	Channel_post int32
	Post_author  string
}

type TL_messageActionPinMessage struct {
}

type TL_peerSettings struct {
	Flags int32
	// Report_spam	bool // flags_0?true
}

type TL_updateChannelPinnedMessage struct {
	Channel_id int32
	Id         int32
}

type TL_keyboardButtonUrl struct {
	Text string
	Url  string
}

type TL_keyboardButtonCallback struct {
	Text string
	Data []byte
}

type TL_keyboardButtonRequestPhone struct {
	Text string
}

type TL_keyboardButtonRequestGeoLocation struct {
	Text string
}

type TL_auth_codeTypeSms struct {
}

type TL_auth_codeTypeCall struct {
}

type TL_auth_codeTypeFlashCall struct {
}

type TL_auth_sentCodeTypeApp struct {
	Length int32
}

type TL_auth_sentCodeTypeSms struct {
	Length int32
}

type TL_auth_sentCodeTypeCall struct {
	Length int32
}

type TL_auth_sentCodeTypeFlashCall struct {
	Pattern string
}

type TL_keyboardButtonSwitchInline struct {
	Flags int32
	// Same_peer	bool // flags_0?true
	Text  string
	Query string
}

type TL_replyInlineMarkup struct {
	Rows []TL // KeyboardButtonRow
}

type TL_messages_botCallbackAnswer struct {
	Flags int32
	// Alert	bool // flags_1?true
	// Has_url	bool // flags_3?true
	Message    string
	Url        string
	Cache_time int32
}

type TL_updateBotCallbackQuery struct {
	Flags           int32
	Query_id        int64
	User_id         int32
	Peer            TL // Peer
	Msg_id          int32
	Chat_instance   int64
	Data            []byte
	Game_short_name string
}

type TL_messages_messageEditData struct {
	Flags int32
	// Caption	bool // flags_0?true
}

type TL_updateEditMessage struct {
	Message   TL // Message
	Pts       int32
	Pts_count int32
}

type TL_inputBotInlineMessageMediaGeo struct {
	Flags        int32
	Geo_point    TL // InputGeoPoint
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageMediaVenue struct {
	Flags        int32
	Geo_point    TL // InputGeoPoint
	Title        string
	Address      string
	Provider     string
	Venue_id     string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineMessageMediaContact struct {
	Flags        int32
	Phone_number string
	First_name   string
	Last_name    string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaGeo struct {
	Flags        int32
	Geo          TL // GeoPoint
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaVenue struct {
	Flags        int32
	Geo          TL // GeoPoint
	Title        string
	Address      string
	Provider     string
	Venue_id     string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_botInlineMessageMediaContact struct {
	Flags        int32
	Phone_number string
	First_name   string
	Last_name    string
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineResultPhoto struct {
	Id           string
	_Type        string
	Photo        TL // InputPhoto
	Send_message TL // InputBotInlineMessage
}

type TL_inputBotInlineResultDocument struct {
	Flags        int32
	Id           string
	_Type        string
	Title        string
	Description  string
	Document     TL // InputDocument
	Send_message TL // InputBotInlineMessage
}

type TL_botInlineMediaResult struct {
	Flags        int32
	Id           string
	_Type        string
	Photo        TL // flags_0?Photo
	Document     TL // flags_1?Document
	Title        string
	Description  string
	Send_message TL // BotInlineMessage
}

type TL_inputBotInlineMessageID struct {
	Dc_id       int32
	Id          int64
	Access_hash int64
}

type TL_updateInlineBotCallbackQuery struct {
	Flags           int32
	Query_id        int64
	User_id         int32
	Msg_id          TL // InputBotInlineMessageID
	Chat_instance   int64
	Data            []byte
	Game_short_name string
}

type TL_inlineBotSwitchPM struct {
	Text        string
	Start_param string
}

type TL_messageEntityMentionName struct {
	Offset  int32
	Length  int32
	User_id int32
}

type TL_inputMessageEntityMentionName struct {
	Offset  int32
	Length  int32
	User_id TL // InputUser
}

type TL_messages_peerDialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
	State    TL   // updates_State
}

type TL_topPeer struct {
	Peer   TL // Peer
	Rating float64
}

type TL_topPeerCategoryBotsPM struct {
}

type TL_topPeerCategoryBotsInline struct {
}

type TL_topPeerCategoryCorrespondents struct {
}

type TL_topPeerCategoryGroups struct {
}

type TL_topPeerCategoryChannels struct {
}

type TL_topPeerCategoryPeers struct {
	Category TL // TopPeerCategory
	Count    int32
	Peers    []TL // TopPeer
}

type TL_contacts_topPeersNotModified struct {
}

type TL_contacts_topPeers struct {
	Categories []TL // TopPeerCategoryPeers
	Chats      []TL // Chat
	Users      []TL // User
}

type TL_inputMessagesFilterChatPhotos struct {
}

type TL_updateReadChannelOutbox struct {
	Channel_id int32
	Max_id     int32
}

type TL_updateDraftMessage struct {
	Peer  TL // Peer
	Draft TL // DraftMessage
}

type TL_draftMessageEmpty struct {
}

type TL_draftMessage struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	Reply_to_msg_id int32
	Message         string
	Entities        []TL // MessageEntity
	Date            int32
}

type TL_messageActionHistoryClear struct {
}

type TL_updateReadFeaturedStickers struct {
}

type TL_updateRecentStickers struct {
}

type TL_messages_featuredStickersNotModified struct {
}

type TL_messages_featuredStickers struct {
	Hash   int32
	Sets   []TL // StickerSetCovered
	Unread []int64
}

type TL_messages_recentStickersNotModified struct {
}

type TL_messages_recentStickers struct {
	Hash     int32
	Stickers []TL // Document
}

type TL_messages_archivedStickers struct {
	Count int32
	Sets  []TL // StickerSetCovered
}

type TL_messages_stickerSetInstallResultSuccess struct {
}

type TL_messages_stickerSetInstallResultArchive struct {
	Sets []TL // StickerSetCovered
}

type TL_stickerSetCovered struct {
	Set   TL // StickerSet
	Cover TL // Document
}

type TL_inputMediaPhotoExternal struct {
	Flags       int32
	Url         string
	Caption     string
	Ttl_seconds int32
}

type TL_inputMediaDocumentExternal struct {
	Flags       int32
	Url         string
	Caption     string
	Ttl_seconds int32
}

type TL_updateConfig struct {
}

type TL_updatePtsChanged struct {
}

type TL_messageActionGameScore struct {
	Game_id int64
	Score   int32
}

type TL_documentAttributeHasStickers struct {
}

type TL_keyboardButtonGame struct {
	Text string
}

type TL_stickerSetMultiCovered struct {
	Set    TL   // StickerSet
	Covers []TL // Document
}

type TL_maskCoords struct {
	N    int32
	X    float64
	Y    float64
	Zoom float64
}

type TL_inputStickeredMediaPhoto struct {
	Id TL // InputPhoto
}

type TL_inputStickeredMediaDocument struct {
	Id TL // InputDocument
}

type TL_inputMediaGame struct {
	Id TL // InputGame
}

type TL_messageMediaGame struct {
	Game TL // Game
}

type TL_inputBotInlineMessageGame struct {
	Flags        int32
	Reply_markup TL // flags_2?ReplyMarkup
}

type TL_inputBotInlineResultGame struct {
	Id           string
	Short_name   string
	Send_message TL // InputBotInlineMessage
}

type TL_game struct {
	Flags       int32
	Id          int64
	Access_hash int64
	Short_name  string
	Title       string
	Description string
	Photo       TL // Photo
	Document    TL // flags_0?Document
}

type TL_inputGameID struct {
	Id          int64
	Access_hash int64
}

type TL_inputGameShortName struct {
	Bot_id     TL // InputUser
	Short_name string
}

type TL_highScore struct {
	Pos     int32
	User_id int32
	Score   int32
}

type TL_messages_highScores struct {
	Scores []TL // HighScore
	Users  []TL // User
}

type TL_messages_chatsSlice struct {
	Count int32
	Chats []TL // Chat
}

type TL_updateChannelWebPage struct {
	Channel_id int32
	Webpage    TL // WebPage
	Pts        int32
	Pts_count  int32
}

type TL_updates_differenceTooLong struct {
	Pts int32
}

type TL_sendMessageGamePlayAction struct {
}

type TL_webPageNotModified struct {
}

type TL_textEmpty struct {
}

type TL_textPlain struct {
	Text string
}

type TL_textBold struct {
	Text TL // RichText
}

type TL_textItalic struct {
	Text TL // RichText
}

type TL_textUnderline struct {
	Text TL // RichText
}

type TL_textStrike struct {
	Text TL // RichText
}

type TL_textFixed struct {
	Text TL // RichText
}

type TL_textUrl struct {
	Text       TL // RichText
	Url        string
	Webpage_id int64
}

type TL_textEmail struct {
	Text  TL // RichText
	Email string
}

type TL_textConcat struct {
	Texts []TL // RichText
}

type TL_pageBlockTitle struct {
	Text TL // RichText
}

type TL_pageBlockSubtitle struct {
	Text TL // RichText
}

type TL_pageBlockAuthorDate struct {
	Author         TL // RichText
	Published_date int32
}

type TL_pageBlockHeader struct {
	Text TL // RichText
}

type TL_pageBlockSubheader struct {
	Text TL // RichText
}

type TL_pageBlockParagraph struct {
	Text TL // RichText
}

type TL_pageBlockPreformatted struct {
	Text     TL // RichText
	Language string
}

type TL_pageBlockFooter struct {
	Text TL // RichText
}

type TL_pageBlockDivider struct {
}

type TL_pageBlockList struct {
	Ordered TL   // Bool
	Items   []TL // RichText
}

type TL_pageBlockBlockquote struct {
	Text    TL // RichText
	Caption TL // RichText
}

type TL_pageBlockPullquote struct {
	Text    TL // RichText
	Caption TL // RichText
}

type TL_pageBlockPhoto struct {
	Photo_id int64
	Caption  TL // RichText
}

type TL_pageBlockVideo struct {
	Flags int32
	// Autoplay	bool // flags_0?true
	// Loop	bool // flags_1?true
	Video_id int64
	Caption  TL // RichText
}

type TL_pageBlockCover struct {
	Cover TL // PageBlock
}

type TL_pageBlockEmbed struct {
	Flags int32
	// Full_width	bool // flags_0?true
	// Allow_scrolling	bool // flags_3?true
	Url             string
	Html            string
	Poster_photo_id int64
	W               int32
	H               int32
	Caption         TL // RichText
}

type TL_pageBlockEmbedPost struct {
	Url             string
	Webpage_id      int64
	Author_photo_id int64
	Author          string
	Date            int32
	Blocks          []TL // PageBlock
	Caption         TL   // RichText
}

type TL_pageBlockSlideshow struct {
	Items   []TL // PageBlock
	Caption TL   // RichText
}

type TL_pagePart struct {
	Blocks    []TL // PageBlock
	Photos    []TL // Photo
	Documents []TL // Document
}

type TL_pageFull struct {
	Blocks    []TL // PageBlock
	Photos    []TL // Photo
	Documents []TL // Document
}

type TL_updatePhoneCall struct {
	Phone_call TL // PhoneCall
}

type TL_updateDialogPinned struct {
	Flags int32
	// Pinned	bool // flags_0?true
	Peer TL // Peer
}

type TL_updatePinnedDialogs struct {
	Flags int32
	Order []TL // Peer
}

type TL_inputPrivacyKeyPhoneCall struct {
}

type TL_privacyKeyPhoneCall struct {
}

type TL_pageBlockUnsupported struct {
}

type TL_pageBlockAnchor struct {
	Name string
}

type TL_pageBlockCollage struct {
	Items   []TL // PageBlock
	Caption TL   // RichText
}

type TL_inputPhoneCall struct {
	Id          int64
	Access_hash int64
}

type TL_phoneCallEmpty struct {
	Id int64
}

type TL_phoneCallWaiting struct {
	Flags          int32
	Id             int64
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
	Protocol       TL // PhoneCallProtocol
	Receive_date   int32
}

type TL_phoneCallRequested struct {
	Id             int64
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
	G_a_hash       []byte
	Protocol       TL // PhoneCallProtocol
}

type TL_phoneCall struct {
	Id                      int64
	Access_hash             int64
	Date                    int32
	Admin_id                int32
	Participant_id          int32
	G_a_or_b                []byte
	Key_fingerprint         int64
	Protocol                TL   // PhoneCallProtocol
	Connection              TL   // PhoneConnection
	Alternative_connections []TL // PhoneConnection
	Start_date              int32
}

type TL_phoneCallDiscarded struct {
	Flags int32
	// Need_rating	bool // flags_2?true
	// Need_debug	bool // flags_3?true
	Id       int64
	Reason   TL // flags_0?PhoneCallDiscardReason
	Duration int32
}

type TL_phoneConnection struct {
	Id       int64
	Ip       string
	Ipv6     string
	Port     int32
	Peer_tag []byte
}

type TL_phoneCallProtocol struct {
	Flags int32
	// Udp_p2p	bool // flags_0?true
	// Udp_reflector	bool // flags_1?true
	Min_layer int32
	Max_layer int32
}

type TL_phone_phoneCall struct {
	Phone_call TL   // PhoneCall
	Users      []TL // User
}

type TL_phoneCallDiscardReasonMissed struct {
}

type TL_phoneCallDiscardReasonDisconnect struct {
}

type TL_phoneCallDiscardReasonHangup struct {
}

type TL_phoneCallDiscardReasonBusy struct {
}

type TL_inputMessagesFilterPhoneCalls struct {
	Flags int32
	// Missed	bool // flags_0?true
}

type TL_messageActionPhoneCall struct {
	Flags    int32
	Call_id  int64
	Reason   TL // flags_0?PhoneCallDiscardReason
	Duration int32
}

type TL_invoice struct {
	Flags int32
	// Test	bool // flags_0?true
	// Name_requested	bool // flags_1?true
	// Phone_requested	bool // flags_2?true
	// Email_requested	bool // flags_3?true
	// Shipping_address_requested	bool // flags_4?true
	// Flexible	bool // flags_5?true
	Currency string
	Prices   []TL // LabeledPrice
}

type TL_inputMediaInvoice struct {
	Flags       int32
	Title       string
	Description string
	Photo       TL // flags_0?InputWebDocument
	Invoice     TL // Invoice
	Payload     []byte
	Provider    string
	Start_param string
}

type TL_messageActionPaymentSentMe struct {
	Flags              int32
	Currency           string
	Total_amount       int64
	Payload            []byte
	Info               TL // flags_0?PaymentRequestedInfo
	Shipping_option_id string
	Charge             TL // PaymentCharge
}

type TL_messageMediaInvoice struct {
	Flags int32
	// Shipping_address_requested	bool // flags_1?true
	// Test	bool // flags_3?true
	Title          string
	Description    string
	Photo          TL // flags_0?WebDocument
	Receipt_msg_id int32
	Currency       string
	Total_amount   int64
	Start_param    string
}

type TL_keyboardButtonBuy struct {
	Text string
}

type TL_messageActionPaymentSent struct {
	Currency     string
	Total_amount int64
}

type TL_payments_paymentForm struct {
	Flags int32
	// Can_save_credentials	bool // flags_2?true
	// Password_missing	bool // flags_3?true
	Bot_id            int32
	Invoice           TL // Invoice
	Provider_id       int32
	Url               string
	Native_provider   string
	Native_params     TL   // flags_4?DataJSON
	Saved_info        TL   // flags_0?PaymentRequestedInfo
	Saved_credentials TL   // flags_1?PaymentSavedCredentials
	Users             []TL // User
}

type TL_postAddress struct {
	Street_line1 string
	Street_line2 string
	City         string
	State        string
	Country_iso2 string
	Post_code    string
}

type TL_paymentRequestedInfo struct {
	Flags            int32
	Name             string
	Phone            string
	Email            string
	Shipping_address TL // flags_3?PostAddress
}

type TL_updateBotWebhookJSON struct {
	Data TL // DataJSON
}

type TL_updateBotWebhookJSONQuery struct {
	Query_id int64
	Data     TL // DataJSON
	Timeout  int32
}

type TL_updateBotShippingQuery struct {
	Query_id         int64
	User_id          int32
	Payload          []byte
	Shipping_address TL // PostAddress
}

type TL_updateBotPrecheckoutQuery struct {
	Flags              int32
	Query_id           int64
	User_id            int32
	Payload            []byte
	Info               TL // flags_0?PaymentRequestedInfo
	Shipping_option_id string
	Currency           string
	Total_amount       int64
}

type TL_dataJSON struct {
	Data string
}

type TL_labeledPrice struct {
	Label  string
	Amount int64
}

type TL_paymentCharge struct {
	Id                 string
	Provider_charge_id string
}

type TL_paymentSavedCredentialsCard struct {
	Id    string
	Title string
}

type TL_webDocument struct {
	Url         string
	Access_hash int64
	Size        int32
	Mime_type   string
	Attributes  []TL // DocumentAttribute
	Dc_id       int32
}

type TL_inputWebDocument struct {
	Url        string
	Size       int32
	Mime_type  string
	Attributes []TL // DocumentAttribute
}

type TL_inputWebFileLocation struct {
	Url         string
	Access_hash int64
}

type TL_upload_webFile struct {
	Size      int32
	Mime_type string
	File_type TL // storage_FileType
	Mtime     int32
	Bytes     []byte
}

type TL_payments_validatedRequestedInfo struct {
	Flags            int32
	Id               string
	Shipping_options []TL // ShippingOption
}

type TL_payments_paymentResult struct {
	Updates TL // Updates
}

type TL_payments_paymentVerficationNeeded struct {
	Url string
}

type TL_payments_paymentReceipt struct {
	Flags             int32
	Date              int32
	Bot_id            int32
	Invoice           TL // Invoice
	Provider_id       int32
	Info              TL // flags_0?PaymentRequestedInfo
	Shipping          TL // flags_1?ShippingOption
	Currency          string
	Total_amount      int64
	Credentials_title string
	Users             []TL // User
}

type TL_payments_savedInfo struct {
	Flags int32
	// Has_saved_credentials	bool // flags_1?true
	Saved_info TL // flags_0?PaymentRequestedInfo
}

type TL_inputPaymentCredentialsSaved struct {
	Id           string
	Tmp_password []byte
}

type TL_inputPaymentCredentials struct {
	Flags int32
	// Save	bool // flags_0?true
	Data TL // DataJSON
}

type TL_account_tmpPassword struct {
	Tmp_password []byte
	Valid_until  int32
}

type TL_shippingOption struct {
	Id     string
	Title  string
	Prices []TL // LabeledPrice
}

type TL_phoneCallAccepted struct {
	Id             int64
	Access_hash    int64
	Date           int32
	Admin_id       int32
	Participant_id int32
	G_b            []byte
	Protocol       TL // PhoneCallProtocol
}

type TL_inputMessagesFilterRoundVoice struct {
}

type TL_inputMessagesFilterRoundVideo struct {
}

type TL_upload_fileCdnRedirect struct {
	Dc_id           int32
	File_token      []byte
	Encryption_key  []byte
	Encryption_iv   []byte
	Cdn_file_hashes []TL // CdnFileHash
}

type TL_sendMessageRecordRoundAction struct {
}

type TL_sendMessageUploadRoundAction struct {
	Progress int32
}

type TL_upload_cdnFileReuploadNeeded struct {
	Request_token []byte
}

type TL_upload_cdnFile struct {
	Bytes []byte
}

type TL_cdnPublicKey struct {
	Dc_id      int32
	Public_key string
}

type TL_cdnConfig struct {
	Public_keys []TL // CdnPublicKey
}

type TL_updateLangPackTooLong struct {
}

type TL_updateLangPack struct {
	Difference TL // LangPackDifference
}

type TL_pageBlockChannel struct {
	Channel TL // Chat
}

type TL_inputStickerSetItem struct {
	Flags       int32
	Document    TL // InputDocument
	Emoji       string
	Mask_coords TL // flags_0?MaskCoords
}

type TL_langPackString struct {
	Key   string
	Value string
}

type TL_langPackStringPluralized struct {
	Flags       int32
	Key         string
	Zero_value  string
	One_value   string
	Two_value   string
	Few_value   string
	Many_value  string
	Other_value string
}

type TL_langPackStringDeleted struct {
	Key string
}

type TL_langPackDifference struct {
	Lang_code    string
	From_version int32
	Version      int32
	Strings      []TL // LangPackString
}

type TL_langPackLanguage struct {
	Name        string
	Native_name string
	Lang_code   string
}

type TL_channelParticipantAdmin struct {
	Flags int32
	// Can_edit	bool // flags_0?true
	User_id      int32
	Inviter_id   int32
	Promoted_by  int32
	Date         int32
	Admin_rights TL // ChannelAdminRights
}

type TL_channelParticipantBanned struct {
	Flags int32
	// Left	bool // flags_0?true
	User_id       int32
	Kicked_by     int32
	Date          int32
	Banned_rights TL // ChannelBannedRights
}

type TL_channelParticipantsBanned struct {
	Q string
}

type TL_channelParticipantsSearch struct {
	Q string
}

type TL_topPeerCategoryPhoneCalls struct {
}

type TL_pageBlockAudio struct {
	Audio_id int64
	Caption  TL // RichText
}

type TL_channelAdminRights struct {
	Flags int32
	// Change_info	bool // flags_0?true
	// Post_messages	bool // flags_1?true
	// Edit_messages	bool // flags_2?true
	// Delete_messages	bool // flags_3?true
	// Ban_users	bool // flags_4?true
	// Invite_users	bool // flags_5?true
	// Invite_link	bool // flags_6?true
	// Pin_messages	bool // flags_7?true
	// Add_admins	bool // flags_9?true
}

type TL_channelBannedRights struct {
	Flags int32
	// View_messages	bool // flags_0?true
	// Send_messages	bool // flags_1?true
	// Send_media	bool // flags_2?true
	// Send_stickers	bool // flags_3?true
	// Send_gifs	bool // flags_4?true
	// Send_games	bool // flags_5?true
	// Send_inline	bool // flags_6?true
	// Embed_links	bool // flags_7?true
	Until_date int32
}

type TL_channelAdminLogEventActionChangeTitle struct {
	Prev_value string
	New_value  string
}

type TL_channelAdminLogEventActionChangeAbout struct {
	Prev_value string
	New_value  string
}

type TL_channelAdminLogEventActionChangeUsername struct {
	Prev_value string
	New_value  string
}

type TL_channelAdminLogEventActionChangePhoto struct {
	Prev_photo TL // ChatPhoto
	New_photo  TL // ChatPhoto
}

type TL_channelAdminLogEventActionToggleInvites struct {
	New_value TL // Bool
}

type TL_channelAdminLogEventActionToggleSignatures struct {
	New_value TL // Bool
}

type TL_channelAdminLogEventActionUpdatePinned struct {
	Message TL // Message
}

type TL_channelAdminLogEventActionEditMessage struct {
	Prev_message TL // Message
	New_message  TL // Message
}

type TL_channelAdminLogEventActionDeleteMessage struct {
	Message TL // Message
}

type TL_channelAdminLogEventActionParticipantJoin struct {
}

type TL_channelAdminLogEventActionParticipantLeave struct {
}

type TL_channelAdminLogEventActionParticipantInvite struct {
	Participant TL // ChannelParticipant
}

type TL_channelAdminLogEventActionParticipantToggleBan struct {
	Prev_participant TL // ChannelParticipant
	New_participant  TL // ChannelParticipant
}

type TL_channelAdminLogEventActionParticipantToggleAdmin struct {
	Prev_participant TL // ChannelParticipant
	New_participant  TL // ChannelParticipant
}

type TL_channelAdminLogEvent struct {
	Id      int64
	Date    int32
	User_id int32
	Action  TL // ChannelAdminLogEventAction
}

type TL_channels_adminLogResults struct {
	Events []TL // ChannelAdminLogEvent
	Chats  []TL // Chat
	Users  []TL // User
}

type TL_channelAdminLogEventsFilter struct {
	Flags int32
	// Join	bool // flags_0?true
	// Leave	bool // flags_1?true
	// Invite	bool // flags_2?true
	// Ban	bool // flags_3?true
	// Unban	bool // flags_4?true
	// Kick	bool // flags_5?true
	// Unkick	bool // flags_6?true
	// Promote	bool // flags_7?true
	// Demote	bool // flags_8?true
	// Info	bool // flags_9?true
	// Settings	bool // flags_10?true
	// Pinned	bool // flags_11?true
	// Edit	bool // flags_12?true
	// Delete	bool // flags_13?true
}

type TL_messageActionScreenshotTaken struct {
}

type TL_popularContact struct {
	Client_id int64
	Importers int32
}

type TL_cdnFileHash struct {
	Offset int32
	Limit  int32
	Hash   []byte
}

type TL_inputMessagesFilterMyMentions struct {
}

type TL_inputMessagesFilterMyMentionsUnread struct {
}

type TL_updateContactsReset struct {
}

type TL_channelAdminLogEventActionChangeStickerSet struct {
	Prev_stickerset TL // InputStickerSet
	New_stickerset  TL // InputStickerSet
}

type TL_updateFavedStickers struct {
}

type TL_messages_favedStickers struct {
	Hash     int32
	Packs    []TL // StickerPack
	Stickers []TL // Document
}

type TL_messages_favedStickersNotModified struct {
}

type TL_updateChannelReadMessagesContents struct {
	Channel_id int32
	Messages   []int32
}

type TL_invokeAfterMsg struct {
	Msg_id int64
	Query  TL
}

type TL_invokeAfterMsgs struct {
	Msg_ids []int64
	Query   TL
}

type TL_auth_checkPhone struct {
	Phone_number string
}

type TL_auth_sendCode struct {
	Flags int32
	// Allow_flashcall	bool // flags_0?true
	Phone_number   string
	Current_number TL // flags_0?Bool
	Api_id         int32
	Api_hash       string
}

type TL_auth_signUp struct {
	Phone_number    string
	Phone_code_hash string
	Phone_code      string
	First_name      string
	Last_name       string
}

type TL_auth_signIn struct {
	Phone_number    string
	Phone_code_hash string
	Phone_code      string
}

type TL_auth_logOut struct {
}

type TL_auth_resetAuthorizations struct {
}

type TL_auth_sendInvites struct {
	Phone_numbers []string
	Message       string
}

type TL_auth_exportAuthorization struct {
	Dc_id int32
}

type TL_auth_importAuthorization struct {
	Id    int32
	Bytes []byte
}

type TL_account_registerDevice struct {
	Token_type int32
	Token      string
}

type TL_account_unregisterDevice struct {
	Token_type int32
	Token      string
}

type TL_account_updateNotifySettings struct {
	Peer     TL // InputNotifyPeer
	Settings TL // InputPeerNotifySettings
}

type TL_account_getNotifySettings struct {
	Peer TL // InputNotifyPeer
}

type TL_account_resetNotifySettings struct {
}

type TL_account_updateProfile struct {
	Flags      int32
	First_name string
	Last_name  string
	About      string
}

type TL_account_updateStatus struct {
	Offline TL // Bool
}

type TL_account_getWallPapers struct {
}

type TL_users_getUsers struct {
	Id []TL // InputUser
}

type TL_users_getFullUser struct {
	Id TL // InputUser
}

type TL_contacts_getStatuses struct {
}

type TL_contacts_getContacts struct {
	Hash int32
}

type TL_contacts_importContacts struct {
	Contacts []TL // InputContact
}

type TL_contacts_search struct {
	Q     string
	Limit int32
}

type TL_contacts_deleteContact struct {
	Id TL // InputUser
}

type TL_contacts_deleteContacts struct {
	Id []TL // InputUser
}

type TL_contacts_block struct {
	Id TL // InputUser
}

type TL_contacts_unblock struct {
	Id TL // InputUser
}

type TL_contacts_getBlocked struct {
	Offset int32
	Limit  int32
}

type TL_messages_getMessages struct {
	Id []int32
}

type TL_messages_getDialogs struct {
	Flags int32
	// Exclude_pinned	bool // flags_0?true
	Offset_date int32
	Offset_id   int32
	Offset_peer TL // InputPeer
	Limit       int32
}

type TL_messages_getHistory struct {
	Peer        TL // InputPeer
	Offset_id   int32
	Offset_date int32
	Add_offset  int32
	Limit       int32
	Max_id      int32
	Min_id      int32
}

type TL_messages_search struct {
	Flags      int32
	Peer       TL // InputPeer
	Q          string
	From_id    TL // flags_0?InputUser
	Filter     TL // MessagesFilter
	Min_date   int32
	Max_date   int32
	Offset_id  int32
	Add_offset int32
	Limit      int32
	Max_id     int32
	Min_id     int32
}

type TL_messages_readHistory struct {
	Peer   TL // InputPeer
	Max_id int32
}

type TL_messages_deleteHistory struct {
	Flags int32
	// Just_clear	bool // flags_0?true
	Peer   TL // InputPeer
	Max_id int32
}

type TL_messages_deleteMessages struct {
	Flags int32
	// Revoke	bool // flags_0?true
	Id []int32
}

type TL_messages_receivedMessages struct {
	Max_id int32
}

type TL_messages_setTyping struct {
	Peer   TL // InputPeer
	Action TL // SendMessageAction
}

type TL_messages_sendMessage struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	// Silent	bool // flags_5?true
	// Background	bool // flags_6?true
	// Clear_draft	bool // flags_7?true
	Peer            TL // InputPeer
	Reply_to_msg_id int32
	Message         string
	Random_id       int64
	Reply_markup    TL   // flags_2?ReplyMarkup
	Entities        []TL // MessageEntity
}

type TL_messages_sendMedia struct {
	Flags int32
	// Silent	bool // flags_5?true
	// Background	bool // flags_6?true
	// Clear_draft	bool // flags_7?true
	Peer            TL // InputPeer
	Reply_to_msg_id int32
	Media           TL // InputMedia
	Random_id       int64
	Reply_markup    TL // flags_2?ReplyMarkup
}

type TL_messages_forwardMessages struct {
	Flags int32
	// Silent	bool // flags_5?true
	// Background	bool // flags_6?true
	// With_my_score	bool // flags_8?true
	From_peer TL // InputPeer
	Id        []int32
	Random_id []int64
	To_peer   TL // InputPeer
}

type TL_messages_getChats struct {
	Id []int32
}

type TL_messages_getFullChat struct {
	Chat_id int32
}

type TL_messages_editChatTitle struct {
	Chat_id int32
	Title   string
}

type TL_messages_editChatPhoto struct {
	Chat_id int32
	Photo   TL // InputChatPhoto
}

type TL_messages_addChatUser struct {
	Chat_id   int32
	User_id   TL // InputUser
	Fwd_limit int32
}

type TL_messages_deleteChatUser struct {
	Chat_id int32
	User_id TL // InputUser
}

type TL_messages_createChat struct {
	Users []TL // InputUser
	Title string
}

type TL_updates_getState struct {
}

type TL_updates_getDifference struct {
	Flags           int32
	Pts             int32
	Pts_total_limit int32
	Date            int32
	Qts             int32
}

type TL_photos_updateProfilePhoto struct {
	Id TL // InputPhoto
}

type TL_photos_uploadProfilePhoto struct {
	File TL // InputFile
}

type TL_upload_saveFilePart struct {
	File_id   int64
	File_part int32
	Bytes     []byte
}

type TL_upload_getFile struct {
	Location TL // InputFileLocation
	Offset   int32
	Limit    int32
}

type TL_help_getConfig struct {
}

type TL_help_getNearestDc struct {
}

type TL_help_getAppUpdate struct {
}

type TL_help_saveAppLog struct {
	Events []TL // InputAppEvent
}

type TL_help_getInviteText struct {
}

type TL_photos_deletePhotos struct {
	Id []TL // InputPhoto
}

type TL_photos_getUserPhotos struct {
	User_id TL // InputUser
	Offset  int32
	Max_id  int64
	Limit   int32
}

type TL_messages_forwardMessage struct {
	Peer      TL // InputPeer
	Id        int32
	Random_id int64
}

type TL_messages_getDhConfig struct {
	Version       int32
	Random_length int32
}

type TL_messages_requestEncryption struct {
	User_id   TL // InputUser
	Random_id int32
	G_a       []byte
}

type TL_messages_acceptEncryption struct {
	Peer            TL // InputEncryptedChat
	G_b             []byte
	Key_fingerprint int64
}

type TL_messages_discardEncryption struct {
	Chat_id int32
}

type TL_messages_setEncryptedTyping struct {
	Peer   TL // InputEncryptedChat
	Typing TL // Bool
}

type TL_messages_readEncryptedHistory struct {
	Peer     TL // InputEncryptedChat
	Max_date int32
}

type TL_messages_sendEncrypted struct {
	Peer      TL // InputEncryptedChat
	Random_id int64
	Data      []byte
}

type TL_messages_sendEncryptedFile struct {
	Peer      TL // InputEncryptedChat
	Random_id int64
	Data      []byte
	File      TL // InputEncryptedFile
}

type TL_messages_sendEncryptedService struct {
	Peer      TL // InputEncryptedChat
	Random_id int64
	Data      []byte
}

type TL_messages_receivedQueue struct {
	Max_qts int32
}

type TL_upload_saveBigFilePart struct {
	File_id          int64
	File_part        int32
	File_total_parts int32
	Bytes            []byte
}

type TL_initConnection struct {
	Api_id           int32
	Device_model     string
	System_version   string
	App_version      string
	System_lang_code string
	Lang_pack        string
	Lang_code        string
	Query            TL
}

type TL_help_getSupport struct {
}

type TL_auth_bindTempAuthKey struct {
	Perm_auth_key_id  int64
	Nonce             int64
	Expires_at        int32
	Encrypted_message []byte
}

type TL_contacts_exportCard struct {
}

type TL_contacts_importCard struct {
	Export_card []int32
}

type TL_messages_readMessageContents struct {
	Id []int32
}

type TL_account_checkUsername struct {
	Username string
}

type TL_account_updateUsername struct {
	Username string
}

type TL_account_getPrivacy struct {
	Key TL // InputPrivacyKey
}

type TL_account_setPrivacy struct {
	Key   TL   // InputPrivacyKey
	Rules []TL // InputPrivacyRule
}

type TL_account_deleteAccount struct {
	Reason string
}

type TL_account_getAccountTTL struct {
}

type TL_account_setAccountTTL struct {
	Ttl TL // AccountDaysTTL
}

type TL_invokeWithLayer struct {
	Layer int32
	Query TL
}

type TL_contacts_resolveUsername struct {
	Username string
}

type TL_account_sendChangePhoneCode struct {
	Flags int32
	// Allow_flashcall	bool // flags_0?true
	Phone_number   string
	Current_number TL // flags_0?Bool
}

type TL_account_changePhone struct {
	Phone_number    string
	Phone_code_hash string
	Phone_code      string
}

type TL_messages_getAllStickers struct {
	Hash int32
}

type TL_account_updateDeviceLocked struct {
	Period int32
}

type TL_account_getPassword struct {
}

type TL_auth_checkPassword struct {
	Password_hash []byte
}

type TL_messages_getWebPagePreview struct {
	Message string
}

type TL_account_getAuthorizations struct {
}

type TL_account_resetAuthorization struct {
	Hash int64
}

type TL_account_getPasswordSettings struct {
	Current_password_hash []byte
}

type TL_account_updatePasswordSettings struct {
	Current_password_hash []byte
	New_settings          TL // account_PasswordInputSettings
}

type TL_auth_requestPasswordRecovery struct {
}

type TL_auth_recoverPassword struct {
	Code string
}

type TL_invokeWithoutUpdates struct {
	Query TL
}

type TL_messages_exportChatInvite struct {
	Chat_id int32
}

type TL_messages_checkChatInvite struct {
	Hash string
}

type TL_messages_importChatInvite struct {
	Hash string
}

type TL_messages_getStickerSet struct {
	Stickerset TL // InputStickerSet
}

type TL_messages_installStickerSet struct {
	Stickerset TL // InputStickerSet
	Archived   TL // Bool
}

type TL_messages_uninstallStickerSet struct {
	Stickerset TL // InputStickerSet
}

type TL_auth_importBotAuthorization struct {
	Flags          int32
	Api_id         int32
	Api_hash       string
	Bot_auth_token string
}

type TL_messages_startBot struct {
	Bot         TL // InputUser
	Peer        TL // InputPeer
	Random_id   int64
	Start_param string
}

type TL_help_getAppChangelog struct {
	Prev_app_version string
}

type TL_messages_reportSpam struct {
	Peer TL // InputPeer
}

type TL_messages_getMessagesViews struct {
	Peer      TL // InputPeer
	Id        []int32
	Increment TL // Bool
}

type TL_updates_getChannelDifference struct {
	Flags int32
	// Force	bool // flags_0?true
	Channel TL // InputChannel
	Filter  TL // ChannelMessagesFilter
	Pts     int32
	Limit   int32
}

type TL_channels_readHistory struct {
	Channel TL // InputChannel
	Max_id  int32
}

type TL_channels_deleteMessages struct {
	Channel TL // InputChannel
	Id      []int32
}

type TL_channels_deleteUserHistory struct {
	Channel TL // InputChannel
	User_id TL // InputUser
}

type TL_channels_reportSpam struct {
	Channel TL // InputChannel
	User_id TL // InputUser
	Id      []int32
}

type TL_channels_getMessages struct {
	Channel TL // InputChannel
	Id      []int32
}

type TL_channels_getParticipants struct {
	Channel TL // InputChannel
	Filter  TL // ChannelParticipantsFilter
	Offset  int32
	Limit   int32
}

type TL_channels_getParticipant struct {
	Channel TL // InputChannel
	User_id TL // InputUser
}

type TL_channels_getChannels struct {
	Id []TL // InputChannel
}

type TL_channels_getFullChannel struct {
	Channel TL // InputChannel
}

type TL_channels_createChannel struct {
	Flags int32
	// Broadcast	bool // flags_0?true
	// Megagroup	bool // flags_1?true
	Title string
	About string
}

type TL_channels_editAbout struct {
	Channel TL // InputChannel
	About   string
}

type TL_channels_editAdmin struct {
	Channel      TL // InputChannel
	User_id      TL // InputUser
	Admin_rights TL // ChannelAdminRights
}

type TL_channels_editTitle struct {
	Channel TL // InputChannel
	Title   string
}

type TL_channels_editPhoto struct {
	Channel TL // InputChannel
	Photo   TL // InputChatPhoto
}

type TL_channels_checkUsername struct {
	Channel  TL // InputChannel
	Username string
}

type TL_channels_updateUsername struct {
	Channel  TL // InputChannel
	Username string
}

type TL_channels_joinChannel struct {
	Channel TL // InputChannel
}

type TL_channels_leaveChannel struct {
	Channel TL // InputChannel
}

type TL_channels_inviteToChannel struct {
	Channel TL   // InputChannel
	Users   []TL // InputUser
}

type TL_channels_exportInvite struct {
	Channel TL // InputChannel
}

type TL_channels_deleteChannel struct {
	Channel TL // InputChannel
}

type TL_messages_toggleChatAdmins struct {
	Chat_id int32
	Enabled TL // Bool
}

type TL_messages_editChatAdmin struct {
	Chat_id  int32
	User_id  TL // InputUser
	Is_admin TL // Bool
}

type TL_messages_migrateChat struct {
	Chat_id int32
}

type TL_messages_searchGlobal struct {
	Q           string
	Offset_date int32
	Offset_peer TL // InputPeer
	Offset_id   int32
	Limit       int32
}

type TL_account_reportPeer struct {
	Peer   TL // InputPeer
	Reason TL // ReportReason
}

type TL_messages_reorderStickerSets struct {
	Flags int32
	// Masks	bool // flags_0?true
	Order []int64
}

type TL_help_getTermsOfService struct {
}

type TL_messages_getDocumentByHash struct {
	Sha256    []byte
	Size      int32
	Mime_type string
}

type TL_messages_searchGifs struct {
	Q      string
	Offset int32
}

type TL_messages_getSavedGifs struct {
	Hash int32
}

type TL_messages_saveGif struct {
	Id     TL // InputDocument
	Unsave TL // Bool
}

type TL_messages_getInlineBotResults struct {
	Flags     int32
	Bot       TL // InputUser
	Peer      TL // InputPeer
	Geo_point TL // flags_0?InputGeoPoint
	Query     string
	Offset    string
}

type TL_messages_setInlineBotResults struct {
	Flags int32
	// Gallery	bool // flags_0?true
	// Private	bool // flags_1?true
	Query_id    int64
	Results     []TL // InputBotInlineResult
	Cache_time  int32
	Next_offset string
	Switch_pm   TL // flags_3?InlineBotSwitchPM
}

type TL_messages_sendInlineBotResult struct {
	Flags int32
	// Silent	bool // flags_5?true
	// Background	bool // flags_6?true
	// Clear_draft	bool // flags_7?true
	Peer            TL // InputPeer
	Reply_to_msg_id int32
	Random_id       int64
	Query_id        int64
	Id              string
}

type TL_channels_toggleInvites struct {
	Channel TL // InputChannel
	Enabled TL // Bool
}

type TL_channels_exportMessageLink struct {
	Channel TL // InputChannel
	Id      int32
}

type TL_channels_toggleSignatures struct {
	Channel TL // InputChannel
	Enabled TL // Bool
}

type TL_messages_hideReportSpam struct {
	Peer TL // InputPeer
}

type TL_messages_getPeerSettings struct {
	Peer TL // InputPeer
}

type TL_channels_updatePinnedMessage struct {
	Flags int32
	// Silent	bool // flags_0?true
	Channel TL // InputChannel
	Id      int32
}

type TL_auth_resendCode struct {
	Phone_number    string
	Phone_code_hash string
}

type TL_auth_cancelCode struct {
	Phone_number    string
	Phone_code_hash string
}

type TL_messages_getMessageEditData struct {
	Peer TL // InputPeer
	Id   int32
}

type TL_messages_editMessage struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	Peer         TL // InputPeer
	Id           int32
	Message      string
	Reply_markup TL   // flags_2?ReplyMarkup
	Entities     []TL // MessageEntity
}

type TL_messages_editInlineBotMessage struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	Id           TL // InputBotInlineMessageID
	Message      string
	Reply_markup TL   // flags_2?ReplyMarkup
	Entities     []TL // MessageEntity
}

type TL_messages_getBotCallbackAnswer struct {
	Flags int32
	// Game	bool // flags_1?true
	Peer   TL // InputPeer
	Msg_id int32
	Data   []byte
}

type TL_messages_setBotCallbackAnswer struct {
	Flags int32
	// Alert	bool // flags_1?true
	Query_id   int64
	Message    string
	Url        string
	Cache_time int32
}

type TL_contacts_getTopPeers struct {
	Flags int32
	// Correspondents	bool // flags_0?true
	// Bots_pm	bool // flags_1?true
	// Bots_inline	bool // flags_2?true
	// Phone_calls	bool // flags_3?true
	// Groups	bool // flags_10?true
	// Channels	bool // flags_15?true
	Offset int32
	Limit  int32
	Hash   int32
}

type TL_contacts_resetTopPeerRating struct {
	Category TL // TopPeerCategory
	Peer     TL // InputPeer
}

type TL_messages_getPeerDialogs struct {
	Peers []TL // InputPeer
}

type TL_messages_saveDraft struct {
	Flags int32
	// No_webpage	bool // flags_1?true
	Reply_to_msg_id int32
	Peer            TL // InputPeer
	Message         string
	Entities        []TL // MessageEntity
}

type TL_messages_getAllDrafts struct {
}

type TL_account_sendConfirmPhoneCode struct {
	Flags int32
	// Allow_flashcall	bool // flags_0?true
	Hash           string
	Current_number TL // flags_0?Bool
}

type TL_account_confirmPhone struct {
	Phone_code_hash string
	Phone_code      string
}

type TL_messages_getFeaturedStickers struct {
	Hash int32
}

type TL_messages_readFeaturedStickers struct {
	Id []int64
}

type TL_messages_getRecentStickers struct {
	Flags int32
	// Attached	bool // flags_0?true
	Hash int32
}

type TL_messages_saveRecentSticker struct {
	Flags int32
	// Attached	bool // flags_0?true
	Id     TL // InputDocument
	Unsave TL // Bool
}

type TL_messages_clearRecentStickers struct {
	Flags int32
	// Attached	bool // flags_0?true
}

type TL_messages_getArchivedStickers struct {
	Flags int32
	// Masks	bool // flags_0?true
	Offset_id int64
	Limit     int32
}

type TL_channels_getAdminedPublicChannels struct {
}

type TL_auth_dropTempAuthKeys struct {
	Except_auth_keys []int64
}

type TL_messages_setGameScore struct {
	Flags int32
	// Edit_message	bool // flags_0?true
	// Force	bool // flags_1?true
	Peer    TL // InputPeer
	Id      int32
	User_id TL // InputUser
	Score   int32
}

type TL_messages_setInlineGameScore struct {
	Flags int32
	// Edit_message	bool // flags_0?true
	// Force	bool // flags_1?true
	Id      TL // InputBotInlineMessageID
	User_id TL // InputUser
	Score   int32
}

type TL_messages_getMaskStickers struct {
	Hash int32
}

type TL_messages_getAttachedStickers struct {
	Media TL // InputStickeredMedia
}

type TL_messages_getGameHighScores struct {
	Peer    TL // InputPeer
	Id      int32
	User_id TL // InputUser
}

type TL_messages_getInlineGameHighScores struct {
	Id      TL // InputBotInlineMessageID
	User_id TL // InputUser
}

type TL_messages_getCommonChats struct {
	User_id TL // InputUser
	Max_id  int32
	Limit   int32
}

type TL_messages_getAllChats struct {
	Except_ids []int32
}

type TL_help_setBotUpdatesStatus struct {
	Pending_updates_count int32
	Message               string
}

type TL_messages_getWebPage struct {
	Url  string
	Hash int32
}

type TL_messages_toggleDialogPin struct {
	Flags int32
	// Pinned	bool // flags_0?true
	Peer TL // InputPeer
}

type TL_messages_reorderPinnedDialogs struct {
	Flags int32
	// Force	bool // flags_0?true
	Order []TL // InputPeer
}

type TL_messages_getPinnedDialogs struct {
}

type TL_phone_requestCall struct {
	User_id   TL // InputUser
	Random_id int32
	G_a_hash  []byte
	Protocol  TL // PhoneCallProtocol
}

type TL_phone_acceptCall struct {
	Peer     TL // InputPhoneCall
	G_b      []byte
	Protocol TL // PhoneCallProtocol
}

type TL_phone_discardCall struct {
	Peer          TL // InputPhoneCall
	Duration      int32
	Reason        TL // PhoneCallDiscardReason
	Connection_id int64
}

type TL_phone_receivedCall struct {
	Peer TL // InputPhoneCall
}

type TL_messages_reportEncryptedSpam struct {
	Peer TL // InputEncryptedChat
}

type TL_payments_getPaymentForm struct {
	Msg_id int32
}

type TL_payments_sendPaymentForm struct {
	Flags              int32
	Msg_id             int32
	Requested_info_id  string
	Shipping_option_id string
	Credentials        TL // InputPaymentCredentials
}

type TL_account_getTmpPassword struct {
	Password_hash []byte
	Period        int32
}

type TL_messages_setBotShippingResults struct {
	Flags            int32
	Query_id         int64
	Error            string
	Shipping_options []TL // ShippingOption
}

type TL_messages_setBotPrecheckoutResults struct {
	Flags int32
	// Success	bool // flags_1?true
	Query_id int64
	Error    string
}

type TL_upload_getWebFile struct {
	Location TL // InputWebFileLocation
	Offset   int32
	Limit    int32
}

type TL_bots_sendCustomRequest struct {
	Custom_method string
	Params        TL // DataJSON
}

type TL_bots_answerWebhookJSONQuery struct {
	Query_id int64
	Data     TL // DataJSON
}

type TL_payments_getPaymentReceipt struct {
	Msg_id int32
}

type TL_payments_validateRequestedInfo struct {
	Flags int32
	// Save	bool // flags_0?true
	Msg_id int32
	Info   TL // PaymentRequestedInfo
}

type TL_payments_getSavedInfo struct {
}

type TL_payments_clearSavedInfo struct {
	Flags int32
	// Credentials	bool // flags_0?true
	// Info	bool // flags_1?true
}

type TL_phone_getCallConfig struct {
}

type TL_phone_confirmCall struct {
	Peer            TL // InputPhoneCall
	G_a             []byte
	Key_fingerprint int64
	Protocol        TL // PhoneCallProtocol
}

type TL_phone_setCallRating struct {
	Peer    TL // InputPhoneCall
	Rating  int32
	Comment string
}

type TL_phone_saveCallDebug struct {
	Peer  TL // InputPhoneCall
	Debug TL // DataJSON
}

type TL_upload_getCdnFile struct {
	File_token []byte
	Offset     int32
	Limit      int32
}

type TL_upload_reuploadCdnFile struct {
	File_token    []byte
	Request_token []byte
}

type TL_help_getCdnConfig struct {
}

type TL_messages_uploadMedia struct {
	Peer  TL // InputPeer
	Media TL // InputMedia
}

type TL_stickers_createStickerSet struct {
	Flags int32
	// Masks	bool // flags_0?true
	User_id    TL // InputUser
	Title      string
	Short_name string
	Stickers   []TL // InputStickerSetItem
}

type TL_langpack_getLangPack struct {
	Lang_code string
}

type TL_langpack_getStrings struct {
	Lang_code string
	Keys      []string
}

type TL_langpack_getDifference struct {
	From_version int32
}

type TL_langpack_getLanguages struct {
}

type TL_channels_editBanned struct {
	Channel       TL // InputChannel
	User_id       TL // InputUser
	Banned_rights TL // ChannelBannedRights
}

type TL_channels_getAdminLog struct {
	Flags         int32
	Channel       TL // InputChannel
	Q             string
	Events_filter TL   // flags_0?ChannelAdminLogEventsFilter
	Admins        []TL // InputUser
	Max_id        int64
	Min_id        int64
	Limit         int32
}

type TL_stickers_removeStickerFromSet struct {
	Sticker TL // InputDocument
}

type TL_stickers_changeStickerPosition struct {
	Sticker  TL // InputDocument
	Position int32
}

type TL_stickers_addStickerToSet struct {
	Stickerset TL // InputStickerSet
	Sticker    TL // InputStickerSetItem
}

type TL_messages_sendScreenshotNotification struct {
	Peer            TL // InputPeer
	Reply_to_msg_id int32
	Random_id       int64
}

type TL_upload_getCdnFileHashes struct {
	File_token []byte
	Offset     int32
}

type TL_messages_getUnreadMentions struct {
	Peer       TL // InputPeer
	Offset_id  int32
	Add_offset int32
	Limit      int32
	Max_id     int32
	Min_id     int32
}

type TL_messages_faveSticker struct {
	Id     TL // InputDocument
	Unfave TL // Bool
}

type TL_channels_setStickers struct {
	Channel    TL // InputChannel
	Stickerset TL // InputStickerSet
}

type TL_contacts_resetSaved struct {
}

type TL_messages_getFavedStickers struct {
	Hash int32
}

type TL_channels_readMessageContents struct {
	Channel TL // InputChannel
	Id      []int32
}

func (e TL_boolFalse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolFalse)
	return x.buf
}

func (e TL_boolTrue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolTrue)
	return x.buf
}

func (e TL_error) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_error)
	x.Int(e.Code)
	x.String(e.Text)
	return x.buf
}

func (e TL_null) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_null)
	return x.buf
}

func (e TL_vector) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_vector)
	x.Vector(e.T)
	return x.buf
}

func (e TL_inputPeerEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerEmpty)
	return x.buf
}

func (e TL_inputPeerSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerSelf)
	return x.buf
}

func (e TL_inputPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChat)
	x.Int(e.Chat_id)
	return x.buf
}

func (e TL_inputUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserEmpty)
	return x.buf
}

func (e TL_inputUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserSelf)
	return x.buf
}

func (e TL_inputPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneContact)
	x.Long(e.Client_id)
	x.String(e.Phone)
	x.String(e.First_name)
	x.String(e.Last_name)
	return x.buf
}

func (e TL_inputFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFile)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	x.String(e.Md5_checksum)
	return x.buf
}

func (e TL_inputMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaEmpty)
	return x.buf
}

func (e TL_inputMediaUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedPhoto)
	x.Int(e.Flags)
	x.Bytes(e.File.encode())
	x.String(e.Caption)
	x.Vector(e.Stickers)
	x.Int(e.Ttl_seconds)
	return x.buf
}

func (e TL_inputMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhoto)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.buf
}

func (e TL_inputMediaGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGeoPoint)
	x.Bytes(e.Geo_point.encode())
	return x.buf
}

func (e TL_inputMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaContact)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	return x.buf
}

func (e TL_inputChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhotoEmpty)
	return x.buf
}

func (e TL_inputChatUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatUploadedPhoto)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_inputChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_inputGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPointEmpty)
	return x.buf
}

func (e TL_inputGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPoint)
	x.Double(e.Lat)
	x.Double(e.Long)
	return x.buf
}

func (e TL_inputPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoEmpty)
	return x.buf
}

func (e TL_inputPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoto)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_inputFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileLocation)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.buf
}

func (e TL_inputAppEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAppEvent)
	x.Double(e.Time)
	x.String(e._Type)
	x.Long(e.Peer)
	x.String(e.Data)
	return x.buf
}

func (e TL_peerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerUser)
	x.Int(e.User_id)
	return x.buf
}

func (e TL_peerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChat)
	x.Int(e.Chat_id)
	return x.buf
}

func (e TL_storage_fileUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileUnknown)
	return x.buf
}

func (e TL_storage_fileJpeg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileJpeg)
	return x.buf
}

func (e TL_storage_fileGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileGif)
	return x.buf
}

func (e TL_storage_filePng) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePng)
	return x.buf
}

func (e TL_storage_fileMp3) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp3)
	return x.buf
}

func (e TL_storage_fileMov) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMov)
	return x.buf
}

func (e TL_storage_filePartial) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePartial)
	return x.buf
}

func (e TL_storage_fileMp4) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp4)
	return x.buf
}

func (e TL_storage_fileWebp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileWebp)
	return x.buf
}

func (e TL_fileLocationUnavailable) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocationUnavailable)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.buf
}

func (e TL_fileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocation)
	x.Int(e.Dc_id)
	x.Long(e.Volume_id)
	x.Int(e.Local_id)
	x.Long(e.Secret)
	return x.buf
}

func (e TL_userEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userEmpty)
	x.Int(e.Id)
	return x.buf
}

func (e TL_userProfilePhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhotoEmpty)
	return x.buf
}

func (e TL_userProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhoto)
	x.Long(e.Photo_id)
	x.Bytes(e.Photo_small.encode())
	x.Bytes(e.Photo_big.encode())
	return x.buf
}

func (e TL_userStatusEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusEmpty)
	return x.buf
}

func (e TL_userStatusOnline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOnline)
	x.Int(e.Expires)
	return x.buf
}

func (e TL_userStatusOffline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOffline)
	x.Int(e.Was_online)
	return x.buf
}

func (e TL_chatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatEmpty)
	x.Int(e.Id)
	return x.buf
}

func (e TL_chat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chat)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.Participants_count)
	x.Int(e.Date)
	x.Int(e.Version)
	x.Bytes(e.Migrated_to.encode())
	return x.buf
}

func (e TL_chatForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatForbidden)
	x.Int(e.Id)
	x.String(e.Title)
	return x.buf
}

func (e TL_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatFull)
	x.Int(e.Id)
	x.Bytes(e.Participants.encode())
	x.Bytes(e.Chat_photo.encode())
	x.Bytes(e.Notify_settings.encode())
	x.Bytes(e.Exported_invite.encode())
	x.Vector(e.Bot_info)
	return x.buf
}

func (e TL_chatParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipant)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.buf
}

func (e TL_chatParticipantsForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantsForbidden)
	x.Int(e.Flags)
	x.Int(e.Chat_id)
	x.Bytes(e.Self_participant.encode())
	return x.buf
}

func (e TL_chatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipants)
	x.Int(e.Chat_id)
	x.Vector(e.Participants)
	x.Int(e.Version)
	return x.buf
}

func (e TL_chatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhotoEmpty)
	return x.buf
}

func (e TL_chatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhoto)
	x.Bytes(e.Photo_small.encode())
	x.Bytes(e.Photo_big.encode())
	return x.buf
}

func (e TL_messageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEmpty)
	x.Int(e.Id)
	return x.buf
}

func (e TL_message) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_message)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.From_id)
	x.Bytes(e.To_id.encode())
	x.Bytes(e.Fwd_from.encode())
	x.Int(e.Via_bot_id)
	x.Int(e.Reply_to_msg_id)
	x.Int(e.Date)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	x.Bytes(e.Reply_markup.encode())
	x.Vector(e.Entities)
	x.Int(e.Views)
	x.Int(e.Edit_date)
	x.String(e.Post_author)
	return x.buf
}

func (e TL_messageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageService)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.From_id)
	x.Bytes(e.To_id.encode())
	x.Int(e.Reply_to_msg_id)
	x.Int(e.Date)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_messageMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaEmpty)
	return x.buf
}

func (e TL_messageMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaPhoto)
	x.Int(e.Flags)
	x.Bytes(e.Photo.encode())
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.buf
}

func (e TL_messageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGeo)
	x.Bytes(e.Geo.encode())
	return x.buf
}

func (e TL_messageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaContact)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.Int(e.User_id)
	return x.buf
}

func (e TL_messageMediaUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaUnsupported)
	return x.buf
}

func (e TL_messageActionEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionEmpty)
	return x.buf
}

func (e TL_messageActionChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatCreate)
	x.String(e.Title)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TL_messageActionChatEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditTitle)
	x.String(e.Title)
	return x.buf
}

func (e TL_messageActionChatEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditPhoto)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TL_messageActionChatDeletePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeletePhoto)
	return x.buf
}

func (e TL_messageActionChatAddUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatAddUser)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TL_messageActionChatDeleteUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeleteUser)
	x.Int(e.User_id)
	return x.buf
}

func (e TL_dialog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialog)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Top_message)
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Int(e.Unread_mentions_count)
	x.Bytes(e.Notify_settings.encode())
	x.Int(e.Pts)
	x.Bytes(e.Draft.encode())
	return x.buf
}

func (e TL_photoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoEmpty)
	x.Long(e.Id)
	return x.buf
}

func (e TL_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photo)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Vector(e.Sizes)
	return x.buf
}

func (e TL_photoSizeEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSizeEmpty)
	x.String(e._Type)
	return x.buf
}

func (e TL_photoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSize)
	x.String(e._Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	return x.buf
}

func (e TL_photoCachedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoCachedSize)
	x.String(e._Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_geoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPointEmpty)
	return x.buf
}

func (e TL_geoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPoint)
	x.Double(e.Long)
	x.Double(e.Lat)
	return x.buf
}

func (e TL_auth_checkedPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkedPhone)
	x.Bytes(e.Phone_registered.encode())
	return x.buf
}

func (e TL_auth_sentCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCode)
	x.Int(e.Flags)
	x.Bytes(e._Type.encode())
	x.String(e.Phone_code_hash)
	x.Bytes(e.Next_type.encode())
	x.Int(e.Timeout)
	return x.buf
}

func (e TL_auth_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_authorization)
	x.Int(e.Flags)
	x.Int(e.Tmp_sessions)
	x.Bytes(e.User.encode())
	return x.buf
}

func (e TL_auth_exportedAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportedAuthorization)
	x.Int(e.Id)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_inputNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_inputNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyUsers)
	return x.buf
}

func (e TL_inputNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyChats)
	return x.buf
}

func (e TL_inputNotifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyAll)
	return x.buf
}

func (e TL_inputPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifySettings)
	x.Int(e.Flags)
	x.Int(e.Mute_until)
	x.String(e.Sound)
	return x.buf
}

func (e TL_peerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsEmpty)
	return x.buf
}

func (e TL_peerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsAll)
	return x.buf
}

func (e TL_peerNotifySettingsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettingsEmpty)
	return x.buf
}

func (e TL_peerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettings)
	x.Int(e.Flags)
	x.Int(e.Mute_until)
	x.String(e.Sound)
	return x.buf
}

func (e TL_wallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaper)
	x.Int(e.Id)
	x.String(e.Title)
	x.Vector(e.Sizes)
	x.Int(e.Color)
	return x.buf
}

func (e TL_userFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userFull)
	x.Int(e.Flags)
	x.Bytes(e.User.encode())
	x.String(e.About)
	x.Bytes(e.Link.encode())
	x.Bytes(e.Profile_photo.encode())
	x.Bytes(e.Notify_settings.encode())
	x.Bytes(e.Bot_info.encode())
	x.Int(e.Common_chats_count)
	return x.buf
}

func (e TL_contact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contact)
	x.Int(e.User_id)
	x.Bytes(e.Mutual.encode())
	return x.buf
}

func (e TL_importedContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_importedContact)
	x.Int(e.User_id)
	x.Long(e.Client_id)
	return x.buf
}

func (e TL_contactBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactBlocked)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.buf
}

func (e TL_contactStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactStatus)
	x.Int(e.User_id)
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TL_contacts_link) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_link)
	x.Bytes(e.My_link.encode())
	x.Bytes(e.Foreign_link.encode())
	x.Bytes(e.User.encode())
	return x.buf
}

func (e TL_contacts_contacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contacts)
	x.Vector(e.Contacts)
	x.Int(e.Saved_count)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_contacts_contactsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contactsNotModified)
	return x.buf
}

func (e TL_contacts_importedContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importedContacts)
	x.Vector(e.Imported)
	x.Vector(e.Popular_invites)
	x.VectorLong(e.Retry_contacts)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_contacts_blocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blocked)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_contacts_blockedSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blockedSlice)
	x.Int(e.Count)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_contacts_found) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_found)
	x.Vector(e.Results)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_dialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_dialogsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogsSlice)
	x.Int(e.Count)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_messages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messages)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_messagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messagesSlice)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_chats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chats)
	x.Vector(e.Chats)
	return x.buf
}

func (e TL_messages_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chatFull)
	x.Bytes(e.Full_chat.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_affectedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedHistory)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Offset)
	return x.buf
}

func (e TL_inputMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterEmpty)
	return x.buf
}

func (e TL_inputMessagesFilterPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotos)
	return x.buf
}

func (e TL_inputMessagesFilterVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVideo)
	return x.buf
}

func (e TL_inputMessagesFilterPhotoVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideo)
	return x.buf
}

func (e TL_updateNewMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_updateMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateMessageID)
	x.Int(e.Id)
	x.Long(e.Random_id)
	return x.buf
}

func (e TL_updateDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteMessages)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_updateUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserTyping)
	x.Int(e.User_id)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_updateChatUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatUserTyping)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_updateChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipants)
	x.Bytes(e.Participants.encode())
	return x.buf
}

func (e TL_updateUserStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserStatus)
	x.Int(e.User_id)
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TL_updateUserName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserName)
	x.Int(e.User_id)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.String(e.Username)
	return x.buf
}

func (e TL_updateUserPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhoto)
	x.Int(e.User_id)
	x.Int(e.Date)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Previous.encode())
	return x.buf
}

func (e TL_updateContactRegistered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactRegistered)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.buf
}

func (e TL_updateContactLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactLink)
	x.Int(e.User_id)
	x.Bytes(e.My_link.encode())
	x.Bytes(e.Foreign_link.encode())
	return x.buf
}

func (e TL_updates_state) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_state)
	x.Int(e.Pts)
	x.Int(e.Qts)
	x.Int(e.Date)
	x.Int(e.Seq)
	x.Int(e.Unread_count)
	return x.buf
}

func (e TL_updates_differenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceEmpty)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e TL_updates_difference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_difference)
	x.Vector(e.New_messages)
	x.Vector(e.New_encrypted_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

func (e TL_updates_differenceSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceSlice)
	x.Vector(e.New_messages)
	x.Vector(e.New_encrypted_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.Intermediate_state.encode())
	return x.buf
}

func (e TL_updatesTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesTooLong)
	return x.buf
}

func (e TL_updateShortMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.User_id)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	x.Bytes(e.Fwd_from.encode())
	x.Int(e.Via_bot_id)
	x.Int(e.Reply_to_msg_id)
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_updateShortChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortChatMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.From_id)
	x.Int(e.Chat_id)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	x.Bytes(e.Fwd_from.encode())
	x.Int(e.Via_bot_id)
	x.Int(e.Reply_to_msg_id)
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_updateShort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShort)
	x.Bytes(e.Update.encode())
	x.Int(e.Date)
	return x.buf
}

func (e TL_updatesCombined) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesCombined)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq_start)
	x.Int(e.Seq)
	return x.buf
}

func (e TL_updates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e TL_photos_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photo)
	x.Bytes(e.Photo.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_upload_file) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_file)
	x.Bytes(e._Type.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_dcOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dcOption)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.Ip_address)
	x.Int(e.Port)
	return x.buf
}

func (e TL_config) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_config)
	x.Int(e.Flags)
	x.Int(e.Date)
	x.Int(e.Expires)
	x.Bytes(e.Test_mode.encode())
	x.Int(e.This_dc)
	x.Vector(e.Dc_options)
	x.Int(e.Chat_size_max)
	x.Int(e.Megagroup_size_max)
	x.Int(e.Forwarded_count_max)
	x.Int(e.Online_update_period_ms)
	x.Int(e.Offline_blur_timeout_ms)
	x.Int(e.Offline_idle_timeout_ms)
	x.Int(e.Online_cloud_timeout_ms)
	x.Int(e.Notify_cloud_delay_ms)
	x.Int(e.Notify_default_delay_ms)
	x.Int(e.Chat_big_size)
	x.Int(e.Push_chat_period_ms)
	x.Int(e.Push_chat_limit)
	x.Int(e.Saved_gifs_limit)
	x.Int(e.Edit_time_limit)
	x.Int(e.Rating_e_decay)
	x.Int(e.Stickers_recent_limit)
	x.Int(e.Stickers_faved_limit)
	x.Int(e.Tmp_sessions)
	x.Int(e.Pinned_dialogs_count_max)
	x.Int(e.Call_receive_timeout_ms)
	x.Int(e.Call_ring_timeout_ms)
	x.Int(e.Call_connect_timeout_ms)
	x.Int(e.Call_packet_timeout_ms)
	x.String(e.Me_url_prefix)
	x.String(e.Suggested_lang_code)
	x.Int(e.Lang_pack_version)
	x.Vector(e.Disabled_features)
	return x.buf
}

func (e TL_nearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_nearestDc)
	x.String(e.Country)
	x.Int(e.This_dc)
	x.Int(e.Nearest_dc)
	return x.buf
}

func (e TL_help_appUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_appUpdate)
	x.Int(e.Id)
	x.Bytes(e.Critical.encode())
	x.String(e.Url)
	x.String(e.Text)
	return x.buf
}

func (e TL_help_noAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_noAppUpdate)
	return x.buf
}

func (e TL_help_inviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_inviteText)
	x.String(e.Message)
	return x.buf
}

func (e TL_inputPeerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsEmpty)
	return x.buf
}

func (e TL_inputPeerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsAll)
	return x.buf
}

func (e TL_photos_photos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photos)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_photos_photosSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photosSlice)
	x.Int(e.Count)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_wallPaperSolid) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaperSolid)
	x.Int(e.Id)
	x.String(e.Title)
	x.Int(e.Bg_color)
	x.Int(e.Color)
	return x.buf
}

func (e TL_updateNewEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewEncryptedMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Qts)
	return x.buf
}

func (e TL_updateEncryptedChatTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedChatTyping)
	x.Int(e.Chat_id)
	return x.buf
}

func (e TL_updateEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryption)
	x.Bytes(e.Chat.encode())
	x.Int(e.Date)
	return x.buf
}

func (e TL_updateEncryptedMessagesRead) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedMessagesRead)
	x.Int(e.Chat_id)
	x.Int(e.Max_date)
	x.Int(e.Date)
	return x.buf
}

func (e TL_encryptedChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatEmpty)
	x.Int(e.Id)
	return x.buf
}

func (e TL_encryptedChatWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatWaiting)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	return x.buf
}

func (e TL_encryptedChatRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatRequested)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_a)
	return x.buf
}

func (e TL_encryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChat)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_a_or_b)
	x.Long(e.Key_fingerprint)
	return x.buf
}

func (e TL_encryptedChatDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatDiscarded)
	x.Int(e.Id)
	return x.buf
}

func (e TL_inputEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedChat)
	x.Int(e.Chat_id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_encryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFileEmpty)
	return x.buf
}

func (e TL_encryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFile)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Size)
	x.Int(e.Dc_id)
	x.Int(e.Key_fingerprint)
	return x.buf
}

func (e TL_inputEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileEmpty)
	return x.buf
}

func (e TL_inputEncryptedFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Md5_checksum)
	x.Int(e.Key_fingerprint)
	return x.buf
}

func (e TL_inputEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFile)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_inputEncryptedFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileLocation)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_encryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessage)
	x.Long(e.Random_id)
	x.Int(e.Chat_id)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_encryptedMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessageService)
	x.Long(e.Random_id)
	x.Int(e.Chat_id)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_messages_dhConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfigNotModified)
	x.StringBytes(e.Random)
	return x.buf
}

func (e TL_messages_dhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfig)
	x.Int(e.G)
	x.StringBytes(e.P)
	x.Int(e.Version)
	x.StringBytes(e.Random)
	return x.buf
}

func (e TL_messages_sentEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedMessage)
	x.Int(e.Date)
	return x.buf
}

func (e TL_messages_sentEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedFile)
	x.Int(e.Date)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_inputFileBig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileBig)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.String(e.Name)
	return x.buf
}

func (e TL_inputEncryptedFileBigUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileBigUploaded)
	x.Long(e.Id)
	x.Int(e.Parts)
	x.Int(e.Key_fingerprint)
	return x.buf
}

func (e TL_storage_filePdf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePdf)
	return x.buf
}

func (e TL_inputMessagesFilterDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterDocument)
	return x.buf
}

func (e TL_inputMessagesFilterPhotoVideoDocuments) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideoDocuments)
	return x.buf
}

func (e TL_updateChatParticipantAdd) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdd)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	x.Int(e.Version)
	return x.buf
}

func (e TL_updateChatParticipantDelete) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantDelete)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Int(e.Version)
	return x.buf
}

func (e TL_updateDcOptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDcOptions)
	x.Vector(e.Dc_options)
	return x.buf
}

func (e TL_inputMediaUploadedDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedDocument)
	x.Int(e.Flags)
	x.Bytes(e.File.encode())
	x.Bytes(e.Thumb.encode())
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	x.String(e.Caption)
	x.Vector(e.Stickers)
	x.Int(e.Ttl_seconds)
	return x.buf
}

func (e TL_inputMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocument)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.buf
}

func (e TL_messageMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaDocument)
	x.Int(e.Flags)
	x.Bytes(e.Document.encode())
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.buf
}

func (e TL_inputDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentEmpty)
	return x.buf
}

func (e TL_inputDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocument)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_inputDocumentFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentFileLocation)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Version)
	return x.buf
}

func (e TL_documentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentEmpty)
	x.Long(e.Id)
	return x.buf
}

func (e TL_document) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_document)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.String(e.Mime_type)
	x.Int(e.Size)
	x.Bytes(e.Thumb.encode())
	x.Int(e.Dc_id)
	x.Int(e.Version)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TL_help_support) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_support)
	x.String(e.Phone_number)
	x.Bytes(e.User.encode())
	return x.buf
}

func (e TL_notifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyAll)
	return x.buf
}

func (e TL_notifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyChats)
	return x.buf
}

func (e TL_notifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_notifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyUsers)
	return x.buf
}

func (e TL_updateUserBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserBlocked)
	x.Int(e.User_id)
	x.Bytes(e.Blocked.encode())
	return x.buf
}

func (e TL_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Notify_settings.encode())
	return x.buf
}

func (e TL_sendMessageTypingAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageTypingAction)
	return x.buf
}

func (e TL_sendMessageCancelAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageCancelAction)
	return x.buf
}

func (e TL_sendMessageRecordVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordVideoAction)
	return x.buf
}

func (e TL_sendMessageUploadVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadVideoAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_sendMessageRecordAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordAudioAction)
	return x.buf
}

func (e TL_sendMessageUploadAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadAudioAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_sendMessageUploadPhotoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadPhotoAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_sendMessageUploadDocumentAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadDocumentAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_sendMessageGeoLocationAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGeoLocationAction)
	return x.buf
}

func (e TL_sendMessageChooseContactAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageChooseContactAction)
	return x.buf
}

func (e TL_updateServiceNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateServiceNotification)
	x.Int(e.Flags)
	x.Int(e.Inbox_date)
	x.String(e._Type)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_userStatusRecently) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusRecently)
	return x.buf
}

func (e TL_userStatusLastWeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastWeek)
	return x.buf
}

func (e TL_userStatusLastMonth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastMonth)
	return x.buf
}

func (e TL_updatePrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

func (e TL_inputPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_privacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_inputPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowAll)
	return x.buf
}

func (e TL_inputPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowUsers)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowAll)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowUsers)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_privacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowContacts)
	return x.buf
}

func (e TL_privacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowAll)
	return x.buf
}

func (e TL_privacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TL_privacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowContacts)
	return x.buf
}

func (e TL_privacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowAll)
	return x.buf
}

func (e TL_privacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TL_account_privacyRules) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_privacyRules)
	x.Vector(e.Rules)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_accountDaysTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountDaysTTL)
	x.Int(e.Days)
	return x.buf
}

func (e TL_updateUserPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhone)
	x.Int(e.User_id)
	x.String(e.Phone)
	return x.buf
}

func (e TL_disabledFeature) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_disabledFeature)
	x.String(e.Feature)
	x.String(e.Description)
	return x.buf
}

func (e TL_documentAttributeImageSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeImageSize)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TL_documentAttributeAnimated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAnimated)
	return x.buf
}

func (e TL_documentAttributeSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeSticker)
	x.Int(e.Flags)
	x.String(e.Alt)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Mask_coords.encode())
	return x.buf
}

func (e TL_documentAttributeVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeVideo)
	x.Int(e.Flags)
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TL_documentAttributeAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAudio)
	x.Int(e.Flags)
	x.Int(e.Duration)
	x.String(e.Title)
	x.String(e.Performer)
	x.StringBytes(e.Waveform)
	return x.buf
}

func (e TL_documentAttributeFilename) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeFilename)
	x.String(e.File_name)
	return x.buf
}

func (e TL_messages_stickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickersNotModified)
	return x.buf
}

func (e TL_messages_stickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickers)
	x.String(e.Hash)
	x.Vector(e.Stickers)
	return x.buf
}

func (e TL_stickerPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerPack)
	x.String(e.Emoticon)
	x.VectorLong(e.Documents)
	return x.buf
}

func (e TL_messages_allStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickersNotModified)
	return x.buf
}

func (e TL_messages_allStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickers)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	return x.buf
}

func (e TL_account_noPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_noPassword)
	x.StringBytes(e.New_salt)
	x.String(e.Email_unconfirmed_pattern)
	return x.buf
}

func (e TL_account_password) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_password)
	x.StringBytes(e.Current_salt)
	x.StringBytes(e.New_salt)
	x.String(e.Hint)
	x.Bytes(e.Has_recovery.encode())
	x.String(e.Email_unconfirmed_pattern)
	return x.buf
}

func (e TL_updateReadHistoryInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryInbox)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_updateReadHistoryOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryOutbox)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_messages_affectedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedMessages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_contactLinkUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkUnknown)
	return x.buf
}

func (e TL_contactLinkNone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkNone)
	return x.buf
}

func (e TL_contactLinkHasPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkHasPhone)
	return x.buf
}

func (e TL_contactLinkContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkContact)
	return x.buf
}

func (e TL_updateWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateWebPage)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_webPageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageEmpty)
	x.Long(e.Id)
	return x.buf
}

func (e TL_webPagePending) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPagePending)
	x.Long(e.Id)
	x.Int(e.Date)
	return x.buf
}

func (e TL_webPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPage)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.String(e.Url)
	x.String(e.Display_url)
	x.Int(e.Hash)
	x.String(e._Type)
	x.String(e.Site_name)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.encode())
	x.String(e.Embed_url)
	x.String(e.Embed_type)
	x.Int(e.Embed_width)
	x.Int(e.Embed_height)
	x.Int(e.Duration)
	x.String(e.Author)
	x.Bytes(e.Document.encode())
	x.Bytes(e.Cached_page.encode())
	return x.buf
}

func (e TL_messageMediaWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaWebPage)
	x.Bytes(e.Webpage.encode())
	return x.buf
}

func (e TL_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authorization)
	x.Long(e.Hash)
	x.Int(e.Flags)
	x.String(e.Device_model)
	x.String(e.Platform)
	x.String(e.System_version)
	x.Int(e.Api_id)
	x.String(e.App_name)
	x.String(e.App_version)
	x.Int(e.Date_created)
	x.Int(e.Date_active)
	x.String(e.Ip)
	x.String(e.Country)
	x.String(e.Region)
	return x.buf
}

func (e TL_account_authorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_authorizations)
	x.Vector(e.Authorizations)
	return x.buf
}

func (e TL_account_passwordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_passwordSettings)
	x.String(e.Email)
	return x.buf
}

func (e TL_account_passwordInputSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_passwordInputSettings)
	x.Int(e.Flags)
	x.StringBytes(e.New_salt)
	x.StringBytes(e.New_password_hash)
	x.String(e.Hint)
	x.String(e.Email)
	return x.buf
}

func (e TL_auth_passwordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_passwordRecovery)
	x.String(e.Email_pattern)
	return x.buf
}

func (e TL_inputMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaVenue)
	x.Bytes(e.Geo_point.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	return x.buf
}

func (e TL_messageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaVenue)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	return x.buf
}

func (e TL_receivedNotifyMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_receivedNotifyMessage)
	x.Int(e.Id)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_chatInviteEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteEmpty)
	return x.buf
}

func (e TL_chatInviteExported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteExported)
	x.String(e.Link)
	return x.buf
}

func (e TL_chatInviteAlready) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteAlready)
	x.Bytes(e.Chat.encode())
	return x.buf
}

func (e TL_chatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInvite)
	x.Int(e.Flags)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.Participants_count)
	x.Vector(e.Participants)
	return x.buf
}

func (e TL_messageActionChatJoinedByLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatJoinedByLink)
	x.Int(e.Inviter_id)
	return x.buf
}

func (e TL_updateReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadMessagesContents)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_inputStickerSetEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetEmpty)
	return x.buf
}

func (e TL_inputStickerSetID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetID)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_inputStickerSetShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetShortName)
	x.String(e.Short_name)
	return x.buf
}

func (e TL_stickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSet)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Title)
	x.String(e.Short_name)
	x.Int(e.Count)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_stickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSet)
	x.Bytes(e.Set.encode())
	x.Vector(e.Packs)
	x.Vector(e.Documents)
	return x.buf
}

func (e TL_user) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_user)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.String(e.Username)
	x.String(e.Phone)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Status.encode())
	x.Int(e.Bot_info_version)
	x.String(e.Restriction_reason)
	x.String(e.Bot_inline_placeholder)
	x.String(e.Lang_code)
	return x.buf
}

func (e TL_botCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botCommand)
	x.String(e.Command)
	x.String(e.Description)
	return x.buf
}

func (e TL_botInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInfo)
	x.Int(e.User_id)
	x.String(e.Description)
	x.Vector(e.Commands)
	return x.buf
}

func (e TL_keyboardButton) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButton)
	x.String(e.Text)
	return x.buf
}

func (e TL_keyboardButtonRow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRow)
	x.Vector(e.Buttons)
	return x.buf
}

func (e TL_replyKeyboardHide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardHide)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_replyKeyboardForceReply) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardForceReply)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_replyKeyboardMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardMarkup)
	x.Int(e.Flags)
	x.Vector(e.Rows)
	return x.buf
}

func (e TL_inputMessagesFilterUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterUrl)
	return x.buf
}

func (e TL_inputPeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerUser)
	x.Int(e.User_id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_inputUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUser)
	x.Int(e.User_id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_messageEntityUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUnknown)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityMention) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMention)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityHashtag) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityHashtag)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityBotCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBotCommand)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityEmail)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBold)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityItalic)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityCode)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityPre) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityPre)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Language)
	return x.buf
}

func (e TL_messageEntityTextUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityTextUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Url)
	return x.buf
}

func (e TL_updateShortSentMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortSentMessage)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	x.Int(e.Date)
	x.Bytes(e.Media.encode())
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_inputPeerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChannel)
	x.Int(e.Channel_id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_peerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChannel)
	x.Int(e.Channel_id)
	return x.buf
}

func (e TL_channel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channel)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Title)
	x.String(e.Username)
	x.Bytes(e.Photo.encode())
	x.Int(e.Date)
	x.Int(e.Version)
	x.String(e.Restriction_reason)
	x.Bytes(e.Admin_rights.encode())
	x.Bytes(e.Banned_rights.encode())
	return x.buf
}

func (e TL_channelForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelForbidden)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Title)
	x.Int(e.Until_date)
	return x.buf
}

func (e TL_channelFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelFull)
	x.Int(e.Flags)
	x.Int(e.Id)
	x.String(e.About)
	x.Int(e.Participants_count)
	x.Int(e.Admins_count)
	x.Int(e.Kicked_count)
	x.Int(e.Banned_count)
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Bytes(e.Chat_photo.encode())
	x.Bytes(e.Notify_settings.encode())
	x.Bytes(e.Exported_invite.encode())
	x.Vector(e.Bot_info)
	x.Int(e.Migrated_from_chat_id)
	x.Int(e.Migrated_from_max_id)
	x.Int(e.Pinned_msg_id)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_messageActionChannelCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelCreate)
	x.String(e.Title)
	return x.buf
}

func (e TL_messages_channelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_channelMessages)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_updateChannelTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelTooLong)
	x.Int(e.Flags)
	x.Int(e.Channel_id)
	x.Int(e.Pts)
	return x.buf
}

func (e TL_updateChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannel)
	x.Int(e.Channel_id)
	return x.buf
}

func (e TL_updateNewChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_updateReadChannelInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelInbox)
	x.Int(e.Channel_id)
	x.Int(e.Max_id)
	return x.buf
}

func (e TL_updateDeleteChannelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteChannelMessages)
	x.Int(e.Channel_id)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_updateChannelMessageViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelMessageViews)
	x.Int(e.Channel_id)
	x.Int(e.Id)
	x.Int(e.Views)
	return x.buf
}

func (e TL_inputChannelEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannelEmpty)
	return x.buf
}

func (e TL_inputChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannel)
	x.Int(e.Channel_id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_contacts_resolvedPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolvedPeer)
	x.Bytes(e.Peer.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messageRange) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageRange)
	x.Int(e.Min_id)
	x.Int(e.Max_id)
	return x.buf
}

func (e TL_updates_channelDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifferenceEmpty)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Timeout)
	return x.buf
}

func (e TL_updates_channelDifferenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifferenceTooLong)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Timeout)
	x.Int(e.Top_message)
	x.Int(e.Read_inbox_max_id)
	x.Int(e.Read_outbox_max_id)
	x.Int(e.Unread_count)
	x.Int(e.Unread_mentions_count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_updates_channelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifference)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Timeout)
	x.Vector(e.New_messages)
	x.Vector(e.Other_updates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_channelMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilterEmpty)
	return x.buf
}

func (e TL_channelMessagesFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilter)
	x.Int(e.Flags)
	x.Vector(e.Ranges)
	return x.buf
}

func (e TL_channelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipant)
	x.Int(e.User_id)
	x.Int(e.Date)
	return x.buf
}

func (e TL_channelParticipantSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantSelf)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.buf
}

func (e TL_channelParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantCreator)
	x.Int(e.User_id)
	return x.buf
}

func (e TL_channelParticipantsRecent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsRecent)
	return x.buf
}

func (e TL_channelParticipantsAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsAdmins)
	return x.buf
}

func (e TL_channelParticipantsKicked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsKicked)
	x.String(e.Q)
	return x.buf
}

func (e TL_channels_channelParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipants)
	x.Int(e.Count)
	x.Vector(e.Participants)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_channels_channelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipant)
	x.Bytes(e.Participant.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_true) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_true)
	return x.buf
}

func (e TL_chatParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantCreator)
	x.Int(e.User_id)
	return x.buf
}

func (e TL_chatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantAdmin)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Date)
	return x.buf
}

func (e TL_updateChatAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatAdmins)
	x.Int(e.Chat_id)
	x.Bytes(e.Enabled.encode())
	x.Int(e.Version)
	return x.buf
}

func (e TL_updateChatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdmin)
	x.Int(e.Chat_id)
	x.Int(e.User_id)
	x.Bytes(e.Is_admin.encode())
	x.Int(e.Version)
	return x.buf
}

func (e TL_messageActionChatMigrateTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatMigrateTo)
	x.Int(e.Channel_id)
	return x.buf
}

func (e TL_messageActionChannelMigrateFrom) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelMigrateFrom)
	x.String(e.Title)
	x.Int(e.Chat_id)
	return x.buf
}

func (e TL_channelParticipantsBots) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsBots)
	return x.buf
}

func (e TL_inputReportReasonSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonSpam)
	return x.buf
}

func (e TL_inputReportReasonViolence) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonViolence)
	return x.buf
}

func (e TL_inputReportReasonPornography) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonPornography)
	return x.buf
}

func (e TL_inputReportReasonOther) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonOther)
	x.String(e.Text)
	return x.buf
}

func (e TL_updateNewStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_updateStickerSetsOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSetsOrder)
	x.Int(e.Flags)
	x.VectorLong(e.Order)
	return x.buf
}

func (e TL_updateStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSets)
	return x.buf
}

func (e TL_help_termsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_termsOfService)
	x.String(e.Text)
	return x.buf
}

func (e TL_foundGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGif)
	x.String(e.Url)
	x.String(e.Thumb_url)
	x.String(e.Content_url)
	x.String(e.Content_type)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TL_inputMediaGifExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGifExternal)
	x.String(e.Url)
	x.String(e.Q)
	return x.buf
}

func (e TL_messages_foundGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_foundGifs)
	x.Int(e.Next_offset)
	x.Vector(e.Results)
	return x.buf
}

func (e TL_inputMessagesFilterGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterGif)
	return x.buf
}

func (e TL_updateSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateSavedGifs)
	return x.buf
}

func (e TL_updateBotInlineQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineQuery)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.String(e.Query)
	x.Bytes(e.Geo.encode())
	x.String(e.Offset)
	return x.buf
}

func (e TL_foundGifCached) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGifCached)
	x.String(e.Url)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Document.encode())
	return x.buf
}

func (e TL_messages_savedGifsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_savedGifsNotModified)
	return x.buf
}

func (e TL_messages_savedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_savedGifs)
	x.Int(e.Hash)
	x.Vector(e.Gifs)
	return x.buf
}

func (e TL_inputBotInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Caption)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageText)
	x.Int(e.Flags)
	x.String(e.Message)
	x.Vector(e.Entities)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e._Type)
	x.String(e.Title)
	x.String(e.Description)
	x.String(e.Url)
	x.String(e.Thumb_url)
	x.String(e.Content_url)
	x.String(e.Content_type)
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Duration)
	x.Bytes(e.Send_message.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Caption)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_botInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageText)
	x.Int(e.Flags)
	x.String(e.Message)
	x.Vector(e.Entities)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_botInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e._Type)
	x.String(e.Title)
	x.String(e.Description)
	x.String(e.Url)
	x.String(e.Thumb_url)
	x.String(e.Content_url)
	x.String(e.Content_type)
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Duration)
	x.Bytes(e.Send_message.encode())
	return x.buf
}

func (e TL_messages_botResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_botResults)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.String(e.Next_offset)
	x.Bytes(e.Switch_pm.encode())
	x.Vector(e.Results)
	x.Int(e.Cache_time)
	return x.buf
}

func (e TL_inputMessagesFilterVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVoice)
	return x.buf
}

func (e TL_inputMessagesFilterMusic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMusic)
	return x.buf
}

func (e TL_updateBotInlineSend) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineSend)
	x.Int(e.Flags)
	x.Int(e.User_id)
	x.String(e.Query)
	x.Bytes(e.Geo.encode())
	x.String(e.Id)
	x.Bytes(e.Msg_id.encode())
	return x.buf
}

func (e TL_inputPrivacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyChatInvite)
	return x.buf
}

func (e TL_privacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyChatInvite)
	return x.buf
}

func (e TL_updateEditChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_exportedMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_exportedMessageLink)
	x.String(e.Link)
	return x.buf
}

func (e TL_messageFwdHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageFwdHeader)
	x.Int(e.Flags)
	x.Int(e.From_id)
	x.Int(e.Date)
	x.Int(e.Channel_id)
	x.Int(e.Channel_post)
	x.String(e.Post_author)
	return x.buf
}

func (e TL_messageActionPinMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPinMessage)
	return x.buf
}

func (e TL_peerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerSettings)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_updateChannelPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelPinnedMessage)
	x.Int(e.Channel_id)
	x.Int(e.Id)
	return x.buf
}

func (e TL_keyboardButtonUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonUrl)
	x.String(e.Text)
	x.String(e.Url)
	return x.buf
}

func (e TL_keyboardButtonCallback) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonCallback)
	x.String(e.Text)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_keyboardButtonRequestPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestPhone)
	x.String(e.Text)
	return x.buf
}

func (e TL_keyboardButtonRequestGeoLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestGeoLocation)
	x.String(e.Text)
	return x.buf
}

func (e TL_auth_codeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeSms)
	return x.buf
}

func (e TL_auth_codeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeCall)
	return x.buf
}

func (e TL_auth_codeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeFlashCall)
	return x.buf
}

func (e TL_auth_sentCodeTypeApp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeApp)
	x.Int(e.Length)
	return x.buf
}

func (e TL_auth_sentCodeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeSms)
	x.Int(e.Length)
	return x.buf
}

func (e TL_auth_sentCodeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeCall)
	x.Int(e.Length)
	return x.buf
}

func (e TL_auth_sentCodeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeFlashCall)
	x.String(e.Pattern)
	return x.buf
}

func (e TL_keyboardButtonSwitchInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonSwitchInline)
	x.Int(e.Flags)
	x.String(e.Text)
	x.String(e.Query)
	return x.buf
}

func (e TL_replyInlineMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyInlineMarkup)
	x.Vector(e.Rows)
	return x.buf
}

func (e TL_messages_botCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_botCallbackAnswer)
	x.Int(e.Flags)
	x.String(e.Message)
	x.String(e.Url)
	x.Int(e.Cache_time)
	return x.buf
}

func (e TL_updateBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.Bytes(e.Peer.encode())
	x.Int(e.Msg_id)
	x.Long(e.Chat_instance)
	x.StringBytes(e.Data)
	x.String(e.Game_short_name)
	return x.buf
}

func (e TL_messages_messageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messageEditData)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_updateEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_inputBotInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.Geo_point.encode())
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.Geo_point.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.Geo.encode())
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.Venue_id)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.Phone_number)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineResultPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultPhoto)
	x.String(e.Id)
	x.String(e._Type)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Send_message.encode())
	return x.buf
}

func (e TL_inputBotInlineResultDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultDocument)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e._Type)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Document.encode())
	x.Bytes(e.Send_message.encode())
	return x.buf
}

func (e TL_botInlineMediaResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMediaResult)
	x.Int(e.Flags)
	x.String(e.Id)
	x.String(e._Type)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Document.encode())
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Send_message.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageID)
	x.Int(e.Dc_id)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_updateInlineBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateInlineBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.Bytes(e.Msg_id.encode())
	x.Long(e.Chat_instance)
	x.StringBytes(e.Data)
	x.String(e.Game_short_name)
	return x.buf
}

func (e TL_inlineBotSwitchPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inlineBotSwitchPM)
	x.String(e.Text)
	x.String(e.Start_param)
	return x.buf
}

func (e TL_messageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Int(e.User_id)
	return x.buf
}

func (e TL_inputMessageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Bytes(e.User_id.encode())
	return x.buf
}

func (e TL_messages_peerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_peerDialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

func (e TL_topPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeer)
	x.Bytes(e.Peer.encode())
	x.Double(e.Rating)
	return x.buf
}

func (e TL_topPeerCategoryBotsPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsPM)
	return x.buf
}

func (e TL_topPeerCategoryBotsInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsInline)
	return x.buf
}

func (e TL_topPeerCategoryCorrespondents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryCorrespondents)
	return x.buf
}

func (e TL_topPeerCategoryGroups) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryGroups)
	return x.buf
}

func (e TL_topPeerCategoryChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryChannels)
	return x.buf
}

func (e TL_topPeerCategoryPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryPeers)
	x.Bytes(e.Category.encode())
	x.Int(e.Count)
	x.Vector(e.Peers)
	return x.buf
}

func (e TL_contacts_topPeersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeersNotModified)
	return x.buf
}

func (e TL_contacts_topPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeers)
	x.Vector(e.Categories)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_inputMessagesFilterChatPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterChatPhotos)
	return x.buf
}

func (e TL_updateReadChannelOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelOutbox)
	x.Int(e.Channel_id)
	x.Int(e.Max_id)
	return x.buf
}

func (e TL_updateDraftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDraftMessage)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Draft.encode())
	return x.buf
}

func (e TL_draftMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessageEmpty)
	return x.buf
}

func (e TL_draftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessage)
	x.Int(e.Flags)
	x.Int(e.Reply_to_msg_id)
	x.String(e.Message)
	x.Vector(e.Entities)
	x.Int(e.Date)
	return x.buf
}

func (e TL_messageActionHistoryClear) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionHistoryClear)
	return x.buf
}

func (e TL_updateReadFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadFeaturedStickers)
	return x.buf
}

func (e TL_updateRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateRecentStickers)
	return x.buf
}

func (e TL_messages_featuredStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_featuredStickersNotModified)
	return x.buf
}

func (e TL_messages_featuredStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_featuredStickers)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	x.VectorLong(e.Unread)
	return x.buf
}

func (e TL_messages_recentStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_recentStickersNotModified)
	return x.buf
}

func (e TL_messages_recentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_recentStickers)
	x.Int(e.Hash)
	x.Vector(e.Stickers)
	return x.buf
}

func (e TL_messages_archivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_archivedStickers)
	x.Int(e.Count)
	x.Vector(e.Sets)
	return x.buf
}

func (e TL_messages_stickerSetInstallResultSuccess) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSetInstallResultSuccess)
	return x.buf
}

func (e TL_messages_stickerSetInstallResultArchive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSetInstallResultArchive)
	x.Vector(e.Sets)
	return x.buf
}

func (e TL_stickerSetCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetCovered)
	x.Bytes(e.Set.encode())
	x.Bytes(e.Cover.encode())
	return x.buf
}

func (e TL_inputMediaPhotoExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhotoExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.buf
}

func (e TL_inputMediaDocumentExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocumentExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	x.String(e.Caption)
	x.Int(e.Ttl_seconds)
	return x.buf
}

func (e TL_updateConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateConfig)
	return x.buf
}

func (e TL_updatePtsChanged) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePtsChanged)
	return x.buf
}

func (e TL_messageActionGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionGameScore)
	x.Long(e.Game_id)
	x.Int(e.Score)
	return x.buf
}

func (e TL_documentAttributeHasStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeHasStickers)
	return x.buf
}

func (e TL_keyboardButtonGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonGame)
	x.String(e.Text)
	return x.buf
}

func (e TL_stickerSetMultiCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetMultiCovered)
	x.Bytes(e.Set.encode())
	x.Vector(e.Covers)
	return x.buf
}

func (e TL_maskCoords) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_maskCoords)
	x.Int(e.N)
	x.Double(e.X)
	x.Double(e.Y)
	x.Double(e.Zoom)
	return x.buf
}

func (e TL_inputStickeredMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickeredMediaPhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_inputStickeredMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickeredMediaDocument)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_inputMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGame)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_messageMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGame)
	x.Bytes(e.Game.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageGame)
	x.Int(e.Flags)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineResultGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultGame)
	x.String(e.Id)
	x.String(e.Short_name)
	x.Bytes(e.Send_message.encode())
	return x.buf
}

func (e TL_game) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_game)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.String(e.Short_name)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Document.encode())
	return x.buf
}

func (e TL_inputGameID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGameID)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_inputGameShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGameShortName)
	x.Bytes(e.Bot_id.encode())
	x.String(e.Short_name)
	return x.buf
}

func (e TL_highScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_highScore)
	x.Int(e.Pos)
	x.Int(e.User_id)
	x.Int(e.Score)
	return x.buf
}

func (e TL_messages_highScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_highScores)
	x.Vector(e.Scores)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_chatsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chatsSlice)
	x.Int(e.Count)
	x.Vector(e.Chats)
	return x.buf
}

func (e TL_updateChannelWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelWebPage)
	x.Int(e.Channel_id)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.Pts_count)
	return x.buf
}

func (e TL_updates_differenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceTooLong)
	x.Int(e.Pts)
	return x.buf
}

func (e TL_sendMessageGamePlayAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGamePlayAction)
	return x.buf
}

func (e TL_webPageNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageNotModified)
	return x.buf
}

func (e TL_textEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textEmpty)
	return x.buf
}

func (e TL_textPlain) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textPlain)
	x.String(e.Text)
	return x.buf
}

func (e TL_textBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textBold)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textItalic)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textUnderline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textUnderline)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textStrike) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textStrike)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textFixed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textFixed)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textUrl)
	x.Bytes(e.Text.encode())
	x.String(e.Url)
	x.Long(e.Webpage_id)
	return x.buf
}

func (e TL_textEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textEmail)
	x.Bytes(e.Text.encode())
	x.String(e.Email)
	return x.buf
}

func (e TL_textConcat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_textConcat)
	x.Vector(e.Texts)
	return x.buf
}

func (e TL_pageBlockTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockTitle)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockSubtitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSubtitle)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockAuthorDate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAuthorDate)
	x.Bytes(e.Author.encode())
	x.Int(e.Published_date)
	return x.buf
}

func (e TL_pageBlockHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockHeader)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockSubheader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSubheader)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockParagraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockParagraph)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockPreformatted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPreformatted)
	x.Bytes(e.Text.encode())
	x.String(e.Language)
	return x.buf
}

func (e TL_pageBlockFooter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockFooter)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockDivider) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockDivider)
	return x.buf
}

func (e TL_pageBlockList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockList)
	x.Bytes(e.Ordered.encode())
	x.Vector(e.Items)
	return x.buf
}

func (e TL_pageBlockBlockquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockBlockquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockPullquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPullquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockPhoto)
	x.Long(e.Photo_id)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockVideo)
	x.Int(e.Flags)
	x.Long(e.Video_id)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockCover) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockCover)
	x.Bytes(e.Cover.encode())
	return x.buf
}

func (e TL_pageBlockEmbed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockEmbed)
	x.Int(e.Flags)
	x.String(e.Url)
	x.String(e.Html)
	x.Long(e.Poster_photo_id)
	x.Int(e.W)
	x.Int(e.H)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockEmbedPost) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockEmbedPost)
	x.String(e.Url)
	x.Long(e.Webpage_id)
	x.Long(e.Author_photo_id)
	x.String(e.Author)
	x.Int(e.Date)
	x.Vector(e.Blocks)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockSlideshow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockSlideshow)
	x.Vector(e.Items)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pagePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pagePart)
	x.Vector(e.Blocks)
	x.Vector(e.Photos)
	x.Vector(e.Documents)
	return x.buf
}

func (e TL_pageFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageFull)
	x.Vector(e.Blocks)
	x.Vector(e.Photos)
	x.Vector(e.Documents)
	return x.buf
}

func (e TL_updatePhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePhoneCall)
	x.Bytes(e.Phone_call.encode())
	return x.buf
}

func (e TL_updateDialogPinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDialogPinned)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_updatePinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePinnedDialogs)
	x.Int(e.Flags)
	x.Vector(e.Order)
	return x.buf
}

func (e TL_inputPrivacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyPhoneCall)
	return x.buf
}

func (e TL_privacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyPhoneCall)
	return x.buf
}

func (e TL_pageBlockUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockUnsupported)
	return x.buf
}

func (e TL_pageBlockAnchor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAnchor)
	x.String(e.Name)
	return x.buf
}

func (e TL_pageBlockCollage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockCollage)
	x.Vector(e.Items)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_inputPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneCall)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_phoneCallEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallEmpty)
	x.Long(e.Id)
	return x.buf
}

func (e TL_phoneCallWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallWaiting)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.Bytes(e.Protocol.encode())
	x.Int(e.Receive_date)
	return x.buf
}

func (e TL_phoneCallRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallRequested)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_a_hash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCall)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_a_or_b)
	x.Long(e.Key_fingerprint)
	x.Bytes(e.Protocol.encode())
	x.Bytes(e.Connection.encode())
	x.Vector(e.Alternative_connections)
	x.Int(e.Start_date)
	return x.buf
}

func (e TL_phoneCallDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscarded)
	x.Int(e.Flags)
	x.Long(e.Id)
	x.Bytes(e.Reason.encode())
	x.Int(e.Duration)
	return x.buf
}

func (e TL_phoneConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneConnection)
	x.Long(e.Id)
	x.String(e.Ip)
	x.String(e.Ipv6)
	x.Int(e.Port)
	x.StringBytes(e.Peer_tag)
	return x.buf
}

func (e TL_phoneCallProtocol) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallProtocol)
	x.Int(e.Flags)
	x.Int(e.Min_layer)
	x.Int(e.Max_layer)
	return x.buf
}

func (e TL_phone_phoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_phoneCall)
	x.Bytes(e.Phone_call.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_phoneCallDiscardReasonMissed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonMissed)
	return x.buf
}

func (e TL_phoneCallDiscardReasonDisconnect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonDisconnect)
	return x.buf
}

func (e TL_phoneCallDiscardReasonHangup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonHangup)
	return x.buf
}

func (e TL_phoneCallDiscardReasonBusy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallDiscardReasonBusy)
	return x.buf
}

func (e TL_inputMessagesFilterPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhoneCalls)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_messageActionPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPhoneCall)
	x.Int(e.Flags)
	x.Long(e.Call_id)
	x.Bytes(e.Reason.encode())
	x.Int(e.Duration)
	return x.buf
}

func (e TL_invoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invoice)
	x.Int(e.Flags)
	x.String(e.Currency)
	x.Vector(e.Prices)
	return x.buf
}

func (e TL_inputMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaInvoice)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Invoice.encode())
	x.StringBytes(e.Payload)
	x.String(e.Provider)
	x.String(e.Start_param)
	return x.buf
}

func (e TL_messageActionPaymentSentMe) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPaymentSentMe)
	x.Int(e.Flags)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.StringBytes(e.Payload)
	x.Bytes(e.Info.encode())
	x.String(e.Shipping_option_id)
	x.Bytes(e.Charge.encode())
	return x.buf
}

func (e TL_messageMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaInvoice)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.encode())
	x.Int(e.Receipt_msg_id)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.String(e.Start_param)
	return x.buf
}

func (e TL_keyboardButtonBuy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonBuy)
	x.String(e.Text)
	return x.buf
}

func (e TL_messageActionPaymentSent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPaymentSent)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	return x.buf
}

func (e TL_payments_paymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentForm)
	x.Int(e.Flags)
	x.Int(e.Bot_id)
	x.Bytes(e.Invoice.encode())
	x.Int(e.Provider_id)
	x.String(e.Url)
	x.String(e.Native_provider)
	x.Bytes(e.Native_params.encode())
	x.Bytes(e.Saved_info.encode())
	x.Bytes(e.Saved_credentials.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_postAddress) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_postAddress)
	x.String(e.Street_line1)
	x.String(e.Street_line2)
	x.String(e.City)
	x.String(e.State)
	x.String(e.Country_iso2)
	x.String(e.Post_code)
	return x.buf
}

func (e TL_paymentRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentRequestedInfo)
	x.Int(e.Flags)
	x.String(e.Name)
	x.String(e.Phone)
	x.String(e.Email)
	x.Bytes(e.Shipping_address.encode())
	return x.buf
}

func (e TL_updateBotWebhookJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotWebhookJSON)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e TL_updateBotWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotWebhookJSONQuery)
	x.Long(e.Query_id)
	x.Bytes(e.Data.encode())
	x.Int(e.Timeout)
	return x.buf
}

func (e TL_updateBotShippingQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotShippingQuery)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.StringBytes(e.Payload)
	x.Bytes(e.Shipping_address.encode())
	return x.buf
}

func (e TL_updateBotPrecheckoutQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotPrecheckoutQuery)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Int(e.User_id)
	x.StringBytes(e.Payload)
	x.Bytes(e.Info.encode())
	x.String(e.Shipping_option_id)
	x.String(e.Currency)
	x.Long(e.Total_amount)
	return x.buf
}

func (e TL_dataJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dataJSON)
	x.String(e.Data)
	return x.buf
}

func (e TL_labeledPrice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_labeledPrice)
	x.String(e.Label)
	x.Long(e.Amount)
	return x.buf
}

func (e TL_paymentCharge) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentCharge)
	x.String(e.Id)
	x.String(e.Provider_charge_id)
	return x.buf
}

func (e TL_paymentSavedCredentialsCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_paymentSavedCredentialsCard)
	x.String(e.Id)
	x.String(e.Title)
	return x.buf
}

func (e TL_webDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webDocument)
	x.String(e.Url)
	x.Long(e.Access_hash)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	x.Int(e.Dc_id)
	return x.buf
}

func (e TL_inputWebDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebDocument)
	x.String(e.Url)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TL_inputWebFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputWebFileLocation)
	x.String(e.Url)
	x.Long(e.Access_hash)
	return x.buf
}

func (e TL_upload_webFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_webFile)
	x.Int(e.Size)
	x.String(e.Mime_type)
	x.Bytes(e.File_type.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_payments_validatedRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_validatedRequestedInfo)
	x.Int(e.Flags)
	x.String(e.Id)
	x.Vector(e.Shipping_options)
	return x.buf
}

func (e TL_payments_paymentResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentResult)
	x.Bytes(e.Updates.encode())
	return x.buf
}

func (e TL_payments_paymentVerficationNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentVerficationNeeded)
	x.String(e.Url)
	return x.buf
}

func (e TL_payments_paymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_paymentReceipt)
	x.Int(e.Flags)
	x.Int(e.Date)
	x.Int(e.Bot_id)
	x.Bytes(e.Invoice.encode())
	x.Int(e.Provider_id)
	x.Bytes(e.Info.encode())
	x.Bytes(e.Shipping.encode())
	x.String(e.Currency)
	x.Long(e.Total_amount)
	x.String(e.Credentials_title)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_payments_savedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_savedInfo)
	x.Int(e.Flags)
	x.Bytes(e.Saved_info.encode())
	return x.buf
}

func (e TL_inputPaymentCredentialsSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentialsSaved)
	x.String(e.Id)
	x.StringBytes(e.Tmp_password)
	return x.buf
}

func (e TL_inputPaymentCredentials) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPaymentCredentials)
	x.Int(e.Flags)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e TL_account_tmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_tmpPassword)
	x.StringBytes(e.Tmp_password)
	x.Int(e.Valid_until)
	return x.buf
}

func (e TL_shippingOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_shippingOption)
	x.String(e.Id)
	x.String(e.Title)
	x.Vector(e.Prices)
	return x.buf
}

func (e TL_phoneCallAccepted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phoneCallAccepted)
	x.Long(e.Id)
	x.Long(e.Access_hash)
	x.Int(e.Date)
	x.Int(e.Admin_id)
	x.Int(e.Participant_id)
	x.StringBytes(e.G_b)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_inputMessagesFilterRoundVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterRoundVoice)
	return x.buf
}

func (e TL_inputMessagesFilterRoundVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterRoundVideo)
	return x.buf
}

func (e TL_upload_fileCdnRedirect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_fileCdnRedirect)
	x.Int(e.Dc_id)
	x.StringBytes(e.File_token)
	x.StringBytes(e.Encryption_key)
	x.StringBytes(e.Encryption_iv)
	x.Vector(e.Cdn_file_hashes)
	return x.buf
}

func (e TL_sendMessageRecordRoundAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordRoundAction)
	return x.buf
}

func (e TL_sendMessageUploadRoundAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadRoundAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_upload_cdnFileReuploadNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_cdnFileReuploadNeeded)
	x.StringBytes(e.Request_token)
	return x.buf
}

func (e TL_upload_cdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_cdnFile)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_cdnPublicKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_cdnPublicKey)
	x.Int(e.Dc_id)
	x.String(e.Public_key)
	return x.buf
}

func (e TL_cdnConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_cdnConfig)
	x.Vector(e.Public_keys)
	return x.buf
}

func (e TL_updateLangPackTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateLangPackTooLong)
	return x.buf
}

func (e TL_updateLangPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateLangPack)
	x.Bytes(e.Difference.encode())
	return x.buf
}

func (e TL_pageBlockChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_inputStickerSetItem) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetItem)
	x.Int(e.Flags)
	x.Bytes(e.Document.encode())
	x.String(e.Emoji)
	x.Bytes(e.Mask_coords.encode())
	return x.buf
}

func (e TL_langPackString) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackString)
	x.String(e.Key)
	x.String(e.Value)
	return x.buf
}

func (e TL_langPackStringPluralized) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackStringPluralized)
	x.Int(e.Flags)
	x.String(e.Key)
	x.String(e.Zero_value)
	x.String(e.One_value)
	x.String(e.Two_value)
	x.String(e.Few_value)
	x.String(e.Many_value)
	x.String(e.Other_value)
	return x.buf
}

func (e TL_langPackStringDeleted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackStringDeleted)
	x.String(e.Key)
	return x.buf
}

func (e TL_langPackDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackDifference)
	x.String(e.Lang_code)
	x.Int(e.From_version)
	x.Int(e.Version)
	x.Vector(e.Strings)
	return x.buf
}

func (e TL_langPackLanguage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langPackLanguage)
	x.String(e.Name)
	x.String(e.Native_name)
	x.String(e.Lang_code)
	return x.buf
}

func (e TL_channelParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantAdmin)
	x.Int(e.Flags)
	x.Int(e.User_id)
	x.Int(e.Inviter_id)
	x.Int(e.Promoted_by)
	x.Int(e.Date)
	x.Bytes(e.Admin_rights.encode())
	return x.buf
}

func (e TL_channelParticipantBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantBanned)
	x.Int(e.Flags)
	x.Int(e.User_id)
	x.Int(e.Kicked_by)
	x.Int(e.Date)
	x.Bytes(e.Banned_rights.encode())
	return x.buf
}

func (e TL_channelParticipantsBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsBanned)
	x.String(e.Q)
	return x.buf
}

func (e TL_channelParticipantsSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsSearch)
	x.String(e.Q)
	return x.buf
}

func (e TL_topPeerCategoryPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryPhoneCalls)
	return x.buf
}

func (e TL_pageBlockAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_pageBlockAudio)
	x.Long(e.Audio_id)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_channelAdminRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminRights)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_channelBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelBannedRights)
	x.Int(e.Flags)
	x.Int(e.Until_date)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeTitle)
	x.String(e.Prev_value)
	x.String(e.New_value)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeAbout)
	x.String(e.Prev_value)
	x.String(e.New_value)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeUsername)
	x.String(e.Prev_value)
	x.String(e.New_value)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangePhoto)
	x.Bytes(e.Prev_photo.encode())
	x.Bytes(e.New_photo.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionToggleInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionToggleInvites)
	x.Bytes(e.New_value.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionToggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionToggleSignatures)
	x.Bytes(e.New_value.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionUpdatePinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionUpdatePinned)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionEditMessage)
	x.Bytes(e.Prev_message.encode())
	x.Bytes(e.New_message.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionDeleteMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionDeleteMessage)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantJoin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantJoin)
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantLeave) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantLeave)
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantInvite)
	x.Bytes(e.Participant.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantToggleBan) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantToggleBan)
	x.Bytes(e.Prev_participant.encode())
	x.Bytes(e.New_participant.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantToggleAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionParticipantToggleAdmin)
	x.Bytes(e.Prev_participant.encode())
	x.Bytes(e.New_participant.encode())
	return x.buf
}

func (e TL_channelAdminLogEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEvent)
	x.Long(e.Id)
	x.Int(e.Date)
	x.Int(e.User_id)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_channels_adminLogResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_adminLogResults)
	x.Vector(e.Events)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_channelAdminLogEventsFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventsFilter)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_messageActionScreenshotTaken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionScreenshotTaken)
	return x.buf
}

func (e TL_popularContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_popularContact)
	x.Long(e.Client_id)
	x.Int(e.Importers)
	return x.buf
}

func (e TL_cdnFileHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_cdnFileHash)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.StringBytes(e.Hash)
	return x.buf
}

func (e TL_inputMessagesFilterMyMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMyMentions)
	return x.buf
}

func (e TL_inputMessagesFilterMyMentionsUnread) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMyMentionsUnread)
	return x.buf
}

func (e TL_updateContactsReset) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactsReset)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelAdminLogEventActionChangeStickerSet)
	x.Bytes(e.Prev_stickerset.encode())
	x.Bytes(e.New_stickerset.encode())
	return x.buf
}

func (e TL_updateFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateFavedStickers)
	return x.buf
}

func (e TL_messages_favedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_favedStickers)
	x.Int(e.Hash)
	x.Vector(e.Packs)
	x.Vector(e.Stickers)
	return x.buf
}

func (e TL_messages_favedStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_favedStickersNotModified)
	return x.buf
}

func (e TL_updateChannelReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelReadMessagesContents)
	x.Int(e.Channel_id)
	x.VectorInt(e.Messages)
	return x.buf
}

func (e TL_invokeAfterMsg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsg)
	x.Long(e.Msg_id)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_invokeAfterMsgs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsgs)
	x.VectorLong(e.Msg_ids)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_auth_checkPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkPhone)
	x.String(e.Phone_number)
	return x.buf
}

func (e TL_auth_sendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendCode)
	x.Int(e.Flags)
	x.String(e.Phone_number)
	x.Bytes(e.Current_number.encode())
	x.Int(e.Api_id)
	x.String(e.Api_hash)
	return x.buf
}

func (e TL_auth_signUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signUp)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	x.String(e.First_name)
	x.String(e.Last_name)
	return x.buf
}

func (e TL_auth_signIn) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signIn)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	return x.buf
}

func (e TL_auth_logOut) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_logOut)
	return x.buf
}

func (e TL_auth_resetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resetAuthorizations)
	return x.buf
}

func (e TL_auth_sendInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendInvites)
	x.VectorString(e.Phone_numbers)
	x.String(e.Message)
	return x.buf
}

func (e TL_auth_exportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportAuthorization)
	x.Int(e.Dc_id)
	return x.buf
}

func (e TL_auth_importAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importAuthorization)
	x.Int(e.Id)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_account_registerDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_registerDevice)
	x.Int(e.Token_type)
	x.String(e.Token)
	return x.buf
}

func (e TL_account_unregisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_unregisterDevice)
	x.Int(e.Token_type)
	x.String(e.Token)
	return x.buf
}

func (e TL_account_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_getNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getNotifySettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_account_resetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetNotifySettings)
	return x.buf
}

func (e TL_account_updateProfile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateProfile)
	x.Int(e.Flags)
	x.String(e.First_name)
	x.String(e.Last_name)
	x.String(e.About)
	return x.buf
}

func (e TL_account_updateStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateStatus)
	x.Bytes(e.Offline.encode())
	return x.buf
}

func (e TL_account_getWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getWallPapers)
	return x.buf
}

func (e TL_users_getUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getUsers)
	x.Vector(e.Id)
	return x.buf
}

func (e TL_users_getFullUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getFullUser)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_contacts_getStatuses) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getStatuses)
	return x.buf
}

func (e TL_contacts_getContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getContacts)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_contacts_importContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importContacts)
	x.Vector(e.Contacts)
	return x.buf
}

func (e TL_contacts_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_search)
	x.String(e.Q)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_contacts_deleteContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContact)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_contacts_deleteContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContacts)
	x.Vector(e.Id)
	return x.buf
}

func (e TL_contacts_block) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_block)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_contacts_unblock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_unblock)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_contacts_getBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getBlocked)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_messages_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessages)
	x.VectorInt(e.Id)
	return x.buf
}

func (e TL_messages_getDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDialogs)
	x.Int(e.Flags)
	x.Int(e.Offset_date)
	x.Int(e.Offset_id)
	x.Bytes(e.Offset_peer.encode())
	x.Int(e.Limit)
	return x.buf
}

func (e TL_messages_getHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Offset_id)
	x.Int(e.Offset_date)
	x.Int(e.Add_offset)
	x.Int(e.Limit)
	x.Int(e.Max_id)
	x.Int(e.Min_id)
	return x.buf
}

func (e TL_messages_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_search)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.String(e.Q)
	x.Bytes(e.From_id.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Min_date)
	x.Int(e.Max_date)
	x.Int(e.Offset_id)
	x.Int(e.Add_offset)
	x.Int(e.Limit)
	x.Int(e.Max_id)
	x.Int(e.Min_id)
	return x.buf
}

func (e TL_messages_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_id)
	return x.buf
}

func (e TL_messages_deleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteHistory)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_id)
	return x.buf
}

func (e TL_messages_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteMessages)
	x.Int(e.Flags)
	x.VectorInt(e.Id)
	return x.buf
}

func (e TL_messages_receivedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedMessages)
	x.Int(e.Max_id)
	return x.buf
}

func (e TL_messages_setTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_messages_sendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMessage)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	if e.Reply_to_msg_id != 0 {
		x.Int(e.Reply_to_msg_id)
	}
	x.String(e.Message)
	x.Long(e.Random_id)
	if e.Reply_markup != nil {
		x.Bytes(e.Reply_markup.encode())
	}
	if len(e.Entities) > 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_messages_sendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMedia)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Reply_to_msg_id)
	x.Bytes(e.Media.encode())
	x.Long(e.Random_id)
	x.Bytes(e.Reply_markup.encode())
	return x.buf
}

func (e TL_messages_forwardMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessages)
	x.Int(e.Flags)
	x.Bytes(e.From_peer.encode())
	x.VectorInt(e.Id)
	x.VectorLong(e.Random_id)
	x.Bytes(e.To_peer.encode())
	return x.buf
}

func (e TL_messages_getChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getChats)
	x.VectorInt(e.Id)
	return x.buf
}

func (e TL_messages_getFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFullChat)
	x.Int(e.Chat_id)
	return x.buf
}

func (e TL_messages_editChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatTitle)
	x.Int(e.Chat_id)
	x.String(e.Title)
	return x.buf
}

func (e TL_messages_editChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatPhoto)
	x.Int(e.Chat_id)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TL_messages_addChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_addChatUser)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.encode())
	x.Int(e.Fwd_limit)
	return x.buf
}

func (e TL_messages_deleteChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteChatUser)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.encode())
	return x.buf
}

func (e TL_messages_createChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_createChat)
	x.Vector(e.Users)
	x.String(e.Title)
	return x.buf
}

func (e TL_updates_getState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getState)
	return x.buf
}

func (e TL_updates_getDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getDifference)
	x.Int(e.Flags)
	x.Int(e.Pts)
	x.Int(e.Pts_total_limit)
	x.Int(e.Date)
	x.Int(e.Qts)
	return x.buf
}

func (e TL_photos_updateProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_updateProfilePhoto)
	x.Bytes(e.Id.encode())
	return x.buf
}

func (e TL_photos_uploadProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_uploadProfilePhoto)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_upload_saveFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveFilePart)
	x.Long(e.File_id)
	x.Int(e.File_part)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_upload_getFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_help_getConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getConfig)
	return x.buf
}

func (e TL_help_getNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getNearestDc)
	return x.buf
}

func (e TL_help_getAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppUpdate)
	return x.buf
}

func (e TL_help_saveAppLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_saveAppLog)
	x.Vector(e.Events)
	return x.buf
}

func (e TL_help_getInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getInviteText)
	return x.buf
}

func (e TL_photos_deletePhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_deletePhotos)
	x.Vector(e.Id)
	return x.buf
}

func (e TL_photos_getUserPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_getUserPhotos)
	x.Bytes(e.User_id.encode())
	x.Int(e.Offset)
	x.Long(e.Max_id)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_messages_forwardMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Long(e.Random_id)
	return x.buf
}

func (e TL_messages_getDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDhConfig)
	x.Int(e.Version)
	x.Int(e.Random_length)
	return x.buf
}

func (e TL_messages_requestEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_requestEncryption)
	x.Bytes(e.User_id.encode())
	x.Int(e.Random_id)
	x.StringBytes(e.G_a)
	return x.buf
}

func (e TL_messages_acceptEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_acceptEncryption)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.G_b)
	x.Long(e.Key_fingerprint)
	return x.buf
}

func (e TL_messages_discardEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_discardEncryption)
	x.Int(e.Chat_id)
	return x.buf
}

func (e TL_messages_setEncryptedTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setEncryptedTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Typing.encode())
	return x.buf
}

func (e TL_messages_readEncryptedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readEncryptedHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Max_date)
	return x.buf
}

func (e TL_messages_sendEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncrypted)
	x.Bytes(e.Peer.encode())
	x.Long(e.Random_id)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_messages_sendEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedFile)
	x.Bytes(e.Peer.encode())
	x.Long(e.Random_id)
	x.StringBytes(e.Data)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_messages_sendEncryptedService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedService)
	x.Bytes(e.Peer.encode())
	x.Long(e.Random_id)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_messages_receivedQueue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedQueue)
	x.Int(e.Max_qts)
	return x.buf
}

func (e TL_upload_saveBigFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveBigFilePart)
	x.Long(e.File_id)
	x.Int(e.File_part)
	x.Int(e.File_total_parts)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_initConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_initConnection)
	x.Int(e.Api_id)
	x.String(e.Device_model)
	x.String(e.System_version)
	x.String(e.App_version)
	x.String(e.System_lang_code)
	x.String(e.Lang_pack)
	x.String(e.Lang_code)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_help_getSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getSupport)
	return x.buf
}

func (e TL_auth_bindTempAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_bindTempAuthKey)
	x.Long(e.Perm_auth_key_id)
	x.Long(e.Nonce)
	x.Int(e.Expires_at)
	x.StringBytes(e.Encrypted_message)
	return x.buf
}

func (e TL_contacts_exportCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_exportCard)
	return x.buf
}

func (e TL_contacts_importCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importCard)
	x.VectorInt(e.Export_card)
	return x.buf
}

func (e TL_messages_readMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readMessageContents)
	x.VectorInt(e.Id)
	return x.buf
}

func (e TL_account_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_checkUsername)
	x.String(e.Username)
	return x.buf
}

func (e TL_account_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateUsername)
	x.String(e.Username)
	return x.buf
}

func (e TL_account_getPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPrivacy)
	x.Bytes(e.Key.encode())
	return x.buf
}

func (e TL_account_setPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setPrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

func (e TL_account_deleteAccount) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_deleteAccount)
	x.String(e.Reason)
	return x.buf
}

func (e TL_account_getAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAccountTTL)
	return x.buf
}

func (e TL_account_setAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setAccountTTL)
	x.Bytes(e.Ttl.encode())
	return x.buf
}

func (e TL_invokeWithLayer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithLayer)
	x.Int(e.Layer)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_contacts_resolveUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolveUsername)
	x.String(e.Username)
	return x.buf
}

func (e TL_account_sendChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendChangePhoneCode)
	x.Int(e.Flags)
	x.String(e.Phone_number)
	x.Bytes(e.Current_number.encode())
	return x.buf
}

func (e TL_account_changePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_changePhone)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	return x.buf
}

func (e TL_messages_getAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_account_updateDeviceLocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateDeviceLocked)
	x.Int(e.Period)
	return x.buf
}

func (e TL_account_getPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPassword)
	return x.buf
}

func (e TL_auth_checkPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkPassword)
	x.StringBytes(e.Password_hash)
	return x.buf
}

func (e TL_messages_getWebPagePreview) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getWebPagePreview)
	x.String(e.Message)
	return x.buf
}

func (e TL_account_getAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAuthorizations)
	return x.buf
}

func (e TL_account_resetAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetAuthorization)
	x.Long(e.Hash)
	return x.buf
}

func (e TL_account_getPasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPasswordSettings)
	x.StringBytes(e.Current_password_hash)
	return x.buf
}

func (e TL_account_updatePasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updatePasswordSettings)
	x.StringBytes(e.Current_password_hash)
	x.Bytes(e.New_settings.encode())
	return x.buf
}

func (e TL_auth_requestPasswordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_requestPasswordRecovery)
	return x.buf
}

func (e TL_auth_recoverPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_recoverPassword)
	x.String(e.Code)
	return x.buf
}

func (e TL_invokeWithoutUpdates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithoutUpdates)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_messages_exportChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_exportChatInvite)
	x.Int(e.Chat_id)
	return x.buf
}

func (e TL_messages_checkChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_checkChatInvite)
	x.String(e.Hash)
	return x.buf
}

func (e TL_messages_importChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_importChatInvite)
	x.String(e.Hash)
	return x.buf
}

func (e TL_messages_getStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_messages_installStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_installStickerSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Archived.encode())
	return x.buf
}

func (e TL_messages_uninstallStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_uninstallStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_auth_importBotAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importBotAuthorization)
	x.Int(e.Flags)
	x.Int(e.Api_id)
	x.String(e.Api_hash)
	x.String(e.Bot_auth_token)
	return x.buf
}

func (e TL_messages_startBot) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_startBot)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	x.Long(e.Random_id)
	x.String(e.Start_param)
	return x.buf
}

func (e TL_help_getAppChangelog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppChangelog)
	x.String(e.Prev_app_version)
	return x.buf
}

func (e TL_messages_reportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reportSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_getMessagesViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessagesViews)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.Id)
	x.Bytes(e.Increment.encode())
	return x.buf
}

func (e TL_updates_getChannelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getChannelDifference)
	x.Int(e.Flags)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Pts)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_channels_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_readHistory)
	x.Bytes(e.Channel.encode())
	x.Int(e.Max_id)
	return x.buf
}

func (e TL_channels_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteMessages)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.Id)
	return x.buf
}

func (e TL_channels_deleteUserHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteUserHistory)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	return x.buf
}

func (e TL_channels_reportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_reportSpam)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	x.VectorInt(e.Id)
	return x.buf
}

func (e TL_channels_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getMessages)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.Id)
	return x.buf
}

func (e TL_channels_getParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getParticipants)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_channels_getParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getParticipant)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	return x.buf
}

func (e TL_channels_getChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getChannels)
	x.Vector(e.Id)
	return x.buf
}

func (e TL_channels_getFullChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getFullChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_channels_createChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_createChannel)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.About)
	return x.buf
}

func (e TL_channels_editAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editAbout)
	x.Bytes(e.Channel.encode())
	x.String(e.About)
	return x.buf
}

func (e TL_channels_editAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editAdmin)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	x.Bytes(e.Admin_rights.encode())
	return x.buf
}

func (e TL_channels_editTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editTitle)
	x.Bytes(e.Channel.encode())
	x.String(e.Title)
	return x.buf
}

func (e TL_channels_editPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editPhoto)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TL_channels_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_checkUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}

func (e TL_channels_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_updateUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}

func (e TL_channels_joinChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_joinChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_channels_leaveChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_leaveChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_channels_inviteToChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_inviteToChannel)
	x.Bytes(e.Channel.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_channels_exportInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_exportInvite)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_channels_deleteChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_messages_toggleChatAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_toggleChatAdmins)
	x.Int(e.Chat_id)
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e TL_messages_editChatAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatAdmin)
	x.Int(e.Chat_id)
	x.Bytes(e.User_id.encode())
	x.Bytes(e.Is_admin.encode())
	return x.buf
}

func (e TL_messages_migrateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_migrateChat)
	x.Int(e.Chat_id)
	return x.buf
}

func (e TL_messages_searchGlobal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchGlobal)
	x.String(e.Q)
	x.Int(e.Offset_date)
	x.Bytes(e.Offset_peer.encode())
	x.Int(e.Offset_id)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_account_reportPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_reportPeer)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Reason.encode())
	return x.buf
}

func (e TL_messages_reorderStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reorderStickerSets)
	x.Int(e.Flags)
	x.VectorLong(e.Order)
	return x.buf
}

func (e TL_help_getTermsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getTermsOfService)
	return x.buf
}

func (e TL_messages_getDocumentByHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDocumentByHash)
	x.StringBytes(e.Sha256)
	x.Int(e.Size)
	x.String(e.Mime_type)
	return x.buf
}

func (e TL_messages_searchGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchGifs)
	x.String(e.Q)
	x.Int(e.Offset)
	return x.buf
}

func (e TL_messages_getSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getSavedGifs)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_saveGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveGif)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

func (e TL_messages_getInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getInlineBotResults)
	x.Int(e.Flags)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Geo_point.encode())
	x.String(e.Query)
	x.String(e.Offset)
	return x.buf
}

func (e TL_messages_setInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setInlineBotResults)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.Vector(e.Results)
	x.Int(e.Cache_time)
	x.String(e.Next_offset)
	x.Bytes(e.Switch_pm.encode())
	return x.buf
}

func (e TL_messages_sendInlineBotResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendInlineBotResult)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Reply_to_msg_id)
	x.Long(e.Random_id)
	x.Long(e.Query_id)
	x.String(e.Id)
	return x.buf
}

func (e TL_channels_toggleInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_toggleInvites)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e TL_channels_exportMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_exportMessageLink)
	x.Bytes(e.Channel.encode())
	x.Int(e.Id)
	return x.buf
}

func (e TL_channels_toggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_toggleSignatures)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e TL_messages_hideReportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_hideReportSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_getPeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPeerSettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_channels_updatePinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_updatePinnedMessage)
	x.Int(e.Flags)
	x.Bytes(e.Channel.encode())
	x.Int(e.Id)
	return x.buf
}

func (e TL_auth_resendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resendCode)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	return x.buf
}

func (e TL_auth_cancelCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_cancelCode)
	x.String(e.Phone_number)
	x.String(e.Phone_code_hash)
	return x.buf
}

func (e TL_messages_getMessageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessageEditData)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	return x.buf
}

func (e TL_messages_editMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editMessage)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.String(e.Message)
	x.Bytes(e.Reply_markup.encode())
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_messages_editInlineBotMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editInlineBotMessage)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.String(e.Message)
	x.Bytes(e.Reply_markup.encode())
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_messages_getBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getBotCallbackAnswer)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Msg_id)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_messages_setBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotCallbackAnswer)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.String(e.Message)
	x.String(e.Url)
	x.Int(e.Cache_time)
	return x.buf
}

func (e TL_contacts_getTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getTopPeers)
	x.Int(e.Flags)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_contacts_resetTopPeerRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resetTopPeerRating)
	x.Bytes(e.Category.encode())
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_getPeerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPeerDialogs)
	x.Vector(e.Peers)
	return x.buf
}

func (e TL_messages_saveDraft) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveDraft)
	x.Int(e.Flags)
	x.Int(e.Reply_to_msg_id)
	x.Bytes(e.Peer.encode())
	x.String(e.Message)
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_messages_getAllDrafts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllDrafts)
	return x.buf
}

func (e TL_account_sendConfirmPhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendConfirmPhoneCode)
	x.Int(e.Flags)
	x.String(e.Hash)
	x.Bytes(e.Current_number.encode())
	return x.buf
}

func (e TL_account_confirmPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_confirmPhone)
	x.String(e.Phone_code_hash)
	x.String(e.Phone_code)
	return x.buf
}

func (e TL_messages_getFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFeaturedStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_readFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readFeaturedStickers)
	x.VectorLong(e.Id)
	return x.buf
}

func (e TL_messages_getRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getRecentStickers)
	x.Int(e.Flags)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_saveRecentSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveRecentSticker)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

func (e TL_messages_clearRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_clearRecentStickers)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_messages_getArchivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getArchivedStickers)
	x.Int(e.Flags)
	x.Long(e.Offset_id)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_channels_getAdminedPublicChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getAdminedPublicChannels)
	return x.buf
}

func (e TL_auth_dropTempAuthKeys) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_dropTempAuthKeys)
	x.VectorLong(e.Except_auth_keys)
	return x.buf
}

func (e TL_messages_setGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setGameScore)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Bytes(e.User_id.encode())
	x.Int(e.Score)
	return x.buf
}

func (e TL_messages_setInlineGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setInlineGameScore)
	x.Int(e.Flags)
	x.Bytes(e.Id.encode())
	x.Bytes(e.User_id.encode())
	x.Int(e.Score)
	return x.buf
}

func (e TL_messages_getMaskStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMaskStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_getAttachedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAttachedStickers)
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e TL_messages_getGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getGameHighScores)
	x.Bytes(e.Peer.encode())
	x.Int(e.Id)
	x.Bytes(e.User_id.encode())
	return x.buf
}

func (e TL_messages_getInlineGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getInlineGameHighScores)
	x.Bytes(e.Id.encode())
	x.Bytes(e.User_id.encode())
	return x.buf
}

func (e TL_messages_getCommonChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getCommonChats)
	x.Bytes(e.User_id.encode())
	x.Int(e.Max_id)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_messages_getAllChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllChats)
	x.VectorInt(e.Except_ids)
	return x.buf
}

func (e TL_help_setBotUpdatesStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_setBotUpdatesStatus)
	x.Int(e.Pending_updates_count)
	x.String(e.Message)
	return x.buf
}

func (e TL_messages_getWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getWebPage)
	x.String(e.Url)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_toggleDialogPin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_toggleDialogPin)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_reorderPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reorderPinnedDialogs)
	x.Int(e.Flags)
	x.Vector(e.Order)
	return x.buf
}

func (e TL_messages_getPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPinnedDialogs)
	return x.buf
}

func (e TL_phone_requestCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_requestCall)
	x.Bytes(e.User_id.encode())
	x.Int(e.Random_id)
	x.StringBytes(e.G_a_hash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phone_acceptCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_acceptCall)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.G_b)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phone_discardCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_discardCall)
	x.Bytes(e.Peer.encode())
	x.Int(e.Duration)
	x.Bytes(e.Reason.encode())
	x.Long(e.Connection_id)
	return x.buf
}

func (e TL_phone_receivedCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_receivedCall)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_reportEncryptedSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reportEncryptedSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_payments_getPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getPaymentForm)
	x.Int(e.Msg_id)
	return x.buf
}

func (e TL_payments_sendPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_sendPaymentForm)
	x.Int(e.Flags)
	x.Int(e.Msg_id)
	x.String(e.Requested_info_id)
	x.String(e.Shipping_option_id)
	x.Bytes(e.Credentials.encode())
	return x.buf
}

func (e TL_account_getTmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getTmpPassword)
	x.StringBytes(e.Password_hash)
	x.Int(e.Period)
	return x.buf
}

func (e TL_messages_setBotShippingResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotShippingResults)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.String(e.Error)
	x.Vector(e.Shipping_options)
	return x.buf
}

func (e TL_messages_setBotPrecheckoutResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotPrecheckoutResults)
	x.Int(e.Flags)
	x.Long(e.Query_id)
	x.String(e.Error)
	return x.buf
}

func (e TL_upload_getWebFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getWebFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_bots_sendCustomRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_bots_sendCustomRequest)
	x.String(e.Custom_method)
	x.Bytes(e.Params.encode())
	return x.buf
}

func (e TL_bots_answerWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_bots_answerWebhookJSONQuery)
	x.Long(e.Query_id)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e TL_payments_getPaymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getPaymentReceipt)
	x.Int(e.Msg_id)
	return x.buf
}

func (e TL_payments_validateRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_validateRequestedInfo)
	x.Int(e.Flags)
	x.Int(e.Msg_id)
	x.Bytes(e.Info.encode())
	return x.buf
}

func (e TL_payments_getSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_getSavedInfo)
	return x.buf
}

func (e TL_payments_clearSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_payments_clearSavedInfo)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_phone_getCallConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_getCallConfig)
	return x.buf
}

func (e TL_phone_confirmCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_confirmCall)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.G_a)
	x.Long(e.Key_fingerprint)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phone_setCallRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_setCallRating)
	x.Bytes(e.Peer.encode())
	x.Int(e.Rating)
	x.String(e.Comment)
	return x.buf
}

func (e TL_phone_saveCallDebug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_phone_saveCallDebug)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Debug.encode())
	return x.buf
}

func (e TL_upload_getCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getCdnFile)
	x.StringBytes(e.File_token)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_upload_reuploadCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_reuploadCdnFile)
	x.StringBytes(e.File_token)
	x.StringBytes(e.Request_token)
	return x.buf
}

func (e TL_help_getCdnConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getCdnConfig)
	return x.buf
}

func (e TL_messages_uploadMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_uploadMedia)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e TL_stickers_createStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickers_createStickerSet)
	x.Int(e.Flags)
	x.Bytes(e.User_id.encode())
	x.String(e.Title)
	x.String(e.Short_name)
	x.Vector(e.Stickers)
	return x.buf
}

func (e TL_langpack_getLangPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getLangPack)
	x.String(e.Lang_code)
	return x.buf
}

func (e TL_langpack_getStrings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getStrings)
	x.String(e.Lang_code)
	x.VectorString(e.Keys)
	return x.buf
}

func (e TL_langpack_getDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getDifference)
	x.Int(e.From_version)
	return x.buf
}

func (e TL_langpack_getLanguages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_langpack_getLanguages)
	return x.buf
}

func (e TL_channels_editBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editBanned)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.User_id.encode())
	x.Bytes(e.Banned_rights.encode())
	return x.buf
}

func (e TL_channels_getAdminLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getAdminLog)
	x.Int(e.Flags)
	x.Bytes(e.Channel.encode())
	x.String(e.Q)
	x.Bytes(e.Events_filter.encode())
	x.Vector(e.Admins)
	x.Long(e.Max_id)
	x.Long(e.Min_id)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_stickers_removeStickerFromSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickers_removeStickerFromSet)
	x.Bytes(e.Sticker.encode())
	return x.buf
}

func (e TL_stickers_changeStickerPosition) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickers_changeStickerPosition)
	x.Bytes(e.Sticker.encode())
	x.Int(e.Position)
	return x.buf
}

func (e TL_stickers_addStickerToSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickers_addStickerToSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Sticker.encode())
	return x.buf
}

func (e TL_messages_sendScreenshotNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendScreenshotNotification)
	x.Bytes(e.Peer.encode())
	x.Int(e.Reply_to_msg_id)
	x.Long(e.Random_id)
	return x.buf
}

func (e TL_upload_getCdnFileHashes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getCdnFileHashes)
	x.StringBytes(e.File_token)
	x.Int(e.Offset)
	return x.buf
}

func (e TL_messages_getUnreadMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getUnreadMentions)
	x.Bytes(e.Peer.encode())
	x.Int(e.Offset_id)
	x.Int(e.Add_offset)
	x.Int(e.Limit)
	x.Int(e.Max_id)
	x.Int(e.Min_id)
	return x.buf
}

func (e TL_messages_faveSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_faveSticker)
	x.Bytes(e.Id.encode())
	x.Bytes(e.Unfave.encode())
	return x.buf
}

func (e TL_channels_setStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_setStickers)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_contacts_resetSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resetSaved)
	return x.buf
}

func (e TL_messages_getFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFavedStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_channels_readMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_readMessageContents)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.Id)
	return x.buf
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case crc_boolFalse:
		r = TL_boolFalse{}

	case crc_boolTrue:
		r = TL_boolTrue{}

	case crc_error:
		r = TL_error{
			m.Int(),
			m.String(),
		}

	case crc_null:
		r = TL_null{}

	case crc_vector:
		r = TL_vector{
			m.TL_Vector(),
		}

	case crc_inputPeerEmpty:
		r = TL_inputPeerEmpty{}

	case crc_inputPeerSelf:
		r = TL_inputPeerSelf{}

	case crc_inputPeerChat:
		r = TL_inputPeerChat{
			m.Int(),
		}

	case crc_inputUserEmpty:
		r = TL_inputUserEmpty{}

	case crc_inputUserSelf:
		r = TL_inputUserSelf{}

	case crc_inputPhoneContact:
		r = TL_inputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputFile:
		r = TL_inputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_inputMediaEmpty:
		r = TL_inputMediaEmpty{}

	case crc_inputMediaUploadedPhoto:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaUploadedPhoto{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedVector(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crc_inputMediaPhoto:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaPhoto{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_inputMediaGeoPoint:
		r = TL_inputMediaGeoPoint{
			m.Object(),
		}

	case crc_inputMediaContact:
		r = TL_inputMediaContact{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputChatPhotoEmpty:
		r = TL_inputChatPhotoEmpty{}

	case crc_inputChatUploadedPhoto:
		r = TL_inputChatUploadedPhoto{
			m.Object(),
		}

	case crc_inputChatPhoto:
		r = TL_inputChatPhoto{
			m.Object(),
		}

	case crc_inputGeoPointEmpty:
		r = TL_inputGeoPointEmpty{}

	case crc_inputGeoPoint:
		r = TL_inputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_inputPhotoEmpty:
		r = TL_inputPhotoEmpty{}

	case crc_inputPhoto:
		r = TL_inputPhoto{
			m.Long(),
			m.Long(),
		}

	case crc_inputFileLocation:
		r = TL_inputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_inputAppEvent:
		r = TL_inputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_peerUser:
		r = TL_peerUser{
			m.Int(),
		}

	case crc_peerChat:
		r = TL_peerChat{
			m.Int(),
		}

	case crc_storage_fileUnknown:
		r = TL_storage_fileUnknown{}

	case crc_storage_fileJpeg:
		r = TL_storage_fileJpeg{}

	case crc_storage_fileGif:
		r = TL_storage_fileGif{}

	case crc_storage_filePng:
		r = TL_storage_filePng{}

	case crc_storage_fileMp3:
		r = TL_storage_fileMp3{}

	case crc_storage_fileMov:
		r = TL_storage_fileMov{}

	case crc_storage_filePartial:
		r = TL_storage_filePartial{}

	case crc_storage_fileMp4:
		r = TL_storage_fileMp4{}

	case crc_storage_fileWebp:
		r = TL_storage_fileWebp{}

	case crc_fileLocationUnavailable:
		r = TL_fileLocationUnavailable{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_fileLocation:
		r = TL_fileLocation{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_userEmpty:
		r = TL_userEmpty{
			m.Int(),
		}

	case crc_userProfilePhotoEmpty:
		r = TL_userProfilePhotoEmpty{}

	case crc_userProfilePhoto:
		r = TL_userProfilePhoto{
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_userStatusEmpty:
		r = TL_userStatusEmpty{}

	case crc_userStatusOnline:
		r = TL_userStatusOnline{
			m.Int(),
		}

	case crc_userStatusOffline:
		r = TL_userStatusOffline{
			m.Int(),
		}

	case crc_chatEmpty:
		r = TL_chatEmpty{
			m.Int(),
		}

	case crc_chat:
		flags := m.Flags()
		_ = flags
		r = TL_chat{
			flags,
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 6),
		}

	case crc_chatForbidden:
		r = TL_chatForbidden{
			m.Int(),
			m.String(),
		}

	case crc_chatFull:
		r = TL_chatFull{
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case crc_chatParticipant:
		r = TL_chatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_chatParticipantsForbidden:
		flags := m.Flags()
		_ = flags
		r = TL_chatParticipantsForbidden{
			flags,
			m.Int(),
			m.FlaggedObject(flags, 0),
		}

	case crc_chatParticipants:
		r = TL_chatParticipants{
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case crc_chatPhotoEmpty:
		r = TL_chatPhotoEmpty{}

	case crc_chatPhoto:
		r = TL_chatPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_messageEmpty:
		r = TL_messageEmpty{
			m.Int(),
		}

	case crc_message:
		flags := m.Flags()
		_ = flags
		r = TL_message{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 8),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 9),
			m.FlaggedObject(flags, 6),
			m.FlaggedVector(flags, 7),
			m.FlaggedInt(flags, 10),
			m.FlaggedInt(flags, 15),
			m.FlaggedString(flags, 16),
		}

	case crc_messageService:
		flags := m.Flags()
		_ = flags
		r = TL_messageService{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 8),
			m.Object(),
			m.FlaggedInt(flags, 3),
			m.Int(),
			m.Object(),
		}

	case crc_messageMediaEmpty:
		r = TL_messageMediaEmpty{}

	case crc_messageMediaPhoto:
		flags := m.Flags()
		_ = flags
		r = TL_messageMediaPhoto{
			flags,
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crc_messageMediaGeo:
		r = TL_messageMediaGeo{
			m.Object(),
		}

	case crc_messageMediaContact:
		r = TL_messageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_messageMediaUnsupported:
		r = TL_messageMediaUnsupported{}

	case crc_messageActionEmpty:
		r = TL_messageActionEmpty{}

	case crc_messageActionChatCreate:
		r = TL_messageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case crc_messageActionChatEditTitle:
		r = TL_messageActionChatEditTitle{
			m.String(),
		}

	case crc_messageActionChatEditPhoto:
		r = TL_messageActionChatEditPhoto{
			m.Object(),
		}

	case crc_messageActionChatDeletePhoto:
		r = TL_messageActionChatDeletePhoto{}

	case crc_messageActionChatAddUser:
		r = TL_messageActionChatAddUser{
			m.VectorInt(),
		}

	case crc_messageActionChatDeleteUser:
		r = TL_messageActionChatDeleteUser{
			m.Int(),
		}

	case crc_dialog:
		flags := m.Flags()
		_ = flags
		r = TL_dialog{
			flags,
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.FlaggedObject(flags, 1),
		}

	case crc_photoEmpty:
		r = TL_photoEmpty{
			m.Long(),
		}

	case crc_photo:
		flags := m.Flags()
		_ = flags
		r = TL_photo{
			flags,
			m.Long(),
			m.Long(),
			m.Int(),
			m.Vector(),
		}

	case crc_photoSizeEmpty:
		r = TL_photoSizeEmpty{
			m.String(),
		}

	case crc_photoSize:
		r = TL_photoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_photoCachedSize:
		r = TL_photoCachedSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_geoPointEmpty:
		r = TL_geoPointEmpty{}

	case crc_geoPoint:
		r = TL_geoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_auth_checkedPhone:
		r = TL_auth_checkedPhone{
			m.Object(),
		}

	case crc_auth_sentCode:
		flags := m.Flags()
		_ = flags
		r = TL_auth_sentCode{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crc_auth_authorization:
		flags := m.Flags()
		_ = flags
		r = TL_auth_authorization{
			flags,
			m.FlaggedInt(flags, 0),
			m.Object(),
		}

	case crc_auth_exportedAuthorization:
		r = TL_auth_exportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputNotifyPeer:
		r = TL_inputNotifyPeer{
			m.Object(),
		}

	case crc_inputNotifyUsers:
		r = TL_inputNotifyUsers{}

	case crc_inputNotifyChats:
		r = TL_inputNotifyChats{}

	case crc_inputNotifyAll:
		r = TL_inputNotifyAll{}

	case crc_inputPeerNotifySettings:
		flags := m.Flags()
		_ = flags
		r = TL_inputPeerNotifySettings{
			flags,
			m.Int(),
			m.String(),
		}

	case crc_peerNotifyEventsEmpty:
		r = TL_peerNotifyEventsEmpty{}

	case crc_peerNotifyEventsAll:
		r = TL_peerNotifyEventsAll{}

	case crc_peerNotifySettingsEmpty:
		r = TL_peerNotifySettingsEmpty{}

	case crc_peerNotifySettings:
		flags := m.Flags()
		_ = flags
		r = TL_peerNotifySettings{
			flags,
			m.Int(),
			m.String(),
		}

	case crc_wallPaper:
		r = TL_wallPaper{
			m.Int(),
			m.String(),
			m.Vector(),
			m.Int(),
		}

	case crc_userFull:
		flags := m.Flags()
		_ = flags
		r = TL_userFull{
			flags,
			m.Object(),
			m.FlaggedString(flags, 1),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.Object(),
			m.FlaggedObject(flags, 3),
			m.Int(),
		}

	case crc_contact:
		r = TL_contact{
			m.Int(),
			m.Object(),
		}

	case crc_importedContact:
		r = TL_importedContact{
			m.Int(),
			m.Long(),
		}

	case crc_contactBlocked:
		r = TL_contactBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_contactStatus:
		r = TL_contactStatus{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_link:
		r = TL_contacts_link{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_contacts_contacts:
		r = TL_contacts_contacts{
			m.Vector(),
			m.Int(),
			m.Vector(),
		}

	case crc_contacts_contactsNotModified:
		r = TL_contacts_contactsNotModified{}

	case crc_contacts_importedContacts:
		r = TL_contacts_importedContacts{
			m.Vector(),
			m.Vector(),
			m.VectorLong(),
			m.Vector(),
		}

	case crc_contacts_blocked:
		r = TL_contacts_blocked{
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_blockedSlice:
		r = TL_contacts_blockedSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_found:
		r = TL_contacts_found{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogs:
		r = TL_messages_dialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogsSlice:
		r = TL_messages_dialogsSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messages:
		r = TL_messages_messages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messagesSlice:
		r = TL_messages_messagesSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_chats:
		r = TL_messages_chats{
			m.Vector(),
		}

	case crc_messages_chatFull:
		r = TL_messages_chatFull{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_affectedHistory:
		r = TL_messages_affectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessagesFilterEmpty:
		r = TL_inputMessagesFilterEmpty{}

	case crc_inputMessagesFilterPhotos:
		r = TL_inputMessagesFilterPhotos{}

	case crc_inputMessagesFilterVideo:
		r = TL_inputMessagesFilterVideo{}

	case crc_inputMessagesFilterPhotoVideo:
		r = TL_inputMessagesFilterPhotoVideo{}

	case crc_updateNewMessage:
		r = TL_updateNewMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateMessageID:
		r = TL_updateMessageID{
			m.Int(),
			m.Long(),
		}

	case crc_updateDeleteMessages:
		r = TL_updateDeleteMessages{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateUserTyping:
		r = TL_updateUserTyping{
			m.Int(),
			m.Object(),
		}

	case crc_updateChatUserTyping:
		r = TL_updateChatUserTyping{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_updateChatParticipants:
		r = TL_updateChatParticipants{
			m.Object(),
		}

	case crc_updateUserStatus:
		r = TL_updateUserStatus{
			m.Int(),
			m.Object(),
		}

	case crc_updateUserName:
		r = TL_updateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_updateUserPhoto:
		r = TL_updateUserPhoto{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updateContactRegistered:
		r = TL_updateContactRegistered{
			m.Int(),
			m.Int(),
		}

	case crc_updateContactLink:
		r = TL_updateContactLink{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updates_state:
		r = TL_updates_state{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_differenceEmpty:
		r = TL_updates_differenceEmpty{
			m.Int(),
			m.Int(),
		}

	case crc_updates_difference:
		r = TL_updates_difference{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updates_differenceSlice:
		r = TL_updates_differenceSlice{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updatesTooLong:
		r = TL_updatesTooLong{}

	case crc_updateShortMessage:
		flags := m.Flags()
		_ = flags
		r = TL_updateShortMessage{
			flags,
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.FlaggedVector(flags, 7),
		}

	case crc_updateShortChatMessage:
		flags := m.Flags()
		_ = flags
		r = TL_updateShortChatMessage{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.FlaggedVector(flags, 7),
		}

	case crc_updateShort:
		r = TL_updateShort{
			m.Object(),
			m.Int(),
		}

	case crc_updatesCombined:
		r = TL_updatesCombined{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates:
		r = TL_updates{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_photos_photo:
		r = TL_photos_photo{
			m.Object(),
			m.Vector(),
		}

	case crc_upload_file:
		r = TL_upload_file{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_dcOption:
		flags := m.Flags()
		_ = flags
		r = TL_dcOption{
			flags,
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_config:
		flags := m.Flags()
		_ = flags
		r = TL_config{
			flags,
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
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
			m.Vector(),
		}

	case crc_nearestDc:
		r = TL_nearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_help_appUpdate:
		r = TL_help_appUpdate{
			m.Int(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_help_noAppUpdate:
		r = TL_help_noAppUpdate{}

	case crc_help_inviteText:
		r = TL_help_inviteText{
			m.String(),
		}

	case crc_inputPeerNotifyEventsEmpty:
		r = TL_inputPeerNotifyEventsEmpty{}

	case crc_inputPeerNotifyEventsAll:
		r = TL_inputPeerNotifyEventsAll{}

	case crc_photos_photos:
		r = TL_photos_photos{
			m.Vector(),
			m.Vector(),
		}

	case crc_photos_photosSlice:
		r = TL_photos_photosSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_wallPaperSolid:
		r = TL_wallPaperSolid{
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_updateNewEncryptedMessage:
		r = TL_updateNewEncryptedMessage{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedChatTyping:
		r = TL_updateEncryptedChatTyping{
			m.Int(),
		}

	case crc_updateEncryption:
		r = TL_updateEncryption{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedMessagesRead:
		r = TL_updateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatEmpty:
		r = TL_encryptedChatEmpty{
			m.Int(),
		}

	case crc_encryptedChatWaiting:
		r = TL_encryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatRequested:
		r = TL_encryptedChatRequested{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_encryptedChat:
		r = TL_encryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_encryptedChatDiscarded:
		r = TL_encryptedChatDiscarded{
			m.Int(),
		}

	case crc_inputEncryptedChat:
		r = TL_inputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case crc_encryptedFileEmpty:
		r = TL_encryptedFileEmpty{}

	case crc_encryptedFile:
		r = TL_encryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputEncryptedFileEmpty:
		r = TL_inputEncryptedFileEmpty{}

	case crc_inputEncryptedFileUploaded:
		r = TL_inputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_inputEncryptedFile:
		r = TL_inputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case crc_inputEncryptedFileLocation:
		r = TL_inputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_encryptedMessage:
		r = TL_encryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_encryptedMessageService:
		r = TL_encryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_dhConfigNotModified:
		r = TL_messages_dhConfigNotModified{
			m.StringBytes(),
		}

	case crc_messages_dhConfig:
		r = TL_messages_dhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_sentEncryptedMessage:
		r = TL_messages_sentEncryptedMessage{
			m.Int(),
		}

	case crc_messages_sentEncryptedFile:
		r = TL_messages_sentEncryptedFile{
			m.Int(),
			m.Object(),
		}

	case crc_inputFileBig:
		r = TL_inputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case crc_inputEncryptedFileBigUploaded:
		r = TL_inputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crc_storage_filePdf:
		r = TL_storage_filePdf{}

	case crc_inputMessagesFilterDocument:
		r = TL_inputMessagesFilterDocument{}

	case crc_inputMessagesFilterPhotoVideoDocuments:
		r = TL_inputMessagesFilterPhotoVideoDocuments{}

	case crc_updateChatParticipantAdd:
		r = TL_updateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantDelete:
		r = TL_updateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateDcOptions:
		r = TL_updateDcOptions{
			m.Vector(),
		}

	case crc_inputMediaUploadedDocument:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaUploadedDocument{
			flags,
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.String(),
			m.Vector(),
			m.String(),
			m.FlaggedVector(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crc_inputMediaDocument:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaDocument{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_messageMediaDocument:
		flags := m.Flags()
		_ = flags
		r = TL_messageMediaDocument{
			flags,
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crc_inputDocumentEmpty:
		r = TL_inputDocumentEmpty{}

	case crc_inputDocument:
		r = TL_inputDocument{
			m.Long(),
			m.Long(),
		}

	case crc_inputDocumentFileLocation:
		r = TL_inputDocumentFileLocation{
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crc_documentEmpty:
		r = TL_documentEmpty{
			m.Long(),
		}

	case crc_document:
		r = TL_document{
			m.Long(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_help_support:
		r = TL_help_support{
			m.String(),
			m.Object(),
		}

	case crc_notifyAll:
		r = TL_notifyAll{}

	case crc_notifyChats:
		r = TL_notifyChats{}

	case crc_notifyPeer:
		r = TL_notifyPeer{
			m.Object(),
		}

	case crc_notifyUsers:
		r = TL_notifyUsers{}

	case crc_updateUserBlocked:
		r = TL_updateUserBlocked{
			m.Int(),
			m.Object(),
		}

	case crc_updateNotifySettings:
		r = TL_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_sendMessageTypingAction:
		r = TL_sendMessageTypingAction{}

	case crc_sendMessageCancelAction:
		r = TL_sendMessageCancelAction{}

	case crc_sendMessageRecordVideoAction:
		r = TL_sendMessageRecordVideoAction{}

	case crc_sendMessageUploadVideoAction:
		r = TL_sendMessageUploadVideoAction{
			m.Int(),
		}

	case crc_sendMessageRecordAudioAction:
		r = TL_sendMessageRecordAudioAction{}

	case crc_sendMessageUploadAudioAction:
		r = TL_sendMessageUploadAudioAction{
			m.Int(),
		}

	case crc_sendMessageUploadPhotoAction:
		r = TL_sendMessageUploadPhotoAction{
			m.Int(),
		}

	case crc_sendMessageUploadDocumentAction:
		r = TL_sendMessageUploadDocumentAction{
			m.Int(),
		}

	case crc_sendMessageGeoLocationAction:
		r = TL_sendMessageGeoLocationAction{}

	case crc_sendMessageChooseContactAction:
		r = TL_sendMessageChooseContactAction{}

	case crc_updateServiceNotification:
		flags := m.Flags()
		_ = flags
		r = TL_updateServiceNotification{
			flags,
			m.FlaggedInt(flags, 1),
			m.String(),
			m.String(),
			m.Object(),
			m.Vector(),
		}

	case crc_userStatusRecently:
		r = TL_userStatusRecently{}

	case crc_userStatusLastWeek:
		r = TL_userStatusLastWeek{}

	case crc_userStatusLastMonth:
		r = TL_userStatusLastMonth{}

	case crc_updatePrivacy:
		r = TL_updatePrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_inputPrivacyKeyStatusTimestamp:
		r = TL_inputPrivacyKeyStatusTimestamp{}

	case crc_privacyKeyStatusTimestamp:
		r = TL_privacyKeyStatusTimestamp{}

	case crc_inputPrivacyValueAllowContacts:
		r = TL_inputPrivacyValueAllowContacts{}

	case crc_inputPrivacyValueAllowAll:
		r = TL_inputPrivacyValueAllowAll{}

	case crc_inputPrivacyValueAllowUsers:
		r = TL_inputPrivacyValueAllowUsers{
			m.Vector(),
		}

	case crc_inputPrivacyValueDisallowContacts:
		r = TL_inputPrivacyValueDisallowContacts{}

	case crc_inputPrivacyValueDisallowAll:
		r = TL_inputPrivacyValueDisallowAll{}

	case crc_inputPrivacyValueDisallowUsers:
		r = TL_inputPrivacyValueDisallowUsers{
			m.Vector(),
		}

	case crc_privacyValueAllowContacts:
		r = TL_privacyValueAllowContacts{}

	case crc_privacyValueAllowAll:
		r = TL_privacyValueAllowAll{}

	case crc_privacyValueAllowUsers:
		r = TL_privacyValueAllowUsers{
			m.VectorInt(),
		}

	case crc_privacyValueDisallowContacts:
		r = TL_privacyValueDisallowContacts{}

	case crc_privacyValueDisallowAll:
		r = TL_privacyValueDisallowAll{}

	case crc_privacyValueDisallowUsers:
		r = TL_privacyValueDisallowUsers{
			m.VectorInt(),
		}

	case crc_account_privacyRules:
		r = TL_account_privacyRules{
			m.Vector(),
			m.Vector(),
		}

	case crc_accountDaysTTL:
		r = TL_accountDaysTTL{
			m.Int(),
		}

	case crc_updateUserPhone:
		r = TL_updateUserPhone{
			m.Int(),
			m.String(),
		}

	case crc_disabledFeature:
		r = TL_disabledFeature{
			m.String(),
			m.String(),
		}

	case crc_documentAttributeImageSize:
		r = TL_documentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAnimated:
		r = TL_documentAttributeAnimated{}

	case crc_documentAttributeSticker:
		flags := m.Flags()
		_ = flags
		r = TL_documentAttributeSticker{
			flags,
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 0),
		}

	case crc_documentAttributeVideo:
		flags := m.Flags()
		_ = flags
		r = TL_documentAttributeVideo{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAudio:
		flags := m.Flags()
		_ = flags
		r = TL_documentAttributeAudio{
			flags,
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedStringBytes(flags, 2),
		}

	case crc_documentAttributeFilename:
		r = TL_documentAttributeFilename{
			m.String(),
		}

	case crc_messages_stickersNotModified:
		r = TL_messages_stickersNotModified{}

	case crc_messages_stickers:
		r = TL_messages_stickers{
			m.String(),
			m.Vector(),
		}

	case crc_stickerPack:
		r = TL_stickerPack{
			m.String(),
			m.VectorLong(),
		}

	case crc_messages_allStickersNotModified:
		r = TL_messages_allStickersNotModified{}

	case crc_messages_allStickers:
		r = TL_messages_allStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_account_noPassword:
		r = TL_account_noPassword{
			m.StringBytes(),
			m.String(),
		}

	case crc_account_password:
		r = TL_account_password{
			m.StringBytes(),
			m.StringBytes(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case crc_updateReadHistoryInbox:
		r = TL_updateReadHistoryInbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadHistoryOutbox:
		r = TL_updateReadHistoryOutbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_affectedMessages:
		r = TL_messages_affectedMessages{
			m.Int(),
			m.Int(),
		}

	case crc_contactLinkUnknown:
		r = TL_contactLinkUnknown{}

	case crc_contactLinkNone:
		r = TL_contactLinkNone{}

	case crc_contactLinkHasPhone:
		r = TL_contactLinkHasPhone{}

	case crc_contactLinkContact:
		r = TL_contactLinkContact{}

	case crc_updateWebPage:
		r = TL_updateWebPage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_webPageEmpty:
		r = TL_webPageEmpty{
			m.Long(),
		}

	case crc_webPagePending:
		r = TL_webPagePending{
			m.Long(),
			m.Int(),
		}

	case crc_webPage:
		flags := m.Flags()
		_ = flags
		r = TL_webPage{
			flags,
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			m.FlaggedString(flags, 8),
			m.FlaggedObject(flags, 9),
			m.FlaggedObject(flags, 10),
		}

	case crc_messageMediaWebPage:
		r = TL_messageMediaWebPage{
			m.Object(),
		}

	case crc_authorization:
		r = TL_authorization{
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

	case crc_account_authorizations:
		r = TL_account_authorizations{
			m.Vector(),
		}

	case crc_account_passwordSettings:
		r = TL_account_passwordSettings{
			m.String(),
		}

	case crc_account_passwordInputSettings:
		flags := m.Flags()
		_ = flags
		r = TL_account_passwordInputSettings{
			flags,
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_auth_passwordRecovery:
		r = TL_auth_passwordRecovery{
			m.String(),
		}

	case crc_inputMediaVenue:
		r = TL_inputMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messageMediaVenue:
		r = TL_messageMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_receivedNotifyMessage:
		r = TL_receivedNotifyMessage{
			m.Int(),
			m.Int(),
		}

	case crc_chatInviteEmpty:
		r = TL_chatInviteEmpty{}

	case crc_chatInviteExported:
		r = TL_chatInviteExported{
			m.String(),
		}

	case crc_chatInviteAlready:
		r = TL_chatInviteAlready{
			m.Object(),
		}

	case crc_chatInvite:
		flags := m.Flags()
		_ = flags
		r = TL_chatInvite{
			flags,
			m.String(),
			m.Object(),
			m.Int(),
			m.FlaggedVector(flags, 4),
		}

	case crc_messageActionChatJoinedByLink:
		r = TL_messageActionChatJoinedByLink{
			m.Int(),
		}

	case crc_updateReadMessagesContents:
		r = TL_updateReadMessagesContents{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_inputStickerSetEmpty:
		r = TL_inputStickerSetEmpty{}

	case crc_inputStickerSetID:
		r = TL_inputStickerSetID{
			m.Long(),
			m.Long(),
		}

	case crc_inputStickerSetShortName:
		r = TL_inputStickerSetShortName{
			m.String(),
		}

	case crc_stickerSet:
		flags := m.Flags()
		_ = flags
		r = TL_stickerSet{
			flags,
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_stickerSet:
		r = TL_messages_stickerSet{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_user:
		flags := m.Flags()
		_ = flags
		r = TL_user{
			flags,
			m.Int(),
			m.FlaggedLong(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedObject(flags, 5),
			m.FlaggedObject(flags, 6),
			m.FlaggedInt(flags, 14),
			m.FlaggedString(flags, 18),
			m.FlaggedString(flags, 19),
			m.FlaggedString(flags, 22),
		}

	case crc_botCommand:
		r = TL_botCommand{
			m.String(),
			m.String(),
		}

	case crc_botInfo:
		r = TL_botInfo{
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_keyboardButton:
		r = TL_keyboardButton{
			m.String(),
		}

	case crc_keyboardButtonRow:
		r = TL_keyboardButtonRow{
			m.Vector(),
		}

	case crc_replyKeyboardHide:
		flags := m.Flags()
		_ = flags
		r = TL_replyKeyboardHide{
			flags,
		}

	case crc_replyKeyboardForceReply:
		flags := m.Flags()
		_ = flags
		r = TL_replyKeyboardForceReply{
			flags,
		}

	case crc_replyKeyboardMarkup:
		flags := m.Flags()
		_ = flags
		r = TL_replyKeyboardMarkup{
			flags,
			m.Vector(),
		}

	case crc_inputMessagesFilterUrl:
		r = TL_inputMessagesFilterUrl{}

	case crc_inputPeerUser:
		r = TL_inputPeerUser{
			m.Int(),
			m.Long(),
		}

	case crc_inputUser:
		r = TL_inputUser{
			m.Int(),
			m.Long(),
		}

	case crc_messageEntityUnknown:
		r = TL_messageEntityUnknown{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityMention:
		r = TL_messageEntityMention{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityHashtag:
		r = TL_messageEntityHashtag{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBotCommand:
		r = TL_messageEntityBotCommand{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityUrl:
		r = TL_messageEntityUrl{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityEmail:
		r = TL_messageEntityEmail{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBold:
		r = TL_messageEntityBold{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityItalic:
		r = TL_messageEntityItalic{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityCode:
		r = TL_messageEntityCode{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityPre:
		r = TL_messageEntityPre{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_messageEntityTextUrl:
		r = TL_messageEntityTextUrl{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_updateShortSentMessage:
		flags := m.Flags()
		_ = flags
		r = TL_updateShortSentMessage{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 9),
			m.FlaggedVector(flags, 7),
		}

	case crc_inputPeerChannel:
		r = TL_inputPeerChannel{
			m.Int(),
			m.Long(),
		}

	case crc_peerChannel:
		r = TL_peerChannel{
			m.Int(),
		}

	case crc_channel:
		flags := m.Flags()
		_ = flags
		r = TL_channel{
			flags,
			m.Int(),
			m.FlaggedLong(flags, 13),
			m.String(),
			m.FlaggedString(flags, 6),
			m.Object(),
			m.Int(),
			m.Int(),
			m.FlaggedString(flags, 9),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 15),
		}

	case crc_channelForbidden:
		flags := m.Flags()
		_ = flags
		r = TL_channelForbidden{
			flags,
			m.Int(),
			m.Long(),
			m.String(),
			m.FlaggedInt(flags, 16),
		}

	case crc_channelFull:
		flags := m.Flags()
		_ = flags
		r = TL_channelFull{
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
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 5),
			m.FlaggedObject(flags, 8),
		}

	case crc_messageActionChannelCreate:
		r = TL_messageActionChannelCreate{
			m.String(),
		}

	case crc_messages_channelMessages:
		flags := m.Flags()
		_ = flags
		r = TL_messages_channelMessages{
			flags,
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updateChannelTooLong:
		flags := m.Flags()
		_ = flags
		r = TL_updateChannelTooLong{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 0),
		}

	case crc_updateChannel:
		r = TL_updateChannel{
			m.Int(),
		}

	case crc_updateNewChannelMessage:
		r = TL_updateNewChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadChannelInbox:
		r = TL_updateReadChannelInbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDeleteChannelMessages:
		r = TL_updateDeleteChannelMessages{
			m.Int(),
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChannelMessageViews:
		r = TL_updateChannelMessageViews{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputChannelEmpty:
		r = TL_inputChannelEmpty{}

	case crc_inputChannel:
		r = TL_inputChannel{
			m.Int(),
			m.Long(),
		}

	case crc_contacts_resolvedPeer:
		r = TL_contacts_resolvedPeer{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messageRange:
		r = TL_messageRange{
			m.Int(),
			m.Int(),
		}

	case crc_updates_channelDifferenceEmpty:
		flags := m.Flags()
		_ = flags
		r = TL_updates_channelDifferenceEmpty{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
		}

	case crc_updates_channelDifferenceTooLong:
		flags := m.Flags()
		_ = flags
		r = TL_updates_channelDifferenceTooLong{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updates_channelDifference:
		flags := m.Flags()
		_ = flags
		r = TL_updates_channelDifference{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channelMessagesFilterEmpty:
		r = TL_channelMessagesFilterEmpty{}

	case crc_channelMessagesFilter:
		flags := m.Flags()
		_ = flags
		r = TL_channelMessagesFilter{
			flags,
			m.Vector(),
		}

	case crc_channelParticipant:
		r = TL_channelParticipant{
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantSelf:
		r = TL_channelParticipantSelf{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantCreator:
		r = TL_channelParticipantCreator{
			m.Int(),
		}

	case crc_channelParticipantsRecent:
		r = TL_channelParticipantsRecent{}

	case crc_channelParticipantsAdmins:
		r = TL_channelParticipantsAdmins{}

	case crc_channelParticipantsKicked:
		r = TL_channelParticipantsKicked{
			m.String(),
		}

	case crc_channels_channelParticipants:
		r = TL_channels_channelParticipants{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channels_channelParticipant:
		r = TL_channels_channelParticipant{
			m.Object(),
			m.Vector(),
		}

	case crc_true:
		r = TL_true{}

	case crc_chatParticipantCreator:
		r = TL_chatParticipantCreator{
			m.Int(),
		}

	case crc_chatParticipantAdmin:
		r = TL_chatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatAdmins:
		r = TL_updateChatAdmins{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_updateChatParticipantAdmin:
		r = TL_updateChatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messageActionChatMigrateTo:
		r = TL_messageActionChatMigrateTo{
			m.Int(),
		}

	case crc_messageActionChannelMigrateFrom:
		r = TL_messageActionChannelMigrateFrom{
			m.String(),
			m.Int(),
		}

	case crc_channelParticipantsBots:
		r = TL_channelParticipantsBots{}

	case crc_inputReportReasonSpam:
		r = TL_inputReportReasonSpam{}

	case crc_inputReportReasonViolence:
		r = TL_inputReportReasonViolence{}

	case crc_inputReportReasonPornography:
		r = TL_inputReportReasonPornography{}

	case crc_inputReportReasonOther:
		r = TL_inputReportReasonOther{
			m.String(),
		}

	case crc_updateNewStickerSet:
		r = TL_updateNewStickerSet{
			m.Object(),
		}

	case crc_updateStickerSetsOrder:
		flags := m.Flags()
		_ = flags
		r = TL_updateStickerSetsOrder{
			flags,
			m.VectorLong(),
		}

	case crc_updateStickerSets:
		r = TL_updateStickerSets{}

	case crc_help_termsOfService:
		r = TL_help_termsOfService{
			m.String(),
		}

	case crc_foundGif:
		r = TL_foundGif{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMediaGifExternal:
		r = TL_inputMediaGifExternal{
			m.String(),
			m.String(),
		}

	case crc_messages_foundGifs:
		r = TL_messages_foundGifs{
			m.Int(),
			m.Vector(),
		}

	case crc_inputMessagesFilterGif:
		r = TL_inputMessagesFilterGif{}

	case crc_updateSavedGifs:
		r = TL_updateSavedGifs{}

	case crc_updateBotInlineQuery:
		flags := m.Flags()
		_ = flags
		r = TL_updateBotInlineQuery{
			flags,
			m.Long(),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.String(),
		}

	case crc_foundGifCached:
		r = TL_foundGifCached{
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_savedGifsNotModified:
		r = TL_messages_savedGifsNotModified{}

	case crc_messages_savedGifs:
		r = TL_messages_savedGifs{
			m.Int(),
			m.Vector(),
		}

	case crc_inputBotInlineMessageMediaAuto:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageMediaAuto{
			flags,
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineMessageText:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageText{
			flags,
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineResult:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineResult{
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
			m.Object(),
		}

	case crc_botInlineMessageMediaAuto:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageMediaAuto{
			flags,
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineMessageText:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageText{
			flags,
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineResult:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineResult{
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
			m.Object(),
		}

	case crc_messages_botResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_botResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 2),
			m.Vector(),
			m.Int(),
		}

	case crc_inputMessagesFilterVoice:
		r = TL_inputMessagesFilterVoice{}

	case crc_inputMessagesFilterMusic:
		r = TL_inputMessagesFilterMusic{}

	case crc_updateBotInlineSend:
		flags := m.Flags()
		_ = flags
		r = TL_updateBotInlineSend{
			flags,
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.FlaggedObject(flags, 1),
		}

	case crc_inputPrivacyKeyChatInvite:
		r = TL_inputPrivacyKeyChatInvite{}

	case crc_privacyKeyChatInvite:
		r = TL_privacyKeyChatInvite{}

	case crc_updateEditChannelMessage:
		r = TL_updateEditChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_exportedMessageLink:
		r = TL_exportedMessageLink{
			m.String(),
		}

	case crc_messageFwdHeader:
		flags := m.Flags()
		_ = flags
		r = TL_messageFwdHeader{
			flags,
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
		}

	case crc_messageActionPinMessage:
		r = TL_messageActionPinMessage{}

	case crc_peerSettings:
		flags := m.Flags()
		_ = flags
		r = TL_peerSettings{
			flags,
		}

	case crc_updateChannelPinnedMessage:
		r = TL_updateChannelPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case crc_keyboardButtonUrl:
		r = TL_keyboardButtonUrl{
			m.String(),
			m.String(),
		}

	case crc_keyboardButtonCallback:
		r = TL_keyboardButtonCallback{
			m.String(),
			m.StringBytes(),
		}

	case crc_keyboardButtonRequestPhone:
		r = TL_keyboardButtonRequestPhone{
			m.String(),
		}

	case crc_keyboardButtonRequestGeoLocation:
		r = TL_keyboardButtonRequestGeoLocation{
			m.String(),
		}

	case crc_auth_codeTypeSms:
		r = TL_auth_codeTypeSms{}

	case crc_auth_codeTypeCall:
		r = TL_auth_codeTypeCall{}

	case crc_auth_codeTypeFlashCall:
		r = TL_auth_codeTypeFlashCall{}

	case crc_auth_sentCodeTypeApp:
		r = TL_auth_sentCodeTypeApp{
			m.Int(),
		}

	case crc_auth_sentCodeTypeSms:
		r = TL_auth_sentCodeTypeSms{
			m.Int(),
		}

	case crc_auth_sentCodeTypeCall:
		r = TL_auth_sentCodeTypeCall{
			m.Int(),
		}

	case crc_auth_sentCodeTypeFlashCall:
		r = TL_auth_sentCodeTypeFlashCall{
			m.String(),
		}

	case crc_keyboardButtonSwitchInline:
		flags := m.Flags()
		_ = flags
		r = TL_keyboardButtonSwitchInline{
			flags,
			m.String(),
			m.String(),
		}

	case crc_replyInlineMarkup:
		r = TL_replyInlineMarkup{
			m.Vector(),
		}

	case crc_messages_botCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = TL_messages_botCallbackAnswer{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case crc_updateBotCallbackQuery:
		flags := m.Flags()
		_ = flags
		r = TL_updateBotCallbackQuery{
			flags,
			m.Long(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_messages_messageEditData:
		flags := m.Flags()
		_ = flags
		r = TL_messages_messageEditData{
			flags,
		}

	case crc_updateEditMessage:
		r = TL_updateEditMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_inputBotInlineMessageMediaGeo:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageMediaGeo{
			flags,
			m.Object(),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineMessageMediaVenue:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageMediaVenue{
			flags,
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineMessageMediaContact:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageMediaContact{
			flags,
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineMessageMediaGeo:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageMediaGeo{
			flags,
			m.Object(),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineMessageMediaVenue:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageMediaVenue{
			flags,
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_botInlineMessageMediaContact:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMessageMediaContact{
			flags,
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineResultPhoto:
		r = TL_inputBotInlineResultPhoto{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputBotInlineResultDocument:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineResultDocument{
			flags,
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.Object(),
			m.Object(),
		}

	case crc_botInlineMediaResult:
		flags := m.Flags()
		_ = flags
		r = TL_botInlineMediaResult{
			flags,
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.Object(),
		}

	case crc_inputBotInlineMessageID:
		r = TL_inputBotInlineMessageID{
			m.Int(),
			m.Long(),
			m.Long(),
		}

	case crc_updateInlineBotCallbackQuery:
		flags := m.Flags()
		_ = flags
		r = TL_updateInlineBotCallbackQuery{
			flags,
			m.Long(),
			m.Int(),
			m.Object(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crc_inlineBotSwitchPM:
		r = TL_inlineBotSwitchPM{
			m.String(),
			m.String(),
		}

	case crc_messageEntityMentionName:
		r = TL_messageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessageEntityMentionName:
		r = TL_inputMessageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_messages_peerDialogs:
		r = TL_messages_peerDialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_topPeer:
		r = TL_topPeer{
			m.Object(),
			m.Double(),
		}

	case crc_topPeerCategoryBotsPM:
		r = TL_topPeerCategoryBotsPM{}

	case crc_topPeerCategoryBotsInline:
		r = TL_topPeerCategoryBotsInline{}

	case crc_topPeerCategoryCorrespondents:
		r = TL_topPeerCategoryCorrespondents{}

	case crc_topPeerCategoryGroups:
		r = TL_topPeerCategoryGroups{}

	case crc_topPeerCategoryChannels:
		r = TL_topPeerCategoryChannels{}

	case crc_topPeerCategoryPeers:
		r = TL_topPeerCategoryPeers{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crc_contacts_topPeersNotModified:
		r = TL_contacts_topPeersNotModified{}

	case crc_contacts_topPeers:
		r = TL_contacts_topPeers{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_inputMessagesFilterChatPhotos:
		r = TL_inputMessagesFilterChatPhotos{}

	case crc_updateReadChannelOutbox:
		r = TL_updateReadChannelOutbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDraftMessage:
		r = TL_updateDraftMessage{
			m.Object(),
			m.Object(),
		}

	case crc_draftMessageEmpty:
		r = TL_draftMessageEmpty{}

	case crc_draftMessage:
		flags := m.Flags()
		_ = flags
		r = TL_draftMessage{
			flags,
			m.FlaggedInt(flags, 0),
			m.String(),
			m.FlaggedVector(flags, 3),
			m.Int(),
		}

	case crc_messageActionHistoryClear:
		r = TL_messageActionHistoryClear{}

	case crc_updateReadFeaturedStickers:
		r = TL_updateReadFeaturedStickers{}

	case crc_updateRecentStickers:
		r = TL_updateRecentStickers{}

	case crc_messages_featuredStickersNotModified:
		r = TL_messages_featuredStickersNotModified{}

	case crc_messages_featuredStickers:
		r = TL_messages_featuredStickers{
			m.Int(),
			m.Vector(),
			m.VectorLong(),
		}

	case crc_messages_recentStickersNotModified:
		r = TL_messages_recentStickersNotModified{}

	case crc_messages_recentStickers:
		r = TL_messages_recentStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_messages_archivedStickers:
		r = TL_messages_archivedStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_messages_stickerSetInstallResultSuccess:
		r = TL_messages_stickerSetInstallResultSuccess{}

	case crc_messages_stickerSetInstallResultArchive:
		r = TL_messages_stickerSetInstallResultArchive{
			m.Vector(),
		}

	case crc_stickerSetCovered:
		r = TL_stickerSetCovered{
			m.Object(),
			m.Object(),
		}

	case crc_inputMediaPhotoExternal:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaPhotoExternal{
			flags,
			m.String(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_inputMediaDocumentExternal:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaDocumentExternal{
			flags,
			m.String(),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crc_updateConfig:
		r = TL_updateConfig{}

	case crc_updatePtsChanged:
		r = TL_updatePtsChanged{}

	case crc_messageActionGameScore:
		r = TL_messageActionGameScore{
			m.Long(),
			m.Int(),
		}

	case crc_documentAttributeHasStickers:
		r = TL_documentAttributeHasStickers{}

	case crc_keyboardButtonGame:
		r = TL_keyboardButtonGame{
			m.String(),
		}

	case crc_stickerSetMultiCovered:
		r = TL_stickerSetMultiCovered{
			m.Object(),
			m.Vector(),
		}

	case crc_maskCoords:
		r = TL_maskCoords{
			m.Int(),
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case crc_inputStickeredMediaPhoto:
		r = TL_inputStickeredMediaPhoto{
			m.Object(),
		}

	case crc_inputStickeredMediaDocument:
		r = TL_inputStickeredMediaDocument{
			m.Object(),
		}

	case crc_inputMediaGame:
		r = TL_inputMediaGame{
			m.Object(),
		}

	case crc_messageMediaGame:
		r = TL_messageMediaGame{
			m.Object(),
		}

	case crc_inputBotInlineMessageGame:
		flags := m.Flags()
		_ = flags
		r = TL_inputBotInlineMessageGame{
			flags,
			m.FlaggedObject(flags, 2),
		}

	case crc_inputBotInlineResultGame:
		r = TL_inputBotInlineResultGame{
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_game:
		flags := m.Flags()
		_ = flags
		r = TL_game{
			flags,
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 0),
		}

	case crc_inputGameID:
		r = TL_inputGameID{
			m.Long(),
			m.Long(),
		}

	case crc_inputGameShortName:
		r = TL_inputGameShortName{
			m.Object(),
			m.String(),
		}

	case crc_highScore:
		r = TL_highScore{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_highScores:
		r = TL_messages_highScores{
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_chatsSlice:
		r = TL_messages_chatsSlice{
			m.Int(),
			m.Vector(),
		}

	case crc_updateChannelWebPage:
		r = TL_updateChannelWebPage{
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_differenceTooLong:
		r = TL_updates_differenceTooLong{
			m.Int(),
		}

	case crc_sendMessageGamePlayAction:
		r = TL_sendMessageGamePlayAction{}

	case crc_webPageNotModified:
		r = TL_webPageNotModified{}

	case crc_textEmpty:
		r = TL_textEmpty{}

	case crc_textPlain:
		r = TL_textPlain{
			m.String(),
		}

	case crc_textBold:
		r = TL_textBold{
			m.Object(),
		}

	case crc_textItalic:
		r = TL_textItalic{
			m.Object(),
		}

	case crc_textUnderline:
		r = TL_textUnderline{
			m.Object(),
		}

	case crc_textStrike:
		r = TL_textStrike{
			m.Object(),
		}

	case crc_textFixed:
		r = TL_textFixed{
			m.Object(),
		}

	case crc_textUrl:
		r = TL_textUrl{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case crc_textEmail:
		r = TL_textEmail{
			m.Object(),
			m.String(),
		}

	case crc_textConcat:
		r = TL_textConcat{
			m.Vector(),
		}

	case crc_pageBlockTitle:
		r = TL_pageBlockTitle{
			m.Object(),
		}

	case crc_pageBlockSubtitle:
		r = TL_pageBlockSubtitle{
			m.Object(),
		}

	case crc_pageBlockAuthorDate:
		r = TL_pageBlockAuthorDate{
			m.Object(),
			m.Int(),
		}

	case crc_pageBlockHeader:
		r = TL_pageBlockHeader{
			m.Object(),
		}

	case crc_pageBlockSubheader:
		r = TL_pageBlockSubheader{
			m.Object(),
		}

	case crc_pageBlockParagraph:
		r = TL_pageBlockParagraph{
			m.Object(),
		}

	case crc_pageBlockPreformatted:
		r = TL_pageBlockPreformatted{
			m.Object(),
			m.String(),
		}

	case crc_pageBlockFooter:
		r = TL_pageBlockFooter{
			m.Object(),
		}

	case crc_pageBlockDivider:
		r = TL_pageBlockDivider{}

	case crc_pageBlockList:
		r = TL_pageBlockList{
			m.Object(),
			m.Vector(),
		}

	case crc_pageBlockBlockquote:
		r = TL_pageBlockBlockquote{
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockPullquote:
		r = TL_pageBlockPullquote{
			m.Object(),
			m.Object(),
		}

	case crc_pageBlockPhoto:
		r = TL_pageBlockPhoto{
			m.Long(),
			m.Object(),
		}

	case crc_pageBlockVideo:
		flags := m.Flags()
		_ = flags
		r = TL_pageBlockVideo{
			flags,
			m.Long(),
			m.Object(),
		}

	case crc_pageBlockCover:
		r = TL_pageBlockCover{
			m.Object(),
		}

	case crc_pageBlockEmbed:
		flags := m.Flags()
		_ = flags
		r = TL_pageBlockEmbed{
			flags,
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedLong(flags, 4),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_pageBlockEmbedPost:
		r = TL_pageBlockEmbedPost{
			m.String(),
			m.Long(),
			m.Long(),
			m.String(),
			m.Int(),
			m.Vector(),
			m.Object(),
		}

	case crc_pageBlockSlideshow:
		r = TL_pageBlockSlideshow{
			m.Vector(),
			m.Object(),
		}

	case crc_pagePart:
		r = TL_pagePart{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_pageFull:
		r = TL_pageFull{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updatePhoneCall:
		r = TL_updatePhoneCall{
			m.Object(),
		}

	case crc_updateDialogPinned:
		flags := m.Flags()
		_ = flags
		r = TL_updateDialogPinned{
			flags,
			m.Object(),
		}

	case crc_updatePinnedDialogs:
		flags := m.Flags()
		_ = flags
		r = TL_updatePinnedDialogs{
			flags,
			m.FlaggedVector(flags, 0),
		}

	case crc_inputPrivacyKeyPhoneCall:
		r = TL_inputPrivacyKeyPhoneCall{}

	case crc_privacyKeyPhoneCall:
		r = TL_privacyKeyPhoneCall{}

	case crc_pageBlockUnsupported:
		r = TL_pageBlockUnsupported{}

	case crc_pageBlockAnchor:
		r = TL_pageBlockAnchor{
			m.String(),
		}

	case crc_pageBlockCollage:
		r = TL_pageBlockCollage{
			m.Vector(),
			m.Object(),
		}

	case crc_inputPhoneCall:
		r = TL_inputPhoneCall{
			m.Long(),
			m.Long(),
		}

	case crc_phoneCallEmpty:
		r = TL_phoneCallEmpty{
			m.Long(),
		}

	case crc_phoneCallWaiting:
		flags := m.Flags()
		_ = flags
		r = TL_phoneCallWaiting{
			flags,
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case crc_phoneCallRequested:
		r = TL_phoneCallRequested{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phoneCall:
		r = TL_phoneCall{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.Int(),
		}

	case crc_phoneCallDiscarded:
		flags := m.Flags()
		_ = flags
		r = TL_phoneCallDiscarded{
			flags,
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crc_phoneConnection:
		r = TL_phoneConnection{
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_phoneCallProtocol:
		flags := m.Flags()
		_ = flags
		r = TL_phoneCallProtocol{
			flags,
			m.Int(),
			m.Int(),
		}

	case crc_phone_phoneCall:
		r = TL_phone_phoneCall{
			m.Object(),
			m.Vector(),
		}

	case crc_phoneCallDiscardReasonMissed:
		r = TL_phoneCallDiscardReasonMissed{}

	case crc_phoneCallDiscardReasonDisconnect:
		r = TL_phoneCallDiscardReasonDisconnect{}

	case crc_phoneCallDiscardReasonHangup:
		r = TL_phoneCallDiscardReasonHangup{}

	case crc_phoneCallDiscardReasonBusy:
		r = TL_phoneCallDiscardReasonBusy{}

	case crc_inputMessagesFilterPhoneCalls:
		flags := m.Flags()
		_ = flags
		r = TL_inputMessagesFilterPhoneCalls{
			flags,
		}

	case crc_messageActionPhoneCall:
		flags := m.Flags()
		_ = flags
		r = TL_messageActionPhoneCall{
			flags,
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crc_invoice:
		flags := m.Flags()
		_ = flags
		r = TL_invoice{
			flags,
			m.String(),
			m.Vector(),
		}

	case crc_inputMediaInvoice:
		flags := m.Flags()
		_ = flags
		r = TL_inputMediaInvoice{
			flags,
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Object(),
			m.StringBytes(),
			m.String(),
			m.String(),
		}

	case crc_messageActionPaymentSentMe:
		flags := m.Flags()
		_ = flags
		r = TL_messageActionPaymentSentMe{
			flags,
			m.String(),
			m.Long(),
			m.StringBytes(),
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.Object(),
		}

	case crc_messageMediaInvoice:
		flags := m.Flags()
		_ = flags
		r = TL_messageMediaInvoice{
			flags,
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 2),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_keyboardButtonBuy:
		r = TL_keyboardButtonBuy{
			m.String(),
		}

	case crc_messageActionPaymentSent:
		r = TL_messageActionPaymentSent{
			m.String(),
			m.Long(),
		}

	case crc_payments_paymentForm:
		flags := m.Flags()
		_ = flags
		r = TL_payments_paymentForm{
			flags,
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 4),
			m.FlaggedObject(flags, 4),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.Vector(),
		}

	case crc_postAddress:
		r = TL_postAddress{
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
		r = TL_paymentRequestedInfo{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case crc_updateBotWebhookJSON:
		r = TL_updateBotWebhookJSON{
			m.Object(),
		}

	case crc_updateBotWebhookJSONQuery:
		r = TL_updateBotWebhookJSONQuery{
			m.Long(),
			m.Object(),
			m.Int(),
		}

	case crc_updateBotShippingQuery:
		r = TL_updateBotShippingQuery{
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_updateBotPrecheckoutQuery:
		flags := m.Flags()
		_ = flags
		r = TL_updateBotPrecheckoutQuery{
			flags,
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Long(),
		}

	case crc_dataJSON:
		r = TL_dataJSON{
			m.String(),
		}

	case crc_labeledPrice:
		r = TL_labeledPrice{
			m.String(),
			m.Long(),
		}

	case crc_paymentCharge:
		r = TL_paymentCharge{
			m.String(),
			m.String(),
		}

	case crc_paymentSavedCredentialsCard:
		r = TL_paymentSavedCredentialsCard{
			m.String(),
			m.String(),
		}

	case crc_webDocument:
		r = TL_webDocument{
			m.String(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Vector(),
			m.Int(),
		}

	case crc_inputWebDocument:
		r = TL_inputWebDocument{
			m.String(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_inputWebFileLocation:
		r = TL_inputWebFileLocation{
			m.String(),
			m.Long(),
		}

	case crc_upload_webFile:
		r = TL_upload_webFile{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_payments_validatedRequestedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_payments_validatedRequestedInfo{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedVector(flags, 1),
		}

	case crc_payments_paymentResult:
		r = TL_payments_paymentResult{
			m.Object(),
		}

	case crc_payments_paymentVerficationNeeded:
		r = TL_payments_paymentVerficationNeeded{
			m.String(),
		}

	case crc_payments_paymentReceipt:
		flags := m.Flags()
		_ = flags
		r = TL_payments_paymentReceipt{
			flags,
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.String(),
			m.Long(),
			m.String(),
			m.Vector(),
		}

	case crc_payments_savedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_payments_savedInfo{
			flags,
			m.FlaggedObject(flags, 0),
		}

	case crc_inputPaymentCredentialsSaved:
		r = TL_inputPaymentCredentialsSaved{
			m.String(),
			m.StringBytes(),
		}

	case crc_inputPaymentCredentials:
		flags := m.Flags()
		_ = flags
		r = TL_inputPaymentCredentials{
			flags,
			m.Object(),
		}

	case crc_account_tmpPassword:
		r = TL_account_tmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case crc_shippingOption:
		r = TL_shippingOption{
			m.String(),
			m.String(),
			m.Vector(),
		}

	case crc_phoneCallAccepted:
		r = TL_phoneCallAccepted{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_inputMessagesFilterRoundVoice:
		r = TL_inputMessagesFilterRoundVoice{}

	case crc_inputMessagesFilterRoundVideo:
		r = TL_inputMessagesFilterRoundVideo{}

	case crc_upload_fileCdnRedirect:
		r = TL_upload_fileCdnRedirect{
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
			m.Vector(),
		}

	case crc_sendMessageRecordRoundAction:
		r = TL_sendMessageRecordRoundAction{}

	case crc_sendMessageUploadRoundAction:
		r = TL_sendMessageUploadRoundAction{
			m.Int(),
		}

	case crc_upload_cdnFileReuploadNeeded:
		r = TL_upload_cdnFileReuploadNeeded{
			m.StringBytes(),
		}

	case crc_upload_cdnFile:
		r = TL_upload_cdnFile{
			m.StringBytes(),
		}

	case crc_cdnPublicKey:
		r = TL_cdnPublicKey{
			m.Int(),
			m.String(),
		}

	case crc_cdnConfig:
		r = TL_cdnConfig{
			m.Vector(),
		}

	case crc_updateLangPackTooLong:
		r = TL_updateLangPackTooLong{}

	case crc_updateLangPack:
		r = TL_updateLangPack{
			m.Object(),
		}

	case crc_pageBlockChannel:
		r = TL_pageBlockChannel{
			m.Object(),
		}

	case crc_inputStickerSetItem:
		flags := m.Flags()
		_ = flags
		r = TL_inputStickerSetItem{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
		}

	case crc_langPackString:
		r = TL_langPackString{
			m.String(),
			m.String(),
		}

	case crc_langPackStringPluralized:
		flags := m.Flags()
		_ = flags
		r = TL_langPackStringPluralized{
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
		r = TL_langPackStringDeleted{
			m.String(),
		}

	case crc_langPackDifference:
		r = TL_langPackDifference{
			m.String(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_langPackLanguage:
		r = TL_langPackLanguage{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_channelParticipantAdmin:
		flags := m.Flags()
		_ = flags
		r = TL_channelParticipantAdmin{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_channelParticipantBanned:
		flags := m.Flags()
		_ = flags
		r = TL_channelParticipantBanned{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_channelParticipantsBanned:
		r = TL_channelParticipantsBanned{
			m.String(),
		}

	case crc_channelParticipantsSearch:
		r = TL_channelParticipantsSearch{
			m.String(),
		}

	case crc_topPeerCategoryPhoneCalls:
		r = TL_topPeerCategoryPhoneCalls{}

	case crc_pageBlockAudio:
		r = TL_pageBlockAudio{
			m.Long(),
			m.Object(),
		}

	case crc_channelAdminRights:
		flags := m.Flags()
		_ = flags
		r = TL_channelAdminRights{
			flags,
		}

	case crc_channelBannedRights:
		flags := m.Flags()
		_ = flags
		r = TL_channelBannedRights{
			flags,
			m.Int(),
		}

	case crc_channelAdminLogEventActionChangeTitle:
		r = TL_channelAdminLogEventActionChangeTitle{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeAbout:
		r = TL_channelAdminLogEventActionChangeAbout{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangeUsername:
		r = TL_channelAdminLogEventActionChangeUsername{
			m.String(),
			m.String(),
		}

	case crc_channelAdminLogEventActionChangePhoto:
		r = TL_channelAdminLogEventActionChangePhoto{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionToggleInvites:
		r = TL_channelAdminLogEventActionToggleInvites{
			m.Object(),
		}

	case crc_channelAdminLogEventActionToggleSignatures:
		r = TL_channelAdminLogEventActionToggleSignatures{
			m.Object(),
		}

	case crc_channelAdminLogEventActionUpdatePinned:
		r = TL_channelAdminLogEventActionUpdatePinned{
			m.Object(),
		}

	case crc_channelAdminLogEventActionEditMessage:
		r = TL_channelAdminLogEventActionEditMessage{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionDeleteMessage:
		r = TL_channelAdminLogEventActionDeleteMessage{
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantJoin:
		r = TL_channelAdminLogEventActionParticipantJoin{}

	case crc_channelAdminLogEventActionParticipantLeave:
		r = TL_channelAdminLogEventActionParticipantLeave{}

	case crc_channelAdminLogEventActionParticipantInvite:
		r = TL_channelAdminLogEventActionParticipantInvite{
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantToggleBan:
		r = TL_channelAdminLogEventActionParticipantToggleBan{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEventActionParticipantToggleAdmin:
		r = TL_channelAdminLogEventActionParticipantToggleAdmin{
			m.Object(),
			m.Object(),
		}

	case crc_channelAdminLogEvent:
		r = TL_channelAdminLogEvent{
			m.Long(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_channels_adminLogResults:
		r = TL_channels_adminLogResults{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channelAdminLogEventsFilter:
		flags := m.Flags()
		_ = flags
		r = TL_channelAdminLogEventsFilter{
			flags,
		}

	case crc_messageActionScreenshotTaken:
		r = TL_messageActionScreenshotTaken{}

	case crc_popularContact:
		r = TL_popularContact{
			m.Long(),
			m.Int(),
		}

	case crc_cdnFileHash:
		r = TL_cdnFileHash{
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputMessagesFilterMyMentions:
		r = TL_inputMessagesFilterMyMentions{}

	case crc_inputMessagesFilterMyMentionsUnread:
		r = TL_inputMessagesFilterMyMentionsUnread{}

	case crc_updateContactsReset:
		r = TL_updateContactsReset{}

	case crc_channelAdminLogEventActionChangeStickerSet:
		r = TL_channelAdminLogEventActionChangeStickerSet{
			m.Object(),
			m.Object(),
		}

	case crc_updateFavedStickers:
		r = TL_updateFavedStickers{}

	case crc_messages_favedStickers:
		r = TL_messages_favedStickers{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_favedStickersNotModified:
		r = TL_messages_favedStickersNotModified{}

	case crc_updateChannelReadMessagesContents:
		r = TL_updateChannelReadMessagesContents{
			m.Int(),
			m.VectorInt(),
		}

	case crc_invokeAfterMsg:
		r = TL_invokeAfterMsg{
			m.Long(),
			m.Object(),
		}

	case crc_invokeAfterMsgs:
		r = TL_invokeAfterMsgs{
			m.VectorLong(),
			m.Object(),
		}

	case crc_auth_checkPhone:
		r = TL_auth_checkPhone{
			m.String(),
		}

	case crc_auth_sendCode:
		flags := m.Flags()
		_ = flags
		r = TL_auth_sendCode{
			flags,
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Int(),
			m.String(),
		}

	case crc_auth_signUp:
		r = TL_auth_signUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_signIn:
		r = TL_auth_signIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_logOut:
		r = TL_auth_logOut{}

	case crc_auth_resetAuthorizations:
		r = TL_auth_resetAuthorizations{}

	case crc_auth_sendInvites:
		r = TL_auth_sendInvites{
			m.VectorString(),
			m.String(),
		}

	case crc_auth_exportAuthorization:
		r = TL_auth_exportAuthorization{
			m.Int(),
		}

	case crc_auth_importAuthorization:
		r = TL_auth_importAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_account_registerDevice:
		r = TL_account_registerDevice{
			m.Int(),
			m.String(),
		}

	case crc_account_unregisterDevice:
		r = TL_account_unregisterDevice{
			m.Int(),
			m.String(),
		}

	case crc_account_updateNotifySettings:
		r = TL_account_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_account_getNotifySettings:
		r = TL_account_getNotifySettings{
			m.Object(),
		}

	case crc_account_resetNotifySettings:
		r = TL_account_resetNotifySettings{}

	case crc_account_updateProfile:
		flags := m.Flags()
		_ = flags
		r = TL_account_updateProfile{
			flags,
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case crc_account_updateStatus:
		r = TL_account_updateStatus{
			m.Object(),
		}

	case crc_account_getWallPapers:
		r = TL_account_getWallPapers{}

	case crc_users_getUsers:
		r = TL_users_getUsers{
			m.Vector(),
		}

	case crc_users_getFullUser:
		r = TL_users_getFullUser{
			m.Object(),
		}

	case crc_contacts_getStatuses:
		r = TL_contacts_getStatuses{}

	case crc_contacts_getContacts:
		r = TL_contacts_getContacts{
			m.Int(),
		}

	case crc_contacts_importContacts:
		r = TL_contacts_importContacts{
			m.Vector(),
		}

	case crc_contacts_search:
		r = TL_contacts_search{
			m.String(),
			m.Int(),
		}

	case crc_contacts_deleteContact:
		r = TL_contacts_deleteContact{
			m.Object(),
		}

	case crc_contacts_deleteContacts:
		r = TL_contacts_deleteContacts{
			m.Vector(),
		}

	case crc_contacts_block:
		r = TL_contacts_block{
			m.Object(),
		}

	case crc_contacts_unblock:
		r = TL_contacts_unblock{
			m.Object(),
		}

	case crc_contacts_getBlocked:
		r = TL_contacts_getBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_messages_getMessages:
		r = TL_messages_getMessages{
			m.VectorInt(),
		}

	case crc_messages_getDialogs:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getDialogs{
			flags,
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_getHistory:
		r = TL_messages_getHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_search:
		flags := m.Flags()
		_ = flags
		r = TL_messages_search{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_readHistory:
		r = TL_messages_readHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteHistory:
		flags := m.Flags()
		_ = flags
		r = TL_messages_deleteHistory{
			flags,
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteMessages:
		flags := m.Flags()
		_ = flags
		r = TL_messages_deleteMessages{
			flags,
			m.VectorInt(),
		}

	case crc_messages_receivedMessages:
		r = TL_messages_receivedMessages{
			m.Int(),
		}

	case crc_messages_setTyping:
		r = TL_messages_setTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendMessage:
		flags := m.Flags()
		_ = flags
		r = TL_messages_sendMessage{
			flags,
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.String(),
			m.Long(),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
		}

	case crc_messages_sendMedia:
		flags := m.Flags()
		_ = flags
		r = TL_messages_sendMedia{
			flags,
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.Long(),
			m.FlaggedObject(flags, 2),
		}

	case crc_messages_forwardMessages:
		flags := m.Flags()
		_ = flags
		r = TL_messages_forwardMessages{
			flags,
			m.Object(),
			m.VectorInt(),
			m.VectorLong(),
			m.Object(),
		}

	case crc_messages_getChats:
		r = TL_messages_getChats{
			m.VectorInt(),
		}

	case crc_messages_getFullChat:
		r = TL_messages_getFullChat{
			m.Int(),
		}

	case crc_messages_editChatTitle:
		r = TL_messages_editChatTitle{
			m.Int(),
			m.String(),
		}

	case crc_messages_editChatPhoto:
		r = TL_messages_editChatPhoto{
			m.Int(),
			m.Object(),
		}

	case crc_messages_addChatUser:
		r = TL_messages_addChatUser{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteChatUser:
		r = TL_messages_deleteChatUser{
			m.Int(),
			m.Object(),
		}

	case crc_messages_createChat:
		r = TL_messages_createChat{
			m.Vector(),
			m.String(),
		}

	case crc_updates_getState:
		r = TL_updates_getState{}

	case crc_updates_getDifference:
		flags := m.Flags()
		_ = flags
		r = TL_updates_getDifference{
			flags,
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
		}

	case crc_photos_updateProfilePhoto:
		r = TL_photos_updateProfilePhoto{
			m.Object(),
		}

	case crc_photos_uploadProfilePhoto:
		r = TL_photos_uploadProfilePhoto{
			m.Object(),
		}

	case crc_upload_saveFilePart:
		r = TL_upload_saveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_upload_getFile:
		r = TL_upload_getFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_help_getConfig:
		r = TL_help_getConfig{}

	case crc_help_getNearestDc:
		r = TL_help_getNearestDc{}

	case crc_help_getAppUpdate:
		r = TL_help_getAppUpdate{}

	case crc_help_saveAppLog:
		r = TL_help_saveAppLog{
			m.Vector(),
		}

	case crc_help_getInviteText:
		r = TL_help_getInviteText{}

	case crc_photos_deletePhotos:
		r = TL_photos_deletePhotos{
			m.Vector(),
		}

	case crc_photos_getUserPhotos:
		r = TL_photos_getUserPhotos{
			m.Object(),
			m.Int(),
			m.Long(),
			m.Int(),
		}

	case crc_messages_forwardMessage:
		r = TL_messages_forwardMessage{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case crc_messages_getDhConfig:
		r = TL_messages_getDhConfig{
			m.Int(),
			m.Int(),
		}

	case crc_messages_requestEncryption:
		r = TL_messages_requestEncryption{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_acceptEncryption:
		r = TL_messages_acceptEncryption{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_messages_discardEncryption:
		r = TL_messages_discardEncryption{
			m.Int(),
		}

	case crc_messages_setEncryptedTyping:
		r = TL_messages_setEncryptedTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_readEncryptedHistory:
		r = TL_messages_readEncryptedHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_sendEncrypted:
		r = TL_messages_sendEncrypted{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_sendEncryptedFile:
		r = TL_messages_sendEncryptedFile{
			m.Object(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_messages_sendEncryptedService:
		r = TL_messages_sendEncryptedService{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_receivedQueue:
		r = TL_messages_receivedQueue{
			m.Int(),
		}

	case crc_upload_saveBigFilePart:
		r = TL_upload_saveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_initConnection:
		r = TL_initConnection{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_help_getSupport:
		r = TL_help_getSupport{}

	case crc_auth_bindTempAuthKey:
		r = TL_auth_bindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_contacts_exportCard:
		r = TL_contacts_exportCard{}

	case crc_contacts_importCard:
		r = TL_contacts_importCard{
			m.VectorInt(),
		}

	case crc_messages_readMessageContents:
		r = TL_messages_readMessageContents{
			m.VectorInt(),
		}

	case crc_account_checkUsername:
		r = TL_account_checkUsername{
			m.String(),
		}

	case crc_account_updateUsername:
		r = TL_account_updateUsername{
			m.String(),
		}

	case crc_account_getPrivacy:
		r = TL_account_getPrivacy{
			m.Object(),
		}

	case crc_account_setPrivacy:
		r = TL_account_setPrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_account_deleteAccount:
		r = TL_account_deleteAccount{
			m.String(),
		}

	case crc_account_getAccountTTL:
		r = TL_account_getAccountTTL{}

	case crc_account_setAccountTTL:
		r = TL_account_setAccountTTL{
			m.Object(),
		}

	case crc_invokeWithLayer:
		r = TL_invokeWithLayer{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_resolveUsername:
		r = TL_contacts_resolveUsername{
			m.String(),
		}

	case crc_account_sendChangePhoneCode:
		flags := m.Flags()
		_ = flags
		r = TL_account_sendChangePhoneCode{
			flags,
			m.String(),
			m.FlaggedObject(flags, 0),
		}

	case crc_account_changePhone:
		r = TL_account_changePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messages_getAllStickers:
		r = TL_messages_getAllStickers{
			m.Int(),
		}

	case crc_account_updateDeviceLocked:
		r = TL_account_updateDeviceLocked{
			m.Int(),
		}

	case crc_account_getPassword:
		r = TL_account_getPassword{}

	case crc_auth_checkPassword:
		r = TL_auth_checkPassword{
			m.StringBytes(),
		}

	case crc_messages_getWebPagePreview:
		r = TL_messages_getWebPagePreview{
			m.String(),
		}

	case crc_account_getAuthorizations:
		r = TL_account_getAuthorizations{}

	case crc_account_resetAuthorization:
		r = TL_account_resetAuthorization{
			m.Long(),
		}

	case crc_account_getPasswordSettings:
		r = TL_account_getPasswordSettings{
			m.StringBytes(),
		}

	case crc_account_updatePasswordSettings:
		r = TL_account_updatePasswordSettings{
			m.StringBytes(),
			m.Object(),
		}

	case crc_auth_requestPasswordRecovery:
		r = TL_auth_requestPasswordRecovery{}

	case crc_auth_recoverPassword:
		r = TL_auth_recoverPassword{
			m.String(),
		}

	case crc_invokeWithoutUpdates:
		r = TL_invokeWithoutUpdates{
			m.Object(),
		}

	case crc_messages_exportChatInvite:
		r = TL_messages_exportChatInvite{
			m.Int(),
		}

	case crc_messages_checkChatInvite:
		r = TL_messages_checkChatInvite{
			m.String(),
		}

	case crc_messages_importChatInvite:
		r = TL_messages_importChatInvite{
			m.String(),
		}

	case crc_messages_getStickerSet:
		r = TL_messages_getStickerSet{
			m.Object(),
		}

	case crc_messages_installStickerSet:
		r = TL_messages_installStickerSet{
			m.Object(),
			m.Object(),
		}

	case crc_messages_uninstallStickerSet:
		r = TL_messages_uninstallStickerSet{
			m.Object(),
		}

	case crc_auth_importBotAuthorization:
		r = TL_auth_importBotAuthorization{
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_messages_startBot:
		r = TL_messages_startBot{
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
		}

	case crc_help_getAppChangelog:
		r = TL_help_getAppChangelog{
			m.String(),
		}

	case crc_messages_reportSpam:
		r = TL_messages_reportSpam{
			m.Object(),
		}

	case crc_messages_getMessagesViews:
		r = TL_messages_getMessagesViews{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case crc_updates_getChannelDifference:
		flags := m.Flags()
		_ = flags
		r = TL_updates_getChannelDifference{
			flags,
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_channels_readHistory:
		r = TL_channels_readHistory{
			m.Object(),
			m.Int(),
		}

	case crc_channels_deleteMessages:
		r = TL_channels_deleteMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_deleteUserHistory:
		r = TL_channels_deleteUserHistory{
			m.Object(),
			m.Object(),
		}

	case crc_channels_reportSpam:
		r = TL_channels_reportSpam{
			m.Object(),
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_getMessages:
		r = TL_channels_getMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_getParticipants:
		r = TL_channels_getParticipants{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_channels_getParticipant:
		r = TL_channels_getParticipant{
			m.Object(),
			m.Object(),
		}

	case crc_channels_getChannels:
		r = TL_channels_getChannels{
			m.Vector(),
		}

	case crc_channels_getFullChannel:
		r = TL_channels_getFullChannel{
			m.Object(),
		}

	case crc_channels_createChannel:
		flags := m.Flags()
		_ = flags
		r = TL_channels_createChannel{
			flags,
			m.String(),
			m.String(),
		}

	case crc_channels_editAbout:
		r = TL_channels_editAbout{
			m.Object(),
			m.String(),
		}

	case crc_channels_editAdmin:
		r = TL_channels_editAdmin{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_channels_editTitle:
		r = TL_channels_editTitle{
			m.Object(),
			m.String(),
		}

	case crc_channels_editPhoto:
		r = TL_channels_editPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_channels_checkUsername:
		r = TL_channels_checkUsername{
			m.Object(),
			m.String(),
		}

	case crc_channels_updateUsername:
		r = TL_channels_updateUsername{
			m.Object(),
			m.String(),
		}

	case crc_channels_joinChannel:
		r = TL_channels_joinChannel{
			m.Object(),
		}

	case crc_channels_leaveChannel:
		r = TL_channels_leaveChannel{
			m.Object(),
		}

	case crc_channels_inviteToChannel:
		r = TL_channels_inviteToChannel{
			m.Object(),
			m.Vector(),
		}

	case crc_channels_exportInvite:
		r = TL_channels_exportInvite{
			m.Object(),
		}

	case crc_channels_deleteChannel:
		r = TL_channels_deleteChannel{
			m.Object(),
		}

	case crc_messages_toggleChatAdmins:
		r = TL_messages_toggleChatAdmins{
			m.Int(),
			m.Object(),
		}

	case crc_messages_editChatAdmin:
		r = TL_messages_editChatAdmin{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_migrateChat:
		r = TL_messages_migrateChat{
			m.Int(),
		}

	case crc_messages_searchGlobal:
		r = TL_messages_searchGlobal{
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_account_reportPeer:
		r = TL_account_reportPeer{
			m.Object(),
			m.Object(),
		}

	case crc_messages_reorderStickerSets:
		flags := m.Flags()
		_ = flags
		r = TL_messages_reorderStickerSets{
			flags,
			m.VectorLong(),
		}

	case crc_help_getTermsOfService:
		r = TL_help_getTermsOfService{}

	case crc_messages_getDocumentByHash:
		r = TL_messages_getDocumentByHash{
			m.StringBytes(),
			m.Int(),
			m.String(),
		}

	case crc_messages_searchGifs:
		r = TL_messages_searchGifs{
			m.String(),
			m.Int(),
		}

	case crc_messages_getSavedGifs:
		r = TL_messages_getSavedGifs{
			m.Int(),
		}

	case crc_messages_saveGif:
		r = TL_messages_saveGif{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getInlineBotResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getInlineBotResults{
			flags,
			m.Object(),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.String(),
		}

	case crc_messages_setInlineBotResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setInlineBotResults{
			flags,
			m.Long(),
			m.Vector(),
			m.Int(),
			m.FlaggedString(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case crc_messages_sendInlineBotResult:
		flags := m.Flags()
		_ = flags
		r = TL_messages_sendInlineBotResult{
			flags,
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Long(),
			m.Long(),
			m.String(),
		}

	case crc_channels_toggleInvites:
		r = TL_channels_toggleInvites{
			m.Object(),
			m.Object(),
		}

	case crc_channels_exportMessageLink:
		r = TL_channels_exportMessageLink{
			m.Object(),
			m.Int(),
		}

	case crc_channels_toggleSignatures:
		r = TL_channels_toggleSignatures{
			m.Object(),
			m.Object(),
		}

	case crc_messages_hideReportSpam:
		r = TL_messages_hideReportSpam{
			m.Object(),
		}

	case crc_messages_getPeerSettings:
		r = TL_messages_getPeerSettings{
			m.Object(),
		}

	case crc_channels_updatePinnedMessage:
		flags := m.Flags()
		_ = flags
		r = TL_channels_updatePinnedMessage{
			flags,
			m.Object(),
			m.Int(),
		}

	case crc_auth_resendCode:
		r = TL_auth_resendCode{
			m.String(),
			m.String(),
		}

	case crc_auth_cancelCode:
		r = TL_auth_cancelCode{
			m.String(),
			m.String(),
		}

	case crc_messages_getMessageEditData:
		r = TL_messages_getMessageEditData{
			m.Object(),
			m.Int(),
		}

	case crc_messages_editMessage:
		flags := m.Flags()
		_ = flags
		r = TL_messages_editMessage{
			flags,
			m.Object(),
			m.Int(),
			m.FlaggedString(flags, 11),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
		}

	case crc_messages_editInlineBotMessage:
		flags := m.Flags()
		_ = flags
		r = TL_messages_editInlineBotMessage{
			flags,
			m.Object(),
			m.FlaggedString(flags, 11),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
		}

	case crc_messages_getBotCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getBotCallbackAnswer{
			flags,
			m.Object(),
			m.Int(),
			m.FlaggedStringBytes(flags, 0),
		}

	case crc_messages_setBotCallbackAnswer:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setBotCallbackAnswer{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case crc_contacts_getTopPeers:
		flags := m.Flags()
		_ = flags
		r = TL_contacts_getTopPeers{
			flags,
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_contacts_resetTopPeerRating:
		r = TL_contacts_resetTopPeerRating{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getPeerDialogs:
		r = TL_messages_getPeerDialogs{
			m.Vector(),
		}

	case crc_messages_saveDraft:
		flags := m.Flags()
		_ = flags
		r = TL_messages_saveDraft{
			flags,
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.String(),
			m.FlaggedVector(flags, 3),
		}

	case crc_messages_getAllDrafts:
		r = TL_messages_getAllDrafts{}

	case crc_account_sendConfirmPhoneCode:
		flags := m.Flags()
		_ = flags
		r = TL_account_sendConfirmPhoneCode{
			flags,
			m.String(),
			m.FlaggedObject(flags, 0),
		}

	case crc_account_confirmPhone:
		r = TL_account_confirmPhone{
			m.String(),
			m.String(),
		}

	case crc_messages_getFeaturedStickers:
		r = TL_messages_getFeaturedStickers{
			m.Int(),
		}

	case crc_messages_readFeaturedStickers:
		r = TL_messages_readFeaturedStickers{
			m.VectorLong(),
		}

	case crc_messages_getRecentStickers:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getRecentStickers{
			flags,
			m.Int(),
		}

	case crc_messages_saveRecentSticker:
		flags := m.Flags()
		_ = flags
		r = TL_messages_saveRecentSticker{
			flags,
			m.Object(),
			m.Object(),
		}

	case crc_messages_clearRecentStickers:
		flags := m.Flags()
		_ = flags
		r = TL_messages_clearRecentStickers{
			flags,
		}

	case crc_messages_getArchivedStickers:
		flags := m.Flags()
		_ = flags
		r = TL_messages_getArchivedStickers{
			flags,
			m.Long(),
			m.Int(),
		}

	case crc_channels_getAdminedPublicChannels:
		r = TL_channels_getAdminedPublicChannels{}

	case crc_auth_dropTempAuthKeys:
		r = TL_auth_dropTempAuthKeys{
			m.VectorLong(),
		}

	case crc_messages_setGameScore:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setGameScore{
			flags,
			m.Object(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_setInlineGameScore:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setInlineGameScore{
			flags,
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_getMaskStickers:
		r = TL_messages_getMaskStickers{
			m.Int(),
		}

	case crc_messages_getAttachedStickers:
		r = TL_messages_getAttachedStickers{
			m.Object(),
		}

	case crc_messages_getGameHighScores:
		r = TL_messages_getGameHighScores{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_messages_getInlineGameHighScores:
		r = TL_messages_getInlineGameHighScores{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getCommonChats:
		r = TL_messages_getCommonChats{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_getAllChats:
		r = TL_messages_getAllChats{
			m.VectorInt(),
		}

	case crc_help_setBotUpdatesStatus:
		r = TL_help_setBotUpdatesStatus{
			m.Int(),
			m.String(),
		}

	case crc_messages_getWebPage:
		r = TL_messages_getWebPage{
			m.String(),
			m.Int(),
		}

	case crc_messages_toggleDialogPin:
		flags := m.Flags()
		_ = flags
		r = TL_messages_toggleDialogPin{
			flags,
			m.Object(),
		}

	case crc_messages_reorderPinnedDialogs:
		flags := m.Flags()
		_ = flags
		r = TL_messages_reorderPinnedDialogs{
			flags,
			m.Vector(),
		}

	case crc_messages_getPinnedDialogs:
		r = TL_messages_getPinnedDialogs{}

	case crc_phone_requestCall:
		r = TL_phone_requestCall{
			m.Object(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phone_acceptCall:
		r = TL_phone_acceptCall{
			m.Object(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_phone_discardCall:
		r = TL_phone_discardCall{
			m.Object(),
			m.Int(),
			m.Object(),
			m.Long(),
		}

	case crc_phone_receivedCall:
		r = TL_phone_receivedCall{
			m.Object(),
		}

	case crc_messages_reportEncryptedSpam:
		r = TL_messages_reportEncryptedSpam{
			m.Object(),
		}

	case crc_payments_getPaymentForm:
		r = TL_payments_getPaymentForm{
			m.Int(),
		}

	case crc_payments_sendPaymentForm:
		flags := m.Flags()
		_ = flags
		r = TL_payments_sendPaymentForm{
			flags,
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.Object(),
		}

	case crc_account_getTmpPassword:
		r = TL_account_getTmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case crc_messages_setBotShippingResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setBotShippingResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedVector(flags, 1),
		}

	case crc_messages_setBotPrecheckoutResults:
		flags := m.Flags()
		_ = flags
		r = TL_messages_setBotPrecheckoutResults{
			flags,
			m.Long(),
			m.FlaggedString(flags, 0),
		}

	case crc_upload_getWebFile:
		r = TL_upload_getWebFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_bots_sendCustomRequest:
		r = TL_bots_sendCustomRequest{
			m.String(),
			m.Object(),
		}

	case crc_bots_answerWebhookJSONQuery:
		r = TL_bots_answerWebhookJSONQuery{
			m.Long(),
			m.Object(),
		}

	case crc_payments_getPaymentReceipt:
		r = TL_payments_getPaymentReceipt{
			m.Int(),
		}

	case crc_payments_validateRequestedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_payments_validateRequestedInfo{
			flags,
			m.Int(),
			m.Object(),
		}

	case crc_payments_getSavedInfo:
		r = TL_payments_getSavedInfo{}

	case crc_payments_clearSavedInfo:
		flags := m.Flags()
		_ = flags
		r = TL_payments_clearSavedInfo{
			flags,
		}

	case crc_phone_getCallConfig:
		r = TL_phone_getCallConfig{}

	case crc_phone_confirmCall:
		r = TL_phone_confirmCall{
			m.Object(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
		}

	case crc_phone_setCallRating:
		r = TL_phone_setCallRating{
			m.Object(),
			m.Int(),
			m.String(),
		}

	case crc_phone_saveCallDebug:
		r = TL_phone_saveCallDebug{
			m.Object(),
			m.Object(),
		}

	case crc_upload_getCdnFile:
		r = TL_upload_getCdnFile{
			m.StringBytes(),
			m.Int(),
			m.Int(),
		}

	case crc_upload_reuploadCdnFile:
		r = TL_upload_reuploadCdnFile{
			m.StringBytes(),
			m.StringBytes(),
		}

	case crc_help_getCdnConfig:
		r = TL_help_getCdnConfig{}

	case crc_messages_uploadMedia:
		r = TL_messages_uploadMedia{
			m.Object(),
			m.Object(),
		}

	case crc_stickers_createStickerSet:
		flags := m.Flags()
		_ = flags
		r = TL_stickers_createStickerSet{
			flags,
			m.Object(),
			m.String(),
			m.String(),
			m.Vector(),
		}

	case crc_langpack_getLangPack:
		r = TL_langpack_getLangPack{
			m.String(),
		}

	case crc_langpack_getStrings:
		r = TL_langpack_getStrings{
			m.String(),
			m.VectorString(),
		}

	case crc_langpack_getDifference:
		r = TL_langpack_getDifference{
			m.Int(),
		}

	case crc_langpack_getLanguages:
		r = TL_langpack_getLanguages{}

	case crc_channels_editBanned:
		r = TL_channels_editBanned{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_channels_getAdminLog:
		flags := m.Flags()
		_ = flags
		r = TL_channels_getAdminLog{
			flags,
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedVector(flags, 1),
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crc_stickers_removeStickerFromSet:
		r = TL_stickers_removeStickerFromSet{
			m.Object(),
		}

	case crc_stickers_changeStickerPosition:
		r = TL_stickers_changeStickerPosition{
			m.Object(),
			m.Int(),
		}

	case crc_stickers_addStickerToSet:
		r = TL_stickers_addStickerToSet{
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendScreenshotNotification:
		r = TL_messages_sendScreenshotNotification{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case crc_upload_getCdnFileHashes:
		r = TL_upload_getCdnFileHashes{
			m.StringBytes(),
			m.Int(),
		}

	case crc_messages_getUnreadMentions:
		r = TL_messages_getUnreadMentions{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_faveSticker:
		r = TL_messages_faveSticker{
			m.Object(),
			m.Object(),
		}

	case crc_channels_setStickers:
		r = TL_channels_setStickers{
			m.Object(),
			m.Object(),
		}

	case crc_contacts_resetSaved:
		r = TL_contacts_resetSaved{}

	case crc_messages_getFavedStickers:
		r = TL_messages_getFavedStickers{
			m.Int(),
		}

	case crc_channels_readMessageContents:
		r = TL_channels_readMessageContents{
			m.Object(),
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
