package internal

type MB struct {
	ID         string  `gorm:"column:id;primaryKey;type:varchar;not null"`
	CustID     string  `gorm:"column:cust_id;type:varchar;not null"`
	Acct       string  `gorm:"column:acct;type:varchar;not null"`
	Br         string  `gorm:"column:br;type:varchar;not null"`
	Segment    string  `gorm:"column:segment;type:varchar;not null"`
	Product    string  `gorm:"column:product;type:varchar;not null"`
	ContCode   string  `gorm:"column:cont_code;type:varchar;not null"`
	ContType   string  `gorm:"column:cont_type;type:varchar;not null"`
	DocNum     *string `gorm:"column:doc_num;type:varchar"`
	SubsLoanTo *string `gorm:"column:subs_loanto;type:varchar"`
	LineType   *string `gorm:"column:line_type;type:varchar"`
	EndDate    *string `gorm:"column:end_date;type:varchar"`
	AmtTng     *string `gorm:"column:amt_tng;type:varchar"`
	OdTng      *string `gorm:"column:od_tng;type:varchar"`
	Stav       *string `gorm:"column:stav;type:varchar"`
	DayPrPr    *string `gorm:"column:day_pr_pr;type:varchar"`
}
