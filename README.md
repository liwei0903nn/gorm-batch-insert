# gorm-batch-insert
function for batch insert with gorm 利用gorm批量插入数据


	type Product struct {
		gorm.Model
		Code  string
		Price int
	}

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
  
	db.LogMode(true)
	db.AutoMigrate(&Product{})

	var product [] Product
	for i:= 0; i < 10; i++{
		product = append(product, Product{Code:strconv.Itoa(i), Price:i})
	}

	for i:= 0; i < 10; i++{
		product[i].Price = 2222
	}

	err = util.GormBatchInsert(db, product, nil)  // nil 表示需要插入所有字段
	if err != nil {
		fmt.Println("error", err)
	}

	var allProdecuList []Product
	err = db.Find(&allProdecuList).Error
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("after insert, all data:%+v\n", allProdecuList)

	for i:= 0; i < len(product); i++{
		product[i].Price = 100+i  // update data
	}

	err = util.GormBatchInsertOnDuplicate(db, product, nil, []string{"price"})
	if err != nil {
		fmt.Println("error", err)
	}

	err = db.Find(&allProdecuList).Error
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("after update, all data:%+v\n", allProdecuList)
