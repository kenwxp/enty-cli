package main

import (
	"entysquare/filer-backend/storage/types"
	"fmt"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/tealeg/xlsx"
	"strconv"
	"testing"
	"time"
)

const insertFilerProductSQL = "" +
	"INSERT INTO filer_product " +
	" (product_id,product_name,node_id,cur_id,period,valid_plan,price,pledge_max,service_rate,node1,node2,shelve_time,create_time,update_time,product_state,is_valid) " +
	" VALUES " +
	" ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;"

const insertFilerOrderSQL = "INSERT INTO filer_order" +
	" (order_id,filer_id,pay_flow,product_id,hold_power,pay_amount,order_time,update_time,valid_time,end_time,order_state)" +
	" VALUES" +
	" ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v') ;"

/*
	TODO 现有用户添加历史订单 生成 SQL 插入语句
	TODO 必填选项 nodeId = 产品关联的节点id ， xlsx表格结尾的用户id
	TODO xlsx表格  ./test_fil/插入案例.xlsx
*/
func TestInsertDB(t *testing.T) {

	nodeId := "node_id1" // TODO() 必填写 ！！！！！！

	var productMap = make(map[string]types.FilerProduct) //产品
	var orders []types.FilerOrder                        //订单
	xlFile, err := xlsx.OpenFile("./test_fil/插入案例.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("表格数据-----------------------------------------------------------------------------")
	fmt.Println()
	for i, row := range xlFile.Sheets[0].Rows {
		//fmt.Println(i, row.Cells)
		name := row.Cells[0]
		filNum, _ := row.Cells[1].Int()
		chargeTime := row.Cells[2]
		powerTime := row.Cells[3]
		tNum, _ := row.Cells[4].Int()
		pledge, _ := row.Cells[5].Float()  //质押返还 0.7
		pledgeStr := row.Cells[5].String() //质押返还 0.7
		reward, _ := row.Cells[6].Float()  //出块分成 0.7 用户得 0.7
		rewardStr := row.Cells[6].String() //出块分成 0.7 用户得 0.7
		userId := row.Cells[7].String()    //出块分成 0.7 用户得 0.7
		if i == 0 {
			fmt.Printf("姓名         充值FIL数     充值时间       算力生效时间     有效算力T      质押返还     产币分成    userID \n")
			continue
		}
		fmt.Printf("%-10v %-10v %-15v %-15v %-10v %-10v %-10v  %-10v\n", name, filNum, chargeTime, powerTime, tNum, pledge, reward, userId)
		price := filNum / tNum
		//区分产品
		if _, ok := productMap["内测-"+pledgeStr+"-"+rewardStr+"-"+strconv.Itoa(price)]; ok {

		} else {
			productMap["内测-"+pledgeStr+"-"+rewardStr+"-"+strconv.Itoa(price)] = types.FilerProduct{
				ProductId:    uuid.NewV4().String(),
				ProductName:  "内测-" + pledgeStr + "-" + rewardStr + "-" + strconv.Itoa(price),
				NodeId:       nodeId,
				CurId:        "fil",
				Period:       "540",
				ValidPlan:    "3",
				Price:        strconv.Itoa(price),
				PledgeMax:    strconv.Itoa(90000 * 1000000000),
				ServiceRate:  strconv.FormatFloat(1.0-reward, 'f', 2, 64),
				Note1:        "",
				Note2:        "",
				ShelveTime:   strconv.FormatInt(time.Now().AddDate(-2, 0, 0).Unix(), 10),
				CreateTime:   strconv.FormatInt(time.Now().AddDate(-2, 0, 0).Unix(), 10),
				UpdateTime:   strconv.FormatInt(time.Now().AddDate(-2, 0, 0).Unix(), 10),
				ProductState: "0",
				IsValid:      "0",
			}
		}
	}

	for i, row := range xlFile.Sheets[0].Rows {
		//fmt.Println(i, row.Cells)
		//name := row.Cells[0]
		filNum, _ := row.Cells[1].Int()
		//chargeTime := row.Cells[2]//充值时间
		powerTime := row.Cells[3]
		tNum, _ := row.Cells[4].Int()
		//pledge, _ := row.Cells[5].Float()  //质押返还 0.7
		pledgeStr := row.Cells[5].String() //质押返还 0.7
		//reward, _ := row.Cells[6].Float()  //出块分成 0.7 用户得 0.7
		rewardStr := row.Cells[6].String() //出块分成 0.7 用户得 0.7
		userId := row.Cells[7].String()    //出块分成 0.7 用户得 0.7
		price := filNum / tNum
		if i == 0 {
			fmt.Printf("姓名         充值FIL数     充值时间       算力生效时间     有效算力T      质押返还     产币分成    userID \n")
			continue
		}
		var cstSh, _ = time.LoadLocation("Asia/Shanghai")                                                //上海
		validPlan, _ := time.ParseInLocation("20060102 15:04:05", powerTime.String()+" 11:11:11", cstSh) //生效时间
		orders = append(orders, types.FilerOrder{
			OrderId:     uuid.NewV4().String(),
			FilerId:     userId,
			PayFlow:     uuid.NewV4().String(), //流水号
			ProductId:   productMap["内测-"+pledgeStr+"-"+rewardStr+"-"+strconv.Itoa(price)].ProductId,
			NodeId:      nodeId,
			Period:      "540",
			ValidPlan:   strconv.FormatInt(validPlan.Unix(), 10),
			ServiceRate: productMap["内测-"+pledgeStr+"-"+rewardStr+"-"+strconv.Itoa(price)].ServiceRate,
			HoldPower:   strconv.Itoa(tNum),
			PayAmount:   strconv.Itoa(filNum * 1000000000),
			OrderTime:   strconv.FormatInt(validPlan.AddDate(0, 0, -3).Unix(), 10),
			UpdateTime:  strconv.FormatInt(time.Now().Unix(), 10),
			ValidTime:   strconv.FormatInt(validPlan.Unix(), 10),
			EndTime:     strconv.FormatInt(validPlan.AddDate(0, 0, 540).Unix(), 10),
			OrderState:  "1", //持仓状态	0-待生效 1-已生效 9-失败
		})
	}

	fmt.Println()
	fmt.Println("区分产品-----------------------------------------------------------------------------")
	fmt.Println()
	for k, v := range productMap {
		fmt.Println(k, v)
	}
	fmt.Println()
	fmt.Println("区分订单-----------------------------------------------------------------------------")
	fmt.Println()
	for k, v := range orders {
		fmt.Println(k, v)
	}
	fmt.Println()
	fmt.Println("SQL-----------------------------------------------------------------------------")
	for _, v := range productMap {
		fmt.Println()
		fmt.Printf(insertFilerProductSQL,
			v.ProductId,
			v.ProductName,
			v.NodeId,
			v.CurId,
			v.Period,
			v.ValidPlan,
			v.Price,
			v.PledgeMax,
			v.ServiceRate,
			v.Note1,
			v.Note2,
			v.ShelveTime,
			v.CreateTime,
			v.UpdateTime,
			v.ProductState,
			v.IsValid,
		)
		//fmt.Println(k, v)
	}
	fmt.Println()
	for _, v := range orders {
		fmt.Println()
		fmt.Printf(insertFilerOrderSQL,
			v.OrderId,
			v.FilerId,
			v.PayFlow,
			v.ProductId,
			v.HoldPower,
			v.PayAmount,
			v.OrderTime,
			v.UpdateTime,
			v.ValidTime,
			v.EndTime,
			v.OrderState,
		)
	}

	fmt.Println()
}
