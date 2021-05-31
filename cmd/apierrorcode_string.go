// Code generated by "stringer -type=APIErrorCode -trimprefix=Err api-errors.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrNone-0]
	_ = x[ErrAccessDenied-1]
	_ = x[ErrBadDigest-2]
	_ = x[ErrEntityTooSmall-3]
	_ = x[ErrEntityTooLarge-4]
	_ = x[ErrPolicyTooLarge-5]
	_ = x[ErrIncompleteBody-6]
	_ = x[ErrInternalError-7]
	_ = x[ErrInvalidAccessKeyID-8]
	_ = x[ErrInvalidBucketName-9]
	_ = x[ErrInvalidDigest-10]
	_ = x[ErrInvalidRange-11]
	_ = x[ErrInvalidRangePartNumber-12]
	_ = x[ErrInvalidCopyPartRange-13]
	_ = x[ErrInvalidCopyPartRangeSource-14]
	_ = x[ErrInvalidMaxKeys-15]
	_ = x[ErrInvalidEncodingMethod-16]
	_ = x[ErrInvalidMaxUploads-17]
	_ = x[ErrInvalidMaxParts-18]
	_ = x[ErrInvalidPartNumberMarker-19]
	_ = x[ErrInvalidPartNumber-20]
	_ = x[ErrInvalidRequestBody-21]
	_ = x[ErrInvalidCopySource-22]
	_ = x[ErrInvalidMetadataDirective-23]
	_ = x[ErrInvalidCopyDest-24]
	_ = x[ErrInvalidPolicyDocument-25]
	_ = x[ErrInvalidObjectState-26]
	_ = x[ErrMalformedXML-27]
	_ = x[ErrMissingContentLength-28]
	_ = x[ErrMissingContentMD5-29]
	_ = x[ErrMissingRequestBodyError-30]
	_ = x[ErrMissingSecurityHeader-31]
	_ = x[ErrNoSuchBucket-32]
	_ = x[ErrNoSuchBucketPolicy-33]
	_ = x[ErrNoSuchBucketLifecycle-34]
	_ = x[ErrNoSuchLifecycleConfiguration-35]
	_ = x[ErrNoSuchBucketSSEConfig-36]
	_ = x[ErrNoSuchCORSConfiguration-37]
	_ = x[ErrNoSuchWebsiteConfiguration-38]
	_ = x[ErrReplicationConfigurationNotFoundError-39]
	_ = x[ErrRemoteDestinationNotFoundError-40]
	_ = x[ErrReplicationDestinationMissingLock-41]
	_ = x[ErrRemoteTargetNotFoundError-42]
	_ = x[ErrReplicationRemoteConnectionError-43]
	_ = x[ErrBucketRemoteIdenticalToSource-44]
	_ = x[ErrBucketRemoteAlreadyExists-45]
	_ = x[ErrBucketRemoteLabelInUse-46]
	_ = x[ErrBucketRemoteArnTypeInvalid-47]
	_ = x[ErrBucketRemoteArnInvalid-48]
	_ = x[ErrBucketRemoteRemoveDisallowed-49]
	_ = x[ErrRemoteTargetNotVersionedError-50]
	_ = x[ErrReplicationSourceNotVersionedError-51]
	_ = x[ErrReplicationNeedsVersioningError-52]
	_ = x[ErrReplicationBucketNeedsVersioningError-53]
	_ = x[ErrObjectRestoreAlreadyInProgress-54]
	_ = x[ErrNoSuchKey-55]
	_ = x[ErrNoSuchUpload-56]
	_ = x[ErrInvalidVersionID-57]
	_ = x[ErrNoSuchVersion-58]
	_ = x[ErrNotImplemented-59]
	_ = x[ErrPreconditionFailed-60]
	_ = x[ErrRequestTimeTooSkewed-61]
	_ = x[ErrSignatureDoesNotMatch-62]
	_ = x[ErrMethodNotAllowed-63]
	_ = x[ErrInvalidPart-64]
	_ = x[ErrInvalidPartOrder-65]
	_ = x[ErrAuthorizationHeaderMalformed-66]
	_ = x[ErrMalformedPOSTRequest-67]
	_ = x[ErrPOSTFileRequired-68]
	_ = x[ErrSignatureVersionNotSupported-69]
	_ = x[ErrBucketNotEmpty-70]
	_ = x[ErrAllAccessDisabled-71]
	_ = x[ErrMalformedPolicy-72]
	_ = x[ErrMissingFields-73]
	_ = x[ErrMissingCredTag-74]
	_ = x[ErrCredMalformed-75]
	_ = x[ErrInvalidRegion-76]
	_ = x[ErrInvalidServiceS3-77]
	_ = x[ErrInvalidServiceSTS-78]
	_ = x[ErrInvalidRequestVersion-79]
	_ = x[ErrMissingSignTag-80]
	_ = x[ErrMissingSignHeadersTag-81]
	_ = x[ErrMalformedDate-82]
	_ = x[ErrMalformedPresignedDate-83]
	_ = x[ErrMalformedCredentialDate-84]
	_ = x[ErrMalformedCredentialRegion-85]
	_ = x[ErrMalformedExpires-86]
	_ = x[ErrNegativeExpires-87]
	_ = x[ErrAuthHeaderEmpty-88]
	_ = x[ErrExpiredPresignRequest-89]
	_ = x[ErrRequestNotReadyYet-90]
	_ = x[ErrUnsignedHeaders-91]
	_ = x[ErrMissingDateHeader-92]
	_ = x[ErrInvalidQuerySignatureAlgo-93]
	_ = x[ErrInvalidQueryParams-94]
	_ = x[ErrBucketAlreadyOwnedByYou-95]
	_ = x[ErrInvalidDuration-96]
	_ = x[ErrBucketAlreadyExists-97]
	_ = x[ErrMetadataTooLarge-98]
	_ = x[ErrUnsupportedMetadata-99]
	_ = x[ErrMaximumExpires-100]
	_ = x[ErrSlowDown-101]
	_ = x[ErrInvalidPrefixMarker-102]
	_ = x[ErrBadRequest-103]
	_ = x[ErrKeyTooLongError-104]
	_ = x[ErrInvalidBucketObjectLockConfiguration-105]
	_ = x[ErrObjectLockConfigurationNotFound-106]
	_ = x[ErrObjectLockConfigurationNotAllowed-107]
	_ = x[ErrNoSuchObjectLockConfiguration-108]
	_ = x[ErrObjectLocked-109]
	_ = x[ErrInvalidRetentionDate-110]
	_ = x[ErrPastObjectLockRetainDate-111]
	_ = x[ErrUnknownWORMModeDirective-112]
	_ = x[ErrBucketTaggingNotFound-113]
	_ = x[ErrObjectLockInvalidHeaders-114]
	_ = x[ErrInvalidTagDirective-115]
	_ = x[ErrInvalidEncryptionMethod-116]
	_ = x[ErrInsecureSSECustomerRequest-117]
	_ = x[ErrSSEMultipartEncrypted-118]
	_ = x[ErrSSEEncryptedObject-119]
	_ = x[ErrInvalidEncryptionParameters-120]
	_ = x[ErrInvalidSSECustomerAlgorithm-121]
	_ = x[ErrInvalidSSECustomerKey-122]
	_ = x[ErrMissingSSECustomerKey-123]
	_ = x[ErrMissingSSECustomerKeyMD5-124]
	_ = x[ErrSSECustomerKeyMD5Mismatch-125]
	_ = x[ErrInvalidSSECustomerParameters-126]
	_ = x[ErrIncompatibleEncryptionMethod-127]
	_ = x[ErrKMSNotConfigured-128]
	_ = x[ErrNoAccessKey-129]
	_ = x[ErrInvalidToken-130]
	_ = x[ErrEventNotification-131]
	_ = x[ErrARNNotification-132]
	_ = x[ErrRegionNotification-133]
	_ = x[ErrOverlappingFilterNotification-134]
	_ = x[ErrFilterNameInvalid-135]
	_ = x[ErrFilterNamePrefix-136]
	_ = x[ErrFilterNameSuffix-137]
	_ = x[ErrFilterValueInvalid-138]
	_ = x[ErrOverlappingConfigs-139]
	_ = x[ErrUnsupportedNotification-140]
	_ = x[ErrContentSHA256Mismatch-141]
	_ = x[ErrReadQuorum-142]
	_ = x[ErrWriteQuorum-143]
	_ = x[ErrParentIsObject-144]
	_ = x[ErrStorageFull-145]
	_ = x[ErrRequestBodyParse-146]
	_ = x[ErrObjectExistsAsDirectory-147]
	_ = x[ErrInvalidObjectName-148]
	_ = x[ErrInvalidObjectNamePrefixSlash-149]
	_ = x[ErrInvalidResourceName-150]
	_ = x[ErrServerNotInitialized-151]
	_ = x[ErrOperationTimedOut-152]
	_ = x[ErrClientDisconnected-153]
	_ = x[ErrOperationMaxedOut-154]
	_ = x[ErrInvalidRequest-155]
	_ = x[ErrTransitionStorageClassNotFoundError-156]
	_ = x[ErrInvalidStorageClass-157]
	_ = x[ErrBackendDown-158]
	_ = x[ErrMalformedJSON-159]
	_ = x[ErrAdminNoSuchUser-160]
	_ = x[ErrAdminNoSuchGroup-161]
	_ = x[ErrAdminGroupNotEmpty-162]
	_ = x[ErrAdminNoSuchPolicy-163]
	_ = x[ErrAdminInvalidArgument-164]
	_ = x[ErrAdminInvalidAccessKey-165]
	_ = x[ErrAdminInvalidSecretKey-166]
	_ = x[ErrAdminConfigNoQuorum-167]
	_ = x[ErrAdminConfigTooLarge-168]
	_ = x[ErrAdminConfigBadJSON-169]
	_ = x[ErrAdminConfigDuplicateKeys-170]
	_ = x[ErrAdminCredentialsMismatch-171]
	_ = x[ErrInsecureClientRequest-172]
	_ = x[ErrObjectTampered-173]
	_ = x[ErrAdminBucketQuotaExceeded-174]
	_ = x[ErrAdminNoSuchQuotaConfiguration-175]
	_ = x[ErrHealNotImplemented-176]
	_ = x[ErrHealNoSuchProcess-177]
	_ = x[ErrHealInvalidClientToken-178]
	_ = x[ErrHealMissingBucket-179]
	_ = x[ErrHealAlreadyRunning-180]
	_ = x[ErrHealOverlappingPaths-181]
	_ = x[ErrIncorrectContinuationToken-182]
	_ = x[ErrEmptyRequestBody-183]
	_ = x[ErrUnsupportedFunction-184]
	_ = x[ErrInvalidExpressionType-185]
	_ = x[ErrBusy-186]
	_ = x[ErrUnauthorizedAccess-187]
	_ = x[ErrExpressionTooLong-188]
	_ = x[ErrIllegalSQLFunctionArgument-189]
	_ = x[ErrInvalidKeyPath-190]
	_ = x[ErrInvalidCompressionFormat-191]
	_ = x[ErrInvalidFileHeaderInfo-192]
	_ = x[ErrInvalidJSONType-193]
	_ = x[ErrInvalidQuoteFields-194]
	_ = x[ErrInvalidRequestParameter-195]
	_ = x[ErrInvalidDataType-196]
	_ = x[ErrInvalidTextEncoding-197]
	_ = x[ErrInvalidDataSource-198]
	_ = x[ErrInvalidTableAlias-199]
	_ = x[ErrMissingRequiredParameter-200]
	_ = x[ErrObjectSerializationConflict-201]
	_ = x[ErrUnsupportedSQLOperation-202]
	_ = x[ErrUnsupportedSQLStructure-203]
	_ = x[ErrUnsupportedSyntax-204]
	_ = x[ErrUnsupportedRangeHeader-205]
	_ = x[ErrLexerInvalidChar-206]
	_ = x[ErrLexerInvalidOperator-207]
	_ = x[ErrLexerInvalidLiteral-208]
	_ = x[ErrLexerInvalidIONLiteral-209]
	_ = x[ErrParseExpectedDatePart-210]
	_ = x[ErrParseExpectedKeyword-211]
	_ = x[ErrParseExpectedTokenType-212]
	_ = x[ErrParseExpected2TokenTypes-213]
	_ = x[ErrParseExpectedNumber-214]
	_ = x[ErrParseExpectedRightParenBuiltinFunctionCall-215]
	_ = x[ErrParseExpectedTypeName-216]
	_ = x[ErrParseExpectedWhenClause-217]
	_ = x[ErrParseUnsupportedToken-218]
	_ = x[ErrParseUnsupportedLiteralsGroupBy-219]
	_ = x[ErrParseExpectedMember-220]
	_ = x[ErrParseUnsupportedSelect-221]
	_ = x[ErrParseUnsupportedCase-222]
	_ = x[ErrParseUnsupportedCaseClause-223]
	_ = x[ErrParseUnsupportedAlias-224]
	_ = x[ErrParseUnsupportedSyntax-225]
	_ = x[ErrParseUnknownOperator-226]
	_ = x[ErrParseMissingIdentAfterAt-227]
	_ = x[ErrParseUnexpectedOperator-228]
	_ = x[ErrParseUnexpectedTerm-229]
	_ = x[ErrParseUnexpectedToken-230]
	_ = x[ErrParseUnexpectedKeyword-231]
	_ = x[ErrParseExpectedExpression-232]
	_ = x[ErrParseExpectedLeftParenAfterCast-233]
	_ = x[ErrParseExpectedLeftParenValueConstructor-234]
	_ = x[ErrParseExpectedLeftParenBuiltinFunctionCall-235]
	_ = x[ErrParseExpectedArgumentDelimiter-236]
	_ = x[ErrParseCastArity-237]
	_ = x[ErrParseInvalidTypeParam-238]
	_ = x[ErrParseEmptySelect-239]
	_ = x[ErrParseSelectMissingFrom-240]
	_ = x[ErrParseExpectedIdentForGroupName-241]
	_ = x[ErrParseExpectedIdentForAlias-242]
	_ = x[ErrParseUnsupportedCallWithStar-243]
	_ = x[ErrParseNonUnaryAgregateFunctionCall-244]
	_ = x[ErrParseMalformedJoin-245]
	_ = x[ErrParseExpectedIdentForAt-246]
	_ = x[ErrParseAsteriskIsNotAloneInSelectList-247]
	_ = x[ErrParseCannotMixSqbAndWildcardInSelectList-248]
	_ = x[ErrParseInvalidContextForWildcardInSelectList-249]
	_ = x[ErrIncorrectSQLFunctionArgumentType-250]
	_ = x[ErrValueParseFailure-251]
	_ = x[ErrEvaluatorInvalidArguments-252]
	_ = x[ErrIntegerOverflow-253]
	_ = x[ErrLikeInvalidInputs-254]
	_ = x[ErrCastFailed-255]
	_ = x[ErrInvalidCast-256]
	_ = x[ErrEvaluatorInvalidTimestampFormatPattern-257]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbolForParsing-258]
	_ = x[ErrEvaluatorTimestampFormatPatternDuplicateFields-259]
	_ = x[ErrEvaluatorTimestampFormatPatternHourClockAmPmMismatch-260]
	_ = x[ErrEvaluatorUnterminatedTimestampFormatPatternToken-261]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternToken-262]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbol-263]
	_ = x[ErrEvaluatorBindingDoesNotExist-264]
	_ = x[ErrMissingHeaders-265]
	_ = x[ErrInvalidColumnIndex-266]
	_ = x[ErrAdminConfigNotificationTargetsFailed-267]
	_ = x[ErrAdminProfilerNotEnabled-268]
	_ = x[ErrInvalidDecompressedSize-269]
	_ = x[ErrAddUserInvalidArgument-270]
	_ = x[ErrAdminAccountNotEligible-271]
	_ = x[ErrAccountNotEligible-272]
	_ = x[ErrAdminServiceAccountNotFound-273]
	_ = x[ErrPostPolicyConditionInvalidFormat-274]
}

const _APIErrorCode_name = "NoneAccessDeniedBadDigestEntityTooSmallEntityTooLargePolicyTooLargeIncompleteBodyInternalErrorInvalidAccessKeyIDInvalidBucketNameInvalidDigestInvalidRangeInvalidRangePartNumberInvalidCopyPartRangeInvalidCopyPartRangeSourceInvalidMaxKeysInvalidEncodingMethodInvalidMaxUploadsInvalidMaxPartsInvalidPartNumberMarkerInvalidPartNumberInvalidRequestBodyInvalidCopySourceInvalidMetadataDirectiveInvalidCopyDestInvalidPolicyDocumentInvalidObjectStateMalformedXMLMissingContentLengthMissingContentMD5MissingRequestBodyErrorMissingSecurityHeaderNoSuchBucketNoSuchBucketPolicyNoSuchBucketLifecycleNoSuchLifecycleConfigurationNoSuchBucketSSEConfigNoSuchCORSConfigurationNoSuchWebsiteConfigurationReplicationConfigurationNotFoundErrorRemoteDestinationNotFoundErrorReplicationDestinationMissingLockRemoteTargetNotFoundErrorReplicationRemoteConnectionErrorBucketRemoteIdenticalToSourceBucketRemoteAlreadyExistsBucketRemoteLabelInUseBucketRemoteArnTypeInvalidBucketRemoteArnInvalidBucketRemoteRemoveDisallowedRemoteTargetNotVersionedErrorReplicationSourceNotVersionedErrorReplicationNeedsVersioningErrorReplicationBucketNeedsVersioningErrorObjectRestoreAlreadyInProgressNoSuchKeyNoSuchUploadInvalidVersionIDNoSuchVersionNotImplementedPreconditionFailedRequestTimeTooSkewedSignatureDoesNotMatchMethodNotAllowedInvalidPartInvalidPartOrderAuthorizationHeaderMalformedMalformedPOSTRequestPOSTFileRequiredSignatureVersionNotSupportedBucketNotEmptyAllAccessDisabledMalformedPolicyMissingFieldsMissingCredTagCredMalformedInvalidRegionInvalidServiceS3InvalidServiceSTSInvalidRequestVersionMissingSignTagMissingSignHeadersTagMalformedDateMalformedPresignedDateMalformedCredentialDateMalformedCredentialRegionMalformedExpiresNegativeExpiresAuthHeaderEmptyExpiredPresignRequestRequestNotReadyYetUnsignedHeadersMissingDateHeaderInvalidQuerySignatureAlgoInvalidQueryParamsBucketAlreadyOwnedByYouInvalidDurationBucketAlreadyExistsMetadataTooLargeUnsupportedMetadataMaximumExpiresSlowDownInvalidPrefixMarkerBadRequestKeyTooLongErrorInvalidBucketObjectLockConfigurationObjectLockConfigurationNotFoundObjectLockConfigurationNotAllowedNoSuchObjectLockConfigurationObjectLockedInvalidRetentionDatePastObjectLockRetainDateUnknownWORMModeDirectiveBucketTaggingNotFoundObjectLockInvalidHeadersInvalidTagDirectiveInvalidEncryptionMethodInsecureSSECustomerRequestSSEMultipartEncryptedSSEEncryptedObjectInvalidEncryptionParametersInvalidSSECustomerAlgorithmInvalidSSECustomerKeyMissingSSECustomerKeyMissingSSECustomerKeyMD5SSECustomerKeyMD5MismatchInvalidSSECustomerParametersIncompatibleEncryptionMethodKMSNotConfiguredNoAccessKeyInvalidTokenEventNotificationARNNotificationRegionNotificationOverlappingFilterNotificationFilterNameInvalidFilterNamePrefixFilterNameSuffixFilterValueInvalidOverlappingConfigsUnsupportedNotificationContentSHA256MismatchReadQuorumWriteQuorumParentIsObjectStorageFullRequestBodyParseObjectExistsAsDirectoryInvalidObjectNameInvalidObjectNamePrefixSlashInvalidResourceNameServerNotInitializedOperationTimedOutClientDisconnectedOperationMaxedOutInvalidRequestTransitionStorageClassNotFoundErrorInvalidStorageClassBackendDownMalformedJSONAdminNoSuchUserAdminNoSuchGroupAdminGroupNotEmptyAdminNoSuchPolicyAdminInvalidArgumentAdminInvalidAccessKeyAdminInvalidSecretKeyAdminConfigNoQuorumAdminConfigTooLargeAdminConfigBadJSONAdminConfigDuplicateKeysAdminCredentialsMismatchInsecureClientRequestObjectTamperedAdminBucketQuotaExceededAdminNoSuchQuotaConfigurationHealNotImplementedHealNoSuchProcessHealInvalidClientTokenHealMissingBucketHealAlreadyRunningHealOverlappingPathsIncorrectContinuationTokenEmptyRequestBodyUnsupportedFunctionInvalidExpressionTypeBusyUnauthorizedAccessExpressionTooLongIllegalSQLFunctionArgumentInvalidKeyPathInvalidCompressionFormatInvalidFileHeaderInfoInvalidJSONTypeInvalidQuoteFieldsInvalidRequestParameterInvalidDataTypeInvalidTextEncodingInvalidDataSourceInvalidTableAliasMissingRequiredParameterObjectSerializationConflictUnsupportedSQLOperationUnsupportedSQLStructureUnsupportedSyntaxUnsupportedRangeHeaderLexerInvalidCharLexerInvalidOperatorLexerInvalidLiteralLexerInvalidIONLiteralParseExpectedDatePartParseExpectedKeywordParseExpectedTokenTypeParseExpected2TokenTypesParseExpectedNumberParseExpectedRightParenBuiltinFunctionCallParseExpectedTypeNameParseExpectedWhenClauseParseUnsupportedTokenParseUnsupportedLiteralsGroupByParseExpectedMemberParseUnsupportedSelectParseUnsupportedCaseParseUnsupportedCaseClauseParseUnsupportedAliasParseUnsupportedSyntaxParseUnknownOperatorParseMissingIdentAfterAtParseUnexpectedOperatorParseUnexpectedTermParseUnexpectedTokenParseUnexpectedKeywordParseExpectedExpressionParseExpectedLeftParenAfterCastParseExpectedLeftParenValueConstructorParseExpectedLeftParenBuiltinFunctionCallParseExpectedArgumentDelimiterParseCastArityParseInvalidTypeParamParseEmptySelectParseSelectMissingFromParseExpectedIdentForGroupNameParseExpectedIdentForAliasParseUnsupportedCallWithStarParseNonUnaryAgregateFunctionCallParseMalformedJoinParseExpectedIdentForAtParseAsteriskIsNotAloneInSelectListParseCannotMixSqbAndWildcardInSelectListParseInvalidContextForWildcardInSelectListIncorrectSQLFunctionArgumentTypeValueParseFailureEvaluatorInvalidArgumentsIntegerOverflowLikeInvalidInputsCastFailedInvalidCastEvaluatorInvalidTimestampFormatPatternEvaluatorInvalidTimestampFormatPatternSymbolForParsingEvaluatorTimestampFormatPatternDuplicateFieldsEvaluatorTimestampFormatPatternHourClockAmPmMismatchEvaluatorUnterminatedTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternSymbolEvaluatorBindingDoesNotExistMissingHeadersInvalidColumnIndexAdminConfigNotificationTargetsFailedAdminProfilerNotEnabledInvalidDecompressedSizeAddUserInvalidArgumentAdminAccountNotEligibleAccountNotEligibleAdminServiceAccountNotFoundPostPolicyConditionInvalidFormat"

var _APIErrorCode_index = [...]uint16{0, 4, 16, 25, 39, 53, 67, 81, 94, 112, 129, 142, 154, 176, 196, 222, 236, 257, 274, 289, 312, 329, 347, 364, 388, 403, 424, 442, 454, 474, 491, 514, 535, 547, 565, 586, 614, 635, 658, 684, 721, 751, 784, 809, 841, 870, 895, 917, 943, 965, 993, 1022, 1056, 1087, 1124, 1154, 1163, 1175, 1191, 1204, 1218, 1236, 1256, 1277, 1293, 1304, 1320, 1348, 1368, 1384, 1412, 1426, 1443, 1458, 1471, 1485, 1498, 1511, 1527, 1544, 1565, 1579, 1600, 1613, 1635, 1658, 1683, 1699, 1714, 1729, 1750, 1768, 1783, 1800, 1825, 1843, 1866, 1881, 1900, 1916, 1935, 1949, 1957, 1976, 1986, 2001, 2037, 2068, 2101, 2130, 2142, 2162, 2186, 2210, 2231, 2255, 2274, 2297, 2323, 2344, 2362, 2389, 2416, 2437, 2458, 2482, 2507, 2535, 2563, 2579, 2590, 2602, 2619, 2634, 2652, 2681, 2698, 2714, 2730, 2748, 2766, 2789, 2810, 2820, 2831, 2845, 2856, 2872, 2895, 2912, 2940, 2959, 2979, 2996, 3014, 3031, 3045, 3080, 3099, 3110, 3123, 3138, 3154, 3172, 3189, 3209, 3230, 3251, 3270, 3289, 3307, 3331, 3355, 3376, 3390, 3414, 3443, 3461, 3478, 3500, 3517, 3535, 3555, 3581, 3597, 3616, 3637, 3641, 3659, 3676, 3702, 3716, 3740, 3761, 3776, 3794, 3817, 3832, 3851, 3868, 3885, 3909, 3936, 3959, 3982, 3999, 4021, 4037, 4057, 4076, 4098, 4119, 4139, 4161, 4185, 4204, 4246, 4267, 4290, 4311, 4342, 4361, 4383, 4403, 4429, 4450, 4472, 4492, 4516, 4539, 4558, 4578, 4600, 4623, 4654, 4692, 4733, 4763, 4777, 4798, 4814, 4836, 4866, 4892, 4920, 4953, 4971, 4994, 5029, 5069, 5111, 5143, 5160, 5185, 5200, 5217, 5227, 5238, 5276, 5330, 5376, 5428, 5476, 5519, 5563, 5591, 5605, 5623, 5659, 5682, 5705, 5727, 5750, 5768, 5795, 5827}

func (i APIErrorCode) String() string {
	if i < 0 || i >= APIErrorCode(len(_APIErrorCode_index)-1) {
		return "APIErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _APIErrorCode_name[_APIErrorCode_index[i]:_APIErrorCode_index[i+1]]
}
