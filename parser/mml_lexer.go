// Code generated from Mml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MmlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var mmllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mmllexerLexerInit() {
	staticData := &mmllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "':'", "';'", "','", "'='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "ACT", "BLK", "CHECK", "CLB", "DEL", "DLD", "DSP",
		"HIS", "HISTORY", "LST", "MOD", "PING", "POWEROP", "QRY", "RBK", "RMV",
		"RST", "SCAN", "SET", "SHW", "SND", "STOP", "SWITCH", "SYNC", "UNACT",
		"UNBLK", "UPLD", "ADD", "Attribute", "Identifier", "Letter", "Value",
		"Digit", "NumberInt", "IntegerLiteral", "Number", "QuotedString", "ESC",
		"Pattern", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "ACT", "BLK", "CHECK", "CLB", "DEL",
		"DLD", "DSP", "HIS", "HISTORY", "LST", "MOD", "PING", "POWEROP", "QRY",
		"RBK", "RMV", "RST", "SCAN", "SET", "SHW", "SND", "STOP", "SWITCH",
		"SYNC", "UNACT", "UNBLK", "UPLD", "ADD", "Attribute", "Identifier",
		"Letter", "Value", "Digit", "NumberInt", "IntegerLiteral", "Number",
		"QuotedString", "ESC", "Pattern", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 309, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1,
		33, 5, 33, 237, 8, 33, 10, 33, 12, 33, 240, 9, 33, 1, 34, 1, 34, 1, 35,
		1, 35, 3, 35, 246, 8, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 3,
		37, 254, 8, 37, 1, 38, 1, 38, 5, 38, 258, 8, 38, 10, 38, 12, 38, 261, 9,
		38, 1, 39, 4, 39, 264, 8, 39, 11, 39, 12, 39, 265, 1, 39, 5, 39, 269, 8,
		39, 10, 39, 12, 39, 272, 9, 39, 1, 39, 1, 39, 5, 39, 276, 8, 39, 10, 39,
		12, 39, 279, 9, 39, 3, 39, 281, 8, 39, 1, 40, 1, 40, 1, 40, 5, 40, 286,
		8, 40, 10, 40, 12, 40, 289, 9, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 3, 41, 297, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 4, 43, 304, 8,
		43, 11, 43, 12, 43, 305, 1, 43, 1, 43, 1, 287, 0, 44, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31,
		63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40,
		81, 41, 83, 42, 85, 43, 87, 44, 1, 0, 26, 2, 0, 65, 65, 97, 97, 2, 0, 67,
		67, 99, 99, 2, 0, 84, 84, 116, 116, 2, 0, 66, 66, 98, 98, 2, 0, 76, 76,
		108, 108, 2, 0, 75, 75, 107, 107, 2, 0, 72, 72, 104, 104, 2, 0, 69, 69,
		101, 101, 2, 0, 68, 68, 100, 100, 2, 0, 83, 83, 115, 115, 2, 0, 80, 80,
		112, 112, 2, 0, 73, 73, 105, 105, 2, 0, 79, 79, 111, 111, 2, 0, 82, 82,
		114, 114, 2, 0, 89, 89, 121, 121, 2, 0, 77, 77, 109, 109, 2, 0, 78, 78,
		110, 110, 2, 0, 71, 71, 103, 103, 2, 0, 87, 87, 119, 119, 2, 0, 81, 81,
		113, 113, 2, 0, 86, 86, 118, 118, 2, 0, 85, 85, 117, 117, 3, 0, 65, 90,
		97, 122, 124, 124, 1, 0, 48, 57, 1, 0, 49, 57, 3, 0, 9, 10, 13, 13, 32,
		32, 324, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0,
		0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 1, 89, 1, 0, 0, 0, 3, 91, 1, 0,
		0, 0, 5, 93, 1, 0, 0, 0, 7, 95, 1, 0, 0, 0, 9, 97, 1, 0, 0, 0, 11, 101,
		1, 0, 0, 0, 13, 105, 1, 0, 0, 0, 15, 111, 1, 0, 0, 0, 17, 115, 1, 0, 0,
		0, 19, 119, 1, 0, 0, 0, 21, 123, 1, 0, 0, 0, 23, 127, 1, 0, 0, 0, 25, 131,
		1, 0, 0, 0, 27, 139, 1, 0, 0, 0, 29, 143, 1, 0, 0, 0, 31, 147, 1, 0, 0,
		0, 33, 152, 1, 0, 0, 0, 35, 160, 1, 0, 0, 0, 37, 164, 1, 0, 0, 0, 39, 168,
		1, 0, 0, 0, 41, 172, 1, 0, 0, 0, 43, 176, 1, 0, 0, 0, 45, 181, 1, 0, 0,
		0, 47, 185, 1, 0, 0, 0, 49, 189, 1, 0, 0, 0, 51, 193, 1, 0, 0, 0, 53, 198,
		1, 0, 0, 0, 55, 204, 1, 0, 0, 0, 57, 209, 1, 0, 0, 0, 59, 215, 1, 0, 0,
		0, 61, 221, 1, 0, 0, 0, 63, 226, 1, 0, 0, 0, 65, 230, 1, 0, 0, 0, 67, 232,
		1, 0, 0, 0, 69, 241, 1, 0, 0, 0, 71, 245, 1, 0, 0, 0, 73, 247, 1, 0, 0,
		0, 75, 253, 1, 0, 0, 0, 77, 255, 1, 0, 0, 0, 79, 280, 1, 0, 0, 0, 81, 282,
		1, 0, 0, 0, 83, 296, 1, 0, 0, 0, 85, 298, 1, 0, 0, 0, 87, 303, 1, 0, 0,
		0, 89, 90, 5, 58, 0, 0, 90, 2, 1, 0, 0, 0, 91, 92, 5, 59, 0, 0, 92, 4,
		1, 0, 0, 0, 93, 94, 5, 44, 0, 0, 94, 6, 1, 0, 0, 0, 95, 96, 5, 61, 0, 0,
		96, 8, 1, 0, 0, 0, 97, 98, 7, 0, 0, 0, 98, 99, 7, 1, 0, 0, 99, 100, 7,
		2, 0, 0, 100, 10, 1, 0, 0, 0, 101, 102, 7, 3, 0, 0, 102, 103, 7, 4, 0,
		0, 103, 104, 7, 5, 0, 0, 104, 12, 1, 0, 0, 0, 105, 106, 7, 1, 0, 0, 106,
		107, 7, 6, 0, 0, 107, 108, 7, 7, 0, 0, 108, 109, 7, 1, 0, 0, 109, 110,
		7, 5, 0, 0, 110, 14, 1, 0, 0, 0, 111, 112, 7, 1, 0, 0, 112, 113, 7, 4,
		0, 0, 113, 114, 7, 3, 0, 0, 114, 16, 1, 0, 0, 0, 115, 116, 7, 8, 0, 0,
		116, 117, 7, 7, 0, 0, 117, 118, 7, 4, 0, 0, 118, 18, 1, 0, 0, 0, 119, 120,
		7, 8, 0, 0, 120, 121, 7, 4, 0, 0, 121, 122, 7, 8, 0, 0, 122, 20, 1, 0,
		0, 0, 123, 124, 7, 8, 0, 0, 124, 125, 7, 9, 0, 0, 125, 126, 7, 10, 0, 0,
		126, 22, 1, 0, 0, 0, 127, 128, 7, 6, 0, 0, 128, 129, 7, 11, 0, 0, 129,
		130, 7, 9, 0, 0, 130, 24, 1, 0, 0, 0, 131, 132, 7, 6, 0, 0, 132, 133, 7,
		11, 0, 0, 133, 134, 7, 9, 0, 0, 134, 135, 7, 2, 0, 0, 135, 136, 7, 12,
		0, 0, 136, 137, 7, 13, 0, 0, 137, 138, 7, 14, 0, 0, 138, 26, 1, 0, 0, 0,
		139, 140, 7, 4, 0, 0, 140, 141, 7, 9, 0, 0, 141, 142, 7, 2, 0, 0, 142,
		28, 1, 0, 0, 0, 143, 144, 7, 15, 0, 0, 144, 145, 7, 12, 0, 0, 145, 146,
		7, 8, 0, 0, 146, 30, 1, 0, 0, 0, 147, 148, 7, 10, 0, 0, 148, 149, 7, 11,
		0, 0, 149, 150, 7, 16, 0, 0, 150, 151, 7, 17, 0, 0, 151, 32, 1, 0, 0, 0,
		152, 153, 7, 10, 0, 0, 153, 154, 7, 12, 0, 0, 154, 155, 7, 18, 0, 0, 155,
		156, 7, 7, 0, 0, 156, 157, 7, 13, 0, 0, 157, 158, 7, 12, 0, 0, 158, 159,
		7, 10, 0, 0, 159, 34, 1, 0, 0, 0, 160, 161, 7, 19, 0, 0, 161, 162, 7, 13,
		0, 0, 162, 163, 7, 14, 0, 0, 163, 36, 1, 0, 0, 0, 164, 165, 7, 13, 0, 0,
		165, 166, 7, 3, 0, 0, 166, 167, 7, 5, 0, 0, 167, 38, 1, 0, 0, 0, 168, 169,
		7, 13, 0, 0, 169, 170, 7, 15, 0, 0, 170, 171, 7, 20, 0, 0, 171, 40, 1,
		0, 0, 0, 172, 173, 7, 13, 0, 0, 173, 174, 7, 9, 0, 0, 174, 175, 7, 2, 0,
		0, 175, 42, 1, 0, 0, 0, 176, 177, 7, 9, 0, 0, 177, 178, 7, 1, 0, 0, 178,
		179, 7, 0, 0, 0, 179, 180, 7, 16, 0, 0, 180, 44, 1, 0, 0, 0, 181, 182,
		7, 9, 0, 0, 182, 183, 7, 7, 0, 0, 183, 184, 7, 2, 0, 0, 184, 46, 1, 0,
		0, 0, 185, 186, 7, 9, 0, 0, 186, 187, 7, 6, 0, 0, 187, 188, 7, 18, 0, 0,
		188, 48, 1, 0, 0, 0, 189, 190, 7, 9, 0, 0, 190, 191, 7, 16, 0, 0, 191,
		192, 7, 8, 0, 0, 192, 50, 1, 0, 0, 0, 193, 194, 7, 9, 0, 0, 194, 195, 7,
		2, 0, 0, 195, 196, 7, 12, 0, 0, 196, 197, 7, 10, 0, 0, 197, 52, 1, 0, 0,
		0, 198, 199, 7, 9, 0, 0, 199, 200, 7, 18, 0, 0, 200, 201, 7, 11, 0, 0,
		201, 202, 7, 1, 0, 0, 202, 203, 7, 6, 0, 0, 203, 54, 1, 0, 0, 0, 204, 205,
		7, 9, 0, 0, 205, 206, 7, 14, 0, 0, 206, 207, 7, 16, 0, 0, 207, 208, 7,
		1, 0, 0, 208, 56, 1, 0, 0, 0, 209, 210, 7, 21, 0, 0, 210, 211, 7, 16, 0,
		0, 211, 212, 7, 0, 0, 0, 212, 213, 7, 1, 0, 0, 213, 214, 7, 2, 0, 0, 214,
		58, 1, 0, 0, 0, 215, 216, 7, 21, 0, 0, 216, 217, 7, 16, 0, 0, 217, 218,
		7, 3, 0, 0, 218, 219, 7, 4, 0, 0, 219, 220, 7, 5, 0, 0, 220, 60, 1, 0,
		0, 0, 221, 222, 7, 21, 0, 0, 222, 223, 7, 10, 0, 0, 223, 224, 7, 4, 0,
		0, 224, 225, 7, 8, 0, 0, 225, 62, 1, 0, 0, 0, 226, 227, 7, 0, 0, 0, 227,
		228, 7, 8, 0, 0, 228, 229, 7, 8, 0, 0, 229, 64, 1, 0, 0, 0, 230, 231, 3,
		67, 33, 0, 231, 66, 1, 0, 0, 0, 232, 238, 3, 69, 34, 0, 233, 237, 3, 69,
		34, 0, 234, 237, 3, 73, 36, 0, 235, 237, 5, 95, 0, 0, 236, 233, 1, 0, 0,
		0, 236, 234, 1, 0, 0, 0, 236, 235, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238,
		236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 68, 1, 0, 0, 0, 240, 238, 1,
		0, 0, 0, 241, 242, 7, 22, 0, 0, 242, 70, 1, 0, 0, 0, 243, 246, 3, 75, 37,
		0, 244, 246, 3, 85, 42, 0, 245, 243, 1, 0, 0, 0, 245, 244, 1, 0, 0, 0,
		246, 72, 1, 0, 0, 0, 247, 248, 7, 23, 0, 0, 248, 74, 1, 0, 0, 0, 249, 254,
		3, 77, 38, 0, 250, 254, 3, 73, 36, 0, 251, 254, 3, 79, 39, 0, 252, 254,
		3, 81, 40, 0, 253, 249, 1, 0, 0, 0, 253, 250, 1, 0, 0, 0, 253, 251, 1,
		0, 0, 0, 253, 252, 1, 0, 0, 0, 254, 76, 1, 0, 0, 0, 255, 259, 7, 24, 0,
		0, 256, 258, 3, 73, 36, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0,
		259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 78, 1, 0, 0, 0, 261, 259,
		1, 0, 0, 0, 262, 264, 3, 73, 36, 0, 263, 262, 1, 0, 0, 0, 264, 265, 1,
		0, 0, 0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 281, 1, 0, 0,
		0, 267, 269, 3, 73, 36, 0, 268, 267, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0,
		270, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 273, 1, 0, 0, 0, 272,
		270, 1, 0, 0, 0, 273, 277, 5, 46, 0, 0, 274, 276, 3, 73, 36, 0, 275, 274,
		1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0,
		0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 280, 263, 1, 0, 0, 0,
		280, 270, 1, 0, 0, 0, 281, 80, 1, 0, 0, 0, 282, 287, 5, 34, 0, 0, 283,
		286, 3, 83, 41, 0, 284, 286, 9, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 284,
		1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 287, 285, 1, 0,
		0, 0, 288, 290, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 290, 291, 5, 34, 0, 0,
		291, 82, 1, 0, 0, 0, 292, 293, 5, 92, 0, 0, 293, 297, 5, 34, 0, 0, 294,
		295, 5, 92, 0, 0, 295, 297, 5, 92, 0, 0, 296, 292, 1, 0, 0, 0, 296, 294,
		1, 0, 0, 0, 297, 84, 1, 0, 0, 0, 298, 299, 5, 34, 0, 0, 299, 300, 3, 67,
		33, 0, 300, 301, 5, 34, 0, 0, 301, 86, 1, 0, 0, 0, 302, 304, 7, 25, 0,
		0, 303, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305,
		306, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 6, 43, 0, 0, 308, 88,
		1, 0, 0, 0, 14, 0, 236, 238, 245, 253, 259, 265, 270, 277, 280, 285, 287,
		296, 305, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MmlLexerInit initializes any static state used to implement MmlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMmlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MmlLexerInit() {
	staticData := &mmllexerLexerStaticData
	staticData.once.Do(mmllexerLexerInit)
}

// NewMmlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMmlLexer(input antlr.CharStream) *MmlLexer {
	MmlLexerInit()
	l := new(MmlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &mmllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Mml.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MmlLexer tokens.
const (
	MmlLexerT__0           = 1
	MmlLexerT__1           = 2
	MmlLexerT__2           = 3
	MmlLexerT__3           = 4
	MmlLexerACT            = 5
	MmlLexerBLK            = 6
	MmlLexerCHECK          = 7
	MmlLexerCLB            = 8
	MmlLexerDEL            = 9
	MmlLexerDLD            = 10
	MmlLexerDSP            = 11
	MmlLexerHIS            = 12
	MmlLexerHISTORY        = 13
	MmlLexerLST            = 14
	MmlLexerMOD            = 15
	MmlLexerPING           = 16
	MmlLexerPOWEROP        = 17
	MmlLexerQRY            = 18
	MmlLexerRBK            = 19
	MmlLexerRMV            = 20
	MmlLexerRST            = 21
	MmlLexerSCAN           = 22
	MmlLexerSET            = 23
	MmlLexerSHW            = 24
	MmlLexerSND            = 25
	MmlLexerSTOP           = 26
	MmlLexerSWITCH         = 27
	MmlLexerSYNC           = 28
	MmlLexerUNACT          = 29
	MmlLexerUNBLK          = 30
	MmlLexerUPLD           = 31
	MmlLexerADD            = 32
	MmlLexerAttribute      = 33
	MmlLexerIdentifier     = 34
	MmlLexerLetter         = 35
	MmlLexerValue          = 36
	MmlLexerDigit          = 37
	MmlLexerNumberInt      = 38
	MmlLexerIntegerLiteral = 39
	MmlLexerNumber         = 40
	MmlLexerQuotedString   = 41
	MmlLexerESC            = 42
	MmlLexerPattern        = 43
	MmlLexerWS             = 44
)
