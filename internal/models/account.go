package models

import (
	"time"
)

// AccountMove represents a journal entry (invoice, bill, or manual entry).
// Adapted from Odoo account.move model.
type AccountMove struct {
	BaseModel

	Name                string    `gorm:"size:255;not null;index" json:"name"`
	Ref                 string    `gorm:"size:255" json:"ref"`
	Date                time.Time `gorm:"index;not null" json:"date"`
	State               string    `gorm:"size:50;default:'draft';index" json:"state"` // draft, posted, cancel
	MoveType            string    `gorm:"size:50;index" json:"move_type"`             // entry, out_invoice, in_invoice, etc.
	JournalID           uint      `gorm:"index" json:"journal_id"`
	CompanyID           uint      `gorm:"index" json:"company_id"`
	CurrencyID          uint      `json:"currency_id"`
	PartnerID           *uint     `gorm:"index" json:"partner_id,omitempty"`
	CommercialPartnerID *uint     `gorm:"index" json:"commercial_partner_id,omitempty"`

	// Amounts
	AmountUntaxed  float64 `gorm:"type:decimal(20,4)" json:"amount_untaxed"`
	AmountTax      float64 `gorm:"type:decimal(20,4)" json:"amount_tax"`
	AmountTotal    float64 `gorm:"type:decimal(20,4)" json:"amount_total"`
	AmountResidual float64 `gorm:"type:decimal(20,4)" json:"amount_residual"`

	// Payment
	PaymentState   string     `gorm:"size:50" json:"payment_state"` // not_paid, in_payment, paid
	InvoiceDate    *time.Time `json:"invoice_date,omitempty"`
	InvoiceDateDue *time.Time `json:"invoice_date_due,omitempty"`

	// Relationships
	Lines []AccountMoveLine `gorm:"foreignKey:MoveID" json:"lines"`
}

// AccountMoveLine represents a journal item (individual debit/credit line).
// Adapted from Odoo account.move.line model.
type AccountMoveLine struct {
	BaseModel

	MoveID    uint   `gorm:"index;not null" json:"move_id"`
	MoveName  string `gorm:"size:255" json:"move_name"`
	JournalID uint   `gorm:"index" json:"journal_id"`
	CompanyID uint   `gorm:"index" json:"company_id"`
	AccountID uint   `gorm:"index;not null" json:"account_id"`
	PartnerID *uint  `gorm:"index" json:"partner_id,omitempty"`

	Name string `gorm:"size:255" json:"name"` // Label
	Ref  string `gorm:"size:255" json:"ref"`

	// Amounts
	Debit   float64 `gorm:"type:decimal(20,4);default:0" json:"debit"`
	Credit  float64 `gorm:"type:decimal(20,4);default:0" json:"credit"`
	Balance float64 `gorm:"type:decimal(20,4);default:0" json:"balance"` // Debit - Credit

	AmountCurrency float64 `gorm:"type:decimal(20,4)" json:"amount_currency"`
	CurrencyID     *uint   `json:"currency_id,omitempty"`

	// Analytic
	AnalyticDistribution string `gorm:"type:jsonb" json:"analytic_distribution"`

	// Tax
	TaxLineID  *uint `json:"tax_line_id,omitempty"`
	TaxGroupID *uint `json:"tax_group_id,omitempty"`

	// Reconciliation
	Reconciled      bool  `json:"reconciled"`
	FullReconcileID *uint `json:"full_reconcile_id,omitempty"`
}

// TableName overrides the table name used by User to `account_moves`
func (AccountMove) TableName() string {
	return "account_moves"
}

// TableName overrides the table name used by User to `account_move_lines`
func (AccountMoveLine) TableName() string {
	return "account_move_lines"
}
