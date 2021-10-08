package storage

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	host = "127.0.0.1"
	//host     = "156.240.109.163"
	port     = 5532
	user     = "entycli"
	password = "entycli666"
	dbname   = "entycli"
)

// Database represents an account database
type Database struct {
	Db                 *sql.DB
	filerAccount       filerAccountsStatements
	filerBlockIncome   filerBlockIncomeStatements
	filerBlockTemp     filerBlockTempStatements
	filerProduct       filerProductStatements
	filerOrder         filerOrderStatements
	filerOrderIncome   filerOrderIncomeStatements
	filerPool          filerPoolStatements
	filerPoolIncome    filerPoolIncomeStatements
	filerAccountIncome filerAccountIncomeStatements
	filerStatControl   filerStatControlStatements
	filerBalanceIncome filerBalanceIncomeStatements
	filerBalanceFlow   filerBalanceFlowStatements
}

func (d *Database) WithTransaction(fn func(txn *sql.Tx) error) (err error) {
	return util.WithTransaction(d.Db, fn)
}

// NewDatabase creates a new accounts and profiles database
func NewDatabase() (*Database, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	d := &Database{
		Db: db,
	}

	// Create tables before executing migrations so we don't fail if the table is missing,
	// and THEN prepare statements so we don't fail due to referencing new columns
	if err = d.filerAccount.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerAccount.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}

	if err = d.filerBlockTemp.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}

	if err = d.filerBlockTemp.prepare(db); err != nil {
		return nil, err
	}
	if err = d.filerBlockIncome.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerBlockIncome.prepare(db); err != nil {
		return nil, err
	}
	if err = d.filerProduct.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerProduct.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}

	if err = d.filerOrder.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerOrder.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerOrderIncome.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerOrderIncome.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerPool.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerPool.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}

	if err = d.filerPoolIncome.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerPoolIncome.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}

	if err = d.filerAccountIncome.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerAccountIncome.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}

	if err = d.filerStatControl.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerStatControl.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerBalanceIncome.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerBalanceIncome.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerBalanceFlow.execSchema(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	if err = d.filerBalanceFlow.prepare(db); err != nil {
		fmt.Print("err1:", err)
		return nil, err
	}
	return d, nil
}

func (d *Database) SelectAccountByName(ctx context.Context, txn *sql.Tx, name string) (account *types.FilerAccountInfo, err error) {
	account, err = d.filerAccount.selectAccountByName(ctx, txn, name)
	if err != nil {
		return
	}
	return
}

//FilerBlockTemp
func (d *Database) InsertFilerBlockTemp(ctx context.Context, txn *sql.Tx, f *types.FilerBlockTemp) (err error) {
	err = d.filerBlockTemp.insertFilerBlockTemp(ctx, txn, f)
	if err != nil {
		return err
	}
	return
}

func (d *Database) SelectFilerBlockTempMaxBlockHeight(ctx context.Context) (height int64, err error) {
	return d.filerBlockTemp.selectFilerBlockTempMaxBlockHeight(ctx)
}

func (d *Database) SelectFilerBlockTempMaxBlockHeightByNodeId(ctx context.Context, nodeId string) (height int64, err error) {
	return d.filerBlockTemp.selectFilerBlockTempMaxBlockHeightByNodeId(ctx, nodeId)
}

//FilerBlock
func (d *Database) InsertFilerBlockIncome(ctx context.Context, txn *sql.Tx, f *types.FilerBlockIncome) (err error) {
	err = d.filerBlockIncome.insertFilerBlockIncome(ctx, txn, f)
	if err != nil {
		return err
	}
	return
}

func (d *Database) SelectStatisticBlockInfoByNodeId(ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) ([]types.FilerBlockIncome, error) {
	return d.filerBlockTemp.selectStatisticBlockInfoByNodeId(ctx, txn, statTime, nodeId)
}

func (d *Database) UpdateFilerBlockTempState(ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) error {
	return d.filerBlockTemp.updateFilerBlockTempState(ctx, txn, statTime, nodeId)
}

func (d *Database) SelectFilerBlockIncomeByStatTimeAndNodeId(ctx context.Context, startTime string, NodeId string) *types.FilerBlockIncome {
	return d.filerBlockIncome.selectFilerBlockIncomeByStatTimeAndNodeId(ctx, startTime, NodeId)
}
func (d *Database) SelectFilerBlockAndTimeByTime(ctx context.Context, start, end time.Time) ([]types.BlockRecord, error) {
	return d.filerBlockIncome.selectFilerBlockAndTimeByTime(ctx, start, end)
}

//filerProduct
func (d *Database) InsertFilerProduct(ctx context.Context, txn *sql.Tx, f *types.FilerProduct) (err error) {
	err = d.filerProduct.insertFilerProduct(ctx, txn, f)
	if err != nil {
		return err
	}
	return
}
func (d *Database) UpdateFilerProductState(ctx context.Context, txn *sql.Tx, f *types.FilerProduct) (err error) {
	err = d.filerProduct.updateProductState(ctx, txn, f.ProductId, f.ProductState, f.IsValid)
	if err != nil {
		return err
	}
	return
}
func (d *Database) SelectProductInfoById(ctx context.Context, txn *sql.Tx, productId string) (productInfo types.FilerProduct, err error) {
	productInfo, err = d.filerProduct.selectProductInfoByProductID(ctx, txn, productId)
	if err != nil {
		return
	}
	return
}
func (d *Database) SelectProductListByCurId(ctx context.Context, txn *sql.Tx, curId string) (list map[string]types.FilerProduct, err error) {
	list, err = d.filerProduct.selectProductListByCurId(ctx, txn, curId)
	if err != nil {
		return nil, err
	}
	return
}
func (d *Database) SelectProductListByState(ctx context.Context, txn *sql.Tx, state string) (list map[string]types.FilerProduct, err error) {
	list, err = d.filerProduct.selectProductListByState(ctx, txn, state)
	if err != nil {
		return nil, err
	}
	return
}
func (d *Database) SelectFilerBlockIncomeByStatRange(ctx context.Context, txn *sql.Tx, state string, statTime string, nodeId string) ([]types.FilerBlockIncome, error) {
	return d.filerBlockIncome.selectFilerBlockIncomeByStatRange(ctx, txn, state, statTime, nodeId)
}

//filer_pool
func (d *Database) InsertFilerPool(ctx context.Context, txn *sql.Tx, f *types.FilerPool) (err error) {
	return d.filerPool.insertFilerPool(ctx, txn, f)
}

func (d *Database) SelectFilerPool() ([]types.FilerPool, error) {
	return d.filerPool.selectFilerPool()
}

func (d *Database) SelectAllFilerPool(ctx context.Context, txn *sql.Tx) ([]types.FilerPool, error) {
	return d.filerPool.selectAllFilerPool(ctx, txn)
}

func (d *Database) SelectAllInProgressOrderByNodeId(ctx context.Context, txn *sql.Tx, statTimeStr string, nodeId string) ([]types.FilerOrder, error) {
	return d.filerOrder.selectAllInProgressOrderByNodeId(ctx, txn, statTimeStr, nodeId)
}

func (d *Database) SelectFilerOrderIncomeByStatTime(ctx context.Context, txn *sql.Tx, statTime string, orderId string, nodeId string) ([]types.FilerOrderIncome, error) {
	return d.filerOrderIncome.selectFilerOrderIncomeByStatTime(ctx, txn, statTime, orderId, nodeId)
}
func (d *Database) InsertFilerOrderIncome(ctx context.Context, txn *sql.Tx, f *types.FilerOrderIncome) (err error) {
	return d.filerOrderIncome.insertFilerOrderIncome(ctx, txn, f)
}

func (d *Database) UpdateFilerOrderState(ctx context.Context, txn *sql.Tx, orderState string, orderId string) (err error) {
	return d.filerOrder.updateFilerOrderState(ctx, txn, orderState, orderId)
}

//filer_pool_income
func (d *Database) InsertFilerPoolIncome(ctx context.Context, txn *sql.Tx, f *types.FilerPoolIncome) (err error) {
	return d.filerPoolIncome.insertFilerPoolIncome(ctx, txn, f)
}

func (d *Database) SelectFilerPoolIncomeByNodeId(ctx context.Context, txn *sql.Tx, NodeId, statTime string) (*types.FilerPoolIncome, error) {
	return d.filerPoolIncome.selectFilerPoolIncomeByNodeId(ctx, txn, NodeId, statTime)
}
func (d *Database) SelectFilerPoolIncomeAll() ([]types.FilerPoolIncome, error) {
	return d.filerPoolIncome.selectFilerPoolIncomeAll()
}

func (d *Database) SelectOrderListByFilerId(ctx context.Context, txn *sql.Tx, state int, filerId int64) (list map[string]types.FilerOrder, err error) {
	list, err = d.filerOrder.selectOrderListByFilerId(ctx, txn, state, filerId)
	if err != nil {
		return nil, err
	}
	return
}
func (d *Database) SelectFilerOrderIncomeByOrderId(ctx context.Context, txn *sql.Tx, orderId string) (types.FilerOrderIncome, error) {
	return d.filerOrderIncome.selectFilerOrderIncomeByOrderId(ctx, txn, orderId)
}
func (d *Database) SelectFilerAccountIncomeByFilerId(ctx context.Context, txn *sql.Tx, filerId int64) (list map[string]types.FilerAccountIncome, err error) {
	list, err = d.filerAccountIncome.selectFilerAccountIncomeByFilerId(ctx, txn, filerId)
	if err != nil {
		return nil, err
	}
	return
}
func (d *Database) SelectStatisticOrderIncomeByNodeId(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) ([]types.FilerAccountIncome, error) {
	return d.filerOrderIncome.selectStatisticOrderIncomeByNodeId(ctx, txn, statTime, nodeId)
}

func (d *Database) SelectFilerAccountIncomeByStatTime(ctx context.Context, txn *sql.Tx, statTime string, filerId string, nodeId string) ([]types.FilerAccountIncome, error) {
	return d.filerAccountIncome.selectFilerAccountIncomeByStatTime(ctx, txn, statTime, filerId, nodeId)
}

func (d *Database) InsertFilerAccountIncome(ctx context.Context, txn *sql.Tx, f *types.FilerAccountIncome) (err error) {
	return d.filerAccountIncome.insertFilerAccountIncome(ctx, txn, f)
}
func (d *Database) SelectHtFilerPool(ctx context.Context, txn *sql.Tx) (htPool types.FilerPool, err error) {
	htPool, err = d.filerPool.selectHtFilerPool(ctx, txn)
	if err != nil {
		return
	}
	return
}

func (d *Database) SelectFilerPoolIncomeByNodeIdLast(ctx context.Context, NodeId string) *types.FilerPoolIncome {
	return d.filerPoolIncome.selectFilerPoolIncomeByNodeIdLast(ctx, NodeId)
}

func (d *Database) SelectStatControlInfo(ctx context.Context, txn *sql.Tx, statType string, statTimeStr string, nodeId string) (*types.FilerStatControl, error) {
	return d.filerStatControl.selectStatControlInfo(ctx, txn, statType, statTimeStr, nodeId)
}
func (d *Database) SelectAllNodeStatControlInfo(ctx context.Context, txn *sql.Tx, statType string, statTime string) ([]types.FilerStatControl, error) {
	return d.filerStatControl.selectAllNodeStatControlInfo(ctx, txn, statType, statTime)
}

func (d *Database) InsertStatControlInfo(ctx context.Context, txn *sql.Tx, statControl *types.FilerStatControl) (err error) {
	return d.filerStatControl.insertStatControlInfo(ctx, txn, statControl)
}

func (d *Database) UpdateStatControlInfo(ctx context.Context, txn *sql.Tx, statControl *types.FilerStatControl) (err error) {
	return d.filerStatControl.updateStatControlInfo(ctx, txn, statControl)
}

func (d *Database) DeleteStatControlFromStatTimeByNodeId(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	return d.filerStatControl.deleteStatControlFromStatTimeByNodeId(ctx, txn, statTime, nodeId)
}
func (d *Database) DeleteAllStatControlFromStatTime(ctx context.Context, txn *sql.Tx, statTime string) (err error) {
	return d.filerStatControl.deleteAllStatControlFromStatTime(ctx, txn, statTime)
}

func (d *Database) DeleteFilerBlockIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	return d.filerBlockIncome.deleteFilerBlockIncomeFromStatTime(ctx, txn, statTime, nodeId)
}

func (d *Database) DeleteFilerOrderIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	return d.filerOrderIncome.deleteFilerOrderIncomeFromStatTime(ctx, txn, statTime, nodeId)
}

func (d *Database) DeleteFilerAccountIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	return d.filerAccountIncome.deleteFilerAccountIncomeFromStatTime(ctx, txn, statTime, nodeId)
}

func (d *Database) DeleteFilerPoolIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	return d.filerPoolIncome.deleteFilerPoolIncomeFromStatTime(ctx, txn, statTime, nodeId)
}

func (d *Database) SelectFilerBalanceIncomeByFilerId(ctx context.Context, txn *sql.Tx, filerId string) ([]types.FilerBalanceIncome, error) {
	return d.filerBalanceIncome.selectFilerBalanceIncomeByFilerId(ctx, txn, filerId)
}

func (d *Database) SelectFilerBalanceIncomeListByFilerId(ctx context.Context, txn *sql.Tx, filerId string) ([]types.FilerBalanceIncome, error) {
	return d.filerBalanceIncome.selectFilerBalanceIncomeListByFilerId(ctx, txn, filerId)
}
func (d *Database) InsertFilerOrder(ctx context.Context, txn *sql.Tx, f *types.FilerOrder) (err error) {
	err = d.filerOrder.insertFilerOrder(ctx, txn, f)
	if err != nil {
		return err
	}
	return
}

func (d *Database) SelectFilerBalanceIncomeByStatTime(ctx context.Context, txn *sql.Tx, statTime string, filerId string) ([]types.FilerBalanceIncome, error) {

	return d.filerBalanceIncome.selectFilerBalanceIncomeByStatTime(ctx, txn, statTime, filerId)
}

func (d *Database) InsertFilerBalanceIncome(ctx context.Context, txn *sql.Tx, f *types.FilerBalanceIncome) (err error) {
	return d.filerBalanceIncome.insertFilerBalanceIncome(ctx, txn, f)
}
func (d *Database) UpdateFilerBalanceIncomeByUuid(ctx context.Context, txn *sql.Tx, f *types.FilerBalanceIncome) (err error) {
	return d.filerBalanceIncome.updateFilerBalanceIncomeByUuid(ctx, txn, f)
}

func (d *Database) SelectLatestBalanceListForEachFiler(ctx context.Context, txn *sql.Tx) ([]types.FilerBalanceIncome, error) {
	return d.filerBalanceIncome.selectLatestBalanceListForEachFiler(ctx, txn)
}

func (d *Database) DeleteFilerBalanceIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string) (err error) {
	return d.filerBalanceIncome.deleteFilerBalanceIncomeFromStatTime(ctx, txn, statTime)
}

func (d *Database) SelectStatisticBalanceIncome(ctx context.Context, txn *sql.Tx, statTime string) ([]types.FilerBalanceIncome, error) {
	return d.filerAccountIncome.selectStatisticBalanceIncome(ctx, txn, statTime)
}

func (d *Database) SelectStatisticBalanceFlowListByFilerId(ctx context.Context, txn *sql.Tx, filerId string, statTimeStr string) ([]types.FilerBalanceFlow, error) {
	return d.filerBalanceFlow.selectStatisticBalanceFlowListByFilerId(ctx, txn, filerId, statTimeStr)
}

func (d *Database) InsertBalanceFlow(ctx context.Context, txn *sql.Tx, f *types.FilerBalanceFlow) (err error) {
	return d.filerBalanceFlow.insertBalanceFlow(ctx, txn, f)
}

func (d *Database) DeleteBalanceFlowByOperType(ctx context.Context, txn *sql.Tx, filerId string, statTimeStr string, operType string) (err error) {
	return d.filerBalanceFlow.deleteBalanceFlowByOperType(ctx, txn, filerId, statTimeStr, operType)
}
func (d *Database) SelectFilerOrderListWithIncomeInfoByFilerIdAndOrderState(ctx context.Context, txn *sql.Tx, filerId string, orderState string) ([]types.FilerOrderShow, error) {
	return d.filerOrderIncome.selectFilerOrderListWithIncomeInfoByFilerIdAndOrderState(ctx, txn, filerId, orderState)
}

func (d *Database) SelectInvalidOrderListByFilerId(ctx context.Context, txn *sql.Tx, filerId string) ([]types.FilerOrderShow, error) {
	return d.filerOrder.selectInvalidOrderListByFilerId(ctx, txn, filerId)
}
