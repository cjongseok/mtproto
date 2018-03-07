package proto
import "fmt"
const (
crc_boolFalse = 0xbc799737
crc_boolTrue = 0x997275b5
crc_error = 0xc4b9f9bb
crc_null = 0x56730bcc
crc_vector = 0x1cb5c415
crc_inputPeerEmpty = 0x7f3b18ea
crc_inputPeerSelf = 0x7da07ec9
crc_inputPeerChat = 0x179be863
crc_inputUserEmpty = 0xb98886cf
crc_inputUserSelf = 0xf7c1b13f
crc_inputPhoneContact = 0xf392b7f4
crc_inputFile = 0xf52ff27f
crc_inputMediaEmpty = 0x9664f57f
crc_inputMediaUploadedPhoto = 0x2f37e231
crc_inputMediaPhoto = 0x81fa373a
crc_inputMediaGeoPoint = 0xf9c44144
crc_inputMediaContact = 0xa6e45987
crc_inputChatPhotoEmpty = 0x1ca48f57
crc_inputChatUploadedPhoto = 0x927c55b4
crc_inputChatPhoto = 0x8953ad37
crc_inputGeoPointEmpty = 0xe4c123d6
crc_inputGeoPoint = 0xf3b7acc9
crc_inputPhotoEmpty = 0x1cd7bf0d
crc_inputPhoto = 0xfb95c6c4
crc_inputFileLocation = 0x14637196
crc_inputAppEvent = 0x770656a8
crc_peerUser = 0x9db1bc6d
crc_peerChat = 0xbad0e5bb
crc_storageFileUnknown = 0xaa963b05
crc_storageFileJpeg = 0x007efe0e
crc_storageFileGif = 0xcae1aadf
crc_storageFilePng = 0x0a4f63c0
crc_storageFileMp3 = 0x528a0677
crc_storageFileMov = 0x4b09ebbc
crc_storageFilePartial = 0x40bc6f52
crc_storageFileMp4 = 0xb3cea0e4
crc_storageFileWebp = 0x1081464c
crc_fileLocationUnavailable = 0x7c596b46
crc_fileLocation = 0x53d69076
crc_userEmpty = 0x200250ba
crc_userProfilePhotoEmpty = 0x4f11bae1
crc_userProfilePhoto = 0xd559d8c8
crc_userStatusEmpty = 0x09d05049
crc_userStatusOnline = 0xedb93949
crc_userStatusOffline = 0x008c703f
crc_chatEmpty = 0x9ba2d800
crc_chat = 0xd91cdd54
crc_chatForbidden = 0x07328bdb
crc_chatFull = 0x2e02a614
crc_chatParticipant = 0xc8d7493e
crc_chatParticipantsForbidden = 0xfc900c2b
crc_chatParticipants = 0x3f460fed
crc_chatPhotoEmpty = 0x37c1011c
crc_chatPhoto = 0x6153276a
crc_messageEmpty = 0x83e5de54
crc_message = 0x90dddc11
crc_messageService = 0x9e19a1f6
crc_messageMediaEmpty = 0x3ded6320
crc_messageMediaPhoto = 0xb5223b0f
crc_messageMediaGeo = 0x56e0d474
crc_messageMediaContact = 0x5e7d2f39
crc_messageMediaUnsupported = 0x9f84f49e
crc_messageActionEmpty = 0xb6aef7b0
crc_messageActionChatCreate = 0xa6638b9a
crc_messageActionChatEditTitle = 0xb5a1ce5a
crc_messageActionChatEditPhoto = 0x7fcb13a8
crc_messageActionChatDeletePhoto = 0x95e3fbef
crc_messageActionChatAddUser = 0x488a7337
crc_messageActionChatDeleteUser = 0xb2ae9b0c
crc_dialog = 0xe4def5db
crc_photoEmpty = 0x2331b22d
crc_photo = 0x9288dd29
crc_photoSizeEmpty = 0x0e17e23c
crc_photoSize = 0x77bfb61b
crc_photoCachedSize = 0xe9a734fa
crc_geoPointEmpty = 0x1117dd5f
crc_geoPoint = 0x2049d70c
crc_authCheckedPhone = 0x811ea28e
crc_authSentCode = 0x5e002502
crc_authAuthorization = 0xcd050916
crc_authExportedAuthorization = 0xdf969c2d
crc_inputNotifyPeer = 0xb8bc5b0c
crc_inputNotifyUsers = 0x193b4417
crc_inputNotifyChats = 0x4a95e84e
crc_inputNotifyAll = 0xa429b886
crc_inputPeerNotifySettings = 0x38935eb2
crc_peerNotifyEventsEmpty = 0xadd53cb3
crc_peerNotifyEventsAll = 0x6d1ded88
crc_peerNotifySettingsEmpty = 0x70a68512
crc_peerNotifySettings = 0x9acda4c0
crc_wallPaper = 0xccb03657
crc_userFull = 0x0f220f3f
crc_contact = 0xf911c994
crc_importedContact = 0xd0028438
crc_contactBlocked = 0x561bc879
crc_contactStatus = 0xd3680c61
crc_contactsLink = 0x3ace484c
crc_contactsContacts = 0xeae87e42
crc_contactsContactsNotModified = 0xb74ba9d2
crc_contactsImportedContacts = 0x77d01c3b
crc_contactsBlocked = 0x1c138d15
crc_contactsBlockedSlice = 0x900802a1
crc_contactsFound = 0x1aa1f784
crc_messagesDialogs = 0x15ba6c40
crc_messagesDialogsSlice = 0x71e094f3
crc_messagesMessages = 0x8c718e87
crc_messagesMessagesSlice = 0x0b446ae3
crc_messagesChats = 0x64ff9fd5
crc_messagesChatFull = 0xe5d7d19c
crc_messagesAffectedHistory = 0xb45c69d1
crc_inputMessagesFilterEmpty = 0x57e2f66c
crc_inputMessagesFilterPhotos = 0x9609a51c
crc_inputMessagesFilterVideo = 0x9fc00e65
crc_inputMessagesFilterPhotoVideo = 0x56e9f0e4
crc_updateNewMessage = 0x1f2b0afd
crc_updateMessageID = 0x4e90bfd6
crc_updateDeleteMessages = 0xa20db0e5
crc_updateUserTyping = 0x5c486927
crc_updateChatUserTyping = 0x9a65ea1f
crc_updateChatParticipants = 0x07761198
crc_updateUserStatus = 0x1bfbd823
crc_updateUserName = 0xa7332b73
crc_updateUserPhoto = 0x95313b0c
crc_updateContactRegistered = 0x2575bbb9
crc_updateContactLink = 0x9d2e67c5
crc_updatesState = 0xa56c2a3e
crc_updatesDifferenceEmpty = 0x5d75a138
crc_updatesDifference = 0x00f49ca0
crc_updatesDifferenceSlice = 0xa8fb1981
crc_updatesTooLong = 0xe317af7e
crc_updateShortMessage = 0x914fbf11
crc_updateShortChatMessage = 0x16812688
crc_updateShort = 0x78d4dec1
crc_updatesCombined = 0x725b04c3
crc_updates = 0x74ae4240
crc_photosPhoto = 0x20212ca8
crc_uploadFile = 0x096a18d5
crc_dcOption = 0x05d8c6cc
crc_config = 0x8df376a4
crc_nearestDc = 0x8e1a1775
crc_helpAppUpdate = 0x8987f311
crc_helpNoAppUpdate = 0xc45a6536
crc_helpInviteText = 0x18cb9f78
crc_inputPeerNotifyEventsEmpty = 0xf03064d8
crc_inputPeerNotifyEventsAll = 0xe86a2c74
crc_photosPhotos = 0x8dca6aa5
crc_photosPhotosSlice = 0x15051f54
crc_wallPaperSolid = 0x63117f24
crc_updateNewEncryptedMessage = 0x12bcbd9a
crc_updateEncryptedChatTyping = 0x1710f156
crc_updateEncryption = 0xb4a2e88d
crc_updateEncryptedMessagesRead = 0x38fe25b7
crc_encryptedChatEmpty = 0xab7ec0a0
crc_encryptedChatWaiting = 0x3bf703dc
crc_encryptedChatRequested = 0xc878527e
crc_encryptedChat = 0xfa56ce36
crc_encryptedChatDiscarded = 0x13d6dd27
crc_inputEncryptedChat = 0xf141b5e1
crc_encryptedFileEmpty = 0xc21f497e
crc_encryptedFile = 0x4a70994c
crc_inputEncryptedFileEmpty = 0x1837c364
crc_inputEncryptedFileUploaded = 0x64bd0306
crc_inputEncryptedFile = 0x5a17b5e5
crc_inputEncryptedFileLocation = 0xf5235d55
crc_encryptedMessage = 0xed18c118
crc_encryptedMessageService = 0x23734b06
crc_messagesDhConfigNotModified = 0xc0e24635
crc_messagesDhConfig = 0x2c221edd
crc_messagesSentEncryptedMessage = 0x560f8935
crc_messagesSentEncryptedFile = 0x9493ff32
crc_inputFileBig = 0xfa4f0bb5
crc_inputEncryptedFileBigUploaded = 0x2dc173c8
crc_storageFilePdf = 0xae1e508d
crc_inputMessagesFilterDocument = 0x9eddf188
crc_inputMessagesFilterPhotoVideoDocuments = 0xd95e73bb
crc_updateChatParticipantAdd = 0xea4b0e5c
crc_updateChatParticipantDelete = 0x6e5f8c22
crc_updateDcOptions = 0x8e5e9873
crc_inputMediaUploadedDocument = 0xe39621fd
crc_inputMediaDocument = 0x5acb668e
crc_messageMediaDocument = 0x7c4414d3
crc_inputDocumentEmpty = 0x72f0eaae
crc_inputDocument = 0x18798952
crc_inputDocumentFileLocation = 0x430f0724
crc_documentEmpty = 0x36f8c871
crc_document = 0x87232bc7
crc_helpSupport = 0x17c6b5f6
crc_notifyAll = 0x74d07c60
crc_notifyChats = 0xc007cec3
crc_notifyPeer = 0x9fd40bd8
crc_notifyUsers = 0xb4c83b4c
crc_updateUserBlocked = 0x80ece81a
crc_updateNotifySettings = 0xbec268ef
crc_sendMessageTypingAction = 0x16bf744e
crc_sendMessageCancelAction = 0xfd5ec8f5
crc_sendMessageRecordVideoAction = 0xa187d66f
crc_sendMessageUploadVideoAction = 0xe9763aec
crc_sendMessageRecordAudioAction = 0xd52f73f7
crc_sendMessageUploadAudioAction = 0xf351d7ab
crc_sendMessageUploadPhotoAction = 0xd1d34a26
crc_sendMessageUploadDocumentAction = 0xaa0cd9e4
crc_sendMessageGeoLocationAction = 0x176f8ba1
crc_sendMessageChooseContactAction = 0x628cbc6f
crc_updateServiceNotification = 0xebe46819
crc_userStatusRecently = 0xe26f42f1
crc_userStatusLastWeek = 0x07bf09fc
crc_userStatusLastMonth = 0x77ebc742
crc_updatePrivacy = 0xee3b272a
crc_inputPrivacyKeyStatusTimestamp = 0x4f96cb18
crc_privacyKeyStatusTimestamp = 0xbc2eab30
crc_inputPrivacyValueAllowContacts = 0x0d09e07b
crc_inputPrivacyValueAllowAll = 0x184b35ce
crc_inputPrivacyValueAllowUsers = 0x131cc67f
crc_inputPrivacyValueDisallowContacts = 0x0ba52007
crc_inputPrivacyValueDisallowAll = 0xd66b66c9
crc_inputPrivacyValueDisallowUsers = 0x90110467
crc_privacyValueAllowContacts = 0xfffe1bac
crc_privacyValueAllowAll = 0x65427b82
crc_privacyValueAllowUsers = 0x4d5bbe0c
crc_privacyValueDisallowContacts = 0xf888fa1a
crc_privacyValueDisallowAll = 0x8b73e763
crc_privacyValueDisallowUsers = 0x0c7f49b7
crc_accountPrivacyRules = 0x554abb6f
crc_accountDaysTTL = 0xb8d0afdf
crc_updateUserPhone = 0x12b9417b
crc_disabledFeature = 0xae636f24
crc_documentAttributeImageSize = 0x6c37c15c
crc_documentAttributeAnimated = 0x11b58939
crc_documentAttributeSticker = 0x6319d612
crc_documentAttributeVideo = 0x0ef02ce6
crc_documentAttributeAudio = 0x9852f9c6
crc_documentAttributeFilename = 0x15590068
crc_messagesStickersNotModified = 0xf1749a22
crc_messagesStickers = 0x8a8ecd32
crc_stickerPack = 0x12b299d4
crc_messagesAllStickersNotModified = 0xe86602c3
crc_messagesAllStickers = 0xedfd405f
crc_accountNoPassword = 0x96dabc18
crc_accountPassword = 0x7c18141c
crc_updateReadHistoryInbox = 0x9961fd5c
crc_updateReadHistoryOutbox = 0x2f2f21bf
crc_messagesAffectedMessages = 0x84d19185
crc_contactLinkUnknown = 0x5f4f9247
crc_contactLinkNone = 0xfeedd3ad
crc_contactLinkHasPhone = 0x268f3f59
crc_contactLinkContact = 0xd502c2d0
crc_updateWebPage = 0x7f891213
crc_webPageEmpty = 0xeb1477e8
crc_webPagePending = 0xc586da1c
crc_webPage = 0x5f07b4bc
crc_messageMediaWebPage = 0xa32dd600
crc_authorization = 0x7bf2e6f6
crc_accountAuthorizations = 0x1250abde
crc_accountPasswordSettings = 0xb7b72ab3
crc_accountPasswordInputSettings = 0x86916deb
crc_authPasswordRecovery = 0x137948a5
crc_inputMediaVenue = 0x2827a81a
crc_messageMediaVenue = 0x7912b71f
crc_receivedNotifyMessage = 0xa384b779
crc_chatInviteEmpty = 0x69df3769
crc_chatInviteExported = 0xfc2e05bc
crc_chatInviteAlready = 0x5a686d7c
crc_chatInvite = 0xdb74f558
crc_messageActionChatJoinedByLink = 0xf89cf5e8
crc_updateReadMessagesContents = 0x68c13933
crc_inputStickerSetEmpty = 0xffb62b95
crc_inputStickerSetID = 0x9de7a269
crc_inputStickerSetShortName = 0x861cc8a0
crc_stickerSet = 0xcd303b41
crc_messagesStickerSet = 0xb60a24a6
crc_user = 0x2e13f4c3
crc_botCommand = 0xc27ac8c7
crc_botInfo = 0x98e81d3a
crc_keyboardButton = 0xa2fa4880
crc_keyboardButtonRow = 0x77608b83
crc_replyKeyboardHide = 0xa03e5b85
crc_replyKeyboardForceReply = 0xf4108aa0
crc_replyKeyboardMarkup = 0x3502758c
crc_inputMessagesFilterUrl = 0x7ef0dd87
crc_inputPeerUser = 0x7b8e7de6
crc_inputUser = 0xd8292816
crc_messageEntityUnknown = 0xbb92ba95
crc_messageEntityMention = 0xfa04579d
crc_messageEntityHashtag = 0x6f635b0d
crc_messageEntityBotCommand = 0x6cef8ac7
crc_messageEntityUrl = 0x6ed02538
crc_messageEntityEmail = 0x64e475c2
crc_messageEntityBold = 0xbd610bc9
crc_messageEntityItalic = 0x826f8b60
crc_messageEntityCode = 0x28a20571
crc_messageEntityPre = 0x73924be0
crc_messageEntityTextUrl = 0x76a6d327
crc_updateShortSentMessage = 0x11f1331c
crc_inputPeerChannel = 0x20adaef8
crc_peerChannel = 0xbddde532
crc_channel = 0x0cb44b1c
crc_channelForbidden = 0x289da732
crc_channelFull = 0x17f45fcf
crc_messageActionChannelCreate = 0x95d2ac92
crc_messagesChannelMessages = 0x99262e37
crc_updateChannelTooLong = 0xeb0467fb
crc_updateChannel = 0xb6d45656
crc_updateNewChannelMessage = 0x62ba04d9
crc_updateReadChannelInbox = 0x4214f37f
crc_updateDeleteChannelMessages = 0xc37521c9
crc_updateChannelMessageViews = 0x98a12b4b
crc_inputChannelEmpty = 0xee8c1e86
crc_inputChannel = 0xafeb712e
crc_contactsResolvedPeer = 0x7f077ad9
crc_messageRange = 0x0ae30253
crc_updatesChannelDifferenceEmpty = 0x3e11affb
crc_updatesChannelDifferenceTooLong = 0x6a9d7b35
crc_updatesChannelDifference = 0x2064674e
crc_channelMessagesFilterEmpty = 0x94d42ee7
crc_channelMessagesFilter = 0xcd77d957
crc_channelParticipant = 0x15ebac1d
crc_channelParticipantSelf = 0xa3289a6d
crc_channelParticipantCreator = 0xe3e2e1f9
crc_channelParticipantsRecent = 0xde3f3c79
crc_channelParticipantsAdmins = 0xb4608969
crc_channelParticipantsKicked = 0xa3b54985
crc_channelsChannelParticipants = 0xf56ee2a8
crc_channelsChannelParticipant = 0xd0d9b163
crc_true = 0x3fedd339
crc_chatParticipantCreator = 0xda13538a
crc_chatParticipantAdmin = 0xe2d6e436
crc_updateChatAdmins = 0x6e947941
crc_updateChatParticipantAdmin = 0xb6901959
crc_messageActionChatMigrateTo = 0x51bdb021
crc_messageActionChannelMigrateFrom = 0xb055eaee
crc_channelParticipantsBots = 0xb0d1865b
crc_inputReportReasonSpam = 0x58dbcab8
crc_inputReportReasonViolence = 0x1e22c78d
crc_inputReportReasonPornography = 0x2e59d922
crc_inputReportReasonOther = 0xe1746d0a
crc_updateNewStickerSet = 0x688a30aa
crc_updateStickerSetsOrder = 0x0bb2d201
crc_updateStickerSets = 0x43ae3dec
crc_helpTermsOfService = 0xf1ee3e90
crc_foundGif = 0x162ecc1f
crc_inputMediaGifExternal = 0x4843b0fd
crc_messagesFoundGifs = 0x450a1c0a
crc_inputMessagesFilterGif = 0xffc86587
crc_updateSavedGifs = 0x9375341e
crc_updateBotInlineQuery = 0x54826690
crc_foundGifCached = 0x9c750409
crc_messagesSavedGifsNotModified = 0xe8025ca2
crc_messagesSavedGifs = 0x2e0709a5
crc_inputBotInlineMessageMediaAuto = 0x292fed13
crc_inputBotInlineMessageText = 0x3dcd7a87
crc_inputBotInlineResult = 0x2cbbe15a
crc_botInlineMessageMediaAuto = 0x0a74b15b
crc_botInlineMessageText = 0x8c7f65e2
crc_botInlineResult = 0x9bebaeb9
crc_messagesBotResults = 0xccd3563d
crc_inputMessagesFilterVoice = 0x50f5c392
crc_inputMessagesFilterMusic = 0x3751b49e
crc_updateBotInlineSend = 0x0e48f964
crc_inputPrivacyKeyChatInvite = 0xbdfb0426
crc_privacyKeyChatInvite = 0x500e6dfa
crc_updateEditChannelMessage = 0x1b3f4df7
crc_exportedMessageLink = 0x1f486803
crc_messageFwdHeader = 0xfadff4ac
crc_messageActionPinMessage = 0x94bd38ed
crc_peerSettings = 0x818426cd
crc_updateChannelPinnedMessage = 0x98592475
crc_keyboardButtonUrl = 0x258aff05
crc_keyboardButtonCallback = 0x683a5e46
crc_keyboardButtonRequestPhone = 0xb16a6c29
crc_keyboardButtonRequestGeoLocation = 0xfc796b3f
crc_authCodeTypeSms = 0x72a3158c
crc_authCodeTypeCall = 0x741cd3e3
crc_authCodeTypeFlashCall = 0x226ccefb
crc_authSentCodeTypeApp = 0x3dbb5986
crc_authSentCodeTypeSms = 0xc000bba2
crc_authSentCodeTypeCall = 0x5353e5a7
crc_authSentCodeTypeFlashCall = 0xab03c6d9
crc_keyboardButtonSwitchInline = 0x0568a748
crc_replyInlineMarkup = 0x48a30254
crc_messagesBotCallbackAnswer = 0x36585ea4
crc_updateBotCallbackQuery = 0xe73547e1
crc_messagesMessageEditData = 0x26b5dde6
crc_updateEditMessage = 0xe40370a3
crc_inputBotInlineMessageMediaGeo = 0xf4a59de1
crc_inputBotInlineMessageMediaVenue = 0xaaafadc8
crc_inputBotInlineMessageMediaContact = 0x2daf01a7
crc_botInlineMessageMediaGeo = 0x3a8fd8b8
crc_botInlineMessageMediaVenue = 0x4366232e
crc_botInlineMessageMediaContact = 0x35edb4d4
crc_inputBotInlineResultPhoto = 0xa8d864a7
crc_inputBotInlineResultDocument = 0xfff8fdc4
crc_botInlineMediaResult = 0x17db940b
crc_inputBotInlineMessageID = 0x890c3d89
crc_updateInlineBotCallbackQuery = 0xf9d27a5a
crc_inlineBotSwitchPM = 0x3c20629f
crc_messageEntityMentionName = 0x352dca58
crc_inputMessageEntityMentionName = 0x208e68c9
crc_messagesPeerDialogs = 0x3371c354
crc_topPeer = 0xedcdc05b
crc_topPeerCategoryBotsPM = 0xab661b5b
crc_topPeerCategoryBotsInline = 0x148677e2
crc_topPeerCategoryCorrespondents = 0x0637b7ed
crc_topPeerCategoryGroups = 0xbd17a14a
crc_topPeerCategoryChannels = 0x161d9628
crc_topPeerCategoryPeers = 0xfb834291
crc_contactsTopPeersNotModified = 0xde266ef5
crc_contactsTopPeers = 0x70b772a8
crc_inputMessagesFilterChatPhotos = 0x3a20ecb8
crc_updateReadChannelOutbox = 0x25d6c9c7
crc_updateDraftMessage = 0xee2bb969
crc_draftMessageEmpty = 0xba4baec5
crc_draftMessage = 0xfd8e711f
crc_messageActionHistoryClear = 0x9fbab604
crc_updateReadFeaturedStickers = 0x571d2742
crc_updateRecentStickers = 0x9a422c20
crc_messagesFeaturedStickersNotModified = 0x04ede3cf
crc_messagesFeaturedStickers = 0xf89d88e5
crc_messagesRecentStickersNotModified = 0x0b17f890
crc_messagesRecentStickers = 0x5ce20970
crc_messagesArchivedStickers = 0x4fcba9c8
crc_messagesStickerSetInstallResultSuccess = 0x38641628
crc_messagesStickerSetInstallResultArchive = 0x35e410a8
crc_stickerSetCovered = 0x6410a5d2
crc_inputMediaPhotoExternal = 0x0922aec1
crc_inputMediaDocumentExternal = 0xb6f74335
crc_updateConfig = 0xa229dd06
crc_updatePtsChanged = 0x3354678f
crc_messageActionGameScore = 0x92a72876
crc_documentAttributeHasStickers = 0x9801d2f7
crc_keyboardButtonGame = 0x50f41ccf
crc_stickerSetMultiCovered = 0x3407e51b
crc_maskCoords = 0xaed6dbb2
crc_inputStickeredMediaPhoto = 0x4a992157
crc_inputStickeredMediaDocument = 0x0438865b
crc_inputMediaGame = 0xd33f43f3
crc_messageMediaGame = 0xfdb19008
crc_inputBotInlineMessageGame = 0x4b425864
crc_inputBotInlineResultGame = 0x4fa417f2
crc_game = 0xbdf9653b
crc_inputGameID = 0x032c3e77
crc_inputGameShortName = 0xc331e80a
crc_highScore = 0x58fffcd0
crc_messagesHighScores = 0x9a3bfd99
crc_messagesChatsSlice = 0x9cd81144
crc_updateChannelWebPage = 0x40771900
crc_updatesDifferenceTooLong = 0x4afe8f6d
crc_sendMessageGamePlayAction = 0xdd6a8f48
crc_webPageNotModified = 0x85849473
crc_textEmpty = 0xdc3d824f
crc_textPlain = 0x744694e0
crc_textBold = 0x6724abc4
crc_textItalic = 0xd912a59c
crc_textUnderline = 0xc12622c4
crc_textStrike = 0x9bf8bb95
crc_textFixed = 0x6c3f19b9
crc_textUrl = 0x3c2884c1
crc_textEmail = 0xde5a0dd6
crc_textConcat = 0x7e6260d7
crc_pageBlockTitle = 0x70abc3fd
crc_pageBlockSubtitle = 0x8ffa9a1f
crc_pageBlockAuthorDate = 0xbaafe5e0
crc_pageBlockHeader = 0xbfd064ec
crc_pageBlockSubheader = 0xf12bb6e1
crc_pageBlockParagraph = 0x467a0766
crc_pageBlockPreformatted = 0xc070d93e
crc_pageBlockFooter = 0x48870999
crc_pageBlockDivider = 0xdb20b188
crc_pageBlockList = 0x3a58c7f4
crc_pageBlockBlockquote = 0x263d7c26
crc_pageBlockPullquote = 0x4f4456d3
crc_pageBlockPhoto = 0xe9c69982
crc_pageBlockVideo = 0xd9d71866
crc_pageBlockCover = 0x39f23300
crc_pageBlockEmbed = 0xcde200d1
crc_pageBlockEmbedPost = 0x292c7be9
crc_pageBlockSlideshow = 0x130c8963
crc_pagePart = 0x8e3f9ebe
crc_pageFull = 0x556ec7aa
crc_updatePhoneCall = 0xab0f6b1e
crc_updateDialogPinned = 0xd711a2cc
crc_updatePinnedDialogs = 0xd8caf68d
crc_inputPrivacyKeyPhoneCall = 0xfabadc5f
crc_privacyKeyPhoneCall = 0x3d662b7b
crc_pageBlockUnsupported = 0x13567e8a
crc_pageBlockAnchor = 0xce0d37b0
crc_pageBlockCollage = 0x08b31c4f
crc_inputPhoneCall = 0x1e36fded
crc_phoneCallEmpty = 0x5366c915
crc_phoneCallWaiting = 0x1b8f4ad1
crc_phoneCallRequested = 0x83761ce4
crc_phoneCall = 0xffe6ab67
crc_phoneCallDiscarded = 0x50ca4de1
crc_phoneConnection = 0x9d4c17c0
crc_phoneCallProtocol = 0xa2bb35cb
crc_phonePhoneCall = 0xec82e140
crc_phoneCallDiscardReasonMissed = 0x85e42301
crc_phoneCallDiscardReasonDisconnect = 0xe095c1a0
crc_phoneCallDiscardReasonHangup = 0x57adc690
crc_phoneCallDiscardReasonBusy = 0xfaf7e8c9
crc_inputMessagesFilterPhoneCalls = 0x80c99768
crc_messageActionPhoneCall = 0x80e11a7f
crc_invoice = 0xc30aa358
crc_inputMediaInvoice = 0x92153685
crc_messageActionPaymentSentMe = 0x8f31b327
crc_messageMediaInvoice = 0x84551347
crc_keyboardButtonBuy = 0xafd93fbb
crc_messageActionPaymentSent = 0x40699cd0
crc_paymentsPaymentForm = 0x3f56aea3
crc_postAddress = 0x1e8caaeb
crc_paymentRequestedInfo = 0x909c3f94
crc_updateBotWebhookJSON = 0x8317c0c3
crc_updateBotWebhookJSONQuery = 0x9b9240a6
crc_updateBotShippingQuery = 0xe0cdc940
crc_updateBotPrecheckoutQuery = 0x5d2f3aa9
crc_dataJSON = 0x7d748d04
crc_labeledPrice = 0xcb296bf8
crc_paymentCharge = 0xea02c27e
crc_paymentSavedCredentialsCard = 0xcdc27a1f
crc_webDocument = 0xc61acbd8
crc_inputWebDocument = 0x9bed434d
crc_inputWebFileLocation = 0xc239d686
crc_uploadWebFile = 0x21e753bc
crc_paymentsValidatedRequestedInfo = 0xd1451883
crc_paymentsPaymentResult = 0x4e5f810d
crc_paymentsPaymentVerficationNeeded = 0x6b56b921
crc_paymentsPaymentReceipt = 0x500911e1
crc_paymentsSavedInfo = 0xfb8fe43c
crc_inputPaymentCredentialsSaved = 0xc10eb2cf
crc_inputPaymentCredentials = 0x3417d728
crc_accountTmpPassword = 0xdb64fd34
crc_shippingOption = 0xb6213cdf
crc_phoneCallAccepted = 0x6d003d3f
crc_inputMessagesFilterRoundVoice = 0x7a7c17a4
crc_inputMessagesFilterRoundVideo = 0xb549da53
crc_uploadFileCdnRedirect = 0xea52fe5a
crc_sendMessageRecordRoundAction = 0x88f27fbc
crc_sendMessageUploadRoundAction = 0x243e1c66
crc_uploadCdnFileReuploadNeeded = 0xeea8e46e
crc_uploadCdnFile = 0xa99fca4f
crc_cdnPublicKey = 0xc982eaba
crc_cdnConfig = 0x5725e40a
crc_updateLangPackTooLong = 0x10c2404b
crc_updateLangPack = 0x56022f4d
crc_pageBlockChannel = 0xef1751b5
crc_inputStickerSetItem = 0xffa0a496
crc_langPackString = 0xcad181f6
crc_langPackStringPluralized = 0x6c47ac9f
crc_langPackStringDeleted = 0x2979eeb2
crc_langPackDifference = 0xf385c1f6
crc_langPackLanguage = 0x117698f1
crc_channelParticipantAdmin = 0xa82fa898
crc_channelParticipantBanned = 0x222c1886
crc_channelParticipantsBanned = 0x1427a5e1
crc_channelParticipantsSearch = 0x0656ac4b
crc_topPeerCategoryPhoneCalls = 0x1e76a78c
crc_pageBlockAudio = 0x31b81a7f
crc_channelAdminRights = 0x5d7ceba5
crc_channelBannedRights = 0x58cf4249
crc_channelAdminLogEventActionChangeTitle = 0xe6dfb825
crc_channelAdminLogEventActionChangeAbout = 0x55188a2e
crc_channelAdminLogEventActionChangeUsername = 0x6a4afc38
crc_channelAdminLogEventActionChangePhoto = 0xb82f55c3
crc_channelAdminLogEventActionToggleInvites = 0x1b7907ae
crc_channelAdminLogEventActionToggleSignatures = 0x26ae0971
crc_channelAdminLogEventActionUpdatePinned = 0xe9e82c18
crc_channelAdminLogEventActionEditMessage = 0x709b2405
crc_channelAdminLogEventActionDeleteMessage = 0x42e047bb
crc_channelAdminLogEventActionParticipantJoin = 0x183040d3
crc_channelAdminLogEventActionParticipantLeave = 0xf89777f2
crc_channelAdminLogEventActionParticipantInvite = 0xe31c34d8
crc_channelAdminLogEventActionParticipantToggleBan = 0xe6d83d7e
crc_channelAdminLogEventActionParticipantToggleAdmin = 0xd5676710
crc_channelAdminLogEvent = 0x3b5a3e40
crc_channelsAdminLogResults = 0xed8af74d
crc_channelAdminLogEventsFilter = 0xea107ae4
crc_messageActionScreenshotTaken = 0x4792929b
crc_popularContact = 0x5ce14175
crc_cdnFileHash = 0x77eec38f
crc_inputMessagesFilterMyMentions = 0xc1f8e69a
crc_inputMessagesFilterMyMentionsUnread = 0x46caf4a8
crc_updateContactsReset = 0x7084a7be
crc_channelAdminLogEventActionChangeStickerSet = 0xb1c3caa7
crc_updateFavedStickers = 0xe511996d
crc_messagesFavedStickers = 0xf37f2f16
crc_messagesFavedStickersNotModified = 0x9e8fa6d3
crc_updateChannelReadMessagesContents = 0x89893b45
)



// Encode funcs for types
func (e TypeBool) encode() []byte {
	if e.GetValue() != nil {
		return e.GetBoolFalse().encode()
	} else if e.GetValue() != nil {
		return e.GetBoolFalse().encode()
	} else if e.GetValue() != nil {
		return e.GetBoolTrue().encode()
	}
	return nil

}
func (e TypeError) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeNull) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputPeer) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputPeerEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPeerEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPeerSelf().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPeerChat().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPeerUser().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPeerChannel().encode()
	}
	return nil

}
func (e TypeInputUser) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputUserEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputUserEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputUserSelf().encode()
	} else if e.GetValue() != nil {
		return e.GetInputUser().encode()
	}
	return nil

}
func (e TypeInputContact) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputFile) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputFile().encode()
	} else if e.GetValue() != nil {
		return e.GetInputFile().encode()
	} else if e.GetValue() != nil {
		return e.GetInputFileBig().encode()
	}
	return nil

}
func (e TypeInputMedia) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputMediaEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaUploadedPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaGeoPoint().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaContact().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaUploadedDocument().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaDocument().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaVenue().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaGifExternal().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaPhotoExternal().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaDocumentExternal().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaGame().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMediaInvoice().encode()
	}
	return nil

}
func (e TypeInputChatPhoto) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputChatPhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputChatPhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputChatUploadedPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetInputChatPhoto().encode()
	}
	return nil

}
func (e TypeInputGeoPoint) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputGeoPointEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputGeoPointEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputGeoPoint().encode()
	}
	return nil

}
func (e TypeInputPhoto) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputPhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPhoto().encode()
	}
	return nil

}
func (e TypeInputFileLocation) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputFileLocation().encode()
	} else if e.GetValue() != nil {
		return e.GetInputFileLocation().encode()
	} else if e.GetValue() != nil {
		return e.GetInputEncryptedFileLocation().encode()
	} else if e.GetValue() != nil {
		return e.GetInputDocumentFileLocation().encode()
	}
	return nil

}
func (e TypeInputAppEvent) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePeer) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPeerUser().encode()
	} else if e.GetValue() != nil {
		return e.GetPeerUser().encode()
	} else if e.GetValue() != nil {
		return e.GetPeerChat().encode()
	} else if e.GetValue() != nil {
		return e.GetPeerChannel().encode()
	}
	return nil

}
func (e TypeStorageFileType) encode() []byte {
	if e.GetValue() != nil {
		return e.GetStorageFileUnknown().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFileUnknown().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFileJpeg().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFileGif().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFilePng().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFileMp3().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFileMov().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFilePartial().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFileMp4().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFileWebp().encode()
	} else if e.GetValue() != nil {
		return e.GetStorageFilePdf().encode()
	}
	return nil

}
func (e TypeFileLocation) encode() []byte {
	if e.GetValue() != nil {
		return e.GetFileLocationUnavailable().encode()
	} else if e.GetValue() != nil {
		return e.GetFileLocationUnavailable().encode()
	} else if e.GetValue() != nil {
		return e.GetFileLocation().encode()
	}
	return nil

}
func (e TypeUser) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUserEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUserEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUser().encode()
	}
	return nil

}
func (e TypeUserProfilePhoto) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUserProfilePhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUserProfilePhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUserProfilePhoto().encode()
	}
	return nil

}
func (e TypeUserStatus) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUserStatusEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUserStatusEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUserStatusOnline().encode()
	} else if e.GetValue() != nil {
		return e.GetUserStatusOffline().encode()
	} else if e.GetValue() != nil {
		return e.GetUserStatusRecently().encode()
	} else if e.GetValue() != nil {
		return e.GetUserStatusLastWeek().encode()
	} else if e.GetValue() != nil {
		return e.GetUserStatusLastMonth().encode()
	}
	return nil

}
func (e TypeChat) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChatEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetChatEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetChat().encode()
	} else if e.GetValue() != nil {
		return e.GetChatForbidden().encode()
	} else if e.GetValue() != nil {
		return e.GetChannel().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelForbidden().encode()
	}
	return nil

}
func (e TypeChatFull) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChatFull().encode()
	} else if e.GetValue() != nil {
		return e.GetChatFull().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelFull().encode()
	}
	return nil

}
func (e TypeChatParticipant) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChatParticipant().encode()
	} else if e.GetValue() != nil {
		return e.GetChatParticipant().encode()
	} else if e.GetValue() != nil {
		return e.GetChatParticipantCreator().encode()
	} else if e.GetValue() != nil {
		return e.GetChatParticipantAdmin().encode()
	}
	return nil

}
func (e TypeChatParticipants) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChatParticipantsForbidden().encode()
	} else if e.GetValue() != nil {
		return e.GetChatParticipantsForbidden().encode()
	} else if e.GetValue() != nil {
		return e.GetChatParticipants().encode()
	}
	return nil

}
func (e TypeChatPhoto) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChatPhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetChatPhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetChatPhoto().encode()
	}
	return nil

}
func (e TypeMessage) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessageEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageService().encode()
	}
	return nil

}
func (e TypeMessageMedia) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessageMediaEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaGeo().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaContact().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaUnsupported().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaDocument().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaWebPage().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaVenue().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaGame().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageMediaInvoice().encode()
	}
	return nil

}
func (e TypeMessageAction) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessageActionEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChatCreate().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChatEditTitle().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChatEditPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChatDeletePhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChatAddUser().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChatDeleteUser().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChatJoinedByLink().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChannelCreate().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChatMigrateTo().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionChannelMigrateFrom().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionPinMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionHistoryClear().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionGameScore().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionPhoneCall().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionPaymentSentMe().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionPaymentSent().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageActionScreenshotTaken().encode()
	}
	return nil

}
func (e TypeDialog) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePhoto) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPhotoEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoto().encode()
	}
	return nil

}
func (e TypePhotoSize) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPhotoSizeEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPhotoSizeEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPhotoSize().encode()
	} else if e.GetValue() != nil {
		return e.GetPhotoCachedSize().encode()
	}
	return nil

}
func (e TypeGeoPoint) encode() []byte {
	if e.GetValue() != nil {
		return e.GetGeoPointEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetGeoPointEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetGeoPoint().encode()
	}
	return nil

}
func (e TypeAuthCheckedPhone) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAuthSentCode) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAuthAuthorization) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAuthExportedAuthorization) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputNotifyPeer) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputNotifyPeer().encode()
	} else if e.GetValue() != nil {
		return e.GetInputNotifyPeer().encode()
	} else if e.GetValue() != nil {
		return e.GetInputNotifyUsers().encode()
	} else if e.GetValue() != nil {
		return e.GetInputNotifyChats().encode()
	} else if e.GetValue() != nil {
		return e.GetInputNotifyAll().encode()
	}
	return nil

}
func (e TypeInputPeerNotifySettings) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePeerNotifyEvents) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPeerNotifyEventsEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPeerNotifyEventsEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPeerNotifyEventsAll().encode()
	}
	return nil

}
func (e TypePeerNotifySettings) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPeerNotifySettingsEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPeerNotifySettingsEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPeerNotifySettings().encode()
	}
	return nil

}
func (e TypeWallPaper) encode() []byte {
	if e.GetValue() != nil {
		return e.GetWallPaper().encode()
	} else if e.GetValue() != nil {
		return e.GetWallPaper().encode()
	} else if e.GetValue() != nil {
		return e.GetWallPaperSolid().encode()
	}
	return nil

}
func (e TypeUserFull) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeContact) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeImportedContact) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeContactBlocked) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeContactStatus) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeContactsLink) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeContactsContacts) encode() []byte {
	if e.GetValue() != nil {
		return e.GetContactsContacts().encode()
	} else if e.GetValue() != nil {
		return e.GetContactsContacts().encode()
	} else if e.GetValue() != nil {
		return e.GetContactsContactsNotModified().encode()
	}
	return nil

}
func (e TypeContactsImportedContacts) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeContactsBlocked) encode() []byte {
	if e.GetValue() != nil {
		return e.GetContactsBlocked().encode()
	} else if e.GetValue() != nil {
		return e.GetContactsBlocked().encode()
	} else if e.GetValue() != nil {
		return e.GetContactsBlockedSlice().encode()
	}
	return nil

}
func (e TypeContactsFound) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesDialogs) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesDialogs().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesDialogs().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesDialogsSlice().encode()
	}
	return nil

}
func (e TypeMessagesMessages) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesMessages().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesMessages().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesMessagesSlice().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesChannelMessages().encode()
	}
	return nil

}
func (e TypeMessagesChats) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesChats().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesChats().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesChatsSlice().encode()
	}
	return nil

}
func (e TypeMessagesChatFull) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesAffectedHistory) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesFilter) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputMessagesFilterEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterPhotos().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterVideo().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterPhotoVideo().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterDocument().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterPhotoVideoDocuments().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterUrl().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterGif().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterVoice().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterMusic().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterChatPhotos().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterPhoneCalls().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterRoundVoice().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterRoundVideo().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterMyMentions().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessagesFilterMyMentionsUnread().encode()
	}
	return nil

}
func (e TypeUpdate) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUpdateNewMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateNewMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateMessageID().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateDeleteMessages().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateUserTyping().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChatUserTyping().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChatParticipants().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateUserStatus().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateUserName().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateUserPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateContactRegistered().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateContactLink().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateNewEncryptedMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateEncryptedChatTyping().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateEncryption().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateEncryptedMessagesRead().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChatParticipantAdd().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChatParticipantDelete().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateDcOptions().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateUserBlocked().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateNotifySettings().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateServiceNotification().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatePrivacy().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateUserPhone().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateReadHistoryInbox().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateReadHistoryOutbox().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateWebPage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateReadMessagesContents().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChannelTooLong().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChannel().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateNewChannelMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateReadChannelInbox().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateDeleteChannelMessages().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChannelMessageViews().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChatAdmins().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChatParticipantAdmin().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateNewStickerSet().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateStickerSetsOrder().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateStickerSets().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateSavedGifs().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateBotInlineQuery().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateBotInlineSend().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateEditChannelMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChannelPinnedMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateBotCallbackQuery().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateEditMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateInlineBotCallbackQuery().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateReadChannelOutbox().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateDraftMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateReadFeaturedStickers().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateRecentStickers().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateConfig().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatePtsChanged().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChannelWebPage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatePhoneCall().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateDialogPinned().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatePinnedDialogs().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateBotWebhookJSON().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateBotWebhookJSONQuery().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateBotShippingQuery().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateBotPrecheckoutQuery().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateLangPackTooLong().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateLangPack().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateContactsReset().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateFavedStickers().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateChannelReadMessagesContents().encode()
	}
	return nil

}
func (e TypeUpdatesState) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeUpdatesDifference) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUpdatesDifferenceEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesDifferenceEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesDifference().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesDifferenceSlice().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesDifferenceTooLong().encode()
	}
	return nil

}
func (e TypeUpdates) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUpdatesTooLong().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesTooLong().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateShortMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateShortChatMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateShort().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesCombined().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdates().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdateShortSentMessage().encode()
	}
	return nil

}
func (e TypePhotosPhoto) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeUploadFile) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUploadFile().encode()
	} else if e.GetValue() != nil {
		return e.GetUploadFile().encode()
	} else if e.GetValue() != nil {
		return e.GetUploadFileCdnRedirect().encode()
	}
	return nil

}
func (e TypeDcOption) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeConfig) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeNearestDc) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeHelpAppUpdate) encode() []byte {
	if e.GetValue() != nil {
		return e.GetHelpAppUpdate().encode()
	} else if e.GetValue() != nil {
		return e.GetHelpAppUpdate().encode()
	} else if e.GetValue() != nil {
		return e.GetHelpNoAppUpdate().encode()
	}
	return nil

}
func (e TypeHelpInviteText) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputPeerNotifyEvents) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputPeerNotifyEventsEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPeerNotifyEventsEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPeerNotifyEventsAll().encode()
	}
	return nil

}
func (e TypePhotosPhotos) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPhotosPhotos().encode()
	} else if e.GetValue() != nil {
		return e.GetPhotosPhotos().encode()
	} else if e.GetValue() != nil {
		return e.GetPhotosPhotosSlice().encode()
	}
	return nil

}
func (e TypeEncryptedChat) encode() []byte {
	if e.GetValue() != nil {
		return e.GetEncryptedChatEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedChatEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedChatWaiting().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedChatRequested().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedChat().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedChatDiscarded().encode()
	}
	return nil

}
func (e TypeInputEncryptedChat) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeEncryptedFile) encode() []byte {
	if e.GetValue() != nil {
		return e.GetEncryptedFileEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedFileEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedFile().encode()
	}
	return nil

}
func (e TypeInputEncryptedFile) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputEncryptedFileEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputEncryptedFileEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputEncryptedFileUploaded().encode()
	} else if e.GetValue() != nil {
		return e.GetInputEncryptedFile().encode()
	} else if e.GetValue() != nil {
		return e.GetInputEncryptedFileBigUploaded().encode()
	}
	return nil

}
func (e TypeEncryptedMessage) encode() []byte {
	if e.GetValue() != nil {
		return e.GetEncryptedMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetEncryptedMessageService().encode()
	}
	return nil

}
func (e TypeMessagesDhConfig) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesDhConfigNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesDhConfigNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesDhConfig().encode()
	}
	return nil

}
func (e TypeMessagesSentEncryptedMessage) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesSentEncryptedMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesSentEncryptedMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesSentEncryptedFile().encode()
	}
	return nil

}
func (e TypeInputDocument) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputDocumentEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputDocumentEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputDocument().encode()
	}
	return nil

}
func (e TypeDocument) encode() []byte {
	if e.GetValue() != nil {
		return e.GetDocumentEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetDocumentEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetDocument().encode()
	}
	return nil

}
func (e TypeHelpSupport) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeNotifyPeer) encode() []byte {
	if e.GetValue() != nil {
		return e.GetNotifyAll().encode()
	} else if e.GetValue() != nil {
		return e.GetNotifyAll().encode()
	} else if e.GetValue() != nil {
		return e.GetNotifyChats().encode()
	} else if e.GetValue() != nil {
		return e.GetNotifyPeer().encode()
	} else if e.GetValue() != nil {
		return e.GetNotifyUsers().encode()
	}
	return nil

}
func (e TypeSendMessageAction) encode() []byte {
	if e.GetValue() != nil {
		return e.GetSendMessageTypingAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageTypingAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageCancelAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageRecordVideoAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageUploadVideoAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageRecordAudioAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageUploadAudioAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageUploadPhotoAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageUploadDocumentAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageGeoLocationAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageChooseContactAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageGamePlayAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageRecordRoundAction().encode()
	} else if e.GetValue() != nil {
		return e.GetSendMessageUploadRoundAction().encode()
	}
	return nil

}
func (e TypeInputPrivacyKey) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputPrivacyKeyStatusTimestamp().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyKeyStatusTimestamp().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyKeyChatInvite().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyKeyPhoneCall().encode()
	}
	return nil

}
func (e TypePrivacyKey) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPrivacyKeyStatusTimestamp().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyKeyStatusTimestamp().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyKeyChatInvite().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyKeyPhoneCall().encode()
	}
	return nil

}
func (e TypeInputPrivacyRule) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputPrivacyValueAllowContacts().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyValueAllowContacts().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyValueAllowAll().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyValueAllowUsers().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyValueDisallowContacts().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyValueDisallowAll().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPrivacyValueDisallowUsers().encode()
	}
	return nil

}
func (e TypePrivacyRule) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPrivacyValueAllowContacts().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyValueAllowContacts().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyValueAllowAll().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyValueAllowUsers().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyValueDisallowContacts().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyValueDisallowAll().encode()
	} else if e.GetValue() != nil {
		return e.GetPrivacyValueDisallowUsers().encode()
	}
	return nil

}
func (e TypeAccountPrivacyRules) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAccountDaysTTL) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeDisabledFeature) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeDocumentAttribute) encode() []byte {
	if e.GetValue() != nil {
		return e.GetDocumentAttributeImageSize().encode()
	} else if e.GetValue() != nil {
		return e.GetDocumentAttributeImageSize().encode()
	} else if e.GetValue() != nil {
		return e.GetDocumentAttributeAnimated().encode()
	} else if e.GetValue() != nil {
		return e.GetDocumentAttributeSticker().encode()
	} else if e.GetValue() != nil {
		return e.GetDocumentAttributeVideo().encode()
	} else if e.GetValue() != nil {
		return e.GetDocumentAttributeAudio().encode()
	} else if e.GetValue() != nil {
		return e.GetDocumentAttributeFilename().encode()
	} else if e.GetValue() != nil {
		return e.GetDocumentAttributeHasStickers().encode()
	}
	return nil

}
func (e TypeMessagesStickers) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesStickersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesStickersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesStickers().encode()
	}
	return nil

}
func (e TypeStickerPack) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesAllStickers) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesAllStickersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesAllStickersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesAllStickers().encode()
	}
	return nil

}
func (e TypeAccountPassword) encode() []byte {
	if e.GetValue() != nil {
		return e.GetAccountNoPassword().encode()
	} else if e.GetValue() != nil {
		return e.GetAccountNoPassword().encode()
	} else if e.GetValue() != nil {
		return e.GetAccountPassword().encode()
	}
	return nil

}
func (e TypeMessagesAffectedMessages) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeContactLink) encode() []byte {
	if e.GetValue() != nil {
		return e.GetContactLinkUnknown().encode()
	} else if e.GetValue() != nil {
		return e.GetContactLinkUnknown().encode()
	} else if e.GetValue() != nil {
		return e.GetContactLinkNone().encode()
	} else if e.GetValue() != nil {
		return e.GetContactLinkHasPhone().encode()
	} else if e.GetValue() != nil {
		return e.GetContactLinkContact().encode()
	}
	return nil

}
func (e TypeWebPage) encode() []byte {
	if e.GetValue() != nil {
		return e.GetWebPageEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetWebPageEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetWebPagePending().encode()
	} else if e.GetValue() != nil {
		return e.GetWebPage().encode()
	} else if e.GetValue() != nil {
		return e.GetWebPageNotModified().encode()
	}
	return nil

}
func (e TypeAuthorization) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAccountAuthorizations) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAccountPasswordSettings) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAccountPasswordInputSettings) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAuthPasswordRecovery) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeReceivedNotifyMessage) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeExportedChatInvite) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChatInviteEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetChatInviteEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetChatInviteExported().encode()
	}
	return nil

}
func (e TypeChatInvite) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChatInviteAlready().encode()
	} else if e.GetValue() != nil {
		return e.GetChatInviteAlready().encode()
	} else if e.GetValue() != nil {
		return e.GetChatInvite().encode()
	}
	return nil

}
func (e TypeInputStickerSet) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputStickerSetEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputStickerSetEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputStickerSetID().encode()
	} else if e.GetValue() != nil {
		return e.GetInputStickerSetShortName().encode()
	}
	return nil

}
func (e TypeStickerSet) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesStickerSet) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeBotCommand) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeBotInfo) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeKeyboardButton) encode() []byte {
	if e.GetValue() != nil {
		return e.GetKeyboardButton().encode()
	} else if e.GetValue() != nil {
		return e.GetKeyboardButton().encode()
	} else if e.GetValue() != nil {
		return e.GetKeyboardButtonUrl().encode()
	} else if e.GetValue() != nil {
		return e.GetKeyboardButtonCallback().encode()
	} else if e.GetValue() != nil {
		return e.GetKeyboardButtonRequestPhone().encode()
	} else if e.GetValue() != nil {
		return e.GetKeyboardButtonRequestGeoLocation().encode()
	} else if e.GetValue() != nil {
		return e.GetKeyboardButtonSwitchInline().encode()
	} else if e.GetValue() != nil {
		return e.GetKeyboardButtonGame().encode()
	} else if e.GetValue() != nil {
		return e.GetKeyboardButtonBuy().encode()
	}
	return nil

}
func (e TypeKeyboardButtonRow) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeReplyMarkup) encode() []byte {
	if e.GetValue() != nil {
		return e.GetReplyKeyboardHide().encode()
	} else if e.GetValue() != nil {
		return e.GetReplyKeyboardHide().encode()
	} else if e.GetValue() != nil {
		return e.GetReplyKeyboardForceReply().encode()
	} else if e.GetValue() != nil {
		return e.GetReplyKeyboardMarkup().encode()
	} else if e.GetValue() != nil {
		return e.GetReplyInlineMarkup().encode()
	}
	return nil

}
func (e TypeMessageEntity) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessageEntityUnknown().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityUnknown().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityMention().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityHashtag().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityBotCommand().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityUrl().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityEmail().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityBold().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityItalic().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityCode().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityPre().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityTextUrl().encode()
	} else if e.GetValue() != nil {
		return e.GetMessageEntityMentionName().encode()
	} else if e.GetValue() != nil {
		return e.GetInputMessageEntityMentionName().encode()
	}
	return nil

}
func (e TypeInputChannel) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputChannelEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputChannelEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetInputChannel().encode()
	}
	return nil

}
func (e TypeContactsResolvedPeer) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessageRange) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeUpdatesChannelDifference) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUpdatesChannelDifferenceEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesChannelDifferenceEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesChannelDifferenceTooLong().encode()
	} else if e.GetValue() != nil {
		return e.GetUpdatesChannelDifference().encode()
	}
	return nil

}
func (e TypeChannelMessagesFilter) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChannelMessagesFilterEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelMessagesFilterEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelMessagesFilter().encode()
	}
	return nil

}
func (e TypeChannelParticipant) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChannelParticipant().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipant().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantSelf().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantCreator().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantAdmin().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantBanned().encode()
	}
	return nil

}
func (e TypeChannelParticipantsFilter) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChannelParticipantsRecent().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantsRecent().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantsAdmins().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantsKicked().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantsBots().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantsBanned().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelParticipantsSearch().encode()
	}
	return nil

}
func (e TypeChannelsChannelParticipants) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeChannelsChannelParticipant) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeTrue) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeReportReason) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputReportReasonSpam().encode()
	} else if e.GetValue() != nil {
		return e.GetInputReportReasonSpam().encode()
	} else if e.GetValue() != nil {
		return e.GetInputReportReasonViolence().encode()
	} else if e.GetValue() != nil {
		return e.GetInputReportReasonPornography().encode()
	} else if e.GetValue() != nil {
		return e.GetInputReportReasonOther().encode()
	}
	return nil

}
func (e TypeHelpTermsOfService) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeFoundGif) encode() []byte {
	if e.GetValue() != nil {
		return e.GetFoundGif().encode()
	} else if e.GetValue() != nil {
		return e.GetFoundGif().encode()
	} else if e.GetValue() != nil {
		return e.GetFoundGifCached().encode()
	}
	return nil

}
func (e TypeMessagesFoundGifs) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesSavedGifs) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesSavedGifsNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesSavedGifsNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesSavedGifs().encode()
	}
	return nil

}
func (e TypeInputBotInlineMessage) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputBotInlineMessageMediaAuto().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineMessageMediaAuto().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineMessageText().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineMessageMediaGeo().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineMessageMediaVenue().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineMessageMediaContact().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineMessageGame().encode()
	}
	return nil

}
func (e TypeInputBotInlineResult) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputBotInlineResult().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineResult().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineResultPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineResultDocument().encode()
	} else if e.GetValue() != nil {
		return e.GetInputBotInlineResultGame().encode()
	}
	return nil

}
func (e TypeBotInlineMessage) encode() []byte {
	if e.GetValue() != nil {
		return e.GetBotInlineMessageMediaAuto().encode()
	} else if e.GetValue() != nil {
		return e.GetBotInlineMessageMediaAuto().encode()
	} else if e.GetValue() != nil {
		return e.GetBotInlineMessageText().encode()
	} else if e.GetValue() != nil {
		return e.GetBotInlineMessageMediaGeo().encode()
	} else if e.GetValue() != nil {
		return e.GetBotInlineMessageMediaVenue().encode()
	} else if e.GetValue() != nil {
		return e.GetBotInlineMessageMediaContact().encode()
	}
	return nil

}
func (e TypeBotInlineResult) encode() []byte {
	if e.GetValue() != nil {
		return e.GetBotInlineResult().encode()
	} else if e.GetValue() != nil {
		return e.GetBotInlineResult().encode()
	} else if e.GetValue() != nil {
		return e.GetBotInlineMediaResult().encode()
	}
	return nil

}
func (e TypeMessagesBotResults) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeExportedMessageLink) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessageFwdHeader) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePeerSettings) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeAuthCodeType) encode() []byte {
	if e.GetValue() != nil {
		return e.GetAuthCodeTypeSms().encode()
	} else if e.GetValue() != nil {
		return e.GetAuthCodeTypeSms().encode()
	} else if e.GetValue() != nil {
		return e.GetAuthCodeTypeCall().encode()
	} else if e.GetValue() != nil {
		return e.GetAuthCodeTypeFlashCall().encode()
	}
	return nil

}
func (e TypeAuthSentCodeType) encode() []byte {
	if e.GetValue() != nil {
		return e.GetAuthSentCodeTypeApp().encode()
	} else if e.GetValue() != nil {
		return e.GetAuthSentCodeTypeApp().encode()
	} else if e.GetValue() != nil {
		return e.GetAuthSentCodeTypeSms().encode()
	} else if e.GetValue() != nil {
		return e.GetAuthSentCodeTypeCall().encode()
	} else if e.GetValue() != nil {
		return e.GetAuthSentCodeTypeFlashCall().encode()
	}
	return nil

}
func (e TypeMessagesBotCallbackAnswer) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesMessageEditData) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputBotInlineMessageID) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInlineBotSwitchPM) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesPeerDialogs) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeTopPeer) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeTopPeerCategory) encode() []byte {
	if e.GetValue() != nil {
		return e.GetTopPeerCategoryBotsPM().encode()
	} else if e.GetValue() != nil {
		return e.GetTopPeerCategoryBotsPM().encode()
	} else if e.GetValue() != nil {
		return e.GetTopPeerCategoryBotsInline().encode()
	} else if e.GetValue() != nil {
		return e.GetTopPeerCategoryCorrespondents().encode()
	} else if e.GetValue() != nil {
		return e.GetTopPeerCategoryGroups().encode()
	} else if e.GetValue() != nil {
		return e.GetTopPeerCategoryChannels().encode()
	} else if e.GetValue() != nil {
		return e.GetTopPeerCategoryPhoneCalls().encode()
	}
	return nil

}
func (e TypeTopPeerCategoryPeers) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeContactsTopPeers) encode() []byte {
	if e.GetValue() != nil {
		return e.GetContactsTopPeersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetContactsTopPeersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetContactsTopPeers().encode()
	}
	return nil

}
func (e TypeDraftMessage) encode() []byte {
	if e.GetValue() != nil {
		return e.GetDraftMessageEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetDraftMessageEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetDraftMessage().encode()
	}
	return nil

}
func (e TypeMessagesFeaturedStickers) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesFeaturedStickersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesFeaturedStickersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesFeaturedStickers().encode()
	}
	return nil

}
func (e TypeMessagesRecentStickers) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesRecentStickersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesRecentStickersNotModified().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesRecentStickers().encode()
	}
	return nil

}
func (e TypeMessagesArchivedStickers) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesStickerSetInstallResult) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesStickerSetInstallResultSuccess().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesStickerSetInstallResultSuccess().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesStickerSetInstallResultArchive().encode()
	}
	return nil

}
func (e TypeStickerSetCovered) encode() []byte {
	if e.GetValue() != nil {
		return e.GetStickerSetCovered().encode()
	} else if e.GetValue() != nil {
		return e.GetStickerSetCovered().encode()
	} else if e.GetValue() != nil {
		return e.GetStickerSetMultiCovered().encode()
	}
	return nil

}
func (e TypeMaskCoords) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputStickeredMedia) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputStickeredMediaPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetInputStickeredMediaPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetInputStickeredMediaDocument().encode()
	}
	return nil

}
func (e TypeGame) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputGame) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputGameID().encode()
	} else if e.GetValue() != nil {
		return e.GetInputGameID().encode()
	} else if e.GetValue() != nil {
		return e.GetInputGameShortName().encode()
	}
	return nil

}
func (e TypeHighScore) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesHighScores) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeRichText) encode() []byte {
	if e.GetValue() != nil {
		return e.GetTextEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetTextEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetTextPlain().encode()
	} else if e.GetValue() != nil {
		return e.GetTextBold().encode()
	} else if e.GetValue() != nil {
		return e.GetTextItalic().encode()
	} else if e.GetValue() != nil {
		return e.GetTextUnderline().encode()
	} else if e.GetValue() != nil {
		return e.GetTextStrike().encode()
	} else if e.GetValue() != nil {
		return e.GetTextFixed().encode()
	} else if e.GetValue() != nil {
		return e.GetTextUrl().encode()
	} else if e.GetValue() != nil {
		return e.GetTextEmail().encode()
	} else if e.GetValue() != nil {
		return e.GetTextConcat().encode()
	}
	return nil

}
func (e TypePageBlock) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPageBlockTitle().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockTitle().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockSubtitle().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockAuthorDate().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockHeader().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockSubheader().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockParagraph().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockPreformatted().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockFooter().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockDivider().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockList().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockBlockquote().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockPullquote().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockPhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockVideo().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockCover().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockEmbed().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockEmbedPost().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockSlideshow().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockUnsupported().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockAnchor().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockCollage().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockChannel().encode()
	} else if e.GetValue() != nil {
		return e.GetPageBlockAudio().encode()
	}
	return nil

}
func (e TypePage) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPagePart().encode()
	} else if e.GetValue() != nil {
		return e.GetPagePart().encode()
	} else if e.GetValue() != nil {
		return e.GetPageFull().encode()
	}
	return nil

}
func (e TypeInputPhoneCall) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePhoneCall) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPhoneCallEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallEmpty().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallWaiting().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallRequested().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCall().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallDiscarded().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallAccepted().encode()
	}
	return nil

}
func (e TypePhoneConnection) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePhoneCallProtocol) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePhonePhoneCall) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePhoneCallDiscardReason) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPhoneCallDiscardReasonMissed().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallDiscardReasonMissed().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallDiscardReasonDisconnect().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallDiscardReasonHangup().encode()
	} else if e.GetValue() != nil {
		return e.GetPhoneCallDiscardReasonBusy().encode()
	}
	return nil

}
func (e TypeInvoice) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePaymentsPaymentForm) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePostAddress) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePaymentRequestedInfo) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeDataJSON) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeLabeledPrice) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePaymentCharge) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePaymentSavedCredentials) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeWebDocument) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputWebDocument) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputWebFileLocation) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeUploadWebFile) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePaymentsValidatedRequestedInfo) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePaymentsPaymentResult) encode() []byte {
	if e.GetValue() != nil {
		return e.GetPaymentsPaymentResult().encode()
	} else if e.GetValue() != nil {
		return e.GetPaymentsPaymentResult().encode()
	} else if e.GetValue() != nil {
		return e.GetPaymentsPaymentVerficationNeeded().encode()
	}
	return nil

}
func (e TypePaymentsPaymentReceipt) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePaymentsSavedInfo) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputPaymentCredentials) encode() []byte {
	if e.GetValue() != nil {
		return e.GetInputPaymentCredentialsSaved().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPaymentCredentialsSaved().encode()
	} else if e.GetValue() != nil {
		return e.GetInputPaymentCredentials().encode()
	}
	return nil

}
func (e TypeAccountTmpPassword) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeShippingOption) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeUploadCdnFile) encode() []byte {
	if e.GetValue() != nil {
		return e.GetUploadCdnFileReuploadNeeded().encode()
	} else if e.GetValue() != nil {
		return e.GetUploadCdnFileReuploadNeeded().encode()
	} else if e.GetValue() != nil {
		return e.GetUploadCdnFile().encode()
	}
	return nil

}
func (e TypeCdnPublicKey) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeCdnConfig) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeInputStickerSetItem) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeLangPackString) encode() []byte {
	if e.GetValue() != nil {
		return e.GetLangPackString().encode()
	} else if e.GetValue() != nil {
		return e.GetLangPackString().encode()
	} else if e.GetValue() != nil {
		return e.GetLangPackStringPluralized().encode()
	} else if e.GetValue() != nil {
		return e.GetLangPackStringDeleted().encode()
	}
	return nil

}
func (e TypeLangPackDifference) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeLangPackLanguage) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeChannelAdminRights) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeChannelBannedRights) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeChannelAdminLogEventAction) encode() []byte {
	if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionChangeTitle().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionChangeTitle().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionChangeAbout().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionChangeUsername().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionChangePhoto().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionToggleInvites().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionToggleSignatures().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionUpdatePinned().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionEditMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionDeleteMessage().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionParticipantJoin().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionParticipantLeave().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionParticipantInvite().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionParticipantToggleBan().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionParticipantToggleAdmin().encode()
	} else if e.GetValue() != nil {
		return e.GetChannelAdminLogEventActionChangeStickerSet().encode()
	}
	return nil

}
func (e TypeChannelAdminLogEvent) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeChannelsAdminLogResults) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeChannelAdminLogEventsFilter) encode() []byte {
	return e.GetValue().encode()
}
func (e TypePopularContact) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeCdnFileHash) encode() []byte {
	return e.GetValue().encode()
}
func (e TypeMessagesFavedStickers) encode() []byte {
	if e.GetValue() != nil {
		return e.GetMessagesFavedStickers().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesFavedStickers().encode()
	} else if e.GetValue() != nil {
		return e.GetMessagesFavedStickersNotModified().encode()
	}
	return nil

}


// Encode funcs for predicates
func (e PredBoolFalse) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_boolFalse)
return x.buf
}

func (e PredBoolTrue) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_boolTrue)
return x.buf
}

func (e PredError) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_error)
x.Int(e.Code)
x.String(e.Text)
return x.buf
}

func (e PredNull) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_null)
return x.buf
}

func (e PredVector) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_vector)
x.Vector(e.T)
return x.buf
}

func (e PredInputPeerEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPeerEmpty)
return x.buf
}

func (e PredInputPeerSelf) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPeerSelf)
return x.buf
}

func (e PredInputPeerChat) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPeerChat)
x.Int(e.Chat_id)
return x.buf
}

func (e PredInputUserEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputUserEmpty)
return x.buf
}

func (e PredInputUserSelf) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputUserSelf)
return x.buf
}

func (e PredInputPhoneContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPhoneContact)
x.Long(e.Client_id)
x.String(e.Phone)
x.String(e.First_name)
x.String(e.Last_name)
return x.buf
}

func (e PredInputFile) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputFile)
x.Long(e.Id)
x.Int(e.Parts)
x.String(e.Name)
x.String(e.Md5_checksum)
return x.buf
}

func (e PredInputMediaEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaEmpty)
return x.buf
}

func (e PredInputMediaUploadedPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaUploadedPhoto)
x.Int(e.Flags)
x.Bytes(e.File.encode())
x.String(e.Caption)
x.Vector(e.Stickers)
x.Int(e.Ttl_seconds)
return x.buf
}

func (e PredInputMediaPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaPhoto)
x.Int(e.Flags)
x.Bytes(e.Id.encode())
x.String(e.Caption)
x.Int(e.Ttl_seconds)
return x.buf
}

func (e PredInputMediaGeoPoint) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaGeoPoint)
x.Bytes(e.Geo_point.encode())
return x.buf
}

func (e PredInputMediaContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaContact)
x.String(e.Phone_number)
x.String(e.First_name)
x.String(e.Last_name)
return x.buf
}

func (e PredInputChatPhotoEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputChatPhotoEmpty)
return x.buf
}

func (e PredInputChatUploadedPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputChatUploadedPhoto)
x.Bytes(e.File.encode())
return x.buf
}

func (e PredInputChatPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputChatPhoto)
x.Bytes(e.Id.encode())
return x.buf
}

func (e PredInputGeoPointEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputGeoPointEmpty)
return x.buf
}

func (e PredInputGeoPoint) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputGeoPoint)
x.Double(e.Lat)
x.Double(e.Long)
return x.buf
}

func (e PredInputPhotoEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPhotoEmpty)
return x.buf
}

func (e PredInputPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPhoto)
x.Long(e.Id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredInputFileLocation) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputFileLocation)
x.Long(e.Volume_id)
x.Int(e.Local_id)
x.Long(e.Secret)
return x.buf
}

func (e PredInputAppEvent) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputAppEvent)
x.Double(e.Time)
x.String(e.Type)
x.Long(e.Peer)
x.String(e.Data)
return x.buf
}

func (e PredPeerUser) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_peerUser)
x.Int(e.User_id)
return x.buf
}

func (e PredPeerChat) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_peerChat)
x.Int(e.Chat_id)
return x.buf
}

func (e PredStorageFileUnknown) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFileUnknown)
return x.buf
}

func (e PredStorageFileJpeg) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFileJpeg)
return x.buf
}

func (e PredStorageFileGif) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFileGif)
return x.buf
}

func (e PredStorageFilePng) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFilePng)
return x.buf
}

func (e PredStorageFileMp3) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFileMp3)
return x.buf
}

func (e PredStorageFileMov) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFileMov)
return x.buf
}

func (e PredStorageFilePartial) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFilePartial)
return x.buf
}

func (e PredStorageFileMp4) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFileMp4)
return x.buf
}

func (e PredStorageFileWebp) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFileWebp)
return x.buf
}

func (e PredFileLocationUnavailable) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_fileLocationUnavailable)
x.Long(e.Volume_id)
x.Int(e.Local_id)
x.Long(e.Secret)
return x.buf
}

func (e PredFileLocation) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_fileLocation)
x.Int(e.Dc_id)
x.Long(e.Volume_id)
x.Int(e.Local_id)
x.Long(e.Secret)
return x.buf
}

func (e PredUserEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userEmpty)
x.Int(e.Id)
return x.buf
}

func (e PredUserProfilePhotoEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userProfilePhotoEmpty)
return x.buf
}

func (e PredUserProfilePhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userProfilePhoto)
x.Long(e.Photo_id)
x.Bytes(e.Photo_small.encode())
x.Bytes(e.Photo_big.encode())
return x.buf
}

func (e PredUserStatusEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userStatusEmpty)
return x.buf
}

func (e PredUserStatusOnline) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userStatusOnline)
x.Int(e.Expires)
return x.buf
}

func (e PredUserStatusOffline) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userStatusOffline)
x.Int(e.Was_online)
return x.buf
}

func (e PredChatEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatEmpty)
x.Int(e.Id)
return x.buf
}

func (e PredChat) encode() []byte {
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

func (e PredChatForbidden) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatForbidden)
x.Int(e.Id)
x.String(e.Title)
return x.buf
}

func (e PredChatFull) encode() []byte {
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

func (e PredChatParticipant) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatParticipant)
x.Int(e.User_id)
x.Int(e.Inviter_id)
x.Int(e.Date)
return x.buf
}

func (e PredChatParticipantsForbidden) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatParticipantsForbidden)
x.Int(e.Flags)
x.Int(e.Chat_id)
x.Bytes(e.Self_participant.encode())
return x.buf
}

func (e PredChatParticipants) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatParticipants)
x.Int(e.Chat_id)
x.Vector(e.Participants)
x.Int(e.Version)
return x.buf
}

func (e PredChatPhotoEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatPhotoEmpty)
return x.buf
}

func (e PredChatPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatPhoto)
x.Bytes(e.Photo_small.encode())
x.Bytes(e.Photo_big.encode())
return x.buf
}

func (e PredMessageEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEmpty)
x.Int(e.Id)
return x.buf
}

func (e PredMessage) encode() []byte {
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

func (e PredMessageService) encode() []byte {
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

func (e PredMessageMediaEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaEmpty)
return x.buf
}

func (e PredMessageMediaPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaPhoto)
x.Int(e.Flags)
x.Bytes(e.Photo.encode())
x.String(e.Caption)
x.Int(e.Ttl_seconds)
return x.buf
}

func (e PredMessageMediaGeo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaGeo)
x.Bytes(e.Geo.encode())
return x.buf
}

func (e PredMessageMediaContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaContact)
x.String(e.Phone_number)
x.String(e.First_name)
x.String(e.Last_name)
x.Int(e.User_id)
return x.buf
}

func (e PredMessageMediaUnsupported) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaUnsupported)
return x.buf
}

func (e PredMessageActionEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionEmpty)
return x.buf
}

func (e PredMessageActionChatCreate) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChatCreate)
x.String(e.Title)
x.VectorInt(e.Users)
return x.buf
}

func (e PredMessageActionChatEditTitle) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChatEditTitle)
x.String(e.Title)
return x.buf
}

func (e PredMessageActionChatEditPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChatEditPhoto)
x.Bytes(e.Photo.encode())
return x.buf
}

func (e PredMessageActionChatDeletePhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChatDeletePhoto)
return x.buf
}

func (e PredMessageActionChatAddUser) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChatAddUser)
x.VectorInt(e.Users)
return x.buf
}

func (e PredMessageActionChatDeleteUser) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChatDeleteUser)
x.Int(e.User_id)
return x.buf
}

func (e PredDialog) encode() []byte {
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

func (e PredPhotoEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_photoEmpty)
x.Long(e.Id)
return x.buf
}

func (e PredPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_photo)
x.Int(e.Flags)
x.Long(e.Id)
x.Long(e.Access_hash)
x.Int(e.Date)
x.Vector(e.Sizes)
return x.buf
}

func (e PredPhotoSizeEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_photoSizeEmpty)
x.String(e.Type)
return x.buf
}

func (e PredPhotoSize) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_photoSize)
x.String(e.Type)
x.Bytes(e.Location.encode())
x.Int(e.W)
x.Int(e.H)
x.Int(e.Size)
return x.buf
}

func (e PredPhotoCachedSize) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_photoCachedSize)
x.String(e.Type)
x.Bytes(e.Location.encode())
x.Int(e.W)
x.Int(e.H)
x.StringBytes(e.Bytes)
return x.buf
}

func (e PredGeoPointEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_geoPointEmpty)
return x.buf
}

func (e PredGeoPoint) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_geoPoint)
x.Double(e.Long)
x.Double(e.Lat)
return x.buf
}

func (e PredAuthCheckedPhone) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authCheckedPhone)
x.Bytes(e.Phone_registered.encode())
return x.buf
}

func (e PredAuthSentCode) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authSentCode)
x.Int(e.Flags)
x.Bytes(e.Type.encode())
x.String(e.Phone_code_hash)
x.Bytes(e.Next_type.encode())
x.Int(e.Timeout)
return x.buf
}

func (e PredAuthAuthorization) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authAuthorization)
x.Int(e.Flags)
x.Int(e.Tmp_sessions)
x.Bytes(e.User.encode())
return x.buf
}

func (e PredAuthExportedAuthorization) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authExportedAuthorization)
x.Int(e.Id)
x.StringBytes(e.Bytes)
return x.buf
}

func (e PredInputNotifyPeer) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputNotifyPeer)
x.Bytes(e.Peer.encode())
return x.buf
}

func (e PredInputNotifyUsers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputNotifyUsers)
return x.buf
}

func (e PredInputNotifyChats) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputNotifyChats)
return x.buf
}

func (e PredInputNotifyAll) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputNotifyAll)
return x.buf
}

func (e PredInputPeerNotifySettings) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPeerNotifySettings)
x.Int(e.Flags)
x.Int(e.Mute_until)
x.String(e.Sound)
return x.buf
}

func (e PredPeerNotifyEventsEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_peerNotifyEventsEmpty)
return x.buf
}

func (e PredPeerNotifyEventsAll) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_peerNotifyEventsAll)
return x.buf
}

func (e PredPeerNotifySettingsEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_peerNotifySettingsEmpty)
return x.buf
}

func (e PredPeerNotifySettings) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_peerNotifySettings)
x.Int(e.Flags)
x.Int(e.Mute_until)
x.String(e.Sound)
return x.buf
}

func (e PredWallPaper) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_wallPaper)
x.Int(e.Id)
x.String(e.Title)
x.Vector(e.Sizes)
x.Int(e.Color)
return x.buf
}

func (e PredUserFull) encode() []byte {
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

func (e PredContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contact)
x.Int(e.User_id)
x.Bytes(e.Mutual.encode())
return x.buf
}

func (e PredImportedContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_importedContact)
x.Int(e.User_id)
x.Long(e.Client_id)
return x.buf
}

func (e PredContactBlocked) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactBlocked)
x.Int(e.User_id)
x.Int(e.Date)
return x.buf
}

func (e PredContactStatus) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactStatus)
x.Int(e.User_id)
x.Bytes(e.Status.encode())
return x.buf
}

func (e PredContactsLink) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsLink)
x.Bytes(e.My_link.encode())
x.Bytes(e.Foreign_link.encode())
x.Bytes(e.User.encode())
return x.buf
}

func (e PredContactsContacts) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsContacts)
x.Vector(e.Contacts)
x.Int(e.Saved_count)
x.Vector(e.Users)
return x.buf
}

func (e PredContactsContactsNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsContactsNotModified)
return x.buf
}

func (e PredContactsImportedContacts) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsImportedContacts)
x.Vector(e.Imported)
x.Vector(e.Popular_invites)
x.VectorLong(e.Retry_contacts)
x.Vector(e.Users)
return x.buf
}

func (e PredContactsBlocked) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsBlocked)
x.Vector(e.Blocked)
x.Vector(e.Users)
return x.buf
}

func (e PredContactsBlockedSlice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsBlockedSlice)
x.Int(e.Count)
x.Vector(e.Blocked)
x.Vector(e.Users)
return x.buf
}

func (e PredContactsFound) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsFound)
x.Vector(e.Results)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredMessagesDialogs) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesDialogs)
x.Vector(e.Dialogs)
x.Vector(e.Messages)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredMessagesDialogsSlice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesDialogsSlice)
x.Int(e.Count)
x.Vector(e.Dialogs)
x.Vector(e.Messages)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredMessagesMessages) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesMessages)
x.Vector(e.Messages)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredMessagesMessagesSlice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesMessagesSlice)
x.Int(e.Count)
x.Vector(e.Messages)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredMessagesChats) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesChats)
x.Vector(e.Chats)
return x.buf
}

func (e PredMessagesChatFull) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesChatFull)
x.Bytes(e.Full_chat.encode())
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredMessagesAffectedHistory) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesAffectedHistory)
x.Int(e.Pts)
x.Int(e.Pts_count)
x.Int(e.Offset)
return x.buf
}

func (e PredInputMessagesFilterEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterEmpty)
return x.buf
}

func (e PredInputMessagesFilterPhotos) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterPhotos)
return x.buf
}

func (e PredInputMessagesFilterVideo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterVideo)
return x.buf
}

func (e PredInputMessagesFilterPhotoVideo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterPhotoVideo)
return x.buf
}

func (e PredUpdateNewMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateNewMessage)
x.Bytes(e.Message.encode())
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredUpdateMessageID) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateMessageID)
x.Int(e.Id)
x.Long(e.Random_id)
return x.buf
}

func (e PredUpdateDeleteMessages) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateDeleteMessages)
x.VectorInt(e.Messages)
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredUpdateUserTyping) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateUserTyping)
x.Int(e.User_id)
x.Bytes(e.Action.encode())
return x.buf
}

func (e PredUpdateChatUserTyping) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChatUserTyping)
x.Int(e.Chat_id)
x.Int(e.User_id)
x.Bytes(e.Action.encode())
return x.buf
}

func (e PredUpdateChatParticipants) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChatParticipants)
x.Bytes(e.Participants.encode())
return x.buf
}

func (e PredUpdateUserStatus) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateUserStatus)
x.Int(e.User_id)
x.Bytes(e.Status.encode())
return x.buf
}

func (e PredUpdateUserName) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateUserName)
x.Int(e.User_id)
x.String(e.First_name)
x.String(e.Last_name)
x.String(e.Username)
return x.buf
}

func (e PredUpdateUserPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateUserPhoto)
x.Int(e.User_id)
x.Int(e.Date)
x.Bytes(e.Photo.encode())
x.Bytes(e.Previous.encode())
return x.buf
}

func (e PredUpdateContactRegistered) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateContactRegistered)
x.Int(e.User_id)
x.Int(e.Date)
return x.buf
}

func (e PredUpdateContactLink) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateContactLink)
x.Int(e.User_id)
x.Bytes(e.My_link.encode())
x.Bytes(e.Foreign_link.encode())
return x.buf
}

func (e PredUpdatesState) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesState)
x.Int(e.Pts)
x.Int(e.Qts)
x.Int(e.Date)
x.Int(e.Seq)
x.Int(e.Unread_count)
return x.buf
}

func (e PredUpdatesDifferenceEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesDifferenceEmpty)
x.Int(e.Date)
x.Int(e.Seq)
return x.buf
}

func (e PredUpdatesDifference) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesDifference)
x.Vector(e.New_messages)
x.Vector(e.New_encrypted_messages)
x.Vector(e.Other_updates)
x.Vector(e.Chats)
x.Vector(e.Users)
x.Bytes(e.State.encode())
return x.buf
}

func (e PredUpdatesDifferenceSlice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesDifferenceSlice)
x.Vector(e.New_messages)
x.Vector(e.New_encrypted_messages)
x.Vector(e.Other_updates)
x.Vector(e.Chats)
x.Vector(e.Users)
x.Bytes(e.Intermediate_state.encode())
return x.buf
}

func (e PredUpdatesTooLong) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesTooLong)
return x.buf
}

func (e PredUpdateShortMessage) encode() []byte {
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

func (e PredUpdateShortChatMessage) encode() []byte {
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

func (e PredUpdateShort) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateShort)
x.Bytes(e.Update.encode())
x.Int(e.Date)
return x.buf
}

func (e PredUpdatesCombined) encode() []byte {
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

func (e PredUpdates) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updates)
x.Vector(e.Updates)
x.Vector(e.Users)
x.Vector(e.Chats)
x.Int(e.Date)
x.Int(e.Seq)
return x.buf
}

func (e PredPhotosPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_photosPhoto)
x.Bytes(e.Photo.encode())
x.Vector(e.Users)
return x.buf
}

func (e PredUploadFile) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_uploadFile)
x.Bytes(e.Type.encode())
x.Int(e.Mtime)
x.StringBytes(e.Bytes)
return x.buf
}

func (e PredDcOption) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_dcOption)
x.Int(e.Flags)
x.Int(e.Id)
x.String(e.Ip_address)
x.Int(e.Port)
return x.buf
}

func (e PredConfig) encode() []byte {
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

func (e PredNearestDc) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_nearestDc)
x.String(e.Country)
x.Int(e.This_dc)
x.Int(e.Nearest_dc)
return x.buf
}

func (e PredHelpAppUpdate) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_helpAppUpdate)
x.Int(e.Id)
x.Bytes(e.Critical.encode())
x.String(e.Url)
x.String(e.Text)
return x.buf
}

func (e PredHelpNoAppUpdate) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_helpNoAppUpdate)
return x.buf
}

func (e PredHelpInviteText) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_helpInviteText)
x.String(e.Message)
return x.buf
}

func (e PredInputPeerNotifyEventsEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPeerNotifyEventsEmpty)
return x.buf
}

func (e PredInputPeerNotifyEventsAll) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPeerNotifyEventsAll)
return x.buf
}

func (e PredPhotosPhotos) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_photosPhotos)
x.Vector(e.Photos)
x.Vector(e.Users)
return x.buf
}

func (e PredPhotosPhotosSlice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_photosPhotosSlice)
x.Int(e.Count)
x.Vector(e.Photos)
x.Vector(e.Users)
return x.buf
}

func (e PredWallPaperSolid) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_wallPaperSolid)
x.Int(e.Id)
x.String(e.Title)
x.Int(e.Bg_color)
x.Int(e.Color)
return x.buf
}

func (e PredUpdateNewEncryptedMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateNewEncryptedMessage)
x.Bytes(e.Message.encode())
x.Int(e.Qts)
return x.buf
}

func (e PredUpdateEncryptedChatTyping) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateEncryptedChatTyping)
x.Int(e.Chat_id)
return x.buf
}

func (e PredUpdateEncryption) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateEncryption)
x.Bytes(e.Chat.encode())
x.Int(e.Date)
return x.buf
}

func (e PredUpdateEncryptedMessagesRead) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateEncryptedMessagesRead)
x.Int(e.Chat_id)
x.Int(e.Max_date)
x.Int(e.Date)
return x.buf
}

func (e PredEncryptedChatEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_encryptedChatEmpty)
x.Int(e.Id)
return x.buf
}

func (e PredEncryptedChatWaiting) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_encryptedChatWaiting)
x.Int(e.Id)
x.Long(e.Access_hash)
x.Int(e.Date)
x.Int(e.Admin_id)
x.Int(e.Participant_id)
return x.buf
}

func (e PredEncryptedChatRequested) encode() []byte {
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

func (e PredEncryptedChat) encode() []byte {
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

func (e PredEncryptedChatDiscarded) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_encryptedChatDiscarded)
x.Int(e.Id)
return x.buf
}

func (e PredInputEncryptedChat) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputEncryptedChat)
x.Int(e.Chat_id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredEncryptedFileEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_encryptedFileEmpty)
return x.buf
}

func (e PredEncryptedFile) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_encryptedFile)
x.Long(e.Id)
x.Long(e.Access_hash)
x.Int(e.Size)
x.Int(e.Dc_id)
x.Int(e.Key_fingerprint)
return x.buf
}

func (e PredInputEncryptedFileEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputEncryptedFileEmpty)
return x.buf
}

func (e PredInputEncryptedFileUploaded) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputEncryptedFileUploaded)
x.Long(e.Id)
x.Int(e.Parts)
x.String(e.Md5_checksum)
x.Int(e.Key_fingerprint)
return x.buf
}

func (e PredInputEncryptedFile) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputEncryptedFile)
x.Long(e.Id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredInputEncryptedFileLocation) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputEncryptedFileLocation)
x.Long(e.Id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredEncryptedMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_encryptedMessage)
x.Long(e.Random_id)
x.Int(e.Chat_id)
x.Int(e.Date)
x.StringBytes(e.Bytes)
x.Bytes(e.File.encode())
return x.buf
}

func (e PredEncryptedMessageService) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_encryptedMessageService)
x.Long(e.Random_id)
x.Int(e.Chat_id)
x.Int(e.Date)
x.StringBytes(e.Bytes)
return x.buf
}

func (e PredMessagesDhConfigNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesDhConfigNotModified)
x.StringBytes(e.Random)
return x.buf
}

func (e PredMessagesDhConfig) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesDhConfig)
x.Int(e.G)
x.StringBytes(e.P)
x.Int(e.Version)
x.StringBytes(e.Random)
return x.buf
}

func (e PredMessagesSentEncryptedMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesSentEncryptedMessage)
x.Int(e.Date)
return x.buf
}

func (e PredMessagesSentEncryptedFile) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesSentEncryptedFile)
x.Int(e.Date)
x.Bytes(e.File.encode())
return x.buf
}

func (e PredInputFileBig) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputFileBig)
x.Long(e.Id)
x.Int(e.Parts)
x.String(e.Name)
return x.buf
}

func (e PredInputEncryptedFileBigUploaded) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputEncryptedFileBigUploaded)
x.Long(e.Id)
x.Int(e.Parts)
x.Int(e.Key_fingerprint)
return x.buf
}

func (e PredStorageFilePdf) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_storageFilePdf)
return x.buf
}

func (e PredInputMessagesFilterDocument) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterDocument)
return x.buf
}

func (e PredInputMessagesFilterPhotoVideoDocuments) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterPhotoVideoDocuments)
return x.buf
}

func (e PredUpdateChatParticipantAdd) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChatParticipantAdd)
x.Int(e.Chat_id)
x.Int(e.User_id)
x.Int(e.Inviter_id)
x.Int(e.Date)
x.Int(e.Version)
return x.buf
}

func (e PredUpdateChatParticipantDelete) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChatParticipantDelete)
x.Int(e.Chat_id)
x.Int(e.User_id)
x.Int(e.Version)
return x.buf
}

func (e PredUpdateDcOptions) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateDcOptions)
x.Vector(e.Dc_options)
return x.buf
}

func (e PredInputMediaUploadedDocument) encode() []byte {
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

func (e PredInputMediaDocument) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaDocument)
x.Int(e.Flags)
x.Bytes(e.Id.encode())
x.String(e.Caption)
x.Int(e.Ttl_seconds)
return x.buf
}

func (e PredMessageMediaDocument) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaDocument)
x.Int(e.Flags)
x.Bytes(e.Document.encode())
x.String(e.Caption)
x.Int(e.Ttl_seconds)
return x.buf
}

func (e PredInputDocumentEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputDocumentEmpty)
return x.buf
}

func (e PredInputDocument) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputDocument)
x.Long(e.Id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredInputDocumentFileLocation) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputDocumentFileLocation)
x.Long(e.Id)
x.Long(e.Access_hash)
x.Int(e.Version)
return x.buf
}

func (e PredDocumentEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_documentEmpty)
x.Long(e.Id)
return x.buf
}

func (e PredDocument) encode() []byte {
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

func (e PredHelpSupport) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_helpSupport)
x.String(e.Phone_number)
x.Bytes(e.User.encode())
return x.buf
}

func (e PredNotifyAll) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_notifyAll)
return x.buf
}

func (e PredNotifyChats) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_notifyChats)
return x.buf
}

func (e PredNotifyPeer) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_notifyPeer)
x.Bytes(e.Peer.encode())
return x.buf
}

func (e PredNotifyUsers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_notifyUsers)
return x.buf
}

func (e PredUpdateUserBlocked) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateUserBlocked)
x.Int(e.User_id)
x.Bytes(e.Blocked.encode())
return x.buf
}

func (e PredUpdateNotifySettings) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateNotifySettings)
x.Bytes(e.Peer.encode())
x.Bytes(e.Notify_settings.encode())
return x.buf
}

func (e PredSendMessageTypingAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageTypingAction)
return x.buf
}

func (e PredSendMessageCancelAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageCancelAction)
return x.buf
}

func (e PredSendMessageRecordVideoAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageRecordVideoAction)
return x.buf
}

func (e PredSendMessageUploadVideoAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageUploadVideoAction)
x.Int(e.Progress)
return x.buf
}

func (e PredSendMessageRecordAudioAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageRecordAudioAction)
return x.buf
}

func (e PredSendMessageUploadAudioAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageUploadAudioAction)
x.Int(e.Progress)
return x.buf
}

func (e PredSendMessageUploadPhotoAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageUploadPhotoAction)
x.Int(e.Progress)
return x.buf
}

func (e PredSendMessageUploadDocumentAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageUploadDocumentAction)
x.Int(e.Progress)
return x.buf
}

func (e PredSendMessageGeoLocationAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageGeoLocationAction)
return x.buf
}

func (e PredSendMessageChooseContactAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageChooseContactAction)
return x.buf
}

func (e PredUpdateServiceNotification) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateServiceNotification)
x.Int(e.Flags)
x.Int(e.Inbox_date)
x.String(e.Type)
x.String(e.Message)
x.Bytes(e.Media.encode())
x.Vector(e.Entities)
return x.buf
}

func (e PredUserStatusRecently) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userStatusRecently)
return x.buf
}

func (e PredUserStatusLastWeek) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userStatusLastWeek)
return x.buf
}

func (e PredUserStatusLastMonth) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_userStatusLastMonth)
return x.buf
}

func (e PredUpdatePrivacy) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatePrivacy)
x.Bytes(e.Key.encode())
x.Vector(e.Rules)
return x.buf
}

func (e PredInputPrivacyKeyStatusTimestamp) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyKeyStatusTimestamp)
return x.buf
}

func (e PredPrivacyKeyStatusTimestamp) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyKeyStatusTimestamp)
return x.buf
}

func (e PredInputPrivacyValueAllowContacts) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyValueAllowContacts)
return x.buf
}

func (e PredInputPrivacyValueAllowAll) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyValueAllowAll)
return x.buf
}

func (e PredInputPrivacyValueAllowUsers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyValueAllowUsers)
x.Vector(e.Users)
return x.buf
}

func (e PredInputPrivacyValueDisallowContacts) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyValueDisallowContacts)
return x.buf
}

func (e PredInputPrivacyValueDisallowAll) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyValueDisallowAll)
return x.buf
}

func (e PredInputPrivacyValueDisallowUsers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyValueDisallowUsers)
x.Vector(e.Users)
return x.buf
}

func (e PredPrivacyValueAllowContacts) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyValueAllowContacts)
return x.buf
}

func (e PredPrivacyValueAllowAll) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyValueAllowAll)
return x.buf
}

func (e PredPrivacyValueAllowUsers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyValueAllowUsers)
x.VectorInt(e.Users)
return x.buf
}

func (e PredPrivacyValueDisallowContacts) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyValueDisallowContacts)
return x.buf
}

func (e PredPrivacyValueDisallowAll) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyValueDisallowAll)
return x.buf
}

func (e PredPrivacyValueDisallowUsers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyValueDisallowUsers)
x.VectorInt(e.Users)
return x.buf
}

func (e PredAccountPrivacyRules) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_accountPrivacyRules)
x.Vector(e.Rules)
x.Vector(e.Users)
return x.buf
}

func (e PredAccountDaysTTL) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_accountDaysTTL)
x.Int(e.Days)
return x.buf
}

func (e PredUpdateUserPhone) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateUserPhone)
x.Int(e.User_id)
x.String(e.Phone)
return x.buf
}

func (e PredDisabledFeature) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_disabledFeature)
x.String(e.Feature)
x.String(e.Description)
return x.buf
}

func (e PredDocumentAttributeImageSize) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_documentAttributeImageSize)
x.Int(e.W)
x.Int(e.H)
return x.buf
}

func (e PredDocumentAttributeAnimated) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_documentAttributeAnimated)
return x.buf
}

func (e PredDocumentAttributeSticker) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_documentAttributeSticker)
x.Int(e.Flags)
x.String(e.Alt)
x.Bytes(e.Stickerset.encode())
x.Bytes(e.Mask_coords.encode())
return x.buf
}

func (e PredDocumentAttributeVideo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_documentAttributeVideo)
x.Int(e.Flags)
x.Int(e.Duration)
x.Int(e.W)
x.Int(e.H)
return x.buf
}

func (e PredDocumentAttributeAudio) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_documentAttributeAudio)
x.Int(e.Flags)
x.Int(e.Duration)
x.String(e.Title)
x.String(e.Performer)
x.StringBytes(e.Waveform)
return x.buf
}

func (e PredDocumentAttributeFilename) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_documentAttributeFilename)
x.String(e.File_name)
return x.buf
}

func (e PredMessagesStickersNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesStickersNotModified)
return x.buf
}

func (e PredMessagesStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesStickers)
x.String(e.Hash)
x.Vector(e.Stickers)
return x.buf
}

func (e PredStickerPack) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_stickerPack)
x.String(e.Emoticon)
x.VectorLong(e.Documents)
return x.buf
}

func (e PredMessagesAllStickersNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesAllStickersNotModified)
return x.buf
}

func (e PredMessagesAllStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesAllStickers)
x.Int(e.Hash)
x.Vector(e.Sets)
return x.buf
}

func (e PredAccountNoPassword) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_accountNoPassword)
x.StringBytes(e.New_salt)
x.String(e.Email_unconfirmed_pattern)
return x.buf
}

func (e PredAccountPassword) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_accountPassword)
x.StringBytes(e.Current_salt)
x.StringBytes(e.New_salt)
x.String(e.Hint)
x.Bytes(e.Has_recovery.encode())
x.String(e.Email_unconfirmed_pattern)
return x.buf
}

func (e PredUpdateReadHistoryInbox) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateReadHistoryInbox)
x.Bytes(e.Peer.encode())
x.Int(e.Max_id)
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredUpdateReadHistoryOutbox) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateReadHistoryOutbox)
x.Bytes(e.Peer.encode())
x.Int(e.Max_id)
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredMessagesAffectedMessages) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesAffectedMessages)
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredContactLinkUnknown) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactLinkUnknown)
return x.buf
}

func (e PredContactLinkNone) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactLinkNone)
return x.buf
}

func (e PredContactLinkHasPhone) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactLinkHasPhone)
return x.buf
}

func (e PredContactLinkContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactLinkContact)
return x.buf
}

func (e PredUpdateWebPage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateWebPage)
x.Bytes(e.Webpage.encode())
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredWebPageEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_webPageEmpty)
x.Long(e.Id)
return x.buf
}

func (e PredWebPagePending) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_webPagePending)
x.Long(e.Id)
x.Int(e.Date)
return x.buf
}

func (e PredWebPage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_webPage)
x.Int(e.Flags)
x.Long(e.Id)
x.String(e.Url)
x.String(e.Display_url)
x.Int(e.Hash)
x.String(e.Type)
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

func (e PredMessageMediaWebPage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaWebPage)
x.Bytes(e.Webpage.encode())
return x.buf
}

func (e PredAuthorization) encode() []byte {
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

func (e PredAccountAuthorizations) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_accountAuthorizations)
x.Vector(e.Authorizations)
return x.buf
}

func (e PredAccountPasswordSettings) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_accountPasswordSettings)
x.String(e.Email)
return x.buf
}

func (e PredAccountPasswordInputSettings) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_accountPasswordInputSettings)
x.Int(e.Flags)
x.StringBytes(e.New_salt)
x.StringBytes(e.New_password_hash)
x.String(e.Hint)
x.String(e.Email)
return x.buf
}

func (e PredAuthPasswordRecovery) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authPasswordRecovery)
x.String(e.Email_pattern)
return x.buf
}

func (e PredInputMediaVenue) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaVenue)
x.Bytes(e.Geo_point.encode())
x.String(e.Title)
x.String(e.Address)
x.String(e.Provider)
x.String(e.Venue_id)
return x.buf
}

func (e PredMessageMediaVenue) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaVenue)
x.Bytes(e.Geo.encode())
x.String(e.Title)
x.String(e.Address)
x.String(e.Provider)
x.String(e.Venue_id)
return x.buf
}

func (e PredReceivedNotifyMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_receivedNotifyMessage)
x.Int(e.Id)
x.Int(e.Flags)
return x.buf
}

func (e PredChatInviteEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatInviteEmpty)
return x.buf
}

func (e PredChatInviteExported) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatInviteExported)
x.String(e.Link)
return x.buf
}

func (e PredChatInviteAlready) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatInviteAlready)
x.Bytes(e.Chat.encode())
return x.buf
}

func (e PredChatInvite) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatInvite)
x.Int(e.Flags)
x.String(e.Title)
x.Bytes(e.Photo.encode())
x.Int(e.Participants_count)
x.Vector(e.Participants)
return x.buf
}

func (e PredMessageActionChatJoinedByLink) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChatJoinedByLink)
x.Int(e.Inviter_id)
return x.buf
}

func (e PredUpdateReadMessagesContents) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateReadMessagesContents)
x.VectorInt(e.Messages)
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredInputStickerSetEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputStickerSetEmpty)
return x.buf
}

func (e PredInputStickerSetID) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputStickerSetID)
x.Long(e.Id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredInputStickerSetShortName) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputStickerSetShortName)
x.String(e.Short_name)
return x.buf
}

func (e PredStickerSet) encode() []byte {
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

func (e PredMessagesStickerSet) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesStickerSet)
x.Bytes(e.Set.encode())
x.Vector(e.Packs)
x.Vector(e.Documents)
return x.buf
}

func (e PredUser) encode() []byte {
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

func (e PredBotCommand) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_botCommand)
x.String(e.Command)
x.String(e.Description)
return x.buf
}

func (e PredBotInfo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_botInfo)
x.Int(e.User_id)
x.String(e.Description)
x.Vector(e.Commands)
return x.buf
}

func (e PredKeyboardButton) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButton)
x.String(e.Text)
return x.buf
}

func (e PredKeyboardButtonRow) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButtonRow)
x.Vector(e.Buttons)
return x.buf
}

func (e PredReplyKeyboardHide) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_replyKeyboardHide)
x.Int(e.Flags)
return x.buf
}

func (e PredReplyKeyboardForceReply) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_replyKeyboardForceReply)
x.Int(e.Flags)
return x.buf
}

func (e PredReplyKeyboardMarkup) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_replyKeyboardMarkup)
x.Int(e.Flags)
x.Vector(e.Rows)
return x.buf
}

func (e PredInputMessagesFilterUrl) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterUrl)
return x.buf
}

func (e PredInputPeerUser) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPeerUser)
x.Int(e.User_id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredInputUser) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputUser)
x.Int(e.User_id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredMessageEntityUnknown) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityUnknown)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityMention) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityMention)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityHashtag) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityHashtag)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityBotCommand) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityBotCommand)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityUrl) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityUrl)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityEmail) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityEmail)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityBold) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityBold)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityItalic) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityItalic)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityCode) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityCode)
x.Int(e.Offset)
x.Int(e.Length)
return x.buf
}

func (e PredMessageEntityPre) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityPre)
x.Int(e.Offset)
x.Int(e.Length)
x.String(e.Language)
return x.buf
}

func (e PredMessageEntityTextUrl) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityTextUrl)
x.Int(e.Offset)
x.Int(e.Length)
x.String(e.Url)
return x.buf
}

func (e PredUpdateShortSentMessage) encode() []byte {
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

func (e PredInputPeerChannel) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPeerChannel)
x.Int(e.Channel_id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredPeerChannel) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_peerChannel)
x.Int(e.Channel_id)
return x.buf
}

func (e PredChannel) encode() []byte {
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

func (e PredChannelForbidden) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelForbidden)
x.Int(e.Flags)
x.Int(e.Id)
x.Long(e.Access_hash)
x.String(e.Title)
x.Int(e.Until_date)
return x.buf
}

func (e PredChannelFull) encode() []byte {
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

func (e PredMessageActionChannelCreate) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChannelCreate)
x.String(e.Title)
return x.buf
}

func (e PredMessagesChannelMessages) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesChannelMessages)
x.Int(e.Flags)
x.Int(e.Pts)
x.Int(e.Count)
x.Vector(e.Messages)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredUpdateChannelTooLong) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChannelTooLong)
x.Int(e.Flags)
x.Int(e.Channel_id)
x.Int(e.Pts)
return x.buf
}

func (e PredUpdateChannel) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChannel)
x.Int(e.Channel_id)
return x.buf
}

func (e PredUpdateNewChannelMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateNewChannelMessage)
x.Bytes(e.Message.encode())
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredUpdateReadChannelInbox) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateReadChannelInbox)
x.Int(e.Channel_id)
x.Int(e.Max_id)
return x.buf
}

func (e PredUpdateDeleteChannelMessages) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateDeleteChannelMessages)
x.Int(e.Channel_id)
x.VectorInt(e.Messages)
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredUpdateChannelMessageViews) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChannelMessageViews)
x.Int(e.Channel_id)
x.Int(e.Id)
x.Int(e.Views)
return x.buf
}

func (e PredInputChannelEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputChannelEmpty)
return x.buf
}

func (e PredInputChannel) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputChannel)
x.Int(e.Channel_id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredContactsResolvedPeer) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsResolvedPeer)
x.Bytes(e.Peer.encode())
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredMessageRange) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageRange)
x.Int(e.Min_id)
x.Int(e.Max_id)
return x.buf
}

func (e PredUpdatesChannelDifferenceEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesChannelDifferenceEmpty)
x.Int(e.Flags)
x.Int(e.Pts)
x.Int(e.Timeout)
return x.buf
}

func (e PredUpdatesChannelDifferenceTooLong) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesChannelDifferenceTooLong)
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

func (e PredUpdatesChannelDifference) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesChannelDifference)
x.Int(e.Flags)
x.Int(e.Pts)
x.Int(e.Timeout)
x.Vector(e.New_messages)
x.Vector(e.Other_updates)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredChannelMessagesFilterEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelMessagesFilterEmpty)
return x.buf
}

func (e PredChannelMessagesFilter) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelMessagesFilter)
x.Int(e.Flags)
x.Vector(e.Ranges)
return x.buf
}

func (e PredChannelParticipant) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipant)
x.Int(e.User_id)
x.Int(e.Date)
return x.buf
}

func (e PredChannelParticipantSelf) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantSelf)
x.Int(e.User_id)
x.Int(e.Inviter_id)
x.Int(e.Date)
return x.buf
}

func (e PredChannelParticipantCreator) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantCreator)
x.Int(e.User_id)
return x.buf
}

func (e PredChannelParticipantsRecent) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantsRecent)
return x.buf
}

func (e PredChannelParticipantsAdmins) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantsAdmins)
return x.buf
}

func (e PredChannelParticipantsKicked) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantsKicked)
x.String(e.Q)
return x.buf
}

func (e PredChannelsChannelParticipants) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelsChannelParticipants)
x.Int(e.Count)
x.Vector(e.Participants)
x.Vector(e.Users)
return x.buf
}

func (e PredChannelsChannelParticipant) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelsChannelParticipant)
x.Bytes(e.Participant.encode())
x.Vector(e.Users)
return x.buf
}

func (e PredTrue) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_true)
return x.buf
}

func (e PredChatParticipantCreator) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatParticipantCreator)
x.Int(e.User_id)
return x.buf
}

func (e PredChatParticipantAdmin) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_chatParticipantAdmin)
x.Int(e.User_id)
x.Int(e.Inviter_id)
x.Int(e.Date)
return x.buf
}

func (e PredUpdateChatAdmins) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChatAdmins)
x.Int(e.Chat_id)
x.Bytes(e.Enabled.encode())
x.Int(e.Version)
return x.buf
}

func (e PredUpdateChatParticipantAdmin) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChatParticipantAdmin)
x.Int(e.Chat_id)
x.Int(e.User_id)
x.Bytes(e.Is_admin.encode())
x.Int(e.Version)
return x.buf
}

func (e PredMessageActionChatMigrateTo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChatMigrateTo)
x.Int(e.Channel_id)
return x.buf
}

func (e PredMessageActionChannelMigrateFrom) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionChannelMigrateFrom)
x.String(e.Title)
x.Int(e.Chat_id)
return x.buf
}

func (e PredChannelParticipantsBots) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantsBots)
return x.buf
}

func (e PredInputReportReasonSpam) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputReportReasonSpam)
return x.buf
}

func (e PredInputReportReasonViolence) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputReportReasonViolence)
return x.buf
}

func (e PredInputReportReasonPornography) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputReportReasonPornography)
return x.buf
}

func (e PredInputReportReasonOther) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputReportReasonOther)
x.String(e.Text)
return x.buf
}

func (e PredUpdateNewStickerSet) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateNewStickerSet)
x.Bytes(e.Stickerset.encode())
return x.buf
}

func (e PredUpdateStickerSetsOrder) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateStickerSetsOrder)
x.Int(e.Flags)
x.VectorLong(e.Order)
return x.buf
}

func (e PredUpdateStickerSets) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateStickerSets)
return x.buf
}

func (e PredHelpTermsOfService) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_helpTermsOfService)
x.String(e.Text)
return x.buf
}

func (e PredFoundGif) encode() []byte {
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

func (e PredInputMediaGifExternal) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaGifExternal)
x.String(e.Url)
x.String(e.Q)
return x.buf
}

func (e PredMessagesFoundGifs) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesFoundGifs)
x.Int(e.Next_offset)
x.Vector(e.Results)
return x.buf
}

func (e PredInputMessagesFilterGif) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterGif)
return x.buf
}

func (e PredUpdateSavedGifs) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateSavedGifs)
return x.buf
}

func (e PredUpdateBotInlineQuery) encode() []byte {
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

func (e PredFoundGifCached) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_foundGifCached)
x.String(e.Url)
x.Bytes(e.Photo.encode())
x.Bytes(e.Document.encode())
return x.buf
}

func (e PredMessagesSavedGifsNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesSavedGifsNotModified)
return x.buf
}

func (e PredMessagesSavedGifs) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesSavedGifs)
x.Int(e.Hash)
x.Vector(e.Gifs)
return x.buf
}

func (e PredInputBotInlineMessageMediaAuto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineMessageMediaAuto)
x.Int(e.Flags)
x.String(e.Caption)
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredInputBotInlineMessageText) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineMessageText)
x.Int(e.Flags)
x.String(e.Message)
x.Vector(e.Entities)
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredInputBotInlineResult) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineResult)
x.Int(e.Flags)
x.String(e.Id)
x.String(e.Type)
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

func (e PredBotInlineMessageMediaAuto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_botInlineMessageMediaAuto)
x.Int(e.Flags)
x.String(e.Caption)
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredBotInlineMessageText) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_botInlineMessageText)
x.Int(e.Flags)
x.String(e.Message)
x.Vector(e.Entities)
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredBotInlineResult) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_botInlineResult)
x.Int(e.Flags)
x.String(e.Id)
x.String(e.Type)
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

func (e PredMessagesBotResults) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesBotResults)
x.Int(e.Flags)
x.Long(e.Query_id)
x.String(e.Next_offset)
x.Bytes(e.Switch_pm.encode())
x.Vector(e.Results)
x.Int(e.Cache_time)
return x.buf
}

func (e PredInputMessagesFilterVoice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterVoice)
return x.buf
}

func (e PredInputMessagesFilterMusic) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterMusic)
return x.buf
}

func (e PredUpdateBotInlineSend) encode() []byte {
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

func (e PredInputPrivacyKeyChatInvite) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyKeyChatInvite)
return x.buf
}

func (e PredPrivacyKeyChatInvite) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyKeyChatInvite)
return x.buf
}

func (e PredUpdateEditChannelMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateEditChannelMessage)
x.Bytes(e.Message.encode())
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredExportedMessageLink) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_exportedMessageLink)
x.String(e.Link)
return x.buf
}

func (e PredMessageFwdHeader) encode() []byte {
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

func (e PredMessageActionPinMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionPinMessage)
return x.buf
}

func (e PredPeerSettings) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_peerSettings)
x.Int(e.Flags)
return x.buf
}

func (e PredUpdateChannelPinnedMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChannelPinnedMessage)
x.Int(e.Channel_id)
x.Int(e.Id)
return x.buf
}

func (e PredKeyboardButtonUrl) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButtonUrl)
x.String(e.Text)
x.String(e.Url)
return x.buf
}

func (e PredKeyboardButtonCallback) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButtonCallback)
x.String(e.Text)
x.StringBytes(e.Data)
return x.buf
}

func (e PredKeyboardButtonRequestPhone) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButtonRequestPhone)
x.String(e.Text)
return x.buf
}

func (e PredKeyboardButtonRequestGeoLocation) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButtonRequestGeoLocation)
x.String(e.Text)
return x.buf
}

func (e PredAuthCodeTypeSms) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authCodeTypeSms)
return x.buf
}

func (e PredAuthCodeTypeCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authCodeTypeCall)
return x.buf
}

func (e PredAuthCodeTypeFlashCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authCodeTypeFlashCall)
return x.buf
}

func (e PredAuthSentCodeTypeApp) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authSentCodeTypeApp)
x.Int(e.Length)
return x.buf
}

func (e PredAuthSentCodeTypeSms) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authSentCodeTypeSms)
x.Int(e.Length)
return x.buf
}

func (e PredAuthSentCodeTypeCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authSentCodeTypeCall)
x.Int(e.Length)
return x.buf
}

func (e PredAuthSentCodeTypeFlashCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_authSentCodeTypeFlashCall)
x.String(e.Pattern)
return x.buf
}

func (e PredKeyboardButtonSwitchInline) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButtonSwitchInline)
x.Int(e.Flags)
x.String(e.Text)
x.String(e.Query)
return x.buf
}

func (e PredReplyInlineMarkup) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_replyInlineMarkup)
x.Vector(e.Rows)
return x.buf
}

func (e PredMessagesBotCallbackAnswer) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesBotCallbackAnswer)
x.Int(e.Flags)
x.String(e.Message)
x.String(e.Url)
x.Int(e.Cache_time)
return x.buf
}

func (e PredUpdateBotCallbackQuery) encode() []byte {
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

func (e PredMessagesMessageEditData) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesMessageEditData)
x.Int(e.Flags)
return x.buf
}

func (e PredUpdateEditMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateEditMessage)
x.Bytes(e.Message.encode())
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredInputBotInlineMessageMediaGeo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineMessageMediaGeo)
x.Int(e.Flags)
x.Bytes(e.Geo_point.encode())
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredInputBotInlineMessageMediaVenue) encode() []byte {
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

func (e PredInputBotInlineMessageMediaContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineMessageMediaContact)
x.Int(e.Flags)
x.String(e.Phone_number)
x.String(e.First_name)
x.String(e.Last_name)
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredBotInlineMessageMediaGeo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_botInlineMessageMediaGeo)
x.Int(e.Flags)
x.Bytes(e.Geo.encode())
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredBotInlineMessageMediaVenue) encode() []byte {
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

func (e PredBotInlineMessageMediaContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_botInlineMessageMediaContact)
x.Int(e.Flags)
x.String(e.Phone_number)
x.String(e.First_name)
x.String(e.Last_name)
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredInputBotInlineResultPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineResultPhoto)
x.String(e.Id)
x.String(e.Type)
x.Bytes(e.Photo.encode())
x.Bytes(e.Send_message.encode())
return x.buf
}

func (e PredInputBotInlineResultDocument) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineResultDocument)
x.Int(e.Flags)
x.String(e.Id)
x.String(e.Type)
x.String(e.Title)
x.String(e.Description)
x.Bytes(e.Document.encode())
x.Bytes(e.Send_message.encode())
return x.buf
}

func (e PredBotInlineMediaResult) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_botInlineMediaResult)
x.Int(e.Flags)
x.String(e.Id)
x.String(e.Type)
x.Bytes(e.Photo.encode())
x.Bytes(e.Document.encode())
x.String(e.Title)
x.String(e.Description)
x.Bytes(e.Send_message.encode())
return x.buf
}

func (e PredInputBotInlineMessageID) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineMessageID)
x.Int(e.Dc_id)
x.Long(e.Id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredUpdateInlineBotCallbackQuery) encode() []byte {
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

func (e PredInlineBotSwitchPM) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inlineBotSwitchPM)
x.String(e.Text)
x.String(e.Start_param)
return x.buf
}

func (e PredMessageEntityMentionName) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageEntityMentionName)
x.Int(e.Offset)
x.Int(e.Length)
x.Int(e.User_id)
return x.buf
}

func (e PredInputMessageEntityMentionName) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessageEntityMentionName)
x.Int(e.Offset)
x.Int(e.Length)
x.Bytes(e.User_id.encode())
return x.buf
}

func (e PredMessagesPeerDialogs) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesPeerDialogs)
x.Vector(e.Dialogs)
x.Vector(e.Messages)
x.Vector(e.Chats)
x.Vector(e.Users)
x.Bytes(e.State.encode())
return x.buf
}

func (e PredTopPeer) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_topPeer)
x.Bytes(e.Peer.encode())
x.Double(e.Rating)
return x.buf
}

func (e PredTopPeerCategoryBotsPM) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_topPeerCategoryBotsPM)
return x.buf
}

func (e PredTopPeerCategoryBotsInline) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_topPeerCategoryBotsInline)
return x.buf
}

func (e PredTopPeerCategoryCorrespondents) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_topPeerCategoryCorrespondents)
return x.buf
}

func (e PredTopPeerCategoryGroups) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_topPeerCategoryGroups)
return x.buf
}

func (e PredTopPeerCategoryChannels) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_topPeerCategoryChannels)
return x.buf
}

func (e PredTopPeerCategoryPeers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_topPeerCategoryPeers)
x.Bytes(e.Category.encode())
x.Int(e.Count)
x.Vector(e.Peers)
return x.buf
}

func (e PredContactsTopPeersNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsTopPeersNotModified)
return x.buf
}

func (e PredContactsTopPeers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_contactsTopPeers)
x.Vector(e.Categories)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredInputMessagesFilterChatPhotos) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterChatPhotos)
return x.buf
}

func (e PredUpdateReadChannelOutbox) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateReadChannelOutbox)
x.Int(e.Channel_id)
x.Int(e.Max_id)
return x.buf
}

func (e PredUpdateDraftMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateDraftMessage)
x.Bytes(e.Peer.encode())
x.Bytes(e.Draft.encode())
return x.buf
}

func (e PredDraftMessageEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_draftMessageEmpty)
return x.buf
}

func (e PredDraftMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_draftMessage)
x.Int(e.Flags)
x.Int(e.Reply_to_msg_id)
x.String(e.Message)
x.Vector(e.Entities)
x.Int(e.Date)
return x.buf
}

func (e PredMessageActionHistoryClear) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionHistoryClear)
return x.buf
}

func (e PredUpdateReadFeaturedStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateReadFeaturedStickers)
return x.buf
}

func (e PredUpdateRecentStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateRecentStickers)
return x.buf
}

func (e PredMessagesFeaturedStickersNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesFeaturedStickersNotModified)
return x.buf
}

func (e PredMessagesFeaturedStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesFeaturedStickers)
x.Int(e.Hash)
x.Vector(e.Sets)
x.VectorLong(e.Unread)
return x.buf
}

func (e PredMessagesRecentStickersNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesRecentStickersNotModified)
return x.buf
}

func (e PredMessagesRecentStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesRecentStickers)
x.Int(e.Hash)
x.Vector(e.Stickers)
return x.buf
}

func (e PredMessagesArchivedStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesArchivedStickers)
x.Int(e.Count)
x.Vector(e.Sets)
return x.buf
}

func (e PredMessagesStickerSetInstallResultSuccess) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesStickerSetInstallResultSuccess)
return x.buf
}

func (e PredMessagesStickerSetInstallResultArchive) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesStickerSetInstallResultArchive)
x.Vector(e.Sets)
return x.buf
}

func (e PredStickerSetCovered) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_stickerSetCovered)
x.Bytes(e.Set.encode())
x.Bytes(e.Cover.encode())
return x.buf
}

func (e PredInputMediaPhotoExternal) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaPhotoExternal)
x.Int(e.Flags)
x.String(e.Url)
x.String(e.Caption)
x.Int(e.Ttl_seconds)
return x.buf
}

func (e PredInputMediaDocumentExternal) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaDocumentExternal)
x.Int(e.Flags)
x.String(e.Url)
x.String(e.Caption)
x.Int(e.Ttl_seconds)
return x.buf
}

func (e PredUpdateConfig) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateConfig)
return x.buf
}

func (e PredUpdatePtsChanged) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatePtsChanged)
return x.buf
}

func (e PredMessageActionGameScore) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionGameScore)
x.Long(e.Game_id)
x.Int(e.Score)
return x.buf
}

func (e PredDocumentAttributeHasStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_documentAttributeHasStickers)
return x.buf
}

func (e PredKeyboardButtonGame) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButtonGame)
x.String(e.Text)
return x.buf
}

func (e PredStickerSetMultiCovered) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_stickerSetMultiCovered)
x.Bytes(e.Set.encode())
x.Vector(e.Covers)
return x.buf
}

func (e PredMaskCoords) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_maskCoords)
x.Int(e.N)
x.Double(e.X)
x.Double(e.Y)
x.Double(e.Zoom)
return x.buf
}

func (e PredInputStickeredMediaPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputStickeredMediaPhoto)
x.Bytes(e.Id.encode())
return x.buf
}

func (e PredInputStickeredMediaDocument) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputStickeredMediaDocument)
x.Bytes(e.Id.encode())
return x.buf
}

func (e PredInputMediaGame) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMediaGame)
x.Bytes(e.Id.encode())
return x.buf
}

func (e PredMessageMediaGame) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageMediaGame)
x.Bytes(e.Game.encode())
return x.buf
}

func (e PredInputBotInlineMessageGame) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineMessageGame)
x.Int(e.Flags)
x.Bytes(e.Reply_markup.encode())
return x.buf
}

func (e PredInputBotInlineResultGame) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputBotInlineResultGame)
x.String(e.Id)
x.String(e.Short_name)
x.Bytes(e.Send_message.encode())
return x.buf
}

func (e PredGame) encode() []byte {
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

func (e PredInputGameID) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputGameID)
x.Long(e.Id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredInputGameShortName) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputGameShortName)
x.Bytes(e.Bot_id.encode())
x.String(e.Short_name)
return x.buf
}

func (e PredHighScore) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_highScore)
x.Int(e.Pos)
x.Int(e.User_id)
x.Int(e.Score)
return x.buf
}

func (e PredMessagesHighScores) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesHighScores)
x.Vector(e.Scores)
x.Vector(e.Users)
return x.buf
}

func (e PredMessagesChatsSlice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesChatsSlice)
x.Int(e.Count)
x.Vector(e.Chats)
return x.buf
}

func (e PredUpdateChannelWebPage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChannelWebPage)
x.Int(e.Channel_id)
x.Bytes(e.Webpage.encode())
x.Int(e.Pts)
x.Int(e.Pts_count)
return x.buf
}

func (e PredUpdatesDifferenceTooLong) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatesDifferenceTooLong)
x.Int(e.Pts)
return x.buf
}

func (e PredSendMessageGamePlayAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageGamePlayAction)
return x.buf
}

func (e PredWebPageNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_webPageNotModified)
return x.buf
}

func (e PredTextEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textEmpty)
return x.buf
}

func (e PredTextPlain) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textPlain)
x.String(e.Text)
return x.buf
}

func (e PredTextBold) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textBold)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredTextItalic) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textItalic)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredTextUnderline) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textUnderline)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredTextStrike) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textStrike)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredTextFixed) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textFixed)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredTextUrl) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textUrl)
x.Bytes(e.Text.encode())
x.String(e.Url)
x.Long(e.Webpage_id)
return x.buf
}

func (e PredTextEmail) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textEmail)
x.Bytes(e.Text.encode())
x.String(e.Email)
return x.buf
}

func (e PredTextConcat) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_textConcat)
x.Vector(e.Texts)
return x.buf
}

func (e PredPageBlockTitle) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockTitle)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredPageBlockSubtitle) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockSubtitle)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredPageBlockAuthorDate) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockAuthorDate)
x.Bytes(e.Author.encode())
x.Int(e.Published_date)
return x.buf
}

func (e PredPageBlockHeader) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockHeader)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredPageBlockSubheader) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockSubheader)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredPageBlockParagraph) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockParagraph)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredPageBlockPreformatted) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockPreformatted)
x.Bytes(e.Text.encode())
x.String(e.Language)
return x.buf
}

func (e PredPageBlockFooter) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockFooter)
x.Bytes(e.Text.encode())
return x.buf
}

func (e PredPageBlockDivider) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockDivider)
return x.buf
}

func (e PredPageBlockList) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockList)
x.Bytes(e.Ordered.encode())
x.Vector(e.Items)
return x.buf
}

func (e PredPageBlockBlockquote) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockBlockquote)
x.Bytes(e.Text.encode())
x.Bytes(e.Caption.encode())
return x.buf
}

func (e PredPageBlockPullquote) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockPullquote)
x.Bytes(e.Text.encode())
x.Bytes(e.Caption.encode())
return x.buf
}

func (e PredPageBlockPhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockPhoto)
x.Long(e.Photo_id)
x.Bytes(e.Caption.encode())
return x.buf
}

func (e PredPageBlockVideo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockVideo)
x.Int(e.Flags)
x.Long(e.Video_id)
x.Bytes(e.Caption.encode())
return x.buf
}

func (e PredPageBlockCover) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockCover)
x.Bytes(e.Cover.encode())
return x.buf
}

func (e PredPageBlockEmbed) encode() []byte {
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

func (e PredPageBlockEmbedPost) encode() []byte {
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

func (e PredPageBlockSlideshow) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockSlideshow)
x.Vector(e.Items)
x.Bytes(e.Caption.encode())
return x.buf
}

func (e PredPagePart) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pagePart)
x.Vector(e.Blocks)
x.Vector(e.Photos)
x.Vector(e.Documents)
return x.buf
}

func (e PredPageFull) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageFull)
x.Vector(e.Blocks)
x.Vector(e.Photos)
x.Vector(e.Documents)
return x.buf
}

func (e PredUpdatePhoneCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatePhoneCall)
x.Bytes(e.Phone_call.encode())
return x.buf
}

func (e PredUpdateDialogPinned) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateDialogPinned)
x.Int(e.Flags)
x.Bytes(e.Peer.encode())
return x.buf
}

func (e PredUpdatePinnedDialogs) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updatePinnedDialogs)
x.Int(e.Flags)
x.Vector(e.Order)
return x.buf
}

func (e PredInputPrivacyKeyPhoneCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPrivacyKeyPhoneCall)
return x.buf
}

func (e PredPrivacyKeyPhoneCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_privacyKeyPhoneCall)
return x.buf
}

func (e PredPageBlockUnsupported) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockUnsupported)
return x.buf
}

func (e PredPageBlockAnchor) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockAnchor)
x.String(e.Name)
return x.buf
}

func (e PredPageBlockCollage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockCollage)
x.Vector(e.Items)
x.Bytes(e.Caption.encode())
return x.buf
}

func (e PredInputPhoneCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPhoneCall)
x.Long(e.Id)
x.Long(e.Access_hash)
return x.buf
}

func (e PredPhoneCallEmpty) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phoneCallEmpty)
x.Long(e.Id)
return x.buf
}

func (e PredPhoneCallWaiting) encode() []byte {
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

func (e PredPhoneCallRequested) encode() []byte {
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

func (e PredPhoneCall) encode() []byte {
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

func (e PredPhoneCallDiscarded) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phoneCallDiscarded)
x.Int(e.Flags)
x.Long(e.Id)
x.Bytes(e.Reason.encode())
x.Int(e.Duration)
return x.buf
}

func (e PredPhoneConnection) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phoneConnection)
x.Long(e.Id)
x.String(e.Ip)
x.String(e.Ipv6)
x.Int(e.Port)
x.StringBytes(e.Peer_tag)
return x.buf
}

func (e PredPhoneCallProtocol) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phoneCallProtocol)
x.Int(e.Flags)
x.Int(e.Min_layer)
x.Int(e.Max_layer)
return x.buf
}

func (e PredPhonePhoneCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phonePhoneCall)
x.Bytes(e.Phone_call.encode())
x.Vector(e.Users)
return x.buf
}

func (e PredPhoneCallDiscardReasonMissed) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phoneCallDiscardReasonMissed)
return x.buf
}

func (e PredPhoneCallDiscardReasonDisconnect) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phoneCallDiscardReasonDisconnect)
return x.buf
}

func (e PredPhoneCallDiscardReasonHangup) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phoneCallDiscardReasonHangup)
return x.buf
}

func (e PredPhoneCallDiscardReasonBusy) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_phoneCallDiscardReasonBusy)
return x.buf
}

func (e PredInputMessagesFilterPhoneCalls) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterPhoneCalls)
x.Int(e.Flags)
return x.buf
}

func (e PredMessageActionPhoneCall) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionPhoneCall)
x.Int(e.Flags)
x.Long(e.Call_id)
x.Bytes(e.Reason.encode())
x.Int(e.Duration)
return x.buf
}

func (e PredInvoice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_invoice)
x.Int(e.Flags)
x.String(e.Currency)
x.Vector(e.Prices)
return x.buf
}

func (e PredInputMediaInvoice) encode() []byte {
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

func (e PredMessageActionPaymentSentMe) encode() []byte {
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

func (e PredMessageMediaInvoice) encode() []byte {
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

func (e PredKeyboardButtonBuy) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_keyboardButtonBuy)
x.String(e.Text)
return x.buf
}

func (e PredMessageActionPaymentSent) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionPaymentSent)
x.String(e.Currency)
x.Long(e.Total_amount)
return x.buf
}

func (e PredPaymentsPaymentForm) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentsPaymentForm)
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

func (e PredPostAddress) encode() []byte {
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

func (e PredPaymentRequestedInfo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentRequestedInfo)
x.Int(e.Flags)
x.String(e.Name)
x.String(e.Phone)
x.String(e.Email)
x.Bytes(e.Shipping_address.encode())
return x.buf
}

func (e PredUpdateBotWebhookJSON) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateBotWebhookJSON)
x.Bytes(e.Data.encode())
return x.buf
}

func (e PredUpdateBotWebhookJSONQuery) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateBotWebhookJSONQuery)
x.Long(e.Query_id)
x.Bytes(e.Data.encode())
x.Int(e.Timeout)
return x.buf
}

func (e PredUpdateBotShippingQuery) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateBotShippingQuery)
x.Long(e.Query_id)
x.Int(e.User_id)
x.StringBytes(e.Payload)
x.Bytes(e.Shipping_address.encode())
return x.buf
}

func (e PredUpdateBotPrecheckoutQuery) encode() []byte {
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

func (e PredDataJSON) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_dataJSON)
x.String(e.Data)
return x.buf
}

func (e PredLabeledPrice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_labeledPrice)
x.String(e.Label)
x.Long(e.Amount)
return x.buf
}

func (e PredPaymentCharge) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentCharge)
x.String(e.Id)
x.String(e.Provider_charge_id)
return x.buf
}

func (e PredPaymentSavedCredentialsCard) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentSavedCredentialsCard)
x.String(e.Id)
x.String(e.Title)
return x.buf
}

func (e PredWebDocument) encode() []byte {
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

func (e PredInputWebDocument) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputWebDocument)
x.String(e.Url)
x.Int(e.Size)
x.String(e.Mime_type)
x.Vector(e.Attributes)
return x.buf
}

func (e PredInputWebFileLocation) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputWebFileLocation)
x.String(e.Url)
x.Long(e.Access_hash)
return x.buf
}

func (e PredUploadWebFile) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_uploadWebFile)
x.Int(e.Size)
x.String(e.Mime_type)
x.Bytes(e.File_type.encode())
x.Int(e.Mtime)
x.StringBytes(e.Bytes)
return x.buf
}

func (e PredPaymentsValidatedRequestedInfo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentsValidatedRequestedInfo)
x.Int(e.Flags)
x.String(e.Id)
x.Vector(e.Shipping_options)
return x.buf
}

func (e PredPaymentsPaymentResult) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentsPaymentResult)
x.Bytes(e.Updates.encode())
return x.buf
}

func (e PredPaymentsPaymentVerficationNeeded) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentsPaymentVerficationNeeded)
x.String(e.Url)
return x.buf
}

func (e PredPaymentsPaymentReceipt) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentsPaymentReceipt)
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

func (e PredPaymentsSavedInfo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_paymentsSavedInfo)
x.Int(e.Flags)
x.Bytes(e.Saved_info.encode())
return x.buf
}

func (e PredInputPaymentCredentialsSaved) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPaymentCredentialsSaved)
x.String(e.Id)
x.StringBytes(e.Tmp_password)
return x.buf
}

func (e PredInputPaymentCredentials) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputPaymentCredentials)
x.Int(e.Flags)
x.Bytes(e.Data.encode())
return x.buf
}

func (e PredAccountTmpPassword) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_accountTmpPassword)
x.StringBytes(e.Tmp_password)
x.Int(e.Valid_until)
return x.buf
}

func (e PredShippingOption) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_shippingOption)
x.String(e.Id)
x.String(e.Title)
x.Vector(e.Prices)
return x.buf
}

func (e PredPhoneCallAccepted) encode() []byte {
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

func (e PredInputMessagesFilterRoundVoice) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterRoundVoice)
return x.buf
}

func (e PredInputMessagesFilterRoundVideo) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterRoundVideo)
return x.buf
}

func (e PredUploadFileCdnRedirect) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_uploadFileCdnRedirect)
x.Int(e.Dc_id)
x.StringBytes(e.File_token)
x.StringBytes(e.Encryption_key)
x.StringBytes(e.Encryption_iv)
x.Vector(e.Cdn_file_hashes)
return x.buf
}

func (e PredSendMessageRecordRoundAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageRecordRoundAction)
return x.buf
}

func (e PredSendMessageUploadRoundAction) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_sendMessageUploadRoundAction)
x.Int(e.Progress)
return x.buf
}

func (e PredUploadCdnFileReuploadNeeded) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_uploadCdnFileReuploadNeeded)
x.StringBytes(e.Request_token)
return x.buf
}

func (e PredUploadCdnFile) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_uploadCdnFile)
x.StringBytes(e.Bytes)
return x.buf
}

func (e PredCdnPublicKey) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_cdnPublicKey)
x.Int(e.Dc_id)
x.String(e.Public_key)
return x.buf
}

func (e PredCdnConfig) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_cdnConfig)
x.Vector(e.Public_keys)
return x.buf
}

func (e PredUpdateLangPackTooLong) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateLangPackTooLong)
return x.buf
}

func (e PredUpdateLangPack) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateLangPack)
x.Bytes(e.Difference.encode())
return x.buf
}

func (e PredPageBlockChannel) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockChannel)
x.Bytes(e.Channel.encode())
return x.buf
}

func (e PredInputStickerSetItem) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputStickerSetItem)
x.Int(e.Flags)
x.Bytes(e.Document.encode())
x.String(e.Emoji)
x.Bytes(e.Mask_coords.encode())
return x.buf
}

func (e PredLangPackString) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_langPackString)
x.String(e.Key)
x.String(e.Value)
return x.buf
}

func (e PredLangPackStringPluralized) encode() []byte {
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

func (e PredLangPackStringDeleted) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_langPackStringDeleted)
x.String(e.Key)
return x.buf
}

func (e PredLangPackDifference) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_langPackDifference)
x.String(e.Lang_code)
x.Int(e.From_version)
x.Int(e.Version)
x.Vector(e.Strings)
return x.buf
}

func (e PredLangPackLanguage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_langPackLanguage)
x.String(e.Name)
x.String(e.Native_name)
x.String(e.Lang_code)
return x.buf
}

func (e PredChannelParticipantAdmin) encode() []byte {
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

func (e PredChannelParticipantBanned) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantBanned)
x.Int(e.Flags)
x.Int(e.User_id)
x.Int(e.Kicked_by)
x.Int(e.Date)
x.Bytes(e.Banned_rights.encode())
return x.buf
}

func (e PredChannelParticipantsBanned) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantsBanned)
x.String(e.Q)
return x.buf
}

func (e PredChannelParticipantsSearch) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelParticipantsSearch)
x.String(e.Q)
return x.buf
}

func (e PredTopPeerCategoryPhoneCalls) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_topPeerCategoryPhoneCalls)
return x.buf
}

func (e PredPageBlockAudio) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_pageBlockAudio)
x.Long(e.Audio_id)
x.Bytes(e.Caption.encode())
return x.buf
}

func (e PredChannelAdminRights) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminRights)
x.Int(e.Flags)
return x.buf
}

func (e PredChannelBannedRights) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelBannedRights)
x.Int(e.Flags)
x.Int(e.Until_date)
return x.buf
}

func (e PredChannelAdminLogEventActionChangeTitle) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionChangeTitle)
x.String(e.Prev_value)
x.String(e.New_value)
return x.buf
}

func (e PredChannelAdminLogEventActionChangeAbout) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionChangeAbout)
x.String(e.Prev_value)
x.String(e.New_value)
return x.buf
}

func (e PredChannelAdminLogEventActionChangeUsername) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionChangeUsername)
x.String(e.Prev_value)
x.String(e.New_value)
return x.buf
}

func (e PredChannelAdminLogEventActionChangePhoto) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionChangePhoto)
x.Bytes(e.Prev_photo.encode())
x.Bytes(e.New_photo.encode())
return x.buf
}

func (e PredChannelAdminLogEventActionToggleInvites) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionToggleInvites)
x.Bytes(e.New_value.encode())
return x.buf
}

func (e PredChannelAdminLogEventActionToggleSignatures) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionToggleSignatures)
x.Bytes(e.New_value.encode())
return x.buf
}

func (e PredChannelAdminLogEventActionUpdatePinned) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionUpdatePinned)
x.Bytes(e.Message.encode())
return x.buf
}

func (e PredChannelAdminLogEventActionEditMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionEditMessage)
x.Bytes(e.Prev_message.encode())
x.Bytes(e.New_message.encode())
return x.buf
}

func (e PredChannelAdminLogEventActionDeleteMessage) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionDeleteMessage)
x.Bytes(e.Message.encode())
return x.buf
}

func (e PredChannelAdminLogEventActionParticipantJoin) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionParticipantJoin)
return x.buf
}

func (e PredChannelAdminLogEventActionParticipantLeave) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionParticipantLeave)
return x.buf
}

func (e PredChannelAdminLogEventActionParticipantInvite) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionParticipantInvite)
x.Bytes(e.Participant.encode())
return x.buf
}

func (e PredChannelAdminLogEventActionParticipantToggleBan) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionParticipantToggleBan)
x.Bytes(e.Prev_participant.encode())
x.Bytes(e.New_participant.encode())
return x.buf
}

func (e PredChannelAdminLogEventActionParticipantToggleAdmin) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionParticipantToggleAdmin)
x.Bytes(e.Prev_participant.encode())
x.Bytes(e.New_participant.encode())
return x.buf
}

func (e PredChannelAdminLogEvent) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEvent)
x.Long(e.Id)
x.Int(e.Date)
x.Int(e.User_id)
x.Bytes(e.Action.encode())
return x.buf
}

func (e PredChannelsAdminLogResults) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelsAdminLogResults)
x.Vector(e.Events)
x.Vector(e.Chats)
x.Vector(e.Users)
return x.buf
}

func (e PredChannelAdminLogEventsFilter) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventsFilter)
x.Int(e.Flags)
return x.buf
}

func (e PredMessageActionScreenshotTaken) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messageActionScreenshotTaken)
return x.buf
}

func (e PredPopularContact) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_popularContact)
x.Long(e.Client_id)
x.Int(e.Importers)
return x.buf
}

func (e PredCdnFileHash) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_cdnFileHash)
x.Int(e.Offset)
x.Int(e.Limit)
x.StringBytes(e.Hash)
return x.buf
}

func (e PredInputMessagesFilterMyMentions) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterMyMentions)
return x.buf
}

func (e PredInputMessagesFilterMyMentionsUnread) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_inputMessagesFilterMyMentionsUnread)
return x.buf
}

func (e PredUpdateContactsReset) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateContactsReset)
return x.buf
}

func (e PredChannelAdminLogEventActionChangeStickerSet) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_channelAdminLogEventActionChangeStickerSet)
x.Bytes(e.Prev_stickerset.encode())
x.Bytes(e.New_stickerset.encode())
return x.buf
}

func (e PredUpdateFavedStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateFavedStickers)
return x.buf
}

func (e PredMessagesFavedStickers) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesFavedStickers)
x.Int(e.Hash)
x.Vector(e.Packs)
x.Vector(e.Stickers)
return x.buf
}

func (e PredMessagesFavedStickersNotModified) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_messagesFavedStickersNotModified)
return x.buf
}

func (e PredUpdateChannelReadMessagesContents) encode() []byte {
x := NewEncodeBuf(512)
x.UInt(crc_updateChannelReadMessagesContents)
x.Int(e.Channel_id)
x.VectorInt(e.Messages)
return x.buf
}

