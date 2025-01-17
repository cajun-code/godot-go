/*
------------------------------------------------------------------------------
//   This code was generated by template ffi_gdextension_interface.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_gdextension_interface.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------
*/
package gdextensionffi

//revive:disable

import (
	"unsafe"

	"github.com/godot-go/godot-go/pkg/log"
	"go.uber.org/zap"
)

type GDExtensionInterface struct {
	Library      GDExtensionClassLibraryPtr
	Token        unsafe.Pointer
	GodotVersion GDExtensionGodotVersion

	// All of the GDExtension interface functions.
	GetProcAddress                                GDExtensionInterfaceGetProcAddress
	GetGodotVersion                               GDExtensionInterfaceGetGodotVersion
	MemAlloc                                      GDExtensionInterfaceMemAlloc
	MemRealloc                                    GDExtensionInterfaceMemRealloc
	MemFree                                       GDExtensionInterfaceMemFree
	PrintError                                    GDExtensionInterfacePrintError
	PrintErrorWithMessage                         GDExtensionInterfacePrintErrorWithMessage
	PrintWarning                                  GDExtensionInterfacePrintWarning
	PrintWarningWithMessage                       GDExtensionInterfacePrintWarningWithMessage
	PrintScriptError                              GDExtensionInterfacePrintScriptError
	PrintScriptErrorWithMessage                   GDExtensionInterfacePrintScriptErrorWithMessage
	GetNativeStructSize                           GDExtensionInterfaceGetNativeStructSize
	VariantNewCopy                                GDExtensionInterfaceVariantNewCopy
	VariantNewNil                                 GDExtensionInterfaceVariantNewNil
	VariantDestroy                                GDExtensionInterfaceVariantDestroy
	VariantCall                                   GDExtensionInterfaceVariantCall
	VariantCallStatic                             GDExtensionInterfaceVariantCallStatic
	VariantEvaluate                               GDExtensionInterfaceVariantEvaluate
	VariantSet                                    GDExtensionInterfaceVariantSet
	VariantSetNamed                               GDExtensionInterfaceVariantSetNamed
	VariantSetKeyed                               GDExtensionInterfaceVariantSetKeyed
	VariantSetIndexed                             GDExtensionInterfaceVariantSetIndexed
	VariantGet                                    GDExtensionInterfaceVariantGet
	VariantGetNamed                               GDExtensionInterfaceVariantGetNamed
	VariantGetKeyed                               GDExtensionInterfaceVariantGetKeyed
	VariantGetIndexed                             GDExtensionInterfaceVariantGetIndexed
	VariantIterInit                               GDExtensionInterfaceVariantIterInit
	VariantIterNext                               GDExtensionInterfaceVariantIterNext
	VariantIterGet                                GDExtensionInterfaceVariantIterGet
	VariantHash                                   GDExtensionInterfaceVariantHash
	VariantRecursiveHash                          GDExtensionInterfaceVariantRecursiveHash
	VariantHashCompare                            GDExtensionInterfaceVariantHashCompare
	VariantBooleanize                             GDExtensionInterfaceVariantBooleanize
	VariantDuplicate                              GDExtensionInterfaceVariantDuplicate
	VariantStringify                              GDExtensionInterfaceVariantStringify
	VariantGetType                                GDExtensionInterfaceVariantGetType
	VariantHasMethod                              GDExtensionInterfaceVariantHasMethod
	VariantHasMember                              GDExtensionInterfaceVariantHasMember
	VariantHasKey                                 GDExtensionInterfaceVariantHasKey
	VariantGetTypeName                            GDExtensionInterfaceVariantGetTypeName
	VariantCanConvert                             GDExtensionInterfaceVariantCanConvert
	VariantCanConvertStrict                       GDExtensionInterfaceVariantCanConvertStrict
	GetVariantFromTypeConstructor                 GDExtensionInterfaceGetVariantFromTypeConstructor
	GetVariantToTypeConstructor                   GDExtensionInterfaceGetVariantToTypeConstructor
	VariantGetPtrOperatorEvaluator                GDExtensionInterfaceVariantGetPtrOperatorEvaluator
	VariantGetPtrBuiltinMethod                    GDExtensionInterfaceVariantGetPtrBuiltinMethod
	VariantGetPtrConstructor                      GDExtensionInterfaceVariantGetPtrConstructor
	VariantGetPtrDestructor                       GDExtensionInterfaceVariantGetPtrDestructor
	VariantConstruct                              GDExtensionInterfaceVariantConstruct
	VariantGetPtrSetter                           GDExtensionInterfaceVariantGetPtrSetter
	VariantGetPtrGetter                           GDExtensionInterfaceVariantGetPtrGetter
	VariantGetPtrIndexedSetter                    GDExtensionInterfaceVariantGetPtrIndexedSetter
	VariantGetPtrIndexedGetter                    GDExtensionInterfaceVariantGetPtrIndexedGetter
	VariantGetPtrKeyedSetter                      GDExtensionInterfaceVariantGetPtrKeyedSetter
	VariantGetPtrKeyedGetter                      GDExtensionInterfaceVariantGetPtrKeyedGetter
	VariantGetPtrKeyedChecker                     GDExtensionInterfaceVariantGetPtrKeyedChecker
	VariantGetConstantValue                       GDExtensionInterfaceVariantGetConstantValue
	VariantGetPtrUtilityFunction                  GDExtensionInterfaceVariantGetPtrUtilityFunction
	StringNewWithLatin1Chars                      GDExtensionInterfaceStringNewWithLatin1Chars
	StringNewWithUtf8Chars                        GDExtensionInterfaceStringNewWithUtf8Chars
	StringNewWithUtf16Chars                       GDExtensionInterfaceStringNewWithUtf16Chars
	StringNewWithUtf32Chars                       GDExtensionInterfaceStringNewWithUtf32Chars
	StringNewWithWideChars                        GDExtensionInterfaceStringNewWithWideChars
	StringNewWithLatin1CharsAndLen                GDExtensionInterfaceStringNewWithLatin1CharsAndLen
	StringNewWithUtf8CharsAndLen                  GDExtensionInterfaceStringNewWithUtf8CharsAndLen
	StringNewWithUtf16CharsAndLen                 GDExtensionInterfaceStringNewWithUtf16CharsAndLen
	StringNewWithUtf32CharsAndLen                 GDExtensionInterfaceStringNewWithUtf32CharsAndLen
	StringNewWithWideCharsAndLen                  GDExtensionInterfaceStringNewWithWideCharsAndLen
	StringToLatin1Chars                           GDExtensionInterfaceStringToLatin1Chars
	StringToUtf8Chars                             GDExtensionInterfaceStringToUtf8Chars
	StringToUtf16Chars                            GDExtensionInterfaceStringToUtf16Chars
	StringToUtf32Chars                            GDExtensionInterfaceStringToUtf32Chars
	StringToWideChars                             GDExtensionInterfaceStringToWideChars
	StringOperatorIndex                           GDExtensionInterfaceStringOperatorIndex
	StringOperatorIndexConst                      GDExtensionInterfaceStringOperatorIndexConst
	StringOperatorPlusEqString                    GDExtensionInterfaceStringOperatorPlusEqString
	StringOperatorPlusEqChar                      GDExtensionInterfaceStringOperatorPlusEqChar
	StringOperatorPlusEqCstr                      GDExtensionInterfaceStringOperatorPlusEqCstr
	StringOperatorPlusEqWcstr                     GDExtensionInterfaceStringOperatorPlusEqWcstr
	StringOperatorPlusEqC32str                    GDExtensionInterfaceStringOperatorPlusEqC32str
	StringResize                                  GDExtensionInterfaceStringResize
	StringNameNewWithLatin1Chars                  GDExtensionInterfaceStringNameNewWithLatin1Chars
	StringNameNewWithUtf8Chars                    GDExtensionInterfaceStringNameNewWithUtf8Chars
	StringNameNewWithUtf8CharsAndLen              GDExtensionInterfaceStringNameNewWithUtf8CharsAndLen
	XmlParserOpenBuffer                           GDExtensionInterfaceXmlParserOpenBuffer
	FileAccessStoreBuffer                         GDExtensionInterfaceFileAccessStoreBuffer
	FileAccessGetBuffer                           GDExtensionInterfaceFileAccessGetBuffer
	WorkerThreadPoolAddNativeGroupTask            GDExtensionInterfaceWorkerThreadPoolAddNativeGroupTask
	WorkerThreadPoolAddNativeTask                 GDExtensionInterfaceWorkerThreadPoolAddNativeTask
	PackedByteArrayOperatorIndex                  GDExtensionInterfacePackedByteArrayOperatorIndex
	PackedByteArrayOperatorIndexConst             GDExtensionInterfacePackedByteArrayOperatorIndexConst
	PackedColorArrayOperatorIndex                 GDExtensionInterfacePackedColorArrayOperatorIndex
	PackedColorArrayOperatorIndexConst            GDExtensionInterfacePackedColorArrayOperatorIndexConst
	PackedFloat32ArrayOperatorIndex               GDExtensionInterfacePackedFloat32ArrayOperatorIndex
	PackedFloat32ArrayOperatorIndexConst          GDExtensionInterfacePackedFloat32ArrayOperatorIndexConst
	PackedFloat64ArrayOperatorIndex               GDExtensionInterfacePackedFloat64ArrayOperatorIndex
	PackedFloat64ArrayOperatorIndexConst          GDExtensionInterfacePackedFloat64ArrayOperatorIndexConst
	PackedInt32ArrayOperatorIndex                 GDExtensionInterfacePackedInt32ArrayOperatorIndex
	PackedInt32ArrayOperatorIndexConst            GDExtensionInterfacePackedInt32ArrayOperatorIndexConst
	PackedInt64ArrayOperatorIndex                 GDExtensionInterfacePackedInt64ArrayOperatorIndex
	PackedInt64ArrayOperatorIndexConst            GDExtensionInterfacePackedInt64ArrayOperatorIndexConst
	PackedStringArrayOperatorIndex                GDExtensionInterfacePackedStringArrayOperatorIndex
	PackedStringArrayOperatorIndexConst           GDExtensionInterfacePackedStringArrayOperatorIndexConst
	PackedVector2ArrayOperatorIndex               GDExtensionInterfacePackedVector2ArrayOperatorIndex
	PackedVector2ArrayOperatorIndexConst          GDExtensionInterfacePackedVector2ArrayOperatorIndexConst
	PackedVector3ArrayOperatorIndex               GDExtensionInterfacePackedVector3ArrayOperatorIndex
	PackedVector3ArrayOperatorIndexConst          GDExtensionInterfacePackedVector3ArrayOperatorIndexConst
	ArrayOperatorIndex                            GDExtensionInterfaceArrayOperatorIndex
	ArrayOperatorIndexConst                       GDExtensionInterfaceArrayOperatorIndexConst
	ArrayRef                                      GDExtensionInterfaceArrayRef
	ArraySetTyped                                 GDExtensionInterfaceArraySetTyped
	DictionaryOperatorIndex                       GDExtensionInterfaceDictionaryOperatorIndex
	DictionaryOperatorIndexConst                  GDExtensionInterfaceDictionaryOperatorIndexConst
	ObjectMethodBindCall                          GDExtensionInterfaceObjectMethodBindCall
	ObjectMethodBindPtrcall                       GDExtensionInterfaceObjectMethodBindPtrcall
	ObjectDestroy                                 GDExtensionInterfaceObjectDestroy
	GlobalGetSingleton                            GDExtensionInterfaceGlobalGetSingleton
	ObjectGetInstanceBinding                      GDExtensionInterfaceObjectGetInstanceBinding
	ObjectSetInstanceBinding                      GDExtensionInterfaceObjectSetInstanceBinding
	ObjectFreeInstanceBinding                     GDExtensionInterfaceObjectFreeInstanceBinding
	ObjectSetInstance                             GDExtensionInterfaceObjectSetInstance
	ObjectGetClassName                            GDExtensionInterfaceObjectGetClassName
	ObjectCastTo                                  GDExtensionInterfaceObjectCastTo
	ObjectGetInstanceFromId                       GDExtensionInterfaceObjectGetInstanceFromId
	ObjectGetInstanceId                           GDExtensionInterfaceObjectGetInstanceId
	RefGetObject                                  GDExtensionInterfaceRefGetObject
	RefSetObject                                  GDExtensionInterfaceRefSetObject
	ScriptInstanceCreate                          GDExtensionInterfaceScriptInstanceCreate
	ScriptInstanceCreate2                         GDExtensionInterfaceScriptInstanceCreate2
	PlaceHolderScriptInstanceCreate               GDExtensionInterfacePlaceHolderScriptInstanceCreate
	PlaceHolderScriptInstanceUpdate               GDExtensionInterfacePlaceHolderScriptInstanceUpdate
	ObjectGetScriptInstance                       GDExtensionInterfaceObjectGetScriptInstance
	CallableCustomCreate                          GDExtensionInterfaceCallableCustomCreate
	CallableCustomGetUserData                     GDExtensionInterfaceCallableCustomGetUserData
	ClassdbConstructObject                        GDExtensionInterfaceClassdbConstructObject
	ClassdbGetMethodBind                          GDExtensionInterfaceClassdbGetMethodBind
	ClassdbGetClassTag                            GDExtensionInterfaceClassdbGetClassTag
	ClassdbRegisterExtensionClass                 GDExtensionInterfaceClassdbRegisterExtensionClass
	ClassdbRegisterExtensionClass2                GDExtensionInterfaceClassdbRegisterExtensionClass2
	ClassdbRegisterExtensionClassMethod           GDExtensionInterfaceClassdbRegisterExtensionClassMethod
	ClassdbRegisterExtensionClassIntegerConstant  GDExtensionInterfaceClassdbRegisterExtensionClassIntegerConstant
	ClassdbRegisterExtensionClassProperty         GDExtensionInterfaceClassdbRegisterExtensionClassProperty
	ClassdbRegisterExtensionClassPropertyIndexed  GDExtensionInterfaceClassdbRegisterExtensionClassPropertyIndexed
	ClassdbRegisterExtensionClassPropertyGroup    GDExtensionInterfaceClassdbRegisterExtensionClassPropertyGroup
	ClassdbRegisterExtensionClassPropertySubgroup GDExtensionInterfaceClassdbRegisterExtensionClassPropertySubgroup
	ClassdbRegisterExtensionClassSignal           GDExtensionInterfaceClassdbRegisterExtensionClassSignal
	ClassdbUnregisterExtensionClass               GDExtensionInterfaceClassdbUnregisterExtensionClass
	GetLibraryPath                                GDExtensionInterfaceGetLibraryPath
	EditorAddPlugin                               GDExtensionInterfaceEditorAddPlugin
	EditorRemovePlugin                            GDExtensionInterfaceEditorRemovePlugin
}

func (x *GDExtensionInterface) LoadProcAddresses(
	pGetProcAddress GDExtensionInterfaceGetProcAddress,
	pLibrary GDExtensionClassLibraryPtr,
) {
	x.GetProcAddress = pGetProcAddress
	x.Library = pLibrary
	x.Token = unsafe.Pointer(&pLibrary)

	x.GetGodotVersion = (GDExtensionInterfaceGetGodotVersion)(LoadProcAddress("get_godot_version"))
	x.MemAlloc = (GDExtensionInterfaceMemAlloc)(LoadProcAddress("mem_alloc"))
	x.MemRealloc = (GDExtensionInterfaceMemRealloc)(LoadProcAddress("mem_realloc"))
	x.MemFree = (GDExtensionInterfaceMemFree)(LoadProcAddress("mem_free"))
	x.PrintError = (GDExtensionInterfacePrintError)(LoadProcAddress("print_error"))
	x.PrintErrorWithMessage = (GDExtensionInterfacePrintErrorWithMessage)(LoadProcAddress("print_error_with_message"))
	x.PrintWarning = (GDExtensionInterfacePrintWarning)(LoadProcAddress("print_warning"))
	x.PrintWarningWithMessage = (GDExtensionInterfacePrintWarningWithMessage)(LoadProcAddress("print_warning_with_message"))
	x.PrintScriptError = (GDExtensionInterfacePrintScriptError)(LoadProcAddress("print_script_error"))
	x.PrintScriptErrorWithMessage = (GDExtensionInterfacePrintScriptErrorWithMessage)(LoadProcAddress("print_script_error_with_message"))
	x.GetNativeStructSize = (GDExtensionInterfaceGetNativeStructSize)(LoadProcAddress("get_native_struct_size"))
	x.VariantNewCopy = (GDExtensionInterfaceVariantNewCopy)(LoadProcAddress("variant_new_copy"))
	x.VariantNewNil = (GDExtensionInterfaceVariantNewNil)(LoadProcAddress("variant_new_nil"))
	x.VariantDestroy = (GDExtensionInterfaceVariantDestroy)(LoadProcAddress("variant_destroy"))
	x.VariantCall = (GDExtensionInterfaceVariantCall)(LoadProcAddress("variant_call"))
	x.VariantCallStatic = (GDExtensionInterfaceVariantCallStatic)(LoadProcAddress("variant_call_static"))
	x.VariantEvaluate = (GDExtensionInterfaceVariantEvaluate)(LoadProcAddress("variant_evaluate"))
	x.VariantSet = (GDExtensionInterfaceVariantSet)(LoadProcAddress("variant_set"))
	x.VariantSetNamed = (GDExtensionInterfaceVariantSetNamed)(LoadProcAddress("variant_set_named"))
	x.VariantSetKeyed = (GDExtensionInterfaceVariantSetKeyed)(LoadProcAddress("variant_set_keyed"))
	x.VariantSetIndexed = (GDExtensionInterfaceVariantSetIndexed)(LoadProcAddress("variant_set_indexed"))
	x.VariantGet = (GDExtensionInterfaceVariantGet)(LoadProcAddress("variant_get"))
	x.VariantGetNamed = (GDExtensionInterfaceVariantGetNamed)(LoadProcAddress("variant_get_named"))
	x.VariantGetKeyed = (GDExtensionInterfaceVariantGetKeyed)(LoadProcAddress("variant_get_keyed"))
	x.VariantGetIndexed = (GDExtensionInterfaceVariantGetIndexed)(LoadProcAddress("variant_get_indexed"))
	x.VariantIterInit = (GDExtensionInterfaceVariantIterInit)(LoadProcAddress("variant_iter_init"))
	x.VariantIterNext = (GDExtensionInterfaceVariantIterNext)(LoadProcAddress("variant_iter_next"))
	x.VariantIterGet = (GDExtensionInterfaceVariantIterGet)(LoadProcAddress("variant_iter_get"))
	x.VariantHash = (GDExtensionInterfaceVariantHash)(LoadProcAddress("variant_hash"))
	x.VariantRecursiveHash = (GDExtensionInterfaceVariantRecursiveHash)(LoadProcAddress("variant_recursive_hash"))
	x.VariantHashCompare = (GDExtensionInterfaceVariantHashCompare)(LoadProcAddress("variant_hash_compare"))
	x.VariantBooleanize = (GDExtensionInterfaceVariantBooleanize)(LoadProcAddress("variant_booleanize"))
	x.VariantDuplicate = (GDExtensionInterfaceVariantDuplicate)(LoadProcAddress("variant_duplicate"))
	x.VariantStringify = (GDExtensionInterfaceVariantStringify)(LoadProcAddress("variant_stringify"))
	x.VariantGetType = (GDExtensionInterfaceVariantGetType)(LoadProcAddress("variant_get_type"))
	x.VariantHasMethod = (GDExtensionInterfaceVariantHasMethod)(LoadProcAddress("variant_has_method"))
	x.VariantHasMember = (GDExtensionInterfaceVariantHasMember)(LoadProcAddress("variant_has_member"))
	x.VariantHasKey = (GDExtensionInterfaceVariantHasKey)(LoadProcAddress("variant_has_key"))
	x.VariantGetTypeName = (GDExtensionInterfaceVariantGetTypeName)(LoadProcAddress("variant_get_type_name"))
	x.VariantCanConvert = (GDExtensionInterfaceVariantCanConvert)(LoadProcAddress("variant_can_convert"))
	x.VariantCanConvertStrict = (GDExtensionInterfaceVariantCanConvertStrict)(LoadProcAddress("variant_can_convert_strict"))
	x.GetVariantFromTypeConstructor = (GDExtensionInterfaceGetVariantFromTypeConstructor)(LoadProcAddress("get_variant_from_type_constructor"))
	x.GetVariantToTypeConstructor = (GDExtensionInterfaceGetVariantToTypeConstructor)(LoadProcAddress("get_variant_to_type_constructor"))
	x.VariantGetPtrOperatorEvaluator = (GDExtensionInterfaceVariantGetPtrOperatorEvaluator)(LoadProcAddress("variant_get_ptr_operator_evaluator"))
	x.VariantGetPtrBuiltinMethod = (GDExtensionInterfaceVariantGetPtrBuiltinMethod)(LoadProcAddress("variant_get_ptr_builtin_method"))
	x.VariantGetPtrConstructor = (GDExtensionInterfaceVariantGetPtrConstructor)(LoadProcAddress("variant_get_ptr_constructor"))
	x.VariantGetPtrDestructor = (GDExtensionInterfaceVariantGetPtrDestructor)(LoadProcAddress("variant_get_ptr_destructor"))
	x.VariantConstruct = (GDExtensionInterfaceVariantConstruct)(LoadProcAddress("variant_construct"))
	x.VariantGetPtrSetter = (GDExtensionInterfaceVariantGetPtrSetter)(LoadProcAddress("variant_get_ptr_setter"))
	x.VariantGetPtrGetter = (GDExtensionInterfaceVariantGetPtrGetter)(LoadProcAddress("variant_get_ptr_getter"))
	x.VariantGetPtrIndexedSetter = (GDExtensionInterfaceVariantGetPtrIndexedSetter)(LoadProcAddress("variant_get_ptr_indexed_setter"))
	x.VariantGetPtrIndexedGetter = (GDExtensionInterfaceVariantGetPtrIndexedGetter)(LoadProcAddress("variant_get_ptr_indexed_getter"))
	x.VariantGetPtrKeyedSetter = (GDExtensionInterfaceVariantGetPtrKeyedSetter)(LoadProcAddress("variant_get_ptr_keyed_setter"))
	x.VariantGetPtrKeyedGetter = (GDExtensionInterfaceVariantGetPtrKeyedGetter)(LoadProcAddress("variant_get_ptr_keyed_getter"))
	x.VariantGetPtrKeyedChecker = (GDExtensionInterfaceVariantGetPtrKeyedChecker)(LoadProcAddress("variant_get_ptr_keyed_checker"))
	x.VariantGetConstantValue = (GDExtensionInterfaceVariantGetConstantValue)(LoadProcAddress("variant_get_constant_value"))
	x.VariantGetPtrUtilityFunction = (GDExtensionInterfaceVariantGetPtrUtilityFunction)(LoadProcAddress("variant_get_ptr_utility_function"))
	x.StringNewWithLatin1Chars = (GDExtensionInterfaceStringNewWithLatin1Chars)(LoadProcAddress("string_new_with_latin1_chars"))
	x.StringNewWithUtf8Chars = (GDExtensionInterfaceStringNewWithUtf8Chars)(LoadProcAddress("string_new_with_utf8_chars"))
	x.StringNewWithUtf16Chars = (GDExtensionInterfaceStringNewWithUtf16Chars)(LoadProcAddress("string_new_with_utf16_chars"))
	x.StringNewWithUtf32Chars = (GDExtensionInterfaceStringNewWithUtf32Chars)(LoadProcAddress("string_new_with_utf32_chars"))
	x.StringNewWithWideChars = (GDExtensionInterfaceStringNewWithWideChars)(LoadProcAddress("string_new_with_wide_chars"))
	x.StringNewWithLatin1CharsAndLen = (GDExtensionInterfaceStringNewWithLatin1CharsAndLen)(LoadProcAddress("string_new_with_latin1_chars_and_len"))
	x.StringNewWithUtf8CharsAndLen = (GDExtensionInterfaceStringNewWithUtf8CharsAndLen)(LoadProcAddress("string_new_with_utf8_chars_and_len"))
	x.StringNewWithUtf16CharsAndLen = (GDExtensionInterfaceStringNewWithUtf16CharsAndLen)(LoadProcAddress("string_new_with_utf16_chars_and_len"))
	x.StringNewWithUtf32CharsAndLen = (GDExtensionInterfaceStringNewWithUtf32CharsAndLen)(LoadProcAddress("string_new_with_utf32_chars_and_len"))
	x.StringNewWithWideCharsAndLen = (GDExtensionInterfaceStringNewWithWideCharsAndLen)(LoadProcAddress("string_new_with_wide_chars_and_len"))
	x.StringToLatin1Chars = (GDExtensionInterfaceStringToLatin1Chars)(LoadProcAddress("string_to_latin1_chars"))
	x.StringToUtf8Chars = (GDExtensionInterfaceStringToUtf8Chars)(LoadProcAddress("string_to_utf8_chars"))
	x.StringToUtf16Chars = (GDExtensionInterfaceStringToUtf16Chars)(LoadProcAddress("string_to_utf16_chars"))
	x.StringToUtf32Chars = (GDExtensionInterfaceStringToUtf32Chars)(LoadProcAddress("string_to_utf32_chars"))
	x.StringToWideChars = (GDExtensionInterfaceStringToWideChars)(LoadProcAddress("string_to_wide_chars"))
	x.StringOperatorIndex = (GDExtensionInterfaceStringOperatorIndex)(LoadProcAddress("string_operator_index"))
	x.StringOperatorIndexConst = (GDExtensionInterfaceStringOperatorIndexConst)(LoadProcAddress("string_operator_index_const"))
	x.StringOperatorPlusEqString = (GDExtensionInterfaceStringOperatorPlusEqString)(LoadProcAddress("string_operator_plus_eq_string"))
	x.StringOperatorPlusEqChar = (GDExtensionInterfaceStringOperatorPlusEqChar)(LoadProcAddress("string_operator_plus_eq_char"))
	x.StringOperatorPlusEqCstr = (GDExtensionInterfaceStringOperatorPlusEqCstr)(LoadProcAddress("string_operator_plus_eq_cstr"))
	x.StringOperatorPlusEqWcstr = (GDExtensionInterfaceStringOperatorPlusEqWcstr)(LoadProcAddress("string_operator_plus_eq_wcstr"))
	x.StringOperatorPlusEqC32str = (GDExtensionInterfaceStringOperatorPlusEqC32str)(LoadProcAddress("string_operator_plus_eq_c32str"))
	x.StringResize = (GDExtensionInterfaceStringResize)(LoadProcAddress("string_resize"))
	x.StringNameNewWithLatin1Chars = (GDExtensionInterfaceStringNameNewWithLatin1Chars)(LoadProcAddress("string_name_new_with_latin1_chars"))
	x.StringNameNewWithUtf8Chars = (GDExtensionInterfaceStringNameNewWithUtf8Chars)(LoadProcAddress("string_name_new_with_utf8_chars"))
	x.StringNameNewWithUtf8CharsAndLen = (GDExtensionInterfaceStringNameNewWithUtf8CharsAndLen)(LoadProcAddress("string_name_new_with_utf8_chars_and_len"))
	x.XmlParserOpenBuffer = (GDExtensionInterfaceXmlParserOpenBuffer)(LoadProcAddress("xml_parser_open_buffer"))
	x.FileAccessStoreBuffer = (GDExtensionInterfaceFileAccessStoreBuffer)(LoadProcAddress("file_access_store_buffer"))
	x.FileAccessGetBuffer = (GDExtensionInterfaceFileAccessGetBuffer)(LoadProcAddress("file_access_get_buffer"))
	x.WorkerThreadPoolAddNativeGroupTask = (GDExtensionInterfaceWorkerThreadPoolAddNativeGroupTask)(LoadProcAddress("worker_thread_pool_add_native_group_task"))
	x.WorkerThreadPoolAddNativeTask = (GDExtensionInterfaceWorkerThreadPoolAddNativeTask)(LoadProcAddress("worker_thread_pool_add_native_task"))
	x.PackedByteArrayOperatorIndex = (GDExtensionInterfacePackedByteArrayOperatorIndex)(LoadProcAddress("packed_byte_array_operator_index"))
	x.PackedByteArrayOperatorIndexConst = (GDExtensionInterfacePackedByteArrayOperatorIndexConst)(LoadProcAddress("packed_byte_array_operator_index_const"))
	x.PackedColorArrayOperatorIndex = (GDExtensionInterfacePackedColorArrayOperatorIndex)(LoadProcAddress("packed_color_array_operator_index"))
	x.PackedColorArrayOperatorIndexConst = (GDExtensionInterfacePackedColorArrayOperatorIndexConst)(LoadProcAddress("packed_color_array_operator_index_const"))
	x.PackedFloat32ArrayOperatorIndex = (GDExtensionInterfacePackedFloat32ArrayOperatorIndex)(LoadProcAddress("packed_float32_array_operator_index"))
	x.PackedFloat32ArrayOperatorIndexConst = (GDExtensionInterfacePackedFloat32ArrayOperatorIndexConst)(LoadProcAddress("packed_float32_array_operator_index_const"))
	x.PackedFloat64ArrayOperatorIndex = (GDExtensionInterfacePackedFloat64ArrayOperatorIndex)(LoadProcAddress("packed_float64_array_operator_index"))
	x.PackedFloat64ArrayOperatorIndexConst = (GDExtensionInterfacePackedFloat64ArrayOperatorIndexConst)(LoadProcAddress("packed_float64_array_operator_index_const"))
	x.PackedInt32ArrayOperatorIndex = (GDExtensionInterfacePackedInt32ArrayOperatorIndex)(LoadProcAddress("packed_int32_array_operator_index"))
	x.PackedInt32ArrayOperatorIndexConst = (GDExtensionInterfacePackedInt32ArrayOperatorIndexConst)(LoadProcAddress("packed_int32_array_operator_index_const"))
	x.PackedInt64ArrayOperatorIndex = (GDExtensionInterfacePackedInt64ArrayOperatorIndex)(LoadProcAddress("packed_int64_array_operator_index"))
	x.PackedInt64ArrayOperatorIndexConst = (GDExtensionInterfacePackedInt64ArrayOperatorIndexConst)(LoadProcAddress("packed_int64_array_operator_index_const"))
	x.PackedStringArrayOperatorIndex = (GDExtensionInterfacePackedStringArrayOperatorIndex)(LoadProcAddress("packed_string_array_operator_index"))
	x.PackedStringArrayOperatorIndexConst = (GDExtensionInterfacePackedStringArrayOperatorIndexConst)(LoadProcAddress("packed_string_array_operator_index_const"))
	x.PackedVector2ArrayOperatorIndex = (GDExtensionInterfacePackedVector2ArrayOperatorIndex)(LoadProcAddress("packed_vector2_array_operator_index"))
	x.PackedVector2ArrayOperatorIndexConst = (GDExtensionInterfacePackedVector2ArrayOperatorIndexConst)(LoadProcAddress("packed_vector2_array_operator_index_const"))
	x.PackedVector3ArrayOperatorIndex = (GDExtensionInterfacePackedVector3ArrayOperatorIndex)(LoadProcAddress("packed_vector3_array_operator_index"))
	x.PackedVector3ArrayOperatorIndexConst = (GDExtensionInterfacePackedVector3ArrayOperatorIndexConst)(LoadProcAddress("packed_vector3_array_operator_index_const"))
	x.ArrayOperatorIndex = (GDExtensionInterfaceArrayOperatorIndex)(LoadProcAddress("array_operator_index"))
	x.ArrayOperatorIndexConst = (GDExtensionInterfaceArrayOperatorIndexConst)(LoadProcAddress("array_operator_index_const"))
	x.ArrayRef = (GDExtensionInterfaceArrayRef)(LoadProcAddress("array_ref"))
	x.ArraySetTyped = (GDExtensionInterfaceArraySetTyped)(LoadProcAddress("array_set_typed"))
	x.DictionaryOperatorIndex = (GDExtensionInterfaceDictionaryOperatorIndex)(LoadProcAddress("dictionary_operator_index"))
	x.DictionaryOperatorIndexConst = (GDExtensionInterfaceDictionaryOperatorIndexConst)(LoadProcAddress("dictionary_operator_index_const"))
	x.ObjectMethodBindCall = (GDExtensionInterfaceObjectMethodBindCall)(LoadProcAddress("object_method_bind_call"))
	x.ObjectMethodBindPtrcall = (GDExtensionInterfaceObjectMethodBindPtrcall)(LoadProcAddress("object_method_bind_ptrcall"))
	x.ObjectDestroy = (GDExtensionInterfaceObjectDestroy)(LoadProcAddress("object_destroy"))
	x.GlobalGetSingleton = (GDExtensionInterfaceGlobalGetSingleton)(LoadProcAddress("global_get_singleton"))
	x.ObjectGetInstanceBinding = (GDExtensionInterfaceObjectGetInstanceBinding)(LoadProcAddress("object_get_instance_binding"))
	x.ObjectSetInstanceBinding = (GDExtensionInterfaceObjectSetInstanceBinding)(LoadProcAddress("object_set_instance_binding"))
	x.ObjectFreeInstanceBinding = (GDExtensionInterfaceObjectFreeInstanceBinding)(LoadProcAddress("object_free_instance_binding"))
	x.ObjectSetInstance = (GDExtensionInterfaceObjectSetInstance)(LoadProcAddress("object_set_instance"))
	x.ObjectGetClassName = (GDExtensionInterfaceObjectGetClassName)(LoadProcAddress("object_get_class_name"))
	x.ObjectCastTo = (GDExtensionInterfaceObjectCastTo)(LoadProcAddress("object_cast_to"))
	x.ObjectGetInstanceFromId = (GDExtensionInterfaceObjectGetInstanceFromId)(LoadProcAddress("object_get_instance_from_id"))
	x.ObjectGetInstanceId = (GDExtensionInterfaceObjectGetInstanceId)(LoadProcAddress("object_get_instance_id"))
	x.RefGetObject = (GDExtensionInterfaceRefGetObject)(LoadProcAddress("ref_get_object"))
	x.RefSetObject = (GDExtensionInterfaceRefSetObject)(LoadProcAddress("ref_set_object"))
	x.ScriptInstanceCreate = (GDExtensionInterfaceScriptInstanceCreate)(LoadProcAddress("script_instance_create"))
	x.ScriptInstanceCreate2 = (GDExtensionInterfaceScriptInstanceCreate2)(LoadProcAddress("script_instance_create2"))
	x.PlaceHolderScriptInstanceCreate = (GDExtensionInterfacePlaceHolderScriptInstanceCreate)(LoadProcAddress("placeholder_script_instance_create"))
	x.PlaceHolderScriptInstanceUpdate = (GDExtensionInterfacePlaceHolderScriptInstanceUpdate)(LoadProcAddress("placeholder_script_instance_update"))
	x.ObjectGetScriptInstance = (GDExtensionInterfaceObjectGetScriptInstance)(LoadProcAddress("object_get_script_instance"))
	x.CallableCustomCreate = (GDExtensionInterfaceCallableCustomCreate)(LoadProcAddress("callable_custom_create"))
	x.CallableCustomGetUserData = (GDExtensionInterfaceCallableCustomGetUserData)(LoadProcAddress("callable_custom_get_user_data"))
	x.ClassdbConstructObject = (GDExtensionInterfaceClassdbConstructObject)(LoadProcAddress("classdb_construct_object"))
	x.ClassdbGetMethodBind = (GDExtensionInterfaceClassdbGetMethodBind)(LoadProcAddress("classdb_get_method_bind"))
	x.ClassdbGetClassTag = (GDExtensionInterfaceClassdbGetClassTag)(LoadProcAddress("classdb_get_class_tag"))
	x.ClassdbRegisterExtensionClass = (GDExtensionInterfaceClassdbRegisterExtensionClass)(LoadProcAddress("classdb_register_extension_class"))
	x.ClassdbRegisterExtensionClass2 = (GDExtensionInterfaceClassdbRegisterExtensionClass2)(LoadProcAddress("classdb_register_extension_class2"))
	x.ClassdbRegisterExtensionClassMethod = (GDExtensionInterfaceClassdbRegisterExtensionClassMethod)(LoadProcAddress("classdb_register_extension_class_method"))
	x.ClassdbRegisterExtensionClassIntegerConstant = (GDExtensionInterfaceClassdbRegisterExtensionClassIntegerConstant)(LoadProcAddress("classdb_register_extension_class_integer_constant"))
	x.ClassdbRegisterExtensionClassProperty = (GDExtensionInterfaceClassdbRegisterExtensionClassProperty)(LoadProcAddress("classdb_register_extension_class_property"))
	x.ClassdbRegisterExtensionClassPropertyIndexed = (GDExtensionInterfaceClassdbRegisterExtensionClassPropertyIndexed)(LoadProcAddress("classdb_register_extension_class_property_indexed"))
	x.ClassdbRegisterExtensionClassPropertyGroup = (GDExtensionInterfaceClassdbRegisterExtensionClassPropertyGroup)(LoadProcAddress("classdb_register_extension_class_property_group"))
	x.ClassdbRegisterExtensionClassPropertySubgroup = (GDExtensionInterfaceClassdbRegisterExtensionClassPropertySubgroup)(LoadProcAddress("classdb_register_extension_class_property_subgroup"))
	x.ClassdbRegisterExtensionClassSignal = (GDExtensionInterfaceClassdbRegisterExtensionClassSignal)(LoadProcAddress("classdb_register_extension_class_signal"))
	x.ClassdbUnregisterExtensionClass = (GDExtensionInterfaceClassdbUnregisterExtensionClass)(LoadProcAddress("classdb_unregister_extension_class"))
	x.GetLibraryPath = (GDExtensionInterfaceGetLibraryPath)(LoadProcAddress("get_library_path"))
	x.EditorAddPlugin = (GDExtensionInterfaceEditorAddPlugin)(LoadProcAddress("editor_add_plugin"))
	x.EditorRemovePlugin = (GDExtensionInterfaceEditorRemovePlugin)(LoadProcAddress("editor_remove_plugin"))
}

var (
	FFI GDExtensionInterface
)

func LoadProcAddress(funcName string) unsafe.Pointer {
	ret := CallFunc_GDExtensionInterfaceGetProcAddress(funcName)
	if ret == nil {
		log.Warn("GDExtensionInterfaceGetProcAddress Error",
			zap.String("name", funcName),
		)
	}
	return unsafe.Pointer(ret)
}

func (gv GDExtensionGodotVersion) GetMajor() int32 {
	return int32(gv.major)
}

func (gv GDExtensionGodotVersion) GetMinor() int32 {
	return int32(gv.minor)
}
