package models

// Account represents a chart of accounts entry.
// Adapted from Odoo account.account model.
type Account struct {
	BaseModel

	Code          string `gorm:"size:64;not null;index" json:"code"`
	Name          string `gorm:"size:255;not null" json:"name"`
	AccountType   string `gorm:"size:50;not null" json:"account_type"` // asset_receivable, liability_payable, etc.
	InternalGroup string `gorm:"size:50" json:"internal_group"`        // equity, asset, liability, income, expense

	Reconcile  bool `json:"reconcile"`
	Deprecated bool `json:"deprecated"`

	CurrencyID *uint `json:"currency_id,omitempty"`
	CompanyID  uint  `gorm:"index" json:"company_id"`

	Note string `gorm:"type:text" json:"note"`
}

// TableName overrides the table name used by User to `accounts`
func (Account) TableName() string {
	return "accounts"
}
