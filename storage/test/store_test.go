package main

//
//func TestInsertAccount(t *testing.T) {
//
//	acc := &types.Account{
//		UserName:    "酱爆",
//		Password:    "123456",
//		PhoneNumber: "13888888888",
//		Email:       "abc@163.com",
//		Gender:      1,
//		Age:         20,
//		GoogleKey:   "Y6XEWOZ7HNFNNBAGCM65CHCZMI3NIDEO",
//	}
//	d, _ := storage.NewDatabase(true)
//	err := d.CreateAccount(context.TODO(), acc)
//	if err != nil {
//		fmt.Print("insert error:", err)
//	}
//}
//
//func TestSelectAccount(t *testing.T) {
//	d, _ := storage.NewDatabase(true)
//	acc, err := d.GetAccountDataByMail("abc@163.com")
//	if err != nil {
//		fmt.Print("select account error:", err)
//	}
//	fmt.Print(*acc)
//}
//
//func TestInsertPresence(t *testing.T) {
//	pre := &types.UserPre{
//		Device: 1,
//		Token:  "333333",
//	}
//	d, _ := storage.NewDatabase(true)
//	err := d.CreateUserPresence(context.TODO(), pre)
//	if err != nil {
//		fmt.Print("insert error:", err)
//	}
//}
//
//func TestSelectAccountPresence(t *testing.T) {
//	db, _ := storage.NewDatabase(true)
//	err := db.InsertPayLog(context.TODO(), types.PayLog{
//		//Id:            "1",
//		Form:          "1",
//		To:            "1",
//		CodeType:      "1",
//		Nun:           "1",
//		PayType:       "1",
//		State:         "1",
//		ServiceCharge: "1",
//		Hash:          "1",
//		UpdateTime:    nil,
//		OrderNum:      "1",
//		LogUserId:     1,
//		Callback:      "1",
//	})
//
//	fmt.Println(err)
//}
