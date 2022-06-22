package sqldb

import (
	"bytes"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	outputsMap map[int64][]SpentInfo
)

// func CallContract(outputsMap map[int64][]sqldb.SpentInfo, tx Tx) {
//	_txInputs := GetUnusedOutputsMap(outputsMap, tx.KeyID)
//	smartVM := &SmartVM{TxSmart: tx, TxInputs: _txInputs}
//	_, _ = VMCallContract(smartVM)
//	txInputsCtx := smartVM.TxInputs
//	txOutputsCtx := smartVM.TxOutputs
//	if len(txInputsCtx) > 0 && len(txOutputsCtx) > 0 {
//		updateTxInputs(outputsMap, tx.Hash, txInputsCtx)
//		insertTxOutputs(outputsMap, tx.Hash, txOutputsCtx)
//	}
// }

func InsertTxOutputs(ecosystem int64, outputTxHash []byte, txOutputsCtx []SpentInfo) {
	for index, txOutput := range txOutputsCtx {
		spentInfos := outputsMap[txOutput.OutputKeyId]
		txOutput.OutputTxHash = outputTxHash
		txOutput.OutputIndex = int32(index)
		// txOutput.Height=height
		spentInfos = append(spentInfos, txOutput)
		outputsMap[txOutput.OutputKeyId] = spentInfos
	}
}
func InsertTxOutputsChange(ecosystem int64, outputTxHash []byte, inputChange SpentInfo, txOutputsCtx []SpentInfo) {
	spentInfosChange := outputsMap[inputChange.OutputKeyId]
	var outputIndex int32
	for index, info := range spentInfosChange {
		if bytes.EqualFold(info.OutputTxHash, outputTxHash) && strings.EqualFold(info.Action, "change") {
			spentInfosChange = append(spentInfosChange[:index], spentInfosChange[index+1:]...) // delete change
			outputIndex = info.OutputIndex
			break
		}
	}
	outputsMap[inputChange.OutputKeyId] = spentInfosChange

	for _, txOutput := range txOutputsCtx {
		spentInfos := outputsMap[txOutput.OutputKeyId]
		txOutput.OutputTxHash = outputTxHash
		txOutput.OutputIndex = outputIndex
		outputIndex++
		spentInfos = append(spentInfos, txOutput)
		outputsMap[txOutput.OutputKeyId] = spentInfos
	}
}

func UpdateTxInputs(ecosystem int64, inputTxHash []byte, txInputsCtx []SpentInfo) {
	var inputIndex int32
	for _, txInput := range txInputsCtx {
		// spentInfos := GetUnusedOutputsMap(outputsMap, txInput.OutputKeyId)
		spentInfos := outputsMap[txInput.OutputKeyId]
		for i, info := range spentInfos {
			if bytes.EqualFold(info.OutputTxHash, txInput.OutputTxHash) &&
				info.OutputKeyId == txInput.OutputKeyId &&
				info.OutputIndex == txInput.OutputIndex &&
				len(txInput.InputTxHash) == 0 && len(info.InputTxHash) == 0 {
				outputsMap[txInput.OutputKeyId][i].InputTxHash = inputTxHash
				outputsMap[txInput.OutputKeyId][i].InputIndex = inputIndex
				inputIndex++
			}
		}
	}
}

func PutAllOutputsMap(outputs []SpentInfo) {
	outputsMap = make(map[int64][]SpentInfo)
	for _, output := range outputs {
		spentInfos := outputsMap[output.OutputKeyId]
		spentInfos = append(spentInfos, output)
		var ecosystem int64 // TODO ecosystem
		PutOutputsMap(ecosystem, output.OutputKeyId, spentInfos)
	}
}
func PutOutputsMap(ecosystem int64, keyID int64, outputs []SpentInfo) {
	outputsMap[keyID] = outputs
}

func GetChangeOutputsMap(ecosystem int64, keyID int64, txHash []byte) *SpentInfo {
	spentInfos := outputsMap[keyID]
	for _, info := range spentInfos {
		if info.Ecosystem == ecosystem && bytes.EqualFold(info.OutputTxHash, txHash) && strings.EqualFold(info.Action, "change") {
			return &info
		}
	}
	return nil
}
func GetUnusedOutputsMap(ecosystem int64, keyID int64) []SpentInfo {
	spentInfos := outputsMap[keyID]
	var inputIndex int32
	var list []SpentInfo
	for _, output := range spentInfos {
		if len(output.InputTxHash) == 0 {
			output.InputIndex = inputIndex
			inputIndex++
			list = append(list, output)
		}
	}
	return list
}
func GetBalance(ecosystem int64, keyID int64) *decimal.Decimal {
	txInputs := GetUnusedOutputsMap(ecosystem, keyID)
	balance := decimal.Zero
	if len(txInputs) > 0 {
		for _, input := range txInputs {
			outputValue, _ := decimal.NewFromString(input.OutputValue)
			balance = balance.Add(outputValue)
		}
		return &balance
	}
	return nil

}

func GetAllOutputs() []SpentInfo {
	var list []SpentInfo
	for _, outputs := range outputsMap {
		list = append(list, outputs...)
	}
	outputsMap = make(map[int64][]SpentInfo)
	return list
}
