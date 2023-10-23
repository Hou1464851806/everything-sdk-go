package everything

import (
	"everything-sdk-go/utils"
	"syscall"
)

const (
	OK                     = 0 // no error detected
	ERROR_MEMORY           = 1 // out of memory.
	ERROR_IPC              = 2 // Everything search client is not running
	ERROR_REGISTERCLASSEX  = 3 // unable to register window class.
	ERROR_CREATEWINDOW     = 4 // unable to create listening window
	ERROR_CREATETHREAD     = 5 // unable to create listening thread
	ERROR_INVALIDINDEX     = 6 // invalid index
	ERROR_INVALIDCALL      = 7 // invalid call
	ERROR_INVALIDREQUEST   = 8 // invalid request data, request data first.
	ERROR_INVALIDPARAMETER = 9 // bad parameter.

	REQUEST_FILE_NAME                           = 0x00000001
	REQUEST_PATH                                = 0x00000002
	REQUEST_FULL_PATH_AND_FILE_NAME             = 0x00000004
	REQUEST_EXTENSION                           = 0x00000008
	REQUEST_SIZE                                = 0x00000010
	REQUEST_DATE_CREATED                        = 0x00000020
	REQUEST_DATE_MODIFIED                       = 0x00000040
	REQUEST_DATE_ACCESSED                       = 0x00000080
	REQUEST_ATTRIBUTES                          = 0x00000100
	REQUEST_FILE_LIST_FILE_NAME                 = 0x00000200
	REQUEST_RUN_COUNT                           = 0x00000400
	REQUEST_DATE_RUN                            = 0x00000800
	REQUEST_DATE_RECENTLY_CHANGED               = 0x00001000
	REQUEST_HIGHLIGHTED_FILE_NAME               = 0x00002000
	REQUEST_HIGHLIGHTED_PATH                    = 0x00004000
	REQUEST_HIGHLIGHTED_FULL_PATH_AND_FILE_NAME = 0x00008000

	SORT_NAME_ASCENDING                   = 1
	SORT_NAME_DESCENDING                  = 2
	SORT_PATH_ASCENDING                   = 3
	SORT_PATH_DESCENDING                  = 4
	SORT_SIZE_ASCENDING                   = 5
	SORT_SIZE_DESCENDING                  = 6
	SORT_EXTENSION_ASCENDING              = 7
	SORT_EXTENSION_DESCENDING             = 8
	SORT_TYPE_NAME_ASCENDING              = 9
	SORT_TYPE_NAME_DESCENDING             = 10
	SORT_DATE_CREATED_ASCENDING           = 11
	SORT_DATE_CREATED_DESCENDING          = 12
	SORT_DATE_MODIFIED_ASCENDING          = 13
	SORT_DATE_MODIFIED_DESCENDING         = 14
	SORT_ATTRIBUTES_ASCENDING             = 15
	SORT_ATTRIBUTES_DESCENDING            = 16
	SORT_FILE_LIST_FILENAME_ASCENDING     = 17
	SORT_FILE_LIST_FILENAME_DESCENDING    = 18
	SORT_RUN_COUNT_ASCENDING              = 19
	SORT_RUN_COUNT_DESCENDING             = 20
	SORT_DATE_RECENTLY_CHANGED_ASCENDING  = 21
	SORT_DATE_RECENTLY_CHANGED_DESCENDING = 22
	SORT_DATE_ACCESSED_ASCENDING          = 23
	SORT_DATE_ACCESSED_DESCENDING         = 24
	SORT_DATE_RUN_ASCENDING               = 25
	SORT_DATE_RUN_DESCENDING              = 26

	TARGET_MACHINE_X86 = 1
	TARGET_MACHINE_X64 = 2
	TARGET_MACHINE_ARM = 3
)

var everything = syscall.NewLazyDLL("dll/Everything64.dll")

// write search state

func SetSearch(text string) {
	setSearch := everything.NewProc("Everything_SetSearchW")
	setSearch.Call(utils.Str2Ptr(text))
}

func SetMatchPath(isEnable bool) {
	setMatchPath := everything.NewProc("Everything_SetMatchPath")
	setMatchPath.Call(utils.Bool2Ptr(isEnable))
}

func SetMatchCase(isEnable bool) {
	setMatchCase := everything.NewProc("Everything_SetMatchCase")
	setMatchCase.Call(utils.Bool2Ptr(isEnable))
}

func SetMatchWholeWord(isEnable bool) {
	setMatchWholeWord := everything.NewProc("Everything_SetMatchWholeWord")
	setMatchWholeWord.Call(utils.Bool2Ptr(isEnable))
}

func SetMax(max int) {
	setMax := everything.NewProc("Everything_SetMax")
	setMax.Call(utils.Int2Ptr(max))
}

func SetOffset(offset int) {
	setOffset := everything.NewProc("Everything_SetOffset")
	setOffset.Call(utils.Int2Ptr(offset))
}

//Can't get the handle of windows
//This function doesn't work
//func SetReplyWindow(window uintptr) {
//	setReplyWindow := everything.NewProc("Everything_SetReplyWindow")
//	setReplyWindow.Call(window)
//}

func SetReplyID(id int) {
	setReplyID := everything.NewProc("Everything_SetReplyID")
	setReplyID.Call(utils.Int2Ptr(id))
}

func SetSort(flag int) {
	setSort := everything.NewProc("Everything_SetSort")
	setSort.Call(utils.Int2Ptr(flag))
}

func SetRequestFlags(flag int) {
	setRequestFlags := everything.NewProc("Everything_SetRequestFlags")
	setRequestFlags.Call(utils.Int2Ptr(flag))
}

//read search state

//execute query

func Query(isWait bool) bool {
	query := everything.NewProc("Everything_QueryW")
	ret, _, _ := query.Call(utils.Bool2Ptr(isWait))
	return utils.Ptr2Bool(ret)
}

//query reply

//write result state

//read result state

func GetNumResults() int {
	getNumResults := everything.NewProc("Everything_GetNumResults")
	ret, _, _ := getNumResults.Call()
	return utils.Ptr2Int(ret)
}

func GetResultFileName(index int) string {
	getResultFileName := everything.NewProc("Everything_GetResultFileNameW")
	ret, _, _ := getResultFileName.Call(utils.Int2Ptr(index))
	return utils.Ptr2Str(ret)
}

func GetResultPathW(index int) string {
	getResultPathW := everything.NewProc("Everything_GetResultPathW")
	ret, _, _ := getResultPathW.Call(utils.Int2Ptr(index))
	return utils.Ptr2Str(ret)
}

//reset state and free any allocated memory
