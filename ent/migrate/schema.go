// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CredentialsColumns holds the columns for the "credentials" table.
	CredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "principal", Type: field.TypeString},
		{Name: "secret", Type: field.TypeString, Size: 3000},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"PASSWORD", "KEY", "CERTIFICATE"}},
		{Name: "target_credentials", Type: field.TypeInt, Nullable: true},
	}
	// CredentialsTable holds the schema information for the "credentials" table.
	CredentialsTable = &schema.Table{
		Name:       "credentials",
		Columns:    CredentialsColumns,
		PrimaryKey: []*schema.Column{CredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "credentials_targets_credentials",
				Columns:    []*schema.Column{CredentialsColumns[4]},
				RefColumns: []*schema.Column{TargetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "size", Type: field.TypeInt, Default: 0},
		{Name: "hash", Type: field.TypeString, Size: 100},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeBytes},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// ImplantsColumns holds the columns for the "implants" table.
	ImplantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "session_id", Type: field.TypeString},
		{Name: "process_name", Type: field.TypeString, Nullable: true},
		{Name: "implant_target", Type: field.TypeInt, Nullable: true},
		{Name: "implant_config", Type: field.TypeInt, Nullable: true},
	}
	// ImplantsTable holds the schema information for the "implants" table.
	ImplantsTable = &schema.Table{
		Name:       "implants",
		Columns:    ImplantsColumns,
		PrimaryKey: []*schema.Column{ImplantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "implants_targets_target",
				Columns:    []*schema.Column{ImplantsColumns[3]},
				RefColumns: []*schema.Column{TargetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "implants_implant_configs_config",
				Columns:    []*schema.Column{ImplantsColumns[4]},
				RefColumns: []*schema.Column{ImplantConfigsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImplantCallbackConfigsColumns holds the columns for the "implant_callback_configs" table.
	ImplantCallbackConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uri", Type: field.TypeString},
		{Name: "priority", Type: field.TypeInt, Default: 10},
		{Name: "timeout", Type: field.TypeInt, Default: 5},
		{Name: "interval", Type: field.TypeInt, Default: 60},
		{Name: "jitter", Type: field.TypeInt, Default: 0},
	}
	// ImplantCallbackConfigsTable holds the schema information for the "implant_callback_configs" table.
	ImplantCallbackConfigsTable = &schema.Table{
		Name:       "implant_callback_configs",
		Columns:    ImplantCallbackConfigsColumns,
		PrimaryKey: []*schema.Column{ImplantCallbackConfigsColumns[0]},
	}
	// ImplantConfigsColumns holds the columns for the "implant_configs" table.
	ImplantConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "auth_token", Type: field.TypeString, Unique: true},
	}
	// ImplantConfigsTable holds the schema information for the "implant_configs" table.
	ImplantConfigsTable = &schema.Table{
		Name:       "implant_configs",
		Columns:    ImplantConfigsColumns,
		PrimaryKey: []*schema.Column{ImplantConfigsColumns[0]},
	}
	// ImplantServiceConfigsColumns holds the columns for the "implant_service_configs" table.
	ImplantServiceConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: "System Service"},
		{Name: "executable_path", Type: field.TypeString},
	}
	// ImplantServiceConfigsTable holds the schema information for the "implant_service_configs" table.
	ImplantServiceConfigsTable = &schema.Table{
		Name:       "implant_service_configs",
		Columns:    ImplantServiceConfigsColumns,
		PrimaryKey: []*schema.Column{ImplantServiceConfigsColumns[0]},
	}
	// TargetsColumns holds the columns for the "targets" table.
	TargetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "forward_connect_ip", Type: field.TypeString, Unique: true},
	}
	// TargetsTable holds the schema information for the "targets" table.
	TargetsTable = &schema.Table{
		Name:       "targets",
		Columns:    TargetsColumns,
		PrimaryKey: []*schema.Column{TargetsColumns[0]},
	}
	// ImplantConfigServiceConfigsColumns holds the columns for the "implant_config_serviceConfigs" table.
	ImplantConfigServiceConfigsColumns = []*schema.Column{
		{Name: "implant_config_id", Type: field.TypeInt},
		{Name: "implant_service_config_id", Type: field.TypeInt},
	}
	// ImplantConfigServiceConfigsTable holds the schema information for the "implant_config_serviceConfigs" table.
	ImplantConfigServiceConfigsTable = &schema.Table{
		Name:       "implant_config_serviceConfigs",
		Columns:    ImplantConfigServiceConfigsColumns,
		PrimaryKey: []*schema.Column{ImplantConfigServiceConfigsColumns[0], ImplantConfigServiceConfigsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "implant_config_serviceConfigs_implant_config_id",
				Columns:    []*schema.Column{ImplantConfigServiceConfigsColumns[0]},
				RefColumns: []*schema.Column{ImplantConfigsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "implant_config_serviceConfigs_implant_service_config_id",
				Columns:    []*schema.Column{ImplantConfigServiceConfigsColumns[1]},
				RefColumns: []*schema.Column{ImplantServiceConfigsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ImplantConfigCallbackConfigsColumns holds the columns for the "implant_config_callbackConfigs" table.
	ImplantConfigCallbackConfigsColumns = []*schema.Column{
		{Name: "implant_config_id", Type: field.TypeInt},
		{Name: "implant_callback_config_id", Type: field.TypeInt},
	}
	// ImplantConfigCallbackConfigsTable holds the schema information for the "implant_config_callbackConfigs" table.
	ImplantConfigCallbackConfigsTable = &schema.Table{
		Name:       "implant_config_callbackConfigs",
		Columns:    ImplantConfigCallbackConfigsColumns,
		PrimaryKey: []*schema.Column{ImplantConfigCallbackConfigsColumns[0], ImplantConfigCallbackConfigsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "implant_config_callbackConfigs_implant_config_id",
				Columns:    []*schema.Column{ImplantConfigCallbackConfigsColumns[0]},
				RefColumns: []*schema.Column{ImplantConfigsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "implant_config_callbackConfigs_implant_callback_config_id",
				Columns:    []*schema.Column{ImplantConfigCallbackConfigsColumns[1]},
				RefColumns: []*schema.Column{ImplantCallbackConfigsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CredentialsTable,
		FilesTable,
		ImplantsTable,
		ImplantCallbackConfigsTable,
		ImplantConfigsTable,
		ImplantServiceConfigsTable,
		TargetsTable,
		ImplantConfigServiceConfigsTable,
		ImplantConfigCallbackConfigsTable,
	}
)

func init() {
	CredentialsTable.ForeignKeys[0].RefTable = TargetsTable
	ImplantsTable.ForeignKeys[0].RefTable = TargetsTable
	ImplantsTable.ForeignKeys[1].RefTable = ImplantConfigsTable
	ImplantConfigServiceConfigsTable.ForeignKeys[0].RefTable = ImplantConfigsTable
	ImplantConfigServiceConfigsTable.ForeignKeys[1].RefTable = ImplantServiceConfigsTable
	ImplantConfigCallbackConfigsTable.ForeignKeys[0].RefTable = ImplantConfigsTable
	ImplantConfigCallbackConfigsTable.ForeignKeys[1].RefTable = ImplantCallbackConfigsTable
}
