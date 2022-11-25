package config

import (
	"time"
)

// InspectionConfig includes all the configurations available when fetching inspections
type InspectionConfig struct {
	ModifiedAfter time.Time
	Archived      string
	Completed     string
	Incremental   bool
	Limit         int
	SkipIDs       []string
	WebReportLink string
}

// ExportMediaConfig is a representation of the export.media section from yaml configuration
type ExportMediaConfig struct {
	Export bool   // The flag to export media with CSV and SQL exports or not
	Path   string // The absolute or relative path to save exported data
}

// ExportSiteConfig is a representation of the export.site section from yaml configuration
type ExportSiteConfig struct {
	IncludeDeleted       bool // The flag to include or not include deleted sites in sites table exports
	IncludeFullHierarchy bool // The flag to include full sites hierarchy in table e.g. areas, regions, etc
}

// ExportInspectionConfig is a representation of the export.inspection section from yaml configuration
type ExportInspectionConfig struct {
	Archived             string   // The flag to export archived or active or both archived and active inspections
	Completed            string   // The flag to export completed or incomplete or both completed and incomplete inspections
	IncludeInactiveItems bool     // The flag to include or not include inactive question items in the inspection_items
	BatchLimit           int      // The limit for the number of inspections that gets processed per batch
	SkipIDs              []string // The inspection IDs to skip for inspection exports
	WebReportLink        string   // The flag to export private or public inspection report links
}

// ExportActionConfig is a representation of the export.action section from yaml configuration
type ExportActionConfig struct {
	BatchLimit int // The limit for the number of actions that gets processed per batch
}

// ExportIssueConfig is a representation of the export.issue section from yaml configuration
type ExportIssueConfig struct {
	BatchLimit int // The limit for the number of issues that gets processed per batch
}

// ExportConfig is a representation of the EXPORT section from YAML configuration
type ExportConfig struct {
	Path               string    // The absolute or relative path to save exported data
	Incremental        bool      // The flag to remember or not remember where the last export run
	ModifiedAfter      time.Time // The timestamp in Coordinated Universal Time (UTC) to start inspections and actions exports
	FilterByTemplateID []string  // The template IDs to filter by for inspections and schedules export
	FilterByTableName  []string  // The CSV and/or SQL tables to export. Empty means all tables will be exported
	SchemaOnly         bool
	ActionConfig       *ExportActionConfig
	IssueConfig        *ExportIssueConfig
	InspectionConfig   *ExportInspectionConfig
	SiteConfig         *ExportSiteConfig
	MediaConfig        *ExportMediaConfig
}

// ApiConfig is a representation of the API fields from YAML configuration for SafetyCulture
type ApiConfig struct {
	AccessToken string //the API Token. Maps to access_token
}

// SheqsyApiConfig is a representation of the API fields from YAML configuration for Sheqsy
type SheqsyApiConfig struct {
	UserName  string // maps to sheqsy_username
	CompanyID string // maps to sheqsy_company_id
}

// DBConfig is a representation of the DB section from YAML configuration
type DBConfig struct {
	Dialect          string // The database dialect
	ConnectionString string // The database connection string
}

// CSVConfig is a representation of the CSV section from YAML configuration
type CSVConfig struct {
	MaxRowsPerFile int
}

// ReportConfig is a representation of the REPORT section from YAML configuration
type ReportConfig struct {
	Format             []string // The format for inspection report exports
	PreferenceID       string   // The report layout to apply for PDF and Word inspection reports
	FileNameConvention string   // The naming convention for inspection report exports
	RetryTimeout       int
}

// ConfigurationOptions is an approximate representation of the YAML configuration
// NOTE: while refactoring, it is a partial representation. We will add fields when needed
type ConfigurationOptions struct {
	ApiConfig       *ApiConfig
	SheqsyApiConfig *SheqsyApiConfig
	DBConfig        *DBConfig
	CSVConfig       *CSVConfig
	ExportConfig    *ExportConfig
	ReportConfig    *ReportConfig
}
