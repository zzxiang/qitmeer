package blockdag

import (
	"fmt"
	"github.com/Qitmeer/qitmeer/common/hash"
	_ "github.com/Qitmeer/qitmeer/database/ffldb"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
	exit()
}

func Test_GetFutureSet(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}

	//ph:=ibd.(*Phantom)
	anBlock := tbMap[testData.PH_GetFutureSet.Input]
	bset := NewIdSet()
	bd.getFutureSet(bset, anBlock)
	fmt.Printf("Get %s future set：\n", testData.PH_GetFutureSet.Input)
	printBlockSetTag(bset)
	//
	if !processResult(bset, changeToIDList(testData.PH_GetFutureSet.Output)) {
		t.FailNow()
	}
}

func Test_GetAnticone(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	ph := ibd.(*Phantom)
	//
	anBlock := tbMap[testData.PH_GetAnticone.Input]

	////////////
	bset := ph.bd.getAnticone(anBlock, nil)
	fmt.Printf("Get %s anticone set：\n", testData.PH_GetAnticone.Input)
	printBlockSetTag(bset)
	//
	if !processResult(bset, changeToIDList(testData.PH_GetAnticone.Output)) {
		t.FailNow()
	}

}

func Test_BlueSetFig2(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	ph := ibd.(*Phantom)
	//
	blueSet := ph.GetDiffBlueSet()
	fmt.Println("Fig2 diff blue set：")
	printBlockSetTag(blueSet)
	if !processResult(blueSet, changeToIDList(testData.PH_BlueSetFig2.Output)) {
		t.FailNow()
	}
}

func Test_BlueSetFig4(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig4-blocks")
	if ibd == nil {
		t.FailNow()
	}
	ph := ibd.(*Phantom)
	//
	blueSet := ph.GetDiffBlueSet()
	fmt.Println("Fig4 diff blue set：")
	printBlockSetTag(blueSet)
	if !processResult(blueSet, changeToIDList(testData.PH_BlueSetFig4.Output)) {
		t.FailNow()
	}
}

func Test_OrderFig2(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	ph := ibd.(*Phantom)
	order := []uint{}
	var i uint
	ph.UpdateVirtualBlockOrder()
	for i = 0; i < bd.GetBlockTotal(); i++ {
		order = append(order, bd.getBlockByOrder(uint(i)).GetID())
	}
	fmt.Printf("The Fig.2 Order: ")
	printBlockChainTag(order)

	if !processResult(order, changeToIDList(testData.PH_OrderFig2.Output)) {
		t.FailNow()
	}

	//
	da := ph.GetDiffAnticone()
	fmt.Printf("The diffanticoner: ")
	printBlockSetTag(da)
}

func Test_OrderFig4(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig4-blocks")
	if ibd == nil {
		t.FailNow()
	}
	ph := ibd.(*Phantom)
	order := []uint{}
	var i uint
	ph.UpdateVirtualBlockOrder()
	for i = 0; i < bd.GetBlockTotal(); i++ {
		order = append(order, bd.getBlockByOrder(uint(i)).GetID())
	}
	fmt.Printf("The Fig.4 Order: ")
	printBlockChainTag(order)

	if !processResult(order, changeToIDList(testData.PH_OrderFig4.Output)) {
		t.FailNow()
	}

	//
	da := ph.GetDiffAnticone()
	fmt.Printf("The diffanticoner: ")
	printBlockSetTag(da)
}

func Test_GetLayer(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	var result string = ""
	var i uint
	ph := ibd.(*Phantom)
	ph.UpdateVirtualBlockOrder()
	for i = 0; i < bd.GetBlockTotal(); i++ {
		l := bd.GetLayer(bd.getBlockByOrder(uint(i)).GetID())
		result = fmt.Sprintf("%s%d", result, l)
	}
	if result != testData.PH_GetLayer.Output[0] {
		t.FailNow()
	}
}

func Test_IsOnMainChain(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	if strconv.FormatBool(bd.IsOnMainChain(tbMap[testData.PH_IsOnMainChain.Input].GetID())) != testData.PH_IsOnMainChain.Output[0] {
		t.FailNow()
	}
}

func Test_LocateBlocks(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	gs := NewGraphState()
	gs.GetTips().Add(bd.GetGenesisHash())
	gs.SetTotal(1)
	gs.SetLayer(0)
	lb := bd.locateBlocks(gs, 100)
	lbhs := NewHashSet()
	lbhs.AddList(lb)
	if !processResult(lbhs, changeToIDList(testData.PH_LocateBlocks.Output)) {
		t.FailNow()
	}
}

func Test_LocateMaxBlocks(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	gs := NewGraphState()
	gs.GetTips().Add(bd.GetGenesisHash())
	gs.GetTips().Add(tbMap["G"].GetHash())
	gs.SetTotal(4)
	gs.SetLayer(2)
	lb := bd.locateBlocks(gs, 4)
	//printBlockChainTag(lb,tbMap)
	if !processResult(lb, changeToIDList(testData.PH_LocateMaxBlocks.Output)) {
		t.FailNow()
	}
}

func Test_Confirmations(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	mainTip := bd.GetMainChainTip()
	mainChain := []uint{}
	for cur := mainTip; cur != nil; cur = bd.GetBlockById(cur.GetMainParent()) {
		mainChain = append(mainChain, cur.GetID())
	}
	printBlockChainTag(reverseBlockList(mainChain))

	ph := ibd.(*Phantom)
	ph.UpdateVirtualBlockOrder()
	for i := uint(0); i < bd.GetBlockTotal(); i++ {
		blockHash := bd.getBlockByOrder(uint(i)).GetID()
		fmt.Printf("%s : %d\n", getBlockTag(blockHash), bd.GetConfirmations(blockHash))
	}
}

func Test_IsDAG(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	//ph:=ibd.(*Phantom)
	//
	parentsTag := []string{"I", "G"}
	parents := []*hash.Hash{}
	for _, parent := range parentsTag {
		parents = append(parents, tbMap[parent].GetHash())
	}
	block := buildBlock(parents)
	l, _, ib, _ := bd.AddBlock(block)
	if l != nil && l.Len() > 0 {
		tbMap["L"] = ib
	} else {
		t.Fatalf("Error:%d  L\n", tempHash)
	}

}

func Test_IsHourglass(t *testing.T) {
	ibd := InitBlockDAG(phantom, "CP_Blocks")
	if ibd == nil {
		t.FailNow()
	}
	if !bd.IsHourglass(tbMap["J"].GetID()) {
		t.Fatal()
	}
}

func Test_GetMaturity(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	if bd.GetMaturity(tbMap["D"].GetID(), []uint{tbMap["I"].GetID()}) != 2 {
		t.Fatal()
	}
}

func Test_GetMainParentConcurrency(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}

	//ph:=ibd.(*Phantom)
	anBlock := bd.GetBlock(tbMap[testData.PH_MPConcurrency.Input].GetHash())
	//fmt.Println(bd.GetMainParentConcurrency(anBlock))
	if bd.GetMainParentConcurrency(anBlock) != testData.PH_MPConcurrency.Output {
		t.Fatal()
	}
}

func Test_GetBlockConcurrency(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}

	//ph:=ibd.(*Phantom)
	blueNum, err := bd.GetBlockConcurrency(tbMap[testData.PH_MPConcurrency.Input].GetHash())
	if err != nil {
		t.Fatal(err)
	}
	if blueNum != uint(testData.PH_BConcurrency.Output) {
		t.Fatal()
	}
}

func Test_MainChainTip(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	ph := ibd.(*Phantom)
	ph.UpdateVirtualBlockOrder()

	for _, v := range testData.PH_MainChainTip {
		_, ret := bd.CheckSubMainChainTip(getBlocksByTag(v.Input))
		if ret != v.Output {
			t.Fatalf("Main chain tip check:%v is %v not %v", v.Input, ret, v.Output)
		}
	}
}

func Test_Rollback(t *testing.T) {
	ibd := InitBlockDAG(phantom, "PH_fig2-blocks")
	if ibd == nil {
		t.FailNow()
	}
	ph := ibd.(*Phantom)
	orders := NewIdSet()
	total := bd.GetBlockTotal()
	tips := bd.tips.Clone()

	for i := uint(0); i < bd.GetBlockTotal(); i++ {
		ib := ph.bd.getBlockById(i)
		orders.AddPair(ib.GetID(), ib.GetOrder())
	}

	parents := []*hash.Hash{}
	parents = append(parents, tbMap["I"].GetHash())
	parents = append(parents, tbMap["G"].GetHash())

	block := buildBlock(parents)
	l, _, ib, _ := bd.AddBlock(block)
	if l != nil && l.Len() > 0 {
		tbMap["L"] = ib
	} else {
		t.Fatalf("Error:%d  L\n", tempHash)
		return
	}

	bd.rollback()

	if bd.GetBlockTotal() != total {
		t.Fatalf("Roll back error")
	}
	for i := uint(0); i < bd.GetBlockTotal(); i++ {
		ib := ph.bd.getBlockById(i)
		v := orders.Get(i)
		o, ok := v.(uint)
		if !ok {
			t.Fatalf("Roll back error")
		}
		if o != ib.GetOrder() {
			t.Fatalf("Roll back error")
		}
	}

	if !bd.tips.IsEqual(tips) {
		t.Fatalf("Roll back error")
	}
}
