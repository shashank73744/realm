// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BeaconsColumns holds the columns for the "beacons" table.
	BeaconsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "principal", Type: field.TypeString, Nullable: true},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "agent_identifier", Type: field.TypeString, Nullable: true},
		{Name: "last_seen_at", Type: field.TypeTime, Nullable: true},
		{Name: "beacon_host", Type: field.TypeInt},
	}
	// BeaconsTable holds the schema information for the "beacons" table.
	BeaconsTable = &schema.Table{
		Name:       "beacons",
		Columns:    BeaconsColumns,
		PrimaryKey: []*schema.Column{BeaconsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "beacons_hosts_host",
				Columns:    []*schema.Column{BeaconsColumns[6]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "size", Type: field.TypeInt, Default: 0},
		{Name: "hash", Type: field.TypeString, Size: 100},
		{Name: "content", Type: field.TypeBytes},
		{Name: "tome_files", Type: field.TypeInt, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_tomes_files",
				Columns:    []*schema.Column{FilesColumns[7]},
				RefColumns: []*schema.Column{TomesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HostsColumns holds the columns for the "hosts" table.
	HostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "primary_ip", Type: field.TypeString, Nullable: true},
		{Name: "platform", Type: field.TypeEnum, Enums: []string{"Windows", "Linux", "MacOS", "BSD", "Unknown"}, Default: "Unknown"},
		{Name: "last_seen_at", Type: field.TypeTime, Nullable: true},
	}
	// HostsTable holds the schema information for the "hosts" table.
	HostsTable = &schema.Table{
		Name:       "hosts",
		Columns:    HostsColumns,
		PrimaryKey: []*schema.Column{HostsColumns[0]},
	}
	// QuestsColumns holds the columns for the "quests" table.
	QuestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "parameters", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
		{Name: "quest_tome", Type: field.TypeInt},
		{Name: "quest_bundle", Type: field.TypeInt, Nullable: true},
		{Name: "quest_creator", Type: field.TypeInt, Nullable: true},
	}
	// QuestsTable holds the schema information for the "quests" table.
	QuestsTable = &schema.Table{
		Name:       "quests",
		Columns:    QuestsColumns,
		PrimaryKey: []*schema.Column{QuestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "quests_tomes_tome",
				Columns:    []*schema.Column{QuestsColumns[5]},
				RefColumns: []*schema.Column{TomesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "quests_files_bundle",
				Columns:    []*schema.Column{QuestsColumns[6]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "quests_users_creator",
				Columns:    []*schema.Column{QuestsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"group", "service"}},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "claimed_at", Type: field.TypeTime, Nullable: true},
		{Name: "exec_started_at", Type: field.TypeTime, Nullable: true},
		{Name: "exec_finished_at", Type: field.TypeTime, Nullable: true},
		{Name: "output", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "quest_tasks", Type: field.TypeInt},
		{Name: "task_beacon", Type: field.TypeInt},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_quests_tasks",
				Columns:    []*schema.Column{TasksColumns[8]},
				RefColumns: []*schema.Column{QuestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tasks_beacons_beacon",
				Columns:    []*schema.Column{TasksColumns[9]},
				RefColumns: []*schema.Column{BeaconsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TomesColumns holds the columns for the "tomes" table.
	TomesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "param_defs", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
		{Name: "hash", Type: field.TypeString, Size: 100},
		{Name: "eldritch", Type: field.TypeString, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
	}
	// TomesTable holds the schema information for the "tomes" table.
	TomesTable = &schema.Table{
		Name:       "tomes",
		Columns:    TomesColumns,
		PrimaryKey: []*schema.Column{TomesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 25},
		{Name: "oauth_id", Type: field.TypeString, Unique: true},
		{Name: "photo_url", Type: field.TypeString, SchemaType: map[string]string{"mysql": "MEDIUMTEXT"}},
		{Name: "session_token", Type: field.TypeString, Size: 200},
		{Name: "is_activated", Type: field.TypeBool, Default: false},
		{Name: "is_admin", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// HostTagsColumns holds the columns for the "host_tags" table.
	HostTagsColumns = []*schema.Column{
		{Name: "host_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// HostTagsTable holds the schema information for the "host_tags" table.
	HostTagsTable = &schema.Table{
		Name:       "host_tags",
		Columns:    HostTagsColumns,
		PrimaryKey: []*schema.Column{HostTagsColumns[0], HostTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "host_tags_host_id",
				Columns:    []*schema.Column{HostTagsColumns[0]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "host_tags_tag_id",
				Columns:    []*schema.Column{HostTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BeaconsTable,
		FilesTable,
		HostsTable,
		QuestsTable,
		TagsTable,
		TasksTable,
		TomesTable,
		UsersTable,
		HostTagsTable,
	}
)

func init() {
	BeaconsTable.ForeignKeys[0].RefTable = HostsTable
	FilesTable.ForeignKeys[0].RefTable = TomesTable
	QuestsTable.ForeignKeys[0].RefTable = TomesTable
	QuestsTable.ForeignKeys[1].RefTable = FilesTable
	QuestsTable.ForeignKeys[2].RefTable = UsersTable
	TasksTable.ForeignKeys[0].RefTable = QuestsTable
	TasksTable.ForeignKeys[1].RefTable = BeaconsTable
	HostTagsTable.ForeignKeys[0].RefTable = HostsTable
	HostTagsTable.ForeignKeys[1].RefTable = TagsTable
}
