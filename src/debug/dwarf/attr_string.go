// Code generated by "stringer -type Attr -trimprefix=Attr"; DO NOT EDIT.

package dwarf

import "strconv"

const _Attr_name = "SiblingLocationNameOrderingByteSizeBitOffsetBitSizeStmtListLowpcHighpcLanguageDiscrDiscrValueVisibilityImportStringLengthCommonRefCompDirConstValueContainingTypeDefaultValueInlineIsOptionalLowerBoundProducerPrototypedReturnAddrStartScopeStrideSizeUpperBoundAbstractOriginAccessibilityAddrClassArtificialBaseTypesCallingCountDataMemberLocDeclColumnDeclFileDeclLineDeclarationDiscrListEncodingExternalFrameBaseFriendIdentifierCaseMacroInfoNamelistItemPrioritySegmentSpecificationStaticLinkTypeUseLocationVarParamVirtualityVtableElemLocAllocatedAssociatedDataLocationStrideEntrypcUseUTF8ExtensionRangesTrampolineCallColumnCallFileCallLineDescription"

var _Attr_map = map[Attr]string{
	1:  _Attr_name[0:7],
	2:  _Attr_name[7:15],
	3:  _Attr_name[15:19],
	9:  _Attr_name[19:27],
	11: _Attr_name[27:35],
	12: _Attr_name[35:44],
	13: _Attr_name[44:51],
	16: _Attr_name[51:59],
	17: _Attr_name[59:64],
	18: _Attr_name[64:70],
	19: _Attr_name[70:78],
	21: _Attr_name[78:83],
	22: _Attr_name[83:93],
	23: _Attr_name[93:103],
	24: _Attr_name[103:109],
	25: _Attr_name[109:121],
	26: _Attr_name[121:130],
	27: _Attr_name[130:137],
	28: _Attr_name[137:147],
	29: _Attr_name[147:161],
	30: _Attr_name[161:173],
	32: _Attr_name[173:179],
	33: _Attr_name[179:189],
	34: _Attr_name[189:199],
	37: _Attr_name[199:207],
	39: _Attr_name[207:217],
	42: _Attr_name[217:227],
	44: _Attr_name[227:237],
	46: _Attr_name[237:247],
	47: _Attr_name[247:257],
	49: _Attr_name[257:271],
	50: _Attr_name[271:284],
	51: _Attr_name[284:293],
	52: _Attr_name[293:303],
	53: _Attr_name[303:312],
	54: _Attr_name[312:319],
	55: _Attr_name[319:324],
	56: _Attr_name[324:337],
	57: _Attr_name[337:347],
	58: _Attr_name[347:355],
	59: _Attr_name[355:363],
	60: _Attr_name[363:374],
	61: _Attr_name[374:383],
	62: _Attr_name[383:391],
	63: _Attr_name[391:399],
	64: _Attr_name[399:408],
	65: _Attr_name[408:414],
	66: _Attr_name[414:428],
	67: _Attr_name[428:437],
	68: _Attr_name[437:449],
	69: _Attr_name[449:457],
	70: _Attr_name[457:464],
	71: _Attr_name[464:477],
	72: _Attr_name[477:487],
	73: _Attr_name[487:491],
	74: _Attr_name[491:502],
	75: _Attr_name[502:510],
	76: _Attr_name[510:520],
	77: _Attr_name[520:533],
	78: _Attr_name[533:542],
	79: _Attr_name[542:552],
	80: _Attr_name[552:564],
	81: _Attr_name[564:570],
	82: _Attr_name[570:577],
	83: _Attr_name[577:584],
	84: _Attr_name[584:593],
	85: _Attr_name[593:599],
	86: _Attr_name[599:609],
	87: _Attr_name[609:619],
	88: _Attr_name[619:627],
	89: _Attr_name[627:635],
	90: _Attr_name[635:646],
}

func (i Attr) String() string {
	if str, ok := _Attr_map[i]; ok {
		return str
	}
	return "Attr(" + strconv.FormatInt(int64(i), 10) + ")"
}
